/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

type AdminExecutionUpdateRequest struct {
	Id    *CoreWorkflowExecutionIdentifier `json:"id,omitempty"`
	State *AdminExecutionState             `json:"state,omitempty"`
}
