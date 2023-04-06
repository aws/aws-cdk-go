// Receipt Detector Model actions for AWS IoT Events
package awscdkioteventsactionsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkioteventsactionsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdkioteventsactionsalpha/v2/internal"
	"github.com/aws/aws-cdk-go/awscdkioteventsalpha/v2"
)

// The action to write the data to an AWS Lambda function.
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

	if err := validateNewLambdaInvokeActionParameters(func_); err != nil {
		panic(err)
	}
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

