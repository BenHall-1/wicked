package utils

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

func GetBase64FromUrl(url string) string {
	response, err := http.Get(url)

	if err != nil || response.StatusCode != 200 {
		fmt.Println("Error whilst downloading the file")
		return ""
	}
	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Error whilst reading the file")
		return ""
	}

	var base64Encoding string
	mimeType := http.DetectContentType(bytes)

	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	base64Encoding += toBase64(bytes)

	return base64Encoding
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
