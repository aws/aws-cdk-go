package previewawsiotwirelessmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnFuotaTaskPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFuotaTaskMixinProps := &CfnFuotaTaskMixinProps{
//   	AssociateMulticastGroup: jsii.String("associateMulticastGroup"),
//   	AssociateWirelessDevice: jsii.String("associateWirelessDevice"),
//   	Description: jsii.String("description"),
//   	DisassociateMulticastGroup: jsii.String("disassociateMulticastGroup"),
//   	DisassociateWirelessDevice: jsii.String("disassociateWirelessDevice"),
//   	FirmwareUpdateImage: jsii.String("firmwareUpdateImage"),
//   	FirmwareUpdateRole: jsii.String("firmwareUpdateRole"),
//   	LoRaWan: &LoRaWANProperty{
//   		RfRegion: jsii.String("rfRegion"),
//   		StartTime: jsii.String("startTime"),
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-fuotatask.html
//
type CfnFuotaTaskMixinProps struct {
	// The ID of the multicast group to associate with a FUOTA task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-fuotatask.html#cfn-iotwireless-fuotatask-associatemulticastgroup
	//
	AssociateMulticastGroup *string `field:"optional" json:"associateMulticastGroup" yaml:"associateMulticastGroup"`
	// The ID of the wireless device to associate with a multicast group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-fuotatask.html#cfn-iotwireless-fuotatask-associatewirelessdevice
	//
	AssociateWirelessDevice *string `field:"optional" json:"associateWirelessDevice" yaml:"associateWirelessDevice"`
	// The description of the new resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-fuotatask.html#cfn-iotwireless-fuotatask-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of the multicast group to disassociate from a FUOTA task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-fuotatask.html#cfn-iotwireless-fuotatask-disassociatemulticastgroup
	//
	DisassociateMulticastGroup *string `field:"optional" json:"disassociateMulticastGroup" yaml:"disassociateMulticastGroup"`
	// The ID of the wireless device to disassociate from a FUOTA task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-fuotatask.html#cfn-iotwireless-fuotatask-disassociatewirelessdevice
	//
	DisassociateWirelessDevice *string `field:"optional" json:"disassociateWirelessDevice" yaml:"disassociateWirelessDevice"`
	// The S3 URI points to a firmware update image that is to be used with a FUOTA task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-fuotatask.html#cfn-iotwireless-fuotatask-firmwareupdateimage
	//
	FirmwareUpdateImage *string `field:"optional" json:"firmwareUpdateImage" yaml:"firmwareUpdateImage"`
	// The firmware update role that is to be used with a FUOTA task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-fuotatask.html#cfn-iotwireless-fuotatask-firmwareupdaterole
	//
	FirmwareUpdateRole *string `field:"optional" json:"firmwareUpdateRole" yaml:"firmwareUpdateRole"`
	// The LoRaWAN information used with a FUOTA task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-fuotatask.html#cfn-iotwireless-fuotatask-lorawan
	//
	LoRaWan interface{} `field:"optional" json:"loRaWan" yaml:"loRaWan"`
	// The name of a FUOTA task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-fuotatask.html#cfn-iotwireless-fuotatask-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags are an array of key-value pairs to attach to the specified resource.
	//
	// Tags can have a minimum of 0 and a maximum of 50 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-fuotatask.html#cfn-iotwireless-fuotatask-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

