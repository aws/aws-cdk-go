package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Accessor for building type-safe attribute comparisons.
//
// Provides methods for common comparison operators with proper type checking.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   var conditionBuilder ConditionBuilder
//
//   attributeAccessor := bedrock_agentcore_alpha.NewAttributeAccessor(jsii.String("path"), conditionBuilder)
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type AttributeAccessor interface {
	// String contains check.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Contains(value *string) ConditionBuilder
	// Equality comparison (==).
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	EqualTo(value interface{}) ConditionBuilder
	// Greater than comparison (>).
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	GreaterThan(value *float64) ConditionBuilder
	// Greater than or equals comparison (>=).
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	GreaterThanOrEqualTo(value *float64) ConditionBuilder
	// Check if attribute is in a set/list.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	IsIn(values *[]interface{}) ConditionBuilder
	// IP range check - tests if IP address is in CIDR range.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	IsInRange(ipRange *string) ConditionBuilder
	// Less than comparison (<).
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	LessThan(value *float64) ConditionBuilder
	// Less than or equals comparison (<=).
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	LessThanOrEqualTo(value *float64) ConditionBuilder
	// Inequality comparison (!=).
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	NotEqualTo(value interface{}) ConditionBuilder
}

// The jsii proxy struct for AttributeAccessor
type jsiiProxy_AttributeAccessor struct {
	_ byte // padding
}

// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func NewAttributeAccessor(path *string, parent ConditionBuilder) AttributeAccessor {
	_init_.Initialize()

	if err := validateNewAttributeAccessorParameters(path, parent); err != nil {
		panic(err)
	}
	j := jsiiProxy_AttributeAccessor{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.AttributeAccessor",
		[]interface{}{path, parent},
		&j,
	)

	return &j
}

// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func NewAttributeAccessor_Override(a AttributeAccessor, path *string, parent ConditionBuilder) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.AttributeAccessor",
		[]interface{}{path, parent},
		a,
	)
}

func (a *jsiiProxy_AttributeAccessor) Contains(value *string) ConditionBuilder {
	if err := a.validateContainsParameters(value); err != nil {
		panic(err)
	}
	var returns ConditionBuilder

	_jsii_.Invoke(
		a,
		"contains",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttributeAccessor) EqualTo(value interface{}) ConditionBuilder {
	if err := a.validateEqualToParameters(value); err != nil {
		panic(err)
	}
	var returns ConditionBuilder

	_jsii_.Invoke(
		a,
		"equalTo",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttributeAccessor) GreaterThan(value *float64) ConditionBuilder {
	if err := a.validateGreaterThanParameters(value); err != nil {
		panic(err)
	}
	var returns ConditionBuilder

	_jsii_.Invoke(
		a,
		"greaterThan",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttributeAccessor) GreaterThanOrEqualTo(value *float64) ConditionBuilder {
	if err := a.validateGreaterThanOrEqualToParameters(value); err != nil {
		panic(err)
	}
	var returns ConditionBuilder

	_jsii_.Invoke(
		a,
		"greaterThanOrEqualTo",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttributeAccessor) IsIn(values *[]interface{}) ConditionBuilder {
	if err := a.validateIsInParameters(values); err != nil {
		panic(err)
	}
	var returns ConditionBuilder

	_jsii_.Invoke(
		a,
		"isIn",
		[]interface{}{values},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttributeAccessor) IsInRange(ipRange *string) ConditionBuilder {
	if err := a.validateIsInRangeParameters(ipRange); err != nil {
		panic(err)
	}
	var returns ConditionBuilder

	_jsii_.Invoke(
		a,
		"isInRange",
		[]interface{}{ipRange},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttributeAccessor) LessThan(value *float64) ConditionBuilder {
	if err := a.validateLessThanParameters(value); err != nil {
		panic(err)
	}
	var returns ConditionBuilder

	_jsii_.Invoke(
		a,
		"lessThan",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttributeAccessor) LessThanOrEqualTo(value *float64) ConditionBuilder {
	if err := a.validateLessThanOrEqualToParameters(value); err != nil {
		panic(err)
	}
	var returns ConditionBuilder

	_jsii_.Invoke(
		a,
		"lessThanOrEqualTo",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AttributeAccessor) NotEqualTo(value interface{}) ConditionBuilder {
	if err := a.validateNotEqualToParameters(value); err != nil {
		panic(err)
	}
	var returns ConditionBuilder

	_jsii_.Invoke(
		a,
		"notEqualTo",
		[]interface{}{value},
		&returns,
	)

	return returns
}

