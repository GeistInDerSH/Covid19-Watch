package color

type Color int

const (
	Blue Color = iota
	Red
	Green
	Purple
	Pink
	None
)

// SetColor sets the color of the text and then clears the color after the given string
func SetColor(c Color, s string) string {
	switch c {
	case Blue:
		return string("\x1b[38;2;0;255;255m") + " " + s + string("\033[0m")
	case Red:
		return string("\x1b[38;2;226;0;46m") + " " + s + string("\033[0m")
	case Green:
		return string("\x1b[38;2;166;226;46m") + " " + s + string("\033[0m")
	case Purple:
		return string("\x1b[38;2;174;129;255m") + " " + s + string("\033[0m")
	case Pink:
		return string("\x1b[38;2;249;38;114;1m") + " " + s + string("\033[0m")
	case None:
		return string("\033[0m") + " " + s
	default:
		return string("\033[0m") + " " + s
	}
}
