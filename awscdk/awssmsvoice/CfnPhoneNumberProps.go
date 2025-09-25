package awssmsvoice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnPhoneNumber`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPhoneNumberProps := &CfnPhoneNumberProps{
//   	IsoCountryCode: jsii.String("isoCountryCode"),
//   	MandatoryKeywords: &MandatoryKeywordsProperty{
//   		Help: &MandatoryKeywordProperty{
//   			Message: jsii.String("message"),
//   		},
//   		Stop: &MandatoryKeywordProperty{
//   			Message: jsii.String("message"),
//   		},
//   	},
//   	NumberCapabilities: []*string{
//   		jsii.String("numberCapabilities"),
//   	},
//   	NumberType: jsii.String("numberType"),
//
//   	// the properties below are optional
//   	DeletionProtectionEnabled: jsii.Boolean(false),
//   	OptionalKeywords: []interface{}{
//   		&OptionalKeywordProperty{
//   			Action: jsii.String("action"),
//   			Keyword: jsii.String("keyword"),
//   			Message: jsii.String("message"),
//   		},
//   	},
//   	OptOutListName: jsii.String("optOutListName"),
//   	SelfManagedOptOutsEnabled: jsii.Boolean(false),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TwoWay: &TwoWayProperty{
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		ChannelArn: jsii.String("channelArn"),
//   		ChannelRole: jsii.String("channelRole"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-phonenumber.html
//
type CfnPhoneNumberProps struct {
	// The two-character code, in ISO 3166-1 alpha-2 format, for the country or region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-phonenumber.html#cfn-smsvoice-phonenumber-isocountrycode
	//
	IsoCountryCode *string `field:"required" json:"isoCountryCode" yaml:"isoCountryCode"`
	// Creates or updates a `MandatoryKeyword` configuration on an origination phone number For more information, see [Keywords](https://docs.aws.amazon.com/sms-voice/latest/userguide/keywords.html) in the End User Messaging  User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-phonenumber.html#cfn-smsvoice-phonenumber-mandatorykeywords
	//
	MandatoryKeywords interface{} `field:"required" json:"mandatoryKeywords" yaml:"mandatoryKeywords"`
	// Indicates if the phone number will be used for text messages, voice messages, or both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-phonenumber.html#cfn-smsvoice-phonenumber-numbercapabilities
	//
	NumberCapabilities *[]*string `field:"required" json:"numberCapabilities" yaml:"numberCapabilities"`
	// The type of phone number to request.
	//
	// > The `ShortCode` number type is not supported in AWS CloudFormation .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-phonenumber.html#cfn-smsvoice-phonenumber-numbertype
	//
	NumberType *string `field:"required" json:"numberType" yaml:"numberType"`
	// By default this is set to false.
	//
	// When set to true the phone number can't be deleted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-phonenumber.html#cfn-smsvoice-phonenumber-deletionprotectionenabled
	//
	DeletionProtectionEnabled interface{} `field:"optional" json:"deletionProtectionEnabled" yaml:"deletionProtectionEnabled"`
	// A keyword is a word that you can search for on a particular phone number or pool.
	//
	// It is also a specific word or phrase that an end user can send to your number to elicit a response, such as an informational message or a special offer. When your number receives a message that begins with a keyword, End User Messaging  responds with a customizable message. Optional keywords are differentiated from mandatory keywords. For more information, see [Keywords](https://docs.aws.amazon.com/sms-voice/latest/userguide/keywords.html) in the End User Messaging  User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-phonenumber.html#cfn-smsvoice-phonenumber-optionalkeywords
	//
	OptionalKeywords interface{} `field:"optional" json:"optionalKeywords" yaml:"optionalKeywords"`
	// The name of the OptOutList associated with the phone number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-phonenumber.html#cfn-smsvoice-phonenumber-optoutlistname
	//
	OptOutListName *string `field:"optional" json:"optOutListName" yaml:"optOutListName"`
	// When set to false and an end recipient sends a message that begins with HELP or STOP to one of your dedicated numbers, End User Messaging  automatically replies with a customizable message and adds the end recipient to the OptOutList.
	//
	// When set to true you're responsible for responding to HELP and STOP requests. You're also responsible for tracking and honoring opt-out request. For more information see [Self-managed opt-outs](https://docs.aws.amazon.com/sms-voice/latest/userguide/opt-out-list-self-managed.html)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-phonenumber.html#cfn-smsvoice-phonenumber-selfmanagedoptoutsenabled
	//
	SelfManagedOptOutsEnabled interface{} `field:"optional" json:"selfManagedOptOutsEnabled" yaml:"selfManagedOptOutsEnabled"`
	// An array of tags (key and value pairs) to associate with the requested phone number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-phonenumber.html#cfn-smsvoice-phonenumber-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Describes the two-way SMS configuration for a phone number.
	//
	// For more information, see [Two-way SMS messaging](https://docs.aws.amazon.com/sms-voice/latest/userguide/two-way-sms.html) in the End User Messaging  User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-phonenumber.html#cfn-smsvoice-phonenumber-twoway
	//
	TwoWay interface{} `field:"optional" json:"twoWay" yaml:"twoWay"`
}

