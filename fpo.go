package main

import (
	// "fmt"
	"net/http"
	"image"
	"image/draw"
	"image/color"
	"image/png"
	"regexp"
	"strconv"
)

type ImageHandler struct{}

var (
	dimensions_regex, err = regexp.Compile(`/([0-9]+)x([0-9]+)$`)
	gray = color.RGBA{200, 200, 200, 255}
	canvas = &image.Uniform{gray}
)

func (h ImageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf("request path: %v ", r.URL.Path)

	if err != nil {
		http.Error(w, "", 400)
		return
	}

	dims := dimensions_regex.FindAllStringSubmatch(r.URL.Path, -1)

	if (dims == nil) {
		http.Error(w, "", 400)
		return
	}

	height, err := strconv.Atoi(dims[0][1])
	width, err := strconv.Atoi(dims[0][2])

	// fmt.Printf("request for image, h: %v, w: %v\n", height, width)

	if err != nil {
		http.Error(w, "", 400)
		return
	}

	m := image.NewRGBA(image.Rect(0, 0, height, width))
	draw.Draw(m, m.Bounds(), canvas, image.ZP, draw.Src)

        w.Header().Set("Content-type", "image/png")
        w.Header().Set("Cache-control", "public, max-age=259200")

	png.Encode(w, m)
}

func main() {
	var i ImageHandler
	http.ListenAndServe(":4000", i)
}
