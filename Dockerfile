FROM node:14.7.0-buster as reactbuilder

WORKDIR /src

COPY ui/package*.json /src/

RUN npm install

COPY ui/ /src/

RUN npm run build

FROM golang:1.14.6-buster as gobuilder

WORKDIR /go/src/github.com/kwlo/gravity

COPY . .

RUN go test ./...

RUN go build github.com/kwlo/gravity

FROM golang:1.14.6-buster

EXPOSE 8080

WORKDIR /app

COPY --from=reactbuilder /src/build /app/static

COPY --from=gobuilder /go/src/github.com/kwlo/gravity/gravity /app

CMD ["./gravity"]
