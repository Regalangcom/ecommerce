CREATE TABLE products_images (
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    url VARCHAR(255) NOT NULL,
    is_primary BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);



CREATE INDEX idx_products_images_product_id ON products_images(product_id);
CREATE INDEX idx_products_images_is_primary ON products_images(is_primary);
CREATE INDEX idx_products_images_deleted_at ON products_images(deleted_at);