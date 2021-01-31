package controllers

import (
	"bufio"
	"crypto-cli/errors"
	"crypto-cli/views"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func App(a []string, crypto []string) {
	reader := bufio.NewReader(os.Stdin)
	option := 1
	for option != 4 {
		views.WriteMainOptions()
		fmt.Print("Enter option's number: ")
		str, err := reader.ReadString('\n')
		errors.CheckError(err)
		f, err := strconv.Atoi(strings.TrimSpace(str))
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("You choosed option %v \n", f)
			if f > 4 || f < 0 {
				App(a, crypto)
			}
			switch f {
			case 1:
				EncryptSHA256()
				break
			case 2:
				EncryptMD5()
				break
			case 3:
				CryptoSecApp(option, a, crypto)
				break
			case 4:
				option = 4
				break
			}
		}
	}
}
