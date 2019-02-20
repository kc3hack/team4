FROM golang:1.11.5

RUN go get github.com/gin-gonic/gin
RUN go get googlemaps.github.io/maps
RUN go get github.com/kr/pretty

WORKDIR /app
ADD . /app

CMD ["go", "run", "server.go"]