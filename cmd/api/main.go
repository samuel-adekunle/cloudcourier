package main

import (
	"fmt"
	"os"
)

func main() {
	v, err := os.Stat("../../")
	fmt.Println(v, err)
}
