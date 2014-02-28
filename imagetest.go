package main

import (
	"log"
	"os"
	"image/draw"
	"image"
	"image/jpeg"
	"code.google.com/p/graphics-go/graphics"
)

func main() {

	fileImg1, err := os.Open("img_1.jpg")
	defer fileImg1.Close()

	if err != nil {
		log.Fatal(err)
	}

	img1, _, err := image.Decode(fileImg1)

	if err != nil {
		log.Fatal(err)
	}

	fileImg2, err := os.Open("img_2.jpg")
	defer fileImg2.Close()

	if err != nil {
		log.Fatal(err)
	}

	img2, _, err := image.Decode(fileImg2)

	if err != nil {
		log.Fatal(err)
	}

	placeholderImg := image.NewRGBA(image.Rect(0, 0, 700, 500))

	draw.Draw(placeholderImg, placeholderImg.Bounds(), img1, image.Point{0,0}, draw.Src)

	graphics.Rotate(placeholderImg, img2, &graphics.RotateOptions{1})

	toimg, err := os.Create("output.jpg")
	defer toimg.Close()

	if err != nil {
		log.Fatal(err)
	}

	jpeg.Encode(toimg, placeholderImg, &jpeg.Options{jpeg.DefaultQuality})
}