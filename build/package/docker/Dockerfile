FROM golang:1.16-buster as builder
COPY . /build
WORKDIR /build/cmd/hibergate
RUN go build -o /build/hibergate .
FROM gcr.io/distroless/base-debian10:latest
COPY --from=builder /build/hibergate /usr/local/bin/hibergate
RUN chmod +x /usr/local/bin/hibergate
CMD ["/usr/local/bin/hibergate"]