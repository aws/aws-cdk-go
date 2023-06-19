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
//   ec2.NewInstance(this, jsii.String("Instance1"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: ec2.NewAmazonLinuxImage(),
//   })
//
//   // AWS Linux 2
//   // AWS Linux 2
//   ec2.NewInstance(this, jsii.String("Instance2"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: ec2.NewAmazonLinuxImage(&AmazonLinuxImageProps{
//   		Generation: ec2.AmazonLinuxGeneration_AMAZON_LINUX_2,
//   	}),
//   })
//
//   // AWS Linux 2 with kernel 5.x
//   // AWS Linux 2 with kernel 5.x
//   ec2.NewInstance(this, jsii.String("Instance3"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: ec2.NewAmazonLinuxImage(&AmazonLinuxImageProps{
//   		Generation: ec2.AmazonLinuxGeneration_AMAZON_LINUX_2,
//   		Kernel: ec2.AmazonLinuxKernel_KERNEL5_X,
//   	}),
//   })
//
//   // AWS Linux 2022
//   // AWS Linux 2022
//   ec2.NewInstance(this, jsii.String("Instance4"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: ec2.NewAmazonLinuxImage(&AmazonLinuxImageProps{
//   		Generation: ec2.AmazonLinuxGeneration_AMAZON_LINUX_2022,
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

