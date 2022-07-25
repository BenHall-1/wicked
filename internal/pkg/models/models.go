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

func GetAllImagesUploadedByHash(hash string) []ImageUploaded {
	var previousImages []ImageUploaded
	if err := db.Database.Select(&previousImages, "SELECT * FROM images_uploaded WHERE image_hash = ?", hash); err != nil {
		fmt.Println(err)
		return nil
	}

	return previousImages
}

func AddImageUpload(userId, messageId, imageHash string) {
	if _, err := db.Database.Exec("INSERT INTO images_uploaded (user_id, message_id, image_hash) VALUES (?, ?, ?)", userId, messageId, imageHash); err != nil {
		fmt.Println(err)
	}
}
