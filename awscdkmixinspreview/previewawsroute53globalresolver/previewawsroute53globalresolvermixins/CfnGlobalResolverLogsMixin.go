package previewawsroute53globalresolvermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsroute53globalresolver/previewawsroute53globalresolvermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource schema for AWS::Route53GlobalResolver::GlobalResolver.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var logsDelivery ILogsDelivery
//
//   cfnGlobalResolverLogsMixin := awscdkmixinspreview.Mixins.NewCfnGlobalResolverLogsMixin(jsii.String("logType"), logsDelivery)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53globalresolver-globalresolver.html
//
type CfnGlobalResolverLogsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	LogDelivery() previewawslogs.ILogsDelivery
	LogType() *string
	// Apply vended logs configuration to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct (has vendedLogs property).
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnGlobalResolverLogsMixin
type jsiiProxy_CfnGlobalResolverLogsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnGlobalResolverLogsMixin) LogDelivery() previewawslogs.ILogsDelivery {
	var returns previewawslogs.ILogsDelivery
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalResolverLogsMixin) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}


// Create a mixin to enable vended logs for `AWS::Route53GlobalResolver::GlobalResolver`.
func NewCfnGlobalResolverLogsMixin(logType *string, logDelivery previewawslogs.ILogsDelivery) CfnGlobalResolverLogsMixin {
	_init_.Initialize()

	if err := validateNewCfnGlobalResolverLogsMixinParameters(logType, logDelivery); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnGlobalResolverLogsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53globalresolver.mixins.CfnGlobalResolverLogsMixin",
		[]interface{}{logType, logDelivery},
		&j,
	)

	return &j
}

// Create a mixin to enable vended logs for `AWS::Route53GlobalResolver::GlobalResolver`.
func NewCfnGlobalResolverLogsMixin_Override(c CfnGlobalResolverLogsMixin, logType *string, logDelivery previewawslogs.ILogsDelivery) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53globalresolver.mixins.CfnGlobalResolverLogsMixin",
		[]interface{}{logType, logDelivery},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnGlobalResolverLogsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnGlobalResolverLogsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_route53globalresolver.mixins.CfnGlobalResolverLogsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGlobalResolverLogsMixin_GLOBAL_RESOLVER_LOGS() CfnGlobalResolverGlobalResolverLogs {
	_init_.Initialize()
	var returns CfnGlobalResolverGlobalResolverLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_route53globalresolver.mixins.CfnGlobalResolverLogsMixin",
		"GLOBAL_RESOLVER_LOGS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnGlobalResolverLogsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnGlobalResolverLogsMixin) Supports(construct constructs.IConstruct) *bool {
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

