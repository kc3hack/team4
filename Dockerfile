FROM golang:1.11.5

RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-xorm/xorm
RUN go get github.com/go-sql-driver/mysql

WORKDIR /go/src/app

CMD ["go", "run", "server.go"]