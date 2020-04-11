# Binpack

Package binpack provides a method that implements the algorithm for packing 2D rectangles in Go.

## Example
```go
rectangles := []binpack.Rectangle{
    {Name: "one", Width: 50, Height: 20},
    {Name: "two", Width: 40, Height: 80},
    {Name: "three", Width: 60, Height: 25},
    {Name: "four", Width: 45, Height: 78},
}

// pack rectangles in bin with width 100
height, points, err := binpack.Pack(100, rectangles)

// We got bin height = 125
// And left-top points for rectangles:
// {Name:two X:0 Y:0}
// {Name:four X:40 Y:0}
// {Name:three X:0 Y:80}
// {Name:one X:0 Y:105}
```