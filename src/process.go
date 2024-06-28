package src

import (
	"strings"
)

const (
	ZeroWidthSpace     = '\u200B' // Represents binary 0
	ZeroWidthNonJoiner = '\u200C' // Represents binary 1
)

func Encode(coverText string, hiddenText []byte) string {
	var encodedText strings.Builder
	encodedText.WriteString(coverText)

	for _, char := range hiddenText {
		for i := 7; i >= 0; i-- {

			bit := (char >> i) & 1

			if bit == 1 {
				encodedText.WriteRune(ZeroWidthNonJoiner)
			} else {
				encodedText.WriteRune(ZeroWidthSpace)
			}
		}
	}
	return encodedText.String()
}

func Decode(encodedText string) string {
	var hiddenText strings.Builder
	bits := 0
	bitCount := 0

	for _, char := range encodedText {
		if char == ZeroWidthSpace {
			bits = bits << 1 // Add 0
			bitCount++
		} else if char == ZeroWidthNonJoiner {
			bits = (bits << 1) | 1 // Add 1
			bitCount++
		}

		if bitCount == 8 {
			hiddenText.WriteByte(byte(bits))
			bits = 0
			bitCount = 0
		}
	}

	return hiddenText.String()
}
