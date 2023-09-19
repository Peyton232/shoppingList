-- File: 2_create_user_table.up.sql

-- Create the "User" table
CREATE TABLE
IF NOT EXISTS "User"
(
    "id" UUID PRIMARY KEY,
    "item_ids" UUID[] -- An array of UUIDs linking to items
);

-- Create a foreign key constraint to ensure item IDs exist in the items table
ALTER TABLE "User" ADD CONSTRAINT fk_user_item_ids
    FOREIGN KEY ("item_ids") REFERENCES "Item"("id");

-- Create an index on the "item_ids" array for efficient querying
CREATE INDEX idx_user_item_ids ON "User" USING GIN
("item_ids");
