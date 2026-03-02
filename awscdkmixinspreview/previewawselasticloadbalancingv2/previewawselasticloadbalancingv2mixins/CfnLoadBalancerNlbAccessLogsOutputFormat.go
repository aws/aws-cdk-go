package previewawselasticloadbalancingv2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnLoadBalancerNlbAccessLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLoadBalancerNlbAccessLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnLoadBalancerNlbAccessLogsOutputFormat()
//
// Experimental.
type CfnLoadBalancerNlbAccessLogsOutputFormat interface {
}

// The jsii proxy struct for CfnLoadBalancerNlbAccessLogsOutputFormat
type jsiiProxy_CfnLoadBalancerNlbAccessLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnLoadBalancerNlbAccessLogsOutputFormat() CfnLoadBalancerNlbAccessLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnLoadBalancerNlbAccessLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerNlbAccessLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnLoadBalancerNlbAccessLogsOutputFormat_Override(c CfnLoadBalancerNlbAccessLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerNlbAccessLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

