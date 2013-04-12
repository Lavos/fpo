package main

import (
	// "fmt"
	"encoding/hex"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
)

type ImageHandler struct{}

var (
	request_regex, _ = regexp.Compile(`/([0-9]+)x([0-9]+)(?:/([0-9a-f]{6}|random)?)`)
	gray             = color.RGBA{240, 240, 240, 255}
	canvas           = &image.Uniform{gray}
)

func (h ImageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		canvas.C = gray
	}()

	options := request_regex.FindAllStringSubmatch(r.URL.Path, -1)

	if options == nil {
		http.Error(w, "", 400)
		return
	}

	height, err := strconv.Atoi(options[0][1])
	width, err := strconv.Atoi(options[0][2])
	colorhex := options[0][3]

	if len(colorhex) == 0 {
		canvas.C = gray
	} else if colorhex == "random" {
		canvas.C = color.RGBA{uint8(rand.Intn(256)), uint8(rand.Intn(256)), uint8(rand.Intn(256)), 255}
	} else {
		red, green, blue := get_colors(colorhex)
		canvas.C = color.RGBA{red, green, blue, 255}
	}

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

func get_colors(colorhex string) (byte, byte, byte) {
	red, _ := hex.DecodeString(colorhex[0:2])
	green, _ := hex.DecodeString(colorhex[2:4])
	blue, _ := hex.DecodeString(colorhex[4:])

	return red[0], green[0], blue[0]
}

func main() {
	var i ImageHandler
	http.ListenAndServe(":4000", i)
}
