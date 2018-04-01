FROM golang:1.8-alpine

WORKDIR $GOPATH/src/github.com/zerdzhong/coming_movie_server
ADD . $GOPATH/src//github.com/zerdzhong/coming_movie_server
RUN go build .

ENTRYPOINT ["./coming_movie_server"]

EXPOSE 8080

