FROM golang:1.16.3 as builder
WORKDIR /go/src/github.com/pcherednichenko/users
ADD . .

# Swagger documentation
RUN go get github.com/swaggo/swag/cmd/swag
RUN swag init -g cmd/users/users.go

RUN CGO_ENABLED=0 GOOS=linux go build -o users ./cmd/users/users.go

FROM alpine:3.13.5
RUN apk --no-cache add ca-certificates
COPY --from=builder go/src/github.com/pcherednichenko/users/docs/ /docs/
COPY --from=builder /go/src/github.com/pcherednichenko/users/users users

CMD ["./users"]