package awsiotevents


// Properties for options of state transition.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//   var func iFunction
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
//   	OnEnter: []event{
//   		&event{
//   			EventName: jsii.String("test-enter-event"),
//   			Condition: iotevents.Expression_CurrentInput(input),
//   			Actions: []iAction{
//   				actions.NewLambdaInvokeAction(func),
//   			},
//   		},
//   	},
//   	OnInput: []*event{
//   		&event{
//   			 // optional
//   			EventName: jsii.String("test-input-event"),
//   			Actions: []*iAction{
//   				actions.NewLambdaInvokeAction(func),
//   			},
//   		},
//   	},
//   	OnExit: []*event{
//   		&event{
//   			 // optional
//   			EventName: jsii.String("test-exit-event"),
//   			Actions: []*iAction{
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
//   	Executing: []*iAction{
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
type TransitionOptions struct {
	// The condition that is used to determine to cause the state transition and the actions.
	//
	// When this was evaluated to `true`, the state transition and the actions are triggered.
	// Experimental.
	When Expression `field:"required" json:"when" yaml:"when"`
	// The name of the event.
	// Experimental.
	EventName *string `field:"optional" json:"eventName" yaml:"eventName"`
	// The actions to be performed with the transition.
	// Experimental.
	Executing *[]IAction `field:"optional" json:"executing" yaml:"executing"`
}

