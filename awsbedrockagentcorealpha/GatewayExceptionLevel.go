package awsbedrockagentcorealpha


// Exception levels for gateway.
// Experimental.
type GatewayExceptionLevel string

const (
	// Debug mode for granular exception messages.
	//
	// Allows the return of
	// specific error messages related to the gateway target configuration
	// to help you with debugging.
	// Experimental.
	GatewayExceptionLevel_DEBUG GatewayExceptionLevel = "DEBUG"
)

