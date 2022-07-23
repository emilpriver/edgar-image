package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/nfnt/resize"
)

type RequestBody struct {
	Type int `json:"type"`
}

type Response struct {
	Type int `json:"type"`
}

func HandleImageHttp(w http.ResponseWriter, r *http.Request) {
	image_src_query, ok := r.URL.Query()["src"]

	if !ok || len(image_src_query[0]) < 1 {
		log.Println("Url Param 'src' is missing")

		HttpErrorMessage(w, "Url Param 'src' is missing")

		return
	}

	image_width_query, ok := r.URL.Query()["w"]

	if !ok || len(image_width_query[0]) < 1 {
		log.Println("Url Param 'w' is missing")

		HttpErrorMessage(w, "Url Param 'w' is missing")

		return
	}

	image_width, image_src := image_width_query[0], image_src_query[0]

	image_width_int, _ := strconv.ParseUint(image_width, 10, 64)

	downloaded_image, format := DownloadRemoteImageToBuffer(w, image_src)

	fmt.Println(format)

	image, _, err := image.Decode(downloaded_image)

	resize_image := resize.Resize(uint(image_width_int), 0, image, resize.Lanczos3)

	user_accept_header := r.Header.Get("Accept")

	fmt.Println(user_accept_header)

	if strings.Contains(user_accept_header, "image/webp") {
		err = jpeg.Encode(w, resize_image, nil)

		if err != nil {
			log.Println(err.Error())
			return
		}
	}

	if format == "image/jpeg" {
		err = jpeg.Encode(w, resize_image, nil)

		if err != nil {
			log.Println(err.Error())
			return
		}
	}

	if format == "image/png" {
		err = png.Encode(w, resize_image)

		if err != nil {
			log.Println(err.Error())
			return
		}
	}

	return
}
