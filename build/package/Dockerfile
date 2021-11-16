FROM golang:1.17-alpine

ARG TARGETPLATFORM=amd64

RUN addgroup -S identicons && \
    adduser -S identicons -G identicons

RUN mkdir /executables /licenses
COPY identicons_* /executables/
COPY licenses/* /licenses/

RUN apk add --no-cache sed && \
    which sed && \
    SUFFIX=$(echo $TARGETPLATFORM | sed -E 's/^[^\/]+\/([^\/]+)(\/.*)?/\1/') && \
    cp /executables/identicons_$SUFFIX /identicons && \
    rm -r /executables/ && \
    apk del sed

EXPOSE 8080
USER identicons
ENTRYPOINT ["/identicons"]
CMD ["serve"]
