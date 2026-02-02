package previewawsivsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsivs/previewawsivsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::IVS::PlaybackRestrictionPolicy` resource specifies an  playback restriction policy.
//
// A playback restriction policy constrains playback by country and/or origin sites. For more information, see [Undesired Content and Viewers](https://docs.aws.amazon.com/ivs/latest/LowLatencyUserGuide/undesired-content.html) in the *Amazon IVS Low-Latency Streaming User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPlaybackRestrictionPolicyPropsMixin := awscdkmixinspreview.Mixins.NewCfnPlaybackRestrictionPolicyPropsMixin(&CfnPlaybackRestrictionPolicyMixinProps{
//   	AllowedCountries: []*string{
//   		jsii.String("allowedCountries"),
//   	},
//   	AllowedOrigins: []*string{
//   		jsii.String("allowedOrigins"),
//   	},
//   	EnableStrictOriginEnforcement: jsii.Boolean(false),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-playbackrestrictionpolicy.html
//
type CfnPlaybackRestrictionPolicyPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnPlaybackRestrictionPolicyMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPlaybackRestrictionPolicyPropsMixin
type jsiiProxy_CfnPlaybackRestrictionPolicyPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnPlaybackRestrictionPolicyPropsMixin) Props() *CfnPlaybackRestrictionPolicyMixinProps {
	var returns *CfnPlaybackRestrictionPolicyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackRestrictionPolicyPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IVS::PlaybackRestrictionPolicy`.
func NewCfnPlaybackRestrictionPolicyPropsMixin(props *CfnPlaybackRestrictionPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) CfnPlaybackRestrictionPolicyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPlaybackRestrictionPolicyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPlaybackRestrictionPolicyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ivs.mixins.CfnPlaybackRestrictionPolicyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IVS::PlaybackRestrictionPolicy`.
func NewCfnPlaybackRestrictionPolicyPropsMixin_Override(c CfnPlaybackRestrictionPolicyPropsMixin, props *CfnPlaybackRestrictionPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ivs.mixins.CfnPlaybackRestrictionPolicyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPlaybackRestrictionPolicyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPlaybackRestrictionPolicyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ivs.mixins.CfnPlaybackRestrictionPolicyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPlaybackRestrictionPolicyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ivs.mixins.CfnPlaybackRestrictionPolicyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPlaybackRestrictionPolicyPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnPlaybackRestrictionPolicyPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

