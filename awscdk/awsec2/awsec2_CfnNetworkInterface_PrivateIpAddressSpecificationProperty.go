package awsec2


// Describes a secondary private IPv4 address for a network interface.
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
type CfnNetworkInterface_PrivateIpAddressSpecificationProperty struct {
	// Sets the private IP address as the primary private address.
	//
	// You can set only one primary private IP address. If you don't specify a primary private IP address, Amazon EC2 automatically assigns a primary private IP address.
	Primary interface{} `field:"required" json:"primary" yaml:"primary"`
	// The private IP address of the network interface.
	PrivateIpAddress *string `field:"required" json:"privateIpAddress" yaml:"privateIpAddress"`
}

