package awsec2


// Properties for GenericSsmParameterImage.
//
// Example:
//   // Pick the right Amazon Linux edition. All arguments shown are optional
//   // and will default to these values when omitted.
//   amznLinux := ec2.MachineImage_LatestAmazonLinux(&AmazonLinuxImageProps{
//   	Generation: ec2.AmazonLinuxGeneration_AMAZON_LINUX,
//   	Edition: ec2.AmazonLinuxEdition_STANDARD,
//   	Virtualization: ec2.AmazonLinuxVirt_HVM,
//   	Storage: ec2.AmazonLinuxStorage_GENERAL_PURPOSE,
//   	CpuType: ec2.AmazonLinuxCpuType_X86_64,
//   })
//
//   // Pick a Windows edition to use
//   windows := ec2.MachineImage_LatestWindows(ec2.WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_BASE)
//
//   // Read AMI id from SSM parameter store
//   ssm := ec2.MachineImage_FromSsmParameter(jsii.String("/my/ami"), &SsmParameterImageOptions{
//   	Os: ec2.OperatingSystemType_LINUX,
//   })
//
//   // Look up the most recent image matching a set of AMI filters.
//   // In this case, look up the NAT instance AMI, by using a wildcard
//   // in the 'name' field:
//   natAmi := ec2.MachineImage_Lookup(&LookupMachineImageProps{
//   	Name: jsii.String("amzn-ami-vpc-nat-*"),
//   	Owners: []*string{
//   		jsii.String("amazon"),
//   	},
//   })
//
//   // For other custom (Linux) images, instantiate a `GenericLinuxImage` with
//   // a map giving the AMI to in for each region:
//   linux := ec2.MachineImage_GenericLinux(map[string]*string{
//   	"us-east-1": jsii.String("ami-97785bed"),
//   	"eu-west-1": jsii.String("ami-12345678"),
//   })
//
//   // For other custom (Windows) images, instantiate a `GenericWindowsImage` with
//   // a map giving the AMI to in for each region:
//   genericWindows := ec2.MachineImage_GenericWindows(map[string]*string{
//   	"us-east-1": jsii.String("ami-97785bed"),
//   	"eu-west-1": jsii.String("ami-12345678"),
//   })
//
type SsmParameterImageOptions struct {
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
	// Operating system.
	// Default: OperatingSystemType.LINUX
	//
	Os OperatingSystemType `field:"optional" json:"os" yaml:"os"`
	// The version of the SSM parameter.
	// Default: no version specified.
	//
	ParameterVersion *string `field:"optional" json:"parameterVersion" yaml:"parameterVersion"`
	// Custom UserData.
	// Default: - UserData appropriate for the OS.
	//
	UserData UserData `field:"optional" json:"userData" yaml:"userData"`
}

