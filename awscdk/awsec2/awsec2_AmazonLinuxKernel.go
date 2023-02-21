package awsec2


// Amazon Linux Kernel.
//
// Example:
//   var vpc vpc
//   var instanceType instanceType
//
//
//   // AWS Linux
//   // AWS Linux
//   ec2.NewInstance(this, jsii.String("Instance1"), &instanceProps{
//   	vpc: vpc,
//   	instanceType: instanceType,
//   	machineImage: ec2.NewAmazonLinuxImage(),
//   })
//
//   // AWS Linux 2
//   // AWS Linux 2
//   ec2.NewInstance(this, jsii.String("Instance2"), &instanceProps{
//   	vpc: vpc,
//   	instanceType: instanceType,
//   	machineImage: ec2.NewAmazonLinuxImage(&amazonLinuxImageProps{
//   		generation: ec2.amazonLinuxGeneration_AMAZON_LINUX_2,
//   	}),
//   })
//
//   // AWS Linux 2 with kernel 5.x
//   // AWS Linux 2 with kernel 5.x
//   ec2.NewInstance(this, jsii.String("Instance3"), &instanceProps{
//   	vpc: vpc,
//   	instanceType: instanceType,
//   	machineImage: ec2.NewAmazonLinuxImage(&amazonLinuxImageProps{
//   		generation: ec2.*amazonLinuxGeneration_AMAZON_LINUX_2,
//   		kernel: ec2.amazonLinuxKernel_KERNEL5_X,
//   	}),
//   })
//
//   // AWS Linux 2022
//   // AWS Linux 2022
//   ec2.NewInstance(this, jsii.String("Instance4"), &instanceProps{
//   	vpc: vpc,
//   	instanceType: instanceType,
//   	machineImage: ec2.NewAmazonLinuxImage(&amazonLinuxImageProps{
//   		generation: ec2.*amazonLinuxGeneration_AMAZON_LINUX_2022,
//   	}),
//   })
//
// Experimental.
type AmazonLinuxKernel string

const (
	// Standard edition.
	// Experimental.
	AmazonLinuxKernel_KERNEL5_X AmazonLinuxKernel = "KERNEL5_X"
)

