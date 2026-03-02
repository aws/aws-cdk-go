package previewawseksmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnClusterAutoModeComputeLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterAutoModeComputeLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnClusterAutoModeComputeLogsOutputFormat()
//
// Experimental.
type CfnClusterAutoModeComputeLogsOutputFormat interface {
}

// The jsii proxy struct for CfnClusterAutoModeComputeLogsOutputFormat
type jsiiProxy_CfnClusterAutoModeComputeLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnClusterAutoModeComputeLogsOutputFormat() CfnClusterAutoModeComputeLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnClusterAutoModeComputeLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeComputeLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnClusterAutoModeComputeLogsOutputFormat_Override(c CfnClusterAutoModeComputeLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeComputeLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

