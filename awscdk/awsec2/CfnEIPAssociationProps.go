package awsec2


// Properties for defining a `CfnEIPAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEIPAssociationProps := &CfnEIPAssociationProps{
//   	AllocationId: jsii.String("allocationId"),
//   	Eip: jsii.String("eip"),
//   	InstanceId: jsii.String("instanceId"),
//   	NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   	PrivateIpAddress: jsii.String("privateIpAddress"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-eipassociation.html
//
type CfnEIPAssociationProps struct {
	// The allocation ID.
	//
	// This is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-eipassociation.html#cfn-ec2-eipassociation-allocationid
	//
	AllocationId *string `field:"optional" json:"allocationId" yaml:"allocationId"`
	// The Elastic IP address to associate with the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-eipassociation.html#cfn-ec2-eipassociation-eip
	//
	// Deprecated: this property has been deprecated.
	Eip *string `field:"optional" json:"eip" yaml:"eip"`
	// The ID of the instance.
	//
	// The instance must have exactly one attached network interface. You can specify either the instance ID or the network interface ID, but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-eipassociation.html#cfn-ec2-eipassociation-instanceid
	//
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// The ID of the network interface.
	//
	// If the instance has more than one network interface, you must specify a network interface ID.
	//
	// You can specify either the instance ID or the network interface ID, but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-eipassociation.html#cfn-ec2-eipassociation-networkinterfaceid
	//
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// The primary or secondary private IP address to associate with the Elastic IP address.
	//
	// If no private IP address is specified, the Elastic IP address is associated with the primary private IP address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-eipassociation.html#cfn-ec2-eipassociation-privateipaddress
	//
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
}

