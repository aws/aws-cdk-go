// Receipt Detector Model actions for AWS IoT Events
package awscdkioteventsactionsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkioteventsactionsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkioteventsalpha/v2"
)

// The duration of the timer.
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
// Experimental.
type TimerDuration interface {
}

// The jsii proxy struct for TimerDuration
type jsiiProxy_TimerDuration struct {
	_ byte // padding
}

// Experimental.
func NewTimerDuration_Override(t TimerDuration) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iotevents-actions-alpha.TimerDuration",
		nil, // no parameters
		t,
	)
}

// Create a timer-duration from Duration.
//
// The range of the duration is 60-31622400 seconds.
// The evaluated result of the duration expression is rounded down to the nearest whole number.
// For example, if you set the timer to 60.99 seconds, the evaluated result of the duration expression is 60 seconds.
// Experimental.
func TimerDuration_FromDuration(duration awscdk.Duration) TimerDuration {
	_init_.Initialize()

	var returns TimerDuration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iotevents-actions-alpha.TimerDuration",
		"fromDuration",
		[]interface{}{duration},
		&returns,
	)

	return returns
}

// Create a timer-duration from Expression.
//
// You can use a string expression that includes numbers, variables ($variable.<variable-name>),
// and input values ($input.<input-name>.<path-to-datum>) as the duration.
//
// The range of the duration is 60-31622400 seconds.
// The evaluated result of the duration expression is rounded down to the nearest whole number.
// For example, if you set the timer to 60.99 seconds, the evaluated result of the duration expression is 60 seconds.
// Experimental.
func TimerDuration_FromExpression(expression awscdkioteventsalpha.Expression) TimerDuration {
	_init_.Initialize()

	var returns TimerDuration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-iotevents-actions-alpha.TimerDuration",
		"fromExpression",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

