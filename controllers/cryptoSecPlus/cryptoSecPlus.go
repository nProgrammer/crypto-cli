package cryptoSecPlus

import (
	"bufio"
	"crypto-cli/errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func EncryptFromText(a []string, crypto []string) {
	fmt.Print("Create your secret key: ")
	reader := bufio.NewReader(os.Stdin)

	str, err := reader.ReadString('\n')
	errors.CheckError(err)

	keyValue, keyF, keyS := ValueOfKey(str, a)

	str, err = reader.ReadString('\n')
	errors.CheckError(err)

	fmt.Println(keyValue)

	encryptedText := Encrypt(str, keyValue, keyF, keyS, crypto, a)

	Save(encryptedText)

}
func EncryptFromFile() {
	fmt.Println("This function will be added in the future")
}
func DecryptFromText() {
	fmt.Println("This function will be added in the future")
}
func DecryptFromFile() {
	fmt.Println("This function will be added in the future")
}

func Save(encryptedText string) {
	date := time.Now().String()
	fileName := "./crypted/crypted-" + strings.TrimSpace(date) + ".txt"
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	ln, err := io.WriteString(file, encryptedText)
	errors.CheckError(err)

	fmt.Println(ln)
}
