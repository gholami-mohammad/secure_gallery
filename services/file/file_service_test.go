package file

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReadFile(t *testing.T) {
	Convey("pass", t, func() {
		filePath := os.TempDir() + "/file.txt"
		os.Create(filePath)
		defer os.Remove(filePath)

		_, err := ReadFile(filePath)
		So(err, ShouldBeNil)
	})
	Convey("File not exists", t, func() {
		filePath := os.TempDir() + "/somefile.txt"

		_, err := ReadFile(filePath)
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "no such file or directory")
	})
}

func TestWriteFile(t *testing.T) {
	Convey("pass", t, func() {
		filePath := os.TempDir() + "/file.txt"
		os.Create(filePath)
		defer os.Remove(filePath)

		err := WriteFile(filePath, []byte("hello"))
		So(err, ShouldBeNil)
	})
}
