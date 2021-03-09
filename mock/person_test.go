package main

import (
	"testing"

	"mastering-go/mock/equipment"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPerson_Wake(t *testing.T) {
	type fields struct {
		name  string
		phone equipment.Phone
	}

	// 生成mock对象
	mockCtl := gomock.NewController(t)
	mockPhone := equipment.NewMockPhone(mockCtl)
	// 设置mock对象的接口方法返回值
	mockPhone.EXPECT().WeChat().Return(true)
	mockPhone.EXPECT().DingTalk().Return(true)
	mockPhone.EXPECT().ZhiHu().Return(true)

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "case-1",
			fields: fields{
				name:  "iphone6s",
				phone: equipment.NewIphone6s(),
			},
			want: true,
		},
		{
			name: "case-2",
			fields: fields{
				name:  "mocked phone",
				phone: mockPhone,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPerson(tt.fields.name, tt.fields.phone)
			assert.Equal(t, tt.want, p.Wake())
		})
	}
}
