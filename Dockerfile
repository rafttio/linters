FROM golang:1.17 as builder

COPY / /build
WORKDIR /build
RUN CGO_ENABLED=0 go build -trimpath -o raftt-lint ./cmd/all/main.go

# stage 2
FROM golang:1.17
COPY --from=builder /build/raftt-lint /usr/bin/
CMD ["raftt-lint"]
