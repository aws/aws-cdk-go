package awsnetworkmanager


// Describes the VPC options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcOptionsProperty := &vpcOptionsProperty{
//   	ipv6Support: jsii.Boolean(false),
//   }
//
type CfnVpcAttachment_VpcOptionsProperty struct {
	// Indicates whether IPv6 is supported.
	Ipv6Support interface{} `field:"optional" json:"ipv6Support" yaml:"ipv6Support"`
}

