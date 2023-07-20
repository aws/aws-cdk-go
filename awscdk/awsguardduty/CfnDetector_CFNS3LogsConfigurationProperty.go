package awsguardduty


// Describes whether S3 data event logs will be enabled as a data source when the detector is created.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cFNS3LogsConfigurationProperty := &CFNS3LogsConfigurationProperty{
//   	Enable: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-cfns3logsconfiguration.html
//
type CfnDetector_CFNS3LogsConfigurationProperty struct {
	// The status of S3 data event logs as a data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-cfns3logsconfiguration.html#cfn-guardduty-detector-cfns3logsconfiguration-enable
	//
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
}

