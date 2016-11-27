package upload

import (
	//"fmt"
	"io"
	"net/http"
	"os"
)

func UploadFile(filename string, r *http.Request) {
	files := r.MultipartForm.File["file"]
	file := files[0]
	fOpen, err := file.Open()
	defer fOpen.Close()
	if err != nil {
		panic(err)
	}

	dst, err := os.Create(filename)
	defer dst.Close()
	if err != nil {
		panic(err)
	}

	if _, err := io.Copy(dst, fOpen); err != nil {
		panic(err)
	}
}
