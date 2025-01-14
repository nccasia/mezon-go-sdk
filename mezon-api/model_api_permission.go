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

type ApiPermission struct {
	Id string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Slug string `json:"slug,omitempty"`
	Description string `json:"description,omitempty"`
	Active int32 `json:"active,omitempty"`
	Scope int32 `json:"scope,omitempty"`
	Level int32 `json:"level,omitempty"`
}
