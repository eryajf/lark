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

// UpdateDriveCommentPatch 解决或恢复云文档中的评论。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-comment/patch
func (r *DriveService) UpdateDriveCommentPatch(ctx context.Context, request *UpdateDriveCommentPatchReq, options ...MethodOptionFunc) (*UpdateDriveCommentPatchResp, *Response, error) {
	if r.cli.mock.mockDriveUpdateDriveCommentPatch != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#UpdateDriveCommentPatch mock enable")
		return r.cli.mock.mockDriveUpdateDriveCommentPatch(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "UpdateDriveCommentPatch",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/files/:file_token/comments/:comment_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateDriveCommentPatchResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveUpdateDriveCommentPatch mock DriveUpdateDriveCommentPatch method
func (r *Mock) MockDriveUpdateDriveCommentPatch(f func(ctx context.Context, request *UpdateDriveCommentPatchReq, options ...MethodOptionFunc) (*UpdateDriveCommentPatchResp, *Response, error)) {
	r.mockDriveUpdateDriveCommentPatch = f
}

// UnMockDriveUpdateDriveCommentPatch un-mock DriveUpdateDriveCommentPatch method
func (r *Mock) UnMockDriveUpdateDriveCommentPatch() {
	r.mockDriveUpdateDriveCommentPatch = nil
}

// UpdateDriveCommentPatchReq ...
type UpdateDriveCommentPatchReq struct {
	FileType  FileType `query:"file_type" json:"-"` // 文档类型, 示例值："doc", 可选值有: `doc`：文档, `sheet`：表格, `file`：文件, `docx`：新版文档
	FileToken string   `path:"file_token" json:"-"` // 文档token, 示例值："doccnGp4UK1UskrOEJwBXd3****"
	CommentID string   `path:"comment_id" json:"-"` // 评论ID, 示例值："6916106822734578184"
	IsSolved  bool     `json:"is_solved,omitempty"` // 评论解决标志, 示例值：true
}

// updateDriveCommentPatchResp ...
type updateDriveCommentPatchResp struct {
	Code int64                        `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                       `json:"msg,omitempty"`  // 错误描述
	Data *UpdateDriveCommentPatchResp `json:"data,omitempty"`
}

// UpdateDriveCommentPatchResp ...
type UpdateDriveCommentPatchResp struct {
}
