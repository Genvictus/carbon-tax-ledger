FROM golang:1.23-alpine AS builder

LABEL maintainer="Brian Kheng"

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environment variables needed for our image and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o main .

FROM alpine:3.20.3

# Set the working directory for the final image.
WORKDIR /app

# Copy binary and .env.prod from the builder.
COPY --from=builder /build/main .
COPY --from=builder /build/.env .

EXPOSE 5000

CMD ["/app/main"]