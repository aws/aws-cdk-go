package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Builder for condition expressions in Cedar policies.
//
// Conditions define when a policy statement should apply or not apply.
// Supports logical operators (AND, OR) and various comparison operators.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   conditionBuilder := bedrock_agentcore_alpha.NewConditionBuilder()
//
// Experimental.
type ConditionBuilder interface {
	// Logical AND operator - all conditions must be true.
	// Experimental.
	And() ConditionBuilder
	// Access a context attribute for comparison.
	//
	// Context attributes come from the request environment.
	// Common attributes: sourceIp, timestamp, environment, region, etc.
	// Experimental.
	ContextAttribute(attribute *string) AttributeAccessor
	// Logical OR operator - at least one condition must be true.
	// Experimental.
	Or() ConditionBuilder
	// Access a principal attribute for comparison.
	//
	// Principal attributes come from the authenticated user/service making the request.
	// Common attributes: username, role, department, groups, etc.
	// Experimental.
	PrincipalAttribute(attribute *string) AttributeAccessor
	// Access a resource attribute for comparison.
	//
	// Resource attributes come from the target resource being accessed.
	// Common attributes: arn, type, tags, owner, etc.
	// Experimental.
	ResourceAttribute(attribute *string) AttributeAccessor
}

// The jsii proxy struct for ConditionBuilder
type jsiiProxy_ConditionBuilder struct {
	_ byte // padding
}

// Experimental.
func NewConditionBuilder() ConditionBuilder {
	_init_.Initialize()

	j := jsiiProxy_ConditionBuilder{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ConditionBuilder",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewConditionBuilder_Override(c ConditionBuilder) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ConditionBuilder",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_ConditionBuilder) And() ConditionBuilder {
	var returns ConditionBuilder

	_jsii_.Invoke(
		c,
		"and",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionBuilder) ContextAttribute(attribute *string) AttributeAccessor {
	if err := c.validateContextAttributeParameters(attribute); err != nil {
		panic(err)
	}
	var returns AttributeAccessor

	_jsii_.Invoke(
		c,
		"contextAttribute",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionBuilder) Or() ConditionBuilder {
	var returns ConditionBuilder

	_jsii_.Invoke(
		c,
		"or",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionBuilder) PrincipalAttribute(attribute *string) AttributeAccessor {
	if err := c.validatePrincipalAttributeParameters(attribute); err != nil {
		panic(err)
	}
	var returns AttributeAccessor

	_jsii_.Invoke(
		c,
		"principalAttribute",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionBuilder) ResourceAttribute(attribute *string) AttributeAccessor {
	if err := c.validateResourceAttributeParameters(attribute); err != nil {
		panic(err)
	}
	var returns AttributeAccessor

	_jsii_.Invoke(
		c,
		"resourceAttribute",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

