package adapter

import "fmt"

// ThirdPay 第三方支付机构
type ThirdPay struct{}

// ThirdPayOrder 第三方支付订单对象
type ThirdPayOrder struct {
}

// 支付落单
func (p *ThirdPay) thirdPayToPay(o ThirdPayOrder) {
	fmt.Println("第三方机构支付落单")
}

// ThirdAdapter 第三方支付适配器
type ThirdAdapter struct {
	t *ThirdPay
}

// 适配方法转换
func (ta *ThirdAdapter) createPayOrder(o PayOrder) {
	fmt.Println("支付落单参数对象转换")
	to := ThirdPayOrder{}
	ta.t.thirdPayToPay(to)
}
