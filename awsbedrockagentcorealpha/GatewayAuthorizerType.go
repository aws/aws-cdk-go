package awsbedrockagentcorealpha


// Gateway authorizer type.
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type GatewayAuthorizerType string

const (
	// Custom JWT authorizer type.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	GatewayAuthorizerType_CUSTOM_JWT GatewayAuthorizerType = "CUSTOM_JWT"
	// AWS IAM authorizer type.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	GatewayAuthorizerType_AWS_IAM GatewayAuthorizerType = "AWS_IAM"
	// No authorization type.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	GatewayAuthorizerType_NONE GatewayAuthorizerType = "NONE"
)

