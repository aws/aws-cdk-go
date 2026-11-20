package awscloudwatch

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awscloudwatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::CloudWatch::MetricDataQueryDefinition.
//
// A saved CloudWatch metric data query (PromQL, SQL, or Metric Math) with a path that is unique in the account and Region.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnMetricDataQueryDefinitionPropsMixin := awscdkcfnpropertymixins.Aws_cloudwatch.NewCfnMetricDataQueryDefinitionPropsMixin(&CfnMetricDataQueryDefinitionMixinProps{
//   	Expression: jsii.String("expression"),
//   	Path: jsii.String("path"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-metricdataquerydefinition.html
//
type CfnMetricDataQueryDefinitionPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnMetricDataQueryDefinitionMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMetricDataQueryDefinitionPropsMixin
type jsiiProxy_CfnMetricDataQueryDefinitionPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnMetricDataQueryDefinitionPropsMixin) Props() *CfnMetricDataQueryDefinitionMixinProps {
	var returns *CfnMetricDataQueryDefinitionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMetricDataQueryDefinitionPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CloudWatch::MetricDataQueryDefinition`.
func NewCfnMetricDataQueryDefinitionPropsMixin(props *CfnMetricDataQueryDefinitionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnMetricDataQueryDefinitionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMetricDataQueryDefinitionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMetricDataQueryDefinitionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnMetricDataQueryDefinitionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CloudWatch::MetricDataQueryDefinition`.
func NewCfnMetricDataQueryDefinitionPropsMixin_Override(c CfnMetricDataQueryDefinitionPropsMixin, props *CfnMetricDataQueryDefinitionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnMetricDataQueryDefinitionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnMetricDataQueryDefinitionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMetricDataQueryDefinitionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnMetricDataQueryDefinitionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMetricDataQueryDefinitionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_cloudwatch.CfnMetricDataQueryDefinitionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMetricDataQueryDefinitionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnMetricDataQueryDefinitionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

