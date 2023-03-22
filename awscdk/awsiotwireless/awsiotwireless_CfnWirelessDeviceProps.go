package awsiotwireless

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnWirelessDevice`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWirelessDeviceProps := &cfnWirelessDeviceProps{
//   	destinationName: jsii.String("destinationName"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	lastUplinkReceivedAt: jsii.String("lastUplinkReceivedAt"),
//   	loRaWan: &loRaWANDeviceProperty{
//   		abpV10X: &abpV10xProperty{
//   			devAddr: jsii.String("devAddr"),
//   			sessionKeys: &sessionKeysAbpV10xProperty{
//   				appSKey: jsii.String("appSKey"),
//   				nwkSKey: jsii.String("nwkSKey"),
//   			},
//   		},
//   		abpV11: &abpV11Property{
//   			devAddr: jsii.String("devAddr"),
//   			sessionKeys: &sessionKeysAbpV11Property{
//   				appSKey: jsii.String("appSKey"),
//   				fNwkSIntKey: jsii.String("fNwkSIntKey"),
//   				nwkSEncKey: jsii.String("nwkSEncKey"),
//   				sNwkSIntKey: jsii.String("sNwkSIntKey"),
//   			},
//   		},
//   		devEui: jsii.String("devEui"),
//   		deviceProfileId: jsii.String("deviceProfileId"),
//   		otaaV10X: &otaaV10xProperty{
//   			appEui: jsii.String("appEui"),
//   			appKey: jsii.String("appKey"),
//   		},
//   		otaaV11: &otaaV11Property{
//   			appKey: jsii.String("appKey"),
//   			joinEui: jsii.String("joinEui"),
//   			nwkKey: jsii.String("nwkKey"),
//   		},
//   		serviceProfileId: jsii.String("serviceProfileId"),
//   	},
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	thingArn: jsii.String("thingArn"),
//   }
//
type CfnWirelessDeviceProps struct {
	// The name of the destination to assign to the new wireless device.
	//
	// Can have only have alphanumeric, - (hyphen) and _ (underscore) characters and it can't have any spaces.
	DestinationName *string `field:"required" json:"destinationName" yaml:"destinationName"`
	// The wireless device type.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The description of the new resource.
	//
	// Maximum length is 2048.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The date and time when the most recent uplink was received.
	LastUplinkReceivedAt *string `field:"optional" json:"lastUplinkReceivedAt" yaml:"lastUplinkReceivedAt"`
	// The device configuration information to use to create the wireless device.
	//
	// Must be at least one of OtaaV10x, OtaaV11, AbpV11, or AbpV10x.
	LoRaWan interface{} `field:"optional" json:"loRaWan" yaml:"loRaWan"`
	// The name of the new resource.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags are an array of key-value pairs to attach to the specified resource.
	//
	// Tags can have a minimum of 0 and a maximum of 50 items.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ARN of the thing to associate with the wireless device.
	ThingArn *string `field:"optional" json:"thingArn" yaml:"thingArn"`
}

