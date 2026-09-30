package previewawspcsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnClusterPcsScalingLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterPcsScalingLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnClusterPcsScalingLogsOutputFormat()
//
// Experimental.
type CfnClusterPcsScalingLogsOutputFormat interface {
}

// The jsii proxy struct for CfnClusterPcsScalingLogsOutputFormat
type jsiiProxy_CfnClusterPcsScalingLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnClusterPcsScalingLogsOutputFormat() CfnClusterPcsScalingLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnClusterPcsScalingLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsScalingLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnClusterPcsScalingLogsOutputFormat_Override(c CfnClusterPcsScalingLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pcs.mixins.CfnClusterPcsScalingLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

