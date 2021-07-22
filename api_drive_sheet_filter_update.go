// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// UpdateSheetFilter 参数值可参考[筛选指南](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter/filter-user-guide)
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter/update
func (r *DriveService) UpdateSheetFilter(ctx context.Context, request *UpdateSheetFilterReq, options ...MethodOptionFunc) (*UpdateSheetFilterResp, *Response, error) {
	if r.cli.mock.mockDriveUpdateSheetFilter != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#UpdateSheetFilter mock enable")
		return r.cli.mock.mockDriveUpdateSheetFilter(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "UpdateSheetFilter",
		Method:                "PUT",
		URL:                   "https://open.feishu.cn/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateSheetFilterResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDriveUpdateSheetFilter(f func(ctx context.Context, request *UpdateSheetFilterReq, options ...MethodOptionFunc) (*UpdateSheetFilterResp, *Response, error)) {
	r.mockDriveUpdateSheetFilter = f
}

func (r *Mock) UnMockDriveUpdateSheetFilter() {
	r.mockDriveUpdateSheetFilter = nil
}

type UpdateSheetFilterReq struct {
	SpreadSheetToken string                         `path:"spreadsheet_token" json:"-"` // 表格 token, 示例值："shtcnmBA\*****yGehy8"
	SheetID          string                         `path:"sheet_id" json:"-"`          // 子表 id, 示例值："0b\**12"
	Col              string                         `json:"col,omitempty"`              // 更新筛选条件的列, 示例值："E"
	Condition        *UpdateSheetFilterReqCondition `json:"condition,omitempty"`        // 筛选条件
}

type UpdateSheetFilterReqCondition struct {
	FilterType  string   `json:"filter_type,omitempty"`  // 筛选类型, 示例值："number"
	CompareType *string  `json:"compare_type,omitempty"` // 比较类型, 示例值："less"
	Expected    []string `json:"expected,omitempty"`     // 筛选参数, 示例值：6
}

type updateSheetFilterResp struct {
	Code int64                  `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *UpdateSheetFilterResp `json:"data,omitempty"`
}

type UpdateSheetFilterResp struct{}
