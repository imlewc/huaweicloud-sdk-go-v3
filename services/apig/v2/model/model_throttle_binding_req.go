/*
 * APIG
 *
 * API网关（API Gateway）是为开发者、合作伙伴提供的高性能、高可用、高安全的API托管服务，帮助用户轻松构建、管理和发布任意规模的API。
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type ThrottleBindingReq struct {
	// 流控策略编号
	StrategyId string `json:"strategy_id"`
	// API的发布记录编号
	PublishIds []string `json:"publish_ids"`
}

func (o ThrottleBindingReq) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ThrottleBindingReq", string(data)}, " ")
}
