FROM golang:alpine AS builder

LABEL maintainer="Alexandre Awadallak <alexandre.awadallak@gmail.com>"

WORKDIR /app/telegram-bot

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/main cmd/main.go

######## Start a new stage from scratch #######
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/telegram-bot/build/main .

# Expose port 6728 to the outside world
EXPOSE 6728

# Command to run the executable
CMD ["./main"] 




