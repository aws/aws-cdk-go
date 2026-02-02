package previewawsshieldmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsshield/previewawsshieldmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Enables AWS Shield Advanced for a specific AWS resource.
//
// The resource can be an Amazon CloudFront distribution, Amazon Route 53 hosted zone, AWS Global Accelerator standard accelerator, Elastic IP Address, Application Load Balancer, or a Classic Load Balancer. You can protect Amazon EC2 instances and Network Load Balancers by association with protected Amazon EC2 Elastic IP addresses.
//
// *Configure a single `AWS::Shield::Protection`*
//
// Use this protection to protect a single resource at a time.
//
// To configure this Shield Advanced protection through CloudFormation , you must be subscribed to Shield Advanced . You can subscribe through the [Shield Advanced console](https://docs.aws.amazon.com/wafv2/shieldv2#/) and through the APIs. For more information, see [Subscribe to AWS Shield Advanced](https://docs.aws.amazon.com/waf/latest/developerguide/enable-ddos-prem.html) .
//
// See example templates for Shield Advanced in CloudFormation at [aws-samples/aws-shield-advanced-examples](https://docs.aws.amazon.com/https://github.com/aws-samples/aws-shield-advanced-examples) .
//
// *Configure Shield Advanced using AWS CloudFormation and AWS Firewall Manager*
//
// You might be able to use Firewall Manager with AWS CloudFormation to configure Shield Advanced across multiple accounts and protected resources. To do this, your accounts must be part of an organization in AWS Organizations . You can use Firewall Manager to configure Shield Advanced protections for any resource types except for Amazon Route 53 or AWS Global Accelerator .
//
// For an example of this, see the one-click configuration guidance published by the AWS technical community at [One-click deployment of Shield Advanced](https://docs.aws.amazon.com/https://youtu.be/LCA3FwMk_QE) .
//
// *Configure multiple protections through the Shield Advanced console*
//
// You can add protection to multiple resources at once through the [Shield Advanced console](https://docs.aws.amazon.com/wafv2/shieldv2#/) . For more information see [Getting Started with AWS Shield Advanced](https://docs.aws.amazon.com/waf/latest/developerguide/getting-started-ddos.html) and [Managing resource protections in AWS Shield Advanced](https://docs.aws.amazon.com/waf/latest/developerguide/ddos-manage-protected-resources.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var logsDelivery ILogsDelivery
//
//   cfnProtectionLogsMixin := awscdkmixinspreview.Mixins.NewCfnProtectionLogsMixin(jsii.String("logType"), logsDelivery)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-shield-protection.html
//
type CfnProtectionLogsMixin interface {
	core.Mixin
	core.IMixin
	LogDelivery() previewawslogs.ILogsDelivery
	LogType() *string
	// Apply vended logs configuration to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct (has vendedLogs property).
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnProtectionLogsMixin
type jsiiProxy_CfnProtectionLogsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnProtectionLogsMixin) LogDelivery() previewawslogs.ILogsDelivery {
	var returns previewawslogs.ILogsDelivery
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProtectionLogsMixin) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}


// Create a mixin to enable vended logs for `AWS::Shield::Protection`.
func NewCfnProtectionLogsMixin(logType *string, logDelivery previewawslogs.ILogsDelivery) CfnProtectionLogsMixin {
	_init_.Initialize()

	if err := validateNewCfnProtectionLogsMixinParameters(logType, logDelivery); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnProtectionLogsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnProtectionLogsMixin",
		[]interface{}{logType, logDelivery},
		&j,
	)

	return &j
}

// Create a mixin to enable vended logs for `AWS::Shield::Protection`.
func NewCfnProtectionLogsMixin_Override(c CfnProtectionLogsMixin, logType *string, logDelivery previewawslogs.ILogsDelivery) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnProtectionLogsMixin",
		[]interface{}{logType, logDelivery},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnProtectionLogsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnProtectionLogsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnProtectionLogsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnProtectionLogsMixin_FLOW_LOGS() CfnProtectionFlowLogs {
	_init_.Initialize()
	var returns CfnProtectionFlowLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnProtectionLogsMixin",
		"FLOW_LOGS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnProtectionLogsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnProtectionLogsMixin) Supports(construct constructs.IConstruct) *bool {
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

