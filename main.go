package main

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/speaker"
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
	// Load images.
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
	// Init window.
	cfg := pixelgl.WindowConfig{
		Title:  "Alanso says Vraaaarrra :)",
		Bounds: picBounds,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		return errors.WithStack(err)
	}
	// Draw loop.
	played := false
	for !win.Closed() {
		win.Update()
		if win.JustPressed(pixelgl.MouseButton1) {
			sprite.Set(flipper, picBounds)
			// Play sound.
			if !played {
				f := func() {
					played = false
				}
				snd, err := loadSound("macaroni.flac")
				if err != nil {
					return errors.WithStack(err)
				}
				callback := beep.Callback(f)
				seq := beep.Seq(snd, callback)
				speaker.Play(seq)
			}
			played = true
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

func loadSound(path string) (beep.StreamSeekCloser, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	s, format, err := flac.Decode(f)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	return s, nil
}
