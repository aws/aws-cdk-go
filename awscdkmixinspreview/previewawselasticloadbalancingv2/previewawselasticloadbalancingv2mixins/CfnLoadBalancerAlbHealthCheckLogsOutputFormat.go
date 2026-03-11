package previewawselasticloadbalancingv2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnLoadBalancerAlbHealthCheckLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLoadBalancerAlbHealthCheckLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnLoadBalancerAlbHealthCheckLogsOutputFormat()
//
// Experimental.
type CfnLoadBalancerAlbHealthCheckLogsOutputFormat interface {
}

// The jsii proxy struct for CfnLoadBalancerAlbHealthCheckLogsOutputFormat
type jsiiProxy_CfnLoadBalancerAlbHealthCheckLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnLoadBalancerAlbHealthCheckLogsOutputFormat() CfnLoadBalancerAlbHealthCheckLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnLoadBalancerAlbHealthCheckLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbHealthCheckLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnLoadBalancerAlbHealthCheckLogsOutputFormat_Override(c CfnLoadBalancerAlbHealthCheckLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbHealthCheckLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

