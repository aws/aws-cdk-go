package awsec2


// Describes a secondary private IPv4 address for a network interface.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   privateIpAddressSpecificationProperty := &PrivateIpAddressSpecificationProperty{
//   	Primary: jsii.Boolean(false),
//   	PrivateIpAddress: jsii.String("privateIpAddress"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinterface-privateipaddressspecification.html
//
type CfnNetworkInterface_PrivateIpAddressSpecificationProperty struct {
	// Sets the private IP address as the primary private address.
	//
	// You can set only one primary private IP address. If you don't specify a primary private IP address, Amazon EC2 automatically assigns a primary private IP address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinterface-privateipaddressspecification.html#cfn-ec2-networkinterface-privateipaddressspecification-primary
	//
	Primary interface{} `field:"required" json:"primary" yaml:"primary"`
	// The private IP address of the network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinterface-privateipaddressspecification.html#cfn-ec2-networkinterface-privateipaddressspecification-privateipaddress
	//
	PrivateIpAddress *string `field:"required" json:"privateIpAddress" yaml:"privateIpAddress"`
}

