FROM docker.io/golang:1.24.3-bookworm

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o server

CMD [ "./server" ]
