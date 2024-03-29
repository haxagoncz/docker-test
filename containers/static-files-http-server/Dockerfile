FROM alpine:latest

RUN apk update

# Install thttpd
RUN apk add --no-cache thttpd

# Create a non-root user to own the files and run our server
RUN adduser -D static
USER static
WORKDIR /home/static

# Run thttpd
CMD ["thttpd", "-D", "-h", "0.0.0.0", "-p", "8080", "-d", "/home/static", "-u", "static", "-l", "/tmp/logs", "-M", "60"]
