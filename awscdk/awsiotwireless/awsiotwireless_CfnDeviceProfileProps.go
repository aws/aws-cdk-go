package awsiotwireless

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDeviceProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeviceProfileProps := &cfnDeviceProfileProps{
//   	loRaWan: &loRaWANDeviceProfileProperty{
//   		classBTimeout: jsii.Number(123),
//   		classCTimeout: jsii.Number(123),
//   		factoryPresetFreqsList: []interface{}{
//   			jsii.Number(123),
//   		},
//   		macVersion: jsii.String("macVersion"),
//   		maxDutyCycle: jsii.Number(123),
//   		maxEirp: jsii.Number(123),
//   		pingSlotDr: jsii.Number(123),
//   		pingSlotFreq: jsii.Number(123),
//   		pingSlotPeriod: jsii.Number(123),
//   		regParamsRevision: jsii.String("regParamsRevision"),
//   		rfRegion: jsii.String("rfRegion"),
//   		rxDataRate2: jsii.Number(123),
//   		rxDelay1: jsii.Number(123),
//   		rxDrOffset1: jsii.Number(123),
//   		rxFreq2: jsii.Number(123),
//   		supports32BitFCnt: jsii.Boolean(false),
//   		supportsClassB: jsii.Boolean(false),
//   		supportsClassC: jsii.Boolean(false),
//   		supportsJoin: jsii.Boolean(false),
//   	},
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDeviceProfileProps struct {
	// LoRaWAN device profile object.
	LoRaWan interface{} `field:"optional" json:"loRaWan" yaml:"loRaWan"`
	// The name of the new resource.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags are an array of key-value pairs to attach to the specified resource.
	//
	// Tags can have a minimum of 0 and a maximum of 50 items.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

