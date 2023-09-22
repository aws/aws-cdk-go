package awsec2


// What generation of Amazon Linux to use.
//
// Example:
//   // Pick a Windows edition to use
//   windows := ec2.NewWindowsImage(ec2.WindowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_BASE)
//
//   // Pick the right Amazon Linux edition. All arguments shown are optional
//   // and will default to these values when omitted.
//   amznLinux := ec2.NewAmazonLinuxImage(&AmazonLinuxImageProps{
//   	Generation: ec2.AmazonLinuxGeneration_AMAZON_LINUX,
//   	Edition: ec2.AmazonLinuxEdition_STANDARD,
//   	Virtualization: ec2.AmazonLinuxVirt_HVM,
//   	Storage: ec2.AmazonLinuxStorage_GENERAL_PURPOSE,
//   })
//
//   // For other custom (Linux) images, instantiate a `GenericLinuxImage` with
//   // a map giving the AMI to in for each region:
//
//   linux := ec2.NewGenericLinuxImage(map[string]*string{
//   	"us-east-1": jsii.String("ami-97785bed"),
//   	"eu-west-1": jsii.String("ami-12345678"),
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

