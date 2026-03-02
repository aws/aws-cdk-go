package previewawsapsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnScraperApplicationLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnScraperApplicationLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnScraperApplicationLogsOutputFormat()
//
// Experimental.
type CfnScraperApplicationLogsOutputFormat interface {
}

// The jsii proxy struct for CfnScraperApplicationLogsOutputFormat
type jsiiProxy_CfnScraperApplicationLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnScraperApplicationLogsOutputFormat() CfnScraperApplicationLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnScraperApplicationLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnScraperApplicationLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnScraperApplicationLogsOutputFormat_Override(c CfnScraperApplicationLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_aps.mixins.CfnScraperApplicationLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

