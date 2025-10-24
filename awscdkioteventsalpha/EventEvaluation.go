package awscdkioteventsalpha


// Information about the order in which events are evaluated and how actions are executed.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkioteventsalpha"
//   import "github.com/aws/aws-cdk-go/awscdkioteventsactionsalpha"
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//   var func IFunction
//
//
//   input := iotevents.NewInput(this, jsii.String("MyInput"), &InputProps{
//   	InputName: jsii.String("my_input"),
//   	 // optional
//   	AttributeJsonPaths: []*string{
//   		jsii.String("payload.deviceId"),
//   		jsii.String("payload.temperature"),
//   	},
//   })
//
//   warmState := iotevents.NewState(&StateProps{
//   	StateName: jsii.String("warm"),
//   	OnEnter: []Event{
//   		&Event{
//   			EventName: jsii.String("test-enter-event"),
//   			Condition: iotevents.Expression_CurrentInput(input),
//   			Actions: []IAction{
//   				actions.NewLambdaInvokeAction(func),
//   			},
//   		},
//   	},
//   	OnInput: []Event{
//   		&Event{
//   			 // optional
//   			EventName: jsii.String("test-input-event"),
//   			Actions: []IAction{
//   				actions.NewLambdaInvokeAction(func),
//   			},
//   		},
//   	},
//   	OnExit: []Event{
//   		&Event{
//   			 // optional
//   			EventName: jsii.String("test-exit-event"),
//   			Actions: []IAction{
//   				actions.NewLambdaInvokeAction(func),
//   			},
//   		},
//   	},
//   })
//   coldState := iotevents.NewState(&StateProps{
//   	StateName: jsii.String("cold"),
//   })
//
//   // transit to coldState when temperature is less than 15
//   warmState.TransitionTo(coldState, &TransitionOptions{
//   	EventName: jsii.String("to_coldState"),
//   	 // optional property, default by combining the names of the States
//   	When: iotevents.Expression_Lt(iotevents.Expression_InputAttribute(input, jsii.String("payload.temperature")), iotevents.Expression_FromString(jsii.String("15"))),
//   	Executing: []IAction{
//   		actions.NewLambdaInvokeAction(func),
//   	},
//   })
//   // transit to warmState when temperature is greater than or equal to 15
//   coldState.TransitionTo(warmState, &TransitionOptions{
//   	When: iotevents.Expression_Gte(iotevents.Expression_*InputAttribute(input, jsii.String("payload.temperature")), iotevents.Expression_*FromString(jsii.String("15"))),
//   })
//
//   iotevents.NewDetectorModel(this, jsii.String("MyDetectorModel"), &DetectorModelProps{
//   	DetectorModelName: jsii.String("test-detector-model"),
//   	 // optional
//   	Description: jsii.String("test-detector-model-description"),
//   	 // optional property, default is none
//   	EvaluationMethod: iotevents.EventEvaluation_SERIAL,
//   	 // optional property, default is iotevents.EventEvaluation.BATCH
//   	DetectorKey: jsii.String("payload.deviceId"),
//   	 // optional property, default is none and single detector instance will be created and all inputs will be routed to it
//   	InitialState: warmState,
//   })
//
// Experimental.
type EventEvaluation string

const (
	// When setting to BATCH, variables within a state are updated and events within a state are performed only after all event conditions are evaluated.
	// Experimental.
	EventEvaluation_BATCH EventEvaluation = "BATCH"
	// When setting to SERIAL, variables are updated and event conditions are evaluated in the order that the events are defined.
	// Experimental.
	EventEvaluation_SERIAL EventEvaluation = "SERIAL"
)

