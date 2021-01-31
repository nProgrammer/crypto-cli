package controllers

import (
	"bufio"
	"crypto-cli/errors"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func EncryptSHA256() {
	fmt.Print("Pass the string to encrypt: ")
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	errors.CheckError(err)
	str = strings.TrimSpace(str)
	textToHash := []byte(str)
	crypted := sha256.Sum256(textToHash)
	text := hex.EncodeToString(crypted[:])
	Save(text)
}
