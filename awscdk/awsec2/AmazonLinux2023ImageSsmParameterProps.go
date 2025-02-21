package awsec2


// Properties specific to al2023 images.
//
// Example:
//   var vpc vpc
//
//
//   ec2.NewInstance(this, jsii.String("LatestAl2023"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_C7G, ec2.InstanceSize_LARGE),
//   	MachineImage: ec2.MachineImage_LatestAmazonLinux2023(&AmazonLinux2023ImageSsmParameterProps{
//   		CachedInContext: jsii.Boolean(true),
//   	}),
//   })
//
//   // or
//   // or
//   ec2.NewInstance(this, jsii.String("LatestAl2023"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: ec2.InstanceType_*Of(ec2.InstanceClass_C7G, ec2.InstanceSize_LARGE),
//   	// context cache is turned on by default
//   	MachineImage: ec2.NewAmazonLinux2023ImageSsmParameter(),
//   })
//
type AmazonLinux2023ImageSsmParameterProps struct {
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
	// Initial user data.
	// Default: - Empty UserData for Linux machines.
	//
	UserData UserData `field:"optional" json:"userData" yaml:"userData"`
	// CPU Type.
	// Default: AmazonLinuxCpuType.X86_64
	//
	CpuType AmazonLinuxCpuType `field:"optional" json:"cpuType" yaml:"cpuType"`
	// What edition of Amazon Linux to use.
	// Default: AmazonLinuxEdition.Standard
	//
	Edition AmazonLinuxEdition `field:"optional" json:"edition" yaml:"edition"`
	// What kernel version of Amazon Linux to use.
	// Default: AmazonLinux2023Kernel.DEFAULT
	//
	Kernel AmazonLinux2023Kernel `field:"optional" json:"kernel" yaml:"kernel"`
}

