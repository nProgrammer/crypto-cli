package controllers

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

func App(a []string, crypto []string) {
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
				App(a, crypto)
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
	key = strings.ToUpper(key)
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
	content = strings.ToUpper(content)
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
	fileName := "./crypted/crypted-" + strings.TrimSpace(date) + ".txt"
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

func writeOptions() {
	fmt.Print("What do you want to do? \n",
		"	1. Encrypt typed string, \n",
		"	2. Encrypt text from file, \n",
		"	3. Decrypt typed string, \n",
		"	4. Decrypt text from file, \n",
		"	5. Exit. \n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
