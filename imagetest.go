package main

import (
    // "fmt"
    "os"
    "image/draw"
    "image"
    "image/jpeg"
    "code.google.com/p/graphics-go/graphics"
)

func main() {
    fImg1, _ := os.Open("arrow1.jpg")
    defer fImg1.Close()
    img1, _, _ := image.Decode(fImg1)

    fImg2, _ := os.Open("arrow2.jpg")
    defer fImg2.Close()
    img2, _, _ := image.Decode(fImg2)

    m := image.NewRGBA(image.Rect(0, 0, 800, 600))
    draw.Draw(m, m.Bounds(), img1, image.Point{0,0}, draw.Src)
    //draw.Draw(m, m.Bounds(), img2, image.Point{-200,-200}, draw.Src)
    graphics.Rotate(m, img2, &graphics.RotateOptions{3.5})

    toimg, _ := os.Create("new.jpg")
    defer toimg.Close()

    jpeg.Encode(toimg, m, &jpeg.Options{jpeg.DefaultQuality})
}