package previewawslogsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.logs@AWSAPICallViaCloudTrail event types for LogGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aWSAPICallViaCloudTrail := #error#.NewAWSAPICallViaCloudTrail()
//
// Experimental.
type LogGroupEvents_AWSAPICallViaCloudTrail interface {
}

// The jsii proxy struct for LogGroupEvents_AWSAPICallViaCloudTrail
type jsiiProxy_LogGroupEvents_AWSAPICallViaCloudTrail struct {
	_ byte // padding
}

// Experimental.
func NewLogGroupEvents_AWSAPICallViaCloudTrail() LogGroupEvents_AWSAPICallViaCloudTrail {
	_init_.Initialize()

	j := jsiiProxy_LogGroupEvents_AWSAPICallViaCloudTrail{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_logs.events.LogGroupEvents.AWSAPICallViaCloudTrail",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewLogGroupEvents_AWSAPICallViaCloudTrail_Override(l LogGroupEvents_AWSAPICallViaCloudTrail) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_logs.events.LogGroupEvents.AWSAPICallViaCloudTrail",
		nil, // no parameters
		l,
	)
}

