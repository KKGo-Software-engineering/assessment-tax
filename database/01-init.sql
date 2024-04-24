CREATE TYPE tax_deduct_type AS ENUM ('donation', 'personal', 'k-receipt');

CREATE TABLE IF NOT EXISTS tax_deduct (
  id SERIAL PRIMARY KEY,
  type tax_deduct_type UNIQUE NOT NULL,
  min_amount NUMERIC CHECK(min_amount >= 0) NOT NULL,
  max_amount NUMERIC CHECK(max_amount >= min_amount) NOT NULL,
  amount NUMERIC CHECK(amount >= min_amount AND amount <= max_amount) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
  updated_at TIMESTAMPTZ
);

INSERT INTO tax_deduct(type, min_amount, max_amount, amount)
VALUES 
  ('donation', 0, 100000, 100000),
  ('personal', 10000, 100000, 60000),
  ('k-receipt', 1, 100000, 50000);
