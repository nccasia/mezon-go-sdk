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

type ApiPinMessageRequest struct {
	MessageId string `json:"messageId,omitempty"`
	ChannelId string `json:"channelId,omitempty"`
	ClanId string `json:"clanId,omitempty"`
}