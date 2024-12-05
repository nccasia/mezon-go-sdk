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

type ApiChannelSettingListResponse struct {
	ClanId string `json:"clanId,omitempty"`
	ChannelCount int32 `json:"channelCount,omitempty"`
	ThreadCount int32 `json:"threadCount,omitempty"`
	ChannelSettingList []ApiChannelSettingItem `json:"channelSettingList,omitempty"`
}