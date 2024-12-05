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

type ApiOnboardingSteps struct {
	// id.
	Id string `json:"id,omitempty"`
	// user id.
	UserId string `json:"userId,omitempty"`
	// clan id.
	ClanId string `json:"clanId,omitempty"`
	// onboarding step.
	OnboardingStep int32 `json:"onboardingStep,omitempty"`
}
