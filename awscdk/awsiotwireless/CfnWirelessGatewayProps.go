package awsiotwireless

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnWirelessGateway`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWirelessGatewayProps := &CfnWirelessGatewayProps{
//   	LoRaWan: &LoRaWANGatewayProperty{
//   		GatewayEui: jsii.String("gatewayEui"),
//   		RfRegion: jsii.String("rfRegion"),
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	LastUplinkReceivedAt: jsii.String("lastUplinkReceivedAt"),
//   	Name: jsii.String("name"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ThingArn: jsii.String("thingArn"),
//   	ThingName: jsii.String("thingName"),
//   }
//
type CfnWirelessGatewayProps struct {
	// The gateway configuration information to use to create the wireless gateway.
	LoRaWan interface{} `field:"required" json:"loRaWan" yaml:"loRaWan"`
	// The description of the new resource.
	//
	// The maximum length is 2048 characters.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The date and time when the most recent uplink was received.
	LastUplinkReceivedAt *string `field:"optional" json:"lastUplinkReceivedAt" yaml:"lastUplinkReceivedAt"`
	// The name of the new resource.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags are an array of key-value pairs to attach to the specified resource.
	//
	// Tags can have a minimum of 0 and a maximum of 50 items.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ARN of the thing to associate with the wireless gateway.
	ThingArn *string `field:"optional" json:"thingArn" yaml:"thingArn"`
	// The name of the thing associated with the wireless gateway.
	//
	// The value is empty if a thing isn't associated with the gateway.
	ThingName *string `field:"optional" json:"thingName" yaml:"thingName"`
}

