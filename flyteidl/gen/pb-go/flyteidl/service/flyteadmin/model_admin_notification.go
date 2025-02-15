/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Represents a structure for notifications based on execution status. The notification content is configured within flyte admin but can be templatized. Future iterations could expose configuring notifications with custom content.
type AdminNotification struct {
	Phases    []CoreWorkflowExecutionPhase `json:"phases,omitempty"`
	Email     *AdminEmailNotification      `json:"email,omitempty"`
	PagerDuty *AdminPagerDutyNotification  `json:"pager_duty,omitempty"`
	Slack     *AdminSlackNotification      `json:"slack,omitempty"`
}
