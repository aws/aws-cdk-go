package awsstepfunctionstasks


// The authentication method used to call the endpoint.
// Experimental.
type AuthType string

const (
	// Call the API direclty with no authorization method.
	// Experimental.
	AuthType_NO_AUTH AuthType = "NO_AUTH"
	// Use the IAM role associated with the current state machine for authorization.
	// Experimental.
	AuthType_IAM_ROLE AuthType = "IAM_ROLE"
	// Use the resource policy of the API for authorization.
	// Experimental.
	AuthType_RESOURCE_POLICY AuthType = "RESOURCE_POLICY"
)

