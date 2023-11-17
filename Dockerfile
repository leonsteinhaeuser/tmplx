FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY tmplx /tmplx
USER 65532:65532
ENTRYPOINT ["/tmplx"]
