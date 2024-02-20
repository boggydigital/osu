package main

func main() {

	glsvt, err := GitLatestTag()
	if err != nil {
		panic(err)
	}

	glsvt.Increment()

	if err = GitTag(glsvt); err != nil {
		panic(err)
	}

	if err = GitPushOrigin(glsvt); err != nil {
		panic(err)
	}

}
