package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Factory class for creating Gateway Authorizers.
// Experimental.
type GatewayAuthorizer interface {
}

// The jsii proxy struct for GatewayAuthorizer
type jsiiProxy_GatewayAuthorizer struct {
	_ byte // padding
}

// Experimental.
func NewGatewayAuthorizer_Override(g GatewayAuthorizer) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayAuthorizer",
		nil, // no parameters
		g,
	)
}

// AWS IAM authorizer instance.
// Experimental.
func GatewayAuthorizer_UsingAwsIam() IGatewayAuthorizerConfig {
	_init_.Initialize()

	var returns IGatewayAuthorizerConfig

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayAuthorizer",
		"usingAwsIam",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Create a JWT authorizer from Cognito User Pool.
//
// Returns: CustomJwtAuthorizer configured for Cognito.
// Experimental.
func GatewayAuthorizer_UsingCognito(props *CognitoAuthorizerProps) IGatewayAuthorizerConfig {
	_init_.Initialize()

	if err := validateGatewayAuthorizer_UsingCognitoParameters(props); err != nil {
		panic(err)
	}
	var returns IGatewayAuthorizerConfig

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayAuthorizer",
		"usingCognito",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create a custom JWT authorizer.
//
// Returns: IGatewayAuthorizerConfig configured for custom JWT.
// Experimental.
func GatewayAuthorizer_UsingCustomJwt(configuration *CustomJwtConfiguration) IGatewayAuthorizerConfig {
	_init_.Initialize()

	if err := validateGatewayAuthorizer_UsingCustomJwtParameters(configuration); err != nil {
		panic(err)
	}
	var returns IGatewayAuthorizerConfig

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayAuthorizer",
		"usingCustomJwt",
		[]interface{}{configuration},
		&returns,
	)

	return returns
}

// No authorization — the gateway will not perform any inbound authorization.
//
// The gateway endpoint will be publicly accessible without credentials.
// Use this for testing/development, or for production gateways where you have
// implemented compensating controls such as Gateway Interceptors.
//
// Returns: IGatewayAuthorizerConfig configured for no authorization.
// See: https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/gateway-inbound-auth.html#gateway-inbound-auth-none
//
// Experimental.
func GatewayAuthorizer_WithNoAuth() IGatewayAuthorizerConfig {
	_init_.Initialize()

	var returns IGatewayAuthorizerConfig

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayAuthorizer",
		"withNoAuth",
		nil, // no parameters
		&returns,
	)

	return returns
}

