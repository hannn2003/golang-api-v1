package methods
type rect struct {
	width, height int
} 

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*(r.width + r.height)
}

