package awsec2


// Amazon Linux Kernel.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var vpc vpc
//   var instanceType instanceType
//
//
//   // Amazon Linux 1
//   // Amazon Linux 1
//   ec2.NewInstance(this, jsii.String("Instance1"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: ec2.MachineImage_LatestAmazonLinux(),
//   })
//
//   // Amazon Linux 2
//   // Amazon Linux 2
//   ec2.NewInstance(this, jsii.String("Instance2"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: ec2.MachineImage_*LatestAmazonLinux(&AmazonLinuxImageProps{
//   		Generation: ec2.AmazonLinuxGeneration_AMAZON_LINUX_2,
//   	}),
//   })
//
//   // Amazon Linux 2 with kernel 5.x
//   // Amazon Linux 2 with kernel 5.x
//   ec2.NewInstance(this, jsii.String("Instance3"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: ec2.MachineImage_*LatestAmazonLinux(&AmazonLinuxImageProps{
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
//   	MachineImage: ec2.MachineImage_*LatestAmazonLinux(&AmazonLinuxImageProps{
//   		Generation: ec2.AmazonLinuxGeneration_AMAZON_LINUX_2022,
//   	}),
//   })
//
//   // Graviton 3 Processor
//   // Graviton 3 Processor
//   ec2.NewInstance(this, jsii.String("Instance5"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: ec2.*instanceType_Of(ec2.InstanceClass_C7G, ec2.InstanceSize_LARGE),
//   	MachineImage: ec2.MachineImage_*LatestAmazonLinux(&AmazonLinuxImageProps{
//   		Generation: ec2.AmazonLinuxGeneration_AMAZON_LINUX_2,
//   		CpuType: ec2.AmazonLinuxCpuType_ARM_64,
//   	}),
//   })
//
type AmazonLinuxKernel string

const (
	// Standard edition.
	AmazonLinuxKernel_KERNEL5_X AmazonLinuxKernel = "KERNEL5_X"
)

