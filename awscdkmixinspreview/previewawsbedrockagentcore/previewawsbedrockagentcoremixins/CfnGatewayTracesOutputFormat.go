package previewawsbedrockagentcoremixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Output Format options for each destination of CfnGatewayTraces.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGatewayTracesOutputFormat := awscdkmixinspreview.Mixins.NewCfnGatewayTracesOutputFormat()
//
// Experimental.
type CfnGatewayTracesOutputFormat interface {
}

// The jsii proxy struct for CfnGatewayTracesOutputFormat
type jsiiProxy_CfnGatewayTracesOutputFormat struct {
	_ byte // padding
}

// Experimental.
func NewCfnGatewayTracesOutputFormat() CfnGatewayTracesOutputFormat {
	_init_.Initialize()

	j := jsiiProxy_CfnGatewayTracesOutputFormat{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTracesOutputFormat",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnGatewayTracesOutputFormat_Override(c CfnGatewayTracesOutputFormat) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTracesOutputFormat",
		nil, // no parameters
		c,
	)
}

