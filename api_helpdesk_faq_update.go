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

// UpdateHelpdeskFAQ 该接口用于修改知识库。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/faq/patch
func (r *HelpdeskService) UpdateHelpdeskFAQ(ctx context.Context, request *UpdateHelpdeskFAQReq, options ...MethodOptionFunc) (*UpdateHelpdeskFAQResp, *Response, error) {
	if r.cli.mock.mockHelpdeskUpdateHelpdeskFAQ != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#UpdateHelpdeskFAQ mock enable")
		return r.cli.mock.mockHelpdeskUpdateHelpdeskFAQ(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Helpdesk",
		API:                 "UpdateHelpdeskFAQ",
		Method:              "PATCH",
		URL:                 r.cli.openBaseURL + "/open-apis/helpdesk/v1/faqs/:id",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
		NeedHelpdeskAuth:    true,
	}
	resp := new(updateHelpdeskFAQResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHelpdeskUpdateHelpdeskFAQ mock HelpdeskUpdateHelpdeskFAQ method
func (r *Mock) MockHelpdeskUpdateHelpdeskFAQ(f func(ctx context.Context, request *UpdateHelpdeskFAQReq, options ...MethodOptionFunc) (*UpdateHelpdeskFAQResp, *Response, error)) {
	r.mockHelpdeskUpdateHelpdeskFAQ = f
}

// UnMockHelpdeskUpdateHelpdeskFAQ un-mock HelpdeskUpdateHelpdeskFAQ method
func (r *Mock) UnMockHelpdeskUpdateHelpdeskFAQ() {
	r.mockHelpdeskUpdateHelpdeskFAQ = nil
}

// UpdateHelpdeskFAQReq ...
type UpdateHelpdeskFAQReq struct {
	ID  string                   `path:"id" json:"-"`   // 知识库ID, 示例值: "6856395634652479491"
	FAQ *UpdateHelpdeskFAQReqFAQ `json:"faq,omitempty"` // 修改的知识库内容
}

// UpdateHelpdeskFAQReqFAQ ...
type UpdateHelpdeskFAQReqFAQ struct {
	CategoryID     *string  `json:"category_id,omitempty"`     // 知识库分类ID, 示例值: "6836004780707807251"
	Question       string   `json:"question,omitempty"`        // 问题, 示例值: "问题"
	Answer         *string  `json:"answer,omitempty"`          // 答案, 示例值: "答案"
	AnswerRichtext *string  `json:"answer_richtext,omitempty"` // 富文本答案和答案必须有一个必填。Json Array格式, 富文本结构请见[了解更多: 富文本](https://open.feishu.cn/document/ukTMukTMukTM/uITM0YjLyEDN24iMxQjN), 示例值: "[{, "content": "这只是一个测试, 医保问题", "type": "text", }]"
	Tags           []string `json:"tags,omitempty"`            // 相似问题
}

// UpdateHelpdeskFAQResp ...
type UpdateHelpdeskFAQResp struct {
}

// updateHelpdeskFAQResp ...
type updateHelpdeskFAQResp struct {
	Code int64                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *UpdateHelpdeskFAQResp `json:"data,omitempty"`
}