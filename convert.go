package main

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

// 十进制转62进制：[0~9,a~z,A~Z]一共是 10 + 26 + 26 = 62个字符
var decimalToB62 map[int]string = map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f", 16: "g", 17: "h", 18: "i", 19: "j", 20: "k", 21: "l", 22: "m", 23: "n", 24: "o", 25: "p", 26: "q", 27: "r", 28: "s", 29: "t", 30: "u", 31: "v", 32: "w", 33: "x", 34: "y", 35: "z", 36: "A", 37: "B", 38: "C", 39: "D", 40: "E", 41: "F", 42: "G", 43: "H", 44: "I", 45: "J", 46: "K", 47: "L", 48: "M", 49: "N", 50: "O", 51: "P", 52: "Q", 53: "R", 54: "S", 55: "T", 56: "U", 57: "V", 58: "W", 59: "X", 60: "Y", 61: "Z"}

// 十进制数字转换为62进制字符串
func decToB62String(num int) string {
	var result = ""
	radix := 62
	var remainder int
	var remainders string
	for num != 0 {
		remainder = num % radix
		if 9 < remainder && remainder < 61 {
			// char
			remainders = decimalToB62[remainder]
		} else {
			remainders = strconv.Itoa(remainder)
		}
		result = remainders + result
		num = num / radix
	}
	return result
}

func md5String(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
