FROM scratch AS license
COPY NOTICE /NOTICE
COPY LICENSE /LICENSE

FROM busybox AS bin
COPY ./dist /binaries
RUN if [[ "$(arch)" == "x86_64" ]]; then \
        architecture="amd64"; \
    else \
        architecture="arm64"; \
    fi; \
    cp /binaries/deterministic-zip_linux-${architecture} /bin/deterministic-zip && \
    chmod +x /bin/deterministic-zip && \
    chown 1000:1000 /bin/deterministic-zip

FROM alpine
ARG BUILD_TIME \
    BUILD_VERSION \
    BUILD_COMMIT_REF
LABEL org.opencontainers.image.title="deterministic-zip" \
      org.opencontainers.image.description="Simple (almost drop-in) replacement for zip that produces deterministic files." \
      org.opencontainers.image.ref.name="master" \
      org.opencontainers.image.licenses='GPL-3.0' \
      org.opencontainers.image.vendor="Timo Reymann <mail@timo-reymann.de>" \
      org.opencontainers.image.authors="Timo Reymann <mail@timo-reymann.de>" \
      org.opencontainers.image.url="https://github.com/timo-reymann/deterministic-zip" \
      org.opencontainers.image.documentation="https://github.com/timo-reymann/deterministic-zip" \
      org.opencontainers.image.source="https://github.com/timo-reymann/deterministic-zip.git" \
      org.opencontainers.image.created=$BUILD_TIME \
      org.opencontainers.image.version=$BUILD_VERSION \
      org.opencontainers.image.revision=$BUILD_COMMIT_REF

COPY --from=license / /

COPY --from=bin /bin/deterministic-zip /bin/deterministic-zip

WORKDIR /workspace
ENTRYPOINT [ "/bin/deterministic-zip" ]
