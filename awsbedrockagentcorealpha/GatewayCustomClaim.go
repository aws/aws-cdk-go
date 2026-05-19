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
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   gatewayCustomClaim := bedrock_agentcore_alpha.GatewayCustomClaim_WithStringArrayValue(jsii.String("name"), []*string{
//   	jsii.String("values"),
//   }, bedrock_agentcore_alpha.CustomClaimOperator_EQUALS)
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

