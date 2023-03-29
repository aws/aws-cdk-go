package awsec2


// Specifies a secondary private IPv4 address for a network interface.
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
type CfnInstance_PrivateIpAddressSpecificationProperty struct {
	// Indicates whether the private IPv4 address is the primary private IPv4 address.
	//
	// Only one IPv4 address can be designated as primary.
	Primary interface{} `field:"required" json:"primary" yaml:"primary"`
	// The private IPv4 address.
	PrivateIpAddress *string `field:"required" json:"privateIpAddress" yaml:"privateIpAddress"`
}

