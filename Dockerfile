FROM golang:1.11.5

RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-xorm/xorm
RUN go get github.com/go-sql-driver/mysql
# RUN go get github.com/labstack/echo
# RUN go get github.com/labstack/echo/engine/standard
# RUN go get github.com/labstack/echo/middleware
# RUN go get github.com/gocraft/dbr

WORKDIR /go/src/app
# ADD . .

CMD ["go", "run", "server.go"]