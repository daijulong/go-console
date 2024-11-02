package console

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func PrintColor(s string, fg string, bg ...string) {
	c := &ConColor{
		fgColor: fg,
		text:    s,
	}
	if len(bg) > 0 {
		c.bgColor = bg[0]
	}
	if c.fgColor != "" || c.bgColor != "" {
		c.reset = Reset
	}
	c.Print()
}

func PrintBlack(s string, bg ...string) {
	PrintColor(s, FgBLack, bg...)
}

func PrintRed(s string, bg ...string) {
	PrintColor(s, FgRed, bg...)
}

func PrintGreen(s string, bg ...string) {
	PrintColor(s, FgGreen, bg...)
}

func PrintYellow(s string, bg ...string) {
	PrintColor(s, FgYellow, bg...)
}

func PrintBlue(s string, bg ...string) {
	PrintColor(s, FgBlue, bg...)
}

func PrintMagenta(s string, bg ...string) {
	PrintColor(s, FgMagenta, bg...)
}

func PrintCyan(s string, bg ...string) {
	PrintColor(s, FgCyan, bg...)
}

func PrintWhite(s string, bg ...string) {
	PrintColor(s, FgWhite, bg...)
}

func Line(s string) {
	fmt.Println(s)
}

func Lines(s []string) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func BlankLines(n ...int) {
	lines := 1
	if len(n) > 0 {
		lines = n[0]
	}
	for i := 0; i < lines; i++ {
		fmt.Println()
	}
}

func Info(s string) {
	Line(s)
}

func Warning(s string) {
	PrintYellow(s)
}

func Error(s string) {
	PrintRed(s, BgYellow)
}

func ErrorExit(s string) {
	Error(s)
	os.Exit(0)
}

func Table(titles []string, data [][]string) {
	colWidths := make([]int, len(titles))
	for i, title := range titles {
		colWidths[i] = len(title)
	}
	for _, row := range data {
		for i, cell := range row {
			if len(cell) > colWidths[i] {
				colWidths[i] = len(cell)
			}
		}
	}

	tablePrintSeparator(colWidths)
	tablePrintRow(titles, colWidths)
	tablePrintSeparator(colWidths)

	for _, row := range data {
		tablePrintRow(row, colWidths)
	}

	tablePrintSeparator(colWidths)
}

func tableStringWidth(s string) int {
	width := 0
	for _, r := range s {
		if utf8.RuneLen(r) > 1 {
			width += 2
		} else {
			width += 1
		}
	}
	return width
}

func tablePadString(s string, width int) string {
	padding := width - tableStringWidth(s)
	return s + strings.Repeat(" ", padding)
}

func tablePrintRow(row []string, colWidths []int) {
	line := "|"
	for i, cell := range row {
		line += fmt.Sprintf(" %s |", tablePadString(cell, colWidths[i]))
	}
	fmt.Println(line)
}

func tablePrintSeparator(colWidths []int) {
	line := "+"
	for _, width := range colWidths {
		line += strings.Repeat("-", width+2) + "+"
	}
	fmt.Println(line)
}
