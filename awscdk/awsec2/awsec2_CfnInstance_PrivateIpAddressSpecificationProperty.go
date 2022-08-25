package awsec2


// Specifies a secondary private IPv4 address for a network interface.
//
// `PrivateIpAddressSpecification` is a property of the [AWS::EC2::NetworkInterface](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   privateIpAddressSpecificationProperty := &privateIpAddressSpecificationProperty{
//   	primary: jsii.Boolean(false),
//   	privateIpAddress: jsii.String("privateIpAddress"),
//   }
//
type CfnInstance_PrivateIpAddressSpecificationProperty struct {
	// Indicates whether the private IPv4 address is the primary private IPv4 address.
	//
	// Only one IPv4 address can be designated as primary.
	Primary interface{} `field:"required" json:"primary" yaml:"primary"`
	// The private IPv4 addresses.
	PrivateIpAddress *string `field:"required" json:"privateIpAddress" yaml:"privateIpAddress"`
}

