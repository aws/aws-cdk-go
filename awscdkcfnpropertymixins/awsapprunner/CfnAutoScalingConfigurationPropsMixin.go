package awsapprunner

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsapprunner/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specify an AWS App Runner Automatic Scaling configuration by using the `AWS::AppRunner::AutoScalingConfiguration` resource in an AWS CloudFormation template.
//
// The `AWS::AppRunner::AutoScalingConfiguration` resource is an AWS App Runner resource type that specifies an App Runner automatic scaling configuration.
//
// App Runner requires this resource to set non-default auto scaling settings for instances used to process the web requests. You can share an auto scaling configuration across multiple services.
//
// Create multiple revisions of a configuration by calling this action multiple times using the same `AutoScalingConfigurationName` . The call returns incremental `AutoScalingConfigurationRevision` values. When you create a service and configure an auto scaling configuration resource, the service uses the latest active revision of the auto scaling configuration by default. You can optionally configure the service to use a specific revision.
//
// Configure a higher `MinSize` to increase the spread of your App Runner service over more Availability Zones in the AWS Region . The tradeoff is a higher minimal cost.
//
// Configure a lower `MaxSize` to control your cost. The tradeoff is lower responsiveness during peak demand.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnAutoScalingConfigurationPropsMixin := awscdkcfnpropertymixins.Aws_apprunner.NewCfnAutoScalingConfigurationPropsMixin(&CfnAutoScalingConfigurationMixinProps{
//   	AutoScalingConfigurationName: jsii.String("autoScalingConfigurationName"),
//   	MaxConcurrency: jsii.Number(123),
//   	MaxSize: jsii.Number(123),
//   	MinSize: jsii.Number(123),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-autoscalingconfiguration.html
//
type CfnAutoScalingConfigurationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnAutoScalingConfigurationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAutoScalingConfigurationPropsMixin
type jsiiProxy_CfnAutoScalingConfigurationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnAutoScalingConfigurationPropsMixin) Props() *CfnAutoScalingConfigurationMixinProps {
	var returns *CfnAutoScalingConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingConfigurationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AppRunner::AutoScalingConfiguration`.
func NewCfnAutoScalingConfigurationPropsMixin(props *CfnAutoScalingConfigurationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnAutoScalingConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAutoScalingConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAutoScalingConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnAutoScalingConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AppRunner::AutoScalingConfiguration`.
func NewCfnAutoScalingConfigurationPropsMixin_Override(c CfnAutoScalingConfigurationPropsMixin, props *CfnAutoScalingConfigurationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnAutoScalingConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnAutoScalingConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAutoScalingConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnAutoScalingConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAutoScalingConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnAutoScalingConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAutoScalingConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAutoScalingConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

