package awsapprunner

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsapprunner/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specify an AWS App Runner observability configuration by using the `AWS::AppRunner::ObservabilityConfiguration` resource in an AWS CloudFormation template.
//
// The `AWS::AppRunner::ObservabilityConfiguration` resource is an AWS App Runner resource type that specifies an App Runner observability configuration.
//
// App Runner requires this resource when you specify App Runner services and you want to enable non-default observability features. You can share an observability configuration across multiple services.
//
// Create multiple revisions of a configuration by specifying this resource multiple times using the same `ObservabilityConfigurationName` . App Runner creates multiple resources with incremental `ObservabilityConfigurationRevision` values. When you specify a service and configure an observability configuration resource, the service uses the latest active revision of the observability configuration by default. You can optionally configure the service to use a specific revision.
//
// The observability configuration resource is designed to configure multiple features (currently one feature, tracing). This resource takes optional parameters that describe the configuration of these features (currently one parameter, `TraceConfiguration` ). If you don't specify a feature parameter, App Runner doesn't enable the feature.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnObservabilityConfigurationPropsMixin := awscdkcfnpropertymixins.Aws_apprunner.NewCfnObservabilityConfigurationPropsMixin(&CfnObservabilityConfigurationMixinProps{
//   	ObservabilityConfigurationName: jsii.String("observabilityConfigurationName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TraceConfiguration: &TraceConfigurationProperty{
//   		Vendor: jsii.String("vendor"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-observabilityconfiguration.html
//
type CfnObservabilityConfigurationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnObservabilityConfigurationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnObservabilityConfigurationPropsMixin
type jsiiProxy_CfnObservabilityConfigurationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnObservabilityConfigurationPropsMixin) Props() *CfnObservabilityConfigurationMixinProps {
	var returns *CfnObservabilityConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnObservabilityConfigurationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AppRunner::ObservabilityConfiguration`.
func NewCfnObservabilityConfigurationPropsMixin(props *CfnObservabilityConfigurationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnObservabilityConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnObservabilityConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnObservabilityConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnObservabilityConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AppRunner::ObservabilityConfiguration`.
func NewCfnObservabilityConfigurationPropsMixin_Override(c CfnObservabilityConfigurationPropsMixin, props *CfnObservabilityConfigurationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnObservabilityConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnObservabilityConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnObservabilityConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnObservabilityConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnObservabilityConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnObservabilityConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnObservabilityConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnObservabilityConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

