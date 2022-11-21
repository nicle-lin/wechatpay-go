// 微信支付服务商电商收付通
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/open/pay/chapter3_3_0.shtml
// 这里只做了合单小程序下单、退款

package ecommerce

import (
	"encoding/json"
	"fmt"
	"time"
)

// Transaction 电商收付通支付通知结构体
type Transaction struct {
	CombineAppid      *string           `json:"combine_appid"`
	CombineMchid      *string           `json:"combine_mchid"`
	CombineOutTradeNo *string           `json:"combine_out_trade_no"`
	SceneInfo         *SceneInfo        `json:"scene_info,omitempty"`
	SubOrders         []SubOrders       `json:"sub_orders"`
	CombinePayerInfo  *TransactionPayer `json:"combine_payer_info,omitempty"`
}

func (o Transaction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.CombineAppid != nil {
		toSerialize["combine_appid"] = o.CombineAppid
	}

	if o.CombineMchid != nil {
		toSerialize["combine_mchid"] = o.CombineMchid
	}

	if o.CombineOutTradeNo != nil {
		toSerialize["combine_out_trade_no"] = o.CombineOutTradeNo
	}

	if o.SceneInfo != nil {
		toSerialize["scene_info"] = o.SceneInfo
	}

	if o.SubOrders != nil {
		toSerialize["sub_orders"] = o.SubOrders
	}

	if o.CombinePayerInfo != nil {
		toSerialize["combine_payer_info"] = o.CombinePayerInfo
	}

	return json.Marshal(toSerialize)
}

func (o Transaction) String() string {
	var ret string
	if o.CombineAppid == nil {
		ret += "CombineAppid:<nil>, "
	} else {
		ret += fmt.Sprintf("CombineAppid:%v, ", *o.CombineAppid)
	}

	if o.CombineMchid == nil {
		ret += "CombineMchid:<nil>, "
	} else {
		ret += fmt.Sprintf("CombineMchid:%v, ", *o.CombineMchid)
	}

	if o.CombineOutTradeNo == nil {
		ret += "CombineOutTradeNo:<nil>, "
	} else {
		ret += fmt.Sprintf("CombineOutTradeNo:%v, ", *o.CombineOutTradeNo)
	}
	ret += fmt.Sprintf("SceneInfo:%v, ", o.SceneInfo)
	ret += fmt.Sprintf("SubOrders:%v, ", o.SubOrders)
	ret += fmt.Sprintf("CombinePayerInfo:%v", o.CombinePayerInfo)

	return fmt.Sprintf("Transaction{%s}", ret)
}

func (o Transaction) Clone() *Transaction {
	ret := Transaction{}

	if o.CombineAppid != nil {
		ret.CombineAppid = new(string)
		*ret.CombineAppid = *o.CombineAppid
	}

	if o.CombineAppid != nil {
		ret.CombineAppid = new(string)
		*ret.CombineAppid = *o.CombineAppid
	}

	if o.CombineOutTradeNo != nil {
		ret.CombineOutTradeNo = new(string)
		*ret.CombineOutTradeNo = *o.CombineOutTradeNo
	}

	if o.SceneInfo != nil {
		ret.SceneInfo = o.SceneInfo.Clone()
	}

	if o.SubOrders != nil {
		ret.SubOrders = make([]SubOrders, len(o.SubOrders))
		for i, item := range o.SubOrders {
			ret.SubOrders[i] = *item.Clone()
		}
	}

	if o.CombinePayerInfo != nil {
		ret.CombinePayerInfo = o.CombinePayerInfo.Clone()
	}

	return &ret
}

type SceneInfo struct {
	DeviceId *string `json:"device_id,omitempty"`
}

func (o SceneInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceId != nil {
		toSerialize["device_id"] = o.DeviceId
	}
	return json.Marshal(toSerialize)
}

func (o SceneInfo) String() string {
	var ret string
	if o.DeviceId == nil {
		ret += "DeviceId:<nil>"
	} else {
		ret += fmt.Sprintf("DeviceId:%v", *o.DeviceId)
	}
	return ret
}

func (o SceneInfo) Clone() *SceneInfo {
	ret := SceneInfo{}
	if o.DeviceId != nil {
		ret.DeviceId = new(string)
		*ret.DeviceId = *o.DeviceId
	}
	return &ret
}

type SubOrders struct {
	Mchid           *string            `json:"mchid"`
	TradeType       *string            `json:"trade_type"`
	TradeState      *string            `json:"trade_state"`
	BankType        *string            `json:"bank_type,omitempty"`
	Attach          *string            `json:"attach"`
	SuccessTime     *string            `json:"success_time"`
	TransactionId   *string            `json:"transaction_id"`
	OutTradeNo      *string            `json:"out_trade_no"`
	SubMchid        *string            `json:"sub_mchid"`
	SubAppid        *string            `json:"sub_appid,omitempty"`
	Amount          *TransactionAmount `json:"amount"`
	PromotionDetail []PromotionDetail  `json:"promotion_detail,omitempty"`
}

func (o SubOrders) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Mchid != nil {
		toSerialize["mchid"] = o.Mchid
	}
	if o.TradeType != nil {
		toSerialize["trade_type"] = o.TradeType
	}
	if o.TradeState != nil {
		toSerialize["trade_state"] = o.TradeState
	}
	if o.BankType != nil {
		toSerialize["bank_type"] = o.BankType
	}
	if o.Attach != nil {
		toSerialize["attach"] = o.Attach
	}
	if o.SuccessTime != nil {
		toSerialize["success_time"] = o.SuccessTime
	}
	if o.TransactionId != nil {
		toSerialize["transaction_id"] = o.TransactionId
	}
	if o.OutTradeNo != nil {
		toSerialize["out_trade_no"] = o.OutTradeNo
	}
	if o.SubMchid != nil {
		toSerialize["sub_mchid"] = o.SubMchid
	}
	if o.SubAppid != nil {
		toSerialize["sub_appid"] = o.SubAppid
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.PromotionDetail != nil {
		toSerialize["promotion_detail"] = o.PromotionDetail
	}

	return json.Marshal(toSerialize)
}

func (o SubOrders) String() string {
	var ret string
	if o.Mchid == nil {
		ret += "Mchid:<nil>"
	} else {
		ret += fmt.Sprintf("Mchid:%v, ", *o.Mchid)
	}

	if o.TradeType == nil {
		ret += "TradeType:<nil>"
	} else {
		ret += fmt.Sprintf("TradeType:%v, ", *o.TradeType)
	}

	if o.TradeState == nil {
		ret += "TradeState:<nil>"
	} else {
		ret += fmt.Sprintf("TradeState:%v, ", *o.TradeState)
	}
	if o.BankType == nil {
		ret += "BankType:<nil>"
	} else {
		ret += fmt.Sprintf("BankType:%v, ", *o.BankType)
	}
	if o.Attach == nil {
		ret += "Attach:<nil>"
	} else {
		ret += fmt.Sprintf("Attach:%v, ", *o.Attach)
	}
	if o.SuccessTime == nil {
		ret += "SuccessTime:<nil>"
	} else {
		ret += fmt.Sprintf("SuccessTime:%v, ", *o.SuccessTime)
	}

	if o.TransactionId == nil {
		ret += "TransactionId:<nil>"
	} else {
		ret += fmt.Sprintf("TransactionId:%v, ", *o.TransactionId)
	}
	if o.OutTradeNo == nil {
		ret += "OutTradeNo:<nil>"
	} else {
		ret += fmt.Sprintf("OutTradeNo:%v, ", *o.OutTradeNo)
	}
	if o.SubMchid == nil {
		ret += "SubMchid:<nil>"
	} else {
		ret += fmt.Sprintf("SubMchid:%v, ", *o.SubMchid)
	}
	if o.SubAppid == nil {
		ret += "SubAppid:<nil>"
	} else {
		ret += fmt.Sprintf("SubAppid:%v, ", *o.SubAppid)
	}
	ret += fmt.Sprintf("Amount:%v, ", o.Amount)
	ret += fmt.Sprintf("PromotionDetail:%v", o.PromotionDetail)

	return fmt.Sprintf("SubOrders{%s}", ret)
}

func (o SubOrders) Clone() *SubOrders {
	ret := SubOrders{}
	if o.Mchid != nil {
		ret.Mchid = new(string)
		*ret.Mchid = *o.Mchid
	}

	if o.TradeType != nil {
		ret.TradeType = new(string)
		*ret.TradeType = *o.TradeType
	}

	if o.TradeState != nil {
		ret.TradeState = new(string)
		*ret.TradeState = *o.TradeState
	}

	if o.BankType != nil {
		ret.BankType = new(string)
		*ret.BankType = *o.BankType
	}

	if o.Attach != nil {
		ret.Attach = new(string)
		*ret.Attach = *o.Attach
	}

	if o.SuccessTime != nil {
		ret.SuccessTime = new(string)
		*ret.SuccessTime = *o.SuccessTime
	}

	if o.TransactionId != nil {
		ret.TransactionId = new(string)
		*ret.TransactionId = *o.TransactionId
	}

	if o.OutTradeNo != nil {
		ret.OutTradeNo = new(string)
		*ret.OutTradeNo = *o.OutTradeNo
	}

	if o.SubMchid != nil {
		ret.SubMchid = new(string)
		*ret.SubMchid = *o.SubMchid
	}

	if o.SubAppid != nil {
		ret.SubAppid = new(string)
		*ret.SubAppid = *o.SubAppid
	}

	if o.Amount != nil {
		ret.Amount = o.Amount.Clone()
	}

	if o.PromotionDetail != nil {
		ret.PromotionDetail = make([]PromotionDetail, len(o.PromotionDetail))
		for i, item := range o.PromotionDetail {
			ret.PromotionDetail[i] = *item.Clone()
		}
	}

	return &ret
}

type TransactionAmount struct {
	TotalAmount   *int64  `json:"total_amount"`
	Currency      *string `json:"currency"`
	PayerAmount   *int64  `json:"payer_amount"`
	PayerCurrency *string `json:"payer_currency"`
}

func (o TransactionAmount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.TotalAmount != nil {
		toSerialize["total_amount"] = o.TotalAmount
	}

	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}

	if o.PayerAmount != nil {
		toSerialize["payer_amount"] = o.PayerAmount
	}

	if o.PayerCurrency != nil {
		toSerialize["payer_currency"] = o.PayerCurrency
	}

	return json.Marshal(toSerialize)
}

func (o TransactionAmount) String() string {
	var ret string

	if o.TotalAmount == nil {
		ret += "TotalAmount:<nil>"
	} else {
		ret += fmt.Sprintf("TotalAmount:%v, ", *o.TotalAmount)
	}

	if o.Currency == nil {
		ret += "Currency:<nil>, "
	} else {
		ret += fmt.Sprintf("Currency:%v, ", *o.Currency)
	}

	if o.PayerAmount == nil {
		ret += "PayerAmount:<nil>, "
	} else {
		ret += fmt.Sprintf("PayerAmount:%v, ", *o.PayerAmount)
	}

	if o.PayerCurrency == nil {
		ret += "PayerCurrency:<nil>, "
	} else {
		ret += fmt.Sprintf("PayerCurrency:%v", *o.PayerCurrency)
	}

	return fmt.Sprintf("TransactionAmount{%s}", ret)
}

func (o TransactionAmount) Clone() *TransactionAmount {
	ret := TransactionAmount{}

	if o.TotalAmount != nil {
		ret.TotalAmount = new(int64)
		*ret.TotalAmount = *o.TotalAmount
	}

	if o.Currency != nil {
		ret.Currency = new(string)
		*ret.Currency = *o.Currency
	}

	if o.PayerAmount != nil {
		ret.PayerAmount = new(int64)
		*ret.PayerAmount = *o.PayerAmount
	}

	if o.PayerCurrency != nil {
		ret.PayerCurrency = new(string)
		*ret.PayerCurrency = *o.PayerCurrency
	}

	return &ret
}

type PromotionDetail struct {
	// 券ID
	CouponId *string `json:"coupon_id"`
	// 优惠名称
	Name *string `json:"name,omitempty"`
	// GLOBAL：全场代金券；SINGLE：单品优惠
	Scope *string `json:"scope,omitempty"`
	// CASH：充值；NOCASH：预充值。
	Type *string `json:"type,omitempty"`
	// 优惠券面额
	Amount *int64 `json:"amount"`
	// 活动ID，批次ID
	StockId *string `json:"stock_id,omitempty"`
	// 单位为分
	WechatpayContribute *int64 `json:"wechatpay_contribute,omitempty"`
	// 单位为分
	MerchantContribute *int64 `json:"merchant_contribute,omitempty"`
	// 单位为分
	OtherContribute *int64 `json:"other_contribute,omitempty"`
	// CNY：人民币，境内商户号仅支持人民币。
	Currency    *string                `json:"currency,omitempty"`
	GoodsDetail []PromotionGoodsDetail `json:"goods_detail,omitempty"`
}

func (o PromotionDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.CouponId != nil {
		toSerialize["coupon_id"] = o.CouponId
	}

	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}

	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}

	if o.StockId != nil {
		toSerialize["stock_id"] = o.StockId
	}

	if o.WechatpayContribute != nil {
		toSerialize["wechatpay_contribute"] = o.WechatpayContribute
	}

	if o.MerchantContribute != nil {
		toSerialize["merchant_contribute"] = o.MerchantContribute
	}

	if o.OtherContribute != nil {
		toSerialize["other_contribute"] = o.OtherContribute
	}

	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}

	if o.GoodsDetail != nil {
		toSerialize["goods_detail"] = o.GoodsDetail
	}
	return json.Marshal(toSerialize)
}

func (o PromotionDetail) String() string {
	var ret string
	if o.CouponId == nil {
		ret += "CouponId:<nil>, "
	} else {
		ret += fmt.Sprintf("CouponId:%v, ", *o.CouponId)
	}

	if o.Name == nil {
		ret += "Name:<nil>, "
	} else {
		ret += fmt.Sprintf("Name:%v, ", *o.Name)
	}

	if o.Scope == nil {
		ret += "Scope:<nil>, "
	} else {
		ret += fmt.Sprintf("Scope:%v, ", *o.Scope)
	}

	if o.Type == nil {
		ret += "Type:<nil>, "
	} else {
		ret += fmt.Sprintf("Type:%v, ", *o.Type)
	}

	if o.Amount == nil {
		ret += "Amount:<nil>, "
	} else {
		ret += fmt.Sprintf("Amount:%v, ", *o.Amount)
	}

	if o.StockId == nil {
		ret += "StockId:<nil>, "
	} else {
		ret += fmt.Sprintf("StockId:%v, ", *o.StockId)
	}

	if o.WechatpayContribute == nil {
		ret += "WechatpayContribute:<nil>, "
	} else {
		ret += fmt.Sprintf("WechatpayContribute:%v, ", *o.WechatpayContribute)
	}

	if o.MerchantContribute == nil {
		ret += "MerchantContribute:<nil>, "
	} else {
		ret += fmt.Sprintf("MerchantContribute:%v, ", *o.MerchantContribute)
	}

	if o.OtherContribute == nil {
		ret += "OtherContribute:<nil>, "
	} else {
		ret += fmt.Sprintf("OtherContribute:%v, ", *o.OtherContribute)
	}

	if o.Currency == nil {
		ret += "Currency:<nil>, "
	} else {
		ret += fmt.Sprintf("Currency:%v, ", *o.Currency)
	}

	ret += fmt.Sprintf("GoodsDetail:%v", o.GoodsDetail)

	return fmt.Sprintf("PromotionDetail{%s}", ret)
}

func (o PromotionDetail) Clone() *PromotionDetail {
	ret := PromotionDetail{}

	if o.CouponId != nil {
		ret.CouponId = new(string)
		*ret.CouponId = *o.CouponId
	}

	if o.Name != nil {
		ret.Name = new(string)
		*ret.Name = *o.Name
	}

	if o.Scope != nil {
		ret.Scope = new(string)
		*ret.Scope = *o.Scope
	}

	if o.Type != nil {
		ret.Type = new(string)
		*ret.Type = *o.Type
	}

	if o.Amount != nil {
		ret.Amount = new(int64)
		*ret.Amount = *o.Amount
	}

	if o.StockId != nil {
		ret.StockId = new(string)
		*ret.StockId = *o.StockId
	}

	if o.WechatpayContribute != nil {
		ret.WechatpayContribute = new(int64)
		*ret.WechatpayContribute = *o.WechatpayContribute
	}

	if o.MerchantContribute != nil {
		ret.MerchantContribute = new(int64)
		*ret.MerchantContribute = *o.MerchantContribute
	}

	if o.OtherContribute != nil {
		ret.OtherContribute = new(int64)
		*ret.OtherContribute = *o.OtherContribute
	}

	if o.Currency != nil {
		ret.Currency = new(string)
		*ret.Currency = *o.Currency
	}

	if o.GoodsDetail != nil {
		ret.GoodsDetail = make([]PromotionGoodsDetail, len(o.GoodsDetail))
		for i, item := range o.GoodsDetail {
			ret.GoodsDetail[i] = *item.Clone()
		}
	}

	return &ret
}

type PromotionGoodsDetail struct {
	// 商品编码
	GoodsId *string `json:"goods_id"`
	// 商品数量
	Quantity *int64 `json:"quantity"`
	// 商品价格
	UnitPrice *int64 `json:"unit_price"`
	// 商品优惠金额
	DiscountAmount *int64 `json:"discount_amount"`
	// 商品备注
	GoodsRemark *string `json:"goods_remark,omitempty"`
}

func (o PromotionGoodsDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.GoodsId == nil {
		return nil, fmt.Errorf("field `GoodsId` is required and must be specified in PromotionGoodsDetail")
	}
	toSerialize["goods_id"] = o.GoodsId

	if o.Quantity == nil {
		return nil, fmt.Errorf("field `Quantity` is required and must be specified in PromotionGoodsDetail")
	}
	toSerialize["quantity"] = o.Quantity

	if o.UnitPrice == nil {
		return nil, fmt.Errorf("field `UnitPrice` is required and must be specified in PromotionGoodsDetail")
	}
	toSerialize["unit_price"] = o.UnitPrice

	if o.DiscountAmount == nil {
		return nil, fmt.Errorf("field `DiscountAmount` is required and must be specified in PromotionGoodsDetail")
	}
	toSerialize["discount_amount"] = o.DiscountAmount

	if o.GoodsRemark != nil {
		toSerialize["goods_remark"] = o.GoodsRemark
	}
	return json.Marshal(toSerialize)
}

func (o PromotionGoodsDetail) String() string {
	var ret string
	if o.GoodsId == nil {
		ret += "GoodsId:<nil>, "
	} else {
		ret += fmt.Sprintf("GoodsId:%v, ", *o.GoodsId)
	}

	if o.Quantity == nil {
		ret += "Quantity:<nil>, "
	} else {
		ret += fmt.Sprintf("Quantity:%v, ", *o.Quantity)
	}

	if o.UnitPrice == nil {
		ret += "UnitPrice:<nil>, "
	} else {
		ret += fmt.Sprintf("UnitPrice:%v, ", *o.UnitPrice)
	}

	if o.DiscountAmount == nil {
		ret += "DiscountAmount:<nil>, "
	} else {
		ret += fmt.Sprintf("DiscountAmount:%v, ", *o.DiscountAmount)
	}

	if o.GoodsRemark == nil {
		ret += "GoodsRemark:<nil>"
	} else {
		ret += fmt.Sprintf("GoodsRemark:%v", *o.GoodsRemark)
	}

	return fmt.Sprintf("PromotionGoodsDetail{%s}", ret)
}

func (o PromotionGoodsDetail) Clone() *PromotionGoodsDetail {
	ret := PromotionGoodsDetail{}

	if o.GoodsId != nil {
		ret.GoodsId = new(string)
		*ret.GoodsId = *o.GoodsId
	}

	if o.Quantity != nil {
		ret.Quantity = new(int64)
		*ret.Quantity = *o.Quantity
	}

	if o.UnitPrice != nil {
		ret.UnitPrice = new(int64)
		*ret.UnitPrice = *o.UnitPrice
	}

	if o.DiscountAmount != nil {
		ret.DiscountAmount = new(int64)
		*ret.DiscountAmount = *o.DiscountAmount
	}

	if o.GoodsRemark != nil {
		ret.GoodsRemark = new(string)
		*ret.GoodsRemark = *o.GoodsRemark
	}

	return &ret
}

type TransactionPayer struct {
	Openid *string `json:"openid"`
}

func (o TransactionPayer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Openid != nil {
		toSerialize["openid"] = o.Openid
	}
	return json.Marshal(toSerialize)
}

func (o TransactionPayer) String() string {
	var ret string
	if o.Openid == nil {
		ret += "Openid:<nil>, "
	} else {
		ret += fmt.Sprintf("Openid:%v, ", *o.Openid)
	}

	return fmt.Sprintf("TransactionPayer{%s}", ret)
}

func (o TransactionPayer) Clone() *TransactionPayer {
	ret := TransactionPayer{}

	if o.Openid != nil {
		ret.Openid = new(string)
		*ret.Openid = *o.Openid
	}

	return &ret
}

// CloseRequest 合单支付关闭订单参数[body]
type CloseRequest struct {
	// 合单商户appid, 合单发起方的appid。
	CombineAppid *string `json:"combine_appid"`
	// 子单信息
	SubOrders []CloseSubOrders `json:"sub_orders"`
}

func (o CloseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CombineAppid != nil {
		toSerialize["combine_appid"] = o.CombineAppid
	}

	if o.SubOrders != nil {
		toSerialize["sub_orders"] = o.SubOrders
	}

	return json.Marshal(toSerialize)
}

func (o CloseRequest) String() string {
	var ret string

	if o.CombineAppid == nil {
		ret += "CombineAppid:<nil>, "
	} else {
		ret += fmt.Sprintf("CombineAppid:%v, ", *o.CombineAppid)
	}

	ret += fmt.Sprintf("SubOrders:%v, ", o.SubOrders)
	return fmt.Sprintf("CloseOrderRequest{%s}", ret)
}

func (o CloseRequest) Clone() *CloseRequest {
	ret := CloseRequest{}

	if o.CombineAppid != nil {
		ret.CombineAppid = new(string)
		*ret.CombineAppid = *o.CombineAppid
	}

	if o.SubOrders != nil {
		ret.SubOrders = make([]CloseSubOrders, len(o.SubOrders))
		for i, item := range o.SubOrders {
			ret.SubOrders[i] = *item.Clone()
		}
	}

	return &ret
}

// CloseOrderRequest 合单支付关闭订单参数
type CloseOrderRequest struct {
	// 合单商户appid, 合单发起方的appid。
	CombineAppid *string `json:"combine_appid"`
	// 合单商户订单号, 合单支付总订单号
	CombineOutTradeNo *string `json:"combine_out_trade_no"`
	// 子单信息
	SubOrders []CloseSubOrders `json:"sub_orders"`
}

func (o CloseOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CombineAppid != nil {
		toSerialize["combine_appid"] = o.CombineAppid
	}
	if o.CombineOutTradeNo != nil {
		toSerialize["combine_out_trade_no"] = o.CombineOutTradeNo
	}
	if o.SubOrders != nil {
		toSerialize["sub_orders"] = o.SubOrders
	}

	return json.Marshal(toSerialize)
}

func (o CloseOrderRequest) String() string {
	var ret string

	if o.CombineAppid == nil {
		ret += "CombineAppid:<nil>, "
	} else {
		ret += fmt.Sprintf("CombineAppid:%v, ", *o.CombineAppid)
	}

	if o.CombineOutTradeNo == nil {
		ret += "CombineOutTradeNo:<nil>, "
	} else {
		ret += fmt.Sprintf("CombineOutTradeNo:%v, ", *o.CombineOutTradeNo)
	}

	ret += fmt.Sprintf("SubOrders:%v", o.SubOrders)
	return fmt.Sprintf("CloseOrderRequest{%s}", ret)
}

func (o CloseOrderRequest) Clone() *CloseOrderRequest {
	ret := CloseOrderRequest{}

	if o.CombineAppid != nil {
		ret.CombineAppid = new(string)
		*ret.CombineAppid = *o.CombineAppid
	}

	if o.CombineOutTradeNo != nil {
		ret.CombineOutTradeNo = new(string)
		*ret.CombineOutTradeNo = *o.CombineOutTradeNo
	}

	if o.SubOrders != nil {
		ret.SubOrders = make([]CloseSubOrders, len(o.SubOrders))
		for i, item := range o.SubOrders {
			ret.SubOrders[i] = *item.Clone()
		}
	}

	return &ret
}

// CloseSubOrders 合单关闭订单子单信息
type CloseSubOrders struct {
	// 子单商户号, 子单发起方商户号，必须与发起方appid有绑定关系。
	Mchid *string `json:"mchid"`
	// 子单商户订单号
	OutTradeNo *string `json:"out_trade_no"`
	// 二级商户号
	SubMchid *string `json:"sub_mchid"`
	// 子商户应用ID, 服务商模式下，sub_mchid对应的sub_appid
	SubAppid *string `json:"sub_appid,omitempty"`
}

func (o CloseSubOrders) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Mchid != nil {
		toSerialize["mchid"] = o.Mchid
	}
	if o.OutTradeNo != nil {
		toSerialize["out_trade_no"] = o.OutTradeNo
	}
	if o.SubMchid != nil {
		toSerialize["sub_mchid"] = o.SubMchid
	}
	if o.SubAppid != nil {
		toSerialize["sub_appid"] = o.SubAppid
	}

	return json.Marshal(toSerialize)
}

func (o CloseSubOrders) String() string {
	var ret string
	if o.Mchid == nil {
		ret += "Mchid:<nil>, "
	} else {
		ret += fmt.Sprintf("Mchid:%v, ", *o.Mchid)
	}

	if o.OutTradeNo == nil {
		ret += "OutTradeNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutTradeNo:%v, ", *o.OutTradeNo)
	}

	if o.SubMchid == nil {
		ret += "SubMchid:<nil>, "
	} else {
		ret += fmt.Sprintf("SubMchid:%v, ", *o.SubMchid)
	}

	if o.SubAppid == nil {
		ret += "SubAppid:<nil>"
	} else {
		ret += fmt.Sprintf("SubAppid:%v", *o.SubAppid)
	}

	return fmt.Sprintf("CloseSubOrders{%s}", ret)
}

func (o CloseSubOrders) Clone() *CloseSubOrders {
	ret := CloseSubOrders{}
	if o.Mchid != nil {
		ret.Mchid = new(string)
		*ret.Mchid = *o.Mchid
	}

	if o.OutTradeNo != nil {
		ret.OutTradeNo = new(string)
		*ret.OutTradeNo = *o.OutTradeNo
	}

	if o.SubMchid != nil {
		ret.SubMchid = new(string)
		*ret.SubMchid = *o.SubMchid
	}

	if o.SubAppid != nil {
		ret.SubAppid = new(string)
		*ret.SubAppid = *o.SubAppid
	}

	return &ret
}

// QueryOrderByOutTradeNoRequest 合单查询参数
type QueryOrderByOutTradeNoRequest struct {
	// 合单商户订单号
	CombineOutTradeNo *string `json:"combine_out_trade_no"`
}

func (o QueryOrderByOutTradeNoRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.CombineOutTradeNo != nil {
		toSerialize["combine_out_trade_no"] = o.CombineOutTradeNo
	}

	return json.Marshal(toSerialize)
}

func (o QueryOrderByOutTradeNoRequest) String() string {
	var ret string
	if o.CombineOutTradeNo == nil {
		ret += "CombineOutTradeNo:<nil>"
	} else {
		ret += fmt.Sprintf("CombineOutTradeNo:%v", *o.CombineOutTradeNo)
	}
	return fmt.Sprintf("QueryOrderByOutTradeNoRequest{%s}", ret)
}

func (o QueryOrderByOutTradeNoRequest) Clone() *QueryOrderByOutTradeNoRequest {
	ret := QueryOrderByOutTradeNoRequest{}

	if o.CombineOutTradeNo != nil {
		ret.CombineOutTradeNo = new(string)
		*ret.CombineOutTradeNo = *o.CombineOutTradeNo
	}

	return &ret
}

// RefundApplyRequest 退款申请参数
type RefundApplyRequest struct {
	// 二级商户号
	SubMchid *string `json:"sub_mchid"`
	// 电商平台APPID
	SpAppid *string `json:"sp_appid"`
	// 二级商户APPID, 二级商户在微信申请公众号成功后分配的账户ID，需要电商平台侧配置绑定关系才能传参（即二级商户已绑定微信公众号时传入）。
	SubAppid *string `json:"sub_appid,omitempty"`
	// 子微信订单号, 原支付交易对应的微信订单号与OutTradeNo二选一。
	TransactionId *string `json:"transaction_id,omitempty"`
	// 子商户退款单号, 与TransactionId二选一
	OutTradeNo *string `json:"out_trade_no,omitempty"`
	// 商户退款单号
	OutRefundNo *string `json:"out_refund_no"`
	// 退款原因
	Reason *string `json:"reason,omitempty"`
	// 订单金额
	Amount *RefundApplyAmount `json:"amount"`
	// 退款结果回调url
	NotifyUrl *string `json:"notify_url,omitempty"`
	// 退款出资商户, 商平台垫资退款专用参数。需先确认已开通此功能后，才能使用。若需要开通，请联系微信支付客服。
	//枚举值：
	//REFUND_SOURCE_PARTNER_ADVANCE : 电商平台垫付，需要向微信支付申请开通
	//REFUND_SOURCE_SUB_MERCHANT : 二级商户，默认值
	RefundAccount *string `json:"refund_account,omitempty"`
	// 资金账户, 若订单处于待分账状态，且未指定垫资退款（即refund_account未指定为REFUND_SOURCE_PARTNER_ADVANCE），可以传入此参数，指定退款资金来源账户。
	// 当该字段不存在时，默认使用订单交易资金所在账户出款，即待分账时使用不可用余额的资金进行退款，已分账或无分账时使用可用余额的资金进行退款。 AVAILABLE：可用余额
	FundsAccount *string `json:"funds_account,omitempty"`
}

func (o RefundApplyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SubMchid != nil {
		toSerialize["sub_mchid"] = o.SubMchid
	}
	if o.SpAppid != nil {
		toSerialize["sp_appid"] = o.SpAppid
	}
	if o.SubAppid != nil {
		toSerialize["sub_appid"] = o.SubAppid
	}
	if o.TransactionId != nil {
		toSerialize["transaction_id"] = o.TransactionId
	}
	if o.OutTradeNo != nil {
		toSerialize["out_trade_no"] = o.OutTradeNo
	}
	if o.OutRefundNo != nil {
		toSerialize["out_refund_no"] = o.OutRefundNo
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.NotifyUrl != nil {
		toSerialize["notify_url"] = o.NotifyUrl
	}
	if o.RefundAccount != nil {
		toSerialize["refund_account"] = o.RefundAccount
	}
	if o.FundsAccount != nil {
		toSerialize["funds_account"] = o.FundsAccount
	}

	return json.Marshal(toSerialize)
}

func (o RefundApplyRequest) String() string {
	var ret string
	if o.SubMchid == nil {
		ret += "SubMchid:<nil>, "
	} else {
		ret += fmt.Sprintf("SubMchid:%v, ", *o.SubMchid)
	}

	if o.SpAppid == nil {
		ret += "SpAppid:<nil>, "
	} else {
		ret += fmt.Sprintf("SpAppid:%v, ", *o.SpAppid)
	}

	if o.SubAppid == nil {
		ret += "SubAppid:<nil>, "
	} else {
		ret += fmt.Sprintf("SubAppid:%v, ", *o.SubAppid)
	}

	if o.TransactionId == nil {
		ret += "TransactionId:<nil>, "
	} else {
		ret += fmt.Sprintf("TransactionId:%v, ", *o.TransactionId)
	}

	if o.OutTradeNo == nil {
		ret += "OutTradeNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutTradeNo:%v, ", *o.OutTradeNo)
	}

	if o.OutRefundNo == nil {
		ret += "OutRefundNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutRefundNo:%v, ", *o.OutRefundNo)
	}

	if o.Reason == nil {
		ret += "Reason:<nil>, "
	} else {
		ret += fmt.Sprintf("Reason:%v, ", *o.Reason)
	}

	ret += fmt.Sprintf("Amount:%v, ", *o.Amount)

	if o.NotifyUrl == nil {
		ret += "NotifyUrl:<nil>, "
	} else {
		ret += fmt.Sprintf("NotifyUrl:%v, ", *o.NotifyUrl)
	}

	if o.RefundAccount == nil {
		ret += "RefundAccount:<nil>, "
	} else {
		ret += fmt.Sprintf("RefundAccount:%v, ", *o.RefundAccount)
	}

	if o.FundsAccount == nil {
		ret += "FundsAccount:<nil>, "
	} else {
		ret += fmt.Sprintf("FundsAccount:%v, ", *o.FundsAccount)
	}

	return fmt.Sprintf("RefundApplyRequest{%s}", ret)
}

func (o RefundApplyRequest) Clone() *RefundApplyRequest {
	ret := RefundApplyRequest{}
	if o.SubMchid != nil {
		ret.SubMchid = new(string)
		*ret.SubMchid = *o.SubMchid
	}

	if o.SpAppid != nil {
		ret.SpAppid = new(string)
		*ret.SpAppid = *o.SpAppid
	}

	if o.SubAppid != nil {
		ret.SubAppid = new(string)
		*ret.SubAppid = *o.SubAppid
	}

	if o.TransactionId != nil {
		ret.TransactionId = new(string)
		*ret.TransactionId = *o.TransactionId
	}

	if o.OutTradeNo != nil {
		ret.OutTradeNo = new(string)
		*ret.OutTradeNo = *o.OutTradeNo
	}

	if o.OutRefundNo != nil {
		ret.OutRefundNo = new(string)
		*ret.OutRefundNo = *o.OutRefundNo
	}

	if o.Reason != nil {
		ret.Reason = new(string)
		*ret.Reason = *o.Reason
	}

	if o.Amount != nil {
		ret.Amount = o.Amount.Clone()
	}

	if o.NotifyUrl != nil {
		ret.NotifyUrl = new(string)
		*ret.NotifyUrl = *o.NotifyUrl
	}

	if o.RefundAccount != nil {
		ret.RefundAccount = new(string)
		*ret.RefundAccount = *o.RefundAccount
	}

	if o.FundsAccount != nil {
		ret.FundsAccount = new(string)
		*ret.FundsAccount = *o.FundsAccount
	}

	return &ret
}

type RefundApplyAmount struct {
	// 退款金额, 币种的最小单位，只能为整数，不能超过原订单支付金额。
	Refund *int64 `json:"refund"`
	// 原支付交易的订单总金额，币种的最小单位，只能为整数。
	Total *int64 `json:"total"`
	// 符合ISO 4217标准的三位字母代码，目前只支持人民币：CNY。
	Currency *string `json:"currency"`
}

func (o RefundApplyAmount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Refund != nil {
		toSerialize["refund"] = o.Refund
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}

	return json.Marshal(toSerialize)
}

func (o RefundApplyAmount) String() string {
	var ret string
	if o.Refund == nil {
		ret += "Refund:<nil>, "
	} else {
		ret += fmt.Sprintf("Refund:%v, ", *o.Refund)
	}

	if o.Total == nil {
		ret += "Total:<nil>, "
	} else {
		ret += fmt.Sprintf("Total:%v, ", *o.Total)
	}

	if o.Currency == nil {
		ret += "Currency:<nil>"
	} else {
		ret += fmt.Sprintf("Currency:%v", *o.Currency)
	}

	return fmt.Sprintf("RefundAmount{%s}", ret)
}

func (o RefundApplyAmount) Clone() *RefundApplyAmount {
	ret := RefundApplyAmount{}
	if o.Refund != nil {
		ret.Refund = new(int64)
		*ret.Refund = *o.Refund
	}

	if o.Total != nil {
		ret.Total = new(int64)
		*ret.Total = *o.Total
	}

	if o.Currency != nil {
		ret.Currency = new(string)
		*ret.Currency = *o.Currency
	}

	return &ret
}

// Refund 退款申请返回结果
type Refund struct {
	// 微信支付退款号
	RefundId *string `json:"refund_id"`
	// 商户系统内部的退款单号，商户系统内部唯一，同一退款单号多次请求只退一笔。
	OutRefundNo *string `json:"out_refund_no"`
	// 退款受理时间，遵循rfc3339标准格式，格式为YYYY-MM-DDTHH:mm:ss+TIMEZONE，YYYY-MM-DD表示年月日，T出现在字符串中，表示time元素的开头，HH:mm:ss表示时分秒，TIMEZONE表示时区（+08:00表示东八区时间，领先UTC 8小时，即北京时间）。例如：2015-05-20T13:29:35+08:00表示，北京时间2015年5月20日13点29分35秒。
	CreateTime *time.Time `json:"create_time"`
	// 金额详细信息
	Amount *RefundAmount `json:"amount"`
	// 优惠退款信息
	PromotionDetail []RefundPromotionDetail `json:"promotion_detail,omitempty"`
	// 退款出资商户, 商平台垫资退款专用参数。需先确认已开通此功能后，才能使用。若需要开通，请联系微信支付客服。
	//枚举值：
	//REFUND_SOURCE_PARTNER_ADVANCE : 电商平台垫付，需要向微信支付申请开通
	//REFUND_SOURCE_SUB_MERCHANT : 二级商户，默认值
	RefundAccount *string `json:"refund_account,omitempty"`
}

func (o Refund) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RefundId != nil {
		toSerialize["refund_id"] = o.RefundId
	}
	if o.OutRefundNo != nil {
		toSerialize["out_refund_no"] = o.OutRefundNo
	}
	if o.CreateTime != nil {
		toSerialize["create_time"] = o.CreateTime.Format(time.RFC3339)
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.PromotionDetail != nil {
		toSerialize["promotion_detail"] = o.PromotionDetail
	}
	if o.RefundAccount != nil {
		toSerialize["refund_account"] = o.RefundAccount
	}

	return json.Marshal(toSerialize)
}

func (o Refund) String() string {
	var ret string
	if o.RefundId == nil {
		ret += "RefundId:<nil>, "
	} else {
		ret += fmt.Sprintf("RefundId:%v, ", *o.RefundId)
	}

	if o.OutRefundNo == nil {
		ret += "OutRefundNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutRefundNo:%v, ", *o.OutRefundNo)
	}

	if o.CreateTime == nil {
		ret += "CreateTime:<nil>, "
	} else {
		ret += fmt.Sprintf("CreateTime:%v, ", *o.CreateTime)
	}

	ret += fmt.Sprintf("Amount:%v, ", o.Amount)

	ret += fmt.Sprintf("PromotionDetail:%v, ", o.PromotionDetail)

	if o.RefundAccount == nil {
		ret += "RefundAccount:<nil>"
	} else {
		ret += fmt.Sprintf("RefundAccount:%v", *o.RefundAccount)
	}

	return fmt.Sprintf("Refund{%s}", ret)
}

func (o Refund) Clone() *Refund {
	ret := Refund{}
	if o.RefundId != nil {
		ret.RefundId = new(string)
		*ret.RefundId = *o.RefundId
	}

	if o.OutRefundNo != nil {
		ret.OutRefundNo = new(string)
		*ret.OutRefundNo = *o.OutRefundNo
	}

	if o.CreateTime != nil {
		ret.CreateTime = new(time.Time)
		*ret.CreateTime = *o.CreateTime
	}

	if o.Amount != nil {
		ret.Amount = o.Amount.Clone()
	}

	if o.PromotionDetail != nil {
		ret.PromotionDetail = make([]RefundPromotionDetail, len(o.PromotionDetail))
		for i, item := range o.PromotionDetail {
			ret.PromotionDetail[i] = *item.Clone()
		}
	}

	if o.RefundAccount != nil {
		ret.RefundAccount = new(string)
		*ret.RefundAccount = *o.RefundAccount
	}

	return &ret
}

type RefundAmount struct {
	// 退款金额, 币种的最小单位，只能为整数，不能超过原订单支付金额。
	Refund *int64 `json:"refund"`
	// 退款给用户的金额，不包含所有优惠券金额。
	PayerRefund *int64 `json:"payer_refund"`
	// 优惠券的退款金额，原支付单的优惠按比例退款。
	DiscountRefund *int64 `json:"discount_refund"`
	// 符合ISO 4217标准的三位字母代码，目前只支持人民币：CNY。
	Currency *string `json:"currency"`
}

func (o RefundAmount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Refund != nil {
		toSerialize["refund"] = o.Refund
	}
	if o.PayerRefund != nil {
		toSerialize["payer_refund"] = o.PayerRefund
	}
	if o.DiscountRefund != nil {
		toSerialize["discount_refund"] = o.DiscountRefund
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}

	return json.Marshal(toSerialize)
}

func (o RefundAmount) String() string {
	var ret string
	if o.Refund == nil {
		ret += "Refund:<nil>, "
	} else {
		ret += fmt.Sprintf("Refund:%v, ", *o.Refund)
	}

	if o.PayerRefund == nil {
		ret += "PayerRefund:<nil>, "
	} else {
		ret += fmt.Sprintf("PayerRefund:%v, ", *o.PayerRefund)
	}

	if o.DiscountRefund == nil {
		ret += "DiscountRefund:<nil>, "
	} else {
		ret += fmt.Sprintf("DiscountRefund:%v, ", *o.DiscountRefund)
	}

	if o.Currency == nil {
		ret += "Currency:<nil>"
	} else {
		ret += fmt.Sprintf("Currency:%v", *o.Currency)
	}

	return fmt.Sprintf("RefundAmount{%s}", ret)
}

func (o RefundAmount) Clone() *RefundAmount {
	ret := RefundAmount{}
	if o.Refund != nil {
		ret.Refund = new(int64)
		*ret.Refund = *o.Refund
	}

	if o.PayerRefund != nil {
		ret.PayerRefund = new(int64)
		*ret.PayerRefund = *o.PayerRefund
	}

	if o.DiscountRefund != nil {
		ret.DiscountRefund = new(int64)
		*ret.DiscountRefund = *o.DiscountRefund
	}

	if o.Currency != nil {
		ret.Currency = new(string)
		*ret.Currency = *o.Currency
	}

	return &ret
}

type RefundPromotionDetail struct {
	// 券ID
	PromotionId *string `json:"promotion_id"`
	// 优惠范围
	Scope *string `json:"scope"`
	// 优惠类型
	Type *string `json:"type"`
	// 优惠券面额
	Amount *int64 `json:"amount"`
	// 优惠退款金额
	RefundAmount *int64 `json:"refund_amount"`
}

func (o RefundPromotionDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PromotionId != nil {
		toSerialize["promotion_id"] = o.PromotionId
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.RefundAmount != nil {
		toSerialize["refund_amount"] = o.RefundAmount
	}

	return json.Marshal(toSerialize)
}

func (o RefundPromotionDetail) String() string {
	var ret string
	if o.PromotionId == nil {
		ret += "PromotionId:<nil>, "
	} else {
		ret += fmt.Sprintf("PromotionId:%v, ", *o.PromotionId)
	}

	if o.Scope == nil {
		ret += "Scope:<nil>, "
	} else {
		ret += fmt.Sprintf("Scope:%v, ", *o.Scope)
	}

	if o.Type == nil {
		ret += "Type:<nil>, "
	} else {
		ret += fmt.Sprintf("Type:%v, ", *o.Type)
	}

	if o.Amount == nil {
		ret += "Amount:<nil>, "
	} else {
		ret += fmt.Sprintf("Amount:%v, ", *o.Amount)
	}

	if o.RefundAmount == nil {
		ret += "RefundAmount:<nil>"
	} else {
		ret += fmt.Sprintf("RefundAmount:%v", *o.RefundAmount)
	}

	return fmt.Sprintf("RefundPromotionDetail{%s}", ret)
}

func (o RefundPromotionDetail) Clone() *RefundPromotionDetail {
	ret := RefundPromotionDetail{}
	if o.PromotionId != nil {
		ret.PromotionId = new(string)
		*ret.PromotionId = *o.PromotionId
	}

	if o.Scope != nil {
		ret.Scope = new(string)
		*ret.Scope = *o.Scope
	}

	if o.Type != nil {
		ret.Type = new(string)
		*ret.Type = *o.Type
	}

	if o.Amount != nil {
		ret.Amount = new(int64)
		*ret.Amount = *o.Amount
	}

	if o.RefundAmount != nil {
		ret.RefundAmount = new(int64)
		*ret.RefundAmount = *o.RefundAmount
	}

	return &ret
}

// QueryByOutRefundNoRequest 通过商户退款单号查询退款参数
type QueryByOutRefundNoRequest struct {
	// 商户系统内部的退款单号，商户系统内部唯一，同一退款单号多次请求只退一笔。
	OutRefundNo *string `json:"out_refund_no"`
	// 微信支付分配给二级商户的商户号。
	SubMchid *string `json:"sub_mchid"`
}

func (o QueryByOutRefundNoRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.OutRefundNo != nil {
		toSerialize["out_refund_no"] = o.OutRefundNo
	}

	if o.SubMchid != nil {
		toSerialize["sub_mchid"] = o.SubMchid
	}
	return json.Marshal(toSerialize)
}

func (o QueryByOutRefundNoRequest) String() string {
	var ret string
	if o.OutRefundNo == nil {
		ret += "OutRefundNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutRefundNo:%v, ", *o.OutRefundNo)
	}

	if o.SubMchid == nil {
		ret += "SubMchid:<nil>"
	} else {
		ret += fmt.Sprintf("SubMchid:%v", *o.SubMchid)
	}

	return fmt.Sprintf("QueryByOutRefundNoRequest{%s}", ret)
}

func (o QueryByOutRefundNoRequest) Clone() *QueryByOutRefundNoRequest {
	ret := QueryByOutRefundNoRequest{}

	if o.OutRefundNo != nil {
		ret.OutRefundNo = new(string)
		*ret.OutRefundNo = *o.OutRefundNo
	}

	if o.SubMchid != nil {
		ret.SubMchid = new(string)
		*ret.SubMchid = *o.SubMchid
	}

	return &ret
}

// QueryByRefundIdRequest 通过微信支付退款单号查询退款参数
type QueryByRefundIdRequest struct {
	// 商户系统内部的退款单号，商户系统内部唯一，同一退款单号多次请求只退一笔。
	RefundId *string `json:"refund_id"`
	// 微信支付分配给二级商户的商户号。
	SubMchid *string `json:"sub_mchid"`
}

func (o QueryByRefundIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.RefundId != nil {
		toSerialize["refund_id"] = o.RefundId
	}

	if o.SubMchid != nil {
		toSerialize["sub_mchid"] = o.SubMchid
	}
	return json.Marshal(toSerialize)
}

func (o QueryByRefundIdRequest) String() string {
	var ret string
	if o.RefundId == nil {
		ret += "RefundId:<nil>, "
	} else {
		ret += fmt.Sprintf("RefundId:%v, ", *o.RefundId)
	}

	if o.SubMchid == nil {
		ret += "SubMchid:<nil>"
	} else {
		ret += fmt.Sprintf("SubMchid:%v", *o.SubMchid)
	}

	return fmt.Sprintf("QueryByRefundIdRequest{%s}", ret)
}

func (o QueryByRefundIdRequest) Clone() *QueryByRefundIdRequest {
	ret := QueryByRefundIdRequest{}

	if o.RefundId != nil {
		ret.RefundId = new(string)
		*ret.RefundId = *o.RefundId
	}

	if o.SubMchid != nil {
		ret.SubMchid = new(string)
		*ret.SubMchid = *o.SubMchid
	}

	return &ret
}

// RefundNotify 退款回调参数
type RefundNotify struct {
	// 电商平台商户号
	SpMchid *string `json:"sp_mchid"`
	// 二级商户号
	SubMchid *string `json:"sub_mchid"`
	// 商户订单号
	TransactionId *string `json:"transaction_id"`
	// 商户退款单号
	OutRefundNo *string `json:"out_refund_no"`
	// 微信退款单号
	RefundId *string `json:"refund_id"`
	// 退款状态
	RefundStatus *string `json:"refund_status"`
	// 退款成功时间, 遵循rfc3339标准格式，
	SuccessTime *string `json:"success_time"`
	// 退款入账账户
	UserReceivedAccount *string `json:"user_received_account"`
	// 金额信息
	Amount struct {
		// 订单总金额，单位为分，只能为整数，详见支付金额
		Total *int64 `json:"total"`
		// 退款金额，币种的最小单位，只能为整数，不能超过原订单支付金额，如果有使用券，后台会按比例退。
		Refund *int64 `json:"refund"`
		// 用户实际支付金额，单位为分，只能为整数，详见支付金额
		PayerTotal *int64 `json:"payer_total"`
		// 退款给用户的金额，不包含所有优惠券金额
		PayerRefund *int64 `json:"payer_refund"`
	} `json:"amount"`
	// 退款出资商户
	RefundAccount *string `json:"refund_account"`
}
