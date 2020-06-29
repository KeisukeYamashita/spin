/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type EnvironmentArtifactPin struct {

	Comment string `json:"comment,omitempty"`

	PinnedBy string `json:"pinnedBy,omitempty"`

	Reference string `json:"reference,omitempty"`

	TargetEnvironment string `json:"targetEnvironment,omitempty"`

	Version string `json:"version,omitempty"`
}
