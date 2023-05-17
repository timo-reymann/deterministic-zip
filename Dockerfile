FROM busybox as bin
COPY ./dist /binaries
RUN if [[ "$(arch)" == "x86_64" ]]; then \
        architecture="amd64"; \
    else \
        architecture="arm64"; \
    fi; \
    cp /binaries/deterministic-zip_linux-${architecture} /bin/deterministic-zip && \
    chmod +x /bin/deterministic-zip && \
    chown 1000:1000 /bin/deterministic-zip

FROM scratch
COPY --from=bin /bin/deterministic-zip /deterministic-zip
ENTRYPOINT [ "/deterministic-zip" ]
