package formatter

import (
	"fmt"

	"github.com/fatih/color"
)

type Formatter interface {
	Print(message string)
}

type SuccessFormatter struct{}

func (s SuccessFormatter) Print(message string) {
	color.Green(fmt.Sprintf("✅ %s\n", message))
}

type ErrorFormatter struct{}

func (e ErrorFormatter) Print(message string) {
	color.Red(fmt.Sprintf("❌ %s\n", message))
}

type DefaultFormatter struct{}

func (e DefaultFormatter) Print(message string) {
	fmt.Printf("%s\n", message)
}
