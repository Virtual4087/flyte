/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

import (
	"time"
)

type EventWorkflowExecutionEvent struct {
	ExecutionId *CoreWorkflowExecutionIdentifier `json:"execution_id,omitempty"`
	ProducerId  string                           `json:"producer_id,omitempty"`
	Phase       *CoreWorkflowExecutionPhase      `json:"phase,omitempty"`
	// This timestamp represents when the original event occurred, it is generated by the executor of the workflow.
	OccurredAt time.Time `json:"occurred_at,omitempty"`
	// URL to the output of the execution, it encodes all the information including Cloud source provider. ie., s3://...
	OutputUri string              `json:"output_uri,omitempty"`
	Error_    *CoreExecutionError `json:"error,omitempty"`
	// Raw output data produced by this workflow execution.
	OutputData *CoreLiteralMap `json:"output_data,omitempty"`
}
