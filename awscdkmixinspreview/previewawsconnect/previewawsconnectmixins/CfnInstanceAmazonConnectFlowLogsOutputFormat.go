package previewawsconnectmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnInstanceAmazonConnectFlowLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnInstanceAmazonConnectFlowLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnInstanceAmazonConnectFlowLogsOutputFormat()
//
// Experimental.
type CfnInstanceAmazonConnectFlowLogsOutputFormat interface {
}

// The jsii proxy struct for CfnInstanceAmazonConnectFlowLogsOutputFormat
type jsiiProxy_CfnInstanceAmazonConnectFlowLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnInstanceAmazonConnectFlowLogsOutputFormat() CfnInstanceAmazonConnectFlowLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnInstanceAmazonConnectFlowLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnInstanceAmazonConnectFlowLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnInstanceAmazonConnectFlowLogsOutputFormat_Override(c CfnInstanceAmazonConnectFlowLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnInstanceAmazonConnectFlowLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

