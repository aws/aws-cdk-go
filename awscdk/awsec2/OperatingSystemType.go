package awsec2


// The OS type of a particular image.
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
type OperatingSystemType string

const (
	OperatingSystemType_LINUX OperatingSystemType = "LINUX"
	OperatingSystemType_WINDOWS OperatingSystemType = "WINDOWS"
	// Used when the type of the operating system is not known (for example, for imported Auto-Scaling Groups).
	OperatingSystemType_UNKNOWN OperatingSystemType = "UNKNOWN"
)

