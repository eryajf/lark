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

// DeleteMessageReaction 删除指定消息的表情回复（reaction即表情回复，本说明文档统一用“reaction”代称）。
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/home/develop-a-bot-in-5-minutes/create-an-app)
// - 只能删除真实存在的reaction，并且删除reaction请求的操作者必须是reaction的原始添加者
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/delete
func (r *MessageService) DeleteMessageReaction(ctx context.Context, request *DeleteMessageReactionReq, options ...MethodOptionFunc) (*DeleteMessageReactionResp, *Response, error) {
	if r.cli.mock.mockMessageDeleteMessageReaction != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Message#DeleteMessageReaction mock enable")
		return r.cli.mock.mockMessageDeleteMessageReaction(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Message",
		API:                   "DeleteMessageReaction",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/messages/:message_id/reactions/:reaction_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(deleteMessageReactionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMessageDeleteMessageReaction mock MessageDeleteMessageReaction method
func (r *Mock) MockMessageDeleteMessageReaction(f func(ctx context.Context, request *DeleteMessageReactionReq, options ...MethodOptionFunc) (*DeleteMessageReactionResp, *Response, error)) {
	r.mockMessageDeleteMessageReaction = f
}

// UnMockMessageDeleteMessageReaction un-mock MessageDeleteMessageReaction method
func (r *Mock) UnMockMessageDeleteMessageReaction() {
	r.mockMessageDeleteMessageReaction = nil
}

// DeleteMessageReactionReq ...
type DeleteMessageReactionReq struct {
	MessageID  string `path:"message_id" json:"-"`  // 待删除reaction的消息ID, 示例值："om_8964d1b4*********2b31383276113"
	ReactionID string `path:"reaction_id" json:"-"` // 待删除reaction的资源id, 示例值："ZCaCIjUBVVWSrm5L-3ZTw*************sNa8dHVplEzzSfJVUVLMLcS_"
}

// deleteMessageReactionResp ...
type deleteMessageReactionResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *DeleteMessageReactionResp `json:"data,omitempty"`
}

// DeleteMessageReactionResp ...
type DeleteMessageReactionResp struct {
	ReactionID   string                                 `json:"reaction_id,omitempty"`   // reaction资源ID
	Operator     *DeleteMessageReactionRespOperator     `json:"operator,omitempty"`      // 添加reaction的操作人
	ActionTime   string                                 `json:"action_time,omitempty"`   // reaction动作的的unix timestamp(单位:ms)
	ReactionType *DeleteMessageReactionRespReactionType `json:"reaction_type,omitempty"` // reaction资源类型
}

// DeleteMessageReactionRespOperator ...
type DeleteMessageReactionRespOperator struct {
	OperatorID   string `json:"operator_id,omitempty"`   // 操作人ID
	OperatorType string `json:"operator_type,omitempty"` // 操作人身份，用户或应用, 可选值有: `app`："app", `user`："user"
}

// DeleteMessageReactionRespReactionType ...
type DeleteMessageReactionRespReactionType struct {
	EmojiType string `json:"emoji_type,omitempty"` // emoji类型 [emoji类型列举](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message-reaction/emojis-introduce)
}
