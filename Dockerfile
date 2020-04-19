FROM scratch
WORKDIR /app
COPY bl3-bot .
COPY balance_to_inv_key.json .
COPY inventory_raw.json .
ENTRYPOINT ["./bl3-bot"]