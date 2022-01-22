FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

ADD ./backend .

RUN go mod tidy

RUN go build -o /tototodo

EXPOSE 8000

CMD [ "/tototodo" ]