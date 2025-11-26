package previewawssmsvoicemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssmsvoice/previewawssmsvoicemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new pool and associates the specified origination identity to the pool.
//
// A pool can include one or more phone numbers and SenderIds that are associated with your AWS account.
//
// The new pool inherits its configuration from the specified origination identity. This includes keywords, message type, opt-out list, two-way configuration, and self-managed opt-out configuration. Deletion protection isn't inherited from the origination identity and defaults to false.
//
// If the origination identity is a phone number and is already associated with another pool, an error is returned. A sender ID can be associated with multiple pools.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPoolPropsMixin := awscdkmixinspreview.Mixins.NewCfnPoolPropsMixin(&CfnPoolMixinProps{
//   	DeletionProtectionEnabled: jsii.Boolean(false),
//   	MandatoryKeywords: &MandatoryKeywordsProperty{
//   		Help: &MandatoryKeywordProperty{
//   			Message: jsii.String("message"),
//   		},
//   		Stop: &MandatoryKeywordProperty{
//   			Message: jsii.String("message"),
//   		},
//   	},
//   	OptionalKeywords: []interface{}{
//   		&OptionalKeywordProperty{
//   			Action: jsii.String("action"),
//   			Keyword: jsii.String("keyword"),
//   			Message: jsii.String("message"),
//   		},
//   	},
//   	OptOutListName: jsii.String("optOutListName"),
//   	OriginationIdentities: []*string{
//   		jsii.String("originationIdentities"),
//   	},
//   	SelfManagedOptOutsEnabled: jsii.Boolean(false),
//   	SharedRoutesEnabled: jsii.Boolean(false),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-smsvoice-pool.html
//
type CfnPoolPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnPoolMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPoolPropsMixin
type jsiiProxy_CfnPoolPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnPoolPropsMixin) Props() *CfnPoolMixinProps {
	var returns *CfnPoolMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPoolPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SMSVOICE::Pool`.
func NewCfnPoolPropsMixin(props *CfnPoolMixinProps, options *mixins.CfnPropertyMixinOptions) CfnPoolPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPoolPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPoolPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_smsvoice.mixins.CfnPoolPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SMSVOICE::Pool`.
func NewCfnPoolPropsMixin_Override(c CfnPoolPropsMixin, props *CfnPoolMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_smsvoice.mixins.CfnPoolPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPoolPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPoolPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_smsvoice.mixins.CfnPoolPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPoolPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_smsvoice.mixins.CfnPoolPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPoolPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnPoolPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

