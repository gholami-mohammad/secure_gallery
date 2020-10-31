package file

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetFileExetension(t *testing.T) {
	Convey("Pass", t, func() {
		type testcase struct {
			filename string
			wanted   string
		}
		tests := []testcase{
			{"myfile.txt", "txt"},
			{"myfile.tar.gz", "gz"},
			{"myfile", ""},
		}
		for _, tc := range tests {
			ext := GetFileExetension(tc.filename)
			So(ext, ShouldEqual, tc.wanted)
		}
	})
}

func TestGenerateUniqueFilename(t *testing.T) {
	Convey("no same file", t, func() {
		dir := os.TempDir() + "/sg"
		os.RemoveAll(dir)
		got := GenerateUniqueFilename(dir, "image.jpg", 1)

		So(got, ShouldEqual, "image-001.jpg")
	})
	Convey("file with same name exists", t, func() {
		dir := os.TempDir() + "/sg"
		os.RemoveAll(dir)

		_ = os.MkdirAll(dir, os.ModePerm)
		_, _ = os.Create(dir + "/image.jpg")

		got := GenerateUniqueFilename(dir, "image.jpg", 1)

		So(got, ShouldEqual, "image-001.jpg")
	})
	Convey("some files with same name exist", t, func() {
		dir := os.TempDir() + "/sg"
		os.RemoveAll(dir)

		_ = os.MkdirAll(dir, os.ModePerm)
		_, _ = os.Create(dir + "/image.jpg")
		_, _ = os.Create(dir + "/image-001.jpg")
		_, _ = os.Create(dir + "/image-002.jpg")

		got := GenerateUniqueFilename(dir, "image.jpg", 1)

		So(got, ShouldEqual, "image-003.jpg")
	})
}
