// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetAdminDeptStats 该接口用于获取部门维度的用户活跃和功能使用数据，即IM（即时通讯）、日历、云文档、音视频会议功能的使用数据。
//
// - 只有企业自建应用才有权限调用此接口
// - 当天的数据会在第二天的凌晨五点产出（UTC+8）
// - 部门维度的数据最多查询最近366天（包含366天）的数据
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/admin-v1/admin_dept_stat/list
func (r *AdminService) GetAdminDeptStats(ctx context.Context, request *GetAdminDeptStatsReq, options ...MethodOptionFunc) (*GetAdminDeptStatsResp, *Response, error) {
	if r.cli.mock.mockAdminGetAdminDeptStats != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Admin#GetAdminDeptStats mock enable")
		return r.cli.mock.mockAdminGetAdminDeptStats(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Admin",
		API:                   "GetAdminDeptStats",
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/admin/v1/admin_dept_stats",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getAdminDeptStatsResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockAdminGetAdminDeptStats(f func(ctx context.Context, request *GetAdminDeptStatsReq, options ...MethodOptionFunc) (*GetAdminDeptStatsResp, *Response, error)) {
	r.mockAdminGetAdminDeptStats = f
}

func (r *Mock) UnMockAdminGetAdminDeptStats() {
	r.mockAdminGetAdminDeptStats = nil
}

type GetAdminDeptStatsReq struct {
	DepartmentIDType  DepartmentIDType `query:"department_id_type" json:"-"`  // 部门ID类型, 示例值："open_department_id", 可选值有: `department_id`：部门的 ID, `open_department_id`：部门的 Open ID
	StartDate         string           `query:"start_date" json:"-"`          // 起始日期（包含），格式是YYYY-mm-dd, 示例值："2020-02-15"
	EndDate           string           `query:"end_date" json:"-"`            // 终止日期（包含），格式是YYYY-mm-dd，起止日期之间相差不能超过91天（包含91天）, 示例值："2020-02-15"
	DepartmentID      string           `query:"department_id" json:"-"`       // 部门的 ID，取决于department_id_type，仅支持根部门及其下前4级子部门, 示例值："od-382e2793cfc9471f892e8a672987654c"
	ContainsChildDept bool             `query:"contains_child_dept" json:"-"` // 是否包含子部门，如果该值为false，则只查出本部门直属用户活跃和功能使用数据；如果该值为true，则查出该部门以及其子部门（子部门层级最多不超过根部门下的前4级）的用户活跃和功能使用数据, 示例值：false
	PageSize          *int64           `query:"page_size" json:"-"`           // 分页大小，默认是10, 示例值：10, 取值范围：`1` ～ `20`
	PageToken         *string          `query:"page_token" json:"-"`          // 分页标记，第一次请求不填，表示从头开始遍历；当返回的has_more为true时，会返回新的page_token，再次调用接口，传入这个page_token，将获得下一页数据, 示例值："2"
}

type getAdminDeptStatsResp struct {
	Code int64                  `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *GetAdminDeptStatsResp `json:"data,omitempty"`
}

type GetAdminDeptStatsResp struct {
	HasMore   bool                         `json:"has_more,omitempty"`   // 分页查询时返回，代表是否还有更多数据
	PageToken string                       `json:"page_token,omitempty"` // 分页标记，下一页分页的token
	Items     []*GetAdminDeptStatsRespItem `json:"items,omitempty"`      // 数据报表
}

type GetAdminDeptStatsRespItem struct {
	Date                 string `json:"date,omitempty"`                    // 日期
	DepartmentID         string `json:"department_id,omitempty"`           // 部门的department_id 或者open_department_id
	DepartmentName       string `json:"department_name,omitempty"`         // 部门名字
	DepartmentPath       string `json:"department_path,omitempty"`         // 部门路径
	TotalUserNum         int64  `json:"total_user_num,omitempty"`          // 部门总人数
	ActiveUserNum        int64  `json:"active_user_num,omitempty"`         // 激活人数
	ActiveUserRate       string `json:"active_user_rate,omitempty"`        // 激活率
	SuiteDau             int64  `json:"suite_dau,omitempty"`               // 活跃人数
	SuiteActiveRate      string `json:"suite_active_rate,omitempty"`       // 活跃率
	NewUserNum           int64  `json:"new_user_num,omitempty"`            // 新用户数
	NewActiveNum         int64  `json:"new_active_num,omitempty"`          // 新激活数
	ResignUserNum        int64  `json:"resign_user_num,omitempty"`         // 离职人数
	IMDau                int64  `json:"im_dau,omitempty"`                  // 消息活跃人数
	SendMessengerUserNum int64  `json:"send_messenger_user_num,omitempty"` // 发送消息人数
	SendMessengerNum     int64  `json:"send_messenger_num,omitempty"`      // 发送消息数
	AvgSendMessengerNum  string `json:"avg_send_messenger_num,omitempty"`  // 人均发送消息数
	DocsDau              int64  `json:"docs_dau,omitempty"`                // 云文档活跃人数
	CreateDocsUserNum    int64  `json:"create_docs_user_num,omitempty"`    // 创建文件人数
	CreateDocsNum        int64  `json:"create_docs_num,omitempty"`         // 创建文件数
	AvgCreateDocsNum     string `json:"avg_create_docs_num,omitempty"`     // 人均创建文件数
	CalDau               int64  `json:"cal_dau,omitempty"`                 // 日历活跃人数
	CreateCalUserNum     int64  `json:"create_cal_user_num,omitempty"`     // 创建日程人数
	CreateCalNum         int64  `json:"create_cal_num,omitempty"`          // 创建日程数
	AvgCreateCalNum      string `json:"avg_create_cal_num,omitempty"`      // 人均创建日程数
	VCDau                int64  `json:"vc_dau,omitempty"`                  // 音视频会议活跃人数
	VCDuration           int64  `json:"vc_duration,omitempty"`             // 会议时长（分钟）
	AvgVCDuration        string `json:"avg_vc_duration,omitempty"`         // 人均会议时长（分钟）
}
