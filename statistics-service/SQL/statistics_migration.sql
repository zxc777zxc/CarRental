CREATE TABLE statistics (
                            id SERIAL PRIMARY KEY,
                            total_orders BIGINT DEFAULT 0,
                            total_revenue FLOAT DEFAULT 0,
                            average_rating FLOAT DEFAULT 0,
                            total_feedbacks BIGINT DEFAULT 0
);

-- Insert initial empty row
INSERT INTO statistics (id) VALUES (1);
