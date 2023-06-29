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

// SendUrgentSmsMessage 对指定消息进行应用内加急与短信加急。
//
// 特别说明:
// - 通过接口产生的短信加急将消耗企业的加急额度, 请慎重调用
// - 通过[租户管理后台](https://admin.feishu.cn/)-费用中心-短信/电话加急 可以查看当前额度
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)
// - 只能加急机器人自己发送的消息
// - 加急时机器人仍需要在加急消息所在的群组中
// - 加急用户的未读加急总数不能超过200条
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/urgent_sms
// new doc: https://open.feishu.cn/document/server-docs/im-v1/buzz-messages/urgent_sms
func (r *MessageService) SendUrgentSmsMessage(ctx context.Context, request *SendUrgentSmsMessageReq, options ...MethodOptionFunc) (*SendUrgentSmsMessageResp, *Response, error) {
	if r.cli.mock.mockMessageSendUrgentSmsMessage != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Message#SendUrgentSmsMessage mock enable")
		return r.cli.mock.mockMessageSendUrgentSmsMessage(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Message",
		API:                   "SendUrgentSmsMessage",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/messages/:message_id/urgent_sms",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(sendUrgentSmsMessageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMessageSendUrgentSmsMessage mock MessageSendUrgentSmsMessage method
func (r *Mock) MockMessageSendUrgentSmsMessage(f func(ctx context.Context, request *SendUrgentSmsMessageReq, options ...MethodOptionFunc) (*SendUrgentSmsMessageResp, *Response, error)) {
	r.mockMessageSendUrgentSmsMessage = f
}

// UnMockMessageSendUrgentSmsMessage un-mock MessageSendUrgentSmsMessage method
func (r *Mock) UnMockMessageSendUrgentSmsMessage() {
	r.mockMessageSendUrgentSmsMessage = nil
}

// SendUrgentSmsMessageReq ...
type SendUrgentSmsMessageReq struct {
	MessageID  string   `path:"message_id" json:"-"`    // 待加急的消息ID, 详情参见[消息ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/intro#ac79c1c2), 注意: 不支持批量消息ID（bm_xxx）, 示例值: "om_dc13264520392913993dd051dba21dcf"
	UserIDType IDType   `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	UserIDList []string `json:"user_id_list,omitempty"` // 目标用户的ID, 列表不可为空；推荐使用 OpenID, 获取方式可参考文档[如何获取 Open ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), 注意: 请确保所填的用户ID正确, 并且用户在加急消息所在的群组中, 示例值: ["ou_6yf8af6bgb9100449565764t3382b168"]
}

// SendUrgentSmsMessageResp ...
type SendUrgentSmsMessageResp struct {
	InvalidUserIDList []string `json:"invalid_user_id_list,omitempty"` // 无效的用户ID
}

// sendUrgentSmsMessageResp ...
type sendUrgentSmsMessageResp struct {
	Code int64                     `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                    `json:"msg,omitempty"`  // 错误描述
	Data *SendUrgentSmsMessageResp `json:"data,omitempty"`
}
