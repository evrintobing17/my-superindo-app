-- -------------------------------------------------------------
-- TablePlus 6.3.2(586)
--
-- https://tableplus.com/
--
-- Database: superindo
-- Generation Time: 2025-02-25 6:34:55.5530 AM
-- -------------------------------------------------------------


-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS myapp.category_id_seq;

-- Table Definition
CREATE TABLE "myapp"."category" (
    "id" int4 NOT NULL DEFAULT nextval('myapp.category_id_seq'::regclass),
    "name" varchar(100) NOT NULL,
    "description" varchar(100) NOT NULL,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS myapp.mapp_cart_product_id_seq;

-- Table Definition
CREATE TABLE "myapp"."mapp_cart_product" (
    "id" int4 NOT NULL DEFAULT nextval('myapp.mapp_cart_product_id_seq'::regclass),
    "cart_id" int4 NOT NULL,
    "product_id" int4 NOT NULL,
    "total" int4 NOT NULL,
    "is_active" bool NOT NULL DEFAULT true,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS myapp.mapp_cart_user_id_seq;

-- Table Definition
CREATE TABLE "myapp"."mapp_cart_user" (
    "id" int4 NOT NULL DEFAULT nextval('myapp.mapp_cart_user_id_seq'::regclass),
    "user_id" int4 NOT NULL,
    "is_active" bool NOT NULL DEFAULT true,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS myapp.products_id_seq;

-- Table Definition
CREATE TABLE "myapp"."products" (
    "id" int4 NOT NULL DEFAULT nextval('myapp.products_id_seq'::regclass),
    "category_id" int4 NOT NULL,
    "title" varchar(100) NOT NULL,
    "description" varchar(100) NOT NULL,
    PRIMARY KEY ("id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS myapp.users_id_seq;

-- Table Definition
CREATE TABLE "myapp"."users" (
    "id" int4 NOT NULL DEFAULT nextval('myapp.users_id_seq'::regclass),
    "name" varchar(100) NOT NULL,
    "email" varchar(100) NOT NULL,
    "password" varchar(255) NOT NULL,
    PRIMARY KEY ("id")
);

ALTER TABLE "myapp"."mapp_cart_product" ADD FOREIGN KEY ("cart_id") REFERENCES "myapp"."mapp_cart_user"("id") ON DELETE CASCADE;
ALTER TABLE "myapp"."mapp_cart_product" ADD FOREIGN KEY ("product_id") REFERENCES "myapp"."products"("id") ON DELETE CASCADE;
ALTER TABLE "myapp"."mapp_cart_user" ADD FOREIGN KEY ("user_id") REFERENCES "myapp"."users"("id") ON DELETE CASCADE;
ALTER TABLE "myapp"."products" ADD FOREIGN KEY ("category_id") REFERENCES "myapp"."category"("id") ON DELETE CASCADE;


-- Indices
CREATE UNIQUE INDEX users_email_key ON myapp.users USING btree (email);
