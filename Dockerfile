FROM golang:1.14.6-buster

EXPOSE 8080

WORKDIR /go/src/github.com/kwlo/gravity

COPY . .

RUN go test ./...

RUN go install github.com/kwlo/gravity

CMD ["gravity"]
