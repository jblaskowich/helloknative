# Start from golang:alpine with the latest version of Go installed
# and use it as a build environment
FROM golang:alpine AS builder

WORKDIR /go/src/github.com/jblaskowich/helloknative
# Copy the local code to the container
ADD . .

# build du code source
RUN CGO_ENABLED=0 GOOS=linux go build -v -o helloknative

# Transfert the builded binary to alpine
# in order to have the smallest footprint
FROM scratch
COPY --from=builder github.com/jblaskowich/helloknative/helloknative .

# Run the helloknative command when the container starts
ENTRYPOINT ["./helloknative"]

# Document that the service listens on port 8080
EXPOSE 8080
