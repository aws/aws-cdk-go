package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Wrapper class for conditionally building policy statements.
//
// This class allows chaining condition methods and returning to the parent
// PolicyStatement when done. It proxies condition building methods from
// ConditionBuilder.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   var conditionBuilder ConditionBuilder
//   var policyStatement PolicyStatement
//
//   conditionalPolicyStatement := bedrock_agentcore_alpha.NewConditionalPolicyStatement(policyStatement, conditionBuilder)
//
// Experimental.
type ConditionalPolicyStatement interface {
	// Logical AND operator - all conditions must be true.
	// Experimental.
	And() ConditionalPolicyStatement
	// Access a context attribute for comparison.
	// Experimental.
	ContextAttribute(attribute *string) ConditionalAttributeAccessor
	// Complete condition building and return to the PolicyStatement.
	//
	// Use this to finish building when/unless conditions and continue
	// configuring the policy statement.
	// Experimental.
	Done() PolicyStatement
	// Logical OR operator - at least one condition must be true.
	// Experimental.
	Or() ConditionalPolicyStatement
	// Access a principal attribute for comparison.
	// Experimental.
	PrincipalAttribute(attribute *string) ConditionalAttributeAccessor
	// Access a resource attribute for comparison.
	// Experimental.
	ResourceAttribute(attribute *string) ConditionalAttributeAccessor
	// Alias for done() to support fluent unless() chaining.
	// Experimental.
	Unless() ConditionalPolicyStatement
}

// The jsii proxy struct for ConditionalPolicyStatement
type jsiiProxy_ConditionalPolicyStatement struct {
	_ byte // padding
}

// Experimental.
func NewConditionalPolicyStatement(policyStatement PolicyStatement, conditionBuilder ConditionBuilder) ConditionalPolicyStatement {
	_init_.Initialize()

	if err := validateNewConditionalPolicyStatementParameters(policyStatement, conditionBuilder); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConditionalPolicyStatement{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ConditionalPolicyStatement",
		[]interface{}{policyStatement, conditionBuilder},
		&j,
	)

	return &j
}

// Experimental.
func NewConditionalPolicyStatement_Override(c ConditionalPolicyStatement, policyStatement PolicyStatement, conditionBuilder ConditionBuilder) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ConditionalPolicyStatement",
		[]interface{}{policyStatement, conditionBuilder},
		c,
	)
}

func (c *jsiiProxy_ConditionalPolicyStatement) And() ConditionalPolicyStatement {
	var returns ConditionalPolicyStatement

	_jsii_.Invoke(
		c,
		"and",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalPolicyStatement) ContextAttribute(attribute *string) ConditionalAttributeAccessor {
	if err := c.validateContextAttributeParameters(attribute); err != nil {
		panic(err)
	}
	var returns ConditionalAttributeAccessor

	_jsii_.Invoke(
		c,
		"contextAttribute",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalPolicyStatement) Done() PolicyStatement {
	var returns PolicyStatement

	_jsii_.Invoke(
		c,
		"done",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalPolicyStatement) Or() ConditionalPolicyStatement {
	var returns ConditionalPolicyStatement

	_jsii_.Invoke(
		c,
		"or",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalPolicyStatement) PrincipalAttribute(attribute *string) ConditionalAttributeAccessor {
	if err := c.validatePrincipalAttributeParameters(attribute); err != nil {
		panic(err)
	}
	var returns ConditionalAttributeAccessor

	_jsii_.Invoke(
		c,
		"principalAttribute",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalPolicyStatement) ResourceAttribute(attribute *string) ConditionalAttributeAccessor {
	if err := c.validateResourceAttributeParameters(attribute); err != nil {
		panic(err)
	}
	var returns ConditionalAttributeAccessor

	_jsii_.Invoke(
		c,
		"resourceAttribute",
		[]interface{}{attribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConditionalPolicyStatement) Unless() ConditionalPolicyStatement {
	var returns ConditionalPolicyStatement

	_jsii_.Invoke(
		c,
		"unless",
		nil, // no parameters
		&returns,
	)

	return returns
}

