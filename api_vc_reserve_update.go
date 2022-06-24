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

// UpdateVCReserve 更新一个预约
//
// 只能更新归属于自己的预约，不需要更新的字段不传（如果传空则会被更新为空）；可用于续期操作，到期时间距离当前时间不超过30天
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/reserve/update
func (r *VCService) UpdateVCReserve(ctx context.Context, request *UpdateVCReserveReq, options ...MethodOptionFunc) (*UpdateVCReserveResp, *Response, error) {
	if r.cli.mock.mockVCUpdateVCReserve != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] VC#UpdateVCReserve mock enable")
		return r.cli.mock.mockVCUpdateVCReserve(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "VC",
		API:                 "UpdateVCReserve",
		Method:              "PUT",
		URL:                 r.cli.openBaseURL + "/open-apis/vc/v1/reserves/:reserve_id",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(updateVCReserveResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockVCUpdateVCReserve mock VCUpdateVCReserve method
func (r *Mock) MockVCUpdateVCReserve(f func(ctx context.Context, request *UpdateVCReserveReq, options ...MethodOptionFunc) (*UpdateVCReserveResp, *Response, error)) {
	r.mockVCUpdateVCReserve = f
}

// UnMockVCUpdateVCReserve un-mock VCUpdateVCReserve method
func (r *Mock) UnMockVCUpdateVCReserve() {
	r.mockVCUpdateVCReserve = nil
}

// UpdateVCReserveReq ...
type UpdateVCReserveReq struct {
	UserIDType      *IDType                            `query:"user_id_type" json:"-"`     // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	ReserveID       string                             `path:"reserve_id" json:"-"`        // 预约ID（预约的唯一标识）, 示例值："6911188411932033028"
	EndTime         *string                            `json:"end_time,omitempty"`         // 预约到期时间（unix时间，单位sec）, 示例值："1608888867"
	MeetingSettings *UpdateVCReserveReqMeetingSettings `json:"meeting_settings,omitempty"` // 会议设置
}

// UpdateVCReserveReqMeetingSettings ...
type UpdateVCReserveReqMeetingSettings struct {
	Topic              *string                                              `json:"topic,omitempty"`                // 会议主题, 示例值："my meeting"
	ActionPermissions  []*UpdateVCReserveReqMeetingSettingsActionPermission `json:"action_permissions,omitempty"`   // 会议权限配置列表，如果存在相同的权限配置项则它们之间为"逻辑或"的关系（即 有一个为true则拥有该权限）
	MeetingInitialType *int64                                               `json:"meeting_initial_type,omitempty"` // 会议初始类型, 示例值：1, 可选值有: `1`：多人会议, `2`：1v1呼叫
	CallSetting        *UpdateVCReserveReqMeetingSettingsCallSetting        `json:"call_setting,omitempty"`         // 1v1呼叫相关参数
}

// UpdateVCReserveReqMeetingSettingsActionPermission ...
type UpdateVCReserveReqMeetingSettingsActionPermission struct {
	Permission         int64                                                                 `json:"permission,omitempty"`          // 权限项, 示例值：1, 可选值有: `1`：是否能成为主持人, `2`：是否能邀请参会人, `3`：是否能加入会议
	PermissionCheckers []*UpdateVCReserveReqMeetingSettingsActionPermissionPermissionChecker `json:"permission_checkers,omitempty"` // 权限检查器列表，权限检查器之间为"逻辑或"的关系（即 有一个为true则拥有该权限）
}

// UpdateVCReserveReqMeetingSettingsActionPermissionPermissionChecker ...
type UpdateVCReserveReqMeetingSettingsActionPermissionPermissionChecker struct {
	CheckField int64    `json:"check_field,omitempty"` // 检查字段类型, 示例值：1, 可选值有: `1`：用户ID, `2`：用户类型, `3`：租户ID
	CheckMode  int64    `json:"check_mode,omitempty"`  // 检查方式, 示例值：1, 可选值有: `1`：在check_list中为有权限（白名单）, `2`：不在check_list中为有权限（黑名单）
	CheckList  []string `json:"check_list,omitempty"`  // 检查字段列表
}

// UpdateVCReserveReqMeetingSettingsCallSetting ...
type UpdateVCReserveReqMeetingSettingsCallSetting struct {
	Callee *UpdateVCReserveReqMeetingSettingsCallSettingCallee `json:"callee,omitempty"` // 被呼叫的用户
}

// UpdateVCReserveReqMeetingSettingsCallSettingCallee ...
type UpdateVCReserveReqMeetingSettingsCallSettingCallee struct {
	ID          *string                                                        `json:"id,omitempty"`            // 用户ID, 示例值："ou_3ec3f6a28a0d08c45d895276e8e5e19b"
	UserType    int64                                                          `json:"user_type,omitempty"`     // 用户类型，当前仅支持用户类型6(pstn用户), 示例值：1, 可选值有: `1`：lark用户, `2`：rooms用户, `3`：文档用户, `4`：neo单品用户, `5`：neo单品游客用户, `6`：pstn用户, `7`：sip用户
	PstnSipInfo *UpdateVCReserveReqMeetingSettingsCallSettingCalleePstnSipInfo `json:"pstn_sip_info,omitempty"` // pstn/sip信息
}

// UpdateVCReserveReqMeetingSettingsCallSettingCalleePstnSipInfo ...
type UpdateVCReserveReqMeetingSettingsCallSettingCalleePstnSipInfo struct {
	Nickname    *string `json:"nickname,omitempty"`     // 给pstn/sip用户设置的临时昵称, 示例值："dodo"
	MainAddress string  `json:"main_address,omitempty"` // pstn/sip主机号，格式为：[国际冠字]-[电话区号][电话号码]，当前仅支持国内手机及固定电话号码, 示例值："+86-02187654321"
}

// updateVCReserveResp ...
type updateVCReserveResp struct {
	Code int64                `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 错误描述
	Data *UpdateVCReserveResp `json:"data,omitempty"`
}

// UpdateVCReserveResp ...
type UpdateVCReserveResp struct {
	Reserve *UpdateVCReserveRespReserve `json:"reserve,omitempty"` // 预约数据
}

// UpdateVCReserveRespReserve ...
type UpdateVCReserveRespReserve struct {
	ID           string `json:"id,omitempty"`            // 预约ID（预约的唯一标识）
	MeetingNo    string `json:"meeting_no,omitempty"`    // 9位会议号（飞书用户可通过输入9位会议号快捷入会）
	URL          string `json:"url,omitempty"`           // 会议链接（飞书用户可通过点击会议链接快捷入会）
	EndTime      string `json:"end_time,omitempty"`      // 预约到期时间（unix时间，单位sec）
	ExpireStatus int64  `json:"expire_status,omitempty"` // 过期状态, 可选值有: `1`：未过期, `2`：已过期
}
