package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
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
//   }, &GenericWindowsImageProps{
//   	UserData: userData,
//   })
//
type GenericWindowsImage interface {
	IMachineImage
	// Return the image to use in the given context.
	GetImage(scope constructs.Construct) *MachineImageConfig
}

// The jsii proxy struct for GenericWindowsImage
type jsiiProxy_GenericWindowsImage struct {
	jsiiProxy_IMachineImage
}

func NewGenericWindowsImage(amiMap *map[string]*string, props *GenericWindowsImageProps) GenericWindowsImage {
	_init_.Initialize()

	if err := validateNewGenericWindowsImageParameters(amiMap, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_GenericWindowsImage{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.GenericWindowsImage",
		[]interface{}{amiMap, props},
		&j,
	)

	return &j
}

func NewGenericWindowsImage_Override(g GenericWindowsImage, amiMap *map[string]*string, props *GenericWindowsImageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.GenericWindowsImage",
		[]interface{}{amiMap, props},
		g,
	)
}

func (g *jsiiProxy_GenericWindowsImage) GetImage(scope constructs.Construct) *MachineImageConfig {
	if err := g.validateGetImageParameters(scope); err != nil {
		panic(err)
	}
	var returns *MachineImageConfig

	_jsii_.Invoke(
		g,
		"getImage",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

