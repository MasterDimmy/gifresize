package main

import (
	"bufio"
	"bytes"
	"image"
	"image/draw"
	"image/gif"
	"os"

	"testing"

	"github.com/nfnt/resize"
)

func Test_GifResize(t *testing.T) {
	infile, err := os.Open("test/test.gif")
	if err != nil {
		t.Fatal(err)
	}

	img_gif, err := gif.DecodeAll(infile)

	var b bytes.Buffer
	out := bufio.NewWriter(&b)

	for i, img := range img_gif.Image {
		img2 := resize.Resize(50, 50, img, resize.Lanczos3)
		img3 := image.NewPaletted(img2.Bounds(), img.Palette)
		draw.Draw(img3, img2.Bounds(), img2, image.ZP, draw.Src)
		img_gif.Image[i] = img3
	}
	img_gif.Config.Width = 50
	img_gif.Config.Height = 50
	err = gif.EncodeAll(out, img_gif)
	if err != nil {
		t.Fatal(err)
	}

	err = os.WriteFile("D:/temp/done.gif", b.Bytes(), 0755)
	if err != nil {
		t.Fatal(err)
	}
}
