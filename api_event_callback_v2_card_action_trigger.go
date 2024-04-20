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
	"encoding/json"
)

// EventV2CardActionTrigger 卡片回传交互作用于飞书卡片的 请求回调 交互组件。当终端用户点击飞书卡片上的回传交互组件后, 你在开发者后台应用内注册的回调请求地址将会收到 卡片回传交互 回调。该回调包含了用户与卡片之间的交互信息。
//
// 你的业务服务器接收到回调请求后, 需要在 3 秒内响应回调请求, 声明通过弹出 Toast 提示、更新卡片、保持原内容不变等方式响应用户交互。
//
// doc: https://open.larkoffice.com/document/uAjLw4CM/ukzMukzMukzM/feishu-cards/card-callback-communication
func (r *EventCallbackService) HandlerEventV2CardActionTrigger(f EventV2CardActionTriggerHandler) {
	r.cli.eventHandler.eventV2CardActionTriggerHandler = f
}

// EventV2CardActionTriggerHandler event EventV2CardActionTrigger handler
type EventV2CardActionTriggerHandler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2CardActionTrigger) (string, error)

// EventV2CardActionTrigger ...
type EventV2CardActionTrigger struct {
	Operator     *EventV2CardActionTriggerOperator `json:"operator,omitempty"`      // 回调触发者信息。
	Token        string                            `json:"token,omitempty"`         // 更新卡片用的凭证, 有效期为 30 分钟, 最多可更新 2 次。
	Action       *EventV2CardActionTriggerAction   `json:"action,omitempty"`        // 交互信息。
	Host         string                            `json:"host,omitempty"`          // 卡片展示场景。
	DeliveryType string                            `json:"delivery_type,omitempty"` // 卡片分发类型, 固定取值为 `url_preview`, 表示链接预览卡片。仅链接预览卡片有此字段。
	Context      *EventV2CardActionTriggerContext  `json:"context,omitempty"`       // 展示场景上下文。
}

// EventV2CardActionTriggerAction ...
type EventV2CardActionTriggerAction struct {
	Value  json.RawMessage `json:"value,omitempty"`  // 交互元素的 value 字段值。
	Tag    string          `json:"tag,omitempty"`    // 交互元素的 tag 字段值。
	Option json.RawMessage `json:"option,omitempty"` // 选中 option 的 value。（button 元素不适用）
}

// EventV2CardActionTriggerContext ...
type EventV2CardActionTriggerContext struct {
	URL           string `json:"url,omitempty"`             // 链接地址（适用于链接预览场景）。
	PreviewToken  string `json:"preview_token,omitempty"`   // 链接预览的 token（适用于链接预览场景）。
	OpenMessageID string `json:"open_message_id,omitempty"` // 消息 ID。
	OpenChatID    string `json:"open_chat_id,omitempty"`    // 会话 ID。
}

// EventV2CardActionTriggerOperator ...
type EventV2CardActionTriggerOperator struct {
	TenantKey string `json:"tenant_key,omitempty"` // 回调触发者的 tenant key, 即租户唯一标识。
	UserID    string `json:"user_id,omitempty"`    // 回调触发者的 user_id。了解不同的用户 ID, 参见[用户身份概述](https://open.feishu.cn/document/home/user-identity-introduction/introduction)。
	OpenID    string `json:"open_id,omitempty"`    // 回调触发者的 open_id。
}