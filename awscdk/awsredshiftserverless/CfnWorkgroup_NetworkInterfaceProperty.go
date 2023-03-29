package awsredshiftserverless


// Contains information about a network interface in an Amazon Redshift Serverless managed VPC endpoint.
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
type CfnWorkgroup_NetworkInterfaceProperty struct {
	// The availability Zone.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The unique identifier of the network interface.
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// The IPv4 address of the network interface within the subnet.
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// The unique identifier of the subnet.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

