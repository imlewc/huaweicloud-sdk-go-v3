/*
 * cloudide
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowPriceResponse struct {
	// 技术栈价格列表
	Prices *[]ResourcePrice `json:"prices,omitempty"`
	// 状态
	Status *string `json:"status,omitempty"`
}

func (o ShowPriceResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowPriceResponse", string(data)}, " ")
}
