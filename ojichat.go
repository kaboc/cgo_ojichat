package main

import (
	"C"
	"github.com/greymd/ojichat/generator"
)

//export getMessage
func getMessage(name *C.char) *C.char {
	cnf := generator.Config{
		TargetName:       C.GoString(name),
		EmojiNum:         3,
		PunctuationLevel: 2,
	}

	msg, err := generator.Start(cnf)
	if err != nil {
		return C.CString("")
	}
	return C.CString(msg)
}

func main() {}
