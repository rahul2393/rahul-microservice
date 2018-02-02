FROM golang:1.9.0 as builder

WORKDIR /go/src/github.com/rahul-microservice/user-service

COPY . .

RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build  -o user-service -a -installsuffix cgo main.go repository.go handler.go database.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/rahul-microservice/user-service/user-service .

CMD ["./user-service"]