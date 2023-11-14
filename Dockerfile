# Stage 1: Build the Go application
FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go application
RUN go get modernc.org/sqlite
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o todoapp cmd/todoapp/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o migration cmd/migration/main.go

# Stage 2: Create a minimal image to run the application
FROM alpine:latest

WORKDIR /app
RUN apk --no-cache add curl

# Copy only the necessary files from the builder stage
COPY --from=builder /app/todoapp .
COPY --from=builder /app/migration .
COPY --from=builder /app/.env .
COPY --from=builder /app/web/template/index.html ./web/template/index.html

# Expose any necessary ports
EXPOSE 8018

# Command to run the application
CMD /app/migration up && /app/todoapp
