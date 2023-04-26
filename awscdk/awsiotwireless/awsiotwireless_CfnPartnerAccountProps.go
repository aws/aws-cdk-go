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
//   cfnPartnerAccountProps := &CfnPartnerAccountProps{
//   	AccountLinked: jsii.Boolean(false),
//   	PartnerAccountId: jsii.String("partnerAccountId"),
//   	PartnerType: jsii.String("partnerType"),
//   	Sidewalk: &SidewalkAccountInfoProperty{
//   		AppServerPrivateKey: jsii.String("appServerPrivateKey"),
//   	},
//   	SidewalkResponse: &SidewalkAccountInfoWithFingerprintProperty{
//   		AmazonId: jsii.String("amazonId"),
//   		Arn: jsii.String("arn"),
//   		Fingerprint: jsii.String("fingerprint"),
//   	},
//   	SidewalkUpdate: &SidewalkUpdateAccountProperty{
//   		AppServerPrivateKey: jsii.String("appServerPrivateKey"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPartnerAccountProps struct {
	// `AWS::IoTWireless::PartnerAccount.AccountLinked`.
	AccountLinked interface{} `field:"optional" json:"accountLinked" yaml:"accountLinked"`
	// The ID of the partner account to update.
	PartnerAccountId *string `field:"optional" json:"partnerAccountId" yaml:"partnerAccountId"`
	// `AWS::IoTWireless::PartnerAccount.PartnerType`.
	PartnerType *string `field:"optional" json:"partnerType" yaml:"partnerType"`
	// The Sidewalk account credentials.
	Sidewalk interface{} `field:"optional" json:"sidewalk" yaml:"sidewalk"`
	// `AWS::IoTWireless::PartnerAccount.SidewalkResponse`.
	SidewalkResponse interface{} `field:"optional" json:"sidewalkResponse" yaml:"sidewalkResponse"`
	// `AWS::IoTWireless::PartnerAccount.SidewalkUpdate`.
	SidewalkUpdate interface{} `field:"optional" json:"sidewalkUpdate" yaml:"sidewalkUpdate"`
	// The tags are an array of key-value pairs to attach to the specified resource.
	//
	// Tags can have a minimum of 0 and a maximum of 50 items.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

