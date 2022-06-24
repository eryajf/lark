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

// DeleteTaskCollaborator 该接口用于删除任务执行者
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-collaborator/delete
func (r *TaskService) DeleteTaskCollaborator(ctx context.Context, request *DeleteTaskCollaboratorReq, options ...MethodOptionFunc) (*DeleteTaskCollaboratorResp, *Response, error) {
	if r.cli.mock.mockTaskDeleteTaskCollaborator != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Task#DeleteTaskCollaborator mock enable")
		return r.cli.mock.mockTaskDeleteTaskCollaborator(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Task",
		API:                   "DeleteTaskCollaborator",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v1/tasks/:task_id/collaborators/:collaborator_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(deleteTaskCollaboratorResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockTaskDeleteTaskCollaborator mock TaskDeleteTaskCollaborator method
func (r *Mock) MockTaskDeleteTaskCollaborator(f func(ctx context.Context, request *DeleteTaskCollaboratorReq, options ...MethodOptionFunc) (*DeleteTaskCollaboratorResp, *Response, error)) {
	r.mockTaskDeleteTaskCollaborator = f
}

// UnMockTaskDeleteTaskCollaborator un-mock TaskDeleteTaskCollaborator method
func (r *Mock) UnMockTaskDeleteTaskCollaborator() {
	r.mockTaskDeleteTaskCollaborator = nil
}

// DeleteTaskCollaboratorReq ...
type DeleteTaskCollaboratorReq struct {
	TaskID         string `path:"task_id" json:"-"`         // 任务 ID, 示例值："83912691-2e43-47fc-94a4-d512e03984fa"
	CollaboratorID string `path:"collaborator_id" json:"-"` // 任务协作者 ID（Open ID）, 示例值："ou_99e1a581b36ecc4862cbfbce123f346a"
}

// deleteTaskCollaboratorResp ...
type deleteTaskCollaboratorResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *DeleteTaskCollaboratorResp `json:"data,omitempty"`
}

// DeleteTaskCollaboratorResp ...
type DeleteTaskCollaboratorResp struct {
}
