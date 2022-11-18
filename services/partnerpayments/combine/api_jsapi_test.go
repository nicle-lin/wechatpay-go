package combine

import (
	"context"
	"fmt"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"log"
	"testing"
)

func TestApiService_JsApiPrepay(t *testing.T) {
	var (
		mchID                      string = "*******" // 商户号
		mchCertificateSerialNumber string = "*******" // 商户证书序列号
		mchAPIv3Key                string = "*******" // 商户APIv3密钥
		SpAppid                           = "*******" // 服务商申请的公众号、小程序appId
		SpMchid                           = "*******" // 服务商商户号

		// 蜂办的
		SubAppid = "*******" // 特约商户(子商户)申请的公众号、小程序appId
		SubMchid = "*******" // 特约商户(子商户)的商户号

		// 仲量联行
		zhongMchid = "*******"
		zhongAppid = "*******"
	)
	fmt.Println(SpAppid, SpMchid, SubAppid, SubMchid, zhongAppid, zhongMchid)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/Users/baolin/cert/wechat_cert/apiclient_key.pem")
	if err != nil {
		log.Fatal("load merchant private key error")
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Fatalf("new wechat pay client err:%s", err)
	}

	// 得到prepay_id，以及调起支付所需的参数和签名

	svc := ApiService{Client: client}
	resp, _, err := svc.JsApiPrepay(ctx, JsApiPrepayRequest{
		CombineAppid:      core.String(SpAppid), // 合单发起方的appid
		CombineMchid:      core.String(SpMchid), // 服务商商户号
		CombineOutTradeNo: core.String("DD2022102415363000"),
		// SceneInfo:         nil,
		SubOrders: []JsApiSubOrders{
			{
				Mchid:  core.String(SpMchid), // 服务商商户号。
				Attach: core.String("测试合单支付"),
				Amount: &JsApiAmount{
					TotalAmount: core.Int64(1),
					Currency:    core.String("CNY"),
				},
				OutTradeNo:  core.String("DD2022102415363000"),
				GoodsTag:    core.String("HY"),
				SubMchid:    core.String(zhongMchid),
				Description: core.String("测试商品"),
				// SettleInfo: JsApiSettleInfo{},
				// SubAppid: core.String(),
			},
		},
		CombinePayerInfo: &JsApiCombinePayerInfo{
			//Openid: core.String("*******"),
			SubOpenid: core.String("o*******"),
			// SubOpenid: nil,
		},
		//TimeStart:        nil,
		//TimeExpire:       nil,
		NotifyUrl: core.String("*******"),
	})
	if err != nil {
		fmt.Println("----------------err:", err.Error())
		return
	}
	fmt.Println("-----------prepay_id:", resp.PrepayId)
}
