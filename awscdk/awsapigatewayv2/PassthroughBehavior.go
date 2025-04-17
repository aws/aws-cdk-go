package awsapigatewayv2


// Integration Passthrough Behavior.
type PassthroughBehavior string

const (
	// Passes the request body for unmapped content types through to the integration back end without transformation.
	PassthroughBehavior_WHEN_NO_MATCH PassthroughBehavior = "WHEN_NO_MATCH"
	// Rejects unmapped content types with an HTTP 415 'Unsupported Media Type' response.
	PassthroughBehavior_NEVER PassthroughBehavior = "NEVER"
	// Allows pass-through when the integration has NO content types mapped to templates.
	//
	// However if there is at least one content type defined,
	// unmapped content types will be rejected with the same 415 response.
	PassthroughBehavior_WHEN_NO_TEMPLATES PassthroughBehavior = "WHEN_NO_TEMPLATES"
)

