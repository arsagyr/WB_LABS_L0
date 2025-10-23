INSERT INTO deliveries (
    name, phone, 
    zip, city, address, 
    region, email
) VALUES (
    'Test Testov', '+9720000000', 
    '2639809', 'Kiryat Mozkin', 'Ploshad Mira 15', 
    'Kraiot', 'test@gmail.com'
);

INSERT INTO payments (
    transaction, request_id, 
    currency, provider, amount, 
    payment_dt, bank, delivery_cost, 
    goods_total,  custom_fee
) VALUES (
    'b563feb7b2b84b6test', '', 
    'USD', 'wbpay', 1817, 
    1637907727, 'alpha', 1500, 
    317, 0
);

INSERT INTO orders (
    order_uid, track_number, entry, 
    locale, internal_signature, customer_id, 
    delivery_service, shardkey, sm_id, 
    date_created, oof_shard, 
    delivery_id, payment_id 
    ) VALUES (
    'b563feb7b2b84b6test', 'WBILMTESTTRACK', 'WBIL', 
    'en', '', 'test',
    'meest', '9', 99, 
    '2021-11-26T06:22:19Z', '1',
    1, 1
);
 
INSERT INTO items (
    chrt_id, track_number, price, rid, name, 
    sale, size, total_price, nm_id, brand, status,
    order_id
) VALUES (
    9934930, 'WBILMTESTTRACK', 453, 
    'ab4219087a764ae0btest', 'Mascaras', 30, '0', 317, 
    2389212, 'Vivienne Sabo', 202,
    1
);