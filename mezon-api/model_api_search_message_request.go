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

type ApiSearchMessageRequest struct {
	Filters []ApiFilterParam `json:"filters,omitempty"`
	From int32 `json:"from,omitempty"`
	Size int32 `json:"size,omitempty"`
	Sorts []ApiSortParam `json:"sorts,omitempty"`
}
