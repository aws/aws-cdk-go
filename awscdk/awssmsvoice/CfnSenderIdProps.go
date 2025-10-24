package awssmsvoice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSenderId`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSenderIdProps := &CfnSenderIdProps{
//   	IsoCountryCode: jsii.String("isoCountryCode"),
//   	SenderId: jsii.String("senderId"),
//
//   	// the properties below are optional
//   	DeletionProtectionEnabled: jsii.Boolean(false),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-senderid.html
//
type CfnSenderIdProps struct {
	// The two-character code, in ISO 3166-1 alpha-2 format, for the country or region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-senderid.html#cfn-smsvoice-senderid-isocountrycode
	//
	IsoCountryCode *string `field:"required" json:"isoCountryCode" yaml:"isoCountryCode"`
	// The sender ID string to request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-senderid.html#cfn-smsvoice-senderid-senderid
	//
	SenderId *string `field:"required" json:"senderId" yaml:"senderId"`
	// By default this is set to false.
	//
	// When set to true the sender ID can't be deleted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-senderid.html#cfn-smsvoice-senderid-deletionprotectionenabled
	//
	DeletionProtectionEnabled interface{} `field:"optional" json:"deletionProtectionEnabled" yaml:"deletionProtectionEnabled"`
	// An array of tags (key and value pairs) to associate with the sender ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-senderid.html#cfn-smsvoice-senderid-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

