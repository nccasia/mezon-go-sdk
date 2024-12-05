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

type ApiUpdateRoleChannelRequest struct {
	// The ID of the role to update.
	RoleId string `json:"roleId,omitempty"`
	// The permissions to add.
	PermissionUpdate []ApiPermissionUpdate `json:"permissionUpdate,omitempty"`
	MaxPermissionId string `json:"maxPermissionId,omitempty"`
	ChannelId string `json:"channelId,omitempty"`
	// The ID of the role to update.
	UserId string `json:"userId,omitempty"`
}