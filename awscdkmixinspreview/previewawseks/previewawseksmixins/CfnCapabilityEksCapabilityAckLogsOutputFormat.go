package previewawseksmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnCapabilityEksCapabilityAckLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCapabilityEksCapabilityAckLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnCapabilityEksCapabilityAckLogsOutputFormat()
//
// Experimental.
type CfnCapabilityEksCapabilityAckLogsOutputFormat interface {
}

// The jsii proxy struct for CfnCapabilityEksCapabilityAckLogsOutputFormat
type jsiiProxy_CfnCapabilityEksCapabilityAckLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnCapabilityEksCapabilityAckLogsOutputFormat() CfnCapabilityEksCapabilityAckLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnCapabilityEksCapabilityAckLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityAckLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnCapabilityEksCapabilityAckLogsOutputFormat_Override(c CfnCapabilityEksCapabilityAckLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityAckLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

