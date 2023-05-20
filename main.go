package main

import (
	"fmt"

	"github.com/codeflames/simplebank/utils"
)

func main() {
	len := 5

	for i := 0; i < len; i++ {
		hello := utils.RandomString(utils.RandomInt(5, 10))

		fmt.Println(hello)
	}
}
