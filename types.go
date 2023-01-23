package goclips

type MovieClipIface interface {
	CreateEmptyMovieClip()
}

type MovieClip struct {
	Children []any
}

// Creates an empty movie clip as a child of the current clip. Returns pointer to created clip.
// P.S. Yes, I know, you aren't supposed to use "this" as a reciever, but I'm trying to remain
// as close to the original flash version as possible
func (this MovieClip) CreateEmptyMovieClip(instanceName string, depth int) *any {
	this.Children = append(this.Children, MovieClip{[]any{}})
	return &this.Children[len(this.Children)-1]
}
