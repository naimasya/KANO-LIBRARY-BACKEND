package util

func ErrorChecker(err error) {
	if err != nil {
		panic(err)
	}
}
