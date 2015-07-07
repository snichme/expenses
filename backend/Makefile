default: test

setup:
	export GOPATH=$(HOME)/go
	go get -u github.com/codegangsta/negroni
	go get -u github.com/gorilla/mux
	go get -u github.com/satori/go.uuid
	go get -u gopkg.in/redis.v2

build:
	CGO_ENABLED=0 go build -ldflags "-s" -a -installsuffix cgo -o main .

run:
	go run *.go

test:
	go test -timeout=5s ./...

clean:
	rm main
