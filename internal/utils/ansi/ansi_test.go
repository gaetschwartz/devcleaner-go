package ansi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRgbColor(t *testing.T) {
	assert.Equal(t, "\033[38;2;31;0;0m", Rgb(31, 0, 0).String())
	assert.Equal(t, "\033[48;2;31;0;0m", BgRgb(31, 0, 0).String())
	assert.Equal(t, "\033[38;2;0;31;0m", Rgb(0, 31, 0).String())
	assert.Equal(t, "\033[48;2;0;31;0m", BgRgb(0, 31, 0).String())
	assert.Equal(t, "\033[38;2;0;0;31m", Rgb(0, 0, 31).String())
	assert.Equal(t, "\033[48;2;0;0;31m", BgRgb(0, 0, 31).String())
}
func TestCodes(t *testing.T) {
	assert.Equal(t, "\033[0m", Reset.String())
	assert.Equal(t, "\033[1m", Bold.String())
	assert.Equal(t, "\033[2m", Dim.String())
	assert.Equal(t, "\033[3m", Italic.String())
	assert.Equal(t, "\033[4m", Underline.String())
	assert.Equal(t, "\033[5m", Blink.String())
	assert.Equal(t, "\033[8m", Hidden.String())
}

func TestAnsiString(t *testing.T) {
	assert.Equal(t,
		Str("\033[31m\033[1mHello World\033[0m"),
		Str("Hello World").Style(Red, Bold),
	)
}
