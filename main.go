// TODO: webserver compliance (nginx/apache)
// TODO: oauthStateString must be regenerate for each request
// TODO: the Username/IconURL seems broken

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

const alphaNum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

type mattermostRequest struct {
	ChannelId   string `form:"channel_id"`
	ChannelName string `json:"channel_name"`
	Command     string `json:"command"`
	ResponseURL string `json:"response_url"`
	TeamDomain  string `json:"team_domain"`
	TeamID      string `json:"team_id"`
	Text        string `json:"text"`
	Token       string `json:"token"`
	UserId      string `json:"user_id"`
	UserName    string `json:"user_name"`
}

type mattermostResponse struct {
	Username     string `json:"username"`
	IconURL      string `json:"icon_url"`
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
}

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "", // redirect URL will be update during the first oAuth step (fishy way)
		ClientID:     os.Getenv("ClientID"),
		ClientSecret: os.Getenv("ClientSecret"),
		Scopes:       []string{"https://www.googleapis.com/auth/calendar"},
		Endpoint:     google.Endpoint,
	}
	oauthStateString = randomStr()
)

// randomStr generate random string for the google Oauth part.
func randomStr() string {
	var bytes = make([]byte, 10)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphaNum[b%byte(len(alphaNum))]
	}
	return string(bytes)
}

// sanityCheck check all mandatory variables.
func sanityCheck() {
	langoustePort := os.Getenv("langoustePort")
	ClientID := os.Getenv("ClientID")
	ClientSecret := os.Getenv("ClientSecret")

	if langoustePort == "" {
		log.Fatal("You must specify a langoustePort.")
	}

	if ClientID == "" {
		log.Fatal("You must specify a google ClientID.")
	}

	if ClientSecret == "" {
		log.Fatal("You must specify a google ClientSecret.")
	}
}

// getClient uses a Context and Config to retrieve a Token
// then generate a Client. It returns the generated Client.
func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
	cacheFile, err := tokenCacheFile()
	if err != nil {
		log.Fatalf("Unable to get path to cached credential file. %v", err)
	}
	tok, err := tokenFromFile(cacheFile)
	if err != nil {
		log.Fatalf("Unable to get the token from file. %v", err)
	}
	return config.Client(ctx, tok)
}

// tokenCacheFile generates credential file path/filename.
// It returns the generated credential path/filename.
func tokenCacheFile() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	tokenCacheDir := filepath.Join(usr.HomeDir, ".credentials")
	os.MkdirAll(tokenCacheDir, 0700)
	return filepath.Join(tokenCacheDir,
		url.QueryEscape("calendar-langouste.json")), err
}

// tokenFromFile retrieves a Token from a given file path.
// It returns the retrieved Token and any read error encountered.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}

// saveToken uses a file path to create a file and store the
// token in it.
func saveToken(file string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.Create(file)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

// insert the calendar event
func insertCalendarEvent(r *http.Request) (calendar.Event, error) {
	now := time.Now()
	hangout_event := &calendar.Event{
		Summary:     fmt.Sprintf("Mattermost - hangout bot - %v", r.FormValue("user_name")),
		Description: "Event automatically create by Langouste.",
		Start: &calendar.EventDateTime{
			DateTime: now.Format(time.RFC3339),
		},
		End: &calendar.EventDateTime{
			DateTime: now.Format(time.RFC3339),
		},
		Location: "Belgrade",
	}

	ctx := context.Background()
	client := getClient(ctx, googleOauthConfig)
	srv, err := calendar.New(client)
	if err != nil {
		log.Printf("Unable to retrieve calendar %v", err)
		return *hangout_event, err
	}

	event, err := srv.Events.Insert("primary", hangout_event).Do()
	if err != nil {
		log.Printf("Unable to create event. %v\n", err)
		return *hangout_event, err
	}
	fmt.Printf("Event created: %s\n", event.HangoutLink)
	return *event, err
}

// loginHandler create a login link and redirects the user to it.
func loginHandler(w http.ResponseWriter, r *http.Request) {
	var scheme string
	scheme = "http"
	if r.TLS != nil {
		scheme = "https"
	}

	env_host := os.Getenv("langousteHost")
	googleOauthConfig.RedirectURL = fmt.Sprintf("%v://%v/callback", scheme, env_host)
	if env_host == "" {
		googleOauthConfig.RedirectURL = fmt.Sprintf("%v://%v/callback", scheme, r.Host)
	}
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// callbackHandler verify the oauth part
// We create the credential cache file also.
func callbackHandler(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	cacheFile, err := tokenCacheFile()
	if err != nil {
		log.Fatalf("Unable to get path to cached credential file. %v", err)
	}
	saveToken(cacheFile, token)
	fmt.Fprintf(w, "Sucessfully save credential file to: %s\n", cacheFile)
}

// eventHandler handle the mattermost slash command and
// create a new calendar event.
func eventHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var response mattermostResponse

	response.IconURL = ""
	response.Username = "Langouste"
	response.ResponseType = "in_channel"

	if r.Body == nil {
		response.Text = "plase send a request body"
		data, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		return
	}

	event, err := insertCalendarEvent(r)
	if err != nil {
		response.Text = err.Error()
		data, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		return
	}

	response.Text = fmt.Sprintf("Hangout link by %v\n\n --- \n\n %v", r.FormValue("user_name"), event.HangoutLink)
	data, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	return
}

func main() {
	sanityCheck()

	port := os.Getenv("langoustePort")
	r := mux.NewRouter()
	r.HandleFunc("/", loginHandler).Methods("GET")
	r.HandleFunc("/callback", callbackHandler).Methods("GET")
	r.HandleFunc("/", eventHandler).Methods("POST")

	loggerHandler := handlers.LoggingHandler(os.Stdout, r)
	errHTTP := http.ListenAndServe(fmt.Sprintf(":%v", port), loggerHandler)
	if errHTTP != nil {
		log.Fatal("ListenAndServe: ", errHTTP)
	}
	log.Printf("ListenAndServe - 0.0.0.0:%v", port)
}
