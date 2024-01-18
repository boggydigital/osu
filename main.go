package main

import (
	"fmt"
)

func main() {

	glsvt, err := GitLatestTag()
	if err != nil {
		panic(err)
	}

	fmt.Println(glsvt)
}
