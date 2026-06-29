FROM golang:1.27rc1-alpine3.24

LABEL version="1.0"

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

EXPOSE 8080

RUN go build -o ./ascii-art-web

CMD [ "./ascii-art-web" ]