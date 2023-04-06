package awsnetworkmanager


// Describes the VPC options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcOptionsProperty := &VpcOptionsProperty{
//   	ApplianceModeSupport: jsii.Boolean(false),
//   	Ipv6Support: jsii.Boolean(false),
//   }
//
type CfnVpcAttachment_VpcOptionsProperty struct {
	// Indicates whether appliance mode is supported.
	//
	// If enabled, traffic flow between a source and destination use the same Availability Zone for the VPC attachment for the lifetime of that flow. The default value is `false` .
	ApplianceModeSupport interface{} `field:"optional" json:"applianceModeSupport" yaml:"applianceModeSupport"`
	// Indicates whether IPv6 is supported.
	Ipv6Support interface{} `field:"optional" json:"ipv6Support" yaml:"ipv6Support"`
}

