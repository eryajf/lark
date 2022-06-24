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

// UpdateSheetDimensionRange
//
// 该接口用于根据 spreadsheetToken 和维度信息更新隐藏行列、单元格大小；单次操作不超过5000行或列。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uYjMzUjL2IzM14iNyMTN
func (r *DriveService) UpdateSheetDimensionRange(ctx context.Context, request *UpdateSheetDimensionRangeReq, options ...MethodOptionFunc) (*UpdateSheetDimensionRangeResp, *Response, error) {
	if r.cli.mock.mockDriveUpdateSheetDimensionRange != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#UpdateSheetDimensionRange mock enable")
		return r.cli.mock.mockDriveUpdateSheetDimensionRange(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "UpdateSheetDimensionRange",
		Method:                "PUT",
		URL:                   r.cli.openBaseURL + "/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/dimension_range",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateSheetDimensionRangeResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveUpdateSheetDimensionRange mock DriveUpdateSheetDimensionRange method
func (r *Mock) MockDriveUpdateSheetDimensionRange(f func(ctx context.Context, request *UpdateSheetDimensionRangeReq, options ...MethodOptionFunc) (*UpdateSheetDimensionRangeResp, *Response, error)) {
	r.mockDriveUpdateSheetDimensionRange = f
}

// UnMockDriveUpdateSheetDimensionRange un-mock DriveUpdateSheetDimensionRange method
func (r *Mock) UnMockDriveUpdateSheetDimensionRange() {
	r.mockDriveUpdateSheetDimensionRange = nil
}

// UpdateSheetDimensionRangeReq ...
type UpdateSheetDimensionRangeReq struct {
	SpreadSheetToken    string                                           `path:"spreadsheetToken" json:"-"`     // spreadsheet 的 token，获取方式见[在线表格开发指南](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview)
	Dimension           *UpdateSheetDimensionRangeReqDimension           `json:"dimension,omitempty"`           // 需要更新行列的维度信息
	DimensionProperties *UpdateSheetDimensionRangeReqDimensionProperties `json:"dimensionProperties,omitempty"` // 需要更新行列的属性
}

// UpdateSheetDimensionRangeReqDimension ...
type UpdateSheetDimensionRangeReqDimension struct {
	SheetID        string  `json:"sheetId,omitempty"`        // sheetId
	MajorDimension *string `json:"majorDimension,omitempty"` // 默认 ROWS ，可选 ROWS、COLUMNS
	StartIndex     int64   `json:"startIndex"`               // 开始的位置
	EndIndex       int64   `json:"endIndex,omitempty"`       // 结束的位置
}

// UpdateSheetDimensionRangeReqDimensionProperties ...
type UpdateSheetDimensionRangeReqDimensionProperties struct {
	Visible   *bool  `json:"visible,omitempty"`   // true 为显示，false 为隐藏行列
	FixedSize *int64 `json:"fixedSize,omitempty"` // 行/列的大小
}

// updateSheetDimensionRangeResp ...
type updateSheetDimensionRangeResp struct {
	Code int64                          `json:"code,omitempty"`
	Msg  string                         `json:"msg,omitempty"`
	Data *UpdateSheetDimensionRangeResp `json:"data,omitempty"`
}

// UpdateSheetDimensionRangeResp ...
type UpdateSheetDimensionRangeResp struct {
}
