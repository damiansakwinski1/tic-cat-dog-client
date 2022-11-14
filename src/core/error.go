package core

import "os"

func HandleError(err error) {
	println(err.Error())
	os.Exit(1)
}
