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

type ApiCreateActivityRequest struct {
	ActivityName string `json:"activityName,omitempty"`
	ActivityType int32 `json:"activityType,omitempty"`
	ActivityDescription string `json:"activityDescription,omitempty"`
	StartTime time.Time `json:"startTime,omitempty"`
	ApplicationId string `json:"applicationId,omitempty"`
	Status int32 `json:"status,omitempty"`
}
