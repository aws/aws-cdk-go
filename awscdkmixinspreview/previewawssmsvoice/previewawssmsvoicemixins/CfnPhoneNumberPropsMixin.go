package previewawssmsvoicemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssmsvoice/previewawssmsvoicemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Request an origination phone number for use in your account.
//
// For more information on phone number request see [Request a phone number](https://docs.aws.amazon.com/sms-voice/latest/userguide/phone-numbers-request.html) in the *End User Messaging  User Guide* .
//
// > Registering phone numbers is not supported by AWS CloudFormation . You can import phone numbers and sender IDs that are automatically provisioned at registration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPhoneNumberPropsMixin := awscdkmixinspreview.Mixins.NewCfnPhoneNumberPropsMixin(&CfnPhoneNumberMixinProps{
//   	DeletionProtectionEnabled: jsii.Boolean(false),
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
//   	OptionalKeywords: []interface{}{
//   		&OptionalKeywordProperty{
//   			Action: jsii.String("action"),
//   			Keyword: jsii.String("keyword"),
//   			Message: jsii.String("message"),
//   		},
//   	},
//   	OptOutListName: jsii.String("optOutListName"),
//   	SelfManagedOptOutsEnabled: jsii.Boolean(false),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TwoWay: &TwoWayProperty{
//   		ChannelArn: jsii.String("channelArn"),
//   		ChannelRole: jsii.String("channelRole"),
//   		Enabled: jsii.Boolean(false),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-phonenumber.html
//
type CfnPhoneNumberPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnPhoneNumberMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPhoneNumberPropsMixin
type jsiiProxy_CfnPhoneNumberPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnPhoneNumberPropsMixin) Props() *CfnPhoneNumberMixinProps {
	var returns *CfnPhoneNumberMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPhoneNumberPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SMSVOICE::PhoneNumber`.
func NewCfnPhoneNumberPropsMixin(props *CfnPhoneNumberMixinProps, options *mixins.CfnPropertyMixinOptions) CfnPhoneNumberPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPhoneNumberPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPhoneNumberPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_smsvoice.mixins.CfnPhoneNumberPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SMSVOICE::PhoneNumber`.
func NewCfnPhoneNumberPropsMixin_Override(c CfnPhoneNumberPropsMixin, props *CfnPhoneNumberMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_smsvoice.mixins.CfnPhoneNumberPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPhoneNumberPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPhoneNumberPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_smsvoice.mixins.CfnPhoneNumberPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPhoneNumberPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_smsvoice.mixins.CfnPhoneNumberPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPhoneNumberPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnPhoneNumberPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

