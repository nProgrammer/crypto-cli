package cryptoSecPlus

import (
	"fmt"
	"strings"
)

func Encrypt(str string, keyValue int, keyF int, keyS int, crypto []string, a []string) string {
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
				if j == 0 {
					value = j + keyF
					newChar := crypto[value]
					encryptedText = append(encryptedText, newChar)
				} else if j == len(a)-1 {
					value = j + keyS
					newChar := crypto[value]
					encryptedText = append(encryptedText, newChar)
				} else {
					value = j + keyValue
					newChar := crypto[value]
					encryptedText = append(encryptedText, newChar)
				}
			}
			j++
		}
		i++
	}
	fmt.Println(encryptedText)
	encryptedTextS := strings.Join(encryptedText, "")
	return encryptedTextS
}
