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

// Send a Facebook Instant Game token to the server. Used with authenticate/link/unlink.
type ApiAccountFacebookInstantGame struct {
	SignedPlayerInfo string `json:"signedPlayerInfo,omitempty"`
	// Extra information that will be bundled in the session token.
	Vars map[string]string `json:"vars,omitempty"`
}