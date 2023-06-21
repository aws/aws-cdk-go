package awsiotwireless

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDeviceProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeviceProfileProps := &CfnDeviceProfileProps{
//   	LoRaWan: &LoRaWANDeviceProfileProperty{
//   		ClassBTimeout: jsii.Number(123),
//   		ClassCTimeout: jsii.Number(123),
//   		FactoryPresetFreqsList: []interface{}{
//   			jsii.Number(123),
//   		},
//   		MacVersion: jsii.String("macVersion"),
//   		MaxDutyCycle: jsii.Number(123),
//   		MaxEirp: jsii.Number(123),
//   		PingSlotDr: jsii.Number(123),
//   		PingSlotFreq: jsii.Number(123),
//   		PingSlotPeriod: jsii.Number(123),
//   		RegParamsRevision: jsii.String("regParamsRevision"),
//   		RfRegion: jsii.String("rfRegion"),
//   		RxDataRate2: jsii.Number(123),
//   		RxDelay1: jsii.Number(123),
//   		RxDrOffset1: jsii.Number(123),
//   		RxFreq2: jsii.Number(123),
//   		Supports32BitFCnt: jsii.Boolean(false),
//   		SupportsClassB: jsii.Boolean(false),
//   		SupportsClassC: jsii.Boolean(false),
//   		SupportsJoin: jsii.Boolean(false),
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
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

