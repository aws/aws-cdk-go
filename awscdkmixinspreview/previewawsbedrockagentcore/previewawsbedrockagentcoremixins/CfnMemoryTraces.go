package previewawsbedrockagentcoremixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
)

// Builder for CfnMemoryLogsMixin to generate TRACES for CfnMemory.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMemoryTraces := awscdkmixinspreview.Mixins.NewCfnMemoryTraces()
//
type CfnMemoryTraces interface {
	// Delivers logs to a pre-created delivery destination.
	//
	// Supported destinations are XRAY
	// You are responsible for setting up the correct permissions for your delivery destination, toDestination() does not set up any permissions for you.
	// Delivery destinations that are imported from another stack using CfnDeliveryDestination.fromDeliveryDestinationArn() or CfnDeliveryDestination.fromDeliveryDestinationName() are supported by toDestination().
	ToDestination(destination interfacesawslogs.IDeliveryDestinationRef) CfnMemoryLogsMixin
	// Send traces to X-Ray.
	ToXRay() CfnMemoryLogsMixin
}

// The jsii proxy struct for CfnMemoryTraces
type jsiiProxy_CfnMemoryTraces struct {
	_ byte // padding
}

// Experimental.
func NewCfnMemoryTraces() CfnMemoryTraces {
	_init_.Initialize()

	j := jsiiProxy_CfnMemoryTraces{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryTraces",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnMemoryTraces_Override(c CfnMemoryTraces) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryTraces",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnMemoryTraces) ToDestination(destination interfacesawslogs.IDeliveryDestinationRef) CfnMemoryLogsMixin {
	if err := c.validateToDestinationParameters(destination); err != nil {
		panic(err)
	}
	var returns CfnMemoryLogsMixin

	_jsii_.Invoke(
		c,
		"toDestination",
		[]interface{}{destination},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMemoryTraces) ToXRay() CfnMemoryLogsMixin {
	var returns CfnMemoryLogsMixin

	_jsii_.Invoke(
		c,
		"toXRay",
		nil, // no parameters
		&returns,
	)

	return returns
}

