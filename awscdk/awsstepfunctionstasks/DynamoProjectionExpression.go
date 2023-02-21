package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Class to generate projection expression.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dynamoProjectionExpression := awscdk.Aws_stepfunctions_tasks.NewDynamoProjectionExpression()
//
type DynamoProjectionExpression interface {
	// Adds the array literal access for passed index.
	AtIndex(index *float64) DynamoProjectionExpression
	// converts and return the string expression.
	ToString() *string
	// Adds the passed attribute to the chain.
	WithAttribute(attr *string) DynamoProjectionExpression
}

// The jsii proxy struct for DynamoProjectionExpression
type jsiiProxy_DynamoProjectionExpression struct {
	_ byte // padding
}

func NewDynamoProjectionExpression() DynamoProjectionExpression {
	_init_.Initialize()

	j := jsiiProxy_DynamoProjectionExpression{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoProjectionExpression",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewDynamoProjectionExpression_Override(d DynamoProjectionExpression) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoProjectionExpression",
		nil, // no parameters
		d,
	)
}

func (d *jsiiProxy_DynamoProjectionExpression) AtIndex(index *float64) DynamoProjectionExpression {
	if err := d.validateAtIndexParameters(index); err != nil {
		panic(err)
	}
	var returns DynamoProjectionExpression

	_jsii_.Invoke(
		d,
		"atIndex",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoProjectionExpression) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoProjectionExpression) WithAttribute(attr *string) DynamoProjectionExpression {
	if err := d.validateWithAttributeParameters(attr); err != nil {
		panic(err)
	}
	var returns DynamoProjectionExpression

	_jsii_.Invoke(
		d,
		"withAttribute",
		[]interface{}{attr},
		&returns,
	)

	return returns
}

