package cryptoSecPlus

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func EncryptFromText(a []string, crypto []string) {
	fmt.Print("Create your secret key: ")
	reader := bufio.NewReader(os.Stdin) // za pomocą tej instrukcji tworzymy czytnik

	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	keyValue := ValueOfKey(str, a)

	str, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(keyValue)

	encryptedText := Encrypt(str, keyValue, crypto, a)

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
	file, _ := os.Create(fileName)
	defer file.Close()
	ln, _ := io.WriteString(file, encryptedText)
	fmt.Println(ln)
}
