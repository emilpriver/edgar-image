package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func DownloadRemoteImageToBuffer(w http.ResponseWriter, src string) (*bytes.Reader, string) {
	request, err := http.Get(src)

	if err != nil {
		HttpErrorMessage(w, "Error downloading image")
	}

	defer request.Body.Close()

	imageBytes, err := ioutil.ReadAll(request.Body)

	contentType := request.Header.Get("Content-Type")

	if len(contentType) == 0 {
		HttpErrorMessage(w, "Can't read format for image")
	}

	return bytes.NewReader(imageBytes), contentType
}
