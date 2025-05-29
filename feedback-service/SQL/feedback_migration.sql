CREATE TABLE feedbacks (
                           id SERIAL PRIMARY KEY,
                           rental_id BIGINT NOT NULL,
                           user_id BIGINT NOT NULL,
                           rating INT NOT NULL,
                           comment TEXT,
                           created_at TIMESTAMP NOT NULL
);
