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

// GetContactJobLevelList 该接口可以获取租户职级列表。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/job_level/list
// new doc: https://open.feishu.cn/document/server-docs/contact-v3/job_level/list
func (r *ContactService) GetContactJobLevelList(ctx context.Context, request *GetContactJobLevelListReq, options ...MethodOptionFunc) (*GetContactJobLevelListResp, *Response, error) {
	if r.cli.mock.mockContactGetContactJobLevelList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Contact#GetContactJobLevelList mock enable")
		return r.cli.mock.mockContactGetContactJobLevelList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "GetContactJobLevelList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/job_levels",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getContactJobLevelListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactGetContactJobLevelList mock ContactGetContactJobLevelList method
func (r *Mock) MockContactGetContactJobLevelList(f func(ctx context.Context, request *GetContactJobLevelListReq, options ...MethodOptionFunc) (*GetContactJobLevelListResp, *Response, error)) {
	r.mockContactGetContactJobLevelList = f
}

// UnMockContactGetContactJobLevelList un-mock ContactGetContactJobLevelList method
func (r *Mock) UnMockContactGetContactJobLevelList() {
	r.mockContactGetContactJobLevelList = nil
}

// GetContactJobLevelListReq ...
type GetContactJobLevelListReq struct {
	PageSize  *int64  `query:"page_size" json:"-"`  // 分页大小, 示例值: 10, 默认值: `10`, 取值范围: `1` ～ `50`
	PageToken *string `query:"page_token" json:"-"` // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: AQD9/Rn9eij9Pm39ED40/RD/cIFmu77WxpxPB/2oHfQLZ+G8JG6tK7+ZnHiT7COhD2hMSICh/eBl7cpzU6JEC3J7COKNe4jrQ8ExwBCR
	Name      *string `query:"name" json:"-"`       // 传入该字段时, 可查询指定职级名称对应的职级信息, 示例值: 高级, 长度范围: `1` ～ `255` 字符
}

// GetContactJobLevelListResp ...
type GetContactJobLevelListResp struct {
	Items     []*GetContactJobLevelListRespItem `json:"items,omitempty"`      // 职级列表
	PageToken string                            `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore   bool                              `json:"has_more,omitempty"`   // 是否还有更多项
}

// GetContactJobLevelListRespItem ...
type GetContactJobLevelListRespItem struct {
	Name            string                                           `json:"name,omitempty"`             // 职级名称
	Description     string                                           `json:"description,omitempty"`      // 职级描述
	Order           int64                                            `json:"order,omitempty"`            // 职级的排序, 可填入自然数100-100000的数值, 系统按照数值大小从小到大排序。不填写该字段时, 默认新增排序在当前职级列表中最后位（最大值）
	Status          bool                                             `json:"status,omitempty"`           // 是否启用
	JobLevelID      string                                           `json:"job_level_id,omitempty"`     // 职级ID
	I18nName        []*GetContactJobLevelListRespItemI18nName        `json:"i18n_name,omitempty"`        // 多语言名称
	I18nDescription []*GetContactJobLevelListRespItemI18nDescription `json:"i18n_description,omitempty"` // 多语言描述
}

// GetContactJobLevelListRespItemI18nDescription ...
type GetContactJobLevelListRespItemI18nDescription struct {
	Locale string `json:"locale,omitempty"` // 语言版本
	Value  string `json:"value,omitempty"`  // 字段名
}

// GetContactJobLevelListRespItemI18nName ...
type GetContactJobLevelListRespItemI18nName struct {
	Locale string `json:"locale,omitempty"` // 语言版本
	Value  string `json:"value,omitempty"`  // 字段名
}

// getContactJobLevelListResp ...
type getContactJobLevelListResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *GetContactJobLevelListResp `json:"data,omitempty"`
}
