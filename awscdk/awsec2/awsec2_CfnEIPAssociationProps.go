package awsec2


// Properties for defining a `CfnEIPAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEIPAssociationProps := &cfnEIPAssociationProps{
//   	allocationId: jsii.String("allocationId"),
//   	eip: jsii.String("eip"),
//   	instanceId: jsii.String("instanceId"),
//   	networkInterfaceId: jsii.String("networkInterfaceId"),
//   	privateIpAddress: jsii.String("privateIpAddress"),
//   }
//
type CfnEIPAssociationProps struct {
	// [EC2-VPC] The allocation ID.
	//
	// This is required for EC2-VPC.
	AllocationId *string `field:"optional" json:"allocationId" yaml:"allocationId"`
	// [EC2-Classic] The Elastic IP address to associate with the instance.
	//
	// This is required for EC2-Classic.
	Eip *string `field:"optional" json:"eip" yaml:"eip"`
	// The ID of the instance.
	//
	// The instance must have exactly one attached network interface. For EC2-VPC, you can specify either the instance ID or the network interface ID, but not both. For EC2-Classic, you must specify an instance ID and the instance must be in the running state.
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// [EC2-VPC] The ID of the network interface.
	//
	// If the instance has more than one network interface, you must specify a network interface ID.
	//
	// For EC2-VPC, you can specify either the instance ID or the network interface ID, but not both.
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// [EC2-VPC] The primary or secondary private IP address to associate with the Elastic IP address.
	//
	// If no private IP address is specified, the Elastic IP address is associated with the primary private IP address.
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
}

