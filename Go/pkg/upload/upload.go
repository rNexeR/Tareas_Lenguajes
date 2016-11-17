package upload

import (
	//"fmt"
	"net/http"
	"os"
	"io"
)

func UploadFile(filename string, r *http.Request) {
	files := r.MultipartForm.File["files"]
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