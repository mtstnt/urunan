DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS participants;
DROP TABLE IF EXISTS bills;
DROP TABLE IF EXISTS items;
DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,

    full_name       VARCHAR(255) NOT NULL,
    email           VARCHAR(255) NOT NULL UNIQUE,
    password        VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS bills (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,

    host_user_id    INTEGER NOT NULL,

    title           VARCHAR(255) NOT NULL,
    description     VARCHAR(255) NOT NULL,

    FOREIGN KEY (host_user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS participants (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,

    bill_id         INTEGER NOT NULL,
    user_id         INTEGER NOT NULL,

    joined_at       INTEGER NOT NULL,
    payment_status  BOOLEAN DEFAULT FALSE,

    FOREIGN KEY (bill_id) REFERENCES bills(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS items (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,

    bill_id         INTEGER NOT NULL,

    name            VARCHAR(255) NOT NULL,
    price           DECIMAL NOT NULL,
    initial_qty     INTEGER NOT NULL,

    FOREIGN KEY (bill_id) REFERENCES bills(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS orders (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,

    participant_id  INTEGER NOT NULL,
    item_id         INTEGER NOT NULL,

    qty             INTEGER NOT NULL,

    FOREIGN KEY (participant_id) REFERENCES participants(id) ON DELETE CASCADE,
    FOREIGN KEY (item_id) REFERENCES items(id) ON DELETE CASCADE
);
