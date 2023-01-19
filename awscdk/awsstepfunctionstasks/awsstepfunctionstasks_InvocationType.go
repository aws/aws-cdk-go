package awsstepfunctionstasks


// Invocation type of a Lambda.
// Deprecated: use `LambdaInvocationType`.
type InvocationType string

const (
	// Invoke synchronously.
	//
	// The API response includes the function response and additional data.
	// Deprecated: use `LambdaInvocationType`.
	InvocationType_REQUEST_RESPONSE InvocationType = "REQUEST_RESPONSE"
	// Invoke asynchronously.
	//
	// Send events that fail multiple times to the function's dead-letter queue (if it's configured).
	// The API response only includes a status code.
	// Deprecated: use `LambdaInvocationType`.
	InvocationType_EVENT InvocationType = "EVENT"
	// TValidate parameter values and verify that the user or role has permission to invoke the function.
	// Deprecated: use `LambdaInvocationType`.
	InvocationType_DRY_RUN InvocationType = "DRY_RUN"
)

