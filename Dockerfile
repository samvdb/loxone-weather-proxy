FROM golang:1.14-alpine
ENV DARKSKY_APIKEY=""
RUN apk --update upgrade
RUN apk add --update gcc musl-dev


# removing apk cache
RUN rm -rf /var/cache/apk/*

RUN mkdir -p /app
COPY . /app/
WORKDIR /app

RUN GOARCH=amd64 CGO_ENABLED=1 go build  -a -ldflags "-linkmode external -extldflags -static"    -o /app/proxy .
# Perform any further action as an unprivileged user.
USER nobody:nobody
# Run the executable
CMD ["/app/proxy"]