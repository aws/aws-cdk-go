package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Machine image representing the latest NAT instance image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   natInstanceImage := awscdk.Aws_ec2.NewNatInstanceImage()
//
type NatInstanceImage interface {
	LookupMachineImage
	// Return the image to use in the given context.
	GetImage(scope constructs.Construct) *MachineImageConfig
}

// The jsii proxy struct for NatInstanceImage
type jsiiProxy_NatInstanceImage struct {
	jsiiProxy_LookupMachineImage
}

func NewNatInstanceImage() NatInstanceImage {
	_init_.Initialize()

	j := jsiiProxy_NatInstanceImage{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.NatInstanceImage",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewNatInstanceImage_Override(n NatInstanceImage) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.NatInstanceImage",
		nil, // no parameters
		n,
	)
}

func (n *jsiiProxy_NatInstanceImage) GetImage(scope constructs.Construct) *MachineImageConfig {
	if err := n.validateGetImageParameters(scope); err != nil {
		panic(err)
	}
	var returns *MachineImageConfig

	_jsii_.Invoke(
		n,
		"getImage",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

