package awsredshift


// Describes a network interface.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkInterfaceProperty := &NetworkInterfaceProperty{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   	PrivateIpAddress: jsii.String("privateIpAddress"),
//   	SubnetId: jsii.String("subnetId"),
//   }
//
type CfnEndpointAccess_NetworkInterfaceProperty struct {
	// The Availability Zone.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The network interface identifier.
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// The IPv4 address of the network interface within the subnet.
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// The subnet identifier.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

