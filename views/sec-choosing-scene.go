package views

import "fmt"

func WriteSecOptions() {
	fmt.Print("What do you want to do? \n",
		"	1. Encrypt typed string, \n",
		"	2. Encrypt text from file, \n",
		"	3. Decrypt typed string, \n",
		"	4. Decrypt text from file, \n",
		"	5. Exit. \n")
}
