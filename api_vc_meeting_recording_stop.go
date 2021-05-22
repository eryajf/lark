// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// StopVCMeetingRecording 在会议中停止录制。
//
// 会议正在录制中，且操作者具有相应权限（如果操作者为用户，必须是会中当前主持人）
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting-recording/stop
func (r *VCService) StopVCMeetingRecording(ctx context.Context, request *StopVCMeetingRecordingReq, options ...MethodOptionFunc) (*StopVCMeetingRecordingResp, *Response, error) {
	if r.cli.mock.mockVCStopVCMeetingRecording != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] VC#StopVCMeetingRecording mock enable")
		return r.cli.mock.mockVCStopVCMeetingRecording(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "VC",
		API:                 "StopVCMeetingRecording",
		Method:              "PATCH",
		URL:                 "https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id/recording/stop",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(stopVCMeetingRecordingResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockVCStopVCMeetingRecording(f func(ctx context.Context, request *StopVCMeetingRecordingReq, options ...MethodOptionFunc) (*StopVCMeetingRecordingResp, *Response, error)) {
	r.mockVCStopVCMeetingRecording = f
}

func (r *Mock) UnMockVCStopVCMeetingRecording() {
	r.mockVCStopVCMeetingRecording = nil
}

type StopVCMeetingRecordingReq struct {
	MeetingID string `path:"meeting_id" json:"-"` // 会议ID, 示例值: "6911188411932033028"
}

type stopVCMeetingRecordingResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *StopVCMeetingRecordingResp `json:"data,omitempty"`
}

type StopVCMeetingRecordingResp struct{}