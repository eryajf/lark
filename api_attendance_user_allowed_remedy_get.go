// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetAttendanceUserAllowedRemedy
//
// 获取用户某天可以补第几次上/下班卡的时间。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/query-user-allowed-remedys
func (r *AttendanceService) GetAttendanceUserAllowedRemedy(ctx context.Context, request *GetAttendanceUserAllowedRemedyReq, options ...MethodOptionFunc) (*GetAttendanceUserAllowedRemedyResp, *Response, error) {
	if r.cli.mock.mockAttendanceGetAttendanceUserAllowedRemedy != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#GetAttendanceUserAllowedRemedy mock enable")
		return r.cli.mock.mockAttendanceGetAttendanceUserAllowedRemedy(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Attendance",
		API:                   "GetAttendanceUserAllowedRemedy",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/attendance/v1/user_task_remedys/query_user_allowed_remedys",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getAttendanceUserAllowedRemedyResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockAttendanceGetAttendanceUserAllowedRemedy(f func(ctx context.Context, request *GetAttendanceUserAllowedRemedyReq, options ...MethodOptionFunc) (*GetAttendanceUserAllowedRemedyResp, *Response, error)) {
	r.mockAttendanceGetAttendanceUserAllowedRemedy = f
}

func (r *Mock) UnMockAttendanceGetAttendanceUserAllowedRemedy() {
	r.mockAttendanceGetAttendanceUserAllowedRemedy = nil
}

type GetAttendanceUserAllowedRemedyReq struct {
	EmployeeType EmployeeType `query:"employee_type" json:"-"` // 请求体中的 user_id 的员工工号类型，必选字段，可用值：【employee_id（员工employeeId），employee_no（员工工号）】，示例值："employee_id"
	UserID       string       `json:"user_id,omitempty"`       // 用户 ID
	RemedyDate   int64        `json:"remedy_date,omitempty"`   // 查询补卡的日期
}

type getAttendanceUserAllowedRemedyResp struct {
	Code int64                               `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                              `json:"msg,omitempty"`  // 错误描述
	Data *GetAttendanceUserAllowedRemedyResp `json:"data,omitempty"` // -
}

type GetAttendanceUserAllowedRemedyResp struct {
	UserAllowedRemedys *GetAttendanceUserAllowedRemedyRespUserAllowedRemedys `json:"user_allowed_remedys,omitempty"`
}

type GetAttendanceUserAllowedRemedyRespUserAllowedRemedys struct {
	UserID          string `json:"user_id,omitempty"`           // 用户 ID
	RemedyDate      int64  `json:"remedy_date,omitempty"`       // 补卡日期
	IsFreePunch     bool   `json:"is_free_punch,omitempty"`     // 是否为自由班次，若为自由班次，则不用选择考虑第几次上下班，直接选择补卡时间即可
	PunchNo         int64  `json:"punch_no,omitempty"`          // 第几次上下班，可用值：【0（第 1 次上下班），1（第 2 次上下班），2（第 3 次上下班）】
	WorkType        int64  `json:"work_type,omitempty"`         // 上班/下班，1：上班，2：下班
	PunchStatus     string `json:"punch_status,omitempty"`      // 打卡状态，可用值【Early（早退），Late（迟到），Lack（缺卡）】
	NormalPunchTime string `json:"normal_punch_time,omitempty"` // 正常的应打卡时间，时间格式为 yyyy-MM-dd HH:mm
	RemedyStartTime string `json:"remedy_start_time,omitempty"` // 可选的补卡时间的最小值，时间格式为 yyyy-MM-dd HH:mm
	RemedyEndTime   string `json:"remedy_end_time,omitempty"`   // 可选的补卡时间的最大值，时间格式为 yyyy-MM-dd HH:mm
}
