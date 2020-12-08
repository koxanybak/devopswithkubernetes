FROM golang:latest as builder

ENV GO111MODULE=on

WORKDIR /bc

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bc

FROM ubuntu:latest
RUN apt-get update && apt-get install -y ca-certificates
COPY --from=builder /bc/ /bc/
WORKDIR /bc
ENTRYPOINT [ "./bc" ]