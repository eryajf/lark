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

// BatchGetCalendarMeetingRoomSummary 通过日程的Uid和Original time, 查询会议室日程主题。
//
// doc: https://open.feishu.cn/document/server-docs/calendar-v4/meeting-room-event/
func (r *CalendarService) BatchGetCalendarMeetingRoomSummary(ctx context.Context, request *BatchGetCalendarMeetingRoomSummaryReq, options ...MethodOptionFunc) (*BatchGetCalendarMeetingRoomSummaryResp, *Response, error) {
	if r.cli.mock.mockCalendarBatchGetCalendarMeetingRoomSummary != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Calendar#BatchGetCalendarMeetingRoomSummary mock enable")
		return r.cli.mock.mockCalendarBatchGetCalendarMeetingRoomSummary(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Calendar",
		API:                   "BatchGetCalendarMeetingRoomSummary",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/meeting_room/summary/batch_get",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(batchGetCalendarMeetingRoomSummaryResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCalendarBatchGetCalendarMeetingRoomSummary mock CalendarBatchGetCalendarMeetingRoomSummary method
func (r *Mock) MockCalendarBatchGetCalendarMeetingRoomSummary(f func(ctx context.Context, request *BatchGetCalendarMeetingRoomSummaryReq, options ...MethodOptionFunc) (*BatchGetCalendarMeetingRoomSummaryResp, *Response, error)) {
	r.mockCalendarBatchGetCalendarMeetingRoomSummary = f
}

// UnMockCalendarBatchGetCalendarMeetingRoomSummary un-mock CalendarBatchGetCalendarMeetingRoomSummary method
func (r *Mock) UnMockCalendarBatchGetCalendarMeetingRoomSummary() {
	r.mockCalendarBatchGetCalendarMeetingRoomSummary = nil
}

// BatchGetCalendarMeetingRoomSummaryReq ...
type BatchGetCalendarMeetingRoomSummaryReq struct {
	EventUids []*BatchGetCalendarMeetingRoomSummaryReqEventUid `json:"EventUids,omitempty"` // 需要查询的日程Uid和Original time
}

// BatchGetCalendarMeetingRoomSummaryReqEventUid ...
type BatchGetCalendarMeetingRoomSummaryReqEventUid struct {
	Uid          string `json:"uid,omitempty"`           // 日程的唯一id
	OriginalTime int64  `json:"original_time,omitempty"` // 日程实例原始时间。非重复性日程和重复性日程, 此处传0；重复性日程的例外, 传对应的original_time
}

// BatchGetCalendarMeetingRoomSummaryResp ...
type BatchGetCalendarMeetingRoomSummaryResp struct {
	EventInfos     []*BatchGetCalendarMeetingRoomSummaryRespEventInfo     `json:"EventInfos,omitempty"`     // 成功查询到的日程信息
	ErrorEventUids []*BatchGetCalendarMeetingRoomSummaryRespErrorEventUid `json:"ErrorEventUids,omitempty"` // 没有查询到的日程
}

// BatchGetCalendarMeetingRoomSummaryRespErrorEventUid ...
type BatchGetCalendarMeetingRoomSummaryRespErrorEventUid struct {
	Uid          string `json:"uid,omitempty"`           // 日程的唯一id
	OriginalTime int64  `json:"original_time,omitempty"` // 日程实例原始时间。非重复性日程和重复性日程, 此处为0；重复性日程的例外, 为对应的original_time
	ErrorMsg     string `json:"error_msg,omitempty"`     // 错误信息
}

// BatchGetCalendarMeetingRoomSummaryRespEventInfo ...
type BatchGetCalendarMeetingRoomSummaryRespEventInfo struct {
	Uid          string                                                `json:"uid,omitempty"`           // 日程的唯一id
	OriginalTime int64                                                 `json:"original_time,omitempty"` // 日程实例原始时间。非重复性日程和重复性日程, 此处为0；重复性日程的例外, 为对应的original_time
	Summary      string                                                `json:"summary,omitempty"`       // 日程主题
	Vchat        *BatchGetCalendarMeetingRoomSummaryRespEventInfoVchat `json:"vchat,omitempty"`         // 视频会议信息
}

// BatchGetCalendarMeetingRoomSummaryRespEventInfoVchat ...
type BatchGetCalendarMeetingRoomSummaryRespEventInfoVchat struct {
	VCType      string `json:"vc_type,omitempty"`     // 视屏会议类型   可选值有: - `vc`: 飞书视频会议, 取该类型时, 其他字段无效。  - `third_party`: 第三方链接视频会议, 取该类型时, icon_type、description、meeting_url字段生效。  - `no_meeting`: 无视频会议, 取该类型时, 其他字段无效。 - `lark_live`: 飞书直播, 内部类型, 飞书客户端使用, API不支持创建, 只读。 - `unknown`: 未知类型, 做兼容使用, 飞书客户端使用, API不支持创建, 只读。
	IconType    string `json:"icon_type,omitempty"`   // 第三方视频会议icon类型；可以为空, 为空展示默认icon。  可选值有: - `vc`: 飞书视频会议icon - `live`: 直播视频会议icon  -  `default`: 默认icon
	Description string `json:"description,omitempty"` // 第三方视频会议文案, 可以为空, 为空展示默认文案
	MeetingURL  string `json:"meeting_url,omitempty"` // 视频会议URL
}

// batchGetCalendarMeetingRoomSummaryResp ...
type batchGetCalendarMeetingRoomSummaryResp struct {
	Code int64                                   `json:"code,omitempty"` // 返回码, 非 0 表示失败
	Msg  string                                  `json:"msg,omitempty"`  // 返回码的描述, "success" 表示成功, 其他为错误提示信息
	Data *BatchGetCalendarMeetingRoomSummaryResp `json:"data,omitempty"` // 返回业务信息
}
