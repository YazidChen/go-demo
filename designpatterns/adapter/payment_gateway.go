package adapter

const (
	UnknownChannel PayChannel = iota // 未知渠道
	Alipay                           // 支付宝 APP 支付
	AlipayWap                        // 支付宝手机网站支付
	AlipayQR                         // 支付宝扫码支付
	AlipayScan                       // 支付宝条码支付
	AlipayLite                       // 支付宝小程序支付
	WX                               // 微信 APP 支付
	WXPub                            // 微信 JSAPI 支付
	WXPubQR                          // 微信 Native 支付
	WXPubScan                        // 微信付款码支付
	WXWap                            // 微信 H5 支付
	WXLite                           // 微信付款码小程序支付
)

const (
	CNY CurrencyType = "cny" // 人民币
	HKD CurrencyType = "hkd" // 港币
	USD CurrencyType = "usd" // 美元
	EUR CurrencyType = "eur" // 欧元
	JPY CurrencyType = "jpy" // 日元
	KRW CurrencyType = "krw" // 韩元
	AUD CurrencyType = "aud" // 澳元
)

// PaymentGateway 支付网关
type PaymentGateway interface {
	createPayOrder(o PayOrder)
}

// PayChannel 支付渠道
type PayChannel int

// CurrencyType 3 位 ISO 货币代码，小写字母
type CurrencyType string

// ChannelExtra 支付渠道扩展字段
type ChannelExtra struct {
}

// WXChannelExtra 支付渠道扩展字段 - 微信 APP 支付
type WXChannelExtra struct {
	ChannelExtra
	LimitPay string // 指定支付方式，指定不能使用信用卡支付可设置为 no_credit
}

// WXPubChannelExtra 支付渠道扩展字段 - 微信 JSAPI 支付
type WXPubChannelExtra struct {
	ChannelExtra
	LimitPay string // 指定支付方式，指定不能使用信用卡支付可设置为 no_credit
	OpenId   string // 用户微信应用下openid
}

// PayOrder 支付订单对象
type PayOrder struct {
	PayOrderNo   string       // 支付订单号
	Amount       int          // 支付金额，单位分
	Currency     CurrencyType // 货币代码
	ClientIp     string       // 发起接口请求的客户端IP
	Subject      string       // 订单标题
	Body         string       // 订单描述
	Channel      PayChannel   // 支付渠道
	ChannelExtra ChannelExtra // 支付渠道扩展字段
}

// 设置支付渠道扩展
func (p *PayOrder) setChannelExtra(ce ChannelExtra) {
	p.ChannelExtra = ce
}

func main() {
	o := PayOrder{}
	t := &ThirdAdapter{}
	t.createPayOrder(o)
}
