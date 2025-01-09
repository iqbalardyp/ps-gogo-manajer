-- Drop tables
DROP TABLE IF EXISTS employees CASCADE;
DROP TABLE IF EXISTS departments CASCADE;
DROP TABLE IF EXISTS users CASCADE;

-- DROP ENUM
DROP TYPE IF EXISTS enum_gender CASCADE;

-- Drop trigger function
DROP FUNCTION IF EXISTS trigger_set_timestamp() CASCADE;