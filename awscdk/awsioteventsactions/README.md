# Actions for AWS::IoTEvents Detector Model

This library contains integration classes to specify actions of state events of Detector Model in `@aws-cdk/aws-iotevents`.
Instances of these classes should be passed to `State` defined in `@aws-cdk/aws-iotevents`
You can define built-in actions to use a timer or set a variable, or send data to other AWS resources.

This library contains integration classes to use a timer or set a variable, or send data to other AWS resources.
AWS IoT Events can trigger actions when it detects a specified event or transition event.

Currently supported are:

* Set variable to detector instanse
* Invoke a Lambda function

## Set variable to detector instanse

The code snippet below creates an Action that set variable to detector instanse
when it is triggered.

```go
// Example automatically generated from non-compiling source. May contain errors.
import iotevents "github.com/aws/aws-cdk-go/awscdk"
import actions "github.com/aws/aws-cdk-go/awscdk"

var input iInput


state := iotevents.NewState(&stateProps{
	stateName: jsii.String("MyState"),
	onEnter: []event{
		&event{
			eventName: jsii.String("test-event"),
			condition: iotevents.expression.currentInput(input),
			actions: []iAction{
				actions,
				[]interface{}{
					actions.NewSetVariableAction(jsii.String("MyVariable"), iotevents.*expression.inputAttribute(input, jsii.String("payload.temperature"))),
				},
			},
		},
	},
})
```

## Invoke a Lambda function

The code snippet below creates an Action that invoke a Lambda function
when it is triggered.

```go
// Example automatically generated from non-compiling source. May contain errors.
import iotevents "github.com/aws/aws-cdk-go/awscdk"
import actions "github.com/aws/aws-cdk-go/awscdk"
import lambda "github.com/aws/aws-cdk-go/awscdk"

var input iInput
var func iFunction


state := iotevents.NewState(&stateProps{
	stateName: jsii.String("MyState"),
	onEnter: []event{
		&event{
			eventName: jsii.String("test-event"),
			condition: iotevents.expression.currentInput(input),
			actions: []iAction{
				actions.NewLambdaInvokeAction(func),
			},
		},
	},
})
```
