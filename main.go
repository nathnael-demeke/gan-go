package main

import (
	"fmt"
	"strings"
)

const originalLetters string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func hashText(key int, toBeHased string) {

}

func ceasarEncrypt(key int, text string) string {
	hashedText := strings.Split(originalLetters, "")
	preHashes := hashedText[len(hashedText)-key:]
	postHashes := hashedText[:len(hashedText)-key]
	preHashes = append(preHashes, postHashes...)
	fmt.Printf("%s \n", strings.Join(preHashes, ""))
	return ""
}
func main() {
	text := "Hello World!"
	encryptedText := ceasarEncrypt(5, text)
	fmt.Printf("Encrypted Text %s", encryptedText)
}
