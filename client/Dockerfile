FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache protobuf-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

RUN protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/greeting.proto

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o client .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN addgroup -g 10014 choreo && \
    adduser -D -s /bin/sh -u 10014 -G choreo choreouser

WORKDIR /home/choreouser

COPY --from=builder /app/client .

RUN chown -R 10014:10014 /home/choreouser

USER 10014

EXPOSE 8080

CMD ["./client"]