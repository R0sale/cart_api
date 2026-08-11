-- +goose Up
-- +goose StatementBegin
CREATE TABLE carts (
    id SERIAL PRIMARY KEY
);

CREATE TABLE carts_items (
    id SERIAL PRIMARY KEY,
    cart_id INT NOT NULL REFERENCES carts(Id) ON DELETE CASCADE,
    product TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE carts;
DROP TABLE carts_items;
-- +goose StatementEnd
