package previewawsrummixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsrum/previewawsrummixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a CloudWatch RUM app monitor, which you can use to collect telemetry data from your application and send it to CloudWatch RUM.
//
// The data includes performance and reliability information such as page load time, client-side errors, and user behavior.
//
// After you create an app monitor, sign in to the CloudWatch RUM console to get the JavaScript code snippet to add to your web application. For more information, see [How do I find a code snippet that I've already generated?](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-find-code-snippet.html)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var logsDelivery ILogsDelivery
//
//   cfnAppMonitorLogsMixin := awscdkmixinspreview.Mixins.NewCfnAppMonitorLogsMixin(jsii.String("logType"), logsDelivery)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rum-appmonitor.html
//
type CfnAppMonitorLogsMixin interface {
	core.Mixin
	core.IMixin
	LogDelivery() previewawslogs.ILogsDelivery
	LogType() *string
	// Apply vended logs configuration to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct (has vendedLogs property).
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAppMonitorLogsMixin
type jsiiProxy_CfnAppMonitorLogsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAppMonitorLogsMixin) LogDelivery() previewawslogs.ILogsDelivery {
	var returns previewawslogs.ILogsDelivery
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAppMonitorLogsMixin) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}


// Create a mixin to enable vended logs for `AWS::RUM::AppMonitor`.
func NewCfnAppMonitorLogsMixin(logType *string, logDelivery previewawslogs.ILogsDelivery) CfnAppMonitorLogsMixin {
	_init_.Initialize()

	if err := validateNewCfnAppMonitorLogsMixinParameters(logType, logDelivery); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAppMonitorLogsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorLogsMixin",
		[]interface{}{logType, logDelivery},
		&j,
	)

	return &j
}

// Create a mixin to enable vended logs for `AWS::RUM::AppMonitor`.
func NewCfnAppMonitorLogsMixin_Override(c CfnAppMonitorLogsMixin, logType *string, logDelivery previewawslogs.ILogsDelivery) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorLogsMixin",
		[]interface{}{logType, logDelivery},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAppMonitorLogsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAppMonitorLogsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorLogsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAppMonitorLogsMixin_RUM_OTEL_LOGS() CfnAppMonitorRumOtelLogs {
	_init_.Initialize()
	var returns CfnAppMonitorRumOtelLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorLogsMixin",
		"RUM_OTEL_LOGS",
		&returns,
	)
	return returns
}

func CfnAppMonitorLogsMixin_RUM_OTEL_SPANS() CfnAppMonitorRumOtelSpans {
	_init_.Initialize()
	var returns CfnAppMonitorRumOtelSpans
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorLogsMixin",
		"RUM_OTEL_SPANS",
		&returns,
	)
	return returns
}

func CfnAppMonitorLogsMixin_RUM_TELEMETRY_LOGS() CfnAppMonitorRumTelemetryLogs {
	_init_.Initialize()
	var returns CfnAppMonitorRumTelemetryLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorLogsMixin",
		"RUM_TELEMETRY_LOGS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAppMonitorLogsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAppMonitorLogsMixin) Supports(construct constructs.IConstruct) *bool {
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

