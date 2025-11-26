package previewawsmediatailormixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsmediatailor/previewawsmediatailormixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The configuration parameters for a channel.
//
// For information about MediaTailor channels, see [Working with channels](https://docs.aws.amazon.com/mediatailor/latest/ug/channel-assembly-channels.html) in the *MediaTailor User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnChannelPropsMixin := awscdkmixinspreview.Mixins.NewCfnChannelPropsMixin(&CfnChannelMixinProps{
//   	Audiences: []*string{
//   		jsii.String("audiences"),
//   	},
//   	ChannelName: jsii.String("channelName"),
//   	FillerSlate: &SlateSourceProperty{
//   		SourceLocationName: jsii.String("sourceLocationName"),
//   		VodSourceName: jsii.String("vodSourceName"),
//   	},
//   	LogConfiguration: &LogConfigurationForChannelProperty{
//   		LogTypes: []*string{
//   			jsii.String("logTypes"),
//   		},
//   	},
//   	Outputs: []interface{}{
//   		&RequestOutputItemProperty{
//   			DashPlaylistSettings: &DashPlaylistSettingsProperty{
//   				ManifestWindowSeconds: jsii.Number(123),
//   				MinBufferTimeSeconds: jsii.Number(123),
//   				MinUpdatePeriodSeconds: jsii.Number(123),
//   				SuggestedPresentationDelaySeconds: jsii.Number(123),
//   			},
//   			HlsPlaylistSettings: &HlsPlaylistSettingsProperty{
//   				AdMarkupType: []*string{
//   					jsii.String("adMarkupType"),
//   				},
//   				ManifestWindowSeconds: jsii.Number(123),
//   			},
//   			ManifestName: jsii.String("manifestName"),
//   			SourceGroup: jsii.String("sourceGroup"),
//   		},
//   	},
//   	PlaybackMode: jsii.String("playbackMode"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Tier: jsii.String("tier"),
//   	TimeShiftConfiguration: &TimeShiftConfigurationProperty{
//   		MaxTimeDelaySeconds: jsii.Number(123),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-channel.html
//
type CfnChannelPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnChannelMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnChannelPropsMixin
type jsiiProxy_CfnChannelPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnChannelPropsMixin) Props() *CfnChannelMixinProps {
	var returns *CfnChannelMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannelPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MediaTailor::Channel`.
func NewCfnChannelPropsMixin(props *CfnChannelMixinProps, options *mixins.CfnPropertyMixinOptions) CfnChannelPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnChannelPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnChannelPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnChannelPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MediaTailor::Channel`.
func NewCfnChannelPropsMixin_Override(c CfnChannelPropsMixin, props *CfnChannelMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnChannelPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnChannelPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnChannelPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnChannelPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnChannelPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnChannelPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnChannelPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnChannelPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

