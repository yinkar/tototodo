FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

ADD ./backend .

COPY ./backend/go.mod ./backend/go.sum ./

RUN go mod tidy

RUN go build .

EXPOSE 8000

CMD [ "./backend" ]