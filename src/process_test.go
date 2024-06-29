package src

import (
	"testing"
)

var (
	globalText string
	hiddenText string = "Secret"
	coverText  string = "This is a cover text."
)

func TestEncoding(t *testing.T) {
	globalText = Encode(coverText, []byte(hiddenText))
	if globalText == coverText {
		t.Error("Error while Encoding!")
	}
}

func TestDecoding(t *testing.T) {
	result := Decode(globalText)

	if result != hiddenText {
		t.Error("Error while Decoding!")
	}
}
