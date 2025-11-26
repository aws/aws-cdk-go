package previewawsbedrockagentcoremixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Builder for CfnGatewayLogsMixin to generate TRACES for CfnGateway.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGatewayTraces := awscdkmixinspreview.Mixins.NewCfnGatewayTraces()
//
type CfnGatewayTraces interface {
	// Send traces to X-Ray.
	ToXRay() CfnGatewayLogsMixin
}

// The jsii proxy struct for CfnGatewayTraces
type jsiiProxy_CfnGatewayTraces struct {
	_ byte // padding
}

// Experimental.
func NewCfnGatewayTraces() CfnGatewayTraces {
	_init_.Initialize()

	j := jsiiProxy_CfnGatewayTraces{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTraces",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnGatewayTraces_Override(c CfnGatewayTraces) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTraces",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnGatewayTraces) ToXRay() CfnGatewayLogsMixin {
	var returns CfnGatewayLogsMixin

	_jsii_.Invoke(
		c,
		"toXRay",
		nil, // no parameters
		&returns,
	)

	return returns
}

