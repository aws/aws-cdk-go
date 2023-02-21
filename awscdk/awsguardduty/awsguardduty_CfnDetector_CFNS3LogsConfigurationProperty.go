package awsguardduty


// Describes whether S3 data event logs will be enabled as a data source when the detector is created.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cFNS3LogsConfigurationProperty := &cFNS3LogsConfigurationProperty{
//   	enable: jsii.Boolean(false),
//   }
//
type CfnDetector_CFNS3LogsConfigurationProperty struct {
	// The status of S3 data event logs as a data source.
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
}

