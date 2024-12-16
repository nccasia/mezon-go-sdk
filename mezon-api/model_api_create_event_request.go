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

import (
	"time"
)

// Create a event within clan.
type ApiCreateEventRequest struct {
	Title string `json:"title,omitempty"`
	Logo string `json:"logo,omitempty"`
	Description string `json:"description,omitempty"`
	ClanId string `json:"clanId,omitempty"`
	ChannelId string `json:"channelId,omitempty"`
	Address string `json:"address,omitempty"`
	StartTime time.Time `json:"startTime,omitempty"`
	EndTime time.Time `json:"endTime,omitempty"`
	EventId string `json:"eventId,omitempty"`
	EventStatus string `json:"eventStatus,omitempty"`
}
