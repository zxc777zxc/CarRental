CREATE TABLE IF NOT EXISTS cars (
                                    id SERIAL PRIMARY KEY,
                                    brand TEXT NOT NULL,
                                    model TEXT NOT NULL,
                                    fuel TEXT NOT NULL,
                                    transmission TEXT NOT NULL,
                                    price_per_day REAL NOT NULL
);