/*
    * Cbr
    *
    * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
    *
*/

package model

// Request Object
type UpdateVaultRequest struct {
	VaultId string `json:"vault_id"`
	Body *VaultUpdateReq `json:"body,omitempty"`
}