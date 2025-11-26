package previewawsmediatailormixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
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
//   var logsDelivery ILogsDelivery
//
//   cfnPlaybackConfigurationLogsMixin := awscdkmixinspreview.Mixins.NewCfnPlaybackConfigurationLogsMixin(jsii.String("logType"), logsDelivery)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-playbackconfiguration.html
//
type CfnPlaybackConfigurationLogsMixin interface {
	core.Mixin
	core.IMixin
	LogDelivery() previewawslogs.ILogsDelivery
	LogType() *string
	// Apply vended logs configuration to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct (has vendedLogs property).
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPlaybackConfigurationLogsMixin
type jsiiProxy_CfnPlaybackConfigurationLogsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnPlaybackConfigurationLogsMixin) LogDelivery() previewawslogs.ILogsDelivery {
	var returns previewawslogs.ILogsDelivery
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPlaybackConfigurationLogsMixin) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}


// Create a mixin to enable vended logs for `AWS::MediaTailor::PlaybackConfiguration`.
func NewCfnPlaybackConfigurationLogsMixin(logType *string, logDelivery previewawslogs.ILogsDelivery) CfnPlaybackConfigurationLogsMixin {
	_init_.Initialize()

	if err := validateNewCfnPlaybackConfigurationLogsMixinParameters(logType, logDelivery); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPlaybackConfigurationLogsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationLogsMixin",
		[]interface{}{logType, logDelivery},
		&j,
	)

	return &j
}

// Create a mixin to enable vended logs for `AWS::MediaTailor::PlaybackConfiguration`.
func NewCfnPlaybackConfigurationLogsMixin_Override(c CfnPlaybackConfigurationLogsMixin, logType *string, logDelivery previewawslogs.ILogsDelivery) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationLogsMixin",
		[]interface{}{logType, logDelivery},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPlaybackConfigurationLogsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPlaybackConfigurationLogsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationLogsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPlaybackConfigurationLogsMixin_AD_DECISION_SERVER_LOGS() CfnPlaybackConfigurationAdDecisionServerLogs {
	_init_.Initialize()
	var returns CfnPlaybackConfigurationAdDecisionServerLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationLogsMixin",
		"AD_DECISION_SERVER_LOGS",
		&returns,
	)
	return returns
}

func CfnPlaybackConfigurationLogsMixin_MANIFEST_SERVICE_LOGS() CfnPlaybackConfigurationManifestServiceLogs {
	_init_.Initialize()
	var returns CfnPlaybackConfigurationManifestServiceLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationLogsMixin",
		"MANIFEST_SERVICE_LOGS",
		&returns,
	)
	return returns
}

func CfnPlaybackConfigurationLogsMixin_TRANSCODE_LOGS() CfnPlaybackConfigurationTranscodeLogs {
	_init_.Initialize()
	var returns CfnPlaybackConfigurationTranscodeLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_mediatailor.mixins.CfnPlaybackConfigurationLogsMixin",
		"TRANSCODE_LOGS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPlaybackConfigurationLogsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnPlaybackConfigurationLogsMixin) Supports(construct constructs.IConstruct) *bool {
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

