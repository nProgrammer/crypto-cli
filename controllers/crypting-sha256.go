package controllers

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func EncryptSHA256() {
	fmt.Print("Pass the string to encrypt: ")
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	textToHash := []byte(str)
	crypted := sha256.Sum256(textToHash)
	text := hex.EncodeToString(crypted[:])
	Save(text)
}
