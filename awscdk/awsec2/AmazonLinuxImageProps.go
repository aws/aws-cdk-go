package awsec2


// Amazon Linux image properties.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var vpc vpc
//   var instanceType instanceType
//
//
//   ec2.NewInstance(this, jsii.String("Instance"), &InstanceProps{
//   	Vpc: Vpc,
//   	InstanceType: InstanceType,
//   	MachineImage: ec2.MachineImage_LatestAmazonLinux(&AmazonLinuxImageProps{
//   		// Amazon Linux 2 uses SystemD
//   		Generation: ec2.AmazonLinuxGeneration,
//   		AMAZON_LINUX_2: AMAZON_LINUX_2,
//   	}),
//
//   	Init: ec2.CloudFormationInit_FromElements([]interface{}{
//   		ec2.InitService_SystemdConfigFile(jsii.String("simpleserver"), &SystemdConfigFileOptions{
//   			Command: jsii.String("/usr/bin/python3 -m http.server 8080"),
//   			Cwd: jsii.String("/var/www/html"),
//   		}),
//   		ec2.InitService_Enable(jsii.String("simpleserver"), &InitServiceOptions{
//   			ServiceManager: ec2.ServiceManager_SYSTEMD,
//   		}),
//   		ec2.InitFile_FromString(jsii.String("/var/www/html/index.html"), jsii.String("Hello! It's working!")),
//   	}),
//   })
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
	CachedInContext *bool `field:"optional" json:"cachedInContext" yaml:"cachedInContext"`
	// CPU Type.
	CpuType AmazonLinuxCpuType `field:"optional" json:"cpuType" yaml:"cpuType"`
	// What edition of Amazon Linux to use.
	Edition AmazonLinuxEdition `field:"optional" json:"edition" yaml:"edition"`
	// What generation of Amazon Linux to use.
	Generation AmazonLinuxGeneration `field:"optional" json:"generation" yaml:"generation"`
	// What kernel version of Amazon Linux to use.
	Kernel AmazonLinuxKernel `field:"optional" json:"kernel" yaml:"kernel"`
	// What storage backed image to use.
	Storage AmazonLinuxStorage `field:"optional" json:"storage" yaml:"storage"`
	// Initial user data.
	UserData UserData `field:"optional" json:"userData" yaml:"userData"`
	// Virtualization type.
	Virtualization AmazonLinuxVirt `field:"optional" json:"virtualization" yaml:"virtualization"`
}

