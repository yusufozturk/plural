MAIN_PACKAGE := plural
BUILT_ON := $(date)
COMMIT_HASH := $(git log -n 1 --pretty=format:"%H")
GO_LINUX := GOOS=linux GOARCH=amd64
GO_OSX := GOOS=darwin GOARCH=amd64
LDFLAGS := '-X "main.builtOn=$(BUILT_ON)" -X "main.commitHash=$(COMMIT_HASH)" -s -w'

test:
	go test -p=1 -cover `go list ./... | sed -n '1!p' | grep -v /vendor/` -v

osx:
	$(GO_OSX) go build -o $(MAIN_PACKAGE) -ldflags $(LDFLAGS) .

linux:
	$(GO_LINUX) go build -o $(MAIN_PACKAGE) -ldflags $(LDFLAGS) .

clean:
	find . -name *_gen.go -type f -exec rm {} \;
	rm -f ./$(MAIN_PACKAGE)

run:
	go run -ldflags $(LDFLAGS) main.go

fmt:
	go fmt ./...
