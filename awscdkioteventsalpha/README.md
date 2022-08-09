# AWS::IoTEvents Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

AWS IoT Events enables you to monitor your equipment or device fleets for
failures or changes in operation, and to trigger actions when such events
occur.

## `DetectorModel`

The following example creates an AWS IoT Events detector model to your stack.
The detector model need a reference to at least one AWS IoT Events input.
AWS IoT Events inputs enable the detector to get MQTT payload values from IoT Core rules.

You can define built-in actions to use a timer or set a variable, or send data to other AWS resources.
See also [@aws-cdk/aws-iotevents-actions](https://docs.aws.amazon.com/cdk/api/v1/docs/aws-iotevents-actions-readme.html) for other actions.

```go
import iotevents "github.com/aws/aws-cdk-go/awscdkioteventsalpha"
import actions "github.com/aws/aws-cdk-go/awscdkioteventsactionsalpha"
import lambda "github.com/aws/aws-cdk-go/awscdk"

var func iFunction


input := iotevents.NewInput(this, jsii.String("MyInput"), &inputProps{
	inputName: jsii.String("my_input"),
	 // optional
	attributeJsonPaths: []*string{
		jsii.String("payload.deviceId"),
		jsii.String("payload.temperature"),
	},
})

warmState := iotevents.NewState(&stateProps{
	stateName: jsii.String("warm"),
	onEnter: []event{
		&event{
			eventName: jsii.String("test-enter-event"),
			condition: iotevents.expression.currentInput(input),
			actions: []iAction{
				actions.NewLambdaInvokeAction(func),
			},
		},
	},
	onInput: []*event{
		&event{
			 // optional
			eventName: jsii.String("test-input-event"),
			actions: []*iAction{
				actions.NewLambdaInvokeAction(func),
			},
		},
	},
	onExit: []*event{
		&event{
			 // optional
			eventName: jsii.String("test-exit-event"),
			actions: []*iAction{
				actions.NewLambdaInvokeAction(func),
			},
		},
	},
})
coldState := iotevents.NewState(&stateProps{
	stateName: jsii.String("cold"),
})

// transit to coldState when temperature is less than 15
warmState.transitionTo(coldState, &transitionOptions{
	eventName: jsii.String("to_coldState"),
	 // optional property, default by combining the names of the States
	when: iotevents.*expression.lt(iotevents.*expression.inputAttribute(input, jsii.String("payload.temperature")), iotevents.*expression.fromString(jsii.String("15"))),
	executing: []*iAction{
		actions.NewLambdaInvokeAction(func),
	},
})
// transit to warmState when temperature is greater than or equal to 15
coldState.transitionTo(warmState, &transitionOptions{
	when: iotevents.*expression.gte(iotevents.*expression.inputAttribute(input, jsii.String("payload.temperature")), iotevents.*expression.fromString(jsii.String("15"))),
})

iotevents.NewDetectorModel(this, jsii.String("MyDetectorModel"), &detectorModelProps{
	detectorModelName: jsii.String("test-detector-model"),
	 // optional
	description: jsii.String("test-detector-model-description"),
	 // optional property, default is none
	evaluationMethod: iotevents.eventEvaluation_SERIAL,
	 // optional property, default is iotevents.EventEvaluation.BATCH
	detectorKey: jsii.String("payload.deviceId"),
	 // optional property, default is none and single detector instance will be created and all inputs will be routed to it
	initialState: warmState,
})
```

To grant permissions to put messages in the input,
you can use the `grantWrite()` method:

```go
import iam "github.com/aws/aws-cdk-go/awscdk"
import iotevents "github.com/aws/aws-cdk-go/awscdkioteventsalpha"

var grantable iGrantable

input := iotevents.input.fromInputName(this, jsii.String("MyInput"), jsii.String("my_input"))

input.grantWrite(grantable)
```
