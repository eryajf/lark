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

// RemoveApprovalComment 删除某审批实例下的全部评论与评论回复。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/instance-comment/remove
func (r *ApprovalService) RemoveApprovalComment(ctx context.Context, request *RemoveApprovalCommentReq, options ...MethodOptionFunc) (*RemoveApprovalCommentResp, *Response, error) {
	if r.cli.mock.mockApprovalRemoveApprovalComment != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Approval#RemoveApprovalComment mock enable")
		return r.cli.mock.mockApprovalRemoveApprovalComment(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Approval",
		API:                   "RemoveApprovalComment",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/approval/v4/instances/:instance_id/comments/remove",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(removeApprovalCommentResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApprovalRemoveApprovalComment mock ApprovalRemoveApprovalComment method
func (r *Mock) MockApprovalRemoveApprovalComment(f func(ctx context.Context, request *RemoveApprovalCommentReq, options ...MethodOptionFunc) (*RemoveApprovalCommentResp, *Response, error)) {
	r.mockApprovalRemoveApprovalComment = f
}

// UnMockApprovalRemoveApprovalComment un-mock ApprovalRemoveApprovalComment method
func (r *Mock) UnMockApprovalRemoveApprovalComment() {
	r.mockApprovalRemoveApprovalComment = nil
}

// RemoveApprovalCommentReq ...
type RemoveApprovalCommentReq struct {
	InstanceID string  `path:"instance_id" json:"-"`   // 审批实例code（或者租户自定义审批实例ID）, 示例值: "6A123516-FB88-470D-A428-9AF58B71B3C0"
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 用户的 open id, union_id: 用户的 union id, user_id: 用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	UserID     *string `query:"user_id" json:"-"`      // 根据user_id_type填写用户ID, 示例值: "ou_806a18fb5bdf525e38ba219733bdbd73"
}

// RemoveApprovalCommentResp ...
type RemoveApprovalCommentResp struct {
	InstanceID string `json:"instance_id,omitempty"` // 审批实例code
	ExternalID string `json:"external_id,omitempty"` // 租户自定义审批实例ID
}

// removeApprovalCommentResp ...
type removeApprovalCommentResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *RemoveApprovalCommentResp `json:"data,omitempty"`
}
