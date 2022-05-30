FROM alpine
ENV DARKSKY_APIKEY=""
RUN mkdir /app

# Run the executable
# Perform any further action as an unprivileged user.
COPY proxy /app/proxy
USER nobody:nobody
CMD ["/app/proxy"]


