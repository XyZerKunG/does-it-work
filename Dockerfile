FROM golang:1.18.2-alpine3.15

WORKDIR /

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /hi-mom

EXPOSE 3000

CMD [ "/hi-mom" ]