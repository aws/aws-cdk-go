package previewawseventsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsevents/previewawseventsmixins/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies an event bus within your account.
//
// This can be a custom event bus which you can use to receive events from your custom applications and services, or it can be a partner event bus which can be matched to a partner event source.
//
// > As an aid to help you jumpstart developing CloudFormation templates, the EventBridge console enables you to create templates from the existing event buses in your account. For more information, see [Generating CloudFormation templates from an EventBridge event bus](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-generate-event-bus-template.html) in the *Amazon EventBridge User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var logsDelivery ILogsDelivery
//
//   cfnEventBusLogsMixin := awscdkmixinspreview.Mixins.NewCfnEventBusLogsMixin(jsii.String("logType"), logsDelivery)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-eventbus.html
//
type CfnEventBusLogsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	LogDelivery() previewawslogs.ILogsDelivery
	LogType() *string
	// Apply vended logs configuration to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct (has vendedLogs property).
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEventBusLogsMixin
type jsiiProxy_CfnEventBusLogsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnEventBusLogsMixin) LogDelivery() previewawslogs.ILogsDelivery {
	var returns previewawslogs.ILogsDelivery
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEventBusLogsMixin) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}


// Create a mixin to enable vended logs for `AWS::Events::EventBus`.
func NewCfnEventBusLogsMixin(logType *string, logDelivery previewawslogs.ILogsDelivery) CfnEventBusLogsMixin {
	_init_.Initialize()

	if err := validateNewCfnEventBusLogsMixinParameters(logType, logDelivery); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEventBusLogsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusLogsMixin",
		[]interface{}{logType, logDelivery},
		&j,
	)

	return &j
}

// Create a mixin to enable vended logs for `AWS::Events::EventBus`.
func NewCfnEventBusLogsMixin_Override(c CfnEventBusLogsMixin, logType *string, logDelivery previewawslogs.ILogsDelivery) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusLogsMixin",
		[]interface{}{logType, logDelivery},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnEventBusLogsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEventBusLogsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusLogsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEventBusLogsMixin_ERROR_LOGS() CfnEventBusErrorLogs {
	_init_.Initialize()
	var returns CfnEventBusErrorLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusLogsMixin",
		"ERROR_LOGS",
		&returns,
	)
	return returns
}

func CfnEventBusLogsMixin_INFO_LOGS() CfnEventBusInfoLogs {
	_init_.Initialize()
	var returns CfnEventBusInfoLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusLogsMixin",
		"INFO_LOGS",
		&returns,
	)
	return returns
}

func CfnEventBusLogsMixin_TRACE_LOGS() CfnEventBusTraceLogs {
	_init_.Initialize()
	var returns CfnEventBusTraceLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_events.mixins.CfnEventBusLogsMixin",
		"TRACE_LOGS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEventBusLogsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnEventBusLogsMixin) Supports(construct constructs.IConstruct) *bool {
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

