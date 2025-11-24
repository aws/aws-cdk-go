package mixinsawsses

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsses/mixinsawsses/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies an identity for using within SES.
//
// An identity is an email address or domain that you use when you send email. Before you can use an identity to send email, you first have to verify it. By verifying an identity, you demonstrate that you're the owner of the identity, and that you've given Amazon SES API v2 permission to send email from the identity.
//
// When you verify an email address, SES sends an email to the address. Your email address is verified as soon as you follow the link in the verification email. When you verify a domain without specifying the `DkimSigningAttributes` properties, OR only the `NextSigningKeyLength` property of `DkimSigningAttributes` , this resource provides a set of CNAME token names and values ( *DkimDNSTokenName1* , *DkimDNSTokenValue1* , *DkimDNSTokenName2* , *DkimDNSTokenValue2* , *DkimDNSTokenName3* , *DkimDNSTokenValue3* ) as outputs. You can then add these to the DNS configuration for your domain. Your domain is verified when Amazon SES detects these records in the DNS configuration for your domain. This verification method is known as Easy DKIM.
//
// Alternatively, you can perform the verification process by providing your own public-private key pair. This verification method is known as Bring Your Own DKIM (BYODKIM). To use BYODKIM, your resource must include `DkimSigningAttributes` properties `DomainSigningSelector` and `DomainSigningPrivateKey` . When you specify this object, you provide a selector ( `DomainSigningSelector` ) (a component of the DNS record name that identifies the public key to use for DKIM authentication) and a private key ( `DomainSigningPrivateKey` ).
//
// Additionally, you can associate an existing configuration set with the email identity that you're verifying.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEmailIdentityPropsMixin := awscdkmixinspreview.Mixins.NewCfnEmailIdentityPropsMixin(&CfnEmailIdentityMixinProps{
//   	ConfigurationSetAttributes: &ConfigurationSetAttributesProperty{
//   		ConfigurationSetName: jsii.String("configurationSetName"),
//   	},
//   	DkimAttributes: &DkimAttributesProperty{
//   		SigningEnabled: jsii.Boolean(false),
//   	},
//   	DkimSigningAttributes: &DkimSigningAttributesProperty{
//   		DomainSigningPrivateKey: jsii.String("domainSigningPrivateKey"),
//   		DomainSigningSelector: jsii.String("domainSigningSelector"),
//   		NextSigningKeyLength: jsii.String("nextSigningKeyLength"),
//   	},
//   	EmailIdentity: jsii.String("emailIdentity"),
//   	FeedbackAttributes: &FeedbackAttributesProperty{
//   		EmailForwardingEnabled: jsii.Boolean(false),
//   	},
//   	MailFromAttributes: &MailFromAttributesProperty{
//   		BehaviorOnMxFailure: jsii.String("behaviorOnMxFailure"),
//   		MailFromDomain: jsii.String("mailFromDomain"),
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-emailidentity.html
//
type CfnEmailIdentityPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnEmailIdentityMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEmailIdentityPropsMixin
type jsiiProxy_CfnEmailIdentityPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnEmailIdentityPropsMixin) Props() *CfnEmailIdentityMixinProps {
	var returns *CfnEmailIdentityMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEmailIdentityPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SES::EmailIdentity`.
func NewCfnEmailIdentityPropsMixin(props *CfnEmailIdentityMixinProps, options *mixins.CfnPropertyMixinOptions) CfnEmailIdentityPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEmailIdentityPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEmailIdentityPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnEmailIdentityPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SES::EmailIdentity`.
func NewCfnEmailIdentityPropsMixin_Override(c CfnEmailIdentityPropsMixin, props *CfnEmailIdentityMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnEmailIdentityPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnEmailIdentityPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEmailIdentityPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnEmailIdentityPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEmailIdentityPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnEmailIdentityPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEmailIdentityPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnEmailIdentityPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

