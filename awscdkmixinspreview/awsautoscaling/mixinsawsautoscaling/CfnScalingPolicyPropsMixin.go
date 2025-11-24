package mixinsawsautoscaling

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsautoscaling/mixinsawsautoscaling/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::AutoScaling::ScalingPolicy` resource specifies an Amazon EC2 Auto Scaling scaling policy so that the Auto Scaling group can scale the number of instances available for your application.
//
// For more information about using scaling policies to scale your Auto Scaling group automatically, see [Dynamic scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scale-based-on-demand.html) and [Predictive scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-predictive-scaling.html) in the *Amazon EC2 Auto Scaling User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnScalingPolicyPropsMixin := awscdkmixinspreview.Mixins.NewCfnScalingPolicyPropsMixin(&CfnScalingPolicyMixinProps{
//   	AdjustmentType: jsii.String("adjustmentType"),
//   	AutoScalingGroupName: jsii.String("autoScalingGroupName"),
//   	Cooldown: jsii.String("cooldown"),
//   	EstimatedInstanceWarmup: jsii.Number(123),
//   	MetricAggregationType: jsii.String("metricAggregationType"),
//   	MinAdjustmentMagnitude: jsii.Number(123),
//   	PolicyType: jsii.String("policyType"),
//   	PredictiveScalingConfiguration: &PredictiveScalingConfigurationProperty{
//   		MaxCapacityBreachBehavior: jsii.String("maxCapacityBreachBehavior"),
//   		MaxCapacityBuffer: jsii.Number(123),
//   		MetricSpecifications: []interface{}{
//   			&PredictiveScalingMetricSpecificationProperty{
//   				CustomizedCapacityMetricSpecification: &PredictiveScalingCustomizedCapacityMetricProperty{
//   					MetricDataQueries: []interface{}{
//   						&MetricDataQueryProperty{
//   							Expression: jsii.String("expression"),
//   							Id: jsii.String("id"),
//   							Label: jsii.String("label"),
//   							MetricStat: &MetricStatProperty{
//   								Metric: &MetricProperty{
//   									Dimensions: []interface{}{
//   										&MetricDimensionProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   									MetricName: jsii.String("metricName"),
//   									Namespace: jsii.String("namespace"),
//   								},
//   								Stat: jsii.String("stat"),
//   								Unit: jsii.String("unit"),
//   							},
//   							ReturnData: jsii.Boolean(false),
//   						},
//   					},
//   				},
//   				CustomizedLoadMetricSpecification: &PredictiveScalingCustomizedLoadMetricProperty{
//   					MetricDataQueries: []interface{}{
//   						&MetricDataQueryProperty{
//   							Expression: jsii.String("expression"),
//   							Id: jsii.String("id"),
//   							Label: jsii.String("label"),
//   							MetricStat: &MetricStatProperty{
//   								Metric: &MetricProperty{
//   									Dimensions: []interface{}{
//   										&MetricDimensionProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   									MetricName: jsii.String("metricName"),
//   									Namespace: jsii.String("namespace"),
//   								},
//   								Stat: jsii.String("stat"),
//   								Unit: jsii.String("unit"),
//   							},
//   							ReturnData: jsii.Boolean(false),
//   						},
//   					},
//   				},
//   				CustomizedScalingMetricSpecification: &PredictiveScalingCustomizedScalingMetricProperty{
//   					MetricDataQueries: []interface{}{
//   						&MetricDataQueryProperty{
//   							Expression: jsii.String("expression"),
//   							Id: jsii.String("id"),
//   							Label: jsii.String("label"),
//   							MetricStat: &MetricStatProperty{
//   								Metric: &MetricProperty{
//   									Dimensions: []interface{}{
//   										&MetricDimensionProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   									MetricName: jsii.String("metricName"),
//   									Namespace: jsii.String("namespace"),
//   								},
//   								Stat: jsii.String("stat"),
//   								Unit: jsii.String("unit"),
//   							},
//   							ReturnData: jsii.Boolean(false),
//   						},
//   					},
//   				},
//   				PredefinedLoadMetricSpecification: &PredictiveScalingPredefinedLoadMetricProperty{
//   					PredefinedMetricType: jsii.String("predefinedMetricType"),
//   					ResourceLabel: jsii.String("resourceLabel"),
//   				},
//   				PredefinedMetricPairSpecification: &PredictiveScalingPredefinedMetricPairProperty{
//   					PredefinedMetricType: jsii.String("predefinedMetricType"),
//   					ResourceLabel: jsii.String("resourceLabel"),
//   				},
//   				PredefinedScalingMetricSpecification: &PredictiveScalingPredefinedScalingMetricProperty{
//   					PredefinedMetricType: jsii.String("predefinedMetricType"),
//   					ResourceLabel: jsii.String("resourceLabel"),
//   				},
//   				TargetValue: jsii.Number(123),
//   			},
//   		},
//   		Mode: jsii.String("mode"),
//   		SchedulingBufferTime: jsii.Number(123),
//   	},
//   	ScalingAdjustment: jsii.Number(123),
//   	StepAdjustments: []interface{}{
//   		&StepAdjustmentProperty{
//   			MetricIntervalLowerBound: jsii.Number(123),
//   			MetricIntervalUpperBound: jsii.Number(123),
//   			ScalingAdjustment: jsii.Number(123),
//   		},
//   	},
//   	TargetTrackingConfiguration: &TargetTrackingConfigurationProperty{
//   		CustomizedMetricSpecification: &CustomizedMetricSpecificationProperty{
//   			Dimensions: []interface{}{
//   				&MetricDimensionProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			MetricName: jsii.String("metricName"),
//   			Metrics: []interface{}{
//   				&TargetTrackingMetricDataQueryProperty{
//   					Expression: jsii.String("expression"),
//   					Id: jsii.String("id"),
//   					Label: jsii.String("label"),
//   					MetricStat: &TargetTrackingMetricStatProperty{
//   						Metric: &MetricProperty{
//   							Dimensions: []interface{}{
//   								&MetricDimensionProperty{
//   									Name: jsii.String("name"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							MetricName: jsii.String("metricName"),
//   							Namespace: jsii.String("namespace"),
//   						},
//   						Period: jsii.Number(123),
//   						Stat: jsii.String("stat"),
//   						Unit: jsii.String("unit"),
//   					},
//   					Period: jsii.Number(123),
//   					ReturnData: jsii.Boolean(false),
//   				},
//   			},
//   			Namespace: jsii.String("namespace"),
//   			Period: jsii.Number(123),
//   			Statistic: jsii.String("statistic"),
//   			Unit: jsii.String("unit"),
//   		},
//   		DisableScaleIn: jsii.Boolean(false),
//   		PredefinedMetricSpecification: &PredefinedMetricSpecificationProperty{
//   			PredefinedMetricType: jsii.String("predefinedMetricType"),
//   			ResourceLabel: jsii.String("resourceLabel"),
//   		},
//   		TargetValue: jsii.Number(123),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-scalingpolicy.html
//
type CfnScalingPolicyPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnScalingPolicyMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnScalingPolicyPropsMixin
type jsiiProxy_CfnScalingPolicyPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnScalingPolicyPropsMixin) Props() *CfnScalingPolicyMixinProps {
	var returns *CfnScalingPolicyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicyPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AutoScaling::ScalingPolicy`.
func NewCfnScalingPolicyPropsMixin(props *CfnScalingPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) CfnScalingPolicyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnScalingPolicyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnScalingPolicyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_autoscaling.mixins.CfnScalingPolicyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AutoScaling::ScalingPolicy`.
func NewCfnScalingPolicyPropsMixin_Override(c CfnScalingPolicyPropsMixin, props *CfnScalingPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_autoscaling.mixins.CfnScalingPolicyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnScalingPolicyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnScalingPolicyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_autoscaling.mixins.CfnScalingPolicyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnScalingPolicyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_autoscaling.mixins.CfnScalingPolicyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnScalingPolicyPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnScalingPolicyPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

