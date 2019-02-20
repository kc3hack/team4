FROM golang:1.11.5

RUN go get github.com/gin-gonic/gin

WORKDIR /go/src/app
# ADD . .

CMD ["go", "run", "server.go"]