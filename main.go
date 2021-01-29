package main

import "fmt"

func main() {
	a := createAlphabetSlice()
	fmt.Println(a)

	crypto := prepareCryptoAlphabetSlice(a)
	fmt.Println(crypto)
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
