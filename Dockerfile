FROM alpine:latest
ENV DARKSKY_APIKEY=""
RUN apk add --no-cache  bash \
                        curl
RUN mkdir /app

# Run the executable
CMD ["/bin/proxy"]
# Perform any further action as an unprivileged user.
USER nobody:nobody
COPY dist/loxone-weather-proxy_linux_amd64/proxy /app/proxy


