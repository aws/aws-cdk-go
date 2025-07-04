package awscdkioteventsalpha


// Properties for defining a state of a detector.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import "github.com/aws/aws-cdk-go/awscdkioteventsalpha"
//   import actions "github.com/aws/aws-cdk-go/awscdkioteventsactionsalpha"
//
//   var input iInput
//
//   state := iotevents.NewState(&StateProps{
//   	StateName: jsii.String("MyState"),
//   	OnEnter: []event{
//   		&event{
//   			EventName: jsii.String("test-event"),
//   			Condition: iotevents.Expression_CurrentInput(input),
//   			Actions: []iAction{
//   				actions.NewSetTimerAction(jsii.String("MyTimer"), map[string]interface{}{
//   					"duration": cdk.Duration_seconds(jsii.Number(60)),
//   				}),
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
	// Default: - no events will trigger on entering this state.
	//
	// Experimental.
	OnEnter *[]*Event `field:"optional" json:"onEnter" yaml:"onEnter"`
	// Specifies the events on exit.
	//
	// The conditions of the events are evaluated when an exiting this state.
	// If the condition evaluates to `true`, the actions of the event will be executed.
	// Default: - no events will trigger on exiting this state.
	//
	// Experimental.
	OnExit *[]*Event `field:"optional" json:"onExit" yaml:"onExit"`
	// Specifies the events on input.
	//
	// The conditions of the events will be evaluated when any input is received.
	// If the condition of the event evaluates to `true`, the actions of the event will be executed.
	// Default: - no events will trigger on input in this state.
	//
	// Experimental.
	OnInput *[]*Event `field:"optional" json:"onInput" yaml:"onInput"`
}

