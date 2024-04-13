CREATE TABLE IF NOT EXISTS tbl_customer (
                              id SERIAL PRIMARY KEY,
                              customer_name VARCHAR NOT NULL,
                              balance DECIMAL,
                              created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              updated_at TIMESTAMP,
                              deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tbl_items (
                           id SERIAL PRIMARY KEY,
                           item_name VARCHAR NOT NULL,
                           cost DECIMAL,
                           price DECIMAL,
                           sort INTEGER,
                           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                           updated_at TIMESTAMP,
                           deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS tbl_transaction (
                                 id SERIAL PRIMARY KEY,
                                 customer_id INTEGER REFERENCES tbl_customer(id),
                                 item_id INTEGER REFERENCES tbl_items(id),
                                 qty INTEGER,
                                 amount DECIMAL,
                                 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                 updated_at TIMESTAMP,
                                 deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS TransactionViews (
                                  id INTEGER PRIMARY KEY,
                                  customer_id INTEGER REFERENCES tbl_customer(id),
                                  customer_name VARCHAR,
                                  item_id INTEGER REFERENCES tbl_items(id),
                                  item_name VARCHAR,
                                  qty INTEGER,
                                  price DECIMAL,
                                  amount DECIMAL,
                                  created_at TIMESTAMP,
                                  updated_at TIMESTAMP,
                                  deleted_at TIMESTAMP
);
