/*
 * kms
 *
 * KMS v1.0 API, open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CancelSelfGrantRequest struct {
	VersionId string                  `json:"version_id"`
	Body      *RevokeGrantRequestBody `json:"body,omitempty"`
}

func (o CancelSelfGrantRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CancelSelfGrantRequest", string(data)}, " ")
}
