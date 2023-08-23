package awsec2


// What generation of Amazon Linux to use.
//
// Example:
//   var vpc vpc
//
//
//   autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
//   	Vpc: Vpc,
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_T3, ec2.InstanceSize_MICRO),
//
//   	// Amazon Linux 2 comes with SSM Agent by default
//   	MachineImage: ec2.MachineImage_LatestAmazonLinux(&AmazonLinuxImageProps{
//   		Generation: ec2.AmazonLinuxGeneration_AMAZON_LINUX_2,
//   	}),
//
//   	// Turn on SSM
//   	SsmSessionPermissions: jsii.Boolean(true),
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

