package ecommerce

import (
	"context"
	"fmt"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	nethttp "net/http"
	neturl "net/url"
	"strings"
)

// JsApiPrepay 小程序合单支付，同时也是JSAPI合单支付，两者接口一样
func (a *ApiService) JsApiPrepay(ctx context.Context, req JsApiPrepayRequest) (resp *JsApiPrepayResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)
	localVarPath := consts.WechatPayAPIServer + "/v3/combine-transactions/jsapi"
	localVarPostBody = req
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)
	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract PrepayResponse from Http Response
	resp = new(JsApiPrepayResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil

}

// CloseOrder 关闭订单, 合单支付关单接口不支持关闭部分子单， 关单的主单商户号、单号、子单个数、子单商户号、子单单号必须与下单时完全一致。
func (a *ApiService) CloseOrder(ctx context.Context, req CloseOrderRequest) (result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.CombineOutTradeNo == nil {
		return nil, fmt.Errorf("field `CombineOutTradeNo` is required and must be specified in CloseOrderRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/combine-transactions/out-trade-no/{combine_out_trade_no}/close"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"combine_out_trade_no"+"}", neturl.PathEscape(core.ParameterToString(*req.CombineOutTradeNo, "")), -1)

	// Make sure All Required Params are properly set

	// Setup Body Params
	localVarPostBody = &CloseRequest{
		CombineAppid: req.CombineAppid,
		SubOrders:    req.SubOrders,
	}

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return result, err
	}

	return result, nil
}

// QueryOrderByOutTradeNo 商户订单号查询订单
func (a *ApiService) QueryOrderByOutTradeNo(ctx context.Context, req QueryOrderByOutTradeNoRequest) (resp *Transaction, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.CombineOutTradeNo == nil {
		return nil, nil, fmt.Errorf("field `CombineOutTradeNo` is required and must be specified in QueryOrderByOutTradeNoRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/combine-transactions/out-trade-no/{combine_out_trade_no}"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"combine_out_trade_no"+"}", neturl.PathEscape(core.ParameterToString(*req.CombineOutTradeNo, "")), -1)

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract partnerpayments.Transaction from Http Response
	resp = new(Transaction)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
