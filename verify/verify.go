package verify

import (
	"fmt"
	"strconv"
	"unicode"
)

var LenErr = fmt.Errorf("长度不对")
var NumErr = fmt.Errorf("不是数字")

var weight = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
var code = []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}

func VerificateID(id string) (bool, error) {
	lastOne, err := getCheckCode(id)
	if err != nil {
		return false, err
	}

	if string(id[17]) == lastOne {
		return true, nil
	}
	return false, nil
}

func getCheckCode(id string) (string, error) {
	if len(id) != 18 {
		return "", LenErr
	}

	sum := 0
	for i := 0; i < 17; i++ {
		if unicode.IsDigit(rune(id[i])) {
			n, err := strconv.Atoi(string(id[i]))
			if err != nil {
				return "", err
			}
			sum += n * weight[i]
		} else {
			return "", NumErr
		}
	}

	m := sum % 11
	return code[m], nil
}
