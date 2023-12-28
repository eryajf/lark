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

// PrepareUploadDriveMedia 发送初始化请求获取上传事务ID和分块策略, 目前是以4MB大小进行定长分片。
//
// 您在24小时内可保存上传事务ID和上传进度, 以便可以恢复上传
// 该接口不支持太高的并发, 且调用频率上限为5QPS
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/upload_prepare
// new doc: https://open.feishu.cn/document/server-docs/docs/drive-v1/media/multipart-upload-media/upload_prepare
func (r *DriveService) PrepareUploadDriveMedia(ctx context.Context, request *PrepareUploadDriveMediaReq, options ...MethodOptionFunc) (*PrepareUploadDriveMediaResp, *Response, error) {
	if r.cli.mock.mockDrivePrepareUploadDriveMedia != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#PrepareUploadDriveMedia mock enable")
		return r.cli.mock.mockDrivePrepareUploadDriveMedia(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "PrepareUploadDriveMedia",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/medias/upload_prepare",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(prepareUploadDriveMediaResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDrivePrepareUploadDriveMedia mock DrivePrepareUploadDriveMedia method
func (r *Mock) MockDrivePrepareUploadDriveMedia(f func(ctx context.Context, request *PrepareUploadDriveMediaReq, options ...MethodOptionFunc) (*PrepareUploadDriveMediaResp, *Response, error)) {
	r.mockDrivePrepareUploadDriveMedia = f
}

// UnMockDrivePrepareUploadDriveMedia un-mock DrivePrepareUploadDriveMedia method
func (r *Mock) UnMockDrivePrepareUploadDriveMedia() {
	r.mockDrivePrepareUploadDriveMedia = nil
}

// PrepareUploadDriveMediaReq ...
type PrepareUploadDriveMediaReq struct {
	FileName   string  `json:"file_name,omitempty"`   // 文件名, 示例值: "demo.jpeg", 最大长度: `250` 字符
	ParentType string  `json:"parent_type,omitempty"` // 上传点类型, 示例值: "doc_image", 可选值有: doc_image: 文档图片。, docx_image: 新版文档图片, sheet_image: 电子表格图片。, doc_file: 文档文件。, docx_file: 新版文档文件, sheet_file: 电子表格文件。, vc_virtual_background: vc虚拟背景(灰度中, 暂未开放)。, bitable_image: 多维表格图片。, bitable_file: 多维表格文件。, moments: 同事圈(灰度中, 暂未开放)。, ccm_import_open: 云文档导入文件。
	ParentNode string  `json:"parent_node,omitempty"` // 上传点 Token, 用于指定素材将要上传到的具体文档或位置。点击 [这里](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/introduction) 了解各上传点类型及其对应的上传点 Token 的说明, 示例值: "doccnFivLCfJfblZjGZtxgabcef"
	Size       int64   `json:"size,omitempty"`        // 文件大小, 示例值: 1024, 最小值: `0`
	Extra      *string `json:"extra,omitempty"`       // 扩展信息(可选), 示例值: "{\"test\":\"test\"}"
}

// PrepareUploadDriveMediaResp ...
type PrepareUploadDriveMediaResp struct {
	UploadID  string `json:"upload_id,omitempty"`  // 分片上传事务ID
	BlockSize int64  `json:"block_size,omitempty"` // 分片大小策略
	BlockNum  int64  `json:"block_num,omitempty"`  // 分片数量
}

// prepareUploadDriveMediaResp ...
type prepareUploadDriveMediaResp struct {
	Code int64                        `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                       `json:"msg,omitempty"`  // 错误描述
	Data *PrepareUploadDriveMediaResp `json:"data,omitempty"`
}
