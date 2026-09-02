package previewawsbedrockagentcoremixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
)

// Builder for CfnBrowserCustomLogsMixin to generate TRACES for CfnBrowserCustom.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBrowserCustomTraces := awscdkmixinspreview.Mixins.NewCfnBrowserCustomTraces()
//
type CfnBrowserCustomTraces interface {
	// Delivers logs to a pre-created delivery destination.
	//
	// Supported destinations are XRAY
	// You are responsible for setting up the correct permissions for your delivery destination, toDestination() does not set up any permissions for you.
	// Delivery destinations that are imported from another stack using CfnDeliveryDestination.fromDeliveryDestinationArn() or CfnDeliveryDestination.fromDeliveryDestinationName() are supported by toDestination().
	ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnBrowserCustomTracesDestProps) CfnBrowserCustomLogsMixin
	// Send traces to X-Ray.
	ToXRay(props *CfnBrowserCustomTracesXRayProps) CfnBrowserCustomLogsMixin
}

// The jsii proxy struct for CfnBrowserCustomTraces
type jsiiProxy_CfnBrowserCustomTraces struct {
	_ byte // padding
}

// Experimental.
func NewCfnBrowserCustomTraces() CfnBrowserCustomTraces {
	_init_.Initialize()

	j := jsiiProxy_CfnBrowserCustomTraces{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomTraces",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnBrowserCustomTraces_Override(c CfnBrowserCustomTraces) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomTraces",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnBrowserCustomTraces) ToDestination(destination interfacesawslogs.IDeliveryDestinationRef, props *CfnBrowserCustomTracesDestProps) CfnBrowserCustomLogsMixin {
	if err := c.validateToDestinationParameters(destination, props); err != nil {
		panic(err)
	}
	var returns CfnBrowserCustomLogsMixin

	_jsii_.Invoke(
		c,
		"toDestination",
		[]interface{}{destination, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnBrowserCustomTraces) ToXRay(props *CfnBrowserCustomTracesXRayProps) CfnBrowserCustomLogsMixin {
	if err := c.validateToXRayParameters(props); err != nil {
		panic(err)
	}
	var returns CfnBrowserCustomLogsMixin

	_jsii_.Invoke(
		c,
		"toXRay",
		[]interface{}{props},
		&returns,
	)

	return returns
}

