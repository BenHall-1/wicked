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

func GetAllImagesUploadedByHash(hash string) []ImageUploaded {
	var previousImages []ImageUploaded
	query := "SELECT * FROM images_uploaded WHERE image_hash = '%s'"
	query = fmt.Sprintf(query, hash)
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
