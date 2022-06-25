-- set timezone
SET TIMEZONE="Africa/Lagos";

-- Create users Table
CREATE TABLE users (
    id SERIAL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE,
    email VARCHAR(200) NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    dob TIMESTAMP WITH TIME ZONE,

    UNIQUE(email),
    PRIMARY KEY (id)
)