FROM alpine:3.9
ENV DARKSKY_APIKEY=""
COPY proxy-linux-amd64 /usr/bin/loxone
USER nobody:nobody
ENTRYPOINT ["loxone"]


