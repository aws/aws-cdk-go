package previewawsbedrockmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
)

// Builder for CfnKnowledgeBaseLogsMixin to generate TRACES for CfnKnowledgeBase.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnKnowledgeBaseTraces := awscdkmixinspreview.Mixins.NewCfnKnowledgeBaseTraces()
//
type CfnKnowledgeBaseTraces interface {
	// Delivers logs to a pre-created delivery destination.
	//
	// Supported destinations are XRAY
	// You are responsible for setting up the correct permissions for your delivery destination, toDestination() does not set up any permissions for you.
	// Delivery destinations that are imported from another stack using CfnDeliveryDestination.fromDeliveryDestinationArn() or CfnDeliveryDestination.fromDeliveryDestinationName() are supported by toDestination().
	ToDestination(destination interfacesawslogs.IDeliveryDestinationRef) CfnKnowledgeBaseLogsMixin
	// Send traces to X-Ray.
	ToXRay() CfnKnowledgeBaseLogsMixin
}

// The jsii proxy struct for CfnKnowledgeBaseTraces
type jsiiProxy_CfnKnowledgeBaseTraces struct {
	_ byte // padding
}

// Experimental.
func NewCfnKnowledgeBaseTraces() CfnKnowledgeBaseTraces {
	_init_.Initialize()

	j := jsiiProxy_CfnKnowledgeBaseTraces{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseTraces",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnKnowledgeBaseTraces_Override(c CfnKnowledgeBaseTraces) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseTraces",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnKnowledgeBaseTraces) ToDestination(destination interfacesawslogs.IDeliveryDestinationRef) CfnKnowledgeBaseLogsMixin {
	if err := c.validateToDestinationParameters(destination); err != nil {
		panic(err)
	}
	var returns CfnKnowledgeBaseLogsMixin

	_jsii_.Invoke(
		c,
		"toDestination",
		[]interface{}{destination},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnKnowledgeBaseTraces) ToXRay() CfnKnowledgeBaseLogsMixin {
	var returns CfnKnowledgeBaseLogsMixin

	_jsii_.Invoke(
		c,
		"toXRay",
		nil, // no parameters
		&returns,
	)

	return returns
}

