-- Таблица для заказов
CREATE TABLE orders (
    order_uid VARCHAR(50) PRIMARY KEY,
    track_number VARCHAR(50) NOT NULL,
    entry VARCHAR(10) NOT NULL,
    deliveries_id INTEGER,
    payments_id INTEGER,
    items_id INTEGER,
    locale VARCHAR(10),
    internal_signature VARCHAR(255),
    customer_id VARCHAR(50),
    delivery_service VARCHAR(50),
    shardkey VARCHAR(10),
    sm_id INTEGER,
    date_created TIMESTAMP WITH TIME ZONE,
    oof_shard VARCHAR(10)
);

-- Таблица для доставки
CREATE TABLE deliveries (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    phone VARCHAR(20),
    zip VARCHAR(20),
    city VARCHAR(100),
    address TEXT,
    region VARCHAR(100),
    email VARCHAR(100)
);

-- Таблица для платежей
CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    transaction VARCHAR(50),
    request_id VARCHAR(100),
    currency VARCHAR(10),
    provider VARCHAR(50),
    amount INTEGER,
    payment_dt BIGINT,
    bank VARCHAR(50),
    delivery_cost INTEGER,
    goods_total INTEGER,
    custom_fee INTEGER
);

-- Таблица для товаров
CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    chrt_id INTEGER NOT NULL,
    track_number VARCHAR(50),
    price INTEGER,
    rid VARCHAR(100),
    name VARCHAR(255),
    sale INTEGER,
    size VARCHAR(10),
    total_price INTEGER,
    nm_id INTEGER,
    brand VARCHAR(100),
    status INTEGER
);
