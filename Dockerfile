FROM golang:1.7

# Create go src directory for this app
RUN mkdir -p /go/src/github.com/politicalrev/accountability-api
ADD . /go/src/github.com/politicalrev/accountability-api/
WORKDIR /go/src/github.com/politicalrev/accountability-api

# Download all dependencies
RUN go get -v -d ./...

# Download tools
RUN go get -v github.com/codegangsta/gin
RUN go get -v github.com/CloudCom/goose/cmd/goose

EXPOSE "4000"
ENTRYPOINT ["go", "run", "app.go"]
