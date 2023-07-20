package awsec2


// Describes a secondary private IPv4 address for a network interface.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   privateIpAddressSpecificationProperty := &PrivateIpAddressSpecificationProperty{
//   	PrivateIpAddress: jsii.String("privateIpAddress"),
//
//   	// the properties below are optional
//   	Primary: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-privateipaddressspecification.html
//
type CfnSpotFleet_PrivateIpAddressSpecificationProperty struct {
	// The private IPv4 address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-privateipaddressspecification.html#cfn-ec2-spotfleet-privateipaddressspecification-privateipaddress
	//
	PrivateIpAddress *string `field:"required" json:"privateIpAddress" yaml:"privateIpAddress"`
	// Indicates whether the private IPv4 address is the primary private IPv4 address.
	//
	// Only one IPv4 address can be designated as primary.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-privateipaddressspecification.html#cfn-ec2-spotfleet-privateipaddressspecification-primary
	//
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
}

