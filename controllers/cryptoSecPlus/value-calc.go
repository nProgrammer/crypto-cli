package cryptoSecPlus

import "strings"

func ValueOfKey(str string, a []string) int {
	key := strings.TrimSpace(str)
	key = strings.ToUpper(key)
	keyA := strings.Split(key, "")

	keyValue := 0
	i := 0
	for i < len(keyA) {
		j := 0
		for j < len(a) {
			if a[j] == keyA[i] {
				keyValue += j
				break
			}
			j++
		}
		i++
	}

	return keyValue
}
