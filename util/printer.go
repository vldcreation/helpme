package util

import "fmt"

const (
	RED_COLOR    = "\033[0;31m"
	BLUE_COLOR   = "\033[0;34m"
	GREEN_COLOR  = "\033[0;32m"
	YELLOW_COLOR = "\033[0;33m"
	RESET_COLOR  = "\033[0m"
)

func PrintlnGreen(a any) {
	fmt.Println(string(GREEN_COLOR), a, string(RESET_COLOR))
}

func PrintlnRed(a any) {
	fmt.Println(string(RED_COLOR), a, string(RESET_COLOR))
}

func PrintlnBlue(a any) {
	fmt.Println(string(BLUE_COLOR), a, string(RESET_COLOR))
}

func PrintlnYellow(a any) {
	fmt.Println(string(YELLOW_COLOR), a, string(RESET_COLOR))
}
