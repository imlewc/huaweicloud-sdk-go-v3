/*
 * ELB
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type CreateL7RuleRequestBody struct {
	Rule *CreateRuleOption `json:"rule"`
}

func (o CreateL7RuleRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateL7RuleRequestBody", string(data)}, " ")
}
