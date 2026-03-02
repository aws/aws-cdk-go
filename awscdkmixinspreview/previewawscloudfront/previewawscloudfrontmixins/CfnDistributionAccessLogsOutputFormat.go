package previewawscloudfrontmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnDistributionAccessLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDistributionAccessLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnDistributionAccessLogsOutputFormat()
//
// Experimental.
type CfnDistributionAccessLogsOutputFormat interface {
}

// The jsii proxy struct for CfnDistributionAccessLogsOutputFormat
type jsiiProxy_CfnDistributionAccessLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnDistributionAccessLogsOutputFormat() CfnDistributionAccessLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnDistributionAccessLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionAccessLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnDistributionAccessLogsOutputFormat_Override(c CfnDistributionAccessLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnDistributionAccessLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

