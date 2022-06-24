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

// UpdateSheetFloatImage 更新已有的浮动图片位置和宽高，包括 range、width、height、offset_x 和 offset_y，不包括 float_image_id 和 float_image_token。
//
// 浮动图片更新参考：[浮动图片指南](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/float-image-user-guide)
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/patch
func (r *DriveService) UpdateSheetFloatImage(ctx context.Context, request *UpdateSheetFloatImageReq, options ...MethodOptionFunc) (*UpdateSheetFloatImageResp, *Response, error) {
	if r.cli.mock.mockDriveUpdateSheetFloatImage != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#UpdateSheetFloatImage mock enable")
		return r.cli.mock.mockDriveUpdateSheetFloatImage(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "UpdateSheetFloatImage",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images/:float_image_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateSheetFloatImageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveUpdateSheetFloatImage mock DriveUpdateSheetFloatImage method
func (r *Mock) MockDriveUpdateSheetFloatImage(f func(ctx context.Context, request *UpdateSheetFloatImageReq, options ...MethodOptionFunc) (*UpdateSheetFloatImageResp, *Response, error)) {
	r.mockDriveUpdateSheetFloatImage = f
}

// UnMockDriveUpdateSheetFloatImage un-mock DriveUpdateSheetFloatImage method
func (r *Mock) UnMockDriveUpdateSheetFloatImage() {
	r.mockDriveUpdateSheetFloatImage = nil
}

// UpdateSheetFloatImageReq ...
type UpdateSheetFloatImageReq struct {
	SpreadSheetToken string   `path:"spreadsheet_token" json:"-"`  // 表格 token, 示例值："shtcnmBA*****yGehy8"
	SheetID          string   `path:"sheet_id" json:"-"`           // 子表 id, 示例值："0b**12"
	FloatImageID     string   `path:"float_image_id" json:"-"`     // 浮动图片 id, 示例值："ye06SS14ph"
	FloatImageToken  *string  `json:"float_image_token,omitempty"` // 【更新时不用传，创建需要】浮动图片 token，需要先上传图片到表格获得此 token 之后再进行浮动图片的相关操作, 示例值："boxbcbQsaSqIXsxxxxx1HCPJFbh"
	Range            *string  `json:"range,omitempty"`             // 浮动图片的左上角单元格定位，只支持一个单元格, 示例值："0b**12!A1:A1"
	Width            *float64 `json:"width,omitempty"`             // 浮动图片的宽度，大于等于 20px, 示例值：100
	Height           *float64 `json:"height,omitempty"`            // 浮动图片的高度，大于等于 20px, 示例值：100
	OffsetX          *float64 `json:"offset_x,omitempty"`          // 浮动图片左上角所在位置相对于所在单元格左上角的横向偏移，大于等于0且小于所在单元格的宽度, 示例值：0
	OffsetY          *float64 `json:"offset_y,omitempty"`          // 浮动图片左上角所在位置相对于所在单元格左上角的纵向偏移，大于等于0且小于所在单元格的高度, 示例值：0
}

// updateSheetFloatImageResp ...
type updateSheetFloatImageResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *UpdateSheetFloatImageResp `json:"data,omitempty"`
}

// UpdateSheetFloatImageResp ...
type UpdateSheetFloatImageResp struct {
	FloatImage *UpdateSheetFloatImageRespFloatImage `json:"float_image,omitempty"` // 浮动图片信息
}

// UpdateSheetFloatImageRespFloatImage ...
type UpdateSheetFloatImageRespFloatImage struct {
	FloatImageID    string  `json:"float_image_id,omitempty"`    // 浮动图片 id
	FloatImageToken string  `json:"float_image_token,omitempty"` // 【更新时不用传，创建需要】浮动图片 token，需要先上传图片到表格获得此 token 之后再进行浮动图片的相关操作
	Range           string  `json:"range,omitempty"`             // 浮动图片的左上角单元格定位，只支持一个单元格
	Width           float64 `json:"width,omitempty"`             // 浮动图片的宽度，大于等于 20px
	Height          float64 `json:"height,omitempty"`            // 浮动图片的高度，大于等于 20px
	OffsetX         float64 `json:"offset_x,omitempty"`          // 浮动图片左上角所在位置相对于所在单元格左上角的横向偏移，大于等于0且小于所在单元格的宽度
	OffsetY         float64 `json:"offset_y,omitempty"`          // 浮动图片左上角所在位置相对于所在单元格左上角的纵向偏移，大于等于0且小于所在单元格的高度
}
