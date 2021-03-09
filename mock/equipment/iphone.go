package equipment

import (
	"fmt"
)

type Iphone6s struct {
}

func NewIphone6s() *Iphone6s {
	return &Iphone6s{}
}

func (p *Iphone6s) WeChat() bool {
	fmt.Println("Iphone6s chat wei xin!")
	return true
}

func (p *Iphone6s) DingTalk() bool {
	fmt.Println("Iphone6s play ding ding!")
	return true
}

func (p *Iphone6s) ZhiHu() bool {
	fmt.Println("Iphone6s read zhi hu!")
	return true
}
