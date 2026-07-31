package awsmediaconnectalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Factory class for creating VPC Interface configurations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//
//   vpcInterface := mediaconnect_alpha.NewVpcInterface()
//
// Experimental.
type VpcInterface interface {
}

// The jsii proxy struct for VpcInterface
type jsiiProxy_VpcInterface struct {
	_ byte // padding
}

// Experimental.
func NewVpcInterface() VpcInterface {
	_init_.Initialize()

	j := jsiiProxy_VpcInterface{}

	_jsii_.Create(
		"@aws-cdk/aws-mediaconnect-alpha.VpcInterface",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewVpcInterface_Override(v VpcInterface) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-mediaconnect-alpha.VpcInterface",
		nil, // no parameters
		v,
	)
}

// Define a new VPC Interface configuration.
//
// MediaConnect will create network interfaces automatically.
//
// Returns: VpcInterfaceConfig configuration object.
// Experimental.
func VpcInterface_Define(props *VpcInterfaceDefineProps) *VpcInterfaceConfig {
	_init_.Initialize()

	if err := validateVpcInterface_DefineParameters(props); err != nil {
		panic(err)
	}
	var returns *VpcInterfaceConfig

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.VpcInterface",
		"define",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Create a VPC Interface configuration using existing network interfaces.
//
// Returns: VpcInterfaceConfig configuration object.
// Experimental.
func VpcInterface_FromNetworkInterfaces(props *VpcInterfaceFromNetworkInterfacesProps) *VpcInterfaceConfig {
	_init_.Initialize()

	if err := validateVpcInterface_FromNetworkInterfacesParameters(props); err != nil {
		panic(err)
	}
	var returns *VpcInterfaceConfig

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediaconnect-alpha.VpcInterface",
		"fromNetworkInterfaces",
		[]interface{}{props},
		&returns,
	)

	return returns
}

