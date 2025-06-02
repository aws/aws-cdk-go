package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// A SSM Parameter that contains the AMI ID for Amazon Linux 2023.
//
// Example:
//   var vpc vpc
//
//
//   ec2.NewInstance(this, jsii.String("LatestAl2023"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_C7G, ec2.InstanceSize_LARGE),
//   	MachineImage: ec2.MachineImage_LatestAmazonLinux2023(&AmazonLinux2023ImageSsmParameterProps{
//   		CachedInContext: jsii.Boolean(true),
//   	}),
//   })
//
//   // or
//   // or
//   ec2.NewInstance(this, jsii.String("LatestAl2023"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: ec2.InstanceType_*Of(ec2.InstanceClass_C7G, ec2.InstanceSize_LARGE),
//   	// context cache is turned on by default
//   	MachineImage: ec2.NewAmazonLinux2023ImageSsmParameter(),
//   })
//
type AmazonLinux2023ImageSsmParameter interface {
	AmazonLinuxImageSsmParameterBase
	// Return the image to use in the given context.
	GetImage(scope constructs.Construct) *MachineImageConfig
}

// The jsii proxy struct for AmazonLinux2023ImageSsmParameter
type jsiiProxy_AmazonLinux2023ImageSsmParameter struct {
	jsiiProxy_AmazonLinuxImageSsmParameterBase
}

func NewAmazonLinux2023ImageSsmParameter(props *AmazonLinux2023ImageSsmParameterProps) AmazonLinux2023ImageSsmParameter {
	_init_.Initialize()

	if err := validateNewAmazonLinux2023ImageSsmParameterParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AmazonLinux2023ImageSsmParameter{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.AmazonLinux2023ImageSsmParameter",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewAmazonLinux2023ImageSsmParameter_Override(a AmazonLinux2023ImageSsmParameter, props *AmazonLinux2023ImageSsmParameterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.AmazonLinux2023ImageSsmParameter",
		[]interface{}{props},
		a,
	)
}

// Generates a SSM Parameter name for a specific amazon linux 2023 AMI.
//
// Example values:
//
//     "/aws/service/ami-amazon-linux-latest/al2023-ami-kernel-6.1-arm64",
//     "/aws/service/ami-amazon-linux-latest/al2023-ami-kernel-6.1-x86_64",
//     "/aws/service/ami-amazon-linux-latest/al2023-ami-minimal-kernel-6.1-arm64",
//     "/aws/service/ami-amazon-linux-latest/al2023-ami-minimal-kernel-6.1-x86_64",
//     "/aws/service/ami-amazon-linux-latest/al2023-ami-minimal-kernel-default-arm64",
//     "/aws/service/ami-amazon-linux-latest/al2023-ami-kernel-default-x86_64",
//     "/aws/service/ami-amazon-linux-latest/al2023-ami-minimal-kernel-default-x86_64",
// "/aws/service/ami-amazon-linux-latest/al2023-ami-kernel-default-arm64",.
func AmazonLinux2023ImageSsmParameter_SsmParameterName(props *AmazonLinux2023ImageSsmParameterProps) *string {
	_init_.Initialize()

	if err := validateAmazonLinux2023ImageSsmParameter_SsmParameterNameParameters(props); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.AmazonLinux2023ImageSsmParameter",
		"ssmParameterName",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AmazonLinux2023ImageSsmParameter) GetImage(scope constructs.Construct) *MachineImageConfig {
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

