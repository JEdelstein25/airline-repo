-- Create the loyalty programs table
CREATE TABLE loyalty_programs (
    id UUID PRIMARY KEY,
    airline_id UUID NOT NULL REFERENCES airlines(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Create index on airline_id for faster lookups
CREATE INDEX loyalty_programs_airline_id_idx ON loyalty_programs(airline_id);

-- Create the loyalty members table
CREATE TABLE loyalty_members (
    id UUID PRIMARY KEY,
    loyalty_program_id UUID NOT NULL REFERENCES loyalty_programs(id) ON DELETE CASCADE,
    passenger_id UUID NOT NULL REFERENCES passengers(id) ON DELETE CASCADE,
    membership_number TEXT NOT NULL,
    tier TEXT NOT NULL DEFAULT 'Basic',
    points INTEGER NOT NULL DEFAULT 0,
    join_date TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Create indexes for faster lookups
CREATE INDEX loyalty_members_program_id_idx ON loyalty_members(loyalty_program_id);
CREATE INDEX loyalty_members_passenger_id_idx ON loyalty_members(passenger_id);

-- Ensure membership numbers are unique within a program
CREATE UNIQUE INDEX loyalty_members_program_membership_idx ON loyalty_members(loyalty_program_id, membership_number);

-- Ensure a passenger can only be enrolled once in a program
CREATE UNIQUE INDEX loyalty_members_program_passenger_idx ON loyalty_members(loyalty_program_id, passenger_id);