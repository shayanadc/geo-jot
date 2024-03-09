FROM golang:alpine

WORKDIR /go/src/app

COPY . .

RUN go get -u

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]