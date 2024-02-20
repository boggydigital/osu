package main

import "fmt"

func main() {

	glsvt, err := GitLatestTag()
	if err != nil {
		panic(err)
	}

	fmt.Println("current:", glsvt)

	glsvt.Increment()

	fmt.Println("tag:", glsvt)

	if err = GitTag(glsvt); err != nil {
		panic(err)
	}

	fmt.Println("push:", glsvt)

	if err = GitPushOrigin(glsvt); err != nil {
		panic(err)
	}

	fmt.Println("done")

}
