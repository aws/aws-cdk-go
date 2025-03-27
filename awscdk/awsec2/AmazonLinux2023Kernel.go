package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Amazon Linux 2023 kernel versions.
//
// Example:
//   var vpc vpc
//
//
//   ec2.NewInstance(this, jsii.String("LatestAl2023"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_C7G, ec2.InstanceSize_LARGE),
//   	// context cache is turned on by default
//   	MachineImage: ec2.NewAmazonLinux2023ImageSsmParameter(&AmazonLinux2023ImageSsmParameterProps{
//   		Kernel: ec2.AmazonLinux2023Kernel_KERNEL_6_1(),
//   	}),
//   })
//
type AmazonLinux2023Kernel interface {
	// Generate a string representation of the kernel.
	ToString() *string
}

// The jsii proxy struct for AmazonLinux2023Kernel
type jsiiProxy_AmazonLinux2023Kernel struct {
	_ byte // padding
}

func NewAmazonLinux2023Kernel(version *string) AmazonLinux2023Kernel {
	_init_.Initialize()

	if err := validateNewAmazonLinux2023KernelParameters(version); err != nil {
		panic(err)
	}
	j := jsiiProxy_AmazonLinux2023Kernel{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.AmazonLinux2023Kernel",
		[]interface{}{version},
		&j,
	)

	return &j
}

func NewAmazonLinux2023Kernel_Override(a AmazonLinux2023Kernel, version *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.AmazonLinux2023Kernel",
		[]interface{}{version},
		a,
	)
}

func AmazonLinux2023Kernel_CDK_LATEST() AmazonLinux2023Kernel {
	_init_.Initialize()
	var returns AmazonLinux2023Kernel
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.AmazonLinux2023Kernel",
		"CDK_LATEST",
		&returns,
	)
	return returns
}

func AmazonLinux2023Kernel_DEFAULT() AmazonLinux2023Kernel {
	_init_.Initialize()
	var returns AmazonLinux2023Kernel
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.AmazonLinux2023Kernel",
		"DEFAULT",
		&returns,
	)
	return returns
}

func AmazonLinux2023Kernel_KERNEL_6_1() AmazonLinux2023Kernel {
	_init_.Initialize()
	var returns AmazonLinux2023Kernel
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.AmazonLinux2023Kernel",
		"KERNEL_6_1",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AmazonLinux2023Kernel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

