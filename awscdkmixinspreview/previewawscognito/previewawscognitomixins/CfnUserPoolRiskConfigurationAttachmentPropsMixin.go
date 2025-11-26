package previewawscognitomixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscognito/previewawscognitomixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Cognito::UserPoolRiskConfigurationAttachment` resource sets the risk configuration that is used for Amazon Cognito advanced security features.
//
// You can specify risk configuration for a single client (with a specific `clientId` ) or for all clients (by setting the `clientId` to `ALL` ). If you specify `ALL` , the default configuration is used for every client that has had no risk configuration set previously. If you specify risk configuration for a particular client, it no longer falls back to the `ALL` configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnUserPoolRiskConfigurationAttachmentPropsMixin := awscdkmixinspreview.Mixins.NewCfnUserPoolRiskConfigurationAttachmentPropsMixin(&CfnUserPoolRiskConfigurationAttachmentMixinProps{
//   	AccountTakeoverRiskConfiguration: &AccountTakeoverRiskConfigurationTypeProperty{
//   		Actions: &AccountTakeoverActionsTypeProperty{
//   			HighAction: &AccountTakeoverActionTypeProperty{
//   				EventAction: jsii.String("eventAction"),
//   				Notify: jsii.Boolean(false),
//   			},
//   			LowAction: &AccountTakeoverActionTypeProperty{
//   				EventAction: jsii.String("eventAction"),
//   				Notify: jsii.Boolean(false),
//   			},
//   			MediumAction: &AccountTakeoverActionTypeProperty{
//   				EventAction: jsii.String("eventAction"),
//   				Notify: jsii.Boolean(false),
//   			},
//   		},
//   		NotifyConfiguration: &NotifyConfigurationTypeProperty{
//   			BlockEmail: &NotifyEmailTypeProperty{
//   				HtmlBody: jsii.String("htmlBody"),
//   				Subject: jsii.String("subject"),
//   				TextBody: jsii.String("textBody"),
//   			},
//   			From: jsii.String("from"),
//   			MfaEmail: &NotifyEmailTypeProperty{
//   				HtmlBody: jsii.String("htmlBody"),
//   				Subject: jsii.String("subject"),
//   				TextBody: jsii.String("textBody"),
//   			},
//   			NoActionEmail: &NotifyEmailTypeProperty{
//   				HtmlBody: jsii.String("htmlBody"),
//   				Subject: jsii.String("subject"),
//   				TextBody: jsii.String("textBody"),
//   			},
//   			ReplyTo: jsii.String("replyTo"),
//   			SourceArn: jsii.String("sourceArn"),
//   		},
//   	},
//   	ClientId: jsii.String("clientId"),
//   	CompromisedCredentialsRiskConfiguration: &CompromisedCredentialsRiskConfigurationTypeProperty{
//   		Actions: &CompromisedCredentialsActionsTypeProperty{
//   			EventAction: jsii.String("eventAction"),
//   		},
//   		EventFilter: []*string{
//   			jsii.String("eventFilter"),
//   		},
//   	},
//   	RiskExceptionConfiguration: &RiskExceptionConfigurationTypeProperty{
//   		BlockedIpRangeList: []*string{
//   			jsii.String("blockedIpRangeList"),
//   		},
//   		SkippedIpRangeList: []*string{
//   			jsii.String("skippedIpRangeList"),
//   		},
//   	},
//   	UserPoolId: jsii.String("userPoolId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolriskconfigurationattachment.html
//
type CfnUserPoolRiskConfigurationAttachmentPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnUserPoolRiskConfigurationAttachmentMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnUserPoolRiskConfigurationAttachmentPropsMixin
type jsiiProxy_CfnUserPoolRiskConfigurationAttachmentPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnUserPoolRiskConfigurationAttachmentPropsMixin) Props() *CfnUserPoolRiskConfigurationAttachmentMixinProps {
	var returns *CfnUserPoolRiskConfigurationAttachmentMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnUserPoolRiskConfigurationAttachmentPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Cognito::UserPoolRiskConfigurationAttachment`.
func NewCfnUserPoolRiskConfigurationAttachmentPropsMixin(props *CfnUserPoolRiskConfigurationAttachmentMixinProps, options *mixins.CfnPropertyMixinOptions) CfnUserPoolRiskConfigurationAttachmentPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnUserPoolRiskConfigurationAttachmentPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnUserPoolRiskConfigurationAttachmentPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolRiskConfigurationAttachmentPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Cognito::UserPoolRiskConfigurationAttachment`.
func NewCfnUserPoolRiskConfigurationAttachmentPropsMixin_Override(c CfnUserPoolRiskConfigurationAttachmentPropsMixin, props *CfnUserPoolRiskConfigurationAttachmentMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolRiskConfigurationAttachmentPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnUserPoolRiskConfigurationAttachmentPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnUserPoolRiskConfigurationAttachmentPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolRiskConfigurationAttachmentPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnUserPoolRiskConfigurationAttachmentPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolRiskConfigurationAttachmentPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnUserPoolRiskConfigurationAttachmentPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnUserPoolRiskConfigurationAttachmentPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

