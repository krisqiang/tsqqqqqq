package util

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"runtime"
	"time"
)

func Upload(f *multipart.FileHeader) (res string) {
	out, err := f.Open()
	if err != nil {
		panic(err)
	}
	defer out.Close()
	ext := path.Ext(f.Filename)
	path := StaticFileRouter()
	fileName := getFileName()
	input, err := os.OpenFile(path+fileName+ext, os.O_CREATE|os.O_RDWR, 0325)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	_, err = io.Copy(input, out)
	if err != nil {
		panic(err)
	}
	return fileName + ext
}

func StaticFileRouter() (res string) {
	os := runtime.GOOS
	fmt.Println(os)
	if os == "windows" {
		res = "E:\\image\\"
	} else if os == "linux" {
		res = "/root/image/"
	}
	return
}

func getFileName() (res string) {
	date := time.Now().UnixNano()
	res = fmt.Sprintf("%v", date)
	return
}
