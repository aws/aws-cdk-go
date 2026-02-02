package previewawsmediapackagev2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsmediapackagev2/previewawsmediapackagev2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies the configuration for a MediaPackage V2 channel group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var logsDelivery ILogsDelivery
//
//   cfnChannelGroupLogsMixin := awscdkmixinspreview.Mixins.NewCfnChannelGroupLogsMixin(jsii.String("logType"), logsDelivery)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-channelgroup.html
//
type CfnChannelGroupLogsMixin interface {
	core.Mixin
	core.IMixin
	LogDelivery() previewawslogs.ILogsDelivery
	LogType() *string
	// Apply vended logs configuration to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct (has vendedLogs property).
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnChannelGroupLogsMixin
type jsiiProxy_CfnChannelGroupLogsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnChannelGroupLogsMixin) LogDelivery() previewawslogs.ILogsDelivery {
	var returns previewawslogs.ILogsDelivery
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnChannelGroupLogsMixin) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}


// Create a mixin to enable vended logs for `AWS::MediaPackageV2::ChannelGroup`.
func NewCfnChannelGroupLogsMixin(logType *string, logDelivery previewawslogs.ILogsDelivery) CfnChannelGroupLogsMixin {
	_init_.Initialize()

	if err := validateNewCfnChannelGroupLogsMixinParameters(logType, logDelivery); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnChannelGroupLogsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupLogsMixin",
		[]interface{}{logType, logDelivery},
		&j,
	)

	return &j
}

// Create a mixin to enable vended logs for `AWS::MediaPackageV2::ChannelGroup`.
func NewCfnChannelGroupLogsMixin_Override(c CfnChannelGroupLogsMixin, logType *string, logDelivery previewawslogs.ILogsDelivery) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupLogsMixin",
		[]interface{}{logType, logDelivery},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnChannelGroupLogsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnChannelGroupLogsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupLogsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnChannelGroupLogsMixin_EGRESS_ACCESS_LOGS() CfnChannelGroupEgressAccessLogs {
	_init_.Initialize()
	var returns CfnChannelGroupEgressAccessLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupLogsMixin",
		"EGRESS_ACCESS_LOGS",
		&returns,
	)
	return returns
}

func CfnChannelGroupLogsMixin_INGRESS_ACCESS_LOGS() CfnChannelGroupIngressAccessLogs {
	_init_.Initialize()
	var returns CfnChannelGroupIngressAccessLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_mediapackagev2.mixins.CfnChannelGroupLogsMixin",
		"INGRESS_ACCESS_LOGS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnChannelGroupLogsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnChannelGroupLogsMixin) Supports(construct constructs.IConstruct) *bool {
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

