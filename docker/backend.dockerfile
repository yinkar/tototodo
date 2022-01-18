FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

ADD ./backend .

COPY ./backend/go.mod ./backend/go.sum ./

RUN go mod download

WORKDIR /app/src

RUN go mod download

RUN go build todolist.go

EXPOSE 8000

CMD [ "./todolist" ]