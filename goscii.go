/*Gooscii

Goscii is a very poor quality ascii art generator written in go.

Examples:
$ goscii                               # defaults, uses pug image, 120 chars
$ goscii -i images/flower.jpg -c 140   # Image path and char column count
*/

package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

const letters = "@$B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'."

func main() {

	var opts struct {
		ImagePath string `short:"i" long:"image" description:"Path to the image file" default:"images/pug.jpg"`
		Chars     int    `short:"c" long:"chars" description:"The output width in characters" default:"120"`
		Verbose   bool   `short:"v" long:"verbose" description:"Print image details and more"`
	}

	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(opts.ImagePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	cols := opts.Chars
	pixels := float32(img.Bounds().Dx()) / float32(cols)
	rows := float32(img.Bounds().Dy()) / pixels

	if opts.Verbose {
		fmt.Println("Image size: ", img.Bounds().Dx(), "x", img.Bounds().Dx())
		fmt.Println("Columns: ", cols)
		fmt.Println("Rows: ", rows)
		fmt.Println("Pixels: ", pixels)
	}

	for i := 1; i <= int(rows); i++ {
		for p := 1; p <= cols; p++ {
			if p == cols {
				fmt.Println("")
			}
			clr := img.At(p*int(pixels)-1, i*int(pixels)-1) // no image filter to computer color, just grab a single pixel
			r, g, b, _ := clr.RGBA()
			lum := 0.299*float32(r) + 0.587*float32(g) + 0.114*float32(b)
			charIndex := int(lum / 950) // 16bit color (65535) to valid letters index
			fmt.Print(string(letters[charIndex]))
		}
	}
}
