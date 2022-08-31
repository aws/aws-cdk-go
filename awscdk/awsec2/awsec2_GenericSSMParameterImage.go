package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// Select the image based on a given SSM parameter.
//
// This Machine Image automatically updates to the latest version on every
// deployment. Be aware this will cause your instances to be replaced when a
// new version of the image becomes available. Do not store stateful information
// on the instance if you are using this image.
//
// The AMI ID is selected using the values published to the SSM parameter store.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var userData userData
//
//   genericSSMParameterImage := awscdk.Aws_ec2.NewGenericSSMParameterImage(jsii.String("parameterName"), awscdk.Aws_ec2.operatingSystemType_LINUX, userData)
//
// Experimental.
type GenericSSMParameterImage interface {
	IMachineImage
	// Name of the SSM parameter we're looking up.
	// Experimental.
	ParameterName() *string
	// Return the image to use in the given context.
	// Experimental.
	GetImage(scope awscdk.Construct) *MachineImageConfig
}

// The jsii proxy struct for GenericSSMParameterImage
type jsiiProxy_GenericSSMParameterImage struct {
	jsiiProxy_IMachineImage
}

func (j *jsiiProxy_GenericSSMParameterImage) ParameterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterName",
		&returns,
	)
	return returns
}


// Experimental.
func NewGenericSSMParameterImage(parameterName *string, os OperatingSystemType, userData UserData) GenericSSMParameterImage {
	_init_.Initialize()

	j := jsiiProxy_GenericSSMParameterImage{}

	_jsii_.Create(
		"monocdk.aws_ec2.GenericSSMParameterImage",
		[]interface{}{parameterName, os, userData},
		&j,
	)

	return &j
}

// Experimental.
func NewGenericSSMParameterImage_Override(g GenericSSMParameterImage, parameterName *string, os OperatingSystemType, userData UserData) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ec2.GenericSSMParameterImage",
		[]interface{}{parameterName, os, userData},
		g,
	)
}

func (g *jsiiProxy_GenericSSMParameterImage) GetImage(scope awscdk.Construct) *MachineImageConfig {
	var returns *MachineImageConfig

	_jsii_.Invoke(
		g,
		"getImage",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

