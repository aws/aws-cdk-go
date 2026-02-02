package previewawsoammixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsoam/previewawsoammixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a link between a source account and a sink that you have created in a monitoring account.
//
// Before you create a link, you must create a sink in the monitoring account. The sink must have a sink policy that permits the source account to link to it. You can grant permission to source accounts by granting permission to an entire organization, an organizational unit, or to individual accounts.
//
// For more information, see [CreateSink](https://docs.aws.amazon.com/OAM/latest/APIReference/API_CreateSink.html) and [PutSinkPolicy](https://docs.aws.amazon.com/OAM/latest/APIReference/API_PutSinkPolicy.html) .
//
// Each monitoring account can be linked to as many as 100,000 source accounts.
//
// Each source account can be linked to as many as five monitoring accounts.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLinkPropsMixin := awscdkmixinspreview.Mixins.NewCfnLinkPropsMixin(&CfnLinkMixinProps{
//   	LabelTemplate: jsii.String("labelTemplate"),
//   	LinkConfiguration: &LinkConfigurationProperty{
//   		LogGroupConfiguration: &LinkFilterProperty{
//   			Filter: jsii.String("filter"),
//   		},
//   		MetricConfiguration: &LinkFilterProperty{
//   			Filter: jsii.String("filter"),
//   		},
//   	},
//   	ResourceTypes: []*string{
//   		jsii.String("resourceTypes"),
//   	},
//   	SinkIdentifier: jsii.String("sinkIdentifier"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-oam-link.html
//
type CfnLinkPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnLinkMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLinkPropsMixin
type jsiiProxy_CfnLinkPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnLinkPropsMixin) Props() *CfnLinkMixinProps {
	var returns *CfnLinkMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLinkPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Oam::Link`.
func NewCfnLinkPropsMixin(props *CfnLinkMixinProps, options *mixins.CfnPropertyMixinOptions) CfnLinkPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLinkPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLinkPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_oam.mixins.CfnLinkPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Oam::Link`.
func NewCfnLinkPropsMixin_Override(c CfnLinkPropsMixin, props *CfnLinkMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_oam.mixins.CfnLinkPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnLinkPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLinkPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_oam.mixins.CfnLinkPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLinkPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_oam.mixins.CfnLinkPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLinkPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnLinkPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

