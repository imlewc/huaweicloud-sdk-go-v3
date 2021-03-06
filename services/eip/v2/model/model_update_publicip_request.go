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
type UpdatePublicipRequest struct {
	PublicipId string                      `json:"publicip_id"`
	Body       *UpdatePublicipsRequestBody `json:"body,omitempty"`
}

func (o UpdatePublicipRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdatePublicipRequest", string(data)}, " ")
}
