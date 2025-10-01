package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Amazon Linux 2 kernel versions.
//
// Example:
//   var vpc vpc
//   var instanceType instanceType
//
//
//   // Amazon Linux 2
//   // Amazon Linux 2
//   ec2.NewInstance(this, jsii.String("Instance2"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: ec2.MachineImage_LatestAmazonLinux2(),
//   })
//
//   // Amazon Linux 2 with kernel 5.x
//   // Amazon Linux 2 with kernel 5.x
//   ec2.NewInstance(this, jsii.String("Instance3"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: ec2.MachineImage_*LatestAmazonLinux2(&AmazonLinux2ImageSsmParameterProps{
//   		Kernel: ec2.AmazonLinux2Kernel_KERNEL_5_10(),
//   	}),
//   })
//
//   // Amazon Linux 2023
//   // Amazon Linux 2023
//   ec2.NewInstance(this, jsii.String("Instance4"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: ec2.MachineImage_LatestAmazonLinux2023(),
//   })
//
//   // Graviton 3 Processor
//   // Graviton 3 Processor
//   ec2.NewInstance(this, jsii.String("Instance5"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: ec2.*instanceType_Of(ec2.InstanceClass_C7G, ec2.InstanceSize_LARGE),
//   	MachineImage: ec2.MachineImage_*LatestAmazonLinux2023(&AmazonLinux2023ImageSsmParameterProps{
//   		CpuType: ec2.AmazonLinuxCpuType_ARM_64,
//   	}),
//   })
//
type AmazonLinux2Kernel interface {
	// Generate a string representation of the kernel.
	ToString() *string
}

// The jsii proxy struct for AmazonLinux2Kernel
type jsiiProxy_AmazonLinux2Kernel struct {
	_ byte // padding
}

func NewAmazonLinux2Kernel(version *string) AmazonLinux2Kernel {
	_init_.Initialize()

	if err := validateNewAmazonLinux2KernelParameters(version); err != nil {
		panic(err)
	}
	j := jsiiProxy_AmazonLinux2Kernel{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.AmazonLinux2Kernel",
		[]interface{}{version},
		&j,
	)

	return &j
}

func NewAmazonLinux2Kernel_Override(a AmazonLinux2Kernel, version *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.AmazonLinux2Kernel",
		[]interface{}{version},
		a,
	)
}

func AmazonLinux2Kernel_CDK_LATEST() AmazonLinux2Kernel {
	_init_.Initialize()
	var returns AmazonLinux2Kernel
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.AmazonLinux2Kernel",
		"CDK_LATEST",
		&returns,
	)
	return returns
}

func AmazonLinux2Kernel_DEFAULT() AmazonLinux2Kernel {
	_init_.Initialize()
	var returns AmazonLinux2Kernel
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.AmazonLinux2Kernel",
		"DEFAULT",
		&returns,
	)
	return returns
}

func AmazonLinux2Kernel_KERNEL_5_10() AmazonLinux2Kernel {
	_init_.Initialize()
	var returns AmazonLinux2Kernel
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.AmazonLinux2Kernel",
		"KERNEL_5_10",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AmazonLinux2Kernel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

