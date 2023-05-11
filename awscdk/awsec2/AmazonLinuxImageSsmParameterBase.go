package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

type AmazonLinuxImageSsmParameterBase interface {
	IMachineImage
	// Return the image to use in the given context.
	GetImage(scope constructs.Construct) *MachineImageConfig
}

// The jsii proxy struct for AmazonLinuxImageSsmParameterBase
type jsiiProxy_AmazonLinuxImageSsmParameterBase struct {
	jsiiProxy_IMachineImage
}

func NewAmazonLinuxImageSsmParameterBase_Override(a AmazonLinuxImageSsmParameterBase, props *AmazonLinuxImageSsmParameterBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.AmazonLinuxImageSsmParameterBase",
		[]interface{}{props},
		a,
	)
}

func (a *jsiiProxy_AmazonLinuxImageSsmParameterBase) GetImage(scope constructs.Construct) *MachineImageConfig {
	if err := a.validateGetImageParameters(scope); err != nil {
		panic(err)
	}
	var returns *MachineImageConfig

	_jsii_.Invoke(
		a,
		"getImage",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

