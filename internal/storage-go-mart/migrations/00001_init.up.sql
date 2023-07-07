BEGIN TRANSACTION;

CREATE TABLE IF NOT EXISTS users(
                        id BIGINT GENERATED ALWAYS AS IDENTITY,
                        login TEXT NOT NULL UNIQUE,
                        password TEXT NOT NULL,
                        PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS orders(
                        id BIGINT GENERATED ALWAYS AS IDENTITY,
                        users_id BIGINT,
                        status TEXT,
                        accrual DOUBLE PRECISION,
                        uploaded_at TIMESTAMP WITH TIME ZONE NOT NULL,
                        PRIMARY KEY(id),
                        FOREIGN KEY(users_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS balance(
                        id BIGINT GENERATED ALWAYS AS IDENTITY,
                        users_id BIGINT,
                        balance DOUBLE PRECISION,
                        FOREIGN KEY(users_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS withdraw(
                        users_id BIGINT NOT NULL PRIMARY KEY REFERENCES users(id),
                        order_id BIGINT NOT NULL,
                        amount DOUBLE PRECISION,
                        processed_at TIMESTAMP WITH TIME ZONE NOT NULL,
                        FOREIGN KEY(order_id) REFERENCES orders(id)
);

COMMIT;