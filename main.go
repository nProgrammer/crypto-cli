package main

import "fmt"

func main() {
	a, crypto := prepareApp()
	app(a, crypto)
}

func app(a []string, crypto []string) {
	option := 1
	for option != 5 {
		writeOptions()
		option = 5
	}
}

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
