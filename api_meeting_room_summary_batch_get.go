// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// BatchGetMeetingRoomSummary 通过日程的Uid和Original time，查询会议室日程主题。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uIjM5UjLyITO14iMykTN/
func (r *MeetingRoomService) BatchGetMeetingRoomSummary(ctx context.Context, request *BatchGetMeetingRoomSummaryReq, options ...MethodOptionFunc) (*BatchGetMeetingRoomSummaryResp, *Response, error) {
	if r.cli.mock.mockMeetingRoomBatchGetMeetingRoomSummary != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] MeetingRoom#BatchGetMeetingRoomSummary mock enable")
		return r.cli.mock.mockMeetingRoomBatchGetMeetingRoomSummary(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "MeetingRoom",
		API:                   "BatchGetMeetingRoomSummary",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/meeting_room/summary/batch_get",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(batchGetMeetingRoomSummaryResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockMeetingRoomBatchGetMeetingRoomSummary(f func(ctx context.Context, request *BatchGetMeetingRoomSummaryReq, options ...MethodOptionFunc) (*BatchGetMeetingRoomSummaryResp, *Response, error)) {
	r.mockMeetingRoomBatchGetMeetingRoomSummary = f
}

func (r *Mock) UnMockMeetingRoomBatchGetMeetingRoomSummary() {
	r.mockMeetingRoomBatchGetMeetingRoomSummary = nil
}

type BatchGetMeetingRoomSummaryReq struct {
	EventUids *BatchGetMeetingRoomSummaryReqEventUids `json:"EventUids,omitempty"` // 需要查询的日程Uid和Original time
}

type BatchGetMeetingRoomSummaryReqEventUids struct {
	Uid          string `json:"uid,omitempty"`           // 日程的唯一id
	OriginalTime int64  `json:"original_time,omitempty"` // 日程实例原始时间，非重复日程必为0。重复日程若为0则表示回复其所有实例，否则表示回复单个实例。
}

type batchGetMeetingRoomSummaryResp struct {
	Code int64                           `json:"code,omitempty"` // 返回码，非 0 表示失败
	Msg  string                          `json:"msg,omitempty"`  // 返回码的描述，"success" 表示成功，其他为错误提示信息
	Data *BatchGetMeetingRoomSummaryResp `json:"data,omitempty"` // 返回业务信息
}

type BatchGetMeetingRoomSummaryResp struct {
	EventInfos     *BatchGetMeetingRoomSummaryRespEventInfos     `json:"EventInfos,omitempty"`     // 成功查询到的日程信息
	ErrorEventUids *BatchGetMeetingRoomSummaryRespErrorEventUids `json:"ErrorEventUids,omitempty"` // 没有查询到的日程
}

type BatchGetMeetingRoomSummaryRespEventInfos struct {
	Uid          string `json:"uid,omitempty"`           // 日程的唯一id
	OriginalTime int64  `json:"original_time,omitempty"` // 日程实例原始时间，非重复日程必为0。重复日程若为0则表示回复其所有实例，否则表示回复单个实例。
	Summary      string `json:"summary,omitempty"`       // 日程主题
}

type BatchGetMeetingRoomSummaryRespErrorEventUids struct {
	Uid          string `json:"uid,omitempty"`           // 日程的唯一id
	OriginalTime int64  `json:"original_time,omitempty"` // 日程实例原始时间，非重复日程必为0。重复日程若为0则表示回复其所有实例，否则表示回复单个实例。
	ErrorMsg     string `json:"error_msg,omitempty"`     // 错误信息
}