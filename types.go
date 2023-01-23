package goclips

type MovieClipIface interface {
	CreateEmptyMovieClip()
}

type MovieClip struct {
	Children []any
}

func (this MovieClip) CreateEmptyMovieClip (instanceName string, depth int) *any {
	// Creates an empty movie clip as a child of the current clip. Returns pointer to created clip.
	this.Children = append(this.Children, MovieClip{[]any{}})
	return &this.Children[len(this.Children) - 1]
}