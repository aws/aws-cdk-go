package awsec2


// Properties for defining a `CfnNetworkInterfaceAttachment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNetworkInterfaceAttachmentProps := &CfnNetworkInterfaceAttachmentProps{
//   	DeviceIndex: jsii.String("deviceIndex"),
//   	InstanceId: jsii.String("instanceId"),
//   	NetworkInterfaceId: jsii.String("networkInterfaceId"),
//
//   	// the properties below are optional
//   	DeleteOnTermination: jsii.Boolean(false),
//   	EnaSrdSpecification: &EnaSrdSpecificationProperty{
//   		EnaSrdEnabled: jsii.Boolean(false),
//   		EnaSrdUdpSpecification: &EnaSrdUdpSpecificationProperty{
//   			EnaSrdUdpEnabled: jsii.Boolean(false),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterfaceattachment.html
//
type CfnNetworkInterfaceAttachmentProps struct {
	// The network interface's position in the attachment order.
	//
	// For example, the first attached network interface has a `DeviceIndex` of 0.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterfaceattachment.html#cfn-ec2-networkinterfaceattachment-deviceindex
	//
	DeviceIndex *string `field:"required" json:"deviceIndex" yaml:"deviceIndex"`
	// The ID of the instance to which you will attach the ENI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterfaceattachment.html#cfn-ec2-networkinterfaceattachment-instanceid
	//
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// The ID of the ENI that you want to attach.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterfaceattachment.html#cfn-ec2-networkinterfaceattachment-networkinterfaceid
	//
	NetworkInterfaceId *string `field:"required" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// Whether to delete the network interface when the instance terminates.
	//
	// By default, this value is set to `true` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterfaceattachment.html#cfn-ec2-networkinterfaceattachment-deleteontermination
	//
	// Default: - true.
	//
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// Configures ENA Express for the network interface that this action attaches to the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterfaceattachment.html#cfn-ec2-networkinterfaceattachment-enasrdspecification
	//
	EnaSrdSpecification interface{} `field:"optional" json:"enaSrdSpecification" yaml:"enaSrdSpecification"`
}

