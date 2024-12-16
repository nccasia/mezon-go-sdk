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

// A list of role description, usually a result of a list operation.
type ApiRoleList struct {
	// A list of role.
	Roles []ApiRole `json:"roles,omitempty"`
	// The cursor to send when retrieving the next page, if any.
	NextCursor string `json:"nextCursor,omitempty"`
	// The cursor to send when retrieving the previous page, if any.
	PrevCursor string `json:"prevCursor,omitempty"`
	// Cacheable cursor to list newer role description. Durable and designed to be stored, unlike next/prev cursors.
	CacheableCursor string `json:"cacheableCursor,omitempty"`
}
