package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a custom claim validation configuration for Gateway JWT authorizers.
//
// Custom claims allow you to validate additional fields in JWT tokens beyond
// the standard audience, client, and scope validations.
//
// Example:
//   // Optional: Create custom claims (CustomClaimOperator and GatewayCustomClaim from agentcore)
//   customClaims := []GatewayCustomClaim{
//   	agentcore.GatewayCustomClaim_WithStringValue(jsii.String("department"), jsii.String("engineering")),
//   	agentcore.GatewayCustomClaim_WithStringArrayValue(jsii.String("roles"), []*string{
//   		jsii.String("admin"),
//   	}, agentcore.CustomClaimOperator_CONTAINS),
//   	agentcore.GatewayCustomClaim_WithStringArrayValue(jsii.String("permissions"), []*string{
//   		jsii.String("read"),
//   		jsii.String("write"),
//   	}, agentcore.CustomClaimOperator_CONTAINS_ANY),
//   }
//
//   gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-gateway"),
//   	AuthorizerConfiguration: agentcore.GatewayAuthorizer_UsingCustomJwt(&CustomJwtConfiguration{
//   		DiscoveryUrl: jsii.String("https://auth.example.com/.well-known/openid-configuration"),
//   		AllowedAudience: []*string{
//   			jsii.String("my-app"),
//   		},
//   		AllowedClients: []*string{
//   			jsii.String("my-client-id"),
//   		},
//   		AllowedScopes: []*string{
//   			jsii.String("read"),
//   			jsii.String("write"),
//   		},
//   		CustomClaims: customClaims,
//   	}),
//   })
//
// Experimental.
type GatewayCustomClaim interface {
}

// The jsii proxy struct for GatewayCustomClaim
type jsiiProxy_GatewayCustomClaim struct {
	_ byte // padding
}

// Create a custom claim with a string array value.
//
// String array claims can use CONTAINS (default) or CONTAINS_ANY operator.
//
// Returns: A GatewayCustomClaim configured for string array validation.
// Experimental.
func GatewayCustomClaim_WithStringArrayValue(name *string, values *[]*string, operator CustomClaimOperator) GatewayCustomClaim {
	_init_.Initialize()

	if err := validateGatewayCustomClaim_WithStringArrayValueParameters(name, values); err != nil {
		panic(err)
	}
	var returns GatewayCustomClaim

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayCustomClaim",
		"withStringArrayValue",
		[]interface{}{name, values, operator},
		&returns,
	)

	return returns
}

// Create a custom claim with a string value.
//
// String claims must use the EQUALS operator.
//
// Returns: A GatewayCustomClaim configured for string validation.
// Experimental.
func GatewayCustomClaim_WithStringValue(name *string, value *string) GatewayCustomClaim {
	_init_.Initialize()

	if err := validateGatewayCustomClaim_WithStringValueParameters(name, value); err != nil {
		panic(err)
	}
	var returns GatewayCustomClaim

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.GatewayCustomClaim",
		"withStringValue",
		[]interface{}{name, value},
		&returns,
	)

	return returns
}

