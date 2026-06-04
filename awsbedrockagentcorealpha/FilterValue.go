package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A typed filter value for online evaluation filtering.
//
// Use the static factory methods to create filter values:
// - `FilterValue.string()` for string comparisons
// - `FilterValue.number()` for numeric comparisons
// - `FilterValue.boolean()` for boolean comparisons
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   filterValue := bedrock_agentcore_alpha.FilterValue_Boolean(jsii.Boolean(false))
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type FilterValue interface {
}

// The jsii proxy struct for FilterValue
type jsiiProxy_FilterValue struct {
	_ byte // padding
}

// Creates a boolean filter value.
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func FilterValue_Boolean(value *bool) FilterValue {
	_init_.Initialize()

	if err := validateFilterValue_BooleanParameters(value); err != nil {
		panic(err)
	}
	var returns FilterValue

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.FilterValue",
		"boolean",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Creates a numeric filter value.
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func FilterValue_Number(value *float64) FilterValue {
	_init_.Initialize()

	if err := validateFilterValue_NumberParameters(value); err != nil {
		panic(err)
	}
	var returns FilterValue

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.FilterValue",
		"number",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Creates a string filter value.
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func FilterValue_String(value *string) FilterValue {
	_init_.Initialize()

	if err := validateFilterValue_StringParameters(value); err != nil {
		panic(err)
	}
	var returns FilterValue

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.FilterValue",
		"string",
		[]interface{}{value},
		&returns,
	)

	return returns
}

