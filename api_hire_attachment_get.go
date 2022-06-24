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

// GetHireAttachment 获取招聘系统中附件的元信息，比如文件名、创建时间、文件url等
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/attachment/get
func (r *HireService) GetHireAttachment(ctx context.Context, request *GetHireAttachmentReq, options ...MethodOptionFunc) (*GetHireAttachmentResp, *Response, error) {
	if r.cli.mock.mockHireGetHireAttachment != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Hire#GetHireAttachment mock enable")
		return r.cli.mock.mockHireGetHireAttachment(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "GetHireAttachment",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/hire/v1/attachments/:attachment_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getHireAttachmentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHireGetHireAttachment mock HireGetHireAttachment method
func (r *Mock) MockHireGetHireAttachment(f func(ctx context.Context, request *GetHireAttachmentReq, options ...MethodOptionFunc) (*GetHireAttachmentResp, *Response, error)) {
	r.mockHireGetHireAttachment = f
}

// UnMockHireGetHireAttachment un-mock HireGetHireAttachment method
func (r *Mock) UnMockHireGetHireAttachment() {
	r.mockHireGetHireAttachment = nil
}

// GetHireAttachmentReq ...
type GetHireAttachmentReq struct {
	AttachmentID string `path:"attachment_id" json:"-"` // 附件id, 示例值："6435242341238"
}

// getHireAttachmentResp ...
type getHireAttachmentResp struct {
	Code int64                  `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *GetHireAttachmentResp `json:"data,omitempty"`
}

// GetHireAttachmentResp ...
type GetHireAttachmentResp struct {
	Attachment *GetHireAttachmentRespAttachment `json:"attachment,omitempty"` // 附件信息
}

// GetHireAttachmentRespAttachment ...
type GetHireAttachmentRespAttachment struct {
	ID         string `json:"id,omitempty"`          // 附件id
	URL        string `json:"url,omitempty"`         // 附件的url
	Name       string `json:"name,omitempty"`        // 附件文件名
	Mime       string `json:"mime,omitempty"`        // 媒体类型/MIME
	CreateTime int64  `json:"create_time,omitempty"` // 附件创建时间（单位ms）
}
