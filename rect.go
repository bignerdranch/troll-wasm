package ebitentest

type Rectangle struct {
	PosX   float64
	PosY   float64
	Width  float64
	Height float64
}

func (r *Rectangle) DimensionsInt() (int, int) {
	return int(r.Width), int(r.Height)
}

func (r *Rectangle) PositionInt() (int, int) {
	return int(r.PosX), int(r.PosY)
}

func (r *Rectangle) PointCollides(x, y float64) bool {
	x2 := r.PosX + r.Width
	y2 := r.PosY + r.Height

	return (r.PosX < x && x < x2) && (r.PosY < y && y < y2)
}
