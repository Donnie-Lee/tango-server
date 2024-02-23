package util

func AssertExec(condition bool, message string) {
	if !condition {
		panic(message)
	}
}

func AssertError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func AssertErrorMessage(err error, message string) {
	if err != nil {
		panic(message + err.Error())
	}
}
