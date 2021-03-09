package main

import (
	"mastering-go/mock/equipment"
)

func main() {
	phone := equipment.NewIphone6s()
	xiaoMing := NewPerson("xiaoMing", phone)
	xiaoMing.Wake()
}
