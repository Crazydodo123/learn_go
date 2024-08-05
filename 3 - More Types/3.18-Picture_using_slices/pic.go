package main

// import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	picture := [][]uint8{}
	for y := range dy {
		picture = append(picture, []uint8{})
		for x := range dx {
			picture[y] = append(picture[y], uint8(x^y))
		}
	}
	return picture
}

func main() {
	// pic.Show(Pic)
}
