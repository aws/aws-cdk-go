package awsbedrockagentcorealpha


// The interception point where the interceptor will be invoked.
// Experimental.
type InterceptionPoint string

const (
	// Execute before the gateway makes a call to the target Useful for request validation, transformation, or custom authorization.
	// Experimental.
	InterceptionPoint_REQUEST InterceptionPoint = "REQUEST"
	// Execute after the target responds but before the gateway sends the response back Useful for response transformation, filtering, or adding custom headers.
	// Experimental.
	InterceptionPoint_RESPONSE InterceptionPoint = "RESPONSE"
)

