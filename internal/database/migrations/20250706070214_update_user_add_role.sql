-- +goose Up
-- +goose StatementBegin
ALTER TYPE USER_ROLE ADD VALUE 'USER';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TYPE USER_ROLE DROP VALUE 'USER';
-- +goose StatementEnd
