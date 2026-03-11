package previewawselasticloadbalancingv2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnLoadBalancerAlbConnectionLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLoadBalancerAlbConnectionLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnLoadBalancerAlbConnectionLogsOutputFormat()
//
// Experimental.
type CfnLoadBalancerAlbConnectionLogsOutputFormat interface {
}

// The jsii proxy struct for CfnLoadBalancerAlbConnectionLogsOutputFormat
type jsiiProxy_CfnLoadBalancerAlbConnectionLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnLoadBalancerAlbConnectionLogsOutputFormat() CfnLoadBalancerAlbConnectionLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnLoadBalancerAlbConnectionLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbConnectionLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnLoadBalancerAlbConnectionLogsOutputFormat_Override(c CfnLoadBalancerAlbConnectionLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbConnectionLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

