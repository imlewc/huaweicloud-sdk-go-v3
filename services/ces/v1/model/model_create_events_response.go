/*
 * CES
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
type CreateEventsResponse struct {
	Body *[]CreateEventsResponseBody `json:"body,omitempty"`
}

func (o CreateEventsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateEventsResponse", string(data)}, " ")
}
