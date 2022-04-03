package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
)

func main() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "Michael Jackson")

	part := make(textproto.MIMEHeader)
	part.Set("Content-Type", "image/jpeg")
	part.Set("Content-Disposition", `form-data; name="thumbnail"; filename="photo/photo.jpg"`)

	fileWriter, err := writer.CreatePart(part)
	if err != nil {
		panic(err)
	}

	readFile, err := os.Open("photo/photo.jpg")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	io.Copy(fileWriter, readFile)
	writer.Close()

	res, err := http.Post("http://localhost:50510", writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	log.Println(string(body))
	log.Println("Status: ", res.Status)
	log.Println("StatusCode: ", res.StatusCode)
	log.Println("Headers: ", res.Header)
	log.Println("Content-Length: ", res.Header.Get("Content-Length"))
}
