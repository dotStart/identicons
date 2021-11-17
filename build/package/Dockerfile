FROM golang:1.17

ARG TARGETPLATFORM=amd64

RUN useradd -U identicons

RUN mkdir /executables /licenses
COPY identicons_* /executables/
COPY licenses/* /licenses/

RUN SUFFIX=$(echo $TARGETPLATFORM | sed -E 's/^[^\/]+\/([^\/]+)(\/.*)?/\1/') && \
    cp /executables/identicons_$SUFFIX /identicons && \
    rm -r /executables/

EXPOSE 8080
USER identicons
ENTRYPOINT ["/identicons"]
CMD ["serve"]
