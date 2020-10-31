package file

import (
	"fmt"
	"os"
	"strings"
)

func GetFileExetension(filename string) string {
	p := strings.Split(filename, ".")
	if len(p) <= 1 {
		return ""
	}

	return p[len(p)-1]
}

func GenerateUniqueFilename(dir, orginalFilename string, duplicationIndex uint) string {
	ext := GetFileExetension(orginalFilename)
	newFilename := strings.TrimRight(orginalFilename, "."+ext)

	filePath := fmt.Sprintf("%s/%s-%03d.%s", dir, newFilename, duplicationIndex, ext)
	if _, err := os.Stat(filePath); err != nil {
		// file not exists
		return fmt.Sprintf("%s-%03d.%s", newFilename, duplicationIndex, ext)
	}

	duplicationIndex++
	return GenerateUniqueFilename(dir, orginalFilename, duplicationIndex)
}
