package previewawseksmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnCapabilityEksCapabilityAckS3Logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCapabilityEksCapabilityAckS3LogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnCapabilityEksCapabilityAckS3LogsOutputFormat()
//
// Experimental.
type CfnCapabilityEksCapabilityAckS3LogsOutputFormat interface {
}

// The jsii proxy struct for CfnCapabilityEksCapabilityAckS3LogsOutputFormat
type jsiiProxy_CfnCapabilityEksCapabilityAckS3LogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnCapabilityEksCapabilityAckS3LogsOutputFormat() CfnCapabilityEksCapabilityAckS3LogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnCapabilityEksCapabilityAckS3LogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityAckS3LogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnCapabilityEksCapabilityAckS3LogsOutputFormat_Override(c CfnCapabilityEksCapabilityAckS3LogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnCapabilityEksCapabilityAckS3LogsOutputFormat",
		nil, // no parameters
		c,
	)
}

