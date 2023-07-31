package middleware

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UploadFile(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("uploadimage") // nangkep file (byte)

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		src, _ := file.Open() // path

		fmt.Println("src:", src)

		defer src.Close() // -> LIFO (last in first out), memory leaks

		tempFile, err := ioutil.TempFile("uploads", "image-*.png")
		if err != nil {
			log.Fatal(err)
		}

		defer tempFile.Close()

		fmt.Println("tempFile:", tempFile)

		writtenCopy, err := io.Copy(tempFile, src)

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		fmt.Println("written copy:", writtenCopy)

		data := tempFile.Name() // uploads/image-123124212.png
		fmt.Println("data name utuh:", data)
		filename := data[8:]

		fmt.Println("filename terpotong", filename)

		c.Set("dataFile", filename) // image-12321321.png

		return next(c)
	}
}
