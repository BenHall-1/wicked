package models

import (
	"fmt"
	"time"

	"github.com/benhall-1/wicked/internal/pkg/db"
)

type ImageUploaded struct {
	UserId      string    `db:"user_id"`
	MessageId   string    `db:"message_id"`
	ImageHash   string    `db:"image_hash"`
	CreatedDate time.Time `db:"created_date"`
}

type Embed struct {
	Title       string
	Description string
}

type BlockedPhrase struct {
	Id          int       `db:"id"`
	GuildId     string    `db:"guild_id"`
	Creator     string    `db:"creator"`
	RawPhrase   string    `db:"raw_phrase"`
	RegexPhrase string    `db:"regex_phrase"`
	CreatedDate time.Time `db:"created_date"`
}

func GetAllImagesUploadedByHash(hash string, author string) []ImageUploaded {
	var previousImages []ImageUploaded
	query := "SELECT * FROM images_uploaded WHERE image_hash = '%s' AND user_id != '%s'"
	query = fmt.Sprintf(query, hash, author)
	if err := db.Database.Select(&previousImages, query); err != nil {
		return nil
	}

	return previousImages
}

func AddImageUpload(userId, messageId, imageHash string) {
	query := "INSERT INTO images_uploaded (user_id, message_id, image_hash) VALUES ('%s', '%s', '%s')"
	query = fmt.Sprintf(query, userId, messageId, imageHash)
	db.Database.Exec(query)
}

func GetAllPhrases(guildId string) []BlockedPhrase {
	var blockedPhrases []BlockedPhrase
	query := "SELECT * FROM blocked_phrases WHERE guild_id = $1"
	if err := db.Database.Select(&blockedPhrases, query); err != nil {
		return nil
	}

	return blockedPhrases
}

func GetPhrase(phrase string, guildId string) *BlockedPhrase {
	var blockedPhrase BlockedPhrase
	query := "SELECT * FROM blocked_phrases WHERE raw_phrase = $1 AND guild_id = $2"
	if err := db.Database.Select(&blockedPhrase, query, phrase); err != nil {
		return nil
	}

	return &blockedPhrase
}

func AddPhrase(guildId string, creator string, phrase string, regexPhrase string) {
	query := "INSERT INTO blocked_phrases (guild_id, creator, raw_phrase, regex_phrase) VALUES ('%s', '%s', '%s', '%s')"
	query = fmt.Sprintf(query, guildId, creator, phrase, regexPhrase)
	db.Database.Exec(query)
}
