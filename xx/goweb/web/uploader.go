package main

import (
	"bytes"
	"mime/multipart"
	"fmt"
	"os"
	"io"
	"net/http"
	"io/ioutil"
)

func main() {
	target_url := "http://localhost:88/upload"
	filename := "uploader.go"
	postFile(filename, target_url)
}
func postFile(fileName string, url string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("uploadFile", fileName)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fh, err := os.Open(fileName)
	if err != nil {
		fmt.Println("OpenFile", err)
		return err
	}
	defer fh.Close()

	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(url, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}
