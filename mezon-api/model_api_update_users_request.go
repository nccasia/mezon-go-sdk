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

// Fetch a batch of zero or more users from the server.
type ApiUpdateUsersRequest struct {
	// The account username of a user.
	DisplayName string `json:"displayName,omitempty"`
	// The avarar_url of a user.
	AvatarUrl string `json:"avatarUrl,omitempty"`
}
