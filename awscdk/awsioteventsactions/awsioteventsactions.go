package awsioteventsactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsiotevents"
	"github.com/aws/aws-cdk-go/awscdk/awsioteventsactions/internal"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/constructs-go/constructs/v3"
)

// The action to write the data to an AWS Lambda function.
//
// Example:
//   import iotevents "github.com/aws/aws-cdk-go/awscdk"import actions "github.com/aws/aws-cdk-go/awscdk"import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//   var func iFunction
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
//   // transit to coldState when temperature is 10
//   warmState.transitionTo(coldState, &transitionOptions{
//   	eventName: jsii.String("to_coldState"),
//   	 // optional property, default by combining the names of the States
//   	when: iotevents.*expression.eq(iotevents.*expression.inputAttribute(input, jsii.String("payload.temperature")), iotevents.*expression.fromString(jsii.String("10"))),
//   	executing: []*iAction{
//   		actions.NewLambdaInvokeAction(func),
//   	},
//   })
//   // transit to warmState when temperature is 20
//   coldState.transitionTo(warmState, &transitionOptions{
//   	when: iotevents.*expression.eq(iotevents.*expression.inputAttribute(input, jsii.String("payload.temperature")), iotevents.*expression.fromString(jsii.String("20"))),
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
	awsiotevents.IAction
	// Returns the AWS IoT Events action specification.
	// Experimental.
	Bind(_scope constructs.Construct, options *awsiotevents.ActionBindOptions) *awsiotevents.ActionConfig
}

// The jsii proxy struct for LambdaInvokeAction
type jsiiProxy_LambdaInvokeAction struct {
	internal.Type__awsioteventsIAction
}

// Experimental.
func NewLambdaInvokeAction(func_ awslambda.IFunction) LambdaInvokeAction {
	_init_.Initialize()

	j := jsiiProxy_LambdaInvokeAction{}

	_jsii_.Create(
		"monocdk.aws_iotevents_actions.LambdaInvokeAction",
		[]interface{}{func_},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaInvokeAction_Override(l LambdaInvokeAction, func_ awslambda.IFunction) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iotevents_actions.LambdaInvokeAction",
		[]interface{}{func_},
		l,
	)
}

func (l *jsiiProxy_LambdaInvokeAction) Bind(_scope constructs.Construct, options *awsiotevents.ActionBindOptions) *awsiotevents.ActionConfig {
	var returns *awsiotevents.ActionConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{_scope, options},
		&returns,
	)

	return returns
}

// The action to create a variable with a specified value.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import iotevents "github.com/aws/aws-cdk-go/awscdk"import actions "github.com/aws/aws-cdk-go/awscdk"
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
// Experimental.
type SetVariableAction interface {
	awsiotevents.IAction
	// Returns the AWS IoT Events action specification.
	// Experimental.
	Bind(_scope constructs.Construct, _options *awsiotevents.ActionBindOptions) *awsiotevents.ActionConfig
}

// The jsii proxy struct for SetVariableAction
type jsiiProxy_SetVariableAction struct {
	internal.Type__awsioteventsIAction
}

// Experimental.
func NewSetVariableAction(variableName *string, value awsiotevents.Expression) SetVariableAction {
	_init_.Initialize()

	j := jsiiProxy_SetVariableAction{}

	_jsii_.Create(
		"monocdk.aws_iotevents_actions.SetVariableAction",
		[]interface{}{variableName, value},
		&j,
	)

	return &j
}

// Experimental.
func NewSetVariableAction_Override(s SetVariableAction, variableName *string, value awsiotevents.Expression) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iotevents_actions.SetVariableAction",
		[]interface{}{variableName, value},
		s,
	)
}

func (s *jsiiProxy_SetVariableAction) Bind(_scope constructs.Construct, _options *awsiotevents.ActionBindOptions) *awsiotevents.ActionConfig {
	var returns *awsiotevents.ActionConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope, _options},
		&returns,
	)

	return returns
}

