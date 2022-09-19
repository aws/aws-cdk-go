// The CDK Construct Library for AWS::IoTEvents
package awscdkioteventsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkioteventsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Expression for events in Detector Model state.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import iotevents "github.com/aws/aws-cdk-go/awscdkioteventsalpha"
//   import actions "github.com/aws/aws-cdk-go/awscdkioteventsactionsalpha"
//
//   var input iInput
//
//   state := iotevents.NewState(&stateProps{
//   	stateName: jsii.String("MyState"),
//   	onEnter: []event{
//   		&event{
//   			eventName: jsii.String("test-event"),
//   			condition: iotevents.expression.currentInput(input),
//   			actions: []iAction{
//   				actions.NewSetTimerAction(jsii.String("MyTimer"), map[string]interface{}{
//   					"duration": cdk.Duration_seconds(jsii.Number(60)),
//   				}),
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
		"@aws-cdk/aws-iotevents-alpha.Expression",
		nil, // no parameters
		e,
	)
}

// Create a expression for the Addition operator.
// Experimental.
func Expression_Add(left Expression, right Expression) Expression {
	_init_.Initialize()

	if err := validateExpression_AddParameters(left, right); err != nil {
		panic(err)
	}
	var returns Expression

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iotevents-alpha.Expression",
		"add",
		[]interface{}{left, right},
		&returns,
	)

	return returns
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
		"@aws-cdk/aws-iotevents-alpha.Expression",
		"and",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

// Create a expression for the Bitwise AND operator.
// Experimental.
func Expression_BitwiseAnd(left Expression, right Expression) Expression {
	_init_.Initialize()

	if err := validateExpression_BitwiseAndParameters(left, right); err != nil {
		panic(err)
	}
	var returns Expression

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iotevents-alpha.Expression",
		"bitwiseAnd",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

// Create a expression for the Bitwise OR operator.
// Experimental.
func Expression_BitwiseOr(left Expression, right Expression) Expression {
	_init_.Initialize()

	if err := validateExpression_BitwiseOrParameters(left, right); err != nil {
		panic(err)
	}
	var returns Expression

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iotevents-alpha.Expression",
		"bitwiseOr",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

// Create a expression for the Bitwise XOR operator.
// Experimental.
func Expression_BitwiseXor(left Expression, right Expression) Expression {
	_init_.Initialize()

	if err := validateExpression_BitwiseXorParameters(left, right); err != nil {
		panic(err)
	}
	var returns Expression

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iotevents-alpha.Expression",
		"bitwiseXor",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

// Create a expression for the String Concatenation operator.
// Experimental.
func Expression_Concat(left Expression, right Expression) Expression {
	_init_.Initialize()

	if err := validateExpression_ConcatParameters(left, right); err != nil {
		panic(err)
	}
	var returns Expression

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iotevents-alpha.Expression",
		"concat",
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
		"@aws-cdk/aws-iotevents-alpha.Expression",
		"currentInput",
		[]interface{}{input},
		&returns,
	)

	return returns
}

// Create a expression for the Division operator.
// Experimental.
func Expression_Divide(left Expression, right Expression) Expression {
	_init_.Initialize()

	if err := validateExpression_DivideParameters(left, right); err != nil {
		panic(err)
	}
	var returns Expression

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iotevents-alpha.Expression",
		"divide",
		[]interface{}{left, right},
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
		"@aws-cdk/aws-iotevents-alpha.Expression",
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
		"@aws-cdk/aws-iotevents-alpha.Expression",
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
		"@aws-cdk/aws-iotevents-alpha.Expression",
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
		"@aws-cdk/aws-iotevents-alpha.Expression",
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
		"@aws-cdk/aws-iotevents-alpha.Expression",
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
		"@aws-cdk/aws-iotevents-alpha.Expression",
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
		"@aws-cdk/aws-iotevents-alpha.Expression",
		"lte",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

// Create a expression for the Multiplication operator.
// Experimental.
func Expression_Multiply(left Expression, right Expression) Expression {
	_init_.Initialize()

	if err := validateExpression_MultiplyParameters(left, right); err != nil {
		panic(err)
	}
	var returns Expression

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iotevents-alpha.Expression",
		"multiply",
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
		"@aws-cdk/aws-iotevents-alpha.Expression",
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
		"@aws-cdk/aws-iotevents-alpha.Expression",
		"or",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

// Create a expression for the Subtraction operator.
// Experimental.
func Expression_Subtract(left Expression, right Expression) Expression {
	_init_.Initialize()

	if err := validateExpression_SubtractParameters(left, right); err != nil {
		panic(err)
	}
	var returns Expression

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iotevents-alpha.Expression",
		"subtract",
		[]interface{}{left, right},
		&returns,
	)

	return returns
}

// Create a expression for function `timeout("timer-name")`.
//
// It is evaluated to true if the specified timer has elapsed.
// You can define a timer only using the `setTimer` action.
// Experimental.
func Expression_Timeout(timerName *string) Expression {
	_init_.Initialize()

	if err := validateExpression_TimeoutParameters(timerName); err != nil {
		panic(err)
	}
	var returns Expression

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iotevents-alpha.Expression",
		"timeout",
		[]interface{}{timerName},
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

