package awsecr


// The image scanning configuration for a repository.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageScanningConfigurationProperty := &imageScanningConfigurationProperty{
//   	scanOnPush: jsii.Boolean(false),
//   }
//
type CfnRepository_ImageScanningConfigurationProperty struct {
	// The setting that determines whether images are scanned after being pushed to a repository.
	//
	// If set to `true` , images will be scanned after being pushed. If this parameter is not specified, it will default to `false` and images will not be scanned unless a scan is manually started.
	ScanOnPush interface{} `field:"optional" json:"scanOnPush" yaml:"scanOnPush"`
}

