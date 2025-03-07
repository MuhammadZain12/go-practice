package structs

type Rectangle struct {
	Width  int
	Height int
}

func Area(rectangle Rectangle) int {
	return rectangle.Height * rectangle.Width
}
