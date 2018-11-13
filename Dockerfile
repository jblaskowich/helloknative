# Start from golang:alpine with the latest version of Go installed
# and use it as a build environment
FROM golang:alpine AS builder

# Copy the local code to the container
ADD helloknative/ /go/src/helloknative

# Change the current work directory to build source code
WORKDIR /go/src/helloknative
RUN go build 

# Transfert the builded binary to alpine
# in order to have the smallest footprint
FROM alpine
COPY --from=builder /go/src/helloknative/helloknative /bin/.

# Run the helloknative command when the container starts
ENTRYPOINT ["/bin/helloknative"]

# Document that the service listens on port 8080
EXPOSE 8080