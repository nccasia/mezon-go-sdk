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

type ApiUserStatusUpdate struct {
	Status string `json:"status,omitempty"`
	Minutes int32 `json:"minutes,omitempty"`
	UntilTurnOn bool `json:"untilTurnOn,omitempty"`
}