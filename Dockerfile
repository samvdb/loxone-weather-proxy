FROM alpine
ENV DARKSKY_APIKEY=""
COPY proxy-linux-amd64 /usr/local/bin/proxy
USER nobody:nobody
ENTRYPOINT ["proxy"]


