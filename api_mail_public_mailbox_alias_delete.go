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

// DeleteMailPublicMailboxAlias 删除公共邮箱别名
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox-alias/delete
func (r *MailService) DeleteMailPublicMailboxAlias(ctx context.Context, request *DeleteMailPublicMailboxAliasReq, options ...MethodOptionFunc) (*DeleteMailPublicMailboxAliasResp, *Response, error) {
	if r.cli.mock.mockMailDeleteMailPublicMailboxAlias != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Mail#DeleteMailPublicMailboxAlias mock enable")
		return r.cli.mock.mockMailDeleteMailPublicMailboxAlias(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Mail",
		API:                   "DeleteMailPublicMailboxAlias",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/mail/v1/public_mailboxes/:public_mailbox_id/aliases/:alias_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deleteMailPublicMailboxAliasResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMailDeleteMailPublicMailboxAlias mock MailDeleteMailPublicMailboxAlias method
func (r *Mock) MockMailDeleteMailPublicMailboxAlias(f func(ctx context.Context, request *DeleteMailPublicMailboxAliasReq, options ...MethodOptionFunc) (*DeleteMailPublicMailboxAliasResp, *Response, error)) {
	r.mockMailDeleteMailPublicMailboxAlias = f
}

// UnMockMailDeleteMailPublicMailboxAlias un-mock MailDeleteMailPublicMailboxAlias method
func (r *Mock) UnMockMailDeleteMailPublicMailboxAlias() {
	r.mockMailDeleteMailPublicMailboxAlias = nil
}

// DeleteMailPublicMailboxAliasReq ...
type DeleteMailPublicMailboxAliasReq struct {
	PublicMailboxID string `path:"public_mailbox_id" json:"-"` // 公共邮箱id或公共邮箱地址, 示例值："xxxxxx 或 xxx@xx.xxx"
	AliasID         string `path:"alias_id" json:"-"`          // 公共邮箱别名, 示例值："xxx@xx.xxx"
}

// deleteMailPublicMailboxAliasResp ...
type deleteMailPublicMailboxAliasResp struct {
	Code int64                             `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                            `json:"msg,omitempty"`  // 错误描述
	Data *DeleteMailPublicMailboxAliasResp `json:"data,omitempty"`
}

// DeleteMailPublicMailboxAliasResp ...
type DeleteMailPublicMailboxAliasResp struct {
}
