package awsdeadline


// The IP addresses for a host.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ipAddressesProperty := &IpAddressesProperty{
//   	IpV4Addresses: []*string{
//   		jsii.String("ipV4Addresses"),
//   	},
//   	IpV6Addresses: []*string{
//   		jsii.String("ipV6Addresses"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-worker-ipaddresses.html
//
type CfnWorker_IpAddressesProperty struct {
	// The IpV4 address of the network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-worker-ipaddresses.html#cfn-deadline-worker-ipaddresses-ipv4addresses
	//
	IpV4Addresses *[]*string `field:"optional" json:"ipV4Addresses" yaml:"ipV4Addresses"`
	// The IpV6 address for the network and node component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-worker-ipaddresses.html#cfn-deadline-worker-ipaddresses-ipv6addresses
	//
	IpV6Addresses *[]*string `field:"optional" json:"ipV6Addresses" yaml:"ipV6Addresses"`
}

