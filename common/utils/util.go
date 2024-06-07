package utils

import (
	"path/filepath"
	"strings"

	"github.com/go-playground/validator"
)

var Validator *validator.Validate

const UploadDir = "./uploads"

// GetFileContentType determines the MIME type of the file based on its extension
func GetFileContentType(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	case ".bmp":
		return "image/bmp"
	case ".mp3":
		return "audio/mpeg"
	default:
		return "application/octet-stream"
	}
}

// remove 1 element in array
func RemoveElementFromArray(arr *[]string, element string) {
	index := -1
	for i, value := range *arr {
		if value == element {
			index = i
			break
		}
	}

	if index == -1 {
		return
	}

	*arr = append((*arr)[:index], (*arr)[index+1:]...)
}
