package awsbedrockagentcorealpha


// Gateway authorizer type.
// Experimental.
type GatewayAuthorizerType string

const (
	// Custom JWT authorizer type.
	// Experimental.
	GatewayAuthorizerType_CUSTOM_JWT GatewayAuthorizerType = "CUSTOM_JWT"
	// AWS IAM authorizer type.
	// Experimental.
	GatewayAuthorizerType_AWS_IAM GatewayAuthorizerType = "AWS_IAM"
)

