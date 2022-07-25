
-- Start the transaction
BEGIN;

CREATE TABLE "images_uploaded" (
	"user_id" BIGINT NOT NULL,
	"message_id" BIGINT NOT NULL,
	"image_hash" VARCHAR(max) NOT NULL,
	"created_date" timestamp with time zone default (now() at time zone 'utc'),
	PRIMARY KEY ("user_id", "message_id")
);

-- Commit the change.
COMMIT;