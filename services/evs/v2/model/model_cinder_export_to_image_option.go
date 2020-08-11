/*
    * EVS
    *
    * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
    *
*/

package model

// 将云硬盘导出为镜像的请求参数
type CinderExportToImageOption struct {
	// 云硬盘导出镜像的容器类型。  目前支持ami、ari、aki、ovf、bare。默认是bare。
	ContainerFormat string `json:"container_format,omitempty"`
	// 云硬盘导出镜像的格式。  目前支持vhd、zvhd、zvhd2、raw、qcow2。默认是vhd。
	DiskFormat string `json:"disk_format,omitempty"`
	// 强制导出镜像的标示，默认值是false。  当force标记为false时，云硬盘处于正在使用状态时，不能强制导出镜像。 当force标记为true时，即使云硬盘处于正在使用状态时，仍可以导出镜像。
	Force bool `json:"force,omitempty"`
	// 云硬盘导出镜像的名称。  名称的长度范围为1～128位。 名称只能包含以下字符：大写字母、小写字母、中文、数字、特殊字符包含“-”、“.”、“_”和空格。
	ImageName string `json:"image_name"`
	// 云硬盘导出镜像的系统类型。目前只支持“windows”和“linux”，默认值是“linux”。  说明： 只有云硬盘的volume_image_metadata信息中无“__os_type”字段且云硬盘状态为“available”时，设置的__os_type才会生效。 如果不传递该参数，则使用默认的“linux”值作为镜像的系统类型。
	OsType string `json:"__os_type,omitempty"`
}