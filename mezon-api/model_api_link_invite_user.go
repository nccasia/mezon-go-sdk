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

// Add link invite users to.
type ApiLinkInviteUser struct {
	ClanId string `json:"clanId,omitempty"`
	// The user to add.
	CreatorId string `json:"creatorId,omitempty"`
	ChannelId string `json:"channelId,omitempty"`
	InviteLink string `json:"inviteLink,omitempty"`
	CreateTime time.Time `json:"createTime,omitempty"`
	ExpiryTime time.Time `json:"expiryTime,omitempty"`
	Id string `json:"id,omitempty"`
}
