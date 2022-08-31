package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// Represents the Wait state which delays a state machine from continuing for a specified time.
//
// Example:
//   convertToSeconds := tasks.NewEvaluateExpression(this, jsii.String("Convert to seconds"), &evaluateExpressionProps{
//   	expression: jsii.String("$.waitMilliseconds / 1000"),
//   	resultPath: jsii.String("$.waitSeconds"),
//   })
//
//   createMessage := tasks.NewEvaluateExpression(this, jsii.String("Create message"), &evaluateExpressionProps{
//   	// Note: this is a string inside a string.
//   	expression: jsii.String("`Now waiting ${$.waitSeconds} seconds...`"),
//   	runtime: lambda.runtime_NODEJS_14_X(),
//   	resultPath: jsii.String("$.message"),
//   })
//
//   publishMessage := tasks.NewSnsPublish(this, jsii.String("Publish message"), &snsPublishProps{
//   	topic: sns.NewTopic(this, jsii.String("cool-topic")),
//   	message: sfn.taskInput.fromJsonPathAt(jsii.String("$.message")),
//   	resultPath: jsii.String("$.sns"),
//   })
//
//   wait := sfn.NewWait(this, jsii.String("Wait"), &waitProps{
//   	time: sfn.waitTime.secondsPath(jsii.String("$.waitSeconds")),
//   })
//
//   sfn.NewStateMachine(this, jsii.String("StateMachine"), &stateMachineProps{
//   	definition: convertToSeconds.next(createMessage).next(publishMessage).next(wait),
//   })
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-wait-state.html
//
// Experimental.
type WaitTime interface {
}

// The jsii proxy struct for WaitTime
type jsiiProxy_WaitTime struct {
	_ byte // padding
}

// Wait a fixed amount of time.
// Experimental.
func WaitTime_Duration(duration awscdk.Duration) WaitTime {
	_init_.Initialize()

	var returns WaitTime

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.WaitTime",
		"duration",
		[]interface{}{duration},
		&returns,
	)

	return returns
}

// Wait for a number of seconds stored in the state object.
//
// Example value: `$.waitSeconds`
// Experimental.
func WaitTime_SecondsPath(path *string) WaitTime {
	_init_.Initialize()

	var returns WaitTime

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.WaitTime",
		"secondsPath",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Wait until the given ISO8601 timestamp.
//
// Example value: `2016-03-14T01:59:00Z`.
// Experimental.
func WaitTime_Timestamp(timestamp *string) WaitTime {
	_init_.Initialize()

	var returns WaitTime

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.WaitTime",
		"timestamp",
		[]interface{}{timestamp},
		&returns,
	)

	return returns
}

// Wait until a timestamp found in the state object.
//
// Example value: `$.waitTimestamp`
// Experimental.
func WaitTime_TimestampPath(path *string) WaitTime {
	_init_.Initialize()

	var returns WaitTime

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.WaitTime",
		"timestampPath",
		[]interface{}{path},
		&returns,
	)

	return returns
}

