package main

import (
	"example/cipher/cipher"
	"fmt"
)

func main() {
	res1 := cipher.Caesar("abc", 1)
	res2 := cipher.Caesar("ABC", 3)
	res3 := cipher.Caesar("xyz", 5)

	fmt.Println("res1", res1)
	fmt.Println("res2", res2)
	fmt.Println("res3", res3)
}
