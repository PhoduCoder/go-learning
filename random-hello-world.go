package main 

import (
	"fmt"
	"math/rand"
)

var salutations_list = []string{
	"Hello, world",
	"Καλημέρα κόσμε",
	"こんにちは世界",
	"سلام دنیا‎",
	"Привет, мир",
	"你好世界",
	"မင်္ဂလာ ပါ ကမ္ဘာ ",
	"Hallo Welt",
	"Olá, mundo",
	"Привет, мир",
	"السلام عليكم يا عالم",
}

func main(){
	for {
		index := rand.Intn(len(salutations_list))
		choice := hello(index)
		fmt.Printf("We are saying %s!\n", choice)
	}
}

func hello(num int) string{
	return salutations_list[num]
}