package main

import (
	"fmt"
	"strings"
)

const originalLetters string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func hashText(key int) []string {
	hashedText := strings.Split(originalLetters, "")
	preHashes := hashedText[len(hashedText)-key:]
	postHashes := hashedText[:len(hashedText)-key]
	preHashes = append(preHashes, postHashes...)
	return preHashes
}

func ceasarEncrypt(key int, text string) string {
	//generate hasing text based on key provided

	encryptedText := ""
	hashedText := hashText(key)
	strings.Map(func(r rune) rune {
		letterIndex := strings.Index(strings.ToLower(originalLetters), strings.ToLower(string(r)))

		if letterIndex != -1 {
			currentIndex := (letterIndex + len(text)) % len(text)
			encryptedText += hashedText[currentIndex]
		}

		return r
	}, text)

	return encryptedText
}

func ceasarDecrypt(key int, encryptedText string) string {
	hasedText := hashText(key)
	decryptedText := ""

	strings.Map(func(letter rune) rune {
		encryptedIndex := strings.Index(strings.Join(hasedText, ""), string([]rune{letter}))
		encryptedIndex = (encryptedIndex + len(originalLetters)) % len(originalLetters)

		decryptedText += string(originalLetters[encryptedIndex])
		fmt.Printf("%d \n", encryptedIndex)
		return 0
	}, encryptedText)
	return decryptedText
}

func main() {
	text := "Hello World"
	encryptedText := ceasarEncrypt(5, text)
	fmt.Printf("%s %s", ceasarDecrypt(5, encryptedText), encryptedText)
}
