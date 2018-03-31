CONTAINER=go-app
VERSION=`git describe --tags`
MAIN_FILES=main.go utils.go handlers.go
LDFLAGS=-ldflags "-X main.Version=${VERSION}"

init:
	go get -d
	go get github.com/codegangsta/gin

# Run the server and reload on file saves aka livereload
watch: 
	gin -a 8080 -i run $(MAIN_FILES) 

run:
	go run $(LDFLAGS) $(MAIN_FILES)

build:
	go build $(LDFLAGS) -o app .

build/container:
	go get -d
	CGO_ENABLED=0 GOOS=linux go build $(LDFLAGS) -a -installsuffix cgo -o app .
	docker build -t $(CONTAINER) .

run/container:
	docker run --publish 8080:8080 --name $(CONTAINER) --rm $(CONTAINER)