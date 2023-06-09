FROM golang:latest as builder

# Set the Current Working Directory inside the container
RUN mkdir /app

WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . ./

# Build the Go app
RUN CGO_ENABLED=0 go build -o main ./cmd/main.go

######## Start a new stage from scratch #######
FROM alpine:latest  

RUN apk --no-cache add ca-certificates curl

WORKDIR /app

# Create user and set ownership and permissions as required
RUN adduser -D xmgateway && chown -R xmgateway /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app .

USER xmgateway

# Expose port 8080 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["./main"] 

