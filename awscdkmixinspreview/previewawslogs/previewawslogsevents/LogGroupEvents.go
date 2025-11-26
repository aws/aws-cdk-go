package previewawslogsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
)

// EventBridge event patterns for LogGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var logGroupRef ILogGroupRef
//
//   logGroupEvents := awscdkmixinspreview.Events.LogGroupEvents_FromLogGroup(logGroupRef)
//
// Experimental.
type LogGroupEvents interface {
	// EventBridge event pattern for LogGroup AWS API Call via CloudTrail.
	// Experimental.
	AwsAPICallViaCloudTrailPattern(options *LogGroupEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) *awsevents.EventPattern
}

// The jsii proxy struct for LogGroupEvents
type jsiiProxy_LogGroupEvents struct {
	_ byte // padding
}

// Create LogGroupEvents from a LogGroup reference.
// Experimental.
func LogGroupEvents_FromLogGroup(logGroupRef interfacesawslogs.ILogGroupRef) LogGroupEvents {
	_init_.Initialize()

	if err := validateLogGroupEvents_FromLogGroupParameters(logGroupRef); err != nil {
		panic(err)
	}
	var returns LogGroupEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_logs.events.LogGroupEvents",
		"fromLogGroup",
		[]interface{}{logGroupRef},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogGroupEvents) AwsAPICallViaCloudTrailPattern(options *LogGroupEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) *awsevents.EventPattern {
	if err := l.validateAwsAPICallViaCloudTrailPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		l,
		"awsAPICallViaCloudTrailPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

