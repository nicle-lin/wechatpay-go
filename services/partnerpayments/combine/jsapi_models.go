package combine

import (
	"encoding/json"
	"fmt"
)

type JsApiPrepayRequest struct {
	CombineAppid      *string                `json:"combine_appid"`
	CombineMchid      *string                `json:"combine_mchid"`
	CombineOutTradeNo *string                `json:"combine_out_trade_no"`
	SceneInfo         *JsApiSceneInfo        `json:"scene_info,omitempty"`
	SubOrders         []JsApiSubOrders       `json:"sub_orders"`
	CombinePayerInfo  *JsApiCombinePayerInfo `json:"combine_payer_info"`
	TimeStart         *string                `json:"time_start,omitempty"`
	TimeExpire        *string                `json:"time_expire,omitempty"`
	NotifyUrl         *string                `json:"notify_url"`
}

func (o JsApiPrepayRequest) MarshalJSON() ([]byte, error) {
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

	if o.TimeStart != nil {
		toSerialize["time_start"] = o.TimeStart
	}

	if o.TimeExpire != nil {
		toSerialize["time_expire"] = o.TimeExpire
	}

	if o.NotifyUrl != nil {
		toSerialize["notify_url"] = o.NotifyUrl
	}

	return json.Marshal(toSerialize)
}

func (o JsApiPrepayRequest) String() string {
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
	ret += fmt.Sprintf("CombinePayerInfo:%v, ", o.CombinePayerInfo)

	if o.TimeStart == nil {
		ret += "TimeStart:<nil>, "
	} else {
		ret += fmt.Sprintf("TimeStart:%v, ", *o.TimeStart)
	}

	if o.TimeExpire == nil {
		ret += "TimeExpire:<nil>, "
	} else {
		ret += fmt.Sprintf("TimeExpire:%v, ", *o.TimeExpire)
	}

	if o.NotifyUrl == nil {
		ret += "NotifyUrl:<nil>, "
	} else {
		ret += fmt.Sprintf("NotifyUrl:%v", *o.NotifyUrl)
	}

	return fmt.Sprintf("Transaction{%s}", ret)
}

func (o JsApiPrepayRequest) Clone() *JsApiPrepayRequest {
	ret := JsApiPrepayRequest{}

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
		ret.SubOrders = make([]JsApiSubOrders, len(o.SubOrders))
		for i, item := range o.SubOrders {
			ret.SubOrders[i] = *item.Clone()
		}
	}

	if o.CombinePayerInfo != nil {
		ret.CombinePayerInfo = o.CombinePayerInfo.Clone()
	}

	if o.TimeStart != nil {
		ret.TimeStart = new(string)
		*ret.TimeStart = *o.TimeStart
	}

	if o.TimeExpire != nil {
		ret.TimeExpire = new(string)
		*ret.TimeExpire = *o.TimeExpire
	}

	if o.NotifyUrl != nil {
		ret.NotifyUrl = new(string)
		*ret.NotifyUrl = *o.NotifyUrl
	}

	return &ret
}

type JsApiSceneInfo struct {
	DeviceId      *string `json:"device_id,omitempty"`
	PayerClientIp *string `json:"payer_client_ip"`
}

func (o JsApiSceneInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.PayerClientIp == nil {
		return nil, fmt.Errorf("field `PayerClientIp` is required and must be specified in SceneInfo")
	}
	toSerialize["payer_client_ip"] = o.PayerClientIp

	if o.DeviceId != nil {
		toSerialize["device_id"] = o.DeviceId
	}

	return json.Marshal(toSerialize)
}

func (o JsApiSceneInfo) String() string {
	var ret string
	if o.PayerClientIp == nil {
		ret += "PayerClientIp:<nil>, "
	} else {
		ret += fmt.Sprintf("PayerClientIp:%v, ", *o.PayerClientIp)
	}

	if o.DeviceId == nil {
		ret += "DeviceId:<nil>, "
	} else {
		ret += fmt.Sprintf("DeviceId:%v", *o.DeviceId)
	}

	return fmt.Sprintf("SceneInfo{%s}", ret)
}

func (o JsApiSceneInfo) Clone() *JsApiSceneInfo {
	ret := JsApiSceneInfo{}

	if o.PayerClientIp != nil {
		ret.PayerClientIp = new(string)
		*ret.PayerClientIp = *o.PayerClientIp
	}

	if o.DeviceId != nil {
		ret.DeviceId = new(string)
		*ret.DeviceId = *o.DeviceId
	}

	return &ret
}

type JsApiSubOrders struct {
	Mchid       *string          `json:"mchid"`
	Attach      *string          `json:"attach"`
	Amount      *JsApiAmount     `json:"amount"`
	OutTradeNo  *string          `json:"out_trade_no"`
	GoodsTag    *string          `json:"goods_tag,omitempty"`
	SubMchid    *string          `json:"sub_mchid"`
	Description *string          `json:"description"`
	SettleInfo  *JsApiSettleInfo `json:"settle_info,omitempty"`
	SubAppid    *string          `json:"sub_appid,omitempty"`
}

func (o JsApiSubOrders) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Mchid != nil {
		toSerialize["mchid"] = o.Mchid
	}

	if o.Attach != nil {
		toSerialize["attach"] = o.Attach
	}

	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}

	if o.OutTradeNo != nil {
		toSerialize["out_trade_no"] = o.OutTradeNo
	}

	if o.GoodsTag != nil {
		toSerialize["goods_tag"] = o.GoodsTag
	}

	if o.SubMchid != nil {
		toSerialize["sub_mchid"] = o.SubMchid
	}

	if o.Description != nil {
		toSerialize["description"] = o.Description
	}

	if o.SettleInfo != nil {
		toSerialize["settle_info"] = o.SettleInfo
	}

	if o.SubAppid != nil {
		toSerialize["sub_appid"] = o.SubAppid
	}

	return json.Marshal(toSerialize)
}

func (o JsApiSubOrders) String() string {
	var ret string
	if o.Mchid == nil {
		ret += "Mchid:<nil>, "
	} else {
		ret += fmt.Sprintf("Mchid:%v, ", *o.Mchid)
	}

	if o.Attach == nil {
		ret += "Attach:<nil>, "
	} else {
		ret += fmt.Sprintf("Attach:%v, ", *o.Attach)
	}

	if o.Amount == nil {
		ret += "Amount:<nil>, "
	} else {
		ret += fmt.Sprintf("Amount:%v, ", o.Amount)
	}

	if o.OutTradeNo == nil {
		ret += "OutTradeNo:<nil>, "
	} else {
		ret += fmt.Sprintf("OutTradeNo:%v, ", *o.OutTradeNo)
	}

	if o.GoodsTag == nil {
		ret += "GoodsTag:<nil>, "
	} else {
		ret += fmt.Sprintf("GoodsTag:%v, ", *o.GoodsTag)
	}

	if o.SubMchid == nil {
		ret += "SubMchid:<nil>, "
	} else {
		ret += fmt.Sprintf("SubMchid:%v, ", *o.SubMchid)
	}

	if o.Description == nil {
		ret += "Description:<nil>, "
	} else {
		ret += fmt.Sprintf("Description:%v, ", *o.Description)
	}

	if o.SettleInfo == nil {
		ret += "SettleInfo:<nil>, "
	} else {
		ret += fmt.Sprintf("SettleInfo:%v, ", o.SettleInfo)
	}

	if o.SubAppid == nil {
		ret += "SubAppid:<nil>"
	} else {
		ret += fmt.Sprintf("SubAppid:%v", *o.SubAppid)
	}

	return fmt.Sprintf("SubOrders{%s}", ret)
}

func (o JsApiSubOrders) Clone() *JsApiSubOrders {
	ret := JsApiSubOrders{}

	if o.Mchid != nil {
		ret.Mchid = new(string)
		*ret.Mchid = *o.Mchid
	}

	if o.Attach != nil {
		ret.Attach = new(string)
		*ret.Attach = *o.Attach
	}

	if o.Amount != nil {
		ret.Amount = o.Amount.Clone()
	}

	if o.OutTradeNo != nil {
		ret.OutTradeNo = new(string)
		*ret.OutTradeNo = *o.OutTradeNo
	}

	if o.GoodsTag != nil {
		ret.GoodsTag = new(string)
		*ret.GoodsTag = *o.GoodsTag
	}

	if o.SubMchid != nil {
		ret.SubMchid = new(string)
		*ret.SubMchid = *o.SubMchid
	}

	if o.Description != nil {
		ret.Description = new(string)
		*ret.Description = *o.Description
	}

	if o.SettleInfo != nil {
		ret.SettleInfo = o.SettleInfo.Clone()
	}

	if o.SubAppid != nil {
		ret.SubAppid = new(string)
		*ret.SubAppid = *o.SubAppid
	}

	return &ret
}

type JsApiAmount struct {
	TotalAmount *int64  `json:"total_amount"`
	Currency    *string `json:"currency"`
}

func (o JsApiAmount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.TotalAmount == nil {
		return nil, fmt.Errorf("field `TotalAmount` is required and must be specified in Amount")
	}
	toSerialize["total_amount"] = o.TotalAmount

	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	return json.Marshal(toSerialize)
}

func (o JsApiAmount) String() string {
	var ret string
	if o.TotalAmount == nil {
		ret += "TotalAmount:<nil>, "
	} else {
		ret += fmt.Sprintf("TotalAmount:%v, ", *o.TotalAmount)
	}

	if o.Currency == nil {
		ret += "Currency:<nil>"
	} else {
		ret += fmt.Sprintf("Currency:%v", *o.Currency)
	}

	return fmt.Sprintf("Amount{%s}", ret)
}

func (o JsApiAmount) Clone() *JsApiAmount {
	ret := JsApiAmount{}

	if o.TotalAmount != nil {
		ret.TotalAmount = new(int64)
		*ret.TotalAmount = *o.TotalAmount
	}

	if o.Currency != nil {
		ret.Currency = new(string)
		*ret.Currency = *o.Currency
	}

	return &ret
}

type JsApiSettleInfo struct {
	ProfitSharing *bool  `json:"profit_sharing,omitempty"`
	SubsidyAmount *int64 `json:"subsidy_amount,omitempty"`
}

func (o JsApiSettleInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.ProfitSharing != nil {
		toSerialize["profit_sharing"] = o.ProfitSharing
	}
	if o.SubsidyAmount != nil {
		toSerialize["subsidy_amount"] = o.SubsidyAmount
	}
	return json.Marshal(toSerialize)
}

func (o JsApiSettleInfo) String() string {
	var ret string
	if o.ProfitSharing == nil {
		ret += "ProfitSharing:<nil>"
	} else {
		ret += fmt.Sprintf("ProfitSharing:%v, ", *o.ProfitSharing)
	}

	if o.SubsidyAmount == nil {
		ret += "SubsidyAmount:<nil>"
	} else {
		ret += fmt.Sprintf("SubsidyAmount:%v", *o.SubsidyAmount)
	}

	return fmt.Sprintf("SettleInfo{%s}", ret)
}

func (o JsApiSettleInfo) Clone() *JsApiSettleInfo {
	ret := JsApiSettleInfo{}

	if o.ProfitSharing != nil {
		ret.ProfitSharing = new(bool)
		*ret.ProfitSharing = *o.ProfitSharing
	}
	if o.SubsidyAmount != nil {
		ret.SubsidyAmount = new(int64)
		*ret.SubsidyAmount = *o.SubsidyAmount
	}

	return &ret
}

type JsApiCombinePayerInfo struct {
	Openid    *string `json:"openid,omitempty"`
	SubOpenid *string `json:"sub_openid,omitempty"`
}

func (o JsApiCombinePayerInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Openid != nil {
		toSerialize["openid"] = o.Openid
	}

	if o.SubOpenid != nil {
		toSerialize["sub_openid"] = o.SubOpenid
	}
	return json.Marshal(toSerialize)
}

func (o JsApiCombinePayerInfo) String() string {
	var ret string
	if o.Openid == nil {
		ret += "Openid:<nil>, "
	} else {
		ret += fmt.Sprintf("Openid:%v, ", *o.Openid)
	}

	if o.SubOpenid == nil {
		ret += "SubOpenid:<nil>"
	} else {
		ret += fmt.Sprintf("SubOpenid:%v", *o.SubOpenid)
	}

	return fmt.Sprintf("Payer{%s}", ret)
}

func (o JsApiCombinePayerInfo) Clone() *JsApiCombinePayerInfo {
	ret := JsApiCombinePayerInfo{}

	if o.Openid != nil {
		ret.Openid = new(string)
		*ret.Openid = *o.Openid
	}

	if o.SubOpenid != nil {
		ret.SubOpenid = new(string)
		*ret.SubOpenid = *o.SubOpenid
	}

	return &ret
}

type JsApiPrepayResponse struct {
	// 预支付交易会话标识
	PrepayId *string `json:"prepay_id"`
}

func (o JsApiPrepayResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.PrepayId == nil {
		return nil, fmt.Errorf("field `PrepayId` is required and must be specified in PrepayResponse")
	}
	toSerialize["prepay_id"] = o.PrepayId
	return json.Marshal(toSerialize)
}

func (o JsApiPrepayResponse) String() string {
	var ret string
	if o.PrepayId == nil {
		ret += "PrepayId:<nil>"
	} else {
		ret += fmt.Sprintf("PrepayId:%v", *o.PrepayId)
	}

	return fmt.Sprintf("PrepayResponse{%s}", ret)
}

func (o JsApiPrepayResponse) Clone() *JsApiPrepayResponse {
	ret := JsApiPrepayResponse{}

	if o.PrepayId != nil {
		ret.PrepayId = new(string)
		*ret.PrepayId = *o.PrepayId
	}

	return &ret
}
