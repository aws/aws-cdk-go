package awswisdom

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnMessageTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMessageTemplateProps := &CfnMessageTemplateProps{
//   	ChannelSubtype: jsii.String("channelSubtype"),
//   	Content: &ContentProperty{
//   		EmailMessageTemplateContent: &EmailMessageTemplateContentProperty{
//   			Body: &EmailMessageTemplateContentBodyProperty{
//   				Html: &MessageTemplateBodyContentProviderProperty{
//   					Content: jsii.String("content"),
//   				},
//   				PlainText: &MessageTemplateBodyContentProviderProperty{
//   					Content: jsii.String("content"),
//   				},
//   			},
//   			Headers: []interface{}{
//   				&EmailMessageTemplateHeaderProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Subject: jsii.String("subject"),
//   		},
//   		SmsMessageTemplateContent: &SmsMessageTemplateContentProperty{
//   			Body: &SmsMessageTemplateContentBodyProperty{
//   				PlainText: &MessageTemplateBodyContentProviderProperty{
//   					Content: jsii.String("content"),
//   				},
//   			},
//   		},
//   	},
//   	KnowledgeBaseArn: jsii.String("knowledgeBaseArn"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	DefaultAttributes: &MessageTemplateAttributesProperty{
//   		AgentAttributes: &AgentAttributesProperty{
//   			FirstName: jsii.String("firstName"),
//   			LastName: jsii.String("lastName"),
//   		},
//   		CustomAttributes: map[string]*string{
//   			"customAttributesKey": jsii.String("customAttributes"),
//   		},
//   		CustomerProfileAttributes: &CustomerProfileAttributesProperty{
//   			AccountNumber: jsii.String("accountNumber"),
//   			AdditionalInformation: jsii.String("additionalInformation"),
//   			Address1: jsii.String("address1"),
//   			Address2: jsii.String("address2"),
//   			Address3: jsii.String("address3"),
//   			Address4: jsii.String("address4"),
//   			BillingAddress1: jsii.String("billingAddress1"),
//   			BillingAddress2: jsii.String("billingAddress2"),
//   			BillingAddress3: jsii.String("billingAddress3"),
//   			BillingAddress4: jsii.String("billingAddress4"),
//   			BillingCity: jsii.String("billingCity"),
//   			BillingCountry: jsii.String("billingCountry"),
//   			BillingCounty: jsii.String("billingCounty"),
//   			BillingPostalCode: jsii.String("billingPostalCode"),
//   			BillingProvince: jsii.String("billingProvince"),
//   			BillingState: jsii.String("billingState"),
//   			BirthDate: jsii.String("birthDate"),
//   			BusinessEmailAddress: jsii.String("businessEmailAddress"),
//   			BusinessName: jsii.String("businessName"),
//   			BusinessPhoneNumber: jsii.String("businessPhoneNumber"),
//   			City: jsii.String("city"),
//   			Country: jsii.String("country"),
//   			County: jsii.String("county"),
//   			Custom: map[string]*string{
//   				"customKey": jsii.String("custom"),
//   			},
//   			EmailAddress: jsii.String("emailAddress"),
//   			FirstName: jsii.String("firstName"),
//   			Gender: jsii.String("gender"),
//   			HomePhoneNumber: jsii.String("homePhoneNumber"),
//   			LastName: jsii.String("lastName"),
//   			MailingAddress1: jsii.String("mailingAddress1"),
//   			MailingAddress2: jsii.String("mailingAddress2"),
//   			MailingAddress3: jsii.String("mailingAddress3"),
//   			MailingAddress4: jsii.String("mailingAddress4"),
//   			MailingCity: jsii.String("mailingCity"),
//   			MailingCountry: jsii.String("mailingCountry"),
//   			MailingCounty: jsii.String("mailingCounty"),
//   			MailingPostalCode: jsii.String("mailingPostalCode"),
//   			MailingProvince: jsii.String("mailingProvince"),
//   			MailingState: jsii.String("mailingState"),
//   			MiddleName: jsii.String("middleName"),
//   			MobilePhoneNumber: jsii.String("mobilePhoneNumber"),
//   			PartyType: jsii.String("partyType"),
//   			PhoneNumber: jsii.String("phoneNumber"),
//   			PostalCode: jsii.String("postalCode"),
//   			ProfileArn: jsii.String("profileArn"),
//   			ProfileId: jsii.String("profileId"),
//   			Province: jsii.String("province"),
//   			ShippingAddress1: jsii.String("shippingAddress1"),
//   			ShippingAddress2: jsii.String("shippingAddress2"),
//   			ShippingAddress3: jsii.String("shippingAddress3"),
//   			ShippingAddress4: jsii.String("shippingAddress4"),
//   			ShippingCity: jsii.String("shippingCity"),
//   			ShippingCountry: jsii.String("shippingCountry"),
//   			ShippingCounty: jsii.String("shippingCounty"),
//   			ShippingPostalCode: jsii.String("shippingPostalCode"),
//   			ShippingProvince: jsii.String("shippingProvince"),
//   			ShippingState: jsii.String("shippingState"),
//   			State: jsii.String("state"),
//   		},
//   		SystemAttributes: &SystemAttributesProperty{
//   			CustomerEndpoint: &SystemEndpointAttributesProperty{
//   				Address: jsii.String("address"),
//   			},
//   			Name: jsii.String("name"),
//   			SystemEndpoint: &SystemEndpointAttributesProperty{
//   				Address: jsii.String("address"),
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	GroupingConfiguration: &GroupingConfigurationProperty{
//   		Criteria: jsii.String("criteria"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	Language: jsii.String("language"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-messagetemplate.html
//
type CfnMessageTemplateProps struct {
	// The channel subtype this message template applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-messagetemplate.html#cfn-wisdom-messagetemplate-channelsubtype
	//
	ChannelSubtype *string `field:"required" json:"channelSubtype" yaml:"channelSubtype"`
	// The content of the message template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-messagetemplate.html#cfn-wisdom-messagetemplate-content
	//
	Content interface{} `field:"required" json:"content" yaml:"content"`
	// The Amazon Resource Name (ARN) of the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-messagetemplate.html#cfn-wisdom-messagetemplate-knowledgebasearn
	//
	KnowledgeBaseArn *string `field:"required" json:"knowledgeBaseArn" yaml:"knowledgeBaseArn"`
	// The name of the message template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-messagetemplate.html#cfn-wisdom-messagetemplate-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// An object that specifies the default values to use for variables in the message template.
	//
	// This object contains different categories of key-value pairs. Each key defines a variable or placeholder in the message template. The corresponding value defines the default value for that variable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-messagetemplate.html#cfn-wisdom-messagetemplate-defaultattributes
	//
	DefaultAttributes interface{} `field:"optional" json:"defaultAttributes" yaml:"defaultAttributes"`
	// The description of the message template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-messagetemplate.html#cfn-wisdom-messagetemplate-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The configuration information of the external data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-messagetemplate.html#cfn-wisdom-messagetemplate-groupingconfiguration
	//
	GroupingConfiguration interface{} `field:"optional" json:"groupingConfiguration" yaml:"groupingConfiguration"`
	// The language code value for the language in which the quick response is written.
	//
	// The supported language codes include `de_DE` , `en_US` , `es_ES` , `fr_FR` , `id_ID` , `it_IT` , `ja_JP` , `ko_KR` , `pt_BR` , `zh_CN` , `zh_TW`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-messagetemplate.html#cfn-wisdom-messagetemplate-language
	//
	Language *string `field:"optional" json:"language" yaml:"language"`
	// The tags used to organize, track, or control access for this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-messagetemplate.html#cfn-wisdom-messagetemplate-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

