// The CDK Construct Library for AWS::IoTEvents
package awscdkioteventsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for defining an AWS IoT Events detector model.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkioteventsalpha"
//   import "github.com/aws/aws-cdk-go/awscdkioteventsactionsalpha"
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
type DetectorModelProps struct {
	// The state that is entered at the creation of each detector.
	// Experimental.
	InitialState State `field:"required" json:"initialState" yaml:"initialState"`
	// A brief description of the detector model.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The value used to identify a detector instance.
	//
	// When a device or system sends input, a new
	// detector instance with a unique key value is created. AWS IoT Events can continue to route
	// input to its corresponding detector instance based on this identifying information.
	//
	// This parameter uses a JSON-path expression to select the attribute-value pair in the message
	// payload that is used for identification. To route the message to the correct detector instance,
	// the device must send a message payload that contains the same attribute-value.
	// Experimental.
	DetectorKey *string `field:"optional" json:"detectorKey" yaml:"detectorKey"`
	// The name of the detector model.
	// Experimental.
	DetectorModelName *string `field:"optional" json:"detectorModelName" yaml:"detectorModelName"`
	// Information about the order in which events are evaluated and how actions are executed.
	//
	// When setting to SERIAL, variables are updated and event conditions are evaluated in the order
	// that the events are defined.
	// When setting to BATCH, variables within a state are updated and events within a state are
	// performed only after all event conditions are evaluated.
	// Experimental.
	EvaluationMethod EventEvaluation `field:"optional" json:"evaluationMethod" yaml:"evaluationMethod"`
	// The role that grants permission to AWS IoT Events to perform its operations.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

