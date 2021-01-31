package controllers

import (
	"bufio"
	"crypto-cli/controllers/cryptoSecPlus"
	"crypto-cli/errors"
	"crypto-cli/views"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func CryptoSecApp(option int, a []string, crypto []string) {
	reader := bufio.NewReader(os.Stdin)
	for option != 5 {
		views.WriteSecOptions()
		fmt.Print("Enter option's number: ")
		str, err := reader.ReadString('\n')
		errors.CheckError(err)

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
				cryptoSecPlus.EncryptFromText(a, crypto)
				break
			case 2:
				cryptoSecPlus.EncryptFromFile()
				break
			case 3:
				cryptoSecPlus.DecryptFromText()
				break
			case 4:
				cryptoSecPlus.DecryptFromFile()
				break
			case 5:
				option = 5
			}
		}
	}
}
