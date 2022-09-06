CREATE TYPE INVENTORY_CHECK_STRATEGY AS ENUM('NEVER', 'ADD_TO_CART');

CREATE TYPE INVENTORY_RESERVATION_STRATEGY AS ENUM('NEVER', 'ADD_TO_CART', 'SUBMIT_ORDER');

CREATE TYPE INVENTORY_TYPE AS ENUM('PHYSICAL', 'VIRTUAL');

CREATE TYPE PRICE_INFO AS (amount DECIMAL, currncy VARCHAR(16));

CREATE TYPE ASSET_TYPE AS ENUM('IMAGE', 'VIDEO', 'EMBED', 'DOC', 'TEXT', 'UNKNOWN');

CREATE TYPE PRODUCT_TYPE AS ENUM (
    'STANDARD',
    'VARIANT_BASED',
    'BUNDLE',
    'SELECTOR',
    'MERCHANDISING'
);

CREATE TABLE IF NOT EXISTS category (
    id VARCHAR(36) NOT NULL,
    context_id VARCHAR(36) NOT NULL,
    parent_context_id VARCHAR(36),
    active_start_date TIMESTAMP WITHOUT TIME ZONE,
    active_end_date TIMESTAMP WITHOUT TIME ZONE,
    archived BOOLEAN,
    attributes JSON,
    title VARCHAR(255),
    short_description TEXT,
    display_template VARCHAR(255),
    url VARCHAR(255),
    meta_title VARCHAR(255),
    meta_description VARCHAR(255)
);

CREATE TABLE if NOT EXISTS category_asset (
    id VARCHAR(36) NOT NULL,
    context_id VARCHAR(36) NOT NULL,
    category_context_id VARCHAR(36) NOT NULL,
    is_primary BOOLEAN,
    sorting VARCHAR(36),
    asset_type ASSET_TYPE,
    url VARCHAR(255),
    alt_text VARCHAR(255),
    title VARCHAR(255)
);

CREATE TABLE if NOT EXISTS category_product (
    id VARCHAR(36) NOT NULL,
    context_id VARCHAR(36) NOT NULL,
    product_context_id VARCHAR(36) NOT NULL,
    category_context_id VARCHAR(36) NOT NULL,
    is_primary BOOLEAN,
    sorting VARCHAR(36)
);

CREATE TABLE IF NOT EXISTS product (
    id VARCHAR(36) NOT NULL,
    context_id VARCHAR(36) NOT NULL,
    active_start_date TIMESTAMP WITHOUT TIME ZONE,
    active_end_date TIMESTAMP WITHOUT TIME ZONE,
    attributes JSON,
    display_template VARCHAR(255),
    available_online BOOLEAN,
    cost PRICE_INFO,
    currncy VARCHAR(16),
    pricing_key VARCHAR(36),
    default_price PRICE_INFO,
    title VARCHAR(255),
    long_description TEXT,
    short_description VARCHAR(255),
    fulfilment_flat_rates JSON,
    included_products JSON,
    indivdiually_sold BOOLEAN,
    inventory_check_strategy INVENTORY_CHECK_STRATEGY,
    inventory_reservation_strategy INVENTORY_RESERVATION_STRATEGY,
    inventory_type INVENTORY_TYPE,
    keywords VARCHAR(36) [ ],
    meta_title VARCHAR(255),
    meta_description VARCHAR(255),
    is_online BOOLEAN,
    options JSON,
    sale_price PRICE_INFO,
    searchable BOOLEAN,
    sku VARCHAR(36),
    upc VARCHAR(36),
    uri VARCHAR(36),
    external_id VARCHAR(36),
    product_type PRODUCT_TYPE,
    tags VARCHAR(16) [ ]
);

CREATE TABLE IF NOT EXISTS product_asset (
    id VARCHAR(36) NOT NULL,
    context_id VARCHAR(36) NOT NULL,
    product_context_id VARCHAR(36) NOT NULL,
    is_primary BOOLEAN,
    sorting VARCHAR(36),
    asset_type ASSET_TYPE,
    url VARCHAR(36),
    alt_text VARCHAR(255),
    title VARCHAR(36)
);

CREATE TABLE IF NOT EXISTS variant (
    id VARCHAR(36) NOT NULL,
    context_id VARCHAR(36) NOT NULL,
    product_context_id VARCHAR(36) NOT NULL,
    active_start_date TIMESTAMP WITHOUT TIME ZONE,
    active_end_date TIMESTAMP WITHOUT TIME ZONE,
    attributes JSON,
    cost PRICE_INFO,
    default_price PRICE_INFO,
    title VARCHAR(255),
    description_long TEXT,
    description_short VARCHAR(255),
    inventory_check_strategy INVENTORY_CHECK_STRATEGY,
    inventory_reservation_strategy INVENTORY_RESERVATION_STRATEGY,
    inventory_type INVENTORY_TYPE,
    is_online BOOLEAN,
    option_values JSON,
    upc VARCHAR(36)
);

CREATE INDEX IF NOT EXISTS not_safe_index_for_porduct_asset_porduct_context_id ON product_asset(product_context_id);

CREATE INDEX IF NOT EXISTS not_safe_index_for_porduct_asset_context_id ON product_asset(context_id);

CREATE INDEX IF NOT EXISTS not_safe_index_for_porduct_asset_sorting ON product_asset(sorting);

CREATE INDEX IF NOT EXISTS not_safe_index_for_porduct_context_id ON product(context_id);

CREATE INDEX IF NOT EXISTS not_safe_index_for_porduct_uri ON product(uri);

CREATE INDEX IF NOT EXISTS not_safe_index_for_porduct_sku ON product(sku);

CREATE INDEX IF NOT EXISTS not_safe_index_for_variant_context_id ON variant(context_id);

CREATE INDEX IF NOT EXISTS not_safe_index_for_variant_product_context_id ON variant(product_context_id);