package awsstepfunctionstasks


// The authentication method used to call the endpoint.
type AuthType string

const (
	// Call the API direclty with no authorization method.
	AuthType_NO_AUTH AuthType = "NO_AUTH"
	// Use the IAM role associated with the current state machine for authorization.
	AuthType_IAM_ROLE AuthType = "IAM_ROLE"
	// Use the resource policy of the API for authorization.
	AuthType_RESOURCE_POLICY AuthType = "RESOURCE_POLICY"
)

