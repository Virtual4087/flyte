/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Defines complete set of information required to trigger an execution on a schedule.
type AdminSchedule struct {
	CronExpression string             `json:"cron_expression,omitempty"`
	Rate           *AdminFixedRate    `json:"rate,omitempty"`
	CronSchedule   *AdminCronSchedule `json:"cron_schedule,omitempty"`
	// Name of the input variable that the kickoff time will be supplied to when the workflow is kicked off.
	KickoffTimeInputArg string `json:"kickoff_time_input_arg,omitempty"`
}
