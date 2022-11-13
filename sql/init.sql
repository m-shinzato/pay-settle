CREATE TABLE debts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    debtor INTEGER NOT NULL,
    lender INTEGER NOT NULL,
    price INTEGER NOT NULL,
    memo TEXT NOT NULL,
    completed INTEGER NOT NULL DEFAULT 0,
    created_at TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime')),
    updated_at TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime'))
);
