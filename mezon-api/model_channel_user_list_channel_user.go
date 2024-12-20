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

// A single user-role pair.
type ChannelUserListChannelUser struct {
	// User.
	UserId string `json:"userId,omitempty"`
	// Their relationship to the role.
	RoleId []string `json:"roleId,omitempty"`
	Id string `json:"id,omitempty"`
	ThreadId string `json:"threadId,omitempty"`
	ClanNick string `json:"clanNick,omitempty"`
	ClanAvatar string `json:"clanAvatar,omitempty"`
	ClanId string `json:"clanId,omitempty"`
}
