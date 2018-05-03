FROM alpine

RUN apk add --no-cache ca-certificates

# Copy the binary file and set it as entrypoint
ADD redis-golang /
ENTRYPOINT ["/redis-golang"]

# The service listens on port 8080 by default.
EXPOSE 8080