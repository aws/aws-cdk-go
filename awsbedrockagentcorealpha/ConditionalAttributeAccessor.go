package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Accessor for building type-safe attribute comparisons within conditional statements.
//
// Returns ConditionalPolicyStatement to allow chaining back to policy building.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   var conditionalPolicyStatement ConditionalPolicyStatement
//   var conditionBuilder ConditionBuilder
//
//   conditionalAttributeAccessor := bedrock_agentcore_alpha.NewConditionalAttributeAccessor(jsii.String("path"), conditionalPolicyStatement, conditionBuilder)
//
// Experimental.
type ConditionalAttributeAccessor interface {
	// String contains check.
	// Experimental.
	Contains(value *string) ConditionalPolicyStatement
	// Equality comparison (==).
	// Experimental.
	EqualTo(value interface{}) ConditionalPolicyStatement
	// Greater than comparison (>).
	// Experimental.
	GreaterThan(value *float64) ConditionalPolicyStatement
	// Greater than or equals comparison (>=).
	// Experimental.
	GreaterThanOrEqualTo(value *float64) ConditionalPolicyStatement
	// Check if attribute is in a set/list.
	// Experimental.
	IsIn(values *[]interface{}) ConditionalPolicyStatement
	// IP range check - tests if IP address is in CIDR range.
	// Experimental.
	IsInRange(ipRange *string) ConditionalPolicyStatement
	// Less than comparison (<).
	// Experimental.
	LessThan(value *float64) ConditionalPolicyStatement
	// Less than or equals comparison (<=).
	// Experimental.
	LessThanOrEqualTo(value *float64) ConditionalPolicyStatement
	// Inequality comparison (!=).
	// Experimental.
	NotEqualTo(value interface{}) ConditionalPolicyStatement
}

// The jsii proxy struct for ConditionalAttributeAccessor
type jsiiProxy_ConditionalAttributeAccessor struct {
	_ byte // padding
}

// Experimental.
func NewConditionalAttributeAccessor(path *string, parent ConditionalPolicyStatement, conditionBuilder ConditionBuilder) ConditionalAttributeAccessor {
	_init_.Initialize()

	if err := validateNewConditionalAttributeAccessorParameters(path, parent, conditionBuilder); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConditionalAttributeAccessor{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ConditionalAttributeAccessor",
		[]interface{}{path, parent, conditionBuilder},
		&j,
	)

	return &j
}

// Experimental.
func NewConditionalAttributeAccessor_Override(c ConditionalAttributeAccessor, path *string, parent ConditionalPolicyStatement, conditionBuilder ConditionBuilder) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ConditionalAttributeAccessor",
		[]interface{}{path, parent, conditionBuilder},
		c,
	)
}

func (c *jsiiProxy_ConditionalAttributeAccessor) Contains(value *string) ConditionalPolicyStatement {
	if err := c.validateContainsParameters(value); err != nil {
		panic(err)
	}
	var returns ConditionalPolicyStatement

	_jsii_.Invoke(
		c,
		"contains",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalAttributeAccessor) EqualTo(value interface{}) ConditionalPolicyStatement {
	if err := c.validateEqualToParameters(value); err != nil {
		panic(err)
	}
	var returns ConditionalPolicyStatement

	_jsii_.Invoke(
		c,
		"equalTo",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalAttributeAccessor) GreaterThan(value *float64) ConditionalPolicyStatement {
	if err := c.validateGreaterThanParameters(value); err != nil {
		panic(err)
	}
	var returns ConditionalPolicyStatement

	_jsii_.Invoke(
		c,
		"greaterThan",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalAttributeAccessor) GreaterThanOrEqualTo(value *float64) ConditionalPolicyStatement {
	if err := c.validateGreaterThanOrEqualToParameters(value); err != nil {
		panic(err)
	}
	var returns ConditionalPolicyStatement

	_jsii_.Invoke(
		c,
		"greaterThanOrEqualTo",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalAttributeAccessor) IsIn(values *[]interface{}) ConditionalPolicyStatement {
	if err := c.validateIsInParameters(values); err != nil {
		panic(err)
	}
	var returns ConditionalPolicyStatement

	_jsii_.Invoke(
		c,
		"isIn",
		[]interface{}{values},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalAttributeAccessor) IsInRange(ipRange *string) ConditionalPolicyStatement {
	if err := c.validateIsInRangeParameters(ipRange); err != nil {
		panic(err)
	}
	var returns ConditionalPolicyStatement

	_jsii_.Invoke(
		c,
		"isInRange",
		[]interface{}{ipRange},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalAttributeAccessor) LessThan(value *float64) ConditionalPolicyStatement {
	if err := c.validateLessThanParameters(value); err != nil {
		panic(err)
	}
	var returns ConditionalPolicyStatement

	_jsii_.Invoke(
		c,
		"lessThan",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalAttributeAccessor) LessThanOrEqualTo(value *float64) ConditionalPolicyStatement {
	if err := c.validateLessThanOrEqualToParameters(value); err != nil {
		panic(err)
	}
	var returns ConditionalPolicyStatement

	_jsii_.Invoke(
		c,
		"lessThanOrEqualTo",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalAttributeAccessor) NotEqualTo(value interface{}) ConditionalPolicyStatement {
	if err := c.validateNotEqualToParameters(value); err != nil {
		panic(err)
	}
	var returns ConditionalPolicyStatement

	_jsii_.Invoke(
		c,
		"notEqualTo",
		[]interface{}{value},
		&returns,
	)

	return returns
}

