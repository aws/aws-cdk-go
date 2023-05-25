package awsquicksight


// The structure that contains information about a network interface.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkInterfaceProperty := &NetworkInterfaceProperty{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	ErrorMessage: jsii.String("errorMessage"),
//   	NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   	Status: jsii.String("status"),
//   	SubnetId: jsii.String("subnetId"),
//   }
//
type CfnVPCConnection_NetworkInterfaceProperty struct {
	// The availability zone that the network interface resides in.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// An error message.
	ErrorMessage *string `field:"optional" json:"errorMessage" yaml:"errorMessage"`
	// The network interface ID.
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// The status of the network interface.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The subnet ID associated with the network interface.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

