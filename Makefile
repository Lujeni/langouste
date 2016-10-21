BINARY=langouste

VERSION=0.1
BUILD=`git rev-parse HEAD`

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

build:
	go build ${LDFLAGS} -o ${BINARY}

install:
	go install

clean:
	if [ -f ${BINARY} ]; then rm ${BINARY}; fi
