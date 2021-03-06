/*
 * VPC
 *
 * VPC Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateSecurityGroupResponse struct {
	// 请求Id
	RequestId     *string            `json:"request_id,omitempty"`
	SecurityGroup *SecurityGroupInfo `json:"security_group,omitempty"`
}

func (o CreateSecurityGroupResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateSecurityGroupResponse", string(data)}, " ")
}
