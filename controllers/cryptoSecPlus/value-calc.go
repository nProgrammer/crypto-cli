package cryptoSecPlus

import "strings"

func ValueOfKey(str string, a []string) (int, int, int) {
	key := strings.TrimSpace(str)
	key = strings.ToUpper(key)
	keyA := strings.Split(key, "")

	keyF := 0
	keyS := 0
	keyValue := 0
	i := 0
	for i < len(keyA) {
		j := 0
		for j < len(a) {
			if a[j] == keyA[i] {
				if key[i] == key[0] {
					keyF = j
				} else if key[i] == key[len(key)-1] {
					keyS = j
				}
				keyValue += j
				break
			}
			j++
		}
		i++
	}

	return keyValue, keyF, keyS
}
