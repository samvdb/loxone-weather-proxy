FROM scratch
ENV DARKSKY_APIKEY=""
RUN mkdir /app

# Run the executable
CMD ["/app/proxy"]
# Perform any further action as an unprivileged user.
USER nobody:nobody
COPY proxy /app/proxy


