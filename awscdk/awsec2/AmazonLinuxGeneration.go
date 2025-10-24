package awsec2


// What generation of Amazon Linux to use.
//
// Example:
//   var vpc Vpc
//
//
//   ec2.NewInstance(this, jsii.String("Instance"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_T3, ec2.InstanceSize_NANO),
//   	MachineImage: ec2.NewAmazonLinuxImage(&AmazonLinuxImageProps{
//   		Generation: ec2.AmazonLinuxGeneration_AMAZON_LINUX_2,
//   	}),
//   	InstanceInitiatedShutdownBehavior: ec2.InstanceInitiatedShutdownBehavior_TERMINATE,
//   })
//
type AmazonLinuxGeneration string

const (
	// Amazon Linux.
	AmazonLinuxGeneration_AMAZON_LINUX AmazonLinuxGeneration = "AMAZON_LINUX"
	// Amazon Linux 2.
	AmazonLinuxGeneration_AMAZON_LINUX_2 AmazonLinuxGeneration = "AMAZON_LINUX_2"
	// Amazon Linux 2022.
	AmazonLinuxGeneration_AMAZON_LINUX_2022 AmazonLinuxGeneration = "AMAZON_LINUX_2022"
	// Amazon Linux 2023.
	AmazonLinuxGeneration_AMAZON_LINUX_2023 AmazonLinuxGeneration = "AMAZON_LINUX_2023"
)

