// Code generated by lark_sdk_gen. DO NOT EDIT.
/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lark

import (
	"context"
)

// GetSheetFilterViewCondition 获取筛选视图某列的筛选条件信息。
//
// 筛选条件含义参考 [筛选视图的筛选条件指南](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/filter-view-condition-user-guide)
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/get
func (r *DriveService) GetSheetFilterViewCondition(ctx context.Context, request *GetSheetFilterViewConditionReq, options ...MethodOptionFunc) (*GetSheetFilterViewConditionResp, *Response, error) {
	if r.cli.mock.mockDriveGetSheetFilterViewCondition != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#GetSheetFilterViewCondition mock enable")
		return r.cli.mock.mockDriveGetSheetFilterViewCondition(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "GetSheetFilterViewCondition",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id/conditions/:condition_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getSheetFilterViewConditionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveGetSheetFilterViewCondition mock DriveGetSheetFilterViewCondition method
func (r *Mock) MockDriveGetSheetFilterViewCondition(f func(ctx context.Context, request *GetSheetFilterViewConditionReq, options ...MethodOptionFunc) (*GetSheetFilterViewConditionResp, *Response, error)) {
	r.mockDriveGetSheetFilterViewCondition = f
}

// UnMockDriveGetSheetFilterViewCondition un-mock DriveGetSheetFilterViewCondition method
func (r *Mock) UnMockDriveGetSheetFilterViewCondition() {
	r.mockDriveGetSheetFilterViewCondition = nil
}

// GetSheetFilterViewConditionReq ...
type GetSheetFilterViewConditionReq struct {
	SpreadSheetToken string `path:"spreadsheet_token" json:"-"` // 表格 token, 示例值："shtcnmBA*****yGehy8"
	SheetID          string `path:"sheet_id" json:"-"`          // 子表 id, 示例值："0b**12"
	FilterViewID     string `path:"filter_view_id" json:"-"`    // 筛选视图 id, 示例值："pH9hbVcCXA"
	ConditionID      string `path:"condition_id" json:"-"`      // 需要查询筛选条件的列字母号, 示例值："E"
}

// getSheetFilterViewConditionResp ...
type getSheetFilterViewConditionResp struct {
	Code int64                            `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                           `json:"msg,omitempty"`  // 错误描述
	Data *GetSheetFilterViewConditionResp `json:"data,omitempty"`
}

// GetSheetFilterViewConditionResp ...
type GetSheetFilterViewConditionResp struct {
	Condition *GetSheetFilterViewConditionRespCondition `json:"condition,omitempty"` // 筛选的条件
}

// GetSheetFilterViewConditionRespCondition ...
type GetSheetFilterViewConditionRespCondition struct {
	ConditionID string   `json:"condition_id,omitempty"` // 设置筛选条件的列，使用字母号
	FilterType  string   `json:"filter_type,omitempty"`  // 筛选类型
	CompareType string   `json:"compare_type,omitempty"` // 比较类型
	Expected    []string `json:"expected,omitempty"`     // 筛选参数
}
