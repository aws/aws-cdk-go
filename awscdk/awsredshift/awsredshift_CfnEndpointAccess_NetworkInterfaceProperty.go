package awsredshift


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkInterfaceProperty := &networkInterfaceProperty{
//   	availabilityZone: jsii.String("availabilityZone"),
//   	networkInterfaceId: jsii.String("networkInterfaceId"),
//   	privateIpAddress: jsii.String("privateIpAddress"),
//   	subnetId: jsii.String("subnetId"),
//   }
//
type CfnEndpointAccess_NetworkInterfaceProperty struct {
	// `CfnEndpointAccess.NetworkInterfaceProperty.AvailabilityZone`.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// `CfnEndpointAccess.NetworkInterfaceProperty.NetworkInterfaceId`.
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// `CfnEndpointAccess.NetworkInterfaceProperty.PrivateIpAddress`.
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// `CfnEndpointAccess.NetworkInterfaceProperty.SubnetId`.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

