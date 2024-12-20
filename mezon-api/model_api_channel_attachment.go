/*
 * Mezon API v2
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.0
 * Contact: hello@heroiclabs.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

type ApiChannelAttachment struct {
	Id string `json:"id,omitempty"`
	Filename string `json:"filename,omitempty"`
	Filetype string `json:"filetype,omitempty"`
	Filesize string `json:"filesize,omitempty"`
	Url string `json:"url,omitempty"`
	Uploader string `json:"uploader,omitempty"`
	// The UNIX time (for gRPC clients) or ISO string (for REST clients) when the group was created.
	CreateTime time.Time `json:"createTime,omitempty"`
}
