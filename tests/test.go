package main

import (
	"fmt"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// mudar o algoritmo para ao inves de gerar aleatoriamente, gerar em base62 + ofuscação hashid
func GeneratorShortId(length int, url string) (string, error) {

	if length == 0 || url == ""{
		return "", nil
	} 
	num := new(big.Int).SetBytes([]byte(url))
	var result []byte
	zero := big.NewInt(0)
	base := big.NewInt(62)

	for num.Cmp(zero) > 0 {
		var remainder big.Int
		num.DivMod(num, base, &remainder)
		result = append(result, charset[remainder.Int64()])
	}
	return string(result[:length]), nil
}

func main() {
	fmt.Println(GeneratorShortId(6, "www.googledasdsadasdasdasdsadsadas"))
	fmt.Println(GeneratorShortId(6, "www.googleeeeeeeeeeeeeeee"))
	fmt.Println(GeneratorShortId(6, "www.google"))

}
