/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// WorkflowMetadataOnFailurePolicy : - FAIL_IMMEDIATELY: FAIL_IMMEDIATELY instructs the system to fail as soon as a node fails in the workflow. It'll automatically abort all currently running nodes and clean up resources before finally marking the workflow executions as failed.  - FAIL_AFTER_EXECUTABLE_NODES_COMPLETE: FAIL_AFTER_EXECUTABLE_NODES_COMPLETE instructs the system to make as much progress as it can. The system will not alter the dependencies of the execution graph so any node that depend on the failed node will not be run. Other nodes that will be executed to completion before cleaning up resources and marking the workflow execution as failed.
type WorkflowMetadataOnFailurePolicy string

// List of WorkflowMetadataOnFailurePolicy
const (
	WorkflowMetadataOnFailurePolicyIMMEDIATELY                     WorkflowMetadataOnFailurePolicy = "FAIL_IMMEDIATELY"
	WorkflowMetadataOnFailurePolicyAFTER_EXECUTABLE_NODES_COMPLETE WorkflowMetadataOnFailurePolicy = "FAIL_AFTER_EXECUTABLE_NODES_COMPLETE"
)
