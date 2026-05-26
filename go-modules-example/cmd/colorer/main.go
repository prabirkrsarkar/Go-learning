package colorer

import (
	"fmt"

	"github.com/prasarkar/go-modules-example/internal/color"
)

func main() {

	redText := "This is a red text"
	blueText := "This is a blue text"

	fmt.Println(color.TextColor(redText, color.Red))
	fmt.Println(color.TextColor(blueText, color.Blue))
}
