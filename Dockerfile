FROM golang:1.20-alpine3.17 as builder

WORKDIR /app

COPY ../../go.mod .
COPY ../../go.sum .

RUN go mod download

COPY ../.. .

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o taskApp ./cmd/app/main.go


FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root

COPY --from=builder /app/taskApp ./
COPY --from=builder /app/.env ./

CMD ["./taskApp"]


