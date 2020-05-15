FROM golang:1.14-alpine
ENV DARKSKY_APIKEY=""
RUN apk add --no-cache  bash \
                        curl
RUN mkdir /app

# Run the executable
CMD ["/bin/proxy"]
# Perform any further action as an unprivileged user.
USER nobody:nobody
COPY proxy /app/proxy


