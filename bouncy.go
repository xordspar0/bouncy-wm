package main

import (
	"log"
	"time"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
)

type direction struct {
	x int16
	y int16
}

func main() {
	X, err := xgb.NewConn()
	if err != nil {
		log.Fatalln(err)
	}

	screen := xproto.Setup(X).DefaultScreen(X)
	root := screen.Root

	geometry, err := xproto.GetGeometry(X, xproto.Drawable(root)).Reply()
	screenWidth := geometry.Width
	screenHeight := geometry.Height

	results, err := xproto.QueryTree(X, root).Reply()
	if err != nil {
		log.Fatalln(err)
	}

	windows := results.Children
	directions := make([]direction, len(windows))
	for i := range directions {
		directions[i].x = 1
		directions[i].y = 1
	}

	for {
		for i, window := range windows {
			geometry, err := xproto.GetGeometry(X, xproto.Drawable(window)).Reply()
			if err != nil {
				log.Println(err)
			}

			if geometry.X < 0 || geometry.X+int16(geometry.Width) >= int16(screenWidth) {
				directions[i].x *= -1
			}
			if geometry.Y < 0 || geometry.Y+int16(geometry.Height) >= int16(screenHeight) {
				directions[i].y *= -1
			}

			newX := uint32(geometry.X + directions[i].x)
			newY := uint32(geometry.Y + directions[i].y)

			xproto.ConfigureWindow(X, window, xproto.ConfigWindowX|xproto.ConfigWindowY, []uint32{newX, newY})
		}

		time.Sleep(16 * time.Millisecond)
	}
}
