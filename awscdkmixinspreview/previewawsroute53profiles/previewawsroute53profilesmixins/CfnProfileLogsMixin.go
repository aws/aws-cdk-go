package previewawsroute53profilesmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsroute53profiles/previewawsroute53profilesmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A complex type that includes settings for a Route 53 Profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var logsDelivery ILogsDelivery
//
//   cfnProfileLogsMixin := awscdkmixinspreview.Mixins.NewCfnProfileLogsMixin(jsii.String("logType"), logsDelivery)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53profiles-profile.html
//
type CfnProfileLogsMixin interface {
	core.Mixin
	core.IMixin
	LogDelivery() previewawslogs.ILogsDelivery
	LogType() *string
	// Apply vended logs configuration to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct (has vendedLogs property).
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnProfileLogsMixin
type jsiiProxy_CfnProfileLogsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnProfileLogsMixin) LogDelivery() previewawslogs.ILogsDelivery {
	var returns previewawslogs.ILogsDelivery
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProfileLogsMixin) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}


// Create a mixin to enable vended logs for `AWS::Route53Profiles::Profile`.
func NewCfnProfileLogsMixin(logType *string, logDelivery previewawslogs.ILogsDelivery) CfnProfileLogsMixin {
	_init_.Initialize()

	if err := validateNewCfnProfileLogsMixinParameters(logType, logDelivery); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnProfileLogsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53profiles.mixins.CfnProfileLogsMixin",
		[]interface{}{logType, logDelivery},
		&j,
	)

	return &j
}

// Create a mixin to enable vended logs for `AWS::Route53Profiles::Profile`.
func NewCfnProfileLogsMixin_Override(c CfnProfileLogsMixin, logType *string, logDelivery previewawslogs.ILogsDelivery) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_route53profiles.mixins.CfnProfileLogsMixin",
		[]interface{}{logType, logDelivery},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnProfileLogsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnProfileLogsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_route53profiles.mixins.CfnProfileLogsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnProfileLogsMixin_ROUTE53_PROFILES_RESOLVER_QUERY_LOGS() CfnProfileRoute53ProfilesResolverQueryLogs {
	_init_.Initialize()
	var returns CfnProfileRoute53ProfilesResolverQueryLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_route53profiles.mixins.CfnProfileLogsMixin",
		"ROUTE53_PROFILES_RESOLVER_QUERY_LOGS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnProfileLogsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnProfileLogsMixin) Supports(construct constructs.IConstruct) *bool {
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

