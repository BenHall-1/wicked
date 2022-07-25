package utils

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func EscapeSpecialCharacters(value string) string {
	replace := []struct {
		raw  string
		safe string
	}{
		{"\\", "\\\\"},   // backslash
		{"'", `\'`},      // single quote
		{`"`, `\"`},      // double quote
		{"\n", "\\n"},    // new line
		{"\r", "\\r"},    // carriage return
		{"\\0", "\\\\0"}, // null
		{"\x1a", "\\Z"},  // control+Z
		{"\x08", "\\b"},  // backspace
		{"%", "\\%"},     // %
	}

	for _, str := range replace {
		value = strings.Replace(value, str.raw, str.safe, -1)
	}

	return value
}

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
