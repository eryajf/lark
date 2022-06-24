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

// FindSheet 按照指定的条件查找子表的某个范围内的数据符合条件的单元格位置。请求体中的 range 和 find 字段为必填。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet/find
func (r *DriveService) FindSheet(ctx context.Context, request *FindSheetReq, options ...MethodOptionFunc) (*FindSheetResp, *Response, error) {
	if r.cli.mock.mockDriveFindSheet != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#FindSheet mock enable")
		return r.cli.mock.mockDriveFindSheet(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "FindSheet",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/find",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(findSheetResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveFindSheet mock DriveFindSheet method
func (r *Mock) MockDriveFindSheet(f func(ctx context.Context, request *FindSheetReq, options ...MethodOptionFunc) (*FindSheetResp, *Response, error)) {
	r.mockDriveFindSheet = f
}

// UnMockDriveFindSheet un-mock DriveFindSheet method
func (r *Mock) UnMockDriveFindSheet() {
	r.mockDriveFindSheet = nil
}

// FindSheetReq ...
type FindSheetReq struct {
	SpreadSheetToken string                     `path:"spreadsheet_token" json:"-"` // 表格的 token, 示例值："shtcnmBA*****yGehy8"
	SheetID          string                     `path:"sheet_id" json:"-"`          // 子表的 id, 示例值："0b**12"
	FindCondition    *FindSheetReqFindCondition `json:"find_condition,omitempty"`   // 查找条件
	Find             string                     `json:"find,omitempty"`             // 查找的字符串, 示例值："hello"
}

// FindSheetReqFindCondition ...
type FindSheetReqFindCondition struct {
	Range           string `json:"range,omitempty"`             // 查找范围, 示例值："0b**12!A1:H10"
	MatchCase       *bool  `json:"match_case,omitempty"`        // 是否忽略大小写, 示例值：true
	MatchEntireCell *bool  `json:"match_entire_cell,omitempty"` // 是否匹配整个单元格, 示例值：false
	SearchByRegex   *bool  `json:"search_by_regex,omitempty"`   // 是否为正则匹配, 示例值：false
	IncludeFormulas *bool  `json:"include_formulas,omitempty"`  // 是否搜索公式内容, 示例值：false
}

// findSheetResp ...
type findSheetResp struct {
	Code int64          `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string         `json:"msg,omitempty"`  // 错误描述
	Data *FindSheetResp `json:"data,omitempty"`
}

// FindSheetResp ...
type FindSheetResp struct {
	FindResult *FindSheetRespFindResult `json:"find_result,omitempty"` // 查找返回符合条件的信息
}

// FindSheetRespFindResult ...
type FindSheetRespFindResult struct {
	MatchedCells        []string `json:"matched_cells,omitempty"`         // 符合查找条件的单元格数组，不包含公式，例如["A1", "A2"...]
	MatchedFormulaCells []string `json:"matched_formula_cells,omitempty"` // 符合查找条件的含有公式的单元格数组，例如["B3", "H7"...]
	RowsCount           int64    `json:"rows_count,omitempty"`            // 符合查找条件的总行数
}
