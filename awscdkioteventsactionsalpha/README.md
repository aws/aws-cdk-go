# Actions for AWS::IoTEvents Detector Model

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

This library contains integration classes to specify actions of state events of Detector Model in `@aws-cdk/aws-iotevents-alpha`.
Instances of these classes should be passed to `State` defined in `@aws-cdk/aws-iotevents-alpha`
You can define built-in actions to use a timer or set a variable, or send data to other AWS resources.

This library contains integration classes to use a timer or set a variable, or send data to other AWS resources.
AWS IoT Events can trigger actions when it detects a specified event or transition event.

Currently supported are:

* Use timer
* Set variable to detector instance
* Invoke a Lambda function

## Use timer

The code snippet below creates an Action that creates the timer with duration in seconds.

```go
// Example automatically generated from non-compiling source. May contain errors.
import "github.com/aws/aws-cdk-go/awscdkioteventsalpha"
import actions "github.com/aws/aws-cdk-go/awscdkioteventsactionsalpha"

var input IInput

state := iotevents.NewState(&StateProps{
	StateName: jsii.String("MyState"),
	OnEnter: []Event{
		&Event{
			EventName: jsii.String("test-event"),
			Condition: iotevents.Expression_CurrentInput(input),
			Actions: []IAction{
				actions.NewSetTimerAction(jsii.String("MyTimer"), map[string]interface{}{
					"duration": cdk.Duration_seconds(jsii.Number(60)),
				}),
			},
		},
	},
})
```

Setting duration by [IoT Events Expression](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html):

```go
// Example automatically generated from non-compiling source. May contain errors.
actions.NewSetTimerAction(jsii.String("MyTimer"), map[string]interface{}{
	"durationExpression": iotevents.Expression_inputAttribute(input, jsii.String("payload.durationSeconds")),
})
```

And the timer can be reset and cleared. Below is an example of general
[Device HeartBeat](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-examples-dhb.html)
Detector Model:

```go
// Example automatically generated from non-compiling source. May contain errors.
online := iotevents.NewState(map[string]interface{}{
	"stateName": jsii.String("Online"),
	"onEnter": []map[string]interface{}{
		map[string]interface{}{
			"eventName": jsii.String("enter-event"),
			"condition": iotevents.Expression_currentInput(input),
			"actions": []interface{}{
				actions.NewSetTimerAction(jsii.String("MyTimer"), map[string]interface{}{
					"duration": cdk.Duration_seconds(jsii.Number(60)),
				}),
			},
		},
	},
	"onInput": []map[string]interface{}{
		map[string]interface{}{
			"eventName": jsii.String("input-event"),
			"condition": iotevents.Expression_currentInput(input),
			"actions": []interface{}{
				actions.NewResetTimerAction(jsii.String("MyTimer")),
			},
		},
	},
	"onExit": []map[string]interface{}{
		map[string]interface{}{
			"eventName": jsii.String("exit-event"),
			"actions": []interface{}{
				actions.NewClearTimerAction(jsii.String("MyTimer")),
			},
		},
	},
})
offline := iotevents.NewState(map[string]*string{
	"stateName": jsii.String("Offline"),
})

online.transitionTo(offline, map[string]interface{}{
	"when": iotevents.Expression_timeout(jsii.String("MyTimer")),
})
offline.transitionTo(online, map[string]interface{}{
	"when": iotevents.Expression_currentInput(input),
})
```

## Set variable to detector instance

The code snippet below creates an Action that set variable to detector instance
when it is triggered.

```go
import "github.com/aws/aws-cdk-go/awscdkioteventsalpha"
import actions "github.com/aws/aws-cdk-go/awscdkioteventsactionsalpha"

var input IInput

state := iotevents.NewState(&StateProps{
	StateName: jsii.String("MyState"),
	OnEnter: []Event{
		&Event{
			EventName: jsii.String("test-event"),
			Condition: iotevents.Expression_CurrentInput(input),
			Actions: []IAction{
				actions.NewSetVariableAction(jsii.String("MyVariable"), iotevents.Expression_InputAttribute(input, jsii.String("payload.temperature"))),
			},
		},
	},
})
```

## Invoke a Lambda function

The code snippet below creates an Action that invoke a Lambda function
when it is triggered.

```go
import "github.com/aws/aws-cdk-go/awscdkioteventsalpha"
import actions "github.com/aws/aws-cdk-go/awscdkioteventsactionsalpha"
import lambda "github.com/aws/aws-cdk-go/awscdk"

var input IInput
var func IFunction

state := iotevents.NewState(&StateProps{
	StateName: jsii.String("MyState"),
	OnEnter: []Event{
		&Event{
			EventName: jsii.String("test-event"),
			Condition: iotevents.Expression_CurrentInput(input),
			Actions: []IAction{
				actions.NewLambdaInvokeAction(func),
			},
		},
	},
})
```
