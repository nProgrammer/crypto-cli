package cryptoSecPlus

import (
	"fmt"
	"strings"
)

func Encrypt(str string, keyValue int, crypto []string, a []string) string {
	var encryptedText = []string{}
	content := strings.TrimSpace(str)
	content = strings.ToUpper(content)
	contentA := strings.Split(content, "")

	i := 0
	value := 0
	for i < len(contentA) {
		j := 0
		for j < len(a) {
			if a[j] == contentA[i] {
				value = j + keyValue
				newChar := crypto[value]
				encryptedText = append(encryptedText, newChar)
			}
			j++
		}
		i++
	}
	fmt.Println(encryptedText)
	encryptedTextS := strings.Join(encryptedText, "")
	return encryptedTextS
}
