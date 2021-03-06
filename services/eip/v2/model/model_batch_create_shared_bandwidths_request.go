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

// Request Object
type BatchCreateSharedBandwidthsRequest struct {
	Body *BatchCreateBandwidthRequestBody `json:"body,omitempty"`
}

func (o BatchCreateSharedBandwidthsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchCreateSharedBandwidthsRequest", string(data)}, " ")
}
