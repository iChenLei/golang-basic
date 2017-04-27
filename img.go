package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	reader, err := os.Open("taiyan.jpg")
	if err != nil {
		log.Fatal(err)
	}
	m, v, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
		fmt.Println("something error...")
	}
	// bounds := m.Bounds()
	fmt.Println("Read a img from dirname...")
	bounds := m.Bounds()

	gray := m.NewGray(bounds)

	imagefile, error := os.Create("img_gray.jpg")
	defer imagefile.Close()
	if error != nil {
		log.Fatal(error)
	}

	erro := jpeg.Encode(imagefile, gray)

	if erro != nil {
		log.Fatal(erro)
	}

	// fmt.Println(bounds)

	// r, g, b, a := m.At(50, 50).RGBA()

	// fmt.Printf("width %v,height %v", m.Bounds().Max.X, m.Bounds().Max.Y)
	// fmt.Printf("At(50,50) rgbs is %v(r) %v(g) %v(b) %v(a)\n", r, g, b, a)

	// rr, gg, bb := r/a*255, g/a*255, b/a*255
	// fmt.Printf("%v %v %v ..",rr,gg,bb)
	fmt.Printf(v)

	defer reader.Close()
}
