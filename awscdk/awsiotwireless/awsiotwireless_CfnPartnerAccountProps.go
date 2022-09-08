package awsiotwireless

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnPartnerAccount`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPartnerAccountProps := &cfnPartnerAccountProps{
//   	accountLinked: jsii.Boolean(false),
//   	fingerprint: jsii.String("fingerprint"),
//   	partnerAccountId: jsii.String("partnerAccountId"),
//   	partnerType: jsii.String("partnerType"),
//   	sidewalk: &sidewalkAccountInfoProperty{
//   		appServerPrivateKey: jsii.String("appServerPrivateKey"),
//   	},
//   	sidewalkUpdate: &sidewalkUpdateAccountProperty{
//   		appServerPrivateKey: jsii.String("appServerPrivateKey"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPartnerAccountProps struct {
	// `AWS::IoTWireless::PartnerAccount.AccountLinked`.
	AccountLinked interface{} `field:"optional" json:"accountLinked" yaml:"accountLinked"`
	// `AWS::IoTWireless::PartnerAccount.Fingerprint`.
	Fingerprint *string `field:"optional" json:"fingerprint" yaml:"fingerprint"`
	// The ID of the partner account to update.
	PartnerAccountId *string `field:"optional" json:"partnerAccountId" yaml:"partnerAccountId"`
	// `AWS::IoTWireless::PartnerAccount.PartnerType`.
	PartnerType *string `field:"optional" json:"partnerType" yaml:"partnerType"`
	// The Sidewalk account credentials.
	Sidewalk interface{} `field:"optional" json:"sidewalk" yaml:"sidewalk"`
	// `AWS::IoTWireless::PartnerAccount.SidewalkUpdate`.
	SidewalkUpdate interface{} `field:"optional" json:"sidewalkUpdate" yaml:"sidewalkUpdate"`
	// The tags are an array of key-value pairs to attach to the specified resource.
	//
	// Tags can have a minimum of 0 and a maximum of 50 items.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

