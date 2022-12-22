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
//   cfnServiceProfileProps := &cfnServiceProfileProps{
//   	loRaWan: &loRaWANServiceProfileProperty{
//   		addGwMetadata: jsii.Boolean(false),
//   		channelMask: jsii.String("channelMask"),
//   		devStatusReqFreq: jsii.Number(123),
//   		dlBucketSize: jsii.Number(123),
//   		dlRate: jsii.Number(123),
//   		dlRatePolicy: jsii.String("dlRatePolicy"),
//   		drMax: jsii.Number(123),
//   		drMin: jsii.Number(123),
//   		hrAllowed: jsii.Boolean(false),
//   		minGwDiversity: jsii.Number(123),
//   		nwkGeoLoc: jsii.Boolean(false),
//   		prAllowed: jsii.Boolean(false),
//   		raAllowed: jsii.Boolean(false),
//   		reportDevStatusBattery: jsii.Boolean(false),
//   		reportDevStatusMargin: jsii.Boolean(false),
//   		targetPer: jsii.Number(123),
//   		ulBucketSize: jsii.Number(123),
//   		ulRate: jsii.Number(123),
//   		ulRatePolicy: jsii.String("ulRatePolicy"),
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

