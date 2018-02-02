FROM golang:1.9.0 as builder

WORKDIR /go/src/github.com/rahul-microservice/user-cli

COPY . .

RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build  -o user-cli -a -installsuffix cgo cli.go


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY user.json /app/user.json
COPY --from=builder /go/src/github.com/rahul-microservice/user-cli/user-cli .

ENTRYPOINT ["./user-cli", "user.json"]
CMD ["./user-cli"]