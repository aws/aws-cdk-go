package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// A SSM Parameter that contains the AMI ID for Amazon Linux 2.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var amazonLinux2Kernel amazonLinux2Kernel
//   var userData userData
//
//   amazonLinux2ImageSsmParameter := awscdk.Aws_ec2.NewAmazonLinux2ImageSsmParameter(&AmazonLinux2ImageSsmParameterProps{
//   	CachedInContext: jsii.Boolean(false),
//   	CpuType: awscdk.*Aws_ec2.AmazonLinuxCpuType_ARM_64,
//   	Edition: awscdk.*Aws_ec2.AmazonLinuxEdition_STANDARD,
//   	Kernel: amazonLinux2Kernel,
//   	Storage: awscdk.*Aws_ec2.AmazonLinuxStorage_EBS,
//   	UserData: userData,
//   	Virtualization: awscdk.*Aws_ec2.AmazonLinuxVirt_HVM,
//   })
//
type AmazonLinux2ImageSsmParameter interface {
	AmazonLinuxImageSsmParameterBase
	// Return the image to use in the given context.
	GetImage(scope constructs.Construct) *MachineImageConfig
}

// The jsii proxy struct for AmazonLinux2ImageSsmParameter
type jsiiProxy_AmazonLinux2ImageSsmParameter struct {
	jsiiProxy_AmazonLinuxImageSsmParameterBase
}

func NewAmazonLinux2ImageSsmParameter(props *AmazonLinux2ImageSsmParameterProps) AmazonLinux2ImageSsmParameter {
	_init_.Initialize()

	if err := validateNewAmazonLinux2ImageSsmParameterParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AmazonLinux2ImageSsmParameter{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.AmazonLinux2ImageSsmParameter",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewAmazonLinux2ImageSsmParameter_Override(a AmazonLinux2ImageSsmParameter, props *AmazonLinux2ImageSsmParameterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.AmazonLinux2ImageSsmParameter",
		[]interface{}{props},
		a,
	)
}

// Generates a SSM Parameter name for a specific amazon linux 2 AMI.
//
// Example values:
//
//     "/aws/service/ami-amazon-linux-latest/amzn2-ami-hvm-x86_64-ebs",
//     "/aws/service/ami-amazon-linux-latest/amzn2-ami-hvm-x86_64-gp2",
//     "/aws/service/ami-amazon-linux-latest/amzn2-ami-kernel-5.10-hvm-x86_64-ebs",
//     "/aws/service/ami-amazon-linux-latest/amzn2-ami-hvm-arm64-gp2",
//     "/aws/service/ami-amazon-linux-latest/amzn2-ami-minimal-hvm-arm64-ebs",
//     "/aws/service/ami-amazon-linux-latest/amzn2-ami-kernel-5.10-hvm-arm64-gp2",
//     "/aws/service/ami-amazon-linux-latest/amzn2-ami-kernel-5.10-hvm-x86_64-gp2",
// "/aws/service/ami-amazon-linux-latest/amzn2-ami-minimal-hvm-x86_64-ebs".
func AmazonLinux2ImageSsmParameter_SsmParameterName(props *AmazonLinux2ImageSsmParameterProps) *string {
	_init_.Initialize()

	if err := validateAmazonLinux2ImageSsmParameter_SsmParameterNameParameters(props); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.AmazonLinux2ImageSsmParameter",
		"ssmParameterName",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AmazonLinux2ImageSsmParameter) GetImage(scope constructs.Construct) *MachineImageConfig {
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

