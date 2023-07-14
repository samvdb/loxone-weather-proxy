FROM alpine:3.9
ENV TOMORROW_APIKEY=""
COPY proxy-linux-amd64 /usr/bin/loxone
RUN apk add libcap && setcap 'cap_net_bind_service=+ep' /usr/bin/loxone
#USER nobody:nobody
ENTRYPOINT ["loxone"]


