// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

// ConversationMidjourney is the golang structure for table c_conversation_midjourney.
type ConversationMidjourney struct {
	ConversationId int64  `orm:"conversation_id,primary" json:"conversationId"` // 对话id
	ActionType     int    `orm:"action_type"             json:"actionType"`     // 行为类型 1生图 2Upsale 3Variate 4Reroll
	FileId         int64  `orm:"file_id"                 json:"fileId"`         // 图片文件ID
	Components     string `orm:"components"              json:"components"`     // 附加组件json 用于u,v,r等按钮及记录
	ErrorData      string `orm:"error_data"              json:"errorData"`      // 错误信息
}