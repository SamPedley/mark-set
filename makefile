CONTAINER=go-app
VERSION=`git describe --tags`
MAIN_FILES=main.go utils.go handlers.go
LDFLAGS=-ldflags "-X main.Version=${VERSION}"

watch: 
	gin run $(MAIN_FILES)

run:
	go run $(MAIN_FILES)

build:
	go build $(CONTAINER) -o app .

container/build:
	go get -d
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
	docker build -t $(CONTAINER) .

container/run:
	docker run --publish 8080:8080 --name $(CONTAINER) --rm $(CONTAINER)