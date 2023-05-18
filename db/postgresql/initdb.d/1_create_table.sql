CREATE TABLE IF NOT EXISTS "users" (
    "id"    UUID NOT NULL,
    "name"  VARCHAR(255) NOT NULL,
    PRIMARY KEY("id")
);

CREATE TABLE IF NOT EXISTS "todos" (
    "id"        UUID NOT NULL,
    "user_id"   UUID NOT NULL,
    "name"      VARCHAR(255) NOT NULL,
    "done"      BOOLEAN DEFAULT 'false' NOT NULL,
    PRIMARY KEY ("id"),
    FOREIGN KEY("user_id") REFERENCES "users"("id") 
);
