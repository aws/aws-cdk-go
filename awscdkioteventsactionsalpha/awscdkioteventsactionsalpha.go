// Receipt Detector Model actions for AWS IoT Events
package awscdkioteventsactionsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkioteventsactionsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdkioteventsactionsalpha/v2/internal"
	"github.com/aws/aws-cdk-go/awscdkioteventsalpha/v2"
)

// The action to delete an existing timer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import iotevents_actions_alpha "github.com/aws/aws-cdk-go/awscdkioteventsactionsalpha"
//
//   clearTimerAction := iotevents_actions_alpha.NewClearTimerAction(jsii.String("timerName"))
//
// Experimental.
type ClearTimerAction interface {
	awscdkioteventsalpha.IAction
}

// The jsii proxy struct for ClearTimerAction
type jsiiProxy_ClearTimerAction struct {
	internal.Type__awscdkioteventsalphaIAction
}

// Experimental.
func NewClearTimerAction(timerName *string) ClearTimerAction {
	_init_.Initialize()

	j := jsiiProxy_ClearTimerAction{}

	_jsii_.Create(
		"@aws-cdk/aws-iotevents-actions-alpha.ClearTimerAction",
		[]interface{}{timerName},
		&j,
	)

	return &j
}

// Experimental.
func NewClearTimerAction_Override(c ClearTimerAction, timerName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iotevents-actions-alpha.ClearTimerAction",
		[]interface{}{timerName},
		c,
	)
}

// The action to write the data to an AWS Lambda function.
//
// Example:
//   import iotevents "github.com/aws/aws-cdk-go/awscdkioteventsalpha"
//   import actions "github.com/aws/aws-cdk-go/awscdkioteventsactionsalpha"
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//   var func iFunction
//
//
//   input := iotevents.NewInput(this, jsii.String("MyInput"), &inputProps{
//   	inputName: jsii.String("my_input"),
//   	 // optional
//   	attributeJsonPaths: []*string{
//   		jsii.String("payload.deviceId"),
//   		jsii.String("payload.temperature"),
//   	},
//   })
//
//   warmState := iotevents.NewState(&stateProps{
//   	stateName: jsii.String("warm"),
//   	onEnter: []event{
//   		&event{
//   			eventName: jsii.String("test-enter-event"),
//   			condition: iotevents.expression.currentInput(input),
//   			actions: []iAction{
//   				actions.NewLambdaInvokeAction(func),
//   			},
//   		},
//   	},
//   	onInput: []*event{
//   		&event{
//   			 // optional
//   			eventName: jsii.String("test-input-event"),
//   			actions: []*iAction{
//   				actions.NewLambdaInvokeAction(func),
//   			},
//   		},
//   	},
//   	onExit: []*event{
//   		&event{
//   			 // optional
//   			eventName: jsii.String("test-exit-event"),
//   			actions: []*iAction{
//   				actions.NewLambdaInvokeAction(func),
//   			},
//   		},
//   	},
//   })
//   coldState := iotevents.NewState(&stateProps{
//   	stateName: jsii.String("cold"),
//   })
//
//   // transit to coldState when temperature is less than 15
//   warmState.transitionTo(coldState, &transitionOptions{
//   	eventName: jsii.String("to_coldState"),
//   	 // optional property, default by combining the names of the States
//   	when: iotevents.*expression.lt(iotevents.*expression.inputAttribute(input, jsii.String("payload.temperature")), iotevents.*expression.fromString(jsii.String("15"))),
//   	executing: []*iAction{
//   		actions.NewLambdaInvokeAction(func),
//   	},
//   })
//   // transit to warmState when temperature is greater than or equal to 15
//   coldState.transitionTo(warmState, &transitionOptions{
//   	when: iotevents.*expression.gte(iotevents.*expression.inputAttribute(input, jsii.String("payload.temperature")), iotevents.*expression.fromString(jsii.String("15"))),
//   })
//
//   iotevents.NewDetectorModel(this, jsii.String("MyDetectorModel"), &detectorModelProps{
//   	detectorModelName: jsii.String("test-detector-model"),
//   	 // optional
//   	description: jsii.String("test-detector-model-description"),
//   	 // optional property, default is none
//   	evaluationMethod: iotevents.eventEvaluation_SERIAL,
//   	 // optional property, default is iotevents.EventEvaluation.BATCH
//   	detectorKey: jsii.String("payload.deviceId"),
//   	 // optional property, default is none and single detector instance will be created and all inputs will be routed to it
//   	initialState: warmState,
//   })
//
// Experimental.
type LambdaInvokeAction interface {
	awscdkioteventsalpha.IAction
}

// The jsii proxy struct for LambdaInvokeAction
type jsiiProxy_LambdaInvokeAction struct {
	internal.Type__awscdkioteventsalphaIAction
}

// Experimental.
func NewLambdaInvokeAction(func_ awslambda.IFunction) LambdaInvokeAction {
	_init_.Initialize()

	j := jsiiProxy_LambdaInvokeAction{}

	_jsii_.Create(
		"@aws-cdk/aws-iotevents-actions-alpha.LambdaInvokeAction",
		[]interface{}{func_},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaInvokeAction_Override(l LambdaInvokeAction, func_ awslambda.IFunction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iotevents-actions-alpha.LambdaInvokeAction",
		[]interface{}{func_},
		l,
	)
}

// The action to reset an existing timer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import iotevents_actions_alpha "github.com/aws/aws-cdk-go/awscdkioteventsactionsalpha"
//
//   resetTimerAction := iotevents_actions_alpha.NewResetTimerAction(jsii.String("timerName"))
//
// Experimental.
type ResetTimerAction interface {
	awscdkioteventsalpha.IAction
}

// The jsii proxy struct for ResetTimerAction
type jsiiProxy_ResetTimerAction struct {
	internal.Type__awscdkioteventsalphaIAction
}

// Experimental.
func NewResetTimerAction(timerName *string) ResetTimerAction {
	_init_.Initialize()

	j := jsiiProxy_ResetTimerAction{}

	_jsii_.Create(
		"@aws-cdk/aws-iotevents-actions-alpha.ResetTimerAction",
		[]interface{}{timerName},
		&j,
	)

	return &j
}

// Experimental.
func NewResetTimerAction_Override(r ResetTimerAction, timerName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iotevents-actions-alpha.ResetTimerAction",
		[]interface{}{timerName},
		r,
	)
}

// The action to create a timer with duration in seconds.
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
type SetTimerAction interface {
	awscdkioteventsalpha.IAction
}

// The jsii proxy struct for SetTimerAction
type jsiiProxy_SetTimerAction struct {
	internal.Type__awscdkioteventsalphaIAction
}

// Experimental.
func NewSetTimerAction(timerName *string, timerDuration TimerDuration) SetTimerAction {
	_init_.Initialize()

	j := jsiiProxy_SetTimerAction{}

	_jsii_.Create(
		"@aws-cdk/aws-iotevents-actions-alpha.SetTimerAction",
		[]interface{}{timerName, timerDuration},
		&j,
	)

	return &j
}

// Experimental.
func NewSetTimerAction_Override(s SetTimerAction, timerName *string, timerDuration TimerDuration) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iotevents-actions-alpha.SetTimerAction",
		[]interface{}{timerName, timerDuration},
		s,
	)
}

// The action to create a variable with a specified value.
//
// Example:
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
//   				actions.NewSetVariableAction(jsii.String("MyVariable"), iotevents.*expression.inputAttribute(input, jsii.String("payload.temperature"))),
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type SetVariableAction interface {
	awscdkioteventsalpha.IAction
}

// The jsii proxy struct for SetVariableAction
type jsiiProxy_SetVariableAction struct {
	internal.Type__awscdkioteventsalphaIAction
}

// Experimental.
func NewSetVariableAction(variableName *string, value awscdkioteventsalpha.Expression) SetVariableAction {
	_init_.Initialize()

	j := jsiiProxy_SetVariableAction{}

	_jsii_.Create(
		"@aws-cdk/aws-iotevents-actions-alpha.SetVariableAction",
		[]interface{}{variableName, value},
		&j,
	)

	return &j
}

// Experimental.
func NewSetVariableAction_Override(s SetVariableAction, variableName *string, value awscdkioteventsalpha.Expression) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iotevents-actions-alpha.SetVariableAction",
		[]interface{}{variableName, value},
		s,
	)
}

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

