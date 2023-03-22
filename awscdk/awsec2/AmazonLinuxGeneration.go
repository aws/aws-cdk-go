package awsec2


// What generation of Amazon Linux to use.
//
// Example:
//   var vpc vpc
//
//
//   mySecurityGroup := ec2.NewSecurityGroup(this, jsii.String("SecurityGroup"), &SecurityGroupProps{
//   	Vpc: Vpc,
//   })
//   autoscaling.NewAutoScalingGroup(this, jsii.String("ASG"), &AutoScalingGroupProps{
//   	Vpc: Vpc,
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE2, ec2.InstanceSize_MICRO),
//   	MachineImage: ec2.MachineImage_LatestAmazonLinux(&AmazonLinuxImageProps{
//   		Generation: ec2.AmazonLinuxGeneration_AMAZON_LINUX_2,
//   	}),
//   	SecurityGroup: mySecurityGroup,
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
)

