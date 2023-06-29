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

// GetAttendanceGroup 通过考勤组 ID 获取考勤组详情。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/group/get
// new doc: https://open.feishu.cn/document/server-docs/attendance-v1/group/get
func (r *AttendanceService) GetAttendanceGroup(ctx context.Context, request *GetAttendanceGroupReq, options ...MethodOptionFunc) (*GetAttendanceGroupResp, *Response, error) {
	if r.cli.mock.mockAttendanceGetAttendanceGroup != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#GetAttendanceGroup mock enable")
		return r.cli.mock.mockAttendanceGetAttendanceGroup(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Attendance",
		API:                   "GetAttendanceGroup",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/attendance/v1/groups/:group_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getAttendanceGroupResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAttendanceGetAttendanceGroup mock AttendanceGetAttendanceGroup method
func (r *Mock) MockAttendanceGetAttendanceGroup(f func(ctx context.Context, request *GetAttendanceGroupReq, options ...MethodOptionFunc) (*GetAttendanceGroupResp, *Response, error)) {
	r.mockAttendanceGetAttendanceGroup = f
}

// UnMockAttendanceGetAttendanceGroup un-mock AttendanceGetAttendanceGroup method
func (r *Mock) UnMockAttendanceGetAttendanceGroup() {
	r.mockAttendanceGetAttendanceGroup = nil
}

// GetAttendanceGroupReq ...
type GetAttendanceGroupReq struct {
	GroupID      string       `path:"group_id" json:"-"`       // 考勤组 ID, 获取方式: 1）[创建或修改考勤组](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/group/create) 2）[按名称查询考勤组](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/group/search) 3）[获取打卡结果](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_task/query), 示例值: "6919358128597097404"
	EmployeeType EmployeeType `query:"employee_type" json:"-"` // 用户 ID 的类型, 示例值: employee_id, 可选值有: employee_id: 员工 employeeId, employee_no: 员工工号
	DeptType     string       `query:"dept_type" json:"-"`     // 部门 ID 的类型, 示例值: od-fcb45c28a45311afd441b8869541ece8, 可选值有: open_id: 暂时只支持部门的 openid
}

// GetAttendanceGroupResp ...
type GetAttendanceGroupResp struct {
	GroupID                 string                                         `json:"group_id,omitempty"`                    // 考勤组 ID（仅修改时提供）, 需要从“获取打卡结果”的接口中获取 groupId
	GroupName               string                                         `json:"group_name,omitempty"`                  // 考勤组名称
	TimeZone                string                                         `json:"time_zone,omitempty"`                   // 时区
	BindDeptIDs             []string                                       `json:"bind_dept_ids,omitempty"`               // 绑定的部门 ID
	ExceptDeptIDs           []string                                       `json:"except_dept_ids,omitempty"`             // 排除的部门 ID
	BindUserIDs             []string                                       `json:"bind_user_ids,omitempty"`               // 绑定的用户 ID
	ExceptUserIDs           []string                                       `json:"except_user_ids,omitempty"`             // 排除的用户 ID
	GroupLeaderIDs          []string                                       `json:"group_leader_ids,omitempty"`            // 考勤主负责人 ID 列表, 必选字段（需至少拥有考勤组管理员权限）
	SubGroupLeaderIDs       []string                                       `json:"sub_group_leader_ids,omitempty"`        // 考勤子负责人 ID 列表
	AllowOutPunch           bool                                           `json:"allow_out_punch,omitempty"`             // 是否允许外勤打卡
	OutPunchNeedApproval    bool                                           `json:"out_punch_need_approval,omitempty"`     // 外勤打卡需审批（需要允许外勤打卡才能设置生效）
	OutPunchNeedRemark      bool                                           `json:"out_punch_need_remark,omitempty"`       // 外勤打卡需填写备注（需要允许外勤打卡才能设置生效）
	OutPunchNeedPhoto       bool                                           `json:"out_punch_need_photo,omitempty"`        // 外勤打卡需拍照（需要允许外勤打卡才能设置生效）
	OutPunchAllowedHideAddr bool                                           `json:"out_punch_allowed_hide_addr,omitempty"` // 外勤打卡允许员工隐藏详细地址（需要允许外勤打卡才能设置生效）
	AllowPcPunch            bool                                           `json:"allow_pc_punch,omitempty"`              // 是否允许 PC 端打卡
	AllowRemedy             bool                                           `json:"allow_remedy,omitempty"`                // 是否限制补卡
	RemedyLimit             bool                                           `json:"remedy_limit,omitempty"`                // 是否限制补卡次数
	RemedyLimitCount        int64                                          `json:"remedy_limit_count,omitempty"`          // 补卡次数
	RemedyDateLimit         bool                                           `json:"remedy_date_limit,omitempty"`           // 是否限制补卡时间
	RemedyDateNum           int64                                          `json:"remedy_date_num,omitempty"`             // 补卡时间, 几天内补卡
	AllowRemedyTypeLack     bool                                           `json:"allow_remedy_type_lack,omitempty"`      // 允许缺卡补卡（需要允许补卡才能设置生效）
	AllowRemedyTypeLate     bool                                           `json:"allow_remedy_type_late,omitempty"`      // 允许迟到补卡（需要允许补卡才能设置生效）
	AllowRemedyTypeEarly    bool                                           `json:"allow_remedy_type_early,omitempty"`     // 允许早退补卡（需要允许补卡才能设置生效）
	AllowRemedyTypeNormal   bool                                           `json:"allow_remedy_type_normal,omitempty"`    // 允许正常补卡（需要允许补卡才能设置生效）
	ShowCumulativeTime      bool                                           `json:"show_cumulative_time,omitempty"`        // 是否展示累计时长
	ShowOverTime            bool                                           `json:"show_over_time,omitempty"`              // 是否展示加班时长
	HideStaffPunchTime      bool                                           `json:"hide_staff_punch_time,omitempty"`       // 是否隐藏员工打卡详情
	FacePunch               bool                                           `json:"face_punch,omitempty"`                  // 是否开启人脸识别打卡
	FacePunchCfg            int64                                          `json:"face_punch_cfg,omitempty"`              // 人脸识别打卡规则, 可选值有: * 1: 每次打卡均需人脸识别, * 2: 疑似作弊打卡时需要人脸识别
	FaceLiveNeedAction      bool                                           `json:"face_live_need_action,omitempty"`       // 人脸打卡规则, 可选值有: * false: 开启活体验证, * true: 动作验证, 仅在 face_punch_cfg = 1 时有效
	FaceDowngrade           bool                                           `json:"face_downgrade,omitempty"`              // 人脸识别失败时是否允许普通拍照打卡
	ReplaceBasicPic         bool                                           `json:"replace_basic_pic,omitempty"`           // 人脸识别失败时是否允许替换基准图片
	Machines                []*GetAttendanceGroupRespMachine               `json:"machines,omitempty"`                    // 考勤机列表
	GpsRange                int64                                          `json:"gps_range,omitempty"`                   // GPS 打卡的有效范围（不建议使用）
	Locations               []*GetAttendanceGroupRespLocation              `json:"locations,omitempty"`                   // 地址列表
	GroupType               int64                                          `json:"group_type,omitempty"`                  // 考勤类型, 0: 固定班制, 2: 排班制, 3: 自由班制
	PunchDayShiftIDs        []string                                       `json:"punch_day_shift_ids,omitempty"`         // 固定班制必须填
	FreePunchCfg            *GetAttendanceGroupRespFreePunchCfg            `json:"free_punch_cfg,omitempty"`              // 配置自由班制
	CalendarID              int64                                          `json:"calendar_id,omitempty"`                 // 国家日历  ID, 0: 不根据国家日历排休, 1: 中国大陆, 2: 美国, 3: 日本, 4: 印度, 5: 新加坡, 默认 1
	NeedPunchSpecialDays    []*GetAttendanceGroupRespNeedPunchSpecialDay   `json:"need_punch_special_days,omitempty"`     // 必须打卡的特殊日期
	NoNeedPunchSpecialDays  []*GetAttendanceGroupRespNoNeedPunchSpecialDay `json:"no_need_punch_special_days,omitempty"`  // 无需打卡的特殊日期
	WorkDayNoPunchAsLack    bool                                           `json:"work_day_no_punch_as_lack,omitempty"`   // 自由班制下工作日不打卡是否记为缺卡
	RemedyPeriodType        int64                                          `json:"remedy_period_type,omitempty"`          // 补卡周期类型
	RemedyPeriodCustomDate  int64                                          `json:"remedy_period_custom_date,omitempty"`   // 补卡自定义周期起始日期
	PunchType               int64                                          `json:"punch_type,omitempty"`                  // 打卡类型, 位运算。1: GPS 打卡, 2: Wi-Fi 打卡, 4: 考勤机打卡, 8: IP 打卡
	EffectTime              string                                         `json:"effect_time,omitempty"`                 // 生效时间, 精确到秒的时间戳
	FixshiftEffectTime      string                                         `json:"fixshift_effect_time,omitempty"`        // 固定班次生效时间, 精确到秒的时间戳
	MemberEffectTime        string                                         `json:"member_effect_time,omitempty"`          // 参加考勤的人员、部门变动生效时间, 精确到秒的时间戳
	RestClockInNeedApproval bool                                           `json:"rest_clockIn_need_approval,omitempty"`  // 休息日打卡需审批
	ClockInNeedPhoto        bool                                           `json:"clockIn_need_photo,omitempty"`          // 每次打卡均需拍照
	MemberStatusChange      *GetAttendanceGroupRespMemberStatusChange      `json:"member_status_change,omitempty"`        // 人员异动打卡设置
	LeaveNeedPunch          bool                                           `json:"leave_need_punch,omitempty"`            // 请假离岗或返岗是否需打卡
	LeaveNeedPunchCfg       *GetAttendanceGroupRespLeaveNeedPunchCfg       `json:"leave_need_punch_cfg,omitempty"`        // 请假离岗或返岗打卡规则
	GoOutNeedPunch          int64                                          `json:"go_out_need_punch,omitempty"`           // 外出期间是否需打卡
	GoOutNeedPunchCfg       *GetAttendanceGroupRespGoOutNeedPunchCfg       `json:"go_out_need_punch_cfg,omitempty"`       // 外出期间打卡规则
	TravelNeedPunch         int64                                          `json:"travel_need_punch,omitempty"`           // 出差期间是否需打卡
	TravelNeedPunchCfg      *GetAttendanceGroupRespTravelNeedPunchCfg      `json:"travel_need_punch_cfg,omitempty"`       // 出差期间打卡规则
	NeedPunchMembers        []*GetAttendanceGroupRespNeedPunchMember       `json:"need_punch_members,omitempty"`          // 需要打卡的人员配置（新）
	NoNeedPunchMembers      []*GetAttendanceGroupRespNoNeedPunchMember     `json:"no_need_punch_members,omitempty"`       // 无需打卡的人员配置（新）
}

// GetAttendanceGroupRespFreePunchCfg ...
type GetAttendanceGroupRespFreePunchCfg struct {
	FreeStartTime        string `json:"free_start_time,omitempty"`           // 自由班制打卡开始时间
	FreeEndTime          string `json:"free_end_time,omitempty"`             // 自由班制打卡结束时间
	PunchDay             int64  `json:"punch_day,omitempty"`                 // 打卡的时间, 为 7 位数字, 每一位依次代表周一到周日, 0 为不上班, 1 为上班
	WorkDayNoPunchAsLack bool   `json:"work_day_no_punch_as_lack,omitempty"` // 工作日不打卡是否记为缺卡
	WorkHoursDemand      bool   `json:"work_hours_demand,omitempty"`         // 工作日出勤是否需满足时长要求
	WorkHours            int64  `json:"work_hours,omitempty"`                // 每日工作时长（分钟), 范围[0, 1440]
}

// GetAttendanceGroupRespGoOutNeedPunchCfg ...
type GetAttendanceGroupRespGoOutNeedPunchCfg struct {
	LateMinutesAsLate   int64 `json:"late_minutes_as_late,omitempty"`   // 晚到超过多久记为迟到
	LateMinutesAsLack   int64 `json:"late_minutes_as_lack,omitempty"`   // 晚到超过多久记为缺卡
	EarlyMinutesAsEarly int64 `json:"early_minutes_as_early,omitempty"` // 早走超过多久记为早退
	EarlyMinutesAsLack  int64 `json:"early_minutes_as_lack,omitempty"`  // 早走超过多久记为缺卡
}

// GetAttendanceGroupRespLeaveNeedPunchCfg ...
type GetAttendanceGroupRespLeaveNeedPunchCfg struct {
	LateMinutesAsLate   int64 `json:"late_minutes_as_late,omitempty"`   // 晚到超过多久记为迟到
	LateMinutesAsLack   int64 `json:"late_minutes_as_lack,omitempty"`   // 晚到超过多久记为缺卡
	EarlyMinutesAsEarly int64 `json:"early_minutes_as_early,omitempty"` // 早走超过多久记为早退
	EarlyMinutesAsLack  int64 `json:"early_minutes_as_lack,omitempty"`  // 早走超过多久记为缺卡
}

// GetAttendanceGroupRespLocation ...
type GetAttendanceGroupRespLocation struct {
	LocationID   string  `json:"location_id,omitempty"`   // 地址 ID
	LocationName string  `json:"location_name,omitempty"` // 地址名称
	LocationType int64   `json:"location_type,omitempty"` // 地址类型, 可选值有: * 1: GPS, * 2: Wi-Fi, * 8: IP
	Latitude     float64 `json:"latitude,omitempty"`      // 地址纬度
	Longitude    float64 `json:"longitude,omitempty"`     // 地址经度
	Ssid         string  `json:"ssid,omitempty"`          // Wi-Fi 名称
	Bssid        string  `json:"bssid,omitempty"`         // Wi-Fi 的 MAC 地址
	MapType      int64   `json:"map_type,omitempty"`      // 地图类型, 1: 高德, 2: 谷歌
	Address      string  `json:"address,omitempty"`       // 地址名称
	Ip           string  `json:"ip,omitempty"`            // IP 地址
	Feature      string  `json:"feature,omitempty"`       // 额外信息, 例如: 运营商信息
	GpsRange     int64   `json:"gps_range,omitempty"`     // GPS 打卡的有效范围
}

// GetAttendanceGroupRespMachine ...
type GetAttendanceGroupRespMachine struct {
	MachineSn   string `json:"machine_sn,omitempty"`   // 考勤机序列号
	MachineName string `json:"machine_name,omitempty"` // 考勤机名称
}

// GetAttendanceGroupRespMemberStatusChange ...
type GetAttendanceGroupRespMemberStatusChange struct {
	OnboardingOnNoNeedPunch   bool `json:"onboarding_on_no_need_punch,omitempty"`   // 是否入职日上班无需打卡
	OnboardingOffNoNeedPunch  bool `json:"onboarding_off_no_need_punch,omitempty"`  // 是否入职日下班无需打卡
	OffboardingOnNoNeedPunch  bool `json:"offboarding_on_no_need_punch,omitempty"`  // 是否离职日上班无需打卡
	OffboardingOffNoNeedPunch bool `json:"offboarding_off_no_need_punch,omitempty"` // 是否离职日下班无需打卡
}

// GetAttendanceGroupRespNeedPunchMember ...
type GetAttendanceGroupRespNeedPunchMember struct {
	RuleScopeType  int64                                                `json:"rule_scope_type,omitempty"`  // 圈人方式, 可选值有: * 0 无, * 1 全部, * 2 自定义
	ScopeGroupList *GetAttendanceGroupRespNeedPunchMemberScopeGroupList `json:"scope_group_list,omitempty"` // 圈人规则列表
}

// GetAttendanceGroupRespNeedPunchMemberScopeGroupList ...
type GetAttendanceGroupRespNeedPunchMemberScopeGroupList struct {
	ScopeValueType int64                                                       `json:"scope_value_type,omitempty"` // 类型, 可选值有: * 1: 部门, * 2:人员, * 3:国家地区, * 4:员工类型, * 5:性别, * 6:工作城市
	OperationType  int64                                                       `json:"operation_type,omitempty"`   // 范围类型（是否包含）
	Right          []*GetAttendanceGroupRespNeedPunchMemberScopeGroupListRight `json:"right,omitempty"`            // 如果是人员/部门类型 不需要使用该字段
	MemberIDs      []string                                                    `json:"member_ids,omitempty"`       // 部门/人员id列表
}

// GetAttendanceGroupRespNeedPunchMemberScopeGroupListRight ...
type GetAttendanceGroupRespNeedPunchMemberScopeGroupListRight struct {
	Key  string `json:"key,omitempty"`  // 标识Key
	Name string `json:"name,omitempty"` // 名称
}

// GetAttendanceGroupRespNeedPunchSpecialDay ...
type GetAttendanceGroupRespNeedPunchSpecialDay struct {
	PunchDay int64  `json:"punch_day,omitempty"` // 打卡日期
	ShiftID  string `json:"shift_id,omitempty"`  // 班次 ID
}

// GetAttendanceGroupRespNoNeedPunchMember ...
type GetAttendanceGroupRespNoNeedPunchMember struct {
	RuleScopeType  int64                                                  `json:"rule_scope_type,omitempty"`  // 圈人方式, 可选值有: * 0 无, * 1 全部, * 2 自定义
	ScopeGroupList *GetAttendanceGroupRespNoNeedPunchMemberScopeGroupList `json:"scope_group_list,omitempty"` // 圈人规则列表
}

// GetAttendanceGroupRespNoNeedPunchMemberScopeGroupList ...
type GetAttendanceGroupRespNoNeedPunchMemberScopeGroupList struct {
	ScopeValueType int64                                                         `json:"scope_value_type,omitempty"` // 类型, 可选值有: * 1: 部门, * 2:人员, * 3:国家地区, * 4:员工类型, * 5:性别, * 6:工作城市
	OperationType  int64                                                         `json:"operation_type,omitempty"`   // 范围类型（是否包含）
	Right          []*GetAttendanceGroupRespNoNeedPunchMemberScopeGroupListRight `json:"right,omitempty"`            // 如果是人员/部门类型 不需要使用该字段
	MemberIDs      []string                                                      `json:"member_ids,omitempty"`       // 部门/人员id列表
}

// GetAttendanceGroupRespNoNeedPunchMemberScopeGroupListRight ...
type GetAttendanceGroupRespNoNeedPunchMemberScopeGroupListRight struct {
	Key  string `json:"key,omitempty"`  // 标识Key
	Name string `json:"name,omitempty"` // 名称
}

// GetAttendanceGroupRespNoNeedPunchSpecialDay ...
type GetAttendanceGroupRespNoNeedPunchSpecialDay struct {
	PunchDay int64  `json:"punch_day,omitempty"` // 打卡日期
	ShiftID  string `json:"shift_id,omitempty"`  // 班次 ID
}

// GetAttendanceGroupRespTravelNeedPunchCfg ...
type GetAttendanceGroupRespTravelNeedPunchCfg struct {
	LateMinutesAsLate   int64 `json:"late_minutes_as_late,omitempty"`   // 晚到超过多久记为迟到
	LateMinutesAsLack   int64 `json:"late_minutes_as_lack,omitempty"`   // 晚到超过多久记为缺卡
	EarlyMinutesAsEarly int64 `json:"early_minutes_as_early,omitempty"` // 早走超过多久记为早退
	EarlyMinutesAsLack  int64 `json:"early_minutes_as_lack,omitempty"`  // 早走超过多久记为缺卡
}

// getAttendanceGroupResp ...
type getAttendanceGroupResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *GetAttendanceGroupResp `json:"data,omitempty"`
}
