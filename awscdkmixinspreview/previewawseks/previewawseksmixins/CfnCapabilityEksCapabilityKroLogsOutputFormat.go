package previewawseksmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnCapabilityEksCapabilityKroLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCapabilityEksCapabilityKroLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnCapabilityEksCapabilityKroLogsOutputFormat()
//
// Experimental.
type CfnCapabilityEksCapabilityKroLogsOutputFormat interface {
}

// The jsii proxy struct for CfnCapabilityEksCapabilityKroLogsOutputFormat
type jsiiProxy_CfnCapabilityEksCapabilityKroLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnCapabilityEksCapabilityKroLogsOutputFormat() CfnCapabilityEksCapabilityKroLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnCapabilityEksCapabilityKroLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityKroLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnCapabilityEksCapabilityKroLogsOutputFormat_Override(c CfnCapabilityEksCapabilityKroLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityKroLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

