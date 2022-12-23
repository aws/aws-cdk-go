package awsec2


// What generation of Amazon Linux to use.
//
// Example:
//   // Pick a Windows edition to use
//   windows := ec2.NewWindowsImage(ec2.windowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_BASE)
//
//   // Pick the right Amazon Linux edition. All arguments shown are optional
//   // and will default to these values when omitted.
//   amznLinux := ec2.NewAmazonLinuxImage(&amazonLinuxImageProps{
//   	generation: ec2.amazonLinuxGeneration_AMAZON_LINUX,
//   	edition: ec2.amazonLinuxEdition_STANDARD,
//   	virtualization: ec2.amazonLinuxVirt_HVM,
//   	storage: ec2.amazonLinuxStorage_GENERAL_PURPOSE,
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
)

