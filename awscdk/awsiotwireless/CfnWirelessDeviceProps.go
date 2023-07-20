package awsiotwireless

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnWirelessDevice`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWirelessDeviceProps := &CfnWirelessDeviceProps{
//   	DestinationName: jsii.String("destinationName"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	LastUplinkReceivedAt: jsii.String("lastUplinkReceivedAt"),
//   	LoRaWan: &LoRaWANDeviceProperty{
//   		AbpV10X: &AbpV10xProperty{
//   			DevAddr: jsii.String("devAddr"),
//   			SessionKeys: &SessionKeysAbpV10xProperty{
//   				AppSKey: jsii.String("appSKey"),
//   				NwkSKey: jsii.String("nwkSKey"),
//   			},
//   		},
//   		AbpV11: &AbpV11Property{
//   			DevAddr: jsii.String("devAddr"),
//   			SessionKeys: &SessionKeysAbpV11Property{
//   				AppSKey: jsii.String("appSKey"),
//   				FNwkSIntKey: jsii.String("fNwkSIntKey"),
//   				NwkSEncKey: jsii.String("nwkSEncKey"),
//   				SNwkSIntKey: jsii.String("sNwkSIntKey"),
//   			},
//   		},
//   		DevEui: jsii.String("devEui"),
//   		DeviceProfileId: jsii.String("deviceProfileId"),
//   		OtaaV10X: &OtaaV10xProperty{
//   			AppEui: jsii.String("appEui"),
//   			AppKey: jsii.String("appKey"),
//   		},
//   		OtaaV11: &OtaaV11Property{
//   			AppKey: jsii.String("appKey"),
//   			JoinEui: jsii.String("joinEui"),
//   			NwkKey: jsii.String("nwkKey"),
//   		},
//   		ServiceProfileId: jsii.String("serviceProfileId"),
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ThingArn: jsii.String("thingArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-wirelessdevice.html
//
type CfnWirelessDeviceProps struct {
	// The name of the destination to assign to the new wireless device.
	//
	// Can have only have alphanumeric, - (hyphen) and _ (underscore) characters and it can't have any spaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-wirelessdevice.html#cfn-iotwireless-wirelessdevice-destinationname
	//
	DestinationName *string `field:"required" json:"destinationName" yaml:"destinationName"`
	// The wireless device type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-wirelessdevice.html#cfn-iotwireless-wirelessdevice-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The description of the new resource.
	//
	// Maximum length is 2048.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-wirelessdevice.html#cfn-iotwireless-wirelessdevice-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The date and time when the most recent uplink was received.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-wirelessdevice.html#cfn-iotwireless-wirelessdevice-lastuplinkreceivedat
	//
	LastUplinkReceivedAt *string `field:"optional" json:"lastUplinkReceivedAt" yaml:"lastUplinkReceivedAt"`
	// The device configuration information to use to create the wireless device.
	//
	// Must be at least one of OtaaV10x, OtaaV11, AbpV11, or AbpV10x.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-wirelessdevice.html#cfn-iotwireless-wirelessdevice-lorawan
	//
	LoRaWan interface{} `field:"optional" json:"loRaWan" yaml:"loRaWan"`
	// The name of the new resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-wirelessdevice.html#cfn-iotwireless-wirelessdevice-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags are an array of key-value pairs to attach to the specified resource.
	//
	// Tags can have a minimum of 0 and a maximum of 50 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-wirelessdevice.html#cfn-iotwireless-wirelessdevice-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ARN of the thing to associate with the wireless device.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-wirelessdevice.html#cfn-iotwireless-wirelessdevice-thingarn
	//
	ThingArn *string `field:"optional" json:"thingArn" yaml:"thingArn"`
}

