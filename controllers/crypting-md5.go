package controllers

import (
	"bufio"
	"crypto-cli/errors"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func EncryptMD5() {
	fmt.Print("Pass the string to encrypt: ")
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	errors.CheckError(err)
	str = strings.TrimSpace(str)
	textToHash := []byte(str)
	crypted := md5.Sum(textToHash)
	text := hex.EncodeToString(crypted[:])
	Save(text)
}
