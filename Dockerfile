FROM gcr.io/distroless/static-debian11

COPY vyconfigure /

CMD ["/vyconfigure"]
