package goclips

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
)

// Root clip
var Root = MovieClip{
	Children: []MovieClip{},
	Events:   ClipFuncs{},
}
var Successful = true

// holds clip events, for convienence
type ClipFuncs struct {
	OnData           func()
	OnDragOut        func()
	OnDragOver       func()
	OnEnterFrame     func() (bool, bool)
	OnKeyDown        func()
	OnKeyUp          func()
	OnKillFocus      func(newFocus any)
	OnLoad           func()
	OnMouseDown      func()
	OnMouseMove      func()
	OnMouseUp        func()
	OnPress          func()
	OnRelease        func()
	OnReleaseOutside func()
	OnRollOut        func()
	OnRollOver       func()
	OnSetFocus       func()
	OnUnload         func()
}

func CallEventFuncs(w screen.Window, eventFuncs *ClipFuncs) {
	e := w.NextEvent()
	logrus.Debugln(e)
}

func MainLoop(w screen.Window, eventFuncs *ClipFuncs) bool {
	go CallEventFuncs(w, eventFuncs)
	for {
		didError, doExit := eventFuncs.OnEnterFrame() // frame is entered, check if an error occured/if we should exit
		if didError {
			logrus.Fatal("GoClips encountered an error in running code!")
			return false
		}
		if doExit {
			return true
		}
	}
}

// Initialize goclips, note that only MainLoop and event funcs get executed after this function is called
// Events is a pointer to a ClipFuncs struct, so that events can be added and removed on the fly
func InitClips(events *ClipFuncs) bool {
	driver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(&screen.NewWindowOptions{
			Title: "GoClips",
		})
		if err != nil {
			logrus.Fatal("GoClips encountered an error!", err)
			Successful = false
			return
		}
		defer w.Release()
		Successful = MainLoop(w, events)

	})
	return Successful
}
