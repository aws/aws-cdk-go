package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// Construct a Windows machine image from an AMI map.
//
// Allows you to create a generic Windows EC2 , manually specify an AMI map.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var userData userData
//
//   genericWindowsImage := awscdk.Aws_ec2.NewGenericWindowsImage(map[string]*string{
//   	"amiMapKey": jsii.String("amiMap"),
//   }, &genericWindowsImageProps{
//   	userData: userData,
//   })
//
// Experimental.
type GenericWindowsImage interface {
	IMachineImage
	// Return the image to use in the given context.
	// Experimental.
	GetImage(scope awscdk.Construct) *MachineImageConfig
}

// The jsii proxy struct for GenericWindowsImage
type jsiiProxy_GenericWindowsImage struct {
	jsiiProxy_IMachineImage
}

// Experimental.
func NewGenericWindowsImage(amiMap *map[string]*string, props *GenericWindowsImageProps) GenericWindowsImage {
	_init_.Initialize()

	j := jsiiProxy_GenericWindowsImage{}

	_jsii_.Create(
		"monocdk.aws_ec2.GenericWindowsImage",
		[]interface{}{amiMap, props},
		&j,
	)

	return &j
}

// Experimental.
func NewGenericWindowsImage_Override(g GenericWindowsImage, amiMap *map[string]*string, props *GenericWindowsImageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ec2.GenericWindowsImage",
		[]interface{}{amiMap, props},
		g,
	)
}

func (g *jsiiProxy_GenericWindowsImage) GetImage(scope awscdk.Construct) *MachineImageConfig {
	var returns *MachineImageConfig

	_jsii_.Invoke(
		g,
		"getImage",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

