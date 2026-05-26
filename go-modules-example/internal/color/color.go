package color

import "fmt"

type Color struct {
	Value int
}

// Declaring variables in a var block allows us to group related variables together,
// making the code more organized and easier to read. In this case, we are declaring variables
// for ANSI foreground color codes used for terminal text formatting.

var (
	Black   = Color{Value: 30}
	Red     = Color{Value: 31}
	Green   = Color{Value: 32}
	Yellow  = Color{Value: 33}
	Blue    = Color{Value: 34}
	Magenta = Color{Value: 35}
	Cyan    = Color{Value: 36}
	White   = Color{Value: 37}
)

// // ANSI background color codes used for terminal text formatting.
var (
	BgBlack   = Color{Value: 40}
	BgRed     = Color{Value: 41}
	BgGreen   = Color{Value: 42}
	BgYellow  = Color{Value: 43}
	BgBlue    = Color{Value: 44}
	BgMagenta = Color{Value: 45}
	BgCyan    = Color{Value: 46}
	BgWhite   = Color{Value: 47}
)

var (
	Bold      = Color{Value: 1}
	Underline = Color{Value: 4}
)

func TextColor(text string, colors ...Color) string {
	if len(colors) == 0 {
		return text
	}

	var colorCodes []int
	for _, color := range colors {
		colorCodes = append(colorCodes, color.Value)
	}

	return fmt.Sprintf("\033[%sm%s\033[0m", join(colorCodes, ";"), text)
}
