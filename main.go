package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	a, crypto := prepareApp()
	app(a, crypto)
}

func app(a []string, crypto []string) {
	reader := bufio.NewReader(os.Stdin)
	option := 1
	for option != 5 {
		writeOptions()
		fmt.Print("Enter option's number: ")
		str, err := reader.ReadString('\n')

		f, err := strconv.Atoi(strings.TrimSpace(str))

		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("You choosed option %v \n", f)
			if f > 5 || f < 0 {
				app(a, crypto)
			}
			switch f {
			case 1:
				encryptFromText(a, crypto)
				break
			case 2:
				encryptFromFile()
				break
			case 3:
				decryptFromText()
				break
			case 4:
				decryptFromFile()
				break
			case 5:
				option = 5
			}
		}
	}
}

func encryptFromText(a []string, crypto []string) {
	var encryptedText = []string{}
	fmt.Print("Create your secret key: ")
	reader := bufio.NewReader(os.Stdin) // za pomocą tej instrukcji tworzymy czytnik

	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	key := strings.TrimSpace(str)
	keyA := strings.Split(key, "")

	keyValue := 0
	i := 0
	for i < len(keyA) {
		j := 0
		for j < len(a) {
			if a[j] == keyA[i] {
				keyValue += j
				break
			}
			j++
		}
		i++
	}

	str, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	content := strings.TrimSpace(str)
	contentA := strings.Split(content, "")

	i = 0
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
	fmt.Println(keyValue)
	fmt.Println(encryptedText)
	encryptedTextS := strings.Join(encryptedText, "")
	date := time.Now().String()
	fileName := "./crypted-" + strings.TrimSpace(date) + ".txt"
	file, err := os.Create(fileName)
	checkError(err)
	defer file.Close()
	ln, err := io.WriteString(file, encryptedTextS)
	checkError(err)
	fmt.Println(ln)

}
func encryptFromFile() {}
func decryptFromText() {}
func decryptFromFile() {}

func prepareApp() ([]string, []string) {
	a := createAlphabetSlice()
	fmt.Println(a)

	crypto := prepareCryptoAlphabetSlice(a)
	fmt.Println(crypto)
	return a, crypto
}

func writeOptions() {
	fmt.Print("What do you want to do? \n",
		"	1. Encrypt typed string, \n",
		"	2. Encrypt text from file, \n",
		"	3. Decrypt typed string, \n",
		"	4. Decrypt text from file, \n",
		"	5. Exit. \n")
}

func createAlphabetSlice() []string {
	var alphabet = []string{"A", "Ą", "B", "C", "Ć", "D", "E", "Ę", "F", "G", "H", "I", "J", "K", "L", "Ł", "M", "N", "Ń", "O", "Ó", "P", "R", "S", "Ś",
		"T", "U", "W", "Y", "Z", "Ź", "Ż"}
	return alphabet
}

func prepareCryptoAlphabetSlice(a []string) []string {
	i := 0
	j := 0
	var cryptoSlice = []string{}
	for i < 288 {
		if j <= 31 {
			cryptoSlice = append(cryptoSlice, a[j])
			j++
		} else {
			j = 0
			cryptoSlice = append(cryptoSlice, a[j])
			j++
		}
		i++
	}
	return cryptoSlice
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
