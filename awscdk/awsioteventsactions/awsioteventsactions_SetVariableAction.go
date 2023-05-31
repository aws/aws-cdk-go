package awsioteventsactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsiotevents"
	"github.com/aws/aws-cdk-go/awscdk/awsioteventsactions/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// The action to create a variable with a specified value.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import actions "github.com/aws/aws-cdk-go/awscdk"
//
//   var input iInput
//
//
//   state := iotevents.NewState(&StateProps{
//   	StateName: jsii.String("MyState"),
//   	OnEnter: []event{
//   		&event{
//   			EventName: jsii.String("test-event"),
//   			Condition: iotevents.Expression_CurrentInput(input),
//   			Actions: []iAction{
//   				actions,
//   				[]interface{}{
//   					actions.NewSetVariableAction(jsii.String("MyVariable"), iotevents.Expression_InputAttribute(input, jsii.String("payload.temperature"))),
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

	if err := validateNewSetVariableActionParameters(variableName, value); err != nil {
		panic(err)
	}
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
	if err := s.validateBindParameters(_scope, _options); err != nil {
		panic(err)
	}
	var returns *awsiotevents.ActionConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope, _options},
		&returns,
	)

	return returns
}

