package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"io"
	"os"
)

func UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("upload")
	filename := header.Filename
	fmt.Println(header.Filename)
	out, err := os.Create("./tmp/" + filename + ".png")
	if err != nil {
		glog.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		glog.Fatal(err)
	}
}
