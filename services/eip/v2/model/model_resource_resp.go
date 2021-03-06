/*
 * EIP
 *
 * 云服务接口
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type ResourceResp struct {
	// 资源配额对象
	Resources []QuotaShowResp `json:"resources"`
}

func (o ResourceResp) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ResourceResp", string(data)}, " ")
}
