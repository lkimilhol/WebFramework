package handler

import "fmt"

type Login struct {
	uid int64
	nickname string
}

func (login *Login) Handler() {

	fmt.Println("login Api Call")
}