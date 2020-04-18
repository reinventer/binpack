// Package binpack provides a method that implements the algorithm for packing 2D rectangles.
package binpack

import "sort"

// Rectangle - type for rectangle with name, width and height
type Rectangle struct {
	Name          string
	Width, Height int
}

// Point - type for left-top point of packed rectangle
type Point struct {
	Name string
	X, Y int
}

// Pack places a rectangle in a bin of width w
func Pack(width int, rectangles []Rectangle) (int, []Point, error) {
	sorted := make([]Rectangle, len(rectangles))
	copy(sorted, rectangles)
	sort.Sort(SortRectangle(sorted))

	bin := &bin{
		width: width,
	}

	for _, r := range sorted {
		if err := bin.addRectangle(r); err != nil {
			return 0, nil, err
		}
	}

	return bin.height(), bin.points(), nil
}
