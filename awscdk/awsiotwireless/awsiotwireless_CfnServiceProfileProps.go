package awsiotwireless

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnServiceProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServiceProfileProps := &CfnServiceProfileProps{
//   	LoRaWan: &LoRaWANServiceProfileProperty{
//   		AddGwMetadata: jsii.Boolean(false),
//   		ChannelMask: jsii.String("channelMask"),
//   		DevStatusReqFreq: jsii.Number(123),
//   		DlBucketSize: jsii.Number(123),
//   		DlRate: jsii.Number(123),
//   		DlRatePolicy: jsii.String("dlRatePolicy"),
//   		DrMax: jsii.Number(123),
//   		DrMin: jsii.Number(123),
//   		HrAllowed: jsii.Boolean(false),
//   		MinGwDiversity: jsii.Number(123),
//   		NwkGeoLoc: jsii.Boolean(false),
//   		PrAllowed: jsii.Boolean(false),
//   		RaAllowed: jsii.Boolean(false),
//   		ReportDevStatusBattery: jsii.Boolean(false),
//   		ReportDevStatusMargin: jsii.Boolean(false),
//   		TargetPer: jsii.Number(123),
//   		UlBucketSize: jsii.Number(123),
//   		UlRate: jsii.Number(123),
//   		UlRatePolicy: jsii.String("ulRatePolicy"),
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
type CfnServiceProfileProps struct {
	// LoRaWAN service profile object.
	LoRaWan interface{} `field:"optional" json:"loRaWan" yaml:"loRaWan"`
	// The name of the new resource.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags are an array of key-value pairs to attach to the specified resource.
	//
	// Tags can have a minimum of 0 and a maximum of 50 items.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

