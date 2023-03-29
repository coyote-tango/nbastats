FROM golang:1.17-alpine as builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o nbastats ./cmd/nbastats

FROM alpine:latest

# Install ca-certificates to enable HTTPS and tzdata for time zones
RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /app/nbastats /app/nbastats

EXPOSE 8080

ENTRYPOINT ["/app/nbastats"]
