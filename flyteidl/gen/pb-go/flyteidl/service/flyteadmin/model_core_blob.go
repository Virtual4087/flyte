/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Refers to an offloaded set of files. It encapsulates the type of the store and a unique uri for where the data is. There are no restrictions on how the uri is formatted since it will depend on how to interact with the store.
type CoreBlob struct {
	Metadata *CoreBlobMetadata `json:"metadata,omitempty"`
	Uri      string            `json:"uri,omitempty"`
}
