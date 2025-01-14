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

// Get clan profile.
type ApiClanProfile struct {
	UserId string `json:"userId,omitempty"`
	NickName string `json:"nickName,omitempty"`
	Avatar string `json:"avatar,omitempty"`
	ClanId string `json:"clanId,omitempty"`
}
