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

// GetHireJobManager 根据职位 ID 获取职位上的招聘人员信息, 如招聘负责人、用人经理。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/job-manager/get
// new doc: https://open.feishu.cn/document/server-docs/hire-v1/recruitment-related-configuration/job/get-2
func (r *HireService) GetHireJobManager(ctx context.Context, request *GetHireJobManagerReq, options ...MethodOptionFunc) (*GetHireJobManagerResp, *Response, error) {
	if r.cli.mock.mockHireGetHireJobManager != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Hire#GetHireJobManager mock enable")
		return r.cli.mock.mockHireGetHireJobManager(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "GetHireJobManager",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/hire/v1/jobs/:job_id/managers/:manager_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getHireJobManagerResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHireGetHireJobManager mock HireGetHireJobManager method
func (r *Mock) MockHireGetHireJobManager(f func(ctx context.Context, request *GetHireJobManagerReq, options ...MethodOptionFunc) (*GetHireJobManagerResp, *Response, error)) {
	r.mockHireGetHireJobManager = f
}

// UnMockHireGetHireJobManager un-mock HireGetHireJobManager method
func (r *Mock) UnMockHireGetHireJobManager() {
	r.mockHireGetHireJobManager = nil
}

// GetHireJobManagerReq ...
type GetHireJobManagerReq struct {
	JobID      string  `path:"job_id" json:"-"`        // 职位 ID, 示例值: "1618209327096"
	ManagerID  string  `path:"manager_id" json:"-"`    // 此处传入职位 ID, 示例值: "1618209327096"
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), people_admin_id: 以 people_admin_id 来识别用户, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetHireJobManagerResp ...
type GetHireJobManagerResp struct {
	Info *GetHireJobManagerRespInfo `json:"info,omitempty"` // 职位负责人
}

// GetHireJobManagerRespInfo ...
type GetHireJobManagerRespInfo struct {
	ID                  string   `json:"id,omitempty"`                     // 职位 ID
	RecruiterID         string   `json:"recruiter_id,omitempty"`           // 招聘负责人 ID, 仅一位, 可通过用户相关接口获取用户 ID
	HiringManagerIDList []string `json:"hiring_manager_id_list,omitempty"` // 用人经理 ID 列表
	AssistantIDList     []string `json:"assistant_id_list,omitempty"`      // 协助人 ID 列表
}

// getHireJobManagerResp ...
type getHireJobManagerResp struct {
	Code int64                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *GetHireJobManagerResp `json:"data,omitempty"`
}
