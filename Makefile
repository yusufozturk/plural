MAIN_PACKAGE := plural
BUILT_ON := $(shell date)
COMMIT_HASH := $(shell git log -n 1 --pretty=format:"%H")
GO_LINUX := GOOS=linux GOARCH=amd64
GO_OSX := GOOS=darwin GOARCH=amd64
LDFLAGS := '-X "main.builtOn=$(BUILT_ON)" -X "main.commitHash=$(COMMIT_HASH)"'

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
