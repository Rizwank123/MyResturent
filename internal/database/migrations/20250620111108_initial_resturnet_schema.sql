-- +goose Up
-- +goose StatementBegin

-- ENUM TYPES
CREATE TYPE USER_ROLE AS ENUM ('OWNER', 'SUPERVISOR');
CREATE TYPE CATEGORY AS ENUM ('VEG', 'NONVEG');
CREATE TYPE MEAL_TYPE AS ENUM ('DINNER', 'LUNCH', 'BREAKFAST');
CREATE TYPE FOOD_TYPE AS ENUM ('STARTER', 'MAINCOURSE');

-- TABLES

CREATE TABLE "public"."resturent" (
    "id" uuid NOT NULL DEFAULT gen_random_uuid(),
    "name" text NOT NULL,
    "address" JSONB DEFAULT '{}'::JSONB,
    "license" text NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT now(),
    "updated_at" timestamp NOT NULL DEFAULT now(),
    "deleted_at" timestamp,
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."users" (
    "id" uuid NOT NULL DEFAULT gen_random_uuid(),
    "name" text NOT NULL,
    "email" text NOT NULL,
    "password" text NOT NULL,
    "role" USER_ROLE NOT NULL,
    "mobile" text NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT now(),
    "updated_at" timestamp NOT NULL DEFAULT now(),
    "resturent_id" uuid NOT NULL,
    "deleted_at" timestamp,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("resturent_id") REFERENCES "resturent" ("id")
);

CREATE TABLE "public"."menu_card" (
    "id" uuid NOT NULL DEFAULT gen_random_uuid(),
    "resturent_id" uuid NOT NULL,
    "name" text,
    "price" float8 NOT NULL,
    "category" CATEGORY NOT NULL,
    "size" text,
    "image" text,
    "food_type" FOOD_TYPE NOT NULL,
    "meal_type" MEAL_TYPE,
    "description" text,
    "is_available" boolean NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT now(),
    "updated_at" timestamp NOT NULL DEFAULT now(),
    "deleted_at" timestamp,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("resturent_id") REFERENCES "resturent" ("id")
);

CREATE TABLE "public"."rating" (
    "id" uuid NOT NULL DEFAULT gen_random_uuid(),
    "resturent_id" uuid NOT NULL,
    "name" text,
    "rating" float8 NOT NULL,
    "review" text,
    "suggestion" text,
    "created_at" timestamp NOT NULL DEFAULT now(),
    "updated_at" timestamp NOT NULL DEFAULT now(),
    "deleted_at" timestamp,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("resturent_id") REFERENCES "resturent" ("id")
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS "rating";
DROP TABLE IF EXISTS "menu_card";
DROP TABLE IF EXISTS "users";
DROP TABLE IF EXISTS "resturent";

DROP TYPE IF EXISTS FOOD_TYPE;
DROP TYPE IF EXISTS MEAL_TYPE;
DROP TYPE IF EXISTS CATEGORY;
DROP TYPE IF EXISTS USER_ROLE;

-- +goose StatementEnd
