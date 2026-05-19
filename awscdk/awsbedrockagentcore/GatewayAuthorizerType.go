package awsbedrockagentcore


// Gateway authorizer type.
type GatewayAuthorizerType string

const (
	// Custom JWT authorizer type.
	GatewayAuthorizerType_CUSTOM_JWT GatewayAuthorizerType = "CUSTOM_JWT"
	// AWS IAM authorizer type.
	GatewayAuthorizerType_AWS_IAM GatewayAuthorizerType = "AWS_IAM"
	// No authorization type.
	GatewayAuthorizerType_NONE GatewayAuthorizerType = "NONE"
)

