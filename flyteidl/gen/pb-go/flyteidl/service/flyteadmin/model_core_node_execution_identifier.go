/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Encapsulation of fields that identify a Flyte node execution entity.
type CoreNodeExecutionIdentifier struct {
	NodeId      string                           `json:"node_id,omitempty"`
	ExecutionId *CoreWorkflowExecutionIdentifier `json:"execution_id,omitempty"`
}
