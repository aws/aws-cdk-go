package awsec2


// Describes a secondary private IPv4 address for a network interface.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   privateIpAddressSpecificationProperty := &privateIpAddressSpecificationProperty{
//   	privateIpAddress: jsii.String("privateIpAddress"),
//
//   	// the properties below are optional
//   	primary: jsii.Boolean(false),
//   }
//
type CfnSpotFleet_PrivateIpAddressSpecificationProperty struct {
	// The private IPv4 addresses.
	PrivateIpAddress *string `field:"required" json:"privateIpAddress" yaml:"privateIpAddress"`
	// Indicates whether the private IPv4 address is the primary private IPv4 address.
	//
	// Only one IPv4 address can be designated as primary.
	Primary interface{} `field:"optional" json:"primary" yaml:"primary"`
}

