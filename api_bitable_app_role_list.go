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

// GetBitableAppRoleList 列出自定义权限
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-role/list
func (r *BitableService) GetBitableAppRoleList(ctx context.Context, request *GetBitableAppRoleListReq, options ...MethodOptionFunc) (*GetBitableAppRoleListResp, *Response, error) {
	if r.cli.mock.mockBitableGetBitableAppRoleList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#GetBitableAppRoleList mock enable")
		return r.cli.mock.mockBitableGetBitableAppRoleList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "GetBitableAppRoleList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/bitable/v1/apps/:app_token/roles",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getBitableAppRoleListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBitableGetBitableAppRoleList mock BitableGetBitableAppRoleList method
func (r *Mock) MockBitableGetBitableAppRoleList(f func(ctx context.Context, request *GetBitableAppRoleListReq, options ...MethodOptionFunc) (*GetBitableAppRoleListResp, *Response, error)) {
	r.mockBitableGetBitableAppRoleList = f
}

// UnMockBitableGetBitableAppRoleList un-mock BitableGetBitableAppRoleList method
func (r *Mock) UnMockBitableGetBitableAppRoleList() {
	r.mockBitableGetBitableAppRoleList = nil
}

// GetBitableAppRoleListReq ...
type GetBitableAppRoleListReq struct {
	PageSize  *int64  `query:"page_size" json:"-"`  // 分页大小, 示例值：10, 最大值：`30`
	PageToken *string `query:"page_token" json:"-"` // 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果, 示例值："roljRpwIUt"
	AppToken  string  `path:"app_token" json:"-"`   // bitable app token, 示例值："appbcbWCzen6D8dezhoCH2RpMAh"
}

// getBitableAppRoleListResp ...
type getBitableAppRoleListResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *GetBitableAppRoleListResp `json:"data,omitempty"`
}

// GetBitableAppRoleListResp ...
type GetBitableAppRoleListResp struct {
	Items     []*GetBitableAppRoleListRespItem `json:"items,omitempty"`      // 自定义权限列表
	PageToken string                           `json:"page_token,omitempty"` // 分页标记，当 has_more 为 true 时，会同时返回新的 page_token，否则不返回 page_token
	HasMore   bool                             `json:"has_more,omitempty"`   // 是否还有更多项
	Total     int64                            `json:"total,omitempty"`      // 总数
}

// GetBitableAppRoleListRespItem ...
type GetBitableAppRoleListRespItem struct {
	RoleName   string                                    `json:"role_name,omitempty"`   // 自定义权限的名字
	RoleID     string                                    `json:"role_id,omitempty"`     // 自定义权限的id
	TableRoles []*GetBitableAppRoleListRespItemTableRole `json:"table_roles,omitempty"` // 数据表权限
}

// GetBitableAppRoleListRespItemTableRole ...
type GetBitableAppRoleListRespItemTableRole struct {
	TableName string                                         `json:"table_name,omitempty"` // 数据表名
	TablePerm int64                                          `json:"table_perm,omitempty"` // 数据表权限，`协作者可编辑自己的记录`和`可编辑指定字段`是`可编辑记录`的特殊情况，可通过指定`rec_rule`或`field_perm`参数实现相同的效果, 可选值有: `0`：无权限, `1`：可阅读, `2`：可编辑记录, `4`：可编辑字段和记录
	RecRule   *GetBitableAppRoleListRespItemTableRoleRecRule `json:"rec_rule,omitempty"`   // 记录筛选条件，在table_perm为1或2时有意义，用于指定可编辑或可阅读某些记录
	FieldPerm map[string]int64                               `json:"field_perm,omitempty"` // 字段权限，仅在table_perm为2时有意义，设置字段可编辑或可阅读。类型为 map，key 是字段名，value 是字段权限。,**value 枚举值有：**, `1`：可阅读, `2`：可编辑
}

// GetBitableAppRoleListRespItemTableRoleRecRule ...
type GetBitableAppRoleListRespItemTableRoleRecRule struct {
	Conditions  []*GetBitableAppRoleListRespItemTableRoleRecRuleCondition `json:"conditions,omitempty"`  // 记录筛选条件
	Conjunction string                                                    `json:"conjunction,omitempty"` // 多个筛选条件的关系, 可选值有: `and`：与, `or`：或
	OtherPerm   int64                                                     `json:"other_perm,omitempty"`  // 其他记录权限，仅在table_perm为2时有意义, 可选值有: `0`：禁止查看, `1`：仅可阅读
}

// GetBitableAppRoleListRespItemTableRoleRecRuleCondition ...
type GetBitableAppRoleListRespItemTableRoleRecRuleCondition struct {
	FieldName string   `json:"field_name,omitempty"` // 字段名，记录筛选条件是`创建人包含访问者本人`时，此参数值为""
	Operator  string   `json:"operator,omitempty"`   // 运算符, 可选值有: `is`：等于, `isNot`：不等于, `contains`：包含, `doesNotContain`：不包含, `isEmpty`：为空, `isNotEmpty`：不为空
	Value     []string `json:"value,omitempty"`      // 单选或多选字段的选项id
	FieldType int64    `json:"field_type,omitempty"` // 字段类型
}
