package awsroute53

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsroute53/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::Route53::QueryLoggingConfig.
//
// Creates a configuration for DNS query logging, which publishes query logs for a public hosted zone to an Amazon CloudWatch Logs log group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnQueryLoggingConfigPropsMixin := awscdkcfnpropertymixins.Aws_route53.NewCfnQueryLoggingConfigPropsMixin(&CfnQueryLoggingConfigMixinProps{
//   	CloudWatchLogsLogGroupArn: jsii.String("cloudWatchLogsLogGroupArn"),
//   	HostedZoneId: jsii.String("hostedZoneId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-queryloggingconfig.html
//
type CfnQueryLoggingConfigPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnQueryLoggingConfigMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnQueryLoggingConfigPropsMixin
type jsiiProxy_CfnQueryLoggingConfigPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnQueryLoggingConfigPropsMixin) Props() *CfnQueryLoggingConfigMixinProps {
	var returns *CfnQueryLoggingConfigMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnQueryLoggingConfigPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Route53::QueryLoggingConfig`.
func NewCfnQueryLoggingConfigPropsMixin(props *CfnQueryLoggingConfigMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnQueryLoggingConfigPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnQueryLoggingConfigPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnQueryLoggingConfigPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnQueryLoggingConfigPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Route53::QueryLoggingConfig`.
func NewCfnQueryLoggingConfigPropsMixin_Override(c CfnQueryLoggingConfigPropsMixin, props *CfnQueryLoggingConfigMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnQueryLoggingConfigPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnQueryLoggingConfigPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnQueryLoggingConfigPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnQueryLoggingConfigPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnQueryLoggingConfigPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_route53.CfnQueryLoggingConfigPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnQueryLoggingConfigPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnQueryLoggingConfigPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

