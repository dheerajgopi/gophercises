package cyoa

// StoryArc represents each branch in the story
type StoryArc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

// IsEndOfStory returns true if the story arc is finished
func (arc *StoryArc) IsEndOfStory() bool {
	if arc.Options == nil || len(arc.Options) == 0 {
		return true
	}

	return false
}
