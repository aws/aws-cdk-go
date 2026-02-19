package previewawsbedrockagentcoremixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
)

// Builder for CfnRuntimeLogsMixin to generate TRACES for CfnRuntime.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnRuntimeTraces := awscdkmixinspreview.Mixins.NewCfnRuntimeTraces()
//
type CfnRuntimeTraces interface {
	// Delivers logs to a pre-created delivery destination.
	//
	// Supported destinations are XRAY
	// You are responsible for setting up the correct permissions for your delivery destination, toDestination() does not set up any permissions for you.
	// Delivery destinations that are imported from another stack using CfnDeliveryDestination.fromDeliveryDestinationArn() or CfnDeliveryDestination.fromDeliveryDestinationName() are supported by toDestination().
	ToDestination(destination interfacesawslogs.IDeliveryDestinationRef) CfnRuntimeLogsMixin
	// Send traces to X-Ray.
	ToXRay() CfnRuntimeLogsMixin
}

// The jsii proxy struct for CfnRuntimeTraces
type jsiiProxy_CfnRuntimeTraces struct {
	_ byte // padding
}

// Experimental.
func NewCfnRuntimeTraces() CfnRuntimeTraces {
	_init_.Initialize()

	j := jsiiProxy_CfnRuntimeTraces{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeTraces",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnRuntimeTraces_Override(c CfnRuntimeTraces) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeTraces",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnRuntimeTraces) ToDestination(destination interfacesawslogs.IDeliveryDestinationRef) CfnRuntimeLogsMixin {
	if err := c.validateToDestinationParameters(destination); err != nil {
		panic(err)
	}
	var returns CfnRuntimeLogsMixin

	_jsii_.Invoke(
		c,
		"toDestination",
		[]interface{}{destination},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRuntimeTraces) ToXRay() CfnRuntimeLogsMixin {
	var returns CfnRuntimeLogsMixin

	_jsii_.Invoke(
		c,
		"toXRay",
		nil, // no parameters
		&returns,
	)

	return returns
}

