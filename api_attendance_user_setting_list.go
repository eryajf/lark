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

// GetAttendanceUserSettingList 批量查询授权内员工的用户设置信息, 包括人脸照片文件 ID、人脸照片更新时间。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_setting/query
// new doc: https://open.feishu.cn/document/server-docs/attendance-v1/user_setting/query
func (r *AttendanceService) GetAttendanceUserSettingList(ctx context.Context, request *GetAttendanceUserSettingListReq, options ...MethodOptionFunc) (*GetAttendanceUserSettingListResp, *Response, error) {
	if r.cli.mock.mockAttendanceGetAttendanceUserSettingList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#GetAttendanceUserSettingList mock enable")
		return r.cli.mock.mockAttendanceGetAttendanceUserSettingList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Attendance",
		API:                   "GetAttendanceUserSettingList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/attendance/v1/user_settings/query",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getAttendanceUserSettingListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAttendanceGetAttendanceUserSettingList mock AttendanceGetAttendanceUserSettingList method
func (r *Mock) MockAttendanceGetAttendanceUserSettingList(f func(ctx context.Context, request *GetAttendanceUserSettingListReq, options ...MethodOptionFunc) (*GetAttendanceUserSettingListResp, *Response, error)) {
	r.mockAttendanceGetAttendanceUserSettingList = f
}

// UnMockAttendanceGetAttendanceUserSettingList un-mock AttendanceGetAttendanceUserSettingList method
func (r *Mock) UnMockAttendanceGetAttendanceUserSettingList() {
	r.mockAttendanceGetAttendanceUserSettingList = nil
}

// GetAttendanceUserSettingListReq ...
type GetAttendanceUserSettingListReq struct {
	EmployeeType EmployeeType `query:"employee_type" json:"-"` // 请求体中的 user_ids 和响应体中的 user_id 的员工工号类型, 示例值: employee_id, 可选值有: employee_id: 员工 employee ID, 即飞书管理后台 > 组织架构 > 成员与部门 > 成员详情中的用户 ID, employee_no: 员工工号, 即飞书管理后台 > 组织架构 > 成员与部门 > 成员详情中的工号
	UserIDs      []string     `json:"user_ids,omitempty"`      // employee_no 或 employee_id 列表, 示例值: ["abd754f7"], 最大长度: `100`
}

// GetAttendanceUserSettingListResp ...
type GetAttendanceUserSettingListResp struct {
	UserSettings []*GetAttendanceUserSettingListRespUserSetting `json:"user_settings,omitempty"` // 用户设置信息列表
}

// GetAttendanceUserSettingListRespUserSetting ...
type GetAttendanceUserSettingListRespUserSetting struct {
	UserID            string `json:"user_id,omitempty"`              // 用户 ID
	FaceKey           string `json:"face_key,omitempty"`             // 人脸照片文件 ID, 获取方式: [文件上传](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/file/upload)
	FaceKeyUpdateTime string `json:"face_key_update_time,omitempty"` // 人脸照片更新时间, 精确到秒的时间戳
}

// getAttendanceUserSettingListResp ...
type getAttendanceUserSettingListResp struct {
	Code int64                             `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                            `json:"msg,omitempty"`  // 错误描述
	Data *GetAttendanceUserSettingListResp `json:"data,omitempty"`
}
