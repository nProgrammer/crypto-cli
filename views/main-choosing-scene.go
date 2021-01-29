package views

import "fmt"

func WriteMainOptions() {
	fmt.Print("What do you want to do? \n",
		"	1. Encrypt typed string to SHA256, \n",
		"	2. Encrypt typed string to MD5, \n",
		"	3. Use crypto-cli-secure+ , \n",
		"	4. Exit. \n")
}
