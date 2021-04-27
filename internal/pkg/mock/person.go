package mock

import (
	"fmt"

	"mastering-go/internal/pkg/mock/equipment"
)

type Person struct {
	name  string
	phone equipment.Phone
}

func NewPerson(name string, phone equipment.Phone) *Person {
	return &Person{
		name:  name,
		phone: phone,
	}
}

func (p *Person) Sleep() {
	fmt.Printf("%s go to sleep!\n", p.name)
}

func (p *Person) Wake() bool {
	fmt.Printf("%s's wake up: \n", p.name)
	if p.phone.WeChat() && p.phone.DingTalk() && p.phone.ZhiHu() {
		p.Sleep()
		return true
	}
	return false
}
