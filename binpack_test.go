package binpack

import "testing"

func TestPack(t *testing.T) {
	rects := []Rectangle{
		{Name: "one", Width: 50, Height: 35},
		{Name: "two", Width: 40, Height: 80},
		{Name: "three", Width: 41, Height: 30},
		{Name: "four", Width: 45, Height: 78},
		{Name: "five", Width: 54, Height: 26},
		{Name: "six", Width: 54, Height: 54},
		{Name: "seven", Width: 21, Height: 50},
		{Name: "eight", Width: 11, Height: 15},
		{Name: "nine", Width: 55, Height: 35},
		{Name: "ten", Width: 40, Height: 40},
	}

	height, points, err := Pack(180, rects)

	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}

	var expHeight = 120
	if height != expHeight {
		t.Errorf("height must be %d, got %d", expHeight, height)
	}

	if len(points) != len(rects) {
		t.Errorf("result slice has wrong len, expected %d, got %d", len(rects), len(points))
	}
}
