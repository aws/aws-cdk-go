package awsec2


// Amazon Linux image properties.
//
// Example:
//   var vpc iVpc
//
//   lb := elb.NewLoadBalancer(this, jsii.String("LB"), &LoadBalancerProps{
//   	Vpc: Vpc,
//   	InternetFacing: jsii.Boolean(true),
//   })
//
//   // instance to add as the target for load balancer.
//   instance := ec2.NewInstance(this, jsii.String("targetInstance"), &InstanceProps{
//   	Vpc: vpc,
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE2, ec2.InstanceSize_MICRO),
//   	MachineImage: ec2.NewAmazonLinuxImage(&AmazonLinuxImageProps{
//   		Generation: ec2.AmazonLinuxGeneration_AMAZON_LINUX_2,
//   	}),
//   })
//   lb.AddTarget(elb.NewInstanceTarget(instance))
//
type AmazonLinuxImageProps struct {
	// Whether the AMI ID is cached to be stable between deployments.
	//
	// By default, the newest image is used on each deployment. This will cause
	// instances to be replaced whenever a new version is released, and may cause
	// downtime if there aren't enough running instances in the AutoScalingGroup
	// to reschedule the tasks on.
	//
	// If set to true, the AMI ID will be cached in `cdk.context.json` and the
	// same value will be used on future runs. Your instances will not be replaced
	// but your AMI version will grow old over time. To refresh the AMI lookup,
	// you will have to evict the value from the cache using the `cdk context`
	// command. See https://docs.aws.amazon.com/cdk/latest/guide/context.html for
	// more information.
	//
	// Can not be set to `true` in environment-agnostic stacks.
	// Default: false.
	//
	CachedInContext *bool `field:"optional" json:"cachedInContext" yaml:"cachedInContext"`
	// CPU Type.
	// Default: X86_64.
	//
	CpuType AmazonLinuxCpuType `field:"optional" json:"cpuType" yaml:"cpuType"`
	// What edition of Amazon Linux to use.
	// Default: Standard.
	//
	Edition AmazonLinuxEdition `field:"optional" json:"edition" yaml:"edition"`
	// What generation of Amazon Linux to use.
	// Default: AmazonLinux.
	//
	Generation AmazonLinuxGeneration `field:"optional" json:"generation" yaml:"generation"`
	// What kernel version of Amazon Linux to use.
	// Default: -.
	//
	Kernel AmazonLinuxKernel `field:"optional" json:"kernel" yaml:"kernel"`
	// What storage backed image to use.
	// Default: GeneralPurpose.
	//
	Storage AmazonLinuxStorage `field:"optional" json:"storage" yaml:"storage"`
	// Initial user data.
	// Default: - Empty UserData for Linux machines.
	//
	UserData UserData `field:"optional" json:"userData" yaml:"userData"`
	// Virtualization type.
	// Default: HVM.
	//
	Virtualization AmazonLinuxVirt `field:"optional" json:"virtualization" yaml:"virtualization"`
}

