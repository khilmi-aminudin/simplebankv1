# FROM golang:1.18-alpine3.15
# WORKDIR /app
# COPY main .
# CMD [ "/app/main" ]

FROM golang:1.18-alpine3.15

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-gs-ping

EXPOSE 8080

CMD [ "/docker-gs-ping" ]
