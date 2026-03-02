package previewawseksmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnClusterAutoModeLoadBalancingLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterAutoModeLoadBalancingLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnClusterAutoModeLoadBalancingLogsOutputFormat()
//
// Experimental.
type CfnClusterAutoModeLoadBalancingLogsOutputFormat interface {
}

// The jsii proxy struct for CfnClusterAutoModeLoadBalancingLogsOutputFormat
type jsiiProxy_CfnClusterAutoModeLoadBalancingLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnClusterAutoModeLoadBalancingLogsOutputFormat() CfnClusterAutoModeLoadBalancingLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnClusterAutoModeLoadBalancingLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeLoadBalancingLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnClusterAutoModeLoadBalancingLogsOutputFormat_Override(c CfnClusterAutoModeLoadBalancingLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_eks.mixins.CfnClusterAutoModeLoadBalancingLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

