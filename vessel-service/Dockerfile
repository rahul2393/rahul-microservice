FROM golang:1.9.0 as builder

WORKDIR /go/src/github.com/rahul2393/rahul-microservice/vessel-service

COPY . .

RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build  -o vessel-service -a -installsuffix cgo main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/rahul2393/rahul-microservice/vessel-service/vessel-service .

CMD ["./vessel-service"]