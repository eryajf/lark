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

// AddBotToChat
//
// 为了更好地提升该接口的安全性，我们对其进行了升级，请尽快迁移至[新版本>>](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/create)
// 拉机器人进群<br>
// **权限说明** ：需要启用机器人能力；机器人的owner需要已经在群里
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uYDO04iN4QjL2gDN
//
// Deprecated
func (r *BotService) AddBotToChat(ctx context.Context, request *AddBotToChatReq, options ...MethodOptionFunc) (*AddBotToChatResp, *Response, error) {
	if r.cli.mock.mockBotAddBotToChat != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bot#AddBotToChat mock enable")
		return r.cli.mock.mockBotAddBotToChat(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bot",
		API:                   "AddBotToChat",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/bot/v4/add",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(addBotToChatResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBotAddBotToChat mock BotAddBotToChat method
func (r *Mock) MockBotAddBotToChat(f func(ctx context.Context, request *AddBotToChatReq, options ...MethodOptionFunc) (*AddBotToChatResp, *Response, error)) {
	r.mockBotAddBotToChat = f
}

// UnMockBotAddBotToChat un-mock BotAddBotToChat method
func (r *Mock) UnMockBotAddBotToChat() {
	r.mockBotAddBotToChat = nil
}

// AddBotToChatReq ...
type AddBotToChatReq struct {
	ChatID string `json:"chat_id,omitempty"` // 群的id
}

// addBotToChatResp ...
type addBotToChatResp struct {
	Code int64             `json:"code,omitempty"` // 返回码，非 0 表示失败
	Msg  string            `json:"msg,omitempty"`  // 返回码描述
	Data *AddBotToChatResp `json:"data,omitempty"`
}

// AddBotToChatResp ...
type AddBotToChatResp struct {
}
