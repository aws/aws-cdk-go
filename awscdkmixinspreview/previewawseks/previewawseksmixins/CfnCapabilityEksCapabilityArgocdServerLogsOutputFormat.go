package previewawseksmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnCapabilityEksCapabilityArgocdServerLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCapabilityEksCapabilityArgocdServerLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnCapabilityEksCapabilityArgocdServerLogsOutputFormat()
//
// Experimental.
type CfnCapabilityEksCapabilityArgocdServerLogsOutputFormat interface {
}

// The jsii proxy struct for CfnCapabilityEksCapabilityArgocdServerLogsOutputFormat
type jsiiProxy_CfnCapabilityEksCapabilityArgocdServerLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnCapabilityEksCapabilityArgocdServerLogsOutputFormat() CfnCapabilityEksCapabilityArgocdServerLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnCapabilityEksCapabilityArgocdServerLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdServerLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnCapabilityEksCapabilityArgocdServerLogsOutputFormat_Override(c CfnCapabilityEksCapabilityArgocdServerLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityArgocdServerLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

