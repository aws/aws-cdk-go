package previewawsmediapackagev2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnChannelGroupIngressAccessLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnChannelGroupIngressAccessLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnChannelGroupIngressAccessLogsOutputFormat()
//
// Experimental.
type CfnChannelGroupIngressAccessLogsOutputFormat interface {
}

// The jsii proxy struct for CfnChannelGroupIngressAccessLogsOutputFormat
type jsiiProxy_CfnChannelGroupIngressAccessLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnChannelGroupIngressAccessLogsOutputFormat() CfnChannelGroupIngressAccessLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnChannelGroupIngressAccessLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupIngressAccessLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnChannelGroupIngressAccessLogsOutputFormat_Override(c CfnChannelGroupIngressAccessLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupIngressAccessLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

