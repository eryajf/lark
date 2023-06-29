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

// GetAttendanceUserTaskRemedyAllowedRemedyList 获取用户某天可以补的第几次上 / 下班卡的时间。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_task_remedy/query_user_allowed_remedys
// new doc: https://open.feishu.cn/document/server-docs/attendance-v1/user_task_remedy/query_user_allowed_remedys
func (r *AttendanceService) GetAttendanceUserTaskRemedyAllowedRemedyList(ctx context.Context, request *GetAttendanceUserTaskRemedyAllowedRemedyListReq, options ...MethodOptionFunc) (*GetAttendanceUserTaskRemedyAllowedRemedyListResp, *Response, error) {
	if r.cli.mock.mockAttendanceGetAttendanceUserTaskRemedyAllowedRemedyList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#GetAttendanceUserTaskRemedyAllowedRemedyList mock enable")
		return r.cli.mock.mockAttendanceGetAttendanceUserTaskRemedyAllowedRemedyList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Attendance",
		API:                   "GetAttendanceUserTaskRemedyAllowedRemedyList",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/attendance/v1/user_task_remedys/query_user_allowed_remedys",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getAttendanceUserTaskRemedyAllowedRemedyListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAttendanceGetAttendanceUserTaskRemedyAllowedRemedyList mock AttendanceGetAttendanceUserTaskRemedyAllowedRemedyList method
func (r *Mock) MockAttendanceGetAttendanceUserTaskRemedyAllowedRemedyList(f func(ctx context.Context, request *GetAttendanceUserTaskRemedyAllowedRemedyListReq, options ...MethodOptionFunc) (*GetAttendanceUserTaskRemedyAllowedRemedyListResp, *Response, error)) {
	r.mockAttendanceGetAttendanceUserTaskRemedyAllowedRemedyList = f
}

// UnMockAttendanceGetAttendanceUserTaskRemedyAllowedRemedyList un-mock AttendanceGetAttendanceUserTaskRemedyAllowedRemedyList method
func (r *Mock) UnMockAttendanceGetAttendanceUserTaskRemedyAllowedRemedyList() {
	r.mockAttendanceGetAttendanceUserTaskRemedyAllowedRemedyList = nil
}

// GetAttendanceUserTaskRemedyAllowedRemedyListReq ...
type GetAttendanceUserTaskRemedyAllowedRemedyListReq struct {
	EmployeeType EmployeeType `query:"employee_type" json:"-"` // 请求体和响应体中的 user_id 的员工工号类型, 示例值: employee_id, 可选值有: employee_id: 员工 employee ID, 即飞书管理后台 > 组织架构 > 成员与部门 > 成员详情中的用户 ID, employee_no: 员工工号, 即飞书管理后台 > 组织架构 > 成员与部门 > 成员详情中的工号
	UserID       string       `json:"user_id,omitempty"`       // 用户 ID, 示例值: "abd754f7"
	RemedyDate   int64        `json:"remedy_date,omitempty"`   // 补卡日期, 示例值: 20210104
}

// GetAttendanceUserTaskRemedyAllowedRemedyListResp ...
type GetAttendanceUserTaskRemedyAllowedRemedyListResp struct {
	UserAllowedRemedys []*GetAttendanceUserTaskRemedyAllowedRemedyListRespUserAllowedRemedy `json:"user_allowed_remedys,omitempty"` // 用户可补卡时间
}

// GetAttendanceUserTaskRemedyAllowedRemedyListRespUserAllowedRemedy ...
type GetAttendanceUserTaskRemedyAllowedRemedyListRespUserAllowedRemedy struct {
	UserID          string `json:"user_id,omitempty"`           // 用户 ID
	RemedyDate      int64  `json:"remedy_date,omitempty"`       // 补卡日期
	IsFreePunch     bool   `json:"is_free_punch,omitempty"`     // 是否为自由班次, 若为自由班次, 则不用选择考虑第几次上下班, 直接选择补卡时间即可
	PunchNo         int64  `json:"punch_no,omitempty"`          // 第几次上下班, 0: 第 1 次上下班, 1: 第 2 次上下班, 2: 第 3 次上下班
	WorkType        int64  `json:"work_type,omitempty"`         // 上班 / 下班, 1: 上班, 2: 下班
	PunchStatus     string `json:"punch_status,omitempty"`      // 打卡状态, Early: 早退, Late: 迟到, Lack: 缺卡
	NormalPunchTime string `json:"normal_punch_time,omitempty"` // 正常的应打卡时间, 时间格式为 yyyy-MM-dd HH:mm
	RemedyStartTime string `json:"remedy_start_time,omitempty"` // 可选的补卡时间的最小值, 时间格式为 yyyy-MM-dd HH:mm
	RemedyEndTime   string `json:"remedy_end_time,omitempty"`   // 可选的补卡时间的最大值, 时间格式为 yyyy-MM-dd HH:mm
}

// getAttendanceUserTaskRemedyAllowedRemedyListResp ...
type getAttendanceUserTaskRemedyAllowedRemedyListResp struct {
	Code int64                                             `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                                            `json:"msg,omitempty"`  // 错误描述
	Data *GetAttendanceUserTaskRemedyAllowedRemedyListResp `json:"data,omitempty"`
}
