package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnRouteServerPeerEventLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRouteServerPeerEventLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnRouteServerPeerEventLogsOutputFormat()
//
// Experimental.
type CfnRouteServerPeerEventLogsOutputFormat interface {
}

// The jsii proxy struct for CfnRouteServerPeerEventLogsOutputFormat
type jsiiProxy_CfnRouteServerPeerEventLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnRouteServerPeerEventLogsOutputFormat() CfnRouteServerPeerEventLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnRouteServerPeerEventLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnRouteServerPeerEventLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnRouteServerPeerEventLogsOutputFormat_Override(c CfnRouteServerPeerEventLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnRouteServerPeerEventLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

