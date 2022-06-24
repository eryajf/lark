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

// HighlightBaikeEntity 传入一句话，智能识别句中对应的词条，并返回词条位置和 entity_id，可在外部系统中快速实现百科词条智能高亮
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/entity/highlight
func (r *BaikeService) HighlightBaikeEntity(ctx context.Context, request *HighlightBaikeEntityReq, options ...MethodOptionFunc) (*HighlightBaikeEntityResp, *Response, error) {
	if r.cli.mock.mockBaikeHighlightBaikeEntity != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Baike#HighlightBaikeEntity mock enable")
		return r.cli.mock.mockBaikeHighlightBaikeEntity(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Baike",
		API:                   "HighlightBaikeEntity",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/baike/v1/entities/highlight",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(highlightBaikeEntityResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBaikeHighlightBaikeEntity mock BaikeHighlightBaikeEntity method
func (r *Mock) MockBaikeHighlightBaikeEntity(f func(ctx context.Context, request *HighlightBaikeEntityReq, options ...MethodOptionFunc) (*HighlightBaikeEntityResp, *Response, error)) {
	r.mockBaikeHighlightBaikeEntity = f
}

// UnMockBaikeHighlightBaikeEntity un-mock BaikeHighlightBaikeEntity method
func (r *Mock) UnMockBaikeHighlightBaikeEntity() {
	r.mockBaikeHighlightBaikeEntity = nil
}

// HighlightBaikeEntityReq ...
type HighlightBaikeEntityReq struct {
	Text string `json:"text,omitempty"` // 需要识别百科词条的内容（不超过1000字）, 示例值："企业百科是飞书提供的一款知识管理工具", 长度范围：`1` ～ `1000` 字符
}

// highlightBaikeEntityResp ...
type highlightBaikeEntityResp struct {
	Code int64                     `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                    `json:"msg,omitempty"`  // 错误描述
	Data *HighlightBaikeEntityResp `json:"data,omitempty"`
}

// HighlightBaikeEntityResp ...
type HighlightBaikeEntityResp struct {
	Phrases []*HighlightBaikeEntityRespPhrase `json:"phrases,omitempty"` // 识别到的词条信息
}

// HighlightBaikeEntityRespPhrase ...
type HighlightBaikeEntityRespPhrase struct {
	Name      string                              `json:"name,omitempty"`       // 识别到的关键词
	EntityIDs []string                            `json:"entity_ids,omitempty"` // 对应的词条 ID
	Span      *HighlightBaikeEntityRespPhraseSpan `json:"span,omitempty"`       // 词条所在位置
}

// HighlightBaikeEntityRespPhraseSpan ...
type HighlightBaikeEntityRespPhraseSpan struct {
	Start int64 `json:"start,omitempty"` // 关键词开始位置，从 0 开始计数（编码格式采用 utf-8）
	End   int64 `json:"end,omitempty"`   // 关键词结束位置，从 0 开始计数（编码格式采用 utf-8）
}
