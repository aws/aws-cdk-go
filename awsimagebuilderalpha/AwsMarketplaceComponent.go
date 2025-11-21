package awsimagebuilderalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsimagebuilderalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Helper class for working with AWS Marketplace components.
//
// Example:
//   marketplaceComponent := imagebuilder.AwsMarketplaceComponent_FromAwsMarketplaceComponentAttributes(this, jsii.String("MarketplaceComponent"), &AwsMarketplaceComponentAttributes{
//   	ComponentName: jsii.String("my-marketplace-component"),
//   	MarketplaceProductId: jsii.String("prod-1234567890abcdef0"),
//   })
//
// Experimental.
type AwsMarketplaceComponent interface {
}

// The jsii proxy struct for AwsMarketplaceComponent
type jsiiProxy_AwsMarketplaceComponent struct {
	_ byte // padding
}

// Experimental.
func NewAwsMarketplaceComponent() AwsMarketplaceComponent {
	_init_.Initialize()

	j := jsiiProxy_AwsMarketplaceComponent{}

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.AwsMarketplaceComponent",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMarketplaceComponent_Override(a AwsMarketplaceComponent) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.AwsMarketplaceComponent",
		nil, // no parameters
		a,
	)
}

// Imports an AWS Marketplace component from its attributes.
// Experimental.
func AwsMarketplaceComponent_FromAwsMarketplaceComponentAttributes(scope constructs.Construct, id *string, attrs *AwsMarketplaceComponentAttributes) IComponent {
	_init_.Initialize()

	if err := validateAwsMarketplaceComponent_FromAwsMarketplaceComponentAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IComponent

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.AwsMarketplaceComponent",
		"fromAwsMarketplaceComponentAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

