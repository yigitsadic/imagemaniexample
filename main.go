package main

import (
	"github.com/fogleman/gg"
	"image/color"
	"log"
)

func main() {
	im, err := gg.LoadPNG("card_template.png")
	if err != nil {
		log.Fatalf("unable to open image due to: %q\n", err)
	}

	dc := gg.NewContextForImage(im)
	dc.Clear()
	dc.SetColor(color.Black)

	if err = dc.LoadFontFace("SyneTactile-Regular.ttf", 30); err != nil {
		log.Fatalf("unable to load font due to: %q\n", err)
	}

	dc.DrawImage(im, 0, 0)
	dc.DrawString("Welcome aboard!", 90, 290)
	dc.Clip()

	if err = dc.SavePNG("out.png"); err != nil {
		log.Fatalf("unable to save file due to: %q\n", err)
	}
}
