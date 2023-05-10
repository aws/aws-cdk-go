package awsquicksight


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
	// `CfnVPCConnection.NetworkInterfaceProperty.AvailabilityZone`.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// `CfnVPCConnection.NetworkInterfaceProperty.ErrorMessage`.
	ErrorMessage *string `field:"optional" json:"errorMessage" yaml:"errorMessage"`
	// `CfnVPCConnection.NetworkInterfaceProperty.NetworkInterfaceId`.
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// `CfnVPCConnection.NetworkInterfaceProperty.Status`.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// `CfnVPCConnection.NetworkInterfaceProperty.SubnetId`.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

