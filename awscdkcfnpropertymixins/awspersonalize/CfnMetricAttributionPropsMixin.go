package awspersonalize

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awspersonalize/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a metric attribution for reporting on recommendation impact in Amazon Personalize.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnMetricAttributionPropsMixin := awscdkcfnpropertymixins.Aws_personalize.NewCfnMetricAttributionPropsMixin(&CfnMetricAttributionMixinProps{
//   	DatasetGroupArn: jsii.String("datasetGroupArn"),
//   	Metrics: []interface{}{
//   		&MetricAttributeProperty{
//   			EventType: jsii.String("eventType"),
//   			Expression: jsii.String("expression"),
//   			MetricName: jsii.String("metricName"),
//   		},
//   	},
//   	MetricsOutputConfig: &MetricsOutputConfigProperty{
//   		RoleArn: jsii.String("roleArn"),
//   		S3DataDestination: &S3DataDestinationProperty{
//   			KmsKeyArn: jsii.String("kmsKeyArn"),
//   			Path: jsii.String("path"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-personalize-metricattribution.html
//
type CfnMetricAttributionPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnMetricAttributionMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMetricAttributionPropsMixin
type jsiiProxy_CfnMetricAttributionPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnMetricAttributionPropsMixin) Props() *CfnMetricAttributionMixinProps {
	var returns *CfnMetricAttributionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMetricAttributionPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Personalize::MetricAttribution`.
func NewCfnMetricAttributionPropsMixin(props *CfnMetricAttributionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnMetricAttributionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMetricAttributionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMetricAttributionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_personalize.CfnMetricAttributionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Personalize::MetricAttribution`.
func NewCfnMetricAttributionPropsMixin_Override(c CfnMetricAttributionPropsMixin, props *CfnMetricAttributionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_personalize.CfnMetricAttributionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnMetricAttributionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMetricAttributionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_personalize.CfnMetricAttributionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMetricAttributionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_personalize.CfnMetricAttributionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMetricAttributionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnMetricAttributionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

