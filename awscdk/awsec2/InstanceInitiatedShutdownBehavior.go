package awsec2


// Provides the options for specifying the instance initiated shutdown behavior.
//
// Example:
//   var vpc vpc
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
// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/terminating-instances.html#Using_ChangingInstanceInitiatedShutdownBehavior
//
type InstanceInitiatedShutdownBehavior string

const (
	// The instance will stop when it initiates a shutdown.
	InstanceInitiatedShutdownBehavior_STOP InstanceInitiatedShutdownBehavior = "STOP"
	// The instance will be terminated when it initiates a shutdown.
	InstanceInitiatedShutdownBehavior_TERMINATE InstanceInitiatedShutdownBehavior = "TERMINATE"
)

