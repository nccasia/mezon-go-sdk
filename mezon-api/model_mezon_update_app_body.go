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

// Update app information.
type MezonUpdateAppBody struct {
	// Username.
	Appname string `json:"appname,omitempty"`
	// Metadata.
	Metadata string `json:"metadata,omitempty"`
	// Avatar URL.
	Applogo string `json:"applogo,omitempty"`
	// Token.
	Token string `json:"token,omitempty"`
	// about the app.
	About string `json:"about,omitempty"`
}