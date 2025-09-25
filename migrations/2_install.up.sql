-- Add description columns
ALTER TABLE companies ADD COLUMN description TEXT NOT NULL;
ALTER TABLE payments ADD COLUMN description TEXT NOT NULL;

-- Create new table
CREATE TABLE IF NOT EXISTS paymentGroup (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT
);

-- Add new column (without constraint)
ALTER TABLE payments ADD COLUMN paymentgroupid INT;
