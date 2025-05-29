CREATE TABLE payments (
                          id SERIAL PRIMARY KEY,
                          rental_id BIGINT NOT NULL,
                          amount NUMERIC(10, 2) NOT NULL,
                          method VARCHAR(20) NOT NULL,
                          status VARCHAR(20) NOT NULL,
                          paid_at TIMESTAMP NOT NULL
);
