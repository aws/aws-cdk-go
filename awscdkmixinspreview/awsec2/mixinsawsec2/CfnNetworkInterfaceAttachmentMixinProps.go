package mixinsawsec2


// Properties for CfnNetworkInterfaceAttachmentPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnNetworkInterfaceAttachmentMixinProps := &CfnNetworkInterfaceAttachmentMixinProps{
//   	DeleteOnTermination: jsii.Boolean(false),
//   	DeviceIndex: jsii.String("deviceIndex"),
//   	EnaQueueCount: jsii.Number(123),
//   	EnaSrdSpecification: &EnaSrdSpecificationProperty{
//   		EnaSrdEnabled: jsii.Boolean(false),
//   		EnaSrdUdpSpecification: &EnaSrdUdpSpecificationProperty{
//   			EnaSrdUdpEnabled: jsii.Boolean(false),
//   		},
//   	},
//   	InstanceId: jsii.String("instanceId"),
//   	NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterfaceattachment.html
//
type CfnNetworkInterfaceAttachmentMixinProps struct {
	// Whether to delete the network interface when the instance terminates.
	//
	// By default, this value is set to `true` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterfaceattachment.html#cfn-ec2-networkinterfaceattachment-deleteontermination
	//
	// Default: - true.
	//
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// The network interface's position in the attachment order.
	//
	// For example, the first attached network interface has a `DeviceIndex` of 0.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterfaceattachment.html#cfn-ec2-networkinterfaceattachment-deviceindex
	//
	DeviceIndex *string `field:"optional" json:"deviceIndex" yaml:"deviceIndex"`
	// The number of ENA queues created with the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterfaceattachment.html#cfn-ec2-networkinterfaceattachment-enaqueuecount
	//
	EnaQueueCount *float64 `field:"optional" json:"enaQueueCount" yaml:"enaQueueCount"`
	// Configures ENA Express for the network interface that this action attaches to the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterfaceattachment.html#cfn-ec2-networkinterfaceattachment-enasrdspecification
	//
	EnaSrdSpecification interface{} `field:"optional" json:"enaSrdSpecification" yaml:"enaSrdSpecification"`
	// The ID of the instance to which you will attach the ENI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterfaceattachment.html#cfn-ec2-networkinterfaceattachment-instanceid
	//
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// The ID of the ENI that you want to attach.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterfaceattachment.html#cfn-ec2-networkinterfaceattachment-networkinterfaceid
	//
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
}

