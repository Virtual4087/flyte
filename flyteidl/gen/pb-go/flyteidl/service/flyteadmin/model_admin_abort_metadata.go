/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Specifies metadata around an aborted workflow execution.
type AdminAbortMetadata struct {
	// In the case of a user-specified abort, this will pass along the user-supplied cause.
	Cause     string `json:"cause,omitempty"`
	Principal string `json:"principal,omitempty"`
}
