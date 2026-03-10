package previewawscloudwatchevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.cloudwatch@CloudWatchAlarmConfigurationChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cloudWatchAlarmConfigurationChange := awscdkmixinspreview.Events.NewCloudWatchAlarmConfigurationChange()
//
// Experimental.
type CloudWatchAlarmConfigurationChange interface {
}

// The jsii proxy struct for CloudWatchAlarmConfigurationChange
type jsiiProxy_CloudWatchAlarmConfigurationChange struct {
	_ byte // padding
}

// Experimental.
func NewCloudWatchAlarmConfigurationChange() CloudWatchAlarmConfigurationChange {
	_init_.Initialize()

	j := jsiiProxy_CloudWatchAlarmConfigurationChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.CloudWatchAlarmConfigurationChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCloudWatchAlarmConfigurationChange_Override(c CloudWatchAlarmConfigurationChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.CloudWatchAlarmConfigurationChange",
		nil, // no parameters
		c,
	)
}

// EventBridge event pattern for CloudWatch Alarm Configuration Change.
// Experimental.
func CloudWatchAlarmConfigurationChange_CloudWatchAlarmConfigurationChangePattern(options *CloudWatchAlarmConfigurationChange_CloudWatchAlarmConfigurationChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateCloudWatchAlarmConfigurationChange_CloudWatchAlarmConfigurationChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cloudwatch.events.CloudWatchAlarmConfigurationChange",
		"cloudWatchAlarmConfigurationChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

