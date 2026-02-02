package previewawsivsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsivs/previewawsivsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::IVS::PlaybackKeyPair` resource specifies an  playback key pair.
//
// uses a public playback key to validate playback tokens that have been signed with the corresponding private key. For more information, see [Setting Up Private Channels](https://docs.aws.amazon.com/ivs/latest/LowLatencyUserGuide/private-channels.html) in the *Amazon IVS Low-Latency Streaming User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPlaybackKeyPairPropsMixin := awscdkmixinspreview.Mixins.NewCfnPlaybackKeyPairPropsMixin(&CfnPlaybackKeyPairMixinProps{
//   	Name: jsii.String("name"),
//   	PublicKeyMaterial: jsii.String("publicKeyMaterial"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-playbackkeypair.html
//
type CfnPlaybackKeyPairPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnPlaybackKeyPairMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPlaybackKeyPairPropsMixin
type jsiiProxy_CfnPlaybackKeyPairPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnPlaybackKeyPairPropsMixin) Props() *CfnPlaybackKeyPairMixinProps {
	var returns *CfnPlaybackKeyPairMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackKeyPairPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IVS::PlaybackKeyPair`.
func NewCfnPlaybackKeyPairPropsMixin(props *CfnPlaybackKeyPairMixinProps, options *mixins.CfnPropertyMixinOptions) CfnPlaybackKeyPairPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPlaybackKeyPairPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPlaybackKeyPairPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ivs.mixins.CfnPlaybackKeyPairPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IVS::PlaybackKeyPair`.
func NewCfnPlaybackKeyPairPropsMixin_Override(c CfnPlaybackKeyPairPropsMixin, props *CfnPlaybackKeyPairMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ivs.mixins.CfnPlaybackKeyPairPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPlaybackKeyPairPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPlaybackKeyPairPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ivs.mixins.CfnPlaybackKeyPairPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPlaybackKeyPairPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ivs.mixins.CfnPlaybackKeyPairPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPlaybackKeyPairPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnPlaybackKeyPairPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

