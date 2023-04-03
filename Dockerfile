FROM golang:1.19.4-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

#instal psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

#build go app
RUN go mod download
RUN go build -o rest-api-shop ./cmd/main.go

CMD ["./rest-api-shop"]