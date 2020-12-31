package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	url := "https://images2.alphacoders.com/606/thumb-1920-606275.jpg"

	err := download(url)
	if err != nil {
		fmt.Println(err.Error)
	} else {
		fmt.Println("Download Berhasil")
	}
}

func download(url string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	fileName := path.Base(response.Request.URL.String())

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)

	return err
}
