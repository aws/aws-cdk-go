package previewawselasticloadbalancingv2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnLoadBalancerAlbAccessLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLoadBalancerAlbAccessLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnLoadBalancerAlbAccessLogsOutputFormat()
//
// Experimental.
type CfnLoadBalancerAlbAccessLogsOutputFormat interface {
}

// The jsii proxy struct for CfnLoadBalancerAlbAccessLogsOutputFormat
type jsiiProxy_CfnLoadBalancerAlbAccessLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnLoadBalancerAlbAccessLogsOutputFormat() CfnLoadBalancerAlbAccessLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnLoadBalancerAlbAccessLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbAccessLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnLoadBalancerAlbAccessLogsOutputFormat_Override(c CfnLoadBalancerAlbAccessLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbAccessLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

