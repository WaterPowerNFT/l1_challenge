package main

import (
	"fmt"
	"time"
)

type Human struct {
	age      uint16
	Height   uint16
	birthday time.Time
}

type Action struct {
	Human // Human встроен в Action
	date  time.Time
}

func (this *Human) WhenBirthday() time.Time {
	return this.birthday
}

func main() {

	var buy_chocolate Action
	buy_chocolate.age = 1
	//можно обращаться к переменной через human и напрямую к наименованию
	fmt.Println(buy_chocolate.age, buy_chocolate.Human.age)
	buy_chocolate.birthday = time.Now()
	//функция WhenBirthday() написана только для структуры Human, но из-за того, что Human встроена в Action, то ее можно вызвать и в Action
	fmt.Println(buy_chocolate.WhenBirthday())
	fmt.Println(buy_chocolate.Human.WhenBirthday())
}
