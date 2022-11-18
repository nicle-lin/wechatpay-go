// 微信支付服务商基础支付
// 合单支付
// 微信支付 API v3 服务商基础支付

package combine

import (
	"encoding/json"
	"fmt"
)

type Transaction struct {
	CombineAppid      *string           `json:"combine_appid,omitempty"`
	CombineMchid      *string           `json:"combine_mchid,omitempty"`
	CombineOutTradeNo *string           `json:"combine_out_trade_no,omitempty"`
	SceneInfo         *SceneInfo        `json:"scene_info,omitempty"`
	SubOrders         []SubOrders       `json:"sub_orders,omitempty"`
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
	Mchid           *string            `json:"mchid,omitempty"`
	TradeType       *string            `json:"trade_type,omitempty"`
	TradeState      *string            `json:"trade_state,omitempty"`
	BankType        *string            `json:"bank_type,omitempty"`
	Attach          *string            `json:"attach,omitempty"`
	SuccessTime     *string            `json:"success_time,omitempty"`
	TransactionId   *string            `json:"transaction_id,omitempty"`
	OutTradeNo      *string            `json:"out_trade_no,omitempty"`
	SubMchid        *string            `json:"sub_mchid,omitempty"`
	SubAppid        *string            `json:"sub_appid,omitempty"`
	Amount          *TransactionAmount `json:"amount,omitempty"`
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
	TotalAmount   *int64  `json:"total_amount,omitempty"`
	Currency      *string `json:"currency,omitempty"`
	PayerAmount   *int64  `json:"payer_amount,omitempty"`
	PayerCurrency *string `json:"payer_currency,omitempty"`
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
	CouponId *string `json:"coupon_id,omitempty"`
	// 优惠名称
	Name *string `json:"name,omitempty"`
	// GLOBAL：全场代金券；SINGLE：单品优惠
	Scope *string `json:"scope,omitempty"`
	// CASH：充值；NOCASH：预充值。
	Type *string `json:"type,omitempty"`
	// 优惠券面额
	Amount *int64 `json:"amount,omitempty"`
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
	Openid *string `json:"openid,omitempty"`
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
