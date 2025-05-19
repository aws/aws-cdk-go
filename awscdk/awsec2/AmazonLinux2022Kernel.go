package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Amazon Linux 2022 kernel versions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   amazonLinux2022Kernel := awscdk.Aws_ec2.AmazonLinux2022Kernel_CDK_LATEST()
//
type AmazonLinux2022Kernel interface {
	// Generate a string representation of the kernel.
	ToString() *string
}

// The jsii proxy struct for AmazonLinux2022Kernel
type jsiiProxy_AmazonLinux2022Kernel struct {
	_ byte // padding
}

func NewAmazonLinux2022Kernel(version *string) AmazonLinux2022Kernel {
	_init_.Initialize()

	if err := validateNewAmazonLinux2022KernelParameters(version); err != nil {
		panic(err)
	}
	j := jsiiProxy_AmazonLinux2022Kernel{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.AmazonLinux2022Kernel",
		[]interface{}{version},
		&j,
	)

	return &j
}

func NewAmazonLinux2022Kernel_Override(a AmazonLinux2022Kernel, version *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.AmazonLinux2022Kernel",
		[]interface{}{version},
		a,
	)
}

func AmazonLinux2022Kernel_CDK_LATEST() AmazonLinux2022Kernel {
	_init_.Initialize()
	var returns AmazonLinux2022Kernel
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.AmazonLinux2022Kernel",
		"CDK_LATEST",
		&returns,
	)
	return returns
}

func AmazonLinux2022Kernel_DEFAULT() AmazonLinux2022Kernel {
	_init_.Initialize()
	var returns AmazonLinux2022Kernel
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.AmazonLinux2022Kernel",
		"DEFAULT",
		&returns,
	)
	return returns
}

func AmazonLinux2022Kernel_KERNEL_5_15() AmazonLinux2022Kernel {
	_init_.Initialize()
	var returns AmazonLinux2022Kernel
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.AmazonLinux2022Kernel",
		"KERNEL_5_15",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AmazonLinux2022Kernel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

