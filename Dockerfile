FROM golang:1.16.0-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./ .

RUN go build -o main .

RUN go get

RUN go install


EXPOSE 8080

CMD ["./main"]