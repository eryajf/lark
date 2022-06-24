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

// PrepareUploadDriveFile 发送初始化请求获取上传事务ID和分块策略，目前是以4MB大小进行定长分片。
//
// 你在24小时内可保存上传事务ID和上传进度，以便可以恢复上传
// 该接口不支持太高的并发，且调用频率上限为5QPS
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/upload_prepare
func (r *DriveService) PrepareUploadDriveFile(ctx context.Context, request *PrepareUploadDriveFileReq, options ...MethodOptionFunc) (*PrepareUploadDriveFileResp, *Response, error) {
	if r.cli.mock.mockDrivePrepareUploadDriveFile != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#PrepareUploadDriveFile mock enable")
		return r.cli.mock.mockDrivePrepareUploadDriveFile(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "PrepareUploadDriveFile",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/files/upload_prepare",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(prepareUploadDriveFileResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDrivePrepareUploadDriveFile mock DrivePrepareUploadDriveFile method
func (r *Mock) MockDrivePrepareUploadDriveFile(f func(ctx context.Context, request *PrepareUploadDriveFileReq, options ...MethodOptionFunc) (*PrepareUploadDriveFileResp, *Response, error)) {
	r.mockDrivePrepareUploadDriveFile = f
}

// UnMockDrivePrepareUploadDriveFile un-mock DrivePrepareUploadDriveFile method
func (r *Mock) UnMockDrivePrepareUploadDriveFile() {
	r.mockDrivePrepareUploadDriveFile = nil
}

// PrepareUploadDriveFileReq ...
type PrepareUploadDriveFileReq struct {
	FileName   string `json:"file_name,omitempty"`   // 文件名, 示例值："test.txt", 最大长度：`250` 字符
	ParentType string `json:"parent_type,omitempty"` // 上传点类型, 示例值："explorer", 可选值有: `explorer`：云空间
	ParentNode string `json:"parent_node,omitempty"` // 文件夹的token, 示例值："fldcnaBcdEfghdis"
	Size       int64  `json:"size,omitempty"`        // 文件大小, 示例值：1024, 最小值：`0`
}

// prepareUploadDriveFileResp ...
type prepareUploadDriveFileResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *PrepareUploadDriveFileResp `json:"data,omitempty"`
}

// PrepareUploadDriveFileResp ...
type PrepareUploadDriveFileResp struct {
	UploadID  string `json:"upload_id,omitempty"`  // 分片上传事务ID
	BlockSize int64  `json:"block_size,omitempty"` // 分片大小策略
	BlockNum  int64  `json:"block_num,omitempty"`  // 分片数量
}
