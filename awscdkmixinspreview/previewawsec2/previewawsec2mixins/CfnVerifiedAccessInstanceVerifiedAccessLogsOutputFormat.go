package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnVerifiedAccessInstanceVerifiedAccessLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVerifiedAccessInstanceVerifiedAccessLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnVerifiedAccessInstanceVerifiedAccessLogsOutputFormat()
//
// Experimental.
type CfnVerifiedAccessInstanceVerifiedAccessLogsOutputFormat interface {
}

// The jsii proxy struct for CfnVerifiedAccessInstanceVerifiedAccessLogsOutputFormat
type jsiiProxy_CfnVerifiedAccessInstanceVerifiedAccessLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnVerifiedAccessInstanceVerifiedAccessLogsOutputFormat() CfnVerifiedAccessInstanceVerifiedAccessLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnVerifiedAccessInstanceVerifiedAccessLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVerifiedAccessInstanceVerifiedAccessLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnVerifiedAccessInstanceVerifiedAccessLogsOutputFormat_Override(c CfnVerifiedAccessInstanceVerifiedAccessLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVerifiedAccessInstanceVerifiedAccessLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

