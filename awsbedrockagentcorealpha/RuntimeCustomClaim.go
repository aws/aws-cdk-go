package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a custom claim validation configuration for Runtime JWT authorizers.
//
// Custom claims allow you to validate additional fields in JWT tokens beyond
// the standard audience, client, and scope validations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   runtimeCustomClaim := bedrock_agentcore_alpha.RuntimeCustomClaim_WithStringArrayValue(jsii.String("name"), []*string{
//   	jsii.String("values"),
//   }, bedrock_agentcore_alpha.CustomClaimOperator_EQUALS)
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type RuntimeCustomClaim interface {
}

// The jsii proxy struct for RuntimeCustomClaim
type jsiiProxy_RuntimeCustomClaim struct {
	_ byte // padding
}

// Create a custom claim with a string array value.
//
// String array claims can use CONTAINS (default) or CONTAINS_ANY operator.
//
// Returns: A RuntimeCustomClaim configured for string array validation.
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func RuntimeCustomClaim_WithStringArrayValue(name *string, values *[]*string, operator CustomClaimOperator) RuntimeCustomClaim {
	_init_.Initialize()

	if err := validateRuntimeCustomClaim_WithStringArrayValueParameters(name, values); err != nil {
		panic(err)
	}
	var returns RuntimeCustomClaim

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.RuntimeCustomClaim",
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
// Returns: A RuntimeCustomClaim configured for string validation.
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func RuntimeCustomClaim_WithStringValue(name *string, value *string) RuntimeCustomClaim {
	_init_.Initialize()

	if err := validateRuntimeCustomClaim_WithStringValueParameters(name, value); err != nil {
		panic(err)
	}
	var returns RuntimeCustomClaim

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.RuntimeCustomClaim",
		"withStringValue",
		[]interface{}{name, value},
		&returns,
	)

	return returns
}

