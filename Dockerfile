## NOTE: This image uses goreleaser to build image
# if building manually please run go build ./cmd/gabby first and then build

# Choose alpine as a base image to make this useful for CI, as many
# CI tools expect an interactive shell inside the container
FROM alpine:latest as production

#COPY --from=builder /build/gabby /usr/bin/gabby
COPY gabby /usr/bin/gabby
RUN chmod +x /usr/bin/gabby

WORKDIR /workdir

ENTRYPOINT ["/usr/bin/gabby"]
CMD ["--help"]
