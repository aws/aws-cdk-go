package previewawscloudwatchevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.cloudwatch@CloudWatchAlarmStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cloudWatchAlarmStateChange := awscdkmixinspreview.Events.NewCloudWatchAlarmStateChange()
//
// Experimental.
type CloudWatchAlarmStateChange interface {
}

// The jsii proxy struct for CloudWatchAlarmStateChange
type jsiiProxy_CloudWatchAlarmStateChange struct {
	_ byte // padding
}

// Experimental.
func NewCloudWatchAlarmStateChange() CloudWatchAlarmStateChange {
	_init_.Initialize()

	j := jsiiProxy_CloudWatchAlarmStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.CloudWatchAlarmStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCloudWatchAlarmStateChange_Override(c CloudWatchAlarmStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.CloudWatchAlarmStateChange",
		nil, // no parameters
		c,
	)
}

// EventBridge event pattern for CloudWatch Alarm State Change.
// Experimental.
func CloudWatchAlarmStateChange_EventPattern(options *CloudWatchAlarmStateChange_CloudWatchAlarmStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateCloudWatchAlarmStateChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.CloudWatchAlarmStateChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

