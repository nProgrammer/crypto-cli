package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	a, crypto := prepareApp()
	app(a, crypto)
}

func app(a []string, crypto []string) {
	reader := bufio.NewReader(os.Stdin) // za pomocą tej instrukcji tworzymy czytnik
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
				encryptFromText()
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

func encryptFromText() {}
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
