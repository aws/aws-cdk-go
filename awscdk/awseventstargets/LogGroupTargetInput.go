package awseventstargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// The input to send to the CloudWatch LogGroup target.
//
// Example:
//   import logs "github.com/aws/aws-cdk-go/awscdk"
//   var logGroup logGroup
//   var rule rule
//
//
//   rule.AddTarget(targets.NewCloudWatchLogGroup(logGroup, &LogGroupProps{
//   	LogEvent: targets.LogGroupTargetInput_FromObject(&LogGroupTargetInputOptions{
//   		Timestamp: events.EventField_FromPath(jsii.String("$.time")),
//   		Message: events.EventField_*FromPath(jsii.String("$.detail-type")),
//   	}),
//   }))
//
type LogGroupTargetInput interface {
	// Return the input properties for this input object.
	Bind(rule awsevents.IRule) *awsevents.RuleTargetInputProperties
}

// The jsii proxy struct for LogGroupTargetInput
type jsiiProxy_LogGroupTargetInput struct {
	_ byte // padding
}

func NewLogGroupTargetInput_Override(l LogGroupTargetInput) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.LogGroupTargetInput",
		nil, // no parameters
		l,
	)
}

// Pass a JSON object to the the log group event target.
//
// May contain strings returned by `EventField.from()` to substitute in parts of the
// matched event.
func LogGroupTargetInput_FromObject(options *LogGroupTargetInputOptions) awsevents.RuleTargetInput {
	_init_.Initialize()

	if err := validateLogGroupTargetInput_FromObjectParameters(options); err != nil {
		panic(err)
	}
	var returns awsevents.RuleTargetInput

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_events_targets.LogGroupTargetInput",
		"fromObject",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogGroupTargetInput) Bind(rule awsevents.IRule) *awsevents.RuleTargetInputProperties {
	if err := l.validateBindParameters(rule); err != nil {
		panic(err)
	}
	var returns *awsevents.RuleTargetInputProperties

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{rule},
		&returns,
	)

	return returns
}

