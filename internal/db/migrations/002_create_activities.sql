CREATE TABLE IF NOT EXISTS activities (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    venue VARCHAR(255),
    host VARCHAR(255),
    status INT CHECK (status BETWEEN 1 AND 5) DEFAULT 1,
    remarks TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

COMMENT ON COLUMN activities.status IS '1: To be conducted, 2: Conducted, 3: Rescheduled, 4: Postponed Indefinitely, 5: Canceled';

CREATE TRIGGER activity_updated_at_trigger
BEFORE UPDATE ON activities
FOR EACH ROW
EXECUTE FUNCTION touch_updated_at();