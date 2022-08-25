package awsguardduty


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cFNScanEc2InstanceWithFindingsConfigurationProperty := &cFNScanEc2InstanceWithFindingsConfigurationProperty{
//   	ebsVolumes: jsii.Boolean(false),
//   }
//
type CfnDetector_CFNScanEc2InstanceWithFindingsConfigurationProperty struct {
	// `CfnDetector.CFNScanEc2InstanceWithFindingsConfigurationProperty.EbsVolumes`.
	EbsVolumes interface{} `field:"optional" json:"ebsVolumes" yaml:"ebsVolumes"`
}

