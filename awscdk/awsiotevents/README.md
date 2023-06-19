# AWS::IoTEvents Construct Library

AWS IoT Events enables you to monitor your equipment or device fleets for
failures or changes in operation, and to trigger actions when such events
occur.

## Installation

Install the module:

```console
$ npm i @aws-cdk/aws-iotevents
```

Import it into your code:

```go
import iotevents "github.com/aws/aws-cdk-go/awscdk"
```

## `DetectorModel`

The following example creates an AWS IoT Events detector model to your stack.
The detector model need a reference to at least one AWS IoT Events input.
AWS IoT Events inputs enable the detector to get MQTT payload values from IoT Core rules.

You can define built-in actions to use a timer or set a variable, or send data to other AWS resources.
See also [@aws-cdk/aws-iotevents-actions](https://docs.aws.amazon.com/cdk/api/v1/docs/aws-iotevents-actions-readme.html) for other actions.

```go
import "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import lambda "github.com/aws/aws-cdk-go/awscdk"

var func iFunction


input := iotevents.NewInput(this, jsii.String("MyInput"), &InputProps{
	InputName: jsii.String("my_input"),
	 // optional
	AttributeJsonPaths: []*string{
		jsii.String("payload.deviceId"),
		jsii.String("payload.temperature"),
	},
})

warmState := iotevents.NewState(&StateProps{
	StateName: jsii.String("warm"),
	OnEnter: []event{
		&event{
			EventName: jsii.String("test-enter-event"),
			Condition: iotevents.Expression_CurrentInput(input),
			Actions: []iAction{
				actions.NewLambdaInvokeAction(func),
			},
		},
	},
	OnInput: []*event{
		&event{
			 // optional
			EventName: jsii.String("test-input-event"),
			Actions: []*iAction{
				actions.NewLambdaInvokeAction(func),
			},
		},
	},
	OnExit: []*event{
		&event{
			 // optional
			EventName: jsii.String("test-exit-event"),
			Actions: []*iAction{
				actions.NewLambdaInvokeAction(func),
			},
		},
	},
})
coldState := iotevents.NewState(&StateProps{
	StateName: jsii.String("cold"),
})

// transit to coldState when temperature is less than 15
warmState.TransitionTo(coldState, &TransitionOptions{
	EventName: jsii.String("to_coldState"),
	 // optional property, default by combining the names of the States
	When: iotevents.Expression_Lt(iotevents.Expression_InputAttribute(input, jsii.String("payload.temperature")), iotevents.Expression_FromString(jsii.String("15"))),
	Executing: []*iAction{
		actions.NewLambdaInvokeAction(func),
	},
})
// transit to warmState when temperature is greater than or equal to 15
coldState.TransitionTo(warmState, &TransitionOptions{
	When: iotevents.Expression_Gte(iotevents.Expression_*InputAttribute(input, jsii.String("payload.temperature")), iotevents.Expression_*FromString(jsii.String("15"))),
})

iotevents.NewDetectorModel(this, jsii.String("MyDetectorModel"), &DetectorModelProps{
	DetectorModelName: jsii.String("test-detector-model"),
	 // optional
	Description: jsii.String("test-detector-model-description"),
	 // optional property, default is none
	EvaluationMethod: iotevents.EventEvaluation_SERIAL,
	 // optional property, default is iotevents.EventEvaluation.BATCH
	DetectorKey: jsii.String("payload.deviceId"),
	 // optional property, default is none and single detector instance will be created and all inputs will be routed to it
	InitialState: warmState,
})
```

To grant permissions to put messages in the input,
you can use the `grantWrite()` method:

```go
import iam "github.com/aws/aws-cdk-go/awscdk"
import iotevents "github.com/aws/aws-cdk-go/awscdk"

var grantable iGrantable

input := iotevents.Input_FromInputName(this, jsii.String("MyInput"), jsii.String("my_input"))

input.GrantWrite(grantable)
```
