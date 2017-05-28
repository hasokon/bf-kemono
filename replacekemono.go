package main

import (
	"strings"
)

func replaceKemono(text []byte) []byte {
	str := string(text)
	str = strings.Replace(str, "たのしー！", ">", -1)
	str = strings.Replace(str, "たーのしー！", "+", -1)
	str = strings.Replace(str, "すごーい！", "<", -1)
	str = strings.Replace(str, "すっごーい！", "-", -1)
	str = strings.Replace(str, "うわー！", "[", -1)
	str = strings.Replace(str, "わーい！", "]", -1)
	str = strings.Replace(str, "なにこれなにこれ！", ".", -1)
	str = strings.Replace(str, "おもしろーい！", ",", -1)

	return []byte(str)
}
