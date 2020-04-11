package binpack

// SortRectangle type for sorting slice of rectangles. Sort backward by height
type SortRectangle []Rectangle

func (s SortRectangle) Len() int { return len(s) }

func (s SortRectangle) Less(i, j int) bool {
	return s[i].Height > s[j].Height
}

func (s SortRectangle) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
