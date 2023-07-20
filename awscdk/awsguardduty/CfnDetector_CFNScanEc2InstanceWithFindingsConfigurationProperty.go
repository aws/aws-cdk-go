package awsguardduty


// Describes whether Malware Protection for EC2 instances with findings will be enabled as a data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cFNScanEc2InstanceWithFindingsConfigurationProperty := &CFNScanEc2InstanceWithFindingsConfigurationProperty{
//   	EbsVolumes: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-cfnscanec2instancewithfindingsconfiguration.html
//
type CfnDetector_CFNScanEc2InstanceWithFindingsConfigurationProperty struct {
	// Describes the configuration for scanning EBS volumes as data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-cfnscanec2instancewithfindingsconfiguration.html#cfn-guardduty-detector-cfnscanec2instancewithfindingsconfiguration-ebsvolumes
	//
	EbsVolumes interface{} `field:"optional" json:"ebsVolumes" yaml:"ebsVolumes"`
}

