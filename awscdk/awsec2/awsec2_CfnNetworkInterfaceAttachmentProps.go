package awsec2


// Properties for defining a `CfnNetworkInterfaceAttachment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNetworkInterfaceAttachmentProps := &cfnNetworkInterfaceAttachmentProps{
//   	deviceIndex: jsii.String("deviceIndex"),
//   	instanceId: jsii.String("instanceId"),
//   	networkInterfaceId: jsii.String("networkInterfaceId"),
//
//   	// the properties below are optional
//   	deleteOnTermination: jsii.Boolean(false),
//   }
//
type CfnNetworkInterfaceAttachmentProps struct {
	// The network interface's position in the attachment order.
	//
	// For example, the first attached network interface has a `DeviceIndex` of 0.
	DeviceIndex *string `field:"required" json:"deviceIndex" yaml:"deviceIndex"`
	// The ID of the instance to which you will attach the ENI.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// The ID of the ENI that you want to attach.
	NetworkInterfaceId *string `field:"required" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// Whether to delete the network interface when the instance terminates.
	//
	// By default, this value is set to `true` .
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
}

