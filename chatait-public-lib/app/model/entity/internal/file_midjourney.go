// Copyright 2023 Anlity <leo@leocode.net>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

// FileMidjourney is the golang structure for table c_file_midjourney.
type FileMidjourney struct {
	Id         int64  `orm:"id,primary"   json:"id"`         // id
	UserId     int64  `orm:"user_id"      json:"userId"`     // 会员id
	QueueId    int64  `orm:"queue_id"     json:"queueId"`    // 生成该图片的队列id
	FileName   string `orm:"file_name"    json:"fileName"`   // 文件名
	Path       string `orm:"path"         json:"path"`       // 本地储存路径
	Prompt     string `orm:"prompt"       json:"prompt"`     // 生成该图片的提示词
	MjFileName string `orm:"mj_file_name" json:"mjFileName"` // midjourney的文件名
	MjUrl      string `orm:"mj_url"       json:"mjUrl"`      // midjourney的路径
	Width      int    `orm:"width"        json:"width"`      // 宽
	Height     int    `orm:"height"       json:"height"`     // 高
	Size       int    `orm:"size"         json:"size"`       // 大小
	CreatedAt  int    `orm:"created_at"   json:"createdAt"`  // 创建时间
}