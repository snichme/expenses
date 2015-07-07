default: test

build:
	CGO_ENABLED=0 go build -ldflags "-s" -a -installsuffix cgo -o main .

run:
	go run *.go

test:
	go test -timeout=5s ./...

clean:
	rm main
