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

type ApiWebhookCreateRequest struct {
	WebhookName string `json:"webhookName,omitempty"`
	ChannelId string `json:"channelId,omitempty"`
	Avatar string `json:"avatar,omitempty"`
	ClanId string `json:"clanId,omitempty"`
}
