/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

type AdminWorkflowCreateRequest struct {
	Id   *CoreIdentifier    `json:"id,omitempty"`
	Spec *AdminWorkflowSpec `json:"spec,omitempty"`
}
