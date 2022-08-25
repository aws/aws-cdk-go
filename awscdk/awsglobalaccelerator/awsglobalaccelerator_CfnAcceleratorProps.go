package awsglobalaccelerator

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnAccelerator`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAcceleratorProps := &cfnAcceleratorProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	enabled: jsii.Boolean(false),
//   	ipAddresses: []*string{
//   		jsii.String("ipAddresses"),
//   	},
//   	ipAddressType: jsii.String("ipAddressType"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnAcceleratorProps struct {
	// The name of the accelerator.
	//
	// The name must contain only alphanumeric characters or hyphens (-), and must not begin or end with a hyphen.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Indicates whether the accelerator is enabled. The value is true or false. The default value is true.
	//
	// If the value is set to true, the accelerator cannot be deleted. If set to false, accelerator can be deleted.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Optionally, if you've added your own IP address pool to Global Accelerator (BYOIP), you can choose IP addresses from your own pool to use for the accelerator's static IP addresses when you create an accelerator.
	//
	// You can specify one or two addresses, separated by a comma. Do not include the /32 suffix.
	//
	// Only one IP address from each of your IP address ranges can be used for each accelerator. If you specify only one IP address from your IP address range, Global Accelerator assigns a second static IP address for the accelerator from the AWS IP address pool.
	//
	// Note that you can't update IP addresses for an existing accelerator. To change them, you must create a new accelerator with the new addresses.
	//
	// For more information, see [Bring Your Own IP Addresses (BYOIP)](https://docs.aws.amazon.com/global-accelerator/latest/dg/using-byoip.html) in the *AWS Global Accelerator Developer Guide* .
	IpAddresses *[]*string `field:"optional" json:"ipAddresses" yaml:"ipAddresses"`
	// The value for the address type must be IPv4.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Create tags for an accelerator.
	//
	// For more information, see [Tagging](https://docs.aws.amazon.com/global-accelerator/latest/dg/tagging-in-global-accelerator.html) in the *AWS Global Accelerator Developer Guide* .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

