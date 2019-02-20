FROM golang:1.11.5

RUN go get github.com/gin-gonic/gin

WORKDIR /app
ADD . /app

CMD ["go", "run", "server.go"]