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

//
type CreatePortRequestBody struct {
	Port *CreatePortOption `json:"port"`
}

func (o CreatePortRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreatePortRequestBody", string(data)}, " ")
}
