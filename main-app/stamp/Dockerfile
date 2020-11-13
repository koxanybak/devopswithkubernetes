FROM golang:latest as builder

ENV G0111MODULE=on

WORKDIR /stamp

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o stamp

FROM ubuntu:latest
COPY --from=builder /stamp/ /app/stamp/
WORKDIR /app/stamp
RUN mkdir ../files && touch ../files/stamp.txt
EXPOSE 8000
ENTRYPOINT [ "./stamp" ]