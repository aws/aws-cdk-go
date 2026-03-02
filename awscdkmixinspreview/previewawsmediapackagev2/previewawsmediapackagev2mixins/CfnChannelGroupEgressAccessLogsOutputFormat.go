package previewawsmediapackagev2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnChannelGroupEgressAccessLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnChannelGroupEgressAccessLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnChannelGroupEgressAccessLogsOutputFormat()
//
// Experimental.
type CfnChannelGroupEgressAccessLogsOutputFormat interface {
}

// The jsii proxy struct for CfnChannelGroupEgressAccessLogsOutputFormat
type jsiiProxy_CfnChannelGroupEgressAccessLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnChannelGroupEgressAccessLogsOutputFormat() CfnChannelGroupEgressAccessLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnChannelGroupEgressAccessLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupEgressAccessLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnChannelGroupEgressAccessLogsOutputFormat_Override(c CfnChannelGroupEgressAccessLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupEgressAccessLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

