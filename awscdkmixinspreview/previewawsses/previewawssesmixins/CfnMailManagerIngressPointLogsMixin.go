package previewawssesmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsses/previewawssesmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource to provision an ingress endpoint for receiving email.
//
// An ingress endpoint serves as the entry point for incoming emails, allowing you to define how emails are received and processed within your AWS environment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var logsDelivery ILogsDelivery
//
//   cfnMailManagerIngressPointLogsMixin := awscdkmixinspreview.Mixins.NewCfnMailManagerIngressPointLogsMixin(jsii.String("logType"), logsDelivery)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-mailmanageringresspoint.html
//
type CfnMailManagerIngressPointLogsMixin interface {
	core.Mixin
	core.IMixin
	LogDelivery() previewawslogs.ILogsDelivery
	LogType() *string
	// Apply vended logs configuration to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct (has vendedLogs property).
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMailManagerIngressPointLogsMixin
type jsiiProxy_CfnMailManagerIngressPointLogsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnMailManagerIngressPointLogsMixin) LogDelivery() previewawslogs.ILogsDelivery {
	var returns previewawslogs.ILogsDelivery
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMailManagerIngressPointLogsMixin) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}


// Create a mixin to enable vended logs for `AWS::SES::MailManagerIngressPoint`.
func NewCfnMailManagerIngressPointLogsMixin(logType *string, logDelivery previewawslogs.ILogsDelivery) CfnMailManagerIngressPointLogsMixin {
	_init_.Initialize()

	if err := validateNewCfnMailManagerIngressPointLogsMixinParameters(logType, logDelivery); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMailManagerIngressPointLogsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointLogsMixin",
		[]interface{}{logType, logDelivery},
		&j,
	)

	return &j
}

// Create a mixin to enable vended logs for `AWS::SES::MailManagerIngressPoint`.
func NewCfnMailManagerIngressPointLogsMixin_Override(c CfnMailManagerIngressPointLogsMixin, logType *string, logDelivery previewawslogs.ILogsDelivery) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointLogsMixin",
		[]interface{}{logType, logDelivery},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnMailManagerIngressPointLogsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMailManagerIngressPointLogsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointLogsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMailManagerIngressPointLogsMixin_APPLICATION_LOGS() CfnMailManagerIngressPointApplicationLogs {
	_init_.Initialize()
	var returns CfnMailManagerIngressPointApplicationLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointLogsMixin",
		"APPLICATION_LOGS",
		&returns,
	)
	return returns
}

func CfnMailManagerIngressPointLogsMixin_TRAFFIC_POLICY_DEBUG_LOGS() CfnMailManagerIngressPointTrafficPolicyDebugLogs {
	_init_.Initialize()
	var returns CfnMailManagerIngressPointTrafficPolicyDebugLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointLogsMixin",
		"TRAFFIC_POLICY_DEBUG_LOGS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMailManagerIngressPointLogsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnMailManagerIngressPointLogsMixin) Supports(construct constructs.IConstruct) *bool {
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

