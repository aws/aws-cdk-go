package awsiotevents


// Properties for defining a state of a detector.
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
type StateProps struct {
	// The name of the state.
	// Experimental.
	StateName *string `field:"required" json:"stateName" yaml:"stateName"`
	// Specifies the events on enter.
	//
	// The conditions of the events will be evaluated when entering this state.
	// If the condition of the event evaluates to `true`, the actions of the event will be executed.
	// Experimental.
	OnEnter *[]*Event `field:"optional" json:"onEnter" yaml:"onEnter"`
	// Specifies the events on exit.
	//
	// The conditions of the events are evaluated when an exiting this state.
	// If the condition evaluates to `true`, the actions of the event will be executed.
	// Experimental.
	OnExit *[]*Event `field:"optional" json:"onExit" yaml:"onExit"`
	// Specifies the events on input.
	//
	// The conditions of the events will be evaluated when any input is received.
	// If the condition of the event evaluates to `true`, the actions of the event will be executed.
	// Experimental.
	OnInput *[]*Event `field:"optional" json:"onInput" yaml:"onInput"`
}

