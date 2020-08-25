/*
 * CES
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

//
type Datapoint struct {
	// 指标值，该字段名称与请求参数中filter使用的查询值相同。
	Average float64 `json:"average"`
	// 指标采集时间。
	Timestamp int64 `json:"timestamp"`
	// 指标单位
	Unit *string `json:"unit,omitempty"`
}

func (o Datapoint) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"Datapoint", string(data)}, " ")
}