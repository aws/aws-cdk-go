package awsiotwireless

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-partneraccount.html
//
type CfnPartnerAccountProps struct {
	// Whether the partner account is linked to the AWS account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-partneraccount.html#cfn-iotwireless-partneraccount-accountlinked
	//
	AccountLinked interface{} `field:"optional" json:"accountLinked" yaml:"accountLinked"`
	// The ID of the partner account to update.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-partneraccount.html#cfn-iotwireless-partneraccount-partneraccountid
	//
	PartnerAccountId *string `field:"optional" json:"partnerAccountId" yaml:"partnerAccountId"`
	// The partner type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-partneraccount.html#cfn-iotwireless-partneraccount-partnertype
	//
	PartnerType *string `field:"optional" json:"partnerType" yaml:"partnerType"`
	// The Sidewalk account credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-partneraccount.html#cfn-iotwireless-partneraccount-sidewalk
	//
	Sidewalk interface{} `field:"optional" json:"sidewalk" yaml:"sidewalk"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-partneraccount.html#cfn-iotwireless-partneraccount-sidewalkresponse
	//
	SidewalkResponse interface{} `field:"optional" json:"sidewalkResponse" yaml:"sidewalkResponse"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-partneraccount.html#cfn-iotwireless-partneraccount-sidewalkupdate
	//
	SidewalkUpdate interface{} `field:"optional" json:"sidewalkUpdate" yaml:"sidewalkUpdate"`
	// The tags are an array of key-value pairs to attach to the specified resource.
	//
	// Tags can have a minimum of 0 and a maximum of 50 items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-partneraccount.html#cfn-iotwireless-partneraccount-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

