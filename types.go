package goclips

// Interface for movieclips
type MovieClipIface interface {
	CreateEmptyMovieClip()
	GetChild()
}

// MovieClip struct, do not create this yourself, instead use goclips.Root.CreateEmptyMovieClip
type MovieClip struct {
	Children     []any
	InstanceName string
	Depth        int
	currentFrame int
}

// Creates an empty movie clip as a child of the current clip. Returns pointer to created clip.
// P.S. Yes, I know, you aren't supposed to use "this" as a reciever, but I'm trying to remain
// as close to the original flash version as possible
func (this MovieClip) CreateEmptyMovieClip(instanceName string, depth int) *any {
	this.Children = append(this.Children, MovieClip{
		Children:     []any{},
		InstanceName: instanceName,
		Depth:        depth,
	})
	return &this.Children[len(this.Children)-1]
}

// Gets a child clip
func (this MovieClip) GetChild(childIndex int) *any {
	return &this.Children[childIndex]
}

func (this MovieClip) NextFrame() {
	// at some point actually display this frame
	this.currentFrame++
}
