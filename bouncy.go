package main

import (
	"log"
	"os"
	"sync"
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

	xproto.ChangeWindowAttributes(X, root, xproto.CwEventMask, []uint32{xproto.EventMaskSubstructureNotify})

	var windowsMutex sync.Mutex
	windows := make(map[xproto.Window]*direction)

	go func() {
		for {
			ev, err := X.WaitForEvent()
			if err != nil {
				log.Println(err)
				continue
			}

			if ev == nil {
				log.Println("X connection closed by server")
				os.Exit(1)
			}

			switch ev := ev.(type) {
			case xproto.CreateNotifyEvent:
				log.Println(ev)
				windowsMutex.Lock()
				windows[ev.Window] = &direction{x: 1, y: 1}
				windowsMutex.Unlock()
			case xproto.DestroyNotifyEvent:
				log.Println(ev)
				windowsMutex.Lock()
				delete(windows, ev.Window)
				windowsMutex.Unlock()
			}
		}
	}()

	geometry, err := xproto.GetGeometry(X, xproto.Drawable(root)).Reply()
	screenWidth := geometry.Width
	screenHeight := geometry.Height

	results, err := xproto.QueryTree(X, root).Reply()
	if err != nil {
		log.Fatalln(err)
	}

	for _, window := range results.Children {
		windows[window] = &direction{x:1, y: 1}
	}

	for {
		for window, direction := range windows {
			geometry, err := xproto.GetGeometry(X, xproto.Drawable(window)).Reply()
			if err != nil {
				log.Println(err)
			}

			if geometry.X < 0 || geometry.X+int16(geometry.Width) >= int16(screenWidth) {
				windowsMutex.Lock()
				direction.x *= -1
				windowsMutex.Unlock()
			}
			if geometry.Y < 0 || geometry.Y+int16(geometry.Height) >= int16(screenHeight) {
				windowsMutex.Lock()
				direction.y *= -1
				windowsMutex.Unlock()
			}

			newX := uint32(geometry.X + direction.x)
			newY := uint32(geometry.Y + direction.y)

			xproto.ConfigureWindow(X, window, xproto.ConfigWindowX|xproto.ConfigWindowY, []uint32{newX, newY})
		}

		time.Sleep(16 * time.Millisecond)
	}
}
