package console

import "fmt"

const (
	FgBLack   = "\033[30m"
	FgRed     = "\033[31m"
	FgGreen   = "\033[32m"
	FgYellow  = "\033[33m"
	FgBlue    = "\033[34m"
	FgMagenta = "\033[35m"
	FgCyan    = "\033[36m"
	FgWhite   = "\033[37m"

	BgBlack   = "\033[40m"
	BgRed     = "\033[41m"
	BgGreen   = "\033[42m"
	BgYellow  = "\033[43m"
	BgBlue    = "\033[44m"
	BgMagenta = "\033[45m"
	BgCyan    = "\033[46m"
	BgWhite   = "\033[47m"

	Reset = "\033[0m"
)

func Color(s string) *ConColor {
	return &ConColor{text: s}
}

type ConColor struct {
	fgColor string
	bgColor string
	reset   string
	text    string
}

func (c *ConColor) fg(color string) *ConColor {
	c.fgColor = color
	c.reset = Reset
	return c
}

func (c *ConColor) bg(color string) *ConColor {
	c.bgColor = color
	c.reset = Reset
	return c
}

func (c *ConColor) Print() {
	fmt.Printf("%s%s%s%s\n", c.fgColor, c.bgColor, c.text, c.reset)
}

func (c *ConColor) Black() *ConColor {
	return c.fg(FgBLack)
}

func (c *ConColor) Red() *ConColor {
	return c.fg(FgRed)
}

func (c *ConColor) Green() *ConColor {
	return c.fg(FgGreen)
}

func (c *ConColor) Yellow() *ConColor {
	return c.fg(FgYellow)
}

func (c *ConColor) Blue() *ConColor {
	return c.fg(FgBlue)
}

func (c *ConColor) Magenta() *ConColor {
	return c.fg(FgMagenta)
}

func (c *ConColor) Cyan() *ConColor {
	return c.fg(FgCyan)
}

func (c *ConColor) White() *ConColor {
	return c.fg(FgWhite)
}

func (c *ConColor) OnBlack() *ConColor {
	return c.bg(BgBlack)
}

func (c *ConColor) OnRed() *ConColor {
	return c.bg(BgRed)
}

func (c *ConColor) OnGreen() *ConColor {
	return c.bg(BgGreen)
}

func (c *ConColor) OnYellow() *ConColor {
	return c.bg(BgYellow)
}

func (c *ConColor) OnBlue() *ConColor {
	return c.bg(BgBlue)
}

func (c *ConColor) OnMagenta() *ConColor {
	return c.bg(BgMagenta)
}

func (c *ConColor) OnCyan() *ConColor {
	return c.bg(BgCyan)
}

func (c *ConColor) OnWhite() *ConColor {
	return c.bg(BgWhite)
}
