FROM golang:1.6.0
COPY . /go/src/app
RUN go install app
CMD ["app"]