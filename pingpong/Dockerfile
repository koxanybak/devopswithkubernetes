FROM golang:latest as builder

ENV GO111MODULE=on

WORKDIR /srv

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o srv

FROM ubuntu:latest
RUN apt-get update && apt-get install -y wait-for-it
COPY --from=builder /srv/ /srv/
WORKDIR /srv
RUN mkdir files
ENV PORT=8000
EXPOSE 8000
ENTRYPOINT [ "wait-for-it", "postgres-svc:5432", "--", "./srv" ]