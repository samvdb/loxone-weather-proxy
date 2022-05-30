FROM alpine:3.9
RUN apk --no-cache --no-progress add ca-certificates
# Adding the grpc_health_probe
RUN GRPC_HEALTH_PROBE_VERSION=v0.3.2 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe
RUN mkdir /uploads
ADD dist/go_linux_amd64/server /server
ENTRYPOINT ["server"]
CMD ["start"]
