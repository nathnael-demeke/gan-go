package main

import (
	"fmt"
	"strings"
)

const originalLetters string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func hashText(key int, toBeHased string) {

}

func ceasarEncrypt(key int, text string) string {
	//generate hasing text based on key provided
	hashedText := strings.Split(originalLetters, "")
	preHashes := hashedText[len(hashedText)-key:]
	postHashes := hashedText[:len(hashedText)-key]
	preHashes = append(preHashes, postHashes...)

	encryptedText := ""

	strings.Map(func(r rune) rune {
		letterIndex := strings.Index(strings.ToLower(originalLetters), strings.ToLower(string(r)))

		if letterIndex != -1 {
			currentIndex := (letterIndex + len(text)) % len(text)
			encryptedText += preHashes[currentIndex]
		}

		return r
	}, text)

	return encryptedText
}
func main() {
	text := "Hello World"
	encryptedText := ceasarEncrypt(5, text)
	fmt.Printf("%s", encryptedText)
}
