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

// Update fields in a given channel.
type ApiChangeChannelPrivateRequest struct {
	// The ID of the channel to update.
	ChannelId string `json:"channelId,omitempty"`
	ChannelPrivate int32 `json:"channelPrivate,omitempty"`
	// The users to add.
	UserIds []string `json:"userIds,omitempty"`
	RoleIds []string `json:"roleIds,omitempty"`
}
