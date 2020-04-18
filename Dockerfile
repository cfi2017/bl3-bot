FROM scratch
WORKDIR /app
COPY bot .
COPY balance_to_inv_key.json .
COPY inventory_raw.json .
ENTRYPOINT ["./bot"]