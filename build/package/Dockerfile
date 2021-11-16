FROM golang:1.17-alpine

RUN addgroup -S identicons && \
    adduser -S identicons -G identicons

RUN mkdir /licenses
COPY identicons /identicons
COPY licenses/* /licenses/

EXPOSE 8080
USER identicons
ENTRYPOINT ["/identicons"]
CMD ["serve"]
