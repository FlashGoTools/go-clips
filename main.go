package goclips

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/lifecycle"
)

// Root clip
var Root = MovieClip{
	Children: []MovieClip{},
}

// Initialize goclips, note that only onFrame gets executed after this function is called
func InitClips(onFrame func(w screen.Window, e any)) {
	driver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(&screen.NewWindowOptions{
			Title: "GoClips",
		})
		if err != nil {
			logrus.Fatal("GoClips encountered an error!", err)
			return
		}
		defer w.Release()
		for {
			e := w.NextEvent()
			switch e := e.(type) {
			case lifecycle.Event:
				if e.To == lifecycle.StageDead {
					return
				}
				onFrame(w, e)
			}
		}
	})
}
