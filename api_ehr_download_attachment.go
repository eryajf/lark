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
	"io"
)

// DownloadEHRAttachments
//
// 根据文件 token 下载文件。
// 调用 「批量获取员工花名册信息」接口的返回值中，「文件」类型的字段 id，即是文件 token
// ![image.png](//sf1-ttcdn-tos.pstatp.com/obj/open-platform-opendoc/bed391d2a8ce6ed2d5985ea69bf92850_9GY1mnuDXP.png)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/ehr/ehr-v1/attachment/get
func (r *EHRService) DownloadEHRAttachments(ctx context.Context, request *DownloadEHRAttachmentsReq, options ...MethodOptionFunc) (*DownloadEHRAttachmentsResp, *Response, error) {
	if r.cli.mock.mockEHRDownloadEHRAttachments != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] EHR#DownloadEHRAttachments mock enable")
		return r.cli.mock.mockEHRDownloadEHRAttachments(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "EHR",
		API:                   "DownloadEHRAttachments",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/ehr/v1/attachments/:token",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(downloadEHRAttachmentsResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockEHRDownloadEHRAttachments mock EHRDownloadEHRAttachments method
func (r *Mock) MockEHRDownloadEHRAttachments(f func(ctx context.Context, request *DownloadEHRAttachmentsReq, options ...MethodOptionFunc) (*DownloadEHRAttachmentsResp, *Response, error)) {
	r.mockEHRDownloadEHRAttachments = f
}

// UnMockEHRDownloadEHRAttachments un-mock EHRDownloadEHRAttachments method
func (r *Mock) UnMockEHRDownloadEHRAttachments() {
	r.mockEHRDownloadEHRAttachments = nil
}

// DownloadEHRAttachmentsReq ...
type DownloadEHRAttachmentsReq struct {
	Token string `path:"token" json:"-"` // 文件 token, 示例值："09bf7b924f9a4a69875788891b5970d8"
}

// downloadEHRAttachmentsResp ...
type downloadEHRAttachmentsResp struct {
	IsFile bool                        `json:"is_file,omitempty"`
	Code   int64                       `json:"code,omitempty"`
	Msg    string                      `json:"msg,omitempty"`
	Data   *DownloadEHRAttachmentsResp `json:"data,omitempty"`
}

func (r *downloadEHRAttachmentsResp) SetReader(file io.Reader) {
	if r.Data == nil {
		r.Data = &DownloadEHRAttachmentsResp{}
	}
	r.Data.File = file
}

// DownloadEHRAttachmentsResp ...
type DownloadEHRAttachmentsResp struct {
	File io.Reader `json:"file,omitempty"`
}
