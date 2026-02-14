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
LABEL 
LABEL org.opencontainers.image.licenses="gpl-3.0"org.opencontainers.image.title="deterministic-zip"
LABEL org.opencontainers.image.description="Simple (almost drop-in) replacement for zip that produces deterministic files."
LABEL org.opencontainers.image.ref.name="master"
LABEL org.opencontainers.image.licenses='"Climate Strike" License Version 1.0 (Draft)'
LABEL org.opencontainers.image.vendor="Timo Reymann <mail@timo-reymann.de>"
LABEL org.opencontainers.image.authors="Timo Reymann <mail@timo-reymann.de>"
LABEL org.opencontainers.image.url="https://github.com/timo-reymann/deterministic-zip"
LABEL org.opencontainers.image.documentation="https://github.com/timo-reymann/deterministic-zip"
LABEL org.opencontainers.image.source="https://github.com/timo-reymann/deterministic-zip.git"

ARG BUILD_TIME
ARG BUILD_VERSION
ARG BUILD_COMMIT_REF
LABEL org.opencontainers.image.created=$BUILD_TIME
LABEL org.opencontainers.image.version=$BUILD_VERSION
LABEL org.opencontainers.image.revision=$BUILD_COMMIT_REF

COPY --from=bin /bin/deterministic-zip /bin/deterministic-zip
WORKDIR /workspace
ENTRYPOINT [ "/bin/deterministic-zip" ]
