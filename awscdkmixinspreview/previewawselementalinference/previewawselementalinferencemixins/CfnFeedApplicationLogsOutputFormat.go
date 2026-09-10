package previewawselementalinferencemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnFeedApplicationLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFeedApplicationLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnFeedApplicationLogsOutputFormat()
//
// Experimental.
type CfnFeedApplicationLogsOutputFormat interface {
}

// The jsii proxy struct for CfnFeedApplicationLogsOutputFormat
type jsiiProxy_CfnFeedApplicationLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnFeedApplicationLogsOutputFormat() CfnFeedApplicationLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnFeedApplicationLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elementalinference.mixins.CfnFeedApplicationLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnFeedApplicationLogsOutputFormat_Override(c CfnFeedApplicationLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elementalinference.mixins.CfnFeedApplicationLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

