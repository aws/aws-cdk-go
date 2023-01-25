package awsiotwireless

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnFuotaTask`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFuotaTaskProps := &cfnFuotaTaskProps{
//   	firmwareUpdateImage: jsii.String("firmwareUpdateImage"),
//   	firmwareUpdateRole: jsii.String("firmwareUpdateRole"),
//   	loRaWan: &loRaWANProperty{
//   		rfRegion: jsii.String("rfRegion"),
//
//   		// the properties below are optional
//   		startTime: jsii.String("startTime"),
//   	},
//
//   	// the properties below are optional
//   	associateMulticastGroup: jsii.String("associateMulticastGroup"),
//   	associateWirelessDevice: jsii.String("associateWirelessDevice"),
//   	description: jsii.String("description"),
//   	disassociateMulticastGroup: jsii.String("disassociateMulticastGroup"),
//   	disassociateWirelessDevice: jsii.String("disassociateWirelessDevice"),
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnFuotaTaskProps struct {
	// The S3 URI points to a firmware update image that is to be used with a FUOTA task.
	FirmwareUpdateImage *string `field:"required" json:"firmwareUpdateImage" yaml:"firmwareUpdateImage"`
	// The firmware update role that is to be used with a FUOTA task.
	FirmwareUpdateRole *string `field:"required" json:"firmwareUpdateRole" yaml:"firmwareUpdateRole"`
	// The LoRaWAN information used with a FUOTA task.
	LoRaWan interface{} `field:"required" json:"loRaWan" yaml:"loRaWan"`
	// The ID of the multicast group to associate with a FUOTA task.
	AssociateMulticastGroup *string `field:"optional" json:"associateMulticastGroup" yaml:"associateMulticastGroup"`
	// The ID of the wireless device to associate with a multicast group.
	AssociateWirelessDevice *string `field:"optional" json:"associateWirelessDevice" yaml:"associateWirelessDevice"`
	// The description of the new resource.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of the multicast group to disassociate from a FUOTA task.
	DisassociateMulticastGroup *string `field:"optional" json:"disassociateMulticastGroup" yaml:"disassociateMulticastGroup"`
	// The ID of the wireless device to disassociate from a FUOTA task.
	DisassociateWirelessDevice *string `field:"optional" json:"disassociateWirelessDevice" yaml:"disassociateWirelessDevice"`
	// The name of a FUOTA task.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags are an array of key-value pairs to attach to the specified resource.
	//
	// Tags can have a minimum of 0 and a maximum of 50 items.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

