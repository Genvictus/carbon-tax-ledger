FROM golang:1.23-alpine

LABEL maintainer="Brian Kheng"

# Move to working directory (/app).
WORKDIR /app

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/air-verse/air@latest

# Copy the code into the container.
COPY . .

# Command to run when starting the container.
CMD ["air"]