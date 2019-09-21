package wxpay

import "testing"

func TestStructToParameter(t *testing.T) {
	value := struct {
		A string `xml:"a"`
		B string `xml:"b"`
		C string `xml:"-"`
	}{
		A: "value a",
		B: "value b",
	}
	params, err := SignStructToParameter(value)
	if err != nil {
		t.Errorf("反射类型错误：%v", err)
		return
	}
	t.Log(params)
}
