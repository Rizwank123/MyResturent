-- +goose Up
-- +goose StatementBegin
ALTER TABLE menu_card ADD COLUMN offer_price float8 NOT NULL DEFAULT 0.0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE menu_card DROP COLUMN offer_price;
-- +goose StatementEnd
