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

// GetHelpdeskAgentScheduleList 该接口用于获取所有客服信息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/agent_schedule/list
func (r *HelpdeskService) GetHelpdeskAgentScheduleList(ctx context.Context, request *GetHelpdeskAgentScheduleListReq, options ...MethodOptionFunc) (*GetHelpdeskAgentScheduleListResp, *Response, error) {
	if r.cli.mock.mockHelpdeskGetHelpdeskAgentScheduleList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#GetHelpdeskAgentScheduleList mock enable")
		return r.cli.mock.mockHelpdeskGetHelpdeskAgentScheduleList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Helpdesk",
		API:                   "GetHelpdeskAgentScheduleList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/helpdesk/v1/agent_schedules",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedHelpdeskAuth:      true,
	}
	resp := new(getHelpdeskAgentScheduleListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHelpdeskGetHelpdeskAgentScheduleList mock HelpdeskGetHelpdeskAgentScheduleList method
func (r *Mock) MockHelpdeskGetHelpdeskAgentScheduleList(f func(ctx context.Context, request *GetHelpdeskAgentScheduleListReq, options ...MethodOptionFunc) (*GetHelpdeskAgentScheduleListResp, *Response, error)) {
	r.mockHelpdeskGetHelpdeskAgentScheduleList = f
}

// UnMockHelpdeskGetHelpdeskAgentScheduleList un-mock HelpdeskGetHelpdeskAgentScheduleList method
func (r *Mock) UnMockHelpdeskGetHelpdeskAgentScheduleList() {
	r.mockHelpdeskGetHelpdeskAgentScheduleList = nil
}

// GetHelpdeskAgentScheduleListReq ...
type GetHelpdeskAgentScheduleListReq struct {
	Status []int64 `query:"status" json:"-"` // 筛选条件, 1 - online客服, 2 - offline(手动)客服, 3 - off duty(下班)客服, 4 - 移除客服, 示例值：status=1&status=2
}

// getHelpdeskAgentScheduleListResp ...
type getHelpdeskAgentScheduleListResp struct {
	Code int64                             `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                            `json:"msg,omitempty"`  // 错误描述
	Data *GetHelpdeskAgentScheduleListResp `json:"data,omitempty"`
}

// GetHelpdeskAgentScheduleListResp ...
type GetHelpdeskAgentScheduleListResp struct {
	AgentSchedules []*GetHelpdeskAgentScheduleListRespAgentSchedule `json:"agent_schedules,omitempty"` // 客服列表
}

// GetHelpdeskAgentScheduleListRespAgentSchedule ...
type GetHelpdeskAgentScheduleListRespAgentSchedule struct {
	Status      int64                                                      `json:"status,omitempty"`       // 客服状态, 1 - online客服, 2 - offline(手动)客服, 3 - off duty(下班)自动处于非服务时间段
	Agent       *GetHelpdeskAgentScheduleListRespAgentScheduleAgent        `json:"agent,omitempty"`        // 客服信息
	Schedule    []*GetHelpdeskAgentScheduleListRespAgentScheduleSchedule   `json:"schedule,omitempty"`     // 工作日程列表
	AgentSkills []*GetHelpdeskAgentScheduleListRespAgentScheduleAgentSkill `json:"agent_skills,omitempty"` // 客服技能
}

// GetHelpdeskAgentScheduleListRespAgentScheduleAgent ...
type GetHelpdeskAgentScheduleListRespAgentScheduleAgent struct {
	ID          string `json:"id,omitempty"`           // 客服 id
	AvatarURL   string `json:"avatar_url,omitempty"`   // avatar url
	Name        string `json:"name,omitempty"`         // 客服名字
	Email       string `json:"email,omitempty"`        // email
	Department  string `json:"department,omitempty"`   // 部门
	CompanyName string `json:"company_name,omitempty"` // 公司名
}

// GetHelpdeskAgentScheduleListRespAgentScheduleSchedule ...
type GetHelpdeskAgentScheduleListRespAgentScheduleSchedule struct {
	StartTime string `json:"start_time,omitempty"` // 开始时间, format 00:00 - 23:59
	EndTime   string `json:"end_time,omitempty"`   // 结束时间, format 00:00 - 23:59
	Weekday   int64  `json:"weekday,omitempty"`    // 星期几, 1 - Monday, 2 - Tuesday, 3 - Wednesday, 4 - Thursday, 5 - Friday, 6 - Saturday, 7 - Sunday, 9 - Everday, 10 - Weekday, 11 - Weekend
}

// GetHelpdeskAgentScheduleListRespAgentScheduleAgentSkill ...
type GetHelpdeskAgentScheduleListRespAgentScheduleAgentSkill struct {
	ID        string `json:"id,omitempty"`         // 客服技能 id
	Name      string `json:"name,omitempty"`       // 客服技能名
	IsDefault bool   `json:"is_default,omitempty"` // 是默认技能
}
