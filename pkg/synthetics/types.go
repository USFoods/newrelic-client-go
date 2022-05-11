// Code generated by tutone: DO NOT EDIT
package synthetics

import (
	"github.com/newrelic/newrelic-client-go/pkg/nrtime"
)

// SyntheticsPrivateLocationMutationErrorType - Types of errors that can be returned from a Private Location mutation request
type SyntheticsPrivateLocationMutationErrorType string

var SyntheticsPrivateLocationMutationErrorTypeTypes = struct {
	// Received a request missing required fields or containing invalid data
	BAD_REQUEST SyntheticsPrivateLocationMutationErrorType
	// An unknown error occured while processing request to purge specified private location job queue
	INTERNAL_SERVER_ERROR SyntheticsPrivateLocationMutationErrorType
	// Private location not found for key (private location does not exist on account or has already been deleted)
	NOT_FOUND SyntheticsPrivateLocationMutationErrorType
	// User does not have authorization to purge job queue for specified private location
	UNAUTHORIZED SyntheticsPrivateLocationMutationErrorType
}{
	// Received a request missing required fields or containing invalid data
	BAD_REQUEST: "BAD_REQUEST",
	// An unknown error occured while processing request to purge specified private location job queue
	INTERNAL_SERVER_ERROR: "INTERNAL_SERVER_ERROR",
	// Private location not found for key (private location does not exist on account or has already been deleted)
	NOT_FOUND: "NOT_FOUND",
	// User does not have authorization to purge job queue for specified private location
	UNAUTHORIZED: "UNAUTHORIZED",
}

// SyntheticsError - Error object for Synthetics mutations
type SyntheticsError struct {
	// Description explaining the cause of the error
	Description string `json:"description,omitempty"`
}

// SyntheticsPrivateLocationDeleteResult - An array containing errors from the deletion of a private location, if any
type SyntheticsPrivateLocationDeleteResult struct {
	// An array container errors resulting from the mutation, if any
	Errors []SyntheticsPrivateLocationMutationError `json:"errors,omitempty"`
}

// SyntheticsPrivateLocationMutationError - Error object for Synthetic Private Location mutation request
type SyntheticsPrivateLocationMutationError struct {
	// String description of error
	Description string `json:"description"`
	// Enum type of error response
	Type SyntheticsPrivateLocationMutationErrorType `json:"type"`
}

// SyntheticsPrivateLocationMutationResult - Result of a private location mutation
type SyntheticsPrivateLocationMutationResult struct {
	// The account associated to the private location
	AccountID int `json:"accountId,omitempty"`
	// A description of the private location
	Description string `json:"description,omitempty"`
	// The private location globally unique identifier
	DomainId string `json:"domainId,omitempty"`
	// An array container errors resulting from the mutation, if any
	Errors []SyntheticsPrivateLocationMutationError `json:"errors,omitempty"`
	// The unique client identifier for the Synthetics private location in New Relic
	GUID EntityGUID `json:"guid,omitempty"`
	// The private locations key
	Key string `json:"key,omitempty"`
	// An alternate identifier based on name
	LocationId string `json:"locationId,omitempty"`
	// The name of the private location
	Name string `json:"name,omitempty"`
	// Specifies whether the private location requires a password for scripted monitors
	VerifiedScriptExecution bool `json:"verifiedScriptExecution,omitempty"`
}

// SyntheticsSecureCredentialMutationResult - The result of a secure credential mutation
type SyntheticsSecureCredentialMutationResult struct {
	// The moment when the secure credential was created, represented in milliseconds since the Unix epoch.
	CreatedAt nrtime.EpochMilliseconds `json:"createdAt,omitempty"`
	// Description of the secure credential, if available
	Description string `json:"description,omitempty"`
	// An array containing errors, if any
	Errors []SyntheticsError `json:"errors,omitempty"`
	// The unique identifier of the secure credential, if available
	Key string `json:"key,omitempty"`
	// The moment when the secure credential was last updated, represented in milliseconds since the Unix epoch.
	LastUpdate nrtime.EpochMilliseconds `json:"lastUpdate,omitempty"`
}

// EntityGUID - An encoded Entity GUID
type EntityGUID string

// SecureValue - The `SecureValue` scalar represents a secure value, ie a password, an API key, etc.
type SecureValue string
