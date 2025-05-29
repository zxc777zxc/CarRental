CREATE TABLE rentals (
                         id SERIAL PRIMARY KEY,
                         user_id BIGINT NOT NULL,
                         car_id BIGINT NOT NULL,
                         start_date TIMESTAMP NOT NULL,
                         end_date TIMESTAMP NOT NULL,
                         total_cost NUMERIC(10,2) NOT NULL,
                         status VARCHAR(20) NOT NULL
);
