BEGIN TRANSACTION;

CREATE TABLE orders(
    id INT PRIMARY KEY,
    status TEXT NOT NULL,
    accrual INT
);

CREATE TABLE product(
    id SERIAL PRIMARY KEY,
    description TEXT,
    price INT
);

CREATE TABLE reward(
    id SERIAL PRIMARY KEY,
    match TEXT NOT NULL,
    reward INT NOT NULL,
    reward_type VARCHAR(2)
);

CREATE TABLE order_product(
    order_product_id SERIAL PRIMARY KEY,
    order_id SERIAL NOT NULL REFERENCES public.orders (id) ON DELETE CASCADE,
    product_id SERIAL NOT NULL REFERENCES public.product (id) ON DELETE CASCADE
);
COMMIT;