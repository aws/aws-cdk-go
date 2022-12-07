package awsiotevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Expression for events in Detector Model state.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import iotevents "github.com/aws/aws-cdk-go/awscdk"
//   import actions "github.com/aws/aws-cdk-go/awscdk"
//
//   var input iInput
//
//
//   state := iotevents.NewState(&stateProps{
//   	stateName: jsii.String("MyState"),
//   	onEnter: []event{
//   		&event{
//   			eventName: jsii.String("test-event"),
//   			condition: iotevents.expression.currentInput(input),
//   			actions: []iAction{
//   				actions,
//   				[]interface{}{
//   					actions.NewSetVariableAction(jsii.String("MyVariable"), iotevents.*expression.inputAttribute(input, jsii.String("payload.temperature"))),
//   				},
//   			},
//   		},
//   	},
//   })
//
// See: https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html
//
// Experimental.
type Expression interface {
	// This is called to evaluate the expression.
	// Experimental.
	Evaluate(parentPriority *float64) *string
}

// The jsii proxy struct for Expression
type jsiiProxy_Expression struct {
	_ byte // padding
}

// Experimental.
func NewExpression_Override(e Expression) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iotevents.Expression",
		nil, // no parameters
		e,
	)
}

// Create a expression for the AND operator.
// Experimental.
func Expression_And(left Expression, right Expression) Expression {
	_init_.Initialize()

	if err := validateExpression_AndParameters(left, right); err != nil {
		panic(err)
	}
	var returns Expression

	_jsii_.StaticInvoke(
		"monocdk.aws_iotevents.Expression",
		"and",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

// Create a expression for function `currentInput()`.
//
// It is evaluated to true if the specified input message was received.
// Experimental.
func Expression_CurrentInput(input IInput) Expression {
	_init_.Initialize()

	if err := validateExpression_CurrentInputParameters(input); err != nil {
		panic(err)
	}
	var returns Expression

	_jsii_.StaticInvoke(
		"monocdk.aws_iotevents.Expression",
		"currentInput",
		[]interface{}{input},
		&returns,
	)

	return returns
}

// Create a expression for the Equal operator.
// Experimental.
func Expression_Eq(left Expression, right Expression) Expression {
	_init_.Initialize()

	if err := validateExpression_EqParameters(left, right); err != nil {
		panic(err)
	}
	var returns Expression

	_jsii_.StaticInvoke(
		"monocdk.aws_iotevents.Expression",
		"eq",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

// Create a expression from the given string.
// Experimental.
func Expression_FromString(value *string) Expression {
	_init_.Initialize()

	if err := validateExpression_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns Expression

	_jsii_.StaticInvoke(
		"monocdk.aws_iotevents.Expression",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Create a expression for the Greater Than operator.
// Experimental.
func Expression_Gt(left Expression, right Expression) Expression {
	_init_.Initialize()

	if err := validateExpression_GtParameters(left, right); err != nil {
		panic(err)
	}
	var returns Expression

	_jsii_.StaticInvoke(
		"monocdk.aws_iotevents.Expression",
		"gt",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

// Create a expression for the Greater Than Or Equal operator.
// Experimental.
func Expression_Gte(left Expression, right Expression) Expression {
	_init_.Initialize()

	if err := validateExpression_GteParameters(left, right); err != nil {
		panic(err)
	}
	var returns Expression

	_jsii_.StaticInvoke(
		"monocdk.aws_iotevents.Expression",
		"gte",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

// Create a expression for get an input attribute as `$input.TemperatureInput.temperatures[2]`.
// Experimental.
func Expression_InputAttribute(input IInput, path *string) Expression {
	_init_.Initialize()

	if err := validateExpression_InputAttributeParameters(input, path); err != nil {
		panic(err)
	}
	var returns Expression

	_jsii_.StaticInvoke(
		"monocdk.aws_iotevents.Expression",
		"inputAttribute",
		[]interface{}{input, path},
		&returns,
	)

	return returns
}

// Create a expression for the Less Than operator.
// Experimental.
func Expression_Lt(left Expression, right Expression) Expression {
	_init_.Initialize()

	if err := validateExpression_LtParameters(left, right); err != nil {
		panic(err)
	}
	var returns Expression

	_jsii_.StaticInvoke(
		"monocdk.aws_iotevents.Expression",
		"lt",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

// Create a expression for the Less Than Or Equal operator.
// Experimental.
func Expression_Lte(left Expression, right Expression) Expression {
	_init_.Initialize()

	if err := validateExpression_LteParameters(left, right); err != nil {
		panic(err)
	}
	var returns Expression

	_jsii_.StaticInvoke(
		"monocdk.aws_iotevents.Expression",
		"lte",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

// Create a expression for the Not Equal operator.
// Experimental.
func Expression_Neq(left Expression, right Expression) Expression {
	_init_.Initialize()

	if err := validateExpression_NeqParameters(left, right); err != nil {
		panic(err)
	}
	var returns Expression

	_jsii_.StaticInvoke(
		"monocdk.aws_iotevents.Expression",
		"neq",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

// Create a expression for the OR operator.
// Experimental.
func Expression_Or(left Expression, right Expression) Expression {
	_init_.Initialize()

	if err := validateExpression_OrParameters(left, right); err != nil {
		panic(err)
	}
	var returns Expression

	_jsii_.StaticInvoke(
		"monocdk.aws_iotevents.Expression",
		"or",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Expression) Evaluate(parentPriority *float64) *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"evaluate",
		[]interface{}{parentPriority},
		&returns,
	)

	return returns
}

