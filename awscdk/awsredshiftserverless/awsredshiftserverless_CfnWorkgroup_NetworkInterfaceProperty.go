package awsredshiftserverless


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
type CfnWorkgroup_NetworkInterfaceProperty struct {
	// `CfnWorkgroup.NetworkInterfaceProperty.AvailabilityZone`.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// `CfnWorkgroup.NetworkInterfaceProperty.NetworkInterfaceId`.
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// `CfnWorkgroup.NetworkInterfaceProperty.PrivateIpAddress`.
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// `CfnWorkgroup.NetworkInterfaceProperty.SubnetId`.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

