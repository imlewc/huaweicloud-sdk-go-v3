/*
 * IAM
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
type RemoveProjectPermissionFromAgencyResponse struct {
}

func (o RemoveProjectPermissionFromAgencyResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"RemoveProjectPermissionFromAgencyResponse", string(data)}, " ")
}
