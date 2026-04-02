package previewawseksmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnCapabilityEksCapabilityArgocdApplicationsetLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCapabilityEksCapabilityArgocdApplicationsetLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnCapabilityEksCapabilityArgocdApplicationsetLogsOutputFormat()
//
// Experimental.
type CfnCapabilityEksCapabilityArgocdApplicationsetLogsOutputFormat interface {
}

// The jsii proxy struct for CfnCapabilityEksCapabilityArgocdApplicationsetLogsOutputFormat
type jsiiProxy_CfnCapabilityEksCapabilityArgocdApplicationsetLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnCapabilityEksCapabilityArgocdApplicationsetLogsOutputFormat() CfnCapabilityEksCapabilityArgocdApplicationsetLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnCapabilityEksCapabilityArgocdApplicationsetLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdApplicationsetLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnCapabilityEksCapabilityArgocdApplicationsetLogsOutputFormat_Override(c CfnCapabilityEksCapabilityArgocdApplicationsetLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdApplicationsetLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

