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

// A list of channel messages, usually a result of a list operation.
type ApiChannelMessageList struct {
	// A list of messages.
	Messages []ApiChannelMessage `json:"messages,omitempty"`
	LastSeenMessage *ApiChannelMessageHeader `json:"lastSeenMessage,omitempty"`
	LastSentMessage *ApiChannelMessageHeader `json:"lastSentMessage,omitempty"`
}
