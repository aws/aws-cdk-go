package previewawswisdommixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawswisdom/previewawswisdommixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an Amazon Q in Connect message template.
//
// The name of the message template has to be unique for each knowledge base. The channel subtype of the message template is immutable and cannot be modified after creation. After the message template is created, you can use the `$LATEST` qualifier to reference the created message template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMessageTemplatePropsMixin := awscdkmixinspreview.Mixins.NewCfnMessageTemplatePropsMixin(&CfnMessageTemplateMixinProps{
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
//   	KnowledgeBaseArn: jsii.String("knowledgeBaseArn"),
//   	Language: jsii.String("language"),
//   	MessageTemplateAttachments: []interface{}{
//   		&MessageTemplateAttachmentProperty{
//   			AttachmentId: jsii.String("attachmentId"),
//   			AttachmentName: jsii.String("attachmentName"),
//   			S3PresignedUrl: jsii.String("s3PresignedUrl"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-messagetemplate.html
//
type CfnMessageTemplatePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnMessageTemplateMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMessageTemplatePropsMixin
type jsiiProxy_CfnMessageTemplatePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnMessageTemplatePropsMixin) Props() *CfnMessageTemplateMixinProps {
	var returns *CfnMessageTemplateMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMessageTemplatePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Wisdom::MessageTemplate`.
func NewCfnMessageTemplatePropsMixin(props *CfnMessageTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) CfnMessageTemplatePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMessageTemplatePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMessageTemplatePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_wisdom.mixins.CfnMessageTemplatePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Wisdom::MessageTemplate`.
func NewCfnMessageTemplatePropsMixin_Override(c CfnMessageTemplatePropsMixin, props *CfnMessageTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_wisdom.mixins.CfnMessageTemplatePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnMessageTemplatePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMessageTemplatePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_wisdom.mixins.CfnMessageTemplatePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMessageTemplatePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_wisdom.mixins.CfnMessageTemplatePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMessageTemplatePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnMessageTemplatePropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

