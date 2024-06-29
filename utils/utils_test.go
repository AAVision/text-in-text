package utils

import "testing"

var (
	enctyptedText string
	password      string = "aavision"
	plainText     string = "Secret"
)

func TestEncrypt(t *testing.T) {
	enctyptedText, _ = Encrypt(plainText, password)
	if enctyptedText == "" {
		t.Error("Error while Encrypting!")
	}
}

func TestDecrypt(t *testing.T) {
	output, _ := Decrypt(enctyptedText, password)
	if output != plainText {
		t.Error("Error while Decrypting!")
	}
}
