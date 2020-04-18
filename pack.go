package binpack

func newPack(width int) *bin {
	return &bin{
		width: width,
	}
}

type bin struct {
	width  int
	levels []*level
}

func (p *bin) height() int {
	var h int
	for _, l := range p.levels {
		h += l.height
	}
	return h
}

func (p *bin) points() []Point {
	var points []Point
	for _, l := range p.levels {
		for _, r := range l.top {
			points = append(points, r.point)
		}

		for _, r := range l.bottom {
			points = append(points, r.point)
		}
	}
	return points
}

func (p *bin) addRectangle(r Rectangle) error {
	if r.Width > p.width {
		return ErrCannotFit
	}

	var packed bool
	for _, l := range p.levels {
		if packed = l.addRectangle(r); packed {
			break
		}
	}

	if !packed {
		p.addLevel(r)
	}

	return nil
}

func (p *bin) addLevel(r Rectangle) {
	height := p.height()
	widthLeft := p.width - r.Width
	p.levels = append(p.levels, &level{
		startFrom: height,
		width:     p.width,
		height:    r.Height,
		top: []packedRectangle{
			{
				rectangle: r,
				point: Point{
					Name: r.Name,
					X:    0,
					Y:    height,
				},
			},
		},
		topWidthLeft:    widthLeft,
		bottomWidthLeft: widthLeft,
	})
}

type level struct {
	startFrom       int
	width           int
	height          int
	top             []packedRectangle
	bottom          []packedRectangle
	topWidthLeft    int
	bottomWidthLeft int
}

func (l *level) addRectangle(r Rectangle) bool {
	// try to insert in top of level
	if l.topWidthLeft >= r.Width {
		newPackedRectangle := packedRectangle{
			rectangle: r,
			point: Point{
				Name: r.Name,
				X:    sumWidth(l.top),
				Y:    l.startFrom,
			},
		}

		var collapsed bool
		for _, rect := range l.bottom {
			if collapsed = isCollapsedLTRB(newPackedRectangle, rect); collapsed {
				break
			}
		}

		if !collapsed {
			l.top = append(l.top, newPackedRectangle)
			l.topWidthLeft -= r.Width
			return true
		}
	}

	// try to insert in bottom of level
	if l.bottomWidthLeft >= r.Width {
		newPackedRectangle := packedRectangle{
			rectangle: r,
			point: Point{
				Name: r.Name,
				X:    l.width - sumWidth(l.bottom) - r.Width,
				Y:    l.startFrom + l.height - r.Height,
			},
		}

		var collapsed bool
		for _, rect := range l.top {
			if collapsed = isCollapsedLTRB(rect, newPackedRectangle); collapsed {
				break
			}
		}

		if !collapsed {
			l.bottom = append(l.bottom, newPackedRectangle)
			l.bottomWidthLeft -= r.Width
			return true
		}
	}

	return false
}

type packedRectangle struct {
	rectangle Rectangle
	point     Point
}

// isCollapsedLTRB - is rectangle in left-top collapsed with rectangle in right-bottom
func isCollapsedLTRB(lt packedRectangle, rb packedRectangle) bool {
	// right-bottom point of left-top rectangle
	pointX := lt.point.X + lt.rectangle.Width
	pointY := lt.point.Y + lt.rectangle.Height

	if pointX > rb.point.X && pointX <= rb.point.X+rb.rectangle.Width &&
		pointY > rb.point.Y && pointY <= rb.point.Y+rb.rectangle.Height {
		return true
	}

	return false
}

func sumWidth(rects []packedRectangle) int {
	var w int
	for _, r := range rects {
		w += r.rectangle.Width
	}
	return w
}
