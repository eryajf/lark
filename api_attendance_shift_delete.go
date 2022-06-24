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

// DeleteAttendanceShift 通过班次 ID 删除班次。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/shift/delete
func (r *AttendanceService) DeleteAttendanceShift(ctx context.Context, request *DeleteAttendanceShiftReq, options ...MethodOptionFunc) (*DeleteAttendanceShiftResp, *Response, error) {
	if r.cli.mock.mockAttendanceDeleteAttendanceShift != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#DeleteAttendanceShift mock enable")
		return r.cli.mock.mockAttendanceDeleteAttendanceShift(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Attendance",
		API:                   "DeleteAttendanceShift",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/attendance/v1/shifts/:shift_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deleteAttendanceShiftResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAttendanceDeleteAttendanceShift mock AttendanceDeleteAttendanceShift method
func (r *Mock) MockAttendanceDeleteAttendanceShift(f func(ctx context.Context, request *DeleteAttendanceShiftReq, options ...MethodOptionFunc) (*DeleteAttendanceShiftResp, *Response, error)) {
	r.mockAttendanceDeleteAttendanceShift = f
}

// UnMockAttendanceDeleteAttendanceShift un-mock AttendanceDeleteAttendanceShift method
func (r *Mock) UnMockAttendanceDeleteAttendanceShift() {
	r.mockAttendanceDeleteAttendanceShift = nil
}

// DeleteAttendanceShiftReq ...
type DeleteAttendanceShiftReq struct {
	ShiftID string `path:"shift_id" json:"-"` // 班次 ID，获取方式：1）[按名称查询班次](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/shift/query) 2）[创建班次](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/shift/create), 示例值："6919358778597097404"
}

// deleteAttendanceShiftResp ...
type deleteAttendanceShiftResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *DeleteAttendanceShiftResp `json:"data,omitempty"`
}

// DeleteAttendanceShiftResp ...
type DeleteAttendanceShiftResp struct {
}
