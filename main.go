package main

import (
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/mewkiz/pkg/imgutil"
	"github.com/pkg/errors"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	if err := f(); err != nil {
		log.Fatalf("%+v", err)
	}
}

func f() error {
	pic, err := loadPicture("penguin.jpeg")
	if err != nil {
		return errors.WithStack(err)
	}
	flipper, err := loadPicture("flipper.jpeg")
	if err != nil {
		return errors.WithStack(err)
	}
	picBounds := pic.Bounds()
	sprite := pixel.NewSprite(pic, picBounds)
	cfg := pixelgl.WindowConfig{
		Title:  "penguin",
		Bounds: picBounds,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		return errors.WithStack(err)
	}
	for !win.Closed() {
		win.Update()
		if win.JustPressed(pixelgl.MouseButton1) {
			sprite.Set(flipper, picBounds)
		} else if win.JustPressed(pixelgl.MouseButton2) {
			sprite.Set(pic, picBounds)
		}
		sprite.Draw(win, pixel.IM.Moved(pic.Bounds().Center()))
	}
	return nil
}

func loadPicture(path string) (pixel.Picture, error) {
	img, err := imgutil.ReadFile(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	pic := pixel.PictureDataFromImage(img)
	return pic, nil
}
