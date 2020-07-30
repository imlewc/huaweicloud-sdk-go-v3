/*
    * IAM
    *
    * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
    *
*/

package model

// 
type Catalog struct {
	// 终端节点信息。
	Endpoints []CatalogEndpoints `json:"endpoints"`
	// 服务ID。
	Id string `json:"id"`
	// 服务名。
	Name string `json:"name"`
	// 服务类型。
	Type string `json:"type"`
}