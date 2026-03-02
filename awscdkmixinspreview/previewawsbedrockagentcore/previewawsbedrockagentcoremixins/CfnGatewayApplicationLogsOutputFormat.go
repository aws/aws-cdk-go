package previewawsbedrockagentcoremixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnGatewayApplicationLogs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGatewayApplicationLogsOutputFormat := awscdkmixinspreview.Mixins.NewCfnGatewayApplicationLogsOutputFormat()
//
// Experimental.
type CfnGatewayApplicationLogsOutputFormat interface {
}

// The jsii proxy struct for CfnGatewayApplicationLogsOutputFormat
type jsiiProxy_CfnGatewayApplicationLogsOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnGatewayApplicationLogsOutputFormat() CfnGatewayApplicationLogsOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnGatewayApplicationLogsOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayApplicationLogsOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnGatewayApplicationLogsOutputFormat_Override(c CfnGatewayApplicationLogsOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayApplicationLogsOutputFormat",
		nil, // no parameters
		c,
	)
}

