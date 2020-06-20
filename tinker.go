package main

import (
	"fmt"
	"wulfheart/brahms/viz"
)

func main(){
	keypoints := []string{
		"#40E0D0", "#FF8C00", "#FF0080",
	}

	palette, err := viz.MakePalette(10, keypoints...)
	if err != nil {
		panic(err)
	}

	for _, p := range palette {
		fmt.Println(p.Hex())
	}
	// h := 1024
	// w := 40
	//
	// keypoints := []string{
	// 	"#40E0D0", "#FF8C00", "#FF0080",
	// }
	//
	// img := image.NewRGBA(image.Rect(0, 0, w, h))
	//
	// palette, err := viz.MakePalette(h, keypoints...)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// for y := 0; y < h -1; y++ {
	// 	draw.Draw(img, image.Rect(0, y, w, y + 1), &image.Uniform{C: palette[y]}, image.Point{},draw.Src)
	// }
	//
	// outpng, err := os.Create("gradientgen.png")
	// if err != nil {
	// 	panic("Error storing png: " + err.Error())
	// }
	// defer outpng.Close()
	//
	// png.Encode(outpng, img)
}
