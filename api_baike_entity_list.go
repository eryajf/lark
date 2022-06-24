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

// GetBaikeEntityList 分页拉取词条列表数据，支持拉取租户内的全部词条
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/entity/list
func (r *BaikeService) GetBaikeEntityList(ctx context.Context, request *GetBaikeEntityListReq, options ...MethodOptionFunc) (*GetBaikeEntityListResp, *Response, error) {
	if r.cli.mock.mockBaikeGetBaikeEntityList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Baike#GetBaikeEntityList mock enable")
		return r.cli.mock.mockBaikeGetBaikeEntityList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Baike",
		API:                   "GetBaikeEntityList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/baike/v1/entities",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getBaikeEntityListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBaikeGetBaikeEntityList mock BaikeGetBaikeEntityList method
func (r *Mock) MockBaikeGetBaikeEntityList(f func(ctx context.Context, request *GetBaikeEntityListReq, options ...MethodOptionFunc) (*GetBaikeEntityListResp, *Response, error)) {
	r.mockBaikeGetBaikeEntityList = f
}

// UnMockBaikeGetBaikeEntityList un-mock BaikeGetBaikeEntityList method
func (r *Mock) UnMockBaikeGetBaikeEntityList() {
	r.mockBaikeGetBaikeEntityList = nil
}

// GetBaikeEntityListReq ...
type GetBaikeEntityListReq struct {
	PageSize   *int64  `query:"page_size" json:"-"`    // 分页大小, 示例值：20, 最大值：`100`
	PageToken  *string `query:"page_token" json:"-"`   // 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果, 示例值："b152fa6e6f62a291019a04c3a93f365f8ac641910506ff15ff4cad6534e087cb4ed8fa2c"
	Provider   *string `query:"provider" json:"-"`     // 相关外部系统【可用来过滤词条数据】, 示例值："星云", 长度范围：`2` ～ `32` 字符
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值："open_id", 可选值有: ,<md-enum>,<md-enum-item key="open_id" >用户的 open id</md-enum-item>,<md-enum-item key="union_id" >用户的 union id</md-enum-item>,<md-enum-item key="user_id" >用户的 user id</md-enum-item>,</md-enum>, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// getBaikeEntityListResp ...
type getBaikeEntityListResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *GetBaikeEntityListResp `json:"data,omitempty"`
}

// GetBaikeEntityListResp ...
type GetBaikeEntityListResp struct {
	Entities  []*GetBaikeEntityListRespEntitie `json:"entities,omitempty"`   // 词条列表
	PageToken string                           `json:"page_token,omitempty"` // 分页标记，当 has_more 为 true 时，会同时返回新的 page_token，否则不返回 page_token
}

// GetBaikeEntityListRespEntitie ...
type GetBaikeEntityListRespEntitie struct {
	ID          string                                    `json:"id,omitempty"`           // 词条 ID （需要更新某个词条时填写，若是创建新词条可不填写）
	MainKeys    []*GetBaikeEntityListRespEntitieMainKey   `json:"main_keys,omitempty"`    // 词条名
	Aliases     []*GetBaikeEntityListRespEntitieAliase    `json:"aliases,omitempty"`      // 别名
	Description string                                    `json:"description,omitempty"`  // 词条释义（纯文本格式）
	CreateTime  string                                    `json:"create_time,omitempty"`  // 词条创建时间
	UpdateTime  string                                    `json:"update_time,omitempty"`  // 词条最近更新时间
	RelatedMeta *GetBaikeEntityListRespEntitieRelatedMeta `json:"related_meta,omitempty"` // 更多相关信息
	Categories  []string                                  `json:"categories,omitempty"`   // 词条标签
	Statistics  *GetBaikeEntityListRespEntitieStatistics  `json:"statistics,omitempty"`   // 当前词条收到的反馈数据
	OuterInfo   *GetBaikeEntityListRespEntitieOuterInfo   `json:"outer_info,omitempty"`   // 外部系统关联数据
	RichText    string                                    `json:"rich_text,omitempty"`    // 富文本格式（当填写富文本内容时，description字段将会失效可不填写），支持的格式参考[企业百科指南](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/overview)中的释义部分
}

// GetBaikeEntityListRespEntitieMainKey ...
type GetBaikeEntityListRespEntitieMainKey struct {
	Key           string                                             `json:"key,omitempty"`            // 名称的值
	DisplayStatus *GetBaikeEntityListRespEntitieMainKeyDisplayStatus `json:"display_status,omitempty"` // 名称展示范围
}

// GetBaikeEntityListRespEntitieMainKeyDisplayStatus ...
type GetBaikeEntityListRespEntitieMainKeyDisplayStatus struct {
	AllowHighlight bool `json:"allow_highlight,omitempty"` // 对应名称是否在消息/云文档高亮
	AllowSearch    bool `json:"allow_search,omitempty"`    // 对应名称是否在搜索结果中展示
}

// GetBaikeEntityListRespEntitieAliase ...
type GetBaikeEntityListRespEntitieAliase struct {
	Key           string                                            `json:"key,omitempty"`            // 名称的值
	DisplayStatus *GetBaikeEntityListRespEntitieAliaseDisplayStatus `json:"display_status,omitempty"` // 名称展示范围
}

// GetBaikeEntityListRespEntitieAliaseDisplayStatus ...
type GetBaikeEntityListRespEntitieAliaseDisplayStatus struct {
	AllowHighlight bool `json:"allow_highlight,omitempty"` // 对应名称是否在消息/云文档高亮
	AllowSearch    bool `json:"allow_search,omitempty"`    // 对应名称是否在搜索结果中展示
}

// GetBaikeEntityListRespEntitieRelatedMeta ...
type GetBaikeEntityListRespEntitieRelatedMeta struct {
	Users           []*GetBaikeEntityListRespEntitieRelatedMetaUser           `json:"users,omitempty"`           // 相关联系人
	Chats           []*GetBaikeEntityListRespEntitieRelatedMetaChat           `json:"chats,omitempty"`           // 相关服务中的相关公开群
	Docs            []*GetBaikeEntityListRespEntitieRelatedMetaDoc            `json:"docs,omitempty"`            // 相关云文档
	Oncalls         []*GetBaikeEntityListRespEntitieRelatedMetaOncall         `json:"oncalls,omitempty"`         // 相关服务中的相关值班号
	Links           []*GetBaikeEntityListRespEntitieRelatedMetaLink           `json:"links,omitempty"`           // 相关链接
	Abbreviations   []*GetBaikeEntityListRespEntitieRelatedMetaAbbreviation   `json:"abbreviations,omitempty"`   // 相关词条
	Classifications []*GetBaikeEntityListRespEntitieRelatedMetaClassification `json:"classifications,omitempty"` // 当前词条所属分类<br>,词条只能属于二级分类，且每个一级分类下只能选择一个二级分类。
}

// GetBaikeEntityListRespEntitieRelatedMetaUser ...
type GetBaikeEntityListRespEntitieRelatedMetaUser struct {
	ID    string `json:"id,omitempty"`    // 对应相关信息 ID
	Title string `json:"title,omitempty"` // 对应相关信息的描述，如相关联系人的描述、相关链接的标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// GetBaikeEntityListRespEntitieRelatedMetaChat ...
type GetBaikeEntityListRespEntitieRelatedMetaChat struct {
	ID    string `json:"id,omitempty"`    // 对应相关信息 ID
	Title string `json:"title,omitempty"` // 对应相关信息的描述，如相关联系人的描述、相关链接的标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// GetBaikeEntityListRespEntitieRelatedMetaDoc ...
type GetBaikeEntityListRespEntitieRelatedMetaDoc struct {
	Title string `json:"title,omitempty"` // 对应相关信息的描述，如相关联系人的描述、相关链接的标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// GetBaikeEntityListRespEntitieRelatedMetaOncall ...
type GetBaikeEntityListRespEntitieRelatedMetaOncall struct {
	ID    string `json:"id,omitempty"`    // 对应相关信息 ID
	Title string `json:"title,omitempty"` // 对应相关信息的描述，如相关联系人的描述、相关链接的标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// GetBaikeEntityListRespEntitieRelatedMetaLink ...
type GetBaikeEntityListRespEntitieRelatedMetaLink struct {
	Title string `json:"title,omitempty"` // 对应相关信息的描述，如相关联系人的描述、相关链接的标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// GetBaikeEntityListRespEntitieRelatedMetaAbbreviation ...
type GetBaikeEntityListRespEntitieRelatedMetaAbbreviation struct {
	ID string `json:"id,omitempty"` // 相关词条 ID
}

// GetBaikeEntityListRespEntitieRelatedMetaClassification ...
type GetBaikeEntityListRespEntitieRelatedMetaClassification struct {
	ID       string `json:"id,omitempty"`        // 二级分类 ID
	Name     string `json:"name,omitempty"`      // 二级分类名称
	FatherID string `json:"father_id,omitempty"` // 对应一级分类 ID
}

// GetBaikeEntityListRespEntitieStatistics ...
type GetBaikeEntityListRespEntitieStatistics struct {
	LikeCount    int64 `json:"like_count,omitempty"`    // 累计点赞
	DislikeCount int64 `json:"dislike_count,omitempty"` // 当前词条版本收到的负反馈数量
}

// GetBaikeEntityListRespEntitieOuterInfo ...
type GetBaikeEntityListRespEntitieOuterInfo struct {
	Provider string `json:"provider,omitempty"` // 外部系统（不能包含中横线 "-"）
	OuterID  string `json:"outer_id,omitempty"` // 词条在外部系统中对应的唯一 ID（不能包含中横线 "-"）
}
