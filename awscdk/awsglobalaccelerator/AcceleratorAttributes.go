package awsglobalaccelerator


// Attributes required to import an existing accelerator to the stack.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   acceleratorAttributes := &AcceleratorAttributes{
//   	AcceleratorArn: jsii.String("acceleratorArn"),
//   	DnsName: jsii.String("dnsName"),
//
//   	// the properties below are optional
//   	DualStackDnsName: jsii.String("dualStackDnsName"),
//   	Ipv4Addresses: []*string{
//   		jsii.String("ipv4Addresses"),
//   	},
//   	Ipv6Addresses: []*string{
//   		jsii.String("ipv6Addresses"),
//   	},
//   }
//
type AcceleratorAttributes struct {
	// The ARN of the accelerator.
	AcceleratorArn *string `field:"required" json:"acceleratorArn" yaml:"acceleratorArn"`
	// The DNS name of the accelerator.
	DnsName *string `field:"required" json:"dnsName" yaml:"dnsName"`
	// The DNS name that points to the dual-stack accelerator's four static IP addresses: two IPv4 addresses and two IPv6 addresses.
	// Default: - undefined.
	//
	DualStackDnsName *string `field:"optional" json:"dualStackDnsName" yaml:"dualStackDnsName"`
	// The array of IPv4 addresses in the IP address set.
	// Default: - undefined.
	//
	Ipv4Addresses *[]*string `field:"optional" json:"ipv4Addresses" yaml:"ipv4Addresses"`
	// The array of IPv6 addresses in the IP address set.
	// Default: - undefined.
	//
	Ipv6Addresses *[]*string `field:"optional" json:"ipv6Addresses" yaml:"ipv6Addresses"`
}

