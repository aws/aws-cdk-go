package previewawsmediatailormixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsmediatailor/previewawsmediatailormixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Adds a new playback configuration to AWS Elemental MediaTailor .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var configurationAliases interface{}
//
//   cfnPlaybackConfigurationPropsMixin := awscdkmixinspreview.Mixins.NewCfnPlaybackConfigurationPropsMixin(&CfnPlaybackConfigurationMixinProps{
//   	AdConditioningConfiguration: &AdConditioningConfigurationProperty{
//   		StreamingMediaFileConditioning: jsii.String("streamingMediaFileConditioning"),
//   	},
//   	AdDecisionServerConfiguration: &AdDecisionServerConfigurationProperty{
//   		HttpRequest: &HttpRequestProperty{
//   			Body: jsii.String("body"),
//   			CompressRequest: jsii.String("compressRequest"),
//   			Headers: map[string]*string{
//   				"headersKey": jsii.String("headers"),
//   			},
//   			HttpMethod: jsii.String("httpMethod"),
//   		},
//   	},
//   	AdDecisionServerUrl: jsii.String("adDecisionServerUrl"),
//   	AvailSuppression: &AvailSuppressionProperty{
//   		FillPolicy: jsii.String("fillPolicy"),
//   		Mode: jsii.String("mode"),
//   		Value: jsii.String("value"),
//   	},
//   	Bumper: &BumperProperty{
//   		EndUrl: jsii.String("endUrl"),
//   		StartUrl: jsii.String("startUrl"),
//   	},
//   	CdnConfiguration: &CdnConfigurationProperty{
//   		AdSegmentUrlPrefix: jsii.String("adSegmentUrlPrefix"),
//   		ContentSegmentUrlPrefix: jsii.String("contentSegmentUrlPrefix"),
//   	},
//   	ConfigurationAliases: map[string]interface{}{
//   		"configurationAliasesKey": configurationAliases,
//   	},
//   	DashConfiguration: &DashConfigurationProperty{
//   		ManifestEndpointPrefix: jsii.String("manifestEndpointPrefix"),
//   		MpdLocation: jsii.String("mpdLocation"),
//   		OriginManifestType: jsii.String("originManifestType"),
//   	},
//   	HlsConfiguration: &HlsConfigurationProperty{
//   		ManifestEndpointPrefix: jsii.String("manifestEndpointPrefix"),
//   	},
//   	InsertionMode: jsii.String("insertionMode"),
//   	LivePreRollConfiguration: &LivePreRollConfigurationProperty{
//   		AdDecisionServerUrl: jsii.String("adDecisionServerUrl"),
//   		MaxDurationSeconds: jsii.Number(123),
//   	},
//   	LogConfiguration: &LogConfigurationProperty{
//   		AdsInteractionLog: &AdsInteractionLogProperty{
//   			ExcludeEventTypes: []*string{
//   				jsii.String("excludeEventTypes"),
//   			},
//   			PublishOptInEventTypes: []*string{
//   				jsii.String("publishOptInEventTypes"),
//   			},
//   		},
//   		EnabledLoggingStrategies: []*string{
//   			jsii.String("enabledLoggingStrategies"),
//   		},
//   		ManifestServiceInteractionLog: &ManifestServiceInteractionLogProperty{
//   			ExcludeEventTypes: []*string{
//   				jsii.String("excludeEventTypes"),
//   			},
//   		},
//   		PercentEnabled: jsii.Number(123),
//   	},
//   	ManifestProcessingRules: &ManifestProcessingRulesProperty{
//   		AdMarkerPassthrough: &AdMarkerPassthroughProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	PersonalizationThresholdSeconds: jsii.Number(123),
//   	SlateAdUrl: jsii.String("slateAdUrl"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TranscodeProfileName: jsii.String("transcodeProfileName"),
//   	VideoContentSourceUrl: jsii.String("videoContentSourceUrl"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-playbackconfiguration.html
//
type CfnPlaybackConfigurationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnPlaybackConfigurationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPlaybackConfigurationPropsMixin
type jsiiProxy_CfnPlaybackConfigurationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnPlaybackConfigurationPropsMixin) Props() *CfnPlaybackConfigurationMixinProps {
	var returns *CfnPlaybackConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfigurationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MediaTailor::PlaybackConfiguration`.
func NewCfnPlaybackConfigurationPropsMixin(props *CfnPlaybackConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnPlaybackConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPlaybackConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPlaybackConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MediaTailor::PlaybackConfiguration`.
func NewCfnPlaybackConfigurationPropsMixin_Override(c CfnPlaybackConfigurationPropsMixin, props *CfnPlaybackConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPlaybackConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPlaybackConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPlaybackConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPlaybackConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnPlaybackConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

