-- File: 1_create_item_table.up.sql

-- Create the "Item" table
CREATE TABLE IF NOT EXISTS "Item" (
    "id" UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    "name" VARCHAR,
    "description" TEXT,
    "amount" VARCHAR,
    "completed" BOOLEAN,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Create an INSERT trigger for "created_at"
CREATE OR REPLACE FUNCTION set_created_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.created_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_created_at_trigger
BEFORE INSERT ON "Item"
FOR EACH ROW
EXECUTE FUNCTION set_created_at();

-- Create an UPDATE trigger for "updated_at"
CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_updated_at_trigger
BEFORE UPDATE ON "Item"
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();
