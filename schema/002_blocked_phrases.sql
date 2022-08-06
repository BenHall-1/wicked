
-- Start the transaction
BEGIN;

CREATE TABLE "blocked_phrases" (
	"id" INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	"guild_id" BIGINT NOT NULL,
	"creator" BIGINT NOT NULL,
	"raw_phrase" VARCHAR(max) NOT NULL,
	"regex_phrase" VARCHAR(max) NOT NULL,
	"created_date" timestamp with time zone default (now() at time zone 'utc'),
);

-- Commit the change.
COMMIT;