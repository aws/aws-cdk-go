package previewawseksmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnCapabilityEksCapabilityArgocdApplicationLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCapabilityEksCapabilityArgocdApplicationLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnCapabilityEksCapabilityArgocdApplicationLogsOutputFormat()
//
// Experimental.
type CfnCapabilityEksCapabilityArgocdApplicationLogsOutputFormat interface {
}

// The jsii proxy struct for CfnCapabilityEksCapabilityArgocdApplicationLogsOutputFormat
type jsiiProxy_CfnCapabilityEksCapabilityArgocdApplicationLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnCapabilityEksCapabilityArgocdApplicationLogsOutputFormat() CfnCapabilityEksCapabilityArgocdApplicationLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnCapabilityEksCapabilityArgocdApplicationLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdApplicationLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnCapabilityEksCapabilityArgocdApplicationLogsOutputFormat_Override(c CfnCapabilityEksCapabilityArgocdApplicationLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdApplicationLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

