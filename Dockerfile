FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch
WORKDIR /app
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY bl3-bot .
COPY balance_to_inv_key.json .
COPY inventory_raw.json .
ENTRYPOINT ["./bl3-bot"]