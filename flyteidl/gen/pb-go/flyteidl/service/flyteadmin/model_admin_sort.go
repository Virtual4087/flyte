/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Specifies sort ordering in a list request.
type AdminSort struct {
	Key       string         `json:"key,omitempty"`
	Direction *SortDirection `json:"direction,omitempty"`
}
