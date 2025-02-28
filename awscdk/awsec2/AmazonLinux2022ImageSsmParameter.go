package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// A SSM Parameter that contains the AMI ID for Amazon Linux 2023.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var amazonLinux2022Kernel amazonLinux2022Kernel
//   var userData userData
//
//   amazonLinux2022ImageSsmParameter := awscdk.Aws_ec2.NewAmazonLinux2022ImageSsmParameter(&AmazonLinux2022ImageSsmParameterProps{
//   	CachedInContext: jsii.Boolean(false),
//   	CpuType: awscdk.*Aws_ec2.AmazonLinuxCpuType_ARM_64,
//   	Edition: awscdk.*Aws_ec2.AmazonLinuxEdition_STANDARD,
//   	Kernel: amazonLinux2022Kernel,
//   	UserData: userData,
//   })
//
type AmazonLinux2022ImageSsmParameter interface {
	AmazonLinuxImageSsmParameterBase
	// Return the image to use in the given context.
	GetImage(scope constructs.Construct) *MachineImageConfig
}

// The jsii proxy struct for AmazonLinux2022ImageSsmParameter
type jsiiProxy_AmazonLinux2022ImageSsmParameter struct {
	jsiiProxy_AmazonLinuxImageSsmParameterBase
}

func NewAmazonLinux2022ImageSsmParameter(props *AmazonLinux2022ImageSsmParameterProps) AmazonLinux2022ImageSsmParameter {
	_init_.Initialize()

	if err := validateNewAmazonLinux2022ImageSsmParameterParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AmazonLinux2022ImageSsmParameter{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.AmazonLinux2022ImageSsmParameter",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewAmazonLinux2022ImageSsmParameter_Override(a AmazonLinux2022ImageSsmParameter, props *AmazonLinux2022ImageSsmParameterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.AmazonLinux2022ImageSsmParameter",
		[]interface{}{props},
		a,
	)
}

// Generates a SSM Parameter name for a specific amazon linux 2022 AMI.
//
// Example values:
//
//     "/aws/service/ami-amazon-linux-latest/al2022-ami-kernel-5.15-x86_64",
//     "/aws/service/ami-amazon-linux-latest/al2022-ami-kernel-default-x86_64",
//     "/aws/service/ami-amazon-linux-latest/al2022-ami-minimal-kernel-5.15-arm64",
//     "/aws/service/ami-amazon-linux-latest/al2022-ami-minimal-kernel-5.15-x86_64",
//     "/aws/service/ami-amazon-linux-latest/al2022-ami-kernel-5.15-arm64",
//     "/aws/service/ami-amazon-linux-latest/al2022-ami-minimal-kernel-default-arm64",
//     "/aws/service/ami-amazon-linux-latest/al2022-ami-minimal-kernel-default-x86_64",
// "/aws/service/ami-amazon-linux-latest/al2022-ami-kernel-default-arm64",.
func AmazonLinux2022ImageSsmParameter_SsmParameterName(props *AmazonLinux2022ImageSsmParameterProps) *string {
	_init_.Initialize()

	if err := validateAmazonLinux2022ImageSsmParameter_SsmParameterNameParameters(props); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.AmazonLinux2022ImageSsmParameter",
		"ssmParameterName",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AmazonLinux2022ImageSsmParameter) GetImage(scope constructs.Construct) *MachineImageConfig {
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

