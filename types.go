package goclips

type EventID int

const (
	OnDataID EventID = iota
	OnDragOutID
	OnDragOverID
	OnEnterFrameID
	OnKeyDownID
	OnKeyUpID
	OnKillFocusID
	OnLoadID
	OnMouseDownID
	OnMouseMoveID
	OnMouseUpID
	OnPressID
	OnReleaseID
	OnRollOutID
	OnRollOverID
	OnSetFocusID
	OnUnloadID
	OnReleaseOutsideID
)

// Interface for movieclips
type MovieClipIface interface {
	CreateEmptyMovieClip()
	GetChild()
	NextFrame()
	PassEventCallToChildren()
}

// MovieClip struct, do not create this yourself, instead use (MovieClip).CreateEmptyMovieClip
type MovieClip struct {
	Children     []MovieClip
	Events       ClipFuncs
	InstanceName string
	Depth        int
	currentFrame int
}

// Creates an empty movie clip as a child of the current clip. Returns pointer to created clip.
// P.S. Yes, I know, you aren't supposed to use "this" as a reciever, but I'm trying to remain
// as close to the original flash version as possible
func (this MovieClip) CreateEmptyMovieClip(instanceName string, depth int) *MovieClip {
	this.Children = append(this.Children, MovieClip{
		Children:     []MovieClip{},
		InstanceName: instanceName,
		Depth:        depth,
	})
	return &this.Children[len(this.Children)-1]
}

// Gets a child clip
func (this MovieClip) GetChild(childIndex int) *MovieClip {
	return &this.Children[childIndex]
}

func (this MovieClip) NextFrame() {
	// at some point actually display this frame
	this.currentFrame++
}

func (this MovieClip) AttachMovie(id string, name string, depth int) *MovieClip {
	return &MovieClip{}
}

func (this MovieClip) PassEventCallToChildren(event EventID, data any) (bool, bool) {
	var didError = false
	var doExit = false
	for _, child := range this.Children {
		switch event {
		case OnDataID:
			child.Events.OnData()
		case OnDragOutID:
			child.Events.OnDragOut()
		case OnDragOverID:
			child.Events.OnDragOver()
		case OnEnterFrameID:
			didError, doExit = child.Events.OnEnterFrame()
		case OnKeyDownID:
			child.Events.OnKeyDown()
		case OnKeyUpID:
			child.Events.OnKeyUp()
		case OnKillFocusID:
			child.Events.OnKillFocus(data)
		case OnLoadID:
			child.Events.OnLoad()
		case OnMouseDownID:
			child.Events.OnMouseDown()
		case OnMouseMoveID:
			child.Events.OnMouseMove()
		case OnMouseUpID:
			child.Events.OnMouseUp()
		case OnPressID:
			child.Events.OnPress()
		case OnReleaseID:
			child.Events.OnRelease()
		case OnReleaseOutsideID:
			child.Events.OnReleaseOutside()
		case OnRollOutID:
			child.Events.OnRollOut()
		case OnRollOverID:
			child.Events.OnRollOver()
		case OnSetFocusID:
			child.Events.OnSetFocus()
		case OnUnloadID:
			child.Events.OnUnload()
		}
	}
	return didError, doExit
}
