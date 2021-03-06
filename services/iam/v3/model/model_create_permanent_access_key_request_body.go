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

//
type CreatePermanentAccessKeyRequestBody struct {
	Credential *CreateCredentialOption `json:"credential"`
}

func (o CreatePermanentAccessKeyRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreatePermanentAccessKeyRequestBody", string(data)}, " ")
}
