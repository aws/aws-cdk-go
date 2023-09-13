package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Represents the Wait state which delays a state machine from continuing for a specified time.
//
// Example:
//   convertToSeconds := tasks.NewEvaluateExpression(this, jsii.String("Convert to seconds"), &EvaluateExpressionProps{
//   	Expression: jsii.String("$.waitMilliseconds / 1000"),
//   	ResultPath: jsii.String("$.waitSeconds"),
//   })
//
//   createMessage := tasks.NewEvaluateExpression(this, jsii.String("Create message"), &EvaluateExpressionProps{
//   	// Note: this is a string inside a string.
//   	Expression: jsii.String("`Now waiting ${$.waitSeconds} seconds...`"),
//   	Runtime: lambda.Runtime_NODEJS_LATEST(),
//   	ResultPath: jsii.String("$.message"),
//   })
//
//   publishMessage := tasks.NewSnsPublish(this, jsii.String("Publish message"), &SnsPublishProps{
//   	Topic: sns.NewTopic(this, jsii.String("cool-topic")),
//   	Message: sfn.TaskInput_FromJsonPathAt(jsii.String("$.message")),
//   	ResultPath: jsii.String("$.sns"),
//   })
//
//   wait := sfn.NewWait(this, jsii.String("Wait"), &WaitProps{
//   	Time: sfn.WaitTime_SecondsPath(jsii.String("$.waitSeconds")),
//   })
//
//   sfn.NewStateMachine(this, jsii.String("StateMachine"), &StateMachineProps{
//   	Definition: convertToSeconds.Next(createMessage).Next(publishMessage).*Next(wait),
//   })
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-wait-state.html
//
type WaitTime interface {
}

// The jsii proxy struct for WaitTime
type jsiiProxy_WaitTime struct {
	_ byte // padding
}

// Wait a fixed amount of time.
func WaitTime_Duration(duration awscdk.Duration) WaitTime {
	_init_.Initialize()

	if err := validateWaitTime_DurationParameters(duration); err != nil {
		panic(err)
	}
	var returns WaitTime

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.WaitTime",
		"duration",
		[]interface{}{duration},
		&returns,
	)

	return returns
}

// Wait for a number of seconds stored in the state object.
//
// Example value: `$.waitSeconds`
func WaitTime_SecondsPath(path *string) WaitTime {
	_init_.Initialize()

	if err := validateWaitTime_SecondsPathParameters(path); err != nil {
		panic(err)
	}
	var returns WaitTime

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.WaitTime",
		"secondsPath",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Wait until the given ISO8601 timestamp.
//
// Example value: `2016-03-14T01:59:00Z`.
func WaitTime_Timestamp(timestamp *string) WaitTime {
	_init_.Initialize()

	if err := validateWaitTime_TimestampParameters(timestamp); err != nil {
		panic(err)
	}
	var returns WaitTime

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.WaitTime",
		"timestamp",
		[]interface{}{timestamp},
		&returns,
	)

	return returns
}

// Wait until a timestamp found in the state object.
//
// Example value: `$.waitTimestamp`
func WaitTime_TimestampPath(path *string) WaitTime {
	_init_.Initialize()

	if err := validateWaitTime_TimestampPathParameters(path); err != nil {
		panic(err)
	}
	var returns WaitTime

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.WaitTime",
		"timestampPath",
		[]interface{}{path},
		&returns,
	)

	return returns
}

