package awssmsvoice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnPool`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPoolProps := &CfnPoolProps{
//   	MandatoryKeywords: &MandatoryKeywordsProperty{
//   		Help: &MandatoryKeywordProperty{
//   			Message: jsii.String("message"),
//   		},
//   		Stop: &MandatoryKeywordProperty{
//   			Message: jsii.String("message"),
//   		},
//   	},
//   	OriginationIdentities: []*string{
//   		jsii.String("originationIdentities"),
//   	},
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
//   	SharedRoutesEnabled: jsii.Boolean(false),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-pool.html
//
type CfnPoolProps struct {
	// Creates or updates the pool's `MandatoryKeyword` configuration.
	//
	// For more information, see [Keywords](https://docs.aws.amazon.com/sms-voice/latest/userguide/keywords.html) in the End User Messaging  User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-pool.html#cfn-smsvoice-pool-mandatorykeywords
	//
	MandatoryKeywords interface{} `field:"required" json:"mandatoryKeywords" yaml:"mandatoryKeywords"`
	// The list of origination identities to apply to the pool, either `PhoneNumberArn` or `SenderIdArn` .
	//
	// For more information, see [Registrations](https://docs.aws.amazon.com/sms-voice/latest/userguide/registrations.html) in the End User Messaging  User Guide.
	//
	// > If you are using a shared End User Messaging  resource then you must use the full Amazon Resource Name (ARN).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-pool.html#cfn-smsvoice-pool-originationidentities
	//
	OriginationIdentities *[]*string `field:"required" json:"originationIdentities" yaml:"originationIdentities"`
	// When set to true the pool can't be deleted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-pool.html#cfn-smsvoice-pool-deletionprotectionenabled
	//
	DeletionProtectionEnabled interface{} `field:"optional" json:"deletionProtectionEnabled" yaml:"deletionProtectionEnabled"`
	// Specifies any optional keywords to associate with the pool.
	//
	// For more information, see [Keywords](https://docs.aws.amazon.com/sms-voice/latest/userguide/keywords.html) in the End User Messaging  User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-pool.html#cfn-smsvoice-pool-optionalkeywords
	//
	OptionalKeywords interface{} `field:"optional" json:"optionalKeywords" yaml:"optionalKeywords"`
	// The name of the OptOutList associated with the pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-pool.html#cfn-smsvoice-pool-optoutlistname
	//
	OptOutListName *string `field:"optional" json:"optOutListName" yaml:"optOutListName"`
	// When set to false, an end recipient sends a message that begins with HELP or STOP to one of your dedicated numbers, End User Messaging  automatically replies with a customizable message and adds the end recipient to the OptOutList.
	//
	// When set to true you're responsible for responding to HELP and STOP requests. You're also responsible for tracking and honoring opt-out requests. For more information see [Self-managed opt-outs](https://docs.aws.amazon.com//pinpoint/latest/userguide/settings-sms-managing.html#settings-account-sms-self-managed-opt-out)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-pool.html#cfn-smsvoice-pool-selfmanagedoptoutsenabled
	//
	SelfManagedOptOutsEnabled interface{} `field:"optional" json:"selfManagedOptOutsEnabled" yaml:"selfManagedOptOutsEnabled"`
	// Allows you to enable shared routes on your pool.
	//
	// By default, this is set to `False` . If you set this value to `True` , your messages are sent using phone numbers or sender IDs (depending on the country) that are shared with other users. In some countries, such as the United States, senders aren't allowed to use shared routes and must use a dedicated phone number or short code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-pool.html#cfn-smsvoice-pool-sharedroutesenabled
	//
	SharedRoutesEnabled interface{} `field:"optional" json:"sharedRoutesEnabled" yaml:"sharedRoutesEnabled"`
	// An array of tags (key and value pairs) associated with the pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-pool.html#cfn-smsvoice-pool-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Describes the two-way SMS configuration for a phone number.
	//
	// For more information, see [Two-way SMS messaging](https://docs.aws.amazon.com/sms-voice/latest/userguide/two-way-sms.html) in the End User Messaging  User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-pool.html#cfn-smsvoice-pool-twoway
	//
	TwoWay interface{} `field:"optional" json:"twoWay" yaml:"twoWay"`
}

