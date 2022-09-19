package awsec2


// Properties for looking up an image.
//
// Example:
//   // Pick the right Amazon Linux edition. All arguments shown are optional
//   // and will default to these values when omitted.
//   amznLinux := ec2.machineImage.latestAmazonLinux(&amazonLinuxImageProps{
//   	generation: ec2.amazonLinuxGeneration_AMAZON_LINUX,
//   	edition: ec2.amazonLinuxEdition_STANDARD,
//   	virtualization: ec2.amazonLinuxVirt_HVM,
//   	storage: ec2.amazonLinuxStorage_GENERAL_PURPOSE,
//   	cpuType: ec2.amazonLinuxCpuType_X86_64,
//   })
//
//   // Pick a Windows edition to use
//   windows := ec2.machineImage.latestWindows(ec2.windowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_BASE)
//
//   // Read AMI id from SSM parameter store
//   ssm := ec2.machineImage.fromSsmParameter(jsii.String("/my/ami"), &ssmParameterImageOptions{
//   	os: ec2.operatingSystemType_LINUX,
//   })
//
//   // Look up the most recent image matching a set of AMI filters.
//   // In this case, look up the NAT instance AMI, by using a wildcard
//   // in the 'name' field:
//   natAmi := ec2.machineImage.lookup(&lookupMachineImageProps{
//   	name: jsii.String("amzn-ami-vpc-nat-*"),
//   	owners: []*string{
//   		jsii.String("amazon"),
//   	},
//   })
//
//   // For other custom (Linux) images, instantiate a `GenericLinuxImage` with
//   // a map giving the AMI to in for each region:
//   linux := ec2.machineImage.genericLinux(map[string]*string{
//   	"us-east-1": jsii.String("ami-97785bed"),
//   	"eu-west-1": jsii.String("ami-12345678"),
//   })
//
//   // For other custom (Windows) images, instantiate a `GenericWindowsImage` with
//   // a map giving the AMI to in for each region:
//   genericWindows := ec2.machineImage.genericWindows(map[string]*string{
//   	"us-east-1": jsii.String("ami-97785bed"),
//   	"eu-west-1": jsii.String("ami-12345678"),
//   })
//
type LookupMachineImageProps struct {
	// Name of the image (may contain wildcards).
	Name *string `field:"required" json:"name" yaml:"name"`
	// Additional filters on the AMI.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeImages.html
	//
	Filters *map[string]*[]*string `field:"optional" json:"filters" yaml:"filters"`
	// Owner account IDs or aliases.
	Owners *[]*string `field:"optional" json:"owners" yaml:"owners"`
	// Custom userdata for this image.
	UserData UserData `field:"optional" json:"userData" yaml:"userData"`
	// Look for Windows images.
	Windows *bool `field:"optional" json:"windows" yaml:"windows"`
}

