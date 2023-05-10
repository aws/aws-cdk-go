package awsiotevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Defines a state of a detector.
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
type State interface {
	// The name of the state.
	// Experimental.
	StateName() *string
	// Add a transition event to the state.
	//
	// The transition event will be triggered if condition is evaluated to `true`.
	// Experimental.
	TransitionTo(targetState State, options *TransitionOptions)
}

// The jsii proxy struct for State
type jsiiProxy_State struct {
	_ byte // padding
}

func (j *jsiiProxy_State) StateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateName",
		&returns,
	)
	return returns
}


// Experimental.
func NewState(props *StateProps) State {
	_init_.Initialize()

	if err := validateNewStateParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_State{}

	_jsii_.Create(
		"monocdk.aws_iotevents.State",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewState_Override(s State, props *StateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iotevents.State",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_State) TransitionTo(targetState State, options *TransitionOptions) {
	if err := s.validateTransitionToParameters(targetState, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"transitionTo",
		[]interface{}{targetState, options},
	)
}

