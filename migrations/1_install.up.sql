CREATE TABLE IF NOT EXISTS companies (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    industry TEXT NOT NULL,
    website TEXT ,
    location TEXT 
);

CREATE TABLE IF NOT EXISTS payments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    cron TEXT NOT NULL,
    url TEXT NOT NULL,
    companyid INTEGER, 
    status BOOLEAN NOT NULL CHECK (status IN (0, 1)) DEFAULT 1
); 
