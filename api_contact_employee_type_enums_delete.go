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

// DeleteEmployeeTypeEnum 删除自定义人员类型
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/employee_type_enum/delete
func (r *ContactService) DeleteEmployeeTypeEnum(ctx context.Context, request *DeleteEmployeeTypeEnumReq, options ...MethodOptionFunc) (*DeleteEmployeeTypeEnumResp, *Response, error) {
	if r.cli.mock.mockContactDeleteEmployeeTypeEnum != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Contact#DeleteEmployeeTypeEnum mock enable")
		return r.cli.mock.mockContactDeleteEmployeeTypeEnum(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "DeleteEmployeeTypeEnum",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/employee_type_enums/:enum_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deleteEmployeeTypeEnumResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactDeleteEmployeeTypeEnum mock ContactDeleteEmployeeTypeEnum method
func (r *Mock) MockContactDeleteEmployeeTypeEnum(f func(ctx context.Context, request *DeleteEmployeeTypeEnumReq, options ...MethodOptionFunc) (*DeleteEmployeeTypeEnumResp, *Response, error)) {
	r.mockContactDeleteEmployeeTypeEnum = f
}

// UnMockContactDeleteEmployeeTypeEnum un-mock ContactDeleteEmployeeTypeEnum method
func (r *Mock) UnMockContactDeleteEmployeeTypeEnum() {
	r.mockContactDeleteEmployeeTypeEnum = nil
}

// DeleteEmployeeTypeEnumReq ...
type DeleteEmployeeTypeEnumReq struct {
	EnumID string `path:"enum_id" json:"-"` // 枚举值id, 示例值："exGeIjow7zIqWMy+ONkFxA=="
}

// deleteEmployeeTypeEnumResp ...
type deleteEmployeeTypeEnumResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *DeleteEmployeeTypeEnumResp `json:"data,omitempty"`
}

// DeleteEmployeeTypeEnumResp ...
type DeleteEmployeeTypeEnumResp struct {
}
