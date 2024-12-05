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

// Send Apple's Game Center account credentials to the server. Used with authenticate/link/unlink.  https://developer.apple.com/documentation/gamekit/gklocalplayer/1515407-generateidentityverificationsign
type ApiAccountGameCenter struct {
	// Player ID (generated by GameCenter).
	PlayerId string `json:"playerId,omitempty"`
	// Bundle ID (generated by GameCenter).
	BundleId string `json:"bundleId,omitempty"`
	// Time since UNIX epoch when the signature was created.
	TimestampSeconds string `json:"timestampSeconds,omitempty"`
	// A random \"NSString\" used to compute the hash and keep it randomized.
	Salt string `json:"salt,omitempty"`
	// The verification signature data generated.
	Signature string `json:"signature,omitempty"`
	// The URL for the public encryption key.
	PublicKeyUrl string `json:"publicKeyUrl,omitempty"`
	// Extra information that will be bundled in the session token.
	Vars map[string]string `json:"vars,omitempty"`
}