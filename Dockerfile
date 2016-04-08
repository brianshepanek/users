FROM golang:1.6.0
COPY . /go/src/github.com/brianshepanek/users
RUN go get -t ./src/github.com/brianshepanek/users
CMD ["users"]