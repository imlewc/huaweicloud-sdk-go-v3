/*
    * IAM
    *
    * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
    *
*/

package model

// 
type Version struct {
	// 版本状态。
	Status string `json:"status"`
	// 最后更新时间。
	Updated string `json:"updated"`
	// 版本的资源链接信息。
	Links []VersionLinks `json:"links"`
	// 版本号，如v3.6。
	Id string `json:"id"`
	// 支持的消息格式。
	MediaTypes []VersionMediatypes `json:"media-types"`
}