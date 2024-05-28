FROM gcr.io/distroless/static:latest
LABEL maintainers="Kubernetes Authors"
LABEL description="CSI External Snapshotter Sidecar"
ARG binary=./bin/csi-snapshotter

COPY ${binary} csi-snapshotter
RUN chmod +x csi-snapshotter
ENTRYPOINT ["/csi-snapshotter"]
