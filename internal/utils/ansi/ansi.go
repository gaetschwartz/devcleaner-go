package ansi

import (
	"fmt"
	"strings"
)

// First byte is a mask to tell if following bytes have a value or not
// # Byte 1:
//
// 0b10000000 - unused
// 0b01000000 - whether byte 1 has a value or not
// 0b00100000 - whether byte 2 has a value or not
// 0b00010000 - whether byte 3 has a value or not
// 0b00001000 - whether byte 4 has a value or not
// 0b00000100 - whether byte 5 has a value or not
// 0b00000010 - whether byte 6 has a value or not
// 0b00000001 - whether byte 7 has a value or not
//
// # Byte 2-8:
//
// code (0-255)
type Code uint64
type Str string

func (c Code) String() string {
	if c&(0xff<<56) == 0 {
		return ""
	}
	var b strings.Builder
	b.WriteString("\033[")
	has_written := false
	for i := 1; i < 8; i++ {
		if has_written {
			b.WriteString(";")
		}
		if (c & (1 << (63 - i))) != 0 {
			// byte i has a value
			bytei := (c >> (8 * (7 - i))) & 0xff
			// fmt.Println("Byte", i, "has value", int(bytei))
			b.WriteString(fmt.Sprintf("%d", bytei))
			has_written = true
		} else {
			has_written = false
		}
	}
	b.WriteByte('m')

	return b.String()
}

func Rgb(r, g, b int) Code {
	return rgbColorFgBase | rgb(r, g, b)
}

func BgRgb(r, g, b int) Code {
	return rgbColorBgBase | rgb(r, g, b)
}

func rgb(r, g, b int) Code {
	return (Code(r&0xff)<<16 | Code(g&0xff)<<8 | Code(b&0xff))
}

func (s Str) Prefix(colors ...Code) Str {
	var b strings.Builder
	for _, c := range colors {
		b.WriteString(c.String())
	}
	b.WriteString(string(s))
	return s
}

func (s Str) Suffix(colors ...Code) Str {
	var b strings.Builder
	b.WriteString(string(s))
	for _, c := range colors {
		b.WriteString(c.String())
	}
	return s
}

func (s Str) Style(codes ...Code) Str {
	var b strings.Builder
	for _, c := range codes {
		b.WriteString(c.String())
	}
	b.WriteString(string(s))
	b.WriteString(resetAnsiCodeString)
	return Str(b.String())
}

func (s Str) String() string {
	return string(s)
}

const (
	singleColorCodeBase = Code(0x01_00_00_00_00_00_00_00)
	resetAnsiCodeString = "\033[0m"
	Empty               = Code(0x00)
	Reset               = singleColorCodeBase | 0x00
	Bold                = singleColorCodeBase | 0x01
	Dim                 = singleColorCodeBase | 0x02
	Italic              = singleColorCodeBase | 0x03
	Underline           = singleColorCodeBase | 0x04
	Blink               = singleColorCodeBase | 0x05
	Hidden              = singleColorCodeBase | 0x08
	Black               = singleColorCodeBase | 30
	BlackBg             = Black + 10
	Red                 = singleColorCodeBase | 31
	RedBg               = Red + 10
	Green               = singleColorCodeBase | 32
	GreenBg             = Green + 10
	Yellow              = singleColorCodeBase | 33
	YellowBg            = Yellow + 10
	Blue                = singleColorCodeBase | 34
	BlueBg              = Blue + 10
	Magenta             = singleColorCodeBase | 35
	MagentaBg           = Magenta + 10
	Cyan                = singleColorCodeBase | 36
	CyanBg              = Cyan + 10
	White               = singleColorCodeBase | 37
	WhiteBg             = White + 10
	BrightBlack         = singleColorCodeBase | 90
	BrightBlackBg       = BrightBlack + 10
	BrightRed           = singleColorCodeBase | 91
	BrightRedBg         = BrightRed + 10
	BrightGreen         = singleColorCodeBase | 92
	BrightGreenBg       = BrightGreen + 10
	BrightYellow        = singleColorCodeBase | 93
	BrightYellowBg      = BrightYellow + 10
	BrightBlue          = singleColorCodeBase | 94
	BrightBlueBg        = BrightBlue + 10
	BrightMagenta       = singleColorCodeBase | 95
	BrightMagentaBg     = BrightMagenta + 10
	BrightCyan          = singleColorCodeBase | 96
	BrightCyanBg        = BrightCyan + 10
	BrightWhite         = singleColorCodeBase | 97
	BrightWhiteBg       = BrightWhite + 10
	rgbColorFgBase      = Code(0x1f_00_00_26_02_00_00_00)
	rgbColorBgBase      = Code(0x1f_00_00_30_02_00_00_00)
)
