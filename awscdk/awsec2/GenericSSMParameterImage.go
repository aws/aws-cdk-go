package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Select the image based on a given SSM parameter at deployment time of the CloudFormation Stack.
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
//   genericSSMParameterImage := awscdk.Aws_ec2.NewGenericSSMParameterImage(jsii.String("parameterName"), awscdk.Aws_ec2.OperatingSystemType_LINUX, userData)
//
type GenericSSMParameterImage interface {
	IMachineImage
	// Name of the SSM parameter we're looking up.
	ParameterName() *string
	// Return the image to use in the given context.
	GetImage(scope constructs.Construct) *MachineImageConfig
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


func NewGenericSSMParameterImage(parameterName *string, os OperatingSystemType, userData UserData) GenericSSMParameterImage {
	_init_.Initialize()

	if err := validateNewGenericSSMParameterImageParameters(parameterName, os); err != nil {
		panic(err)
	}
	j := jsiiProxy_GenericSSMParameterImage{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.GenericSSMParameterImage",
		[]interface{}{parameterName, os, userData},
		&j,
	)

	return &j
}

func NewGenericSSMParameterImage_Override(g GenericSSMParameterImage, parameterName *string, os OperatingSystemType, userData UserData) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.GenericSSMParameterImage",
		[]interface{}{parameterName, os, userData},
		g,
	)
}

func (g *jsiiProxy_GenericSSMParameterImage) GetImage(scope constructs.Construct) *MachineImageConfig {
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

