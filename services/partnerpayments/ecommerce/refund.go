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

// 退款相关

// RefundApply 退款申请
func (a *RefundsApiService) RefundApply(ctx context.Context, req RefundApplyRequest) (resp *Refund, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := consts.WechatPayAPIServer + "/v3/ecommerce/refunds/apply"
	// Make sure All Required Params are properly set

	// Setup Body Params
	localVarPostBody = req

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract Refund from Http Response
	resp = new(Refund)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// QueryByOutRefundNo 通过商户退款单号查询退款
func (a *RefundsApiService) QueryByOutRefundNo(ctx context.Context, req QueryByOutRefundNoRequest) (resp *Refund, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.OutRefundNo == nil {
		return nil, nil, fmt.Errorf("field `OutRefundNo` is required and must be specified in QueryByOutRefundNoRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/ecommerce/refunds/out-refund-no/{out_refund_no}"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"out_refund_no"+"}", neturl.PathEscape(core.ParameterToString(*req.OutRefundNo, "")), -1)

	// Make sure All Required Params are properly set

	// Setup Query Params
	localVarQueryParams = neturl.Values{}
	if req.SubMchid != nil {
		localVarQueryParams.Add("sub_mchid", core.ParameterToString(*req.SubMchid, ""))
	}

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract Refund from Http Response
	resp = new(Refund)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// QueryByRefundId 通过微信支付退款单号查询退款
func (a *RefundsApiService) QueryByRefundId(ctx context.Context, req QueryByRefundIdRequest) (resp *Refund, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.RefundId == nil {
		return nil, nil, fmt.Errorf("field `RefundId` is required and must be specified in QueryByOutRefundNoRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/ecommerce/refunds/id/{refund_id}"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"refund_id"+"}", neturl.PathEscape(core.ParameterToString(*req.RefundId, "")), -1)

	// Make sure All Required Params are properly set

	// Setup Query Params
	localVarQueryParams = neturl.Values{}
	if req.SubMchid != nil {
		localVarQueryParams.Add("sub_mchid", core.ParameterToString(*req.SubMchid, ""))
	}

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract Refund from Http Response
	resp = new(Refund)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
