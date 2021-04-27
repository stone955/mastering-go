package main

import (
	"mastering-go/internal/pkg/mock"
	"mastering-go/internal/pkg/mock/equipment"
)

func main() {
	phone := equipment.NewIphone6s()
	xiaoMing := mock.NewPerson("xiaoMing", phone)
	xiaoMing.Wake()
}
