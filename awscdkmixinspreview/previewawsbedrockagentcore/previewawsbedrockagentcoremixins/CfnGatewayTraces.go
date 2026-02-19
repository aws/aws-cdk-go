package previewawsbedrockagentcoremixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
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
	// Delivers logs to a pre-created delivery destination.
	//
	// Supported destinations are XRAY
	// You are responsible for setting up the correct permissions for your delivery destination, toDestination() does not set up any permissions for you.
	// Delivery destinations that are imported from another stack using CfnDeliveryDestination.fromDeliveryDestinationArn() or CfnDeliveryDestination.fromDeliveryDestinationName() are supported by toDestination().
	ToDestination(destination interfacesawslogs.IDeliveryDestinationRef) CfnGatewayLogsMixin
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

func (c *jsiiProxy_CfnGatewayTraces) ToDestination(destination interfacesawslogs.IDeliveryDestinationRef) CfnGatewayLogsMixin {
	if err := c.validateToDestinationParameters(destination); err != nil {
		panic(err)
	}
	var returns CfnGatewayLogsMixin

	_jsii_.Invoke(
		c,
		"toDestination",
		[]interface{}{destination},
		&returns,
	)

	return returns
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

