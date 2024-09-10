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

// GetPerformanceReviewData 获取绩效结果
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/performance-v1/review_data/query
// new doc: https://open.feishu.cn/document/server-docs/performance-v1/query
func (r *PerformanceService) GetPerformanceReviewData(ctx context.Context, request *GetPerformanceReviewDataReq, options ...MethodOptionFunc) (*GetPerformanceReviewDataResp, *Response, error) {
	if r.cli.mock.mockPerformanceGetPerformanceReviewData != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Performance#GetPerformanceReviewData mock enable")
		return r.cli.mock.mockPerformanceGetPerformanceReviewData(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Performance",
		API:                   "GetPerformanceReviewData",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/performance/v1/review_datas/query",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getPerformanceReviewDataResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockPerformanceGetPerformanceReviewData mock PerformanceGetPerformanceReviewData method
func (r *Mock) MockPerformanceGetPerformanceReviewData(f func(ctx context.Context, request *GetPerformanceReviewDataReq, options ...MethodOptionFunc) (*GetPerformanceReviewDataResp, *Response, error)) {
	r.mockPerformanceGetPerformanceReviewData = f
}

// UnMockPerformanceGetPerformanceReviewData un-mock PerformanceGetPerformanceReviewData method
func (r *Mock) UnMockPerformanceGetPerformanceReviewData() {
	r.mockPerformanceGetPerformanceReviewData = nil
}

// GetPerformanceReviewDataReq ...
type GetPerformanceReviewDataReq struct {
	UserIDType         *IDType  `query:"user_id_type" json:"-"`          // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), people_admin_id: 以 people_admin_id 来识别用户, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	StartTime          string   `json:"start_time,omitempty"`            // 查询范围的开始日期, 毫秒级时间戳, 开始日期不能晚于截止日期, 示例值: "1430425599999"
	EndTime            string   `json:"end_time,omitempty"`              // 查询范围的截止日期, 毫秒级时间戳, 截止日期不能早于开始日期, 示例值: "1630425599999"
	StageTypes         []string `json:"stage_types,omitempty"`           // 环节类型。目前仅支持终评环节、结果沟通环节、查看绩效结果环节（不传默认包含所有的环节）, 示例值: ["leader_review"], 可选值有: leader_review: 终评环节, communication_and_open_result: 最终反馈环节, view_result: 查看绩效结果环节, 最大长度: `50`
	StageProgress      []int64  `json:"stage_progress,omitempty"`        // 查看绩效结果环节状态（不传默认包含所有的状态）, 可选值有: `0`: 已开通, 绩效结果已开通, 未发起复议也无需确认结果, `1`: 待确认, 绩效结果已开通但被评估人还未确认结果, 确认的截止时间还未到达, `2`: 已截止, 绩效结果已开通但被评估人还未确认结果, 确认的截止时间已到达, `3`: 已确认, 绩效结果已开通, 被评估人已确认结果, `4`: 已复议, 绩效结果已开通, 且被评估人已发起复议, 终评环节/结果沟通环节状态（不传默认包含所有的状态）, <!--, 示例值: [1], 可选值有: 0: 未开始, 任务的开始时间未到达, 1: 待完成, 任务的开始时间到达而截止时间未到达, 且任务未完成, 2: 已截止, 任务的截止时间已到达, 且任务未完成, 3: 已完成, 任务已完成, 最大长度: `50`
	SemesterIDList     []string `json:"semester_id_list,omitempty"`      // 评估周期 ID 列表, semester_id 是一个评估周期的唯一标识, 可以通过「我的评估」页面 url 获取, 也可通过本接口的返回值获取, 示例值: ["6992035450862224940"], 最大长度: `50`
	RevieweeUserIDList []string `json:"reviewee_user_id_list,omitempty"` // 被评估人 ID 列表, 示例值: ["ou_3245842393d09e9428ad4655da6e30b3"], 最大长度: `50`
	UpdatedLaterThan   *string  `json:"updated_later_than,omitempty"`    // 环节更新时间晚于, 可筛选出在此时间之后, 有内容提交的环节数据, 示例值: "1630425599999"
}

// GetPerformanceReviewDataResp ...
type GetPerformanceReviewDataResp struct {
	Semesters  []*GetPerformanceReviewDataRespSemester  `json:"semesters,omitempty"`  // 绩效评估周期列表
	Activities []*GetPerformanceReviewDataRespActivitie `json:"activities,omitempty"` // 绩效评估项目列表
	Indicators []*GetPerformanceReviewDataRespIndicator `json:"indicators,omitempty"` // 评估项列表
	Templates  []*GetPerformanceReviewDataRespTemplate  `json:"templates,omitempty"`  // 评估模板列表
	Units      []*GetPerformanceReviewDataRespUnit      `json:"units,omitempty"`      // 评估内容列表
	Fields     []*GetPerformanceReviewDataRespField     `json:"fields,omitempty"`     // 填写项列表
	Datas      []*GetPerformanceReviewDataRespData      `json:"datas,omitempty"`      // 评估数据列表
}

// GetPerformanceReviewDataRespActivitie ...
type GetPerformanceReviewDataRespActivitie struct {
	ID         string                                     `json:"id,omitempty"`          // 绩效评估项目 ID
	Name       *GetPerformanceReviewDataRespActivitieName `json:"name,omitempty"`        // 绩效评估项目名称
	SemesterID string                                     `json:"semester_id,omitempty"` // 绩效评估周期 ID
}

// GetPerformanceReviewDataRespActivitieName ...
type GetPerformanceReviewDataRespActivitieName struct {
	ZhCn string `json:"zh-CN,omitempty"` // 中文
	EnUs string `json:"en-US,omitempty"` // 英文
}

// GetPerformanceReviewDataRespData ...
type GetPerformanceReviewDataRespData struct {
	UserID     *GetPerformanceReviewDataRespDataUserID  `json:"user_id,omitempty"`     // 被评估人 ID
	SemesterID string                                   `json:"semester_id,omitempty"` // 绩效评估周期 ID
	ActivityID string                                   `json:"activity_id,omitempty"` // 绩效评估项目 ID
	Stages     []*GetPerformanceReviewDataRespDataStage `json:"stages,omitempty"`      // 本周期内各环节内容
}

// GetPerformanceReviewDataRespDataStage ...
type GetPerformanceReviewDataRespDataStage struct {
	StageType string                                       `json:"stage_type,omitempty"` // 环节类型
	Progress  int64                                        `json:"progress,omitempty"`   // 环节状态, 可选值有: 0: 未开始, 任务的开始时间未到达, 1: 待完成, 任务的开始时间到达而截止时间未到达, 且任务未完成, 2: 已截止, 任务的截止时间已到达, 且任务未完成, 3: 已完成, 任务已完成, 4: 已复议, 绩效结果已开通, 且被评估人已发起复议
	Data      []*GetPerformanceReviewDataRespDataStageData `json:"data,omitempty"`       // 环节填写内容
}

// GetPerformanceReviewDataRespDataStageData ...
type GetPerformanceReviewDataRespDataStageData struct {
	TemplateID     string                                                   `json:"template_id,omitempty"`      // 评估模板 ID
	UnitID         string                                                   `json:"unit_id,omitempty"`          // 评估内容 ID
	FieldID        string                                                   `json:"field_id,omitempty"`         // 评估控件 ID
	ReviewerUserID *GetPerformanceReviewDataRespDataStageDataReviewerUserID `json:"reviewer_user_id,omitempty"` // 评估人 ID
	SubmitTime     string                                                   `json:"submit_time,omitempty"`      // 最后提交时间
	IndicatorID    string                                                   `json:"indicator_id,omitempty"`     // 评估项 ID, option_id 或 score 有值的时候有值
	OptionID       string                                                   `json:"option_id,omitempty"`        // 评估项结果等级 ID, 有值表示当前数据是评级型评估项数据
	Score          string                                                   `json:"score,omitempty"`            // 评分型评估项填写内容, 有值表示当前数据是评分型评估项数据
	Text           string                                                   `json:"text,omitempty"`             // 填写项填写内容, 有值表示当前数据是填写项数据
}

// GetPerformanceReviewDataRespDataStageDataReviewerUserID ...
type GetPerformanceReviewDataRespDataStageDataReviewerUserID struct {
	OpenID string `json:"open_id,omitempty"` // 用户的 open_id
	UserID string `json:"user_id,omitempty"` // 用户的 user_id, 当 user_id_type 为 people_admin_id 时值为 saas_id, 否则为飞书的 user_id
}

// GetPerformanceReviewDataRespDataUserID ...
type GetPerformanceReviewDataRespDataUserID struct {
	OpenID string `json:"open_id,omitempty"` // 用户的 open_id
	UserID string `json:"user_id,omitempty"` // 用户的 user_id, 当 user_id_type 为 people_admin_id 时值为 saas_id, 否则为飞书的 user_id
}

// GetPerformanceReviewDataRespField ...
type GetPerformanceReviewDataRespField struct {
	ID            string                                 `json:"id,omitempty"`              // 填写项 ID
	Name          *GetPerformanceReviewDataRespFieldName `json:"name,omitempty"`            // 填写项名称
	IndicatorID   string                                 `json:"indicator_id,omitempty"`    // 评估项 ID
	ParentFieldID string                                 `json:"parent_field_id,omitempty"` // 父级填写项 ID
}

// GetPerformanceReviewDataRespFieldName ...
type GetPerformanceReviewDataRespFieldName struct {
	ZhCn string `json:"zh-CN,omitempty"` // 中文
	EnUs string `json:"en-US,omitempty"` // 英文
}

// GetPerformanceReviewDataRespIndicator ...
type GetPerformanceReviewDataRespIndicator struct {
	ID      string                                         `json:"id,omitempty"`      // 评估项 ID
	Name    *GetPerformanceReviewDataRespIndicatorName     `json:"name,omitempty"`    // 评估项名称
	Options []*GetPerformanceReviewDataRespIndicatorOption `json:"options,omitempty"` // 评估项等级列表
}

// GetPerformanceReviewDataRespIndicatorName ...
type GetPerformanceReviewDataRespIndicatorName struct {
	ZhCn string `json:"zh-CN,omitempty"` // 中文
	EnUs string `json:"en-US,omitempty"` // 英文
}

// GetPerformanceReviewDataRespIndicatorOption ...
type GetPerformanceReviewDataRespIndicatorOption struct {
	ID    string                                           `json:"id,omitempty"`    // 等级 ID
	Name  *GetPerformanceReviewDataRespIndicatorOptionName `json:"name,omitempty"`  // 等级名称
	Label string                                           `json:"label,omitempty"` // 等级代号
}

// GetPerformanceReviewDataRespIndicatorOptionName ...
type GetPerformanceReviewDataRespIndicatorOptionName struct {
	ZhCn string `json:"zh-CN,omitempty"` // 中文
	EnUs string `json:"en-US,omitempty"` // 英文
}

// GetPerformanceReviewDataRespSemester ...
type GetPerformanceReviewDataRespSemester struct {
	ID        string                                    `json:"id,omitempty"`         // 绩效评估周期 ID
	Name      *GetPerformanceReviewDataRespSemesterName `json:"name,omitempty"`       // 绩效评估周期名称
	StartTime string                                    `json:"start_time,omitempty"` // 绩效评估周期开始时间
	EndTime   string                                    `json:"end_time,omitempty"`   // 绩效评估周期结束时间
}

// GetPerformanceReviewDataRespSemesterName ...
type GetPerformanceReviewDataRespSemesterName struct {
	ZhCn string `json:"zh-CN,omitempty"` // 中文
	EnUs string `json:"en-US,omitempty"` // 英文
}

// GetPerformanceReviewDataRespTemplate ...
type GetPerformanceReviewDataRespTemplate struct {
	ID    string                                    `json:"id,omitempty"`    // 评估模板 ID
	Name  *GetPerformanceReviewDataRespTemplateName `json:"name,omitempty"`  // 评估模板所在环节名称
	Stage string                                    `json:"stage,omitempty"` // 评估模板环节类型
}

// GetPerformanceReviewDataRespTemplateName ...
type GetPerformanceReviewDataRespTemplateName struct {
	ZhCn string `json:"zh-CN,omitempty"` // 中文
	EnUs string `json:"en-US,omitempty"` // 英文
}

// GetPerformanceReviewDataRespUnit ...
type GetPerformanceReviewDataRespUnit struct {
	ID   string                                `json:"id,omitempty"`   // 评估内容 ID
	Name *GetPerformanceReviewDataRespUnitName `json:"name,omitempty"` // 评估内容名称
}

// GetPerformanceReviewDataRespUnitName ...
type GetPerformanceReviewDataRespUnitName struct {
	ZhCn string `json:"zh-CN,omitempty"` // 中文
	EnUs string `json:"en-US,omitempty"` // 英文
}

// getPerformanceReviewDataResp ...
type getPerformanceReviewDataResp struct {
	Code  int64                         `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                        `json:"msg,omitempty"`  // 错误描述
	Data  *GetPerformanceReviewDataResp `json:"data,omitempty"`
	Error *ErrorDetail                  `json:"error,omitempty"`
}