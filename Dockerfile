FROM golang:alpine AS BUILDER

# Set the Current Working Directory inside the container
WORKDIR /app

# Install basic packages
RUN apk add \
    git gcc g++

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go mod download

# Build image
RUN go build -o ./app .

FROM alpine:latest AS RUNNER

WORKDIR /app

COPY --from=BUILDER /app/app /app/app

RUN chmod +x /app/app
RUN ln -s /app/app /usr/local/bin/app

# Run the executable
CMD ["app"]
