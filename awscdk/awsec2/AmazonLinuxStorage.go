package awsec2


// Available storage options for Amazon Linux images Only applies to Amazon Linux & Amazon Linux 2.
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
type AmazonLinuxStorage string

const (
	// EBS-backed storage.
	AmazonLinuxStorage_EBS AmazonLinuxStorage = "EBS"
	// S3-backed storage.
	AmazonLinuxStorage_S3 AmazonLinuxStorage = "S3"
	// General Purpose-based storage (recommended).
	AmazonLinuxStorage_GENERAL_PURPOSE AmazonLinuxStorage = "GENERAL_PURPOSE"
)

