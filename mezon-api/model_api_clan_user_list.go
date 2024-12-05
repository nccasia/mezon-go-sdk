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

// A list of users belonging to a clan, along with their role.
type ApiClanUserList struct {
	// User-role pairs for a clan.
	ClanUsers []ClanUserListClanUser `json:"clanUsers,omitempty"`
	// Cursor for the next page of results, if any.
	Cursor string `json:"cursor,omitempty"`
	ClanId string `json:"clanId,omitempty"`
}