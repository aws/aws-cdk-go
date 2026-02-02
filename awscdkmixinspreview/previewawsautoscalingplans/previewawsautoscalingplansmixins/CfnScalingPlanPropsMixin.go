package previewawsautoscalingplansmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsautoscalingplans/previewawsautoscalingplansmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::AutoScalingPlans::ScalingPlan` resource defines an AWS Auto Scaling scaling plan.
//
// A scaling plan is used to scale application resources to size them appropriately to ensure that enough resource is available in the application at peak times and to reduce allocated resource during periods of low utilization. The following resources can be added to a scaling plan:
//
// - Amazon EC2 Auto Scaling groups
// - Amazon EC2 Spot Fleet requests
// - Amazon ECS services
// - Amazon DynamoDB tables and global secondary indexes
// - Amazon Aurora Replicas
//
// For more information, see the [Scaling Plans User Guide](https://docs.aws.amazon.com/autoscaling/plans/userguide/what-is-a-scaling-plan.html)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnScalingPlanPropsMixin := awscdkmixinspreview.Mixins.NewCfnScalingPlanPropsMixin(&CfnScalingPlanMixinProps{
//   	ApplicationSource: &ApplicationSourceProperty{
//   		CloudFormationStackArn: jsii.String("cloudFormationStackArn"),
//   		TagFilters: []interface{}{
//   			&TagFilterProperty{
//   				Key: jsii.String("key"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   	},
//   	ScalingInstructions: []interface{}{
//   		&ScalingInstructionProperty{
//   			CustomizedLoadMetricSpecification: &CustomizedLoadMetricSpecificationProperty{
//   				Dimensions: []interface{}{
//   					&MetricDimensionProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				MetricName: jsii.String("metricName"),
//   				Namespace: jsii.String("namespace"),
//   				Statistic: jsii.String("statistic"),
//   				Unit: jsii.String("unit"),
//   			},
//   			DisableDynamicScaling: jsii.Boolean(false),
//   			MaxCapacity: jsii.Number(123),
//   			MinCapacity: jsii.Number(123),
//   			PredefinedLoadMetricSpecification: &PredefinedLoadMetricSpecificationProperty{
//   				PredefinedLoadMetricType: jsii.String("predefinedLoadMetricType"),
//   				ResourceLabel: jsii.String("resourceLabel"),
//   			},
//   			PredictiveScalingMaxCapacityBehavior: jsii.String("predictiveScalingMaxCapacityBehavior"),
//   			PredictiveScalingMaxCapacityBuffer: jsii.Number(123),
//   			PredictiveScalingMode: jsii.String("predictiveScalingMode"),
//   			ResourceId: jsii.String("resourceId"),
//   			ScalableDimension: jsii.String("scalableDimension"),
//   			ScalingPolicyUpdateBehavior: jsii.String("scalingPolicyUpdateBehavior"),
//   			ScheduledActionBufferTime: jsii.Number(123),
//   			ServiceNamespace: jsii.String("serviceNamespace"),
//   			TargetTrackingConfigurations: []interface{}{
//   				&TargetTrackingConfigurationProperty{
//   					CustomizedScalingMetricSpecification: &CustomizedScalingMetricSpecificationProperty{
//   						Dimensions: []interface{}{
//   							&MetricDimensionProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						MetricName: jsii.String("metricName"),
//   						Namespace: jsii.String("namespace"),
//   						Statistic: jsii.String("statistic"),
//   						Unit: jsii.String("unit"),
//   					},
//   					DisableScaleIn: jsii.Boolean(false),
//   					EstimatedInstanceWarmup: jsii.Number(123),
//   					PredefinedScalingMetricSpecification: &PredefinedScalingMetricSpecificationProperty{
//   						PredefinedScalingMetricType: jsii.String("predefinedScalingMetricType"),
//   						ResourceLabel: jsii.String("resourceLabel"),
//   					},
//   					ScaleInCooldown: jsii.Number(123),
//   					ScaleOutCooldown: jsii.Number(123),
//   					TargetValue: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscalingplans-scalingplan.html
//
type CfnScalingPlanPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnScalingPlanMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnScalingPlanPropsMixin
type jsiiProxy_CfnScalingPlanPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnScalingPlanPropsMixin) Props() *CfnScalingPlanMixinProps {
	var returns *CfnScalingPlanMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPlanPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AutoScalingPlans::ScalingPlan`.
func NewCfnScalingPlanPropsMixin(props *CfnScalingPlanMixinProps, options *mixins.CfnPropertyMixinOptions) CfnScalingPlanPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnScalingPlanPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnScalingPlanPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_autoscalingplans.mixins.CfnScalingPlanPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AutoScalingPlans::ScalingPlan`.
func NewCfnScalingPlanPropsMixin_Override(c CfnScalingPlanPropsMixin, props *CfnScalingPlanMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_autoscalingplans.mixins.CfnScalingPlanPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnScalingPlanPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnScalingPlanPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_autoscalingplans.mixins.CfnScalingPlanPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnScalingPlanPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_autoscalingplans.mixins.CfnScalingPlanPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnScalingPlanPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnScalingPlanPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

