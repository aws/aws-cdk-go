package awsapplicationautoscaling

import (
	"time"

	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsapplicationautoscaling/internal"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/constructs-go/constructs/v3"
)

// An adjustment.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import applicationautoscaling "github.com/aws/aws-cdk-go/awscdk/aws_applicationautoscaling"
//   adjustmentTier := &adjustmentTier{
//   	adjustment: jsii.Number(123),
//
//   	// the properties below are optional
//   	lowerBound: jsii.Number(123),
//   	upperBound: jsii.Number(123),
//   }
//
// Experimental.
type AdjustmentTier struct {
	// What number to adjust the capacity with.
	//
	// The number is interpeted as an added capacity, a new fixed capacity or an
	// added percentage depending on the AdjustmentType value of the
	// StepScalingPolicy.
	//
	// Can be positive or negative.
	// Experimental.
	Adjustment *float64 `json:"adjustment" yaml:"adjustment"`
	// Lower bound where this scaling tier applies.
	//
	// The scaling tier applies if the difference between the metric
	// value and its alarm threshold is higher than this value.
	// Experimental.
	LowerBound *float64 `json:"lowerBound" yaml:"lowerBound"`
	// Upper bound where this scaling tier applies.
	//
	// The scaling tier applies if the difference between the metric
	// value and its alarm threshold is lower than this value.
	// Experimental.
	UpperBound *float64 `json:"upperBound" yaml:"upperBound"`
}

// How adjustment numbers are interpreted.
//
// Example:
//   var capacity scalableAttribute
//   var cpuUtilization metric
//
//
//   capacity.scaleOnMetric(jsii.String("ScaleToCPU"), &basicStepScalingPolicyProps{
//   	metric: cpuUtilization,
//   	scalingSteps: []scalingInterval{
//   		&scalingInterval{
//   			upper: jsii.Number(10),
//   			change: -jsii.Number(1),
//   		},
//   		&scalingInterval{
//   			lower: jsii.Number(50),
//   			change: +jsii.Number(1),
//   		},
//   		&scalingInterval{
//   			lower: jsii.Number(70),
//   			change: +jsii.Number(3),
//   		},
//   	},
//
//   	// Change this to AdjustmentType.PercentChangeInCapacity to interpret the
//   	// 'change' numbers before as percentages instead of capacity counts.
//   	adjustmentType: appscaling.adjustmentType_CHANGE_IN_CAPACITY,
//   })
//
// Experimental.
type AdjustmentType string

const (
	// Add the adjustment number to the current capacity.
	//
	// A positive number increases capacity, a negative number decreases capacity.
	// Experimental.
	AdjustmentType_CHANGE_IN_CAPACITY AdjustmentType = "CHANGE_IN_CAPACITY"
	// Add this percentage of the current capacity to itself.
	//
	// The number must be between -100 and 100; a positive number increases
	// capacity and a negative number decreases it.
	// Experimental.
	AdjustmentType_PERCENT_CHANGE_IN_CAPACITY AdjustmentType = "PERCENT_CHANGE_IN_CAPACITY"
	// Make the capacity equal to the exact number given.
	// Experimental.
	AdjustmentType_EXACT_CAPACITY AdjustmentType = "EXACT_CAPACITY"
)

// Represent an attribute for which autoscaling can be configured.
//
// This class is basically a light wrapper around ScalableTarget, but with
// all methods protected instead of public so they can be selectively
// exposed and/or more specific versions of them can be exposed by derived
// classes for individual services support autoscaling.
//
// Typical use cases:
//
// - Hide away the PredefinedMetric enum for target tracking policies.
// - Don't expose all scaling methods (for example Dynamo tables don't support
//    Step Scaling, so the Dynamo subclass won't expose this method).
// Experimental.
type BaseScalableAttribute interface {
	awscdk.Construct
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Experimental.
	Props() *BaseScalableAttributeProps
	// Scale out or in based on a metric value.
	// Experimental.
	DoScaleOnMetric(id *string, props *BasicStepScalingPolicyProps)
	// Scale out or in based on time.
	// Experimental.
	DoScaleOnSchedule(id *string, props *ScalingSchedule)
	// Scale out or in in order to keep a metric around a target value.
	// Experimental.
	DoScaleToTrackMetric(id *string, props *BasicTargetTrackingScalingPolicyProps)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for BaseScalableAttribute
type jsiiProxy_BaseScalableAttribute struct {
	internal.Type__awscdkConstruct
}

func (j *jsiiProxy_BaseScalableAttribute) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseScalableAttribute) Props() *BaseScalableAttributeProps {
	var returns *BaseScalableAttributeProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


// Experimental.
func NewBaseScalableAttribute_Override(b BaseScalableAttribute, scope constructs.Construct, id *string, props *BaseScalableAttributeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_applicationautoscaling.BaseScalableAttribute",
		[]interface{}{scope, id, props},
		b,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func BaseScalableAttribute_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_applicationautoscaling.BaseScalableAttribute",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BaseScalableAttribute) DoScaleOnMetric(id *string, props *BasicStepScalingPolicyProps) {
	_jsii_.InvokeVoid(
		b,
		"doScaleOnMetric",
		[]interface{}{id, props},
	)
}

func (b *jsiiProxy_BaseScalableAttribute) DoScaleOnSchedule(id *string, props *ScalingSchedule) {
	_jsii_.InvokeVoid(
		b,
		"doScaleOnSchedule",
		[]interface{}{id, props},
	)
}

func (b *jsiiProxy_BaseScalableAttribute) DoScaleToTrackMetric(id *string, props *BasicTargetTrackingScalingPolicyProps) {
	_jsii_.InvokeVoid(
		b,
		"doScaleToTrackMetric",
		[]interface{}{id, props},
	)
}

func (b *jsiiProxy_BaseScalableAttribute) OnPrepare() {
	_jsii_.InvokeVoid(
		b,
		"onPrepare",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BaseScalableAttribute) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		b,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (b *jsiiProxy_BaseScalableAttribute) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BaseScalableAttribute) Prepare() {
	_jsii_.InvokeVoid(
		b,
		"prepare",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BaseScalableAttribute) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		b,
		"synthesize",
		[]interface{}{session},
	)
}

func (b *jsiiProxy_BaseScalableAttribute) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BaseScalableAttribute) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a ScalableTableAttribute.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import applicationautoscaling "github.com/aws/aws-cdk-go/awscdk/aws_applicationautoscaling"import awscdk "github.com/aws/aws-cdk-go/awscdk"import iam "github.com/aws/aws-cdk-go/awscdk/aws_iam"
//
//   var role role
//   baseScalableAttributeProps := &baseScalableAttributeProps{
//   	dimension: jsii.String("dimension"),
//   	maxCapacity: jsii.Number(123),
//   	resourceId: jsii.String("resourceId"),
//   	role: role,
//   	serviceNamespace: applicationautoscaling.serviceNamespace_ECS,
//
//   	// the properties below are optional
//   	minCapacity: jsii.Number(123),
//   }
//
// Experimental.
type BaseScalableAttributeProps struct {
	// Maximum capacity to scale to.
	// Experimental.
	MaxCapacity *float64 `json:"maxCapacity" yaml:"maxCapacity"`
	// Minimum capacity to scale to.
	// Experimental.
	MinCapacity *float64 `json:"minCapacity" yaml:"minCapacity"`
	// Scalable dimension of the attribute.
	// Experimental.
	Dimension *string `json:"dimension" yaml:"dimension"`
	// Resource ID of the attribute.
	// Experimental.
	ResourceId *string `json:"resourceId" yaml:"resourceId"`
	// Role to use for scaling.
	// Experimental.
	Role awsiam.IRole `json:"role" yaml:"role"`
	// Service namespace of the scalable attribute.
	// Experimental.
	ServiceNamespace ServiceNamespace `json:"serviceNamespace" yaml:"serviceNamespace"`
}

// Base interface for target tracking props.
//
// Contains the attributes that are common to target tracking policies,
// except the ones relating to the metric and to the scalable target.
//
// This interface is reused by more specific target tracking props objects
// in other services.
//
// Example:
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"import awscdk "github.com/aws/aws-cdk-go/awscdk"import applicationautoscaling "github.com/aws/aws-cdk-go/awscdk/aws_applicationautoscaling"
//
//   var duration duration
//   baseTargetTrackingProps := &baseTargetTrackingProps{
//   	disableScaleIn: jsii.Boolean(false),
//   	policyName: jsii.String("policyName"),
//   	scaleInCooldown: duration,
//   	scaleOutCooldown: duration,
//   }
//
// Experimental.
type BaseTargetTrackingProps struct {
	// Indicates whether scale in by the target tracking policy is disabled.
	//
	// If the value is true, scale in is disabled and the target tracking policy
	// won't remove capacity from the scalable resource. Otherwise, scale in is
	// enabled and the target tracking policy can remove capacity from the
	// scalable resource.
	// Experimental.
	DisableScaleIn *bool `json:"disableScaleIn" yaml:"disableScaleIn"`
	// A name for the scaling policy.
	// Experimental.
	PolicyName *string `json:"policyName" yaml:"policyName"`
	// Period after a scale in activity completes before another scale in activity can start.
	// Experimental.
	ScaleInCooldown awscdk.Duration `json:"scaleInCooldown" yaml:"scaleInCooldown"`
	// Period after a scale out activity completes before another scale out activity can start.
	// Experimental.
	ScaleOutCooldown awscdk.Duration `json:"scaleOutCooldown" yaml:"scaleOutCooldown"`
}

// Example:
//   var capacity scalableAttribute
//   var cpuUtilization metric
//
//
//   capacity.scaleOnMetric(jsii.String("ScaleToCPU"), &basicStepScalingPolicyProps{
//   	metric: cpuUtilization,
//   	scalingSteps: []scalingInterval{
//   		&scalingInterval{
//   			upper: jsii.Number(10),
//   			change: -jsii.Number(1),
//   		},
//   		&scalingInterval{
//   			lower: jsii.Number(50),
//   			change: +jsii.Number(1),
//   		},
//   		&scalingInterval{
//   			lower: jsii.Number(70),
//   			change: +jsii.Number(3),
//   		},
//   	},
//
//   	// Change this to AdjustmentType.PercentChangeInCapacity to interpret the
//   	// 'change' numbers before as percentages instead of capacity counts.
//   	adjustmentType: appscaling.adjustmentType_CHANGE_IN_CAPACITY,
//   })
//
// Experimental.
type BasicStepScalingPolicyProps struct {
	// Metric to scale on.
	// Experimental.
	Metric awscloudwatch.IMetric `json:"metric" yaml:"metric"`
	// The intervals for scaling.
	//
	// Maps a range of metric values to a particular scaling behavior.
	// Experimental.
	ScalingSteps *[]*ScalingInterval `json:"scalingSteps" yaml:"scalingSteps"`
	// How the adjustment numbers inside 'intervals' are interpreted.
	// Experimental.
	AdjustmentType AdjustmentType `json:"adjustmentType" yaml:"adjustmentType"`
	// Grace period after scaling activity.
	//
	// Subsequent scale outs during the cooldown period are squashed so that only
	// the biggest scale out happens.
	//
	// Subsequent scale ins during the cooldown period are ignored.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_StepScalingPolicyConfiguration.html
	//
	// Experimental.
	Cooldown awscdk.Duration `json:"cooldown" yaml:"cooldown"`
	// The number of data points out of the evaluation periods that must be breaching to trigger a scaling action.
	//
	// Creates an "M out of N" alarm, where this property is the M and the value set for
	// `evaluationPeriods` is the N value.
	//
	// Only has meaning if `evaluationPeriods != 1`.
	// Experimental.
	DatapointsToAlarm *float64 `json:"datapointsToAlarm" yaml:"datapointsToAlarm"`
	// How many evaluation periods of the metric to wait before triggering a scaling action.
	//
	// Raising this value can be used to smooth out the metric, at the expense
	// of slower response times.
	//
	// If `datapointsToAlarm` is not set, then all data points in the evaluation period
	// must meet the criteria to trigger a scaling action.
	// Experimental.
	EvaluationPeriods *float64 `json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// Aggregation to apply to all data points over the evaluation periods.
	//
	// Only has meaning if `evaluationPeriods != 1`.
	// Experimental.
	MetricAggregationType MetricAggregationType `json:"metricAggregationType" yaml:"metricAggregationType"`
	// Minimum absolute number to adjust capacity with as result of percentage scaling.
	//
	// Only when using AdjustmentType = PercentChangeInCapacity, this number controls
	// the minimum absolute effect size.
	// Experimental.
	MinAdjustmentMagnitude *float64 `json:"minAdjustmentMagnitude" yaml:"minAdjustmentMagnitude"`
}

// Properties for a Target Tracking policy that include the metric but exclude the target.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//   var code code
//
//
//   handler := lambda.NewFunction(this, jsii.String("MyFunction"), &functionProps{
//   	runtime: lambda.runtime_PYTHON_3_7(),
//   	handler: jsii.String("index.handler"),
//   	code: code,
//
//   	reservedConcurrentExecutions: jsii.Number(2),
//   })
//
//   fnVer := handler.addVersion(jsii.String("CDKLambdaVersion"), undefined, jsii.String("demo alias"), jsii.Number(10))
//
//   target := appscaling.NewScalableTarget(this, jsii.String("ScalableTarget"), &scalableTargetProps{
//   	serviceNamespace: appscaling.serviceNamespace_LAMBDA,
//   	maxCapacity: jsii.Number(100),
//   	minCapacity: jsii.Number(10),
//   	resourceId: fmt.Sprintf("function:%v:%v", handler.functionName, fnVer.version),
//   	scalableDimension: jsii.String("lambda:function:ProvisionedConcurrency"),
//   })
//
//   target.scaleToTrackMetric(jsii.String("PceTracking"), &basicTargetTrackingScalingPolicyProps{
//   	targetValue: jsii.Number(0.9),
//   	predefinedMetric: appscaling.predefinedMetric_LAMBDA_PROVISIONED_CONCURRENCY_UTILIZATION,
//   })
//
// Experimental.
type BasicTargetTrackingScalingPolicyProps struct {
	// Indicates whether scale in by the target tracking policy is disabled.
	//
	// If the value is true, scale in is disabled and the target tracking policy
	// won't remove capacity from the scalable resource. Otherwise, scale in is
	// enabled and the target tracking policy can remove capacity from the
	// scalable resource.
	// Experimental.
	DisableScaleIn *bool `json:"disableScaleIn" yaml:"disableScaleIn"`
	// A name for the scaling policy.
	// Experimental.
	PolicyName *string `json:"policyName" yaml:"policyName"`
	// Period after a scale in activity completes before another scale in activity can start.
	// Experimental.
	ScaleInCooldown awscdk.Duration `json:"scaleInCooldown" yaml:"scaleInCooldown"`
	// Period after a scale out activity completes before another scale out activity can start.
	// Experimental.
	ScaleOutCooldown awscdk.Duration `json:"scaleOutCooldown" yaml:"scaleOutCooldown"`
	// The target value for the metric.
	// Experimental.
	TargetValue *float64 `json:"targetValue" yaml:"targetValue"`
	// A custom metric for application autoscaling.
	//
	// The metric must track utilization. Scaling out will happen if the metric is higher than
	// the target value, scaling in will happen in the metric is lower than the target value.
	//
	// Exactly one of customMetric or predefinedMetric must be specified.
	// Experimental.
	CustomMetric awscloudwatch.IMetric `json:"customMetric" yaml:"customMetric"`
	// A predefined metric for application autoscaling.
	//
	// The metric must track utilization. Scaling out will happen if the metric is higher than
	// the target value, scaling in will happen in the metric is lower than the target value.
	//
	// Exactly one of customMetric or predefinedMetric must be specified.
	// Experimental.
	PredefinedMetric PredefinedMetric `json:"predefinedMetric" yaml:"predefinedMetric"`
	// Identify the resource associated with the metric type.
	//
	// Only used for predefined metric ALBRequestCountPerTarget.
	//
	// Example value: `app/<load-balancer-name>/<load-balancer-id>/targetgroup/<target-group-name>/<target-group-id>`.
	// Experimental.
	ResourceLabel *string `json:"resourceLabel" yaml:"resourceLabel"`
}

// A CloudFormation `AWS::ApplicationAutoScaling::ScalableTarget`.
//
// The `AWS::ApplicationAutoScaling::ScalableTarget` resource specifies a resource that Application Auto Scaling can scale, such as an AWS::DynamoDB::Table or AWS::ECS::Service resource.
//
// > If the resource that you want Application Auto Scaling to scale is not yet created in your account, add a dependency on the resource when registering it as a scalable target using the [DependsOn](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) attribute.
//
// For more information, see [RegisterScalableTarget](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html) in the *Application Auto Scaling API Reference* .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import applicationautoscaling "github.com/aws/aws-cdk-go/awscdk/aws_applicationautoscaling"
//   cfnScalableTarget := applicationautoscaling.NewCfnScalableTarget(this, jsii.String("MyCfnScalableTarget"), &cfnScalableTargetProps{
//   	maxCapacity: jsii.Number(123),
//   	minCapacity: jsii.Number(123),
//   	resourceId: jsii.String("resourceId"),
//   	roleArn: jsii.String("roleArn"),
//   	scalableDimension: jsii.String("scalableDimension"),
//   	serviceNamespace: jsii.String("serviceNamespace"),
//
//   	// the properties below are optional
//   	scheduledActions: []interface{}{
//   		&scheduledActionProperty{
//   			schedule: jsii.String("schedule"),
//   			scheduledActionName: jsii.String("scheduledActionName"),
//
//   			// the properties below are optional
//   			endTime: NewDate(),
//   			scalableTargetAction: &scalableTargetActionProperty{
//   				maxCapacity: jsii.Number(123),
//   				minCapacity: jsii.Number(123),
//   			},
//   			startTime: NewDate(),
//   			timezone: jsii.String("timezone"),
//   		},
//   	},
//   	suspendedState: &suspendedStateProperty{
//   		dynamicScalingInSuspended: jsii.Boolean(false),
//   		dynamicScalingOutSuspended: jsii.Boolean(false),
//   		scheduledScalingSuspended: jsii.Boolean(false),
//   	},
//   })
//
type CfnScalableTarget interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The maximum value that you plan to scale out to.
	//
	// When a scaling policy is in effect, Application Auto Scaling can scale out (expand) as needed to the maximum capacity limit in response to changing demand.
	MaxCapacity() *float64
	SetMaxCapacity(val *float64)
	// The minimum value that you plan to scale in to.
	//
	// When a scaling policy is in effect, Application Auto Scaling can scale in (contract) as needed to the minimum capacity limit in response to changing demand.
	MinCapacity() *float64
	SetMinCapacity(val *float64)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The identifier of the resource associated with the scalable target.
	//
	// This string consists of the resource type and unique identifier.
	//
	// - ECS service - The resource type is `service` and the unique identifier is the cluster name and service name. Example: `service/default/sample-webapp` .
	// - Spot Fleet - The resource type is `spot-fleet-request` and the unique identifier is the Spot Fleet request ID. Example: `spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE` .
	// - EMR cluster - The resource type is `instancegroup` and the unique identifier is the cluster ID and instance group ID. Example: `instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0` .
	// - AppStream 2.0 fleet - The resource type is `fleet` and the unique identifier is the fleet name. Example: `fleet/sample-fleet` .
	// - DynamoDB table - The resource type is `table` and the unique identifier is the table name. Example: `table/my-table` .
	// - DynamoDB global secondary index - The resource type is `index` and the unique identifier is the index name. Example: `table/my-table/index/my-table-index` .
	// - Aurora DB cluster - The resource type is `cluster` and the unique identifier is the cluster name. Example: `cluster:my-db-cluster` .
	// - SageMaker endpoint variant - The resource type is `variant` and the unique identifier is the resource ID. Example: `endpoint/my-end-point/variant/KMeansClustering` .
	// - Custom resources are not supported with a resource type. This parameter must specify the `OutputValue` from the CloudFormation template stack used to access the resources. The unique identifier is defined by the service provider. More information is available in our [GitHub repository](https://docs.aws.amazon.com/https://github.com/aws/aws-auto-scaling-custom-resource) .
	// - Amazon Comprehend document classification endpoint - The resource type and unique identifier are specified using the endpoint ARN. Example: `arn:aws:comprehend:us-west-2:123456789012:document-classifier-endpoint/EXAMPLE` .
	// - Amazon Comprehend entity recognizer endpoint - The resource type and unique identifier are specified using the endpoint ARN. Example: `arn:aws:comprehend:us-west-2:123456789012:entity-recognizer-endpoint/EXAMPLE` .
	// - Lambda provisioned concurrency - The resource type is `function` and the unique identifier is the function name with a function version or alias name suffix that is not `$LATEST` . Example: `function:my-function:prod` or `function:my-function:1` .
	// - Amazon Keyspaces table - The resource type is `table` and the unique identifier is the table name. Example: `keyspace/mykeyspace/table/mytable` .
	// - Amazon MSK cluster - The resource type and unique identifier are specified using the cluster ARN. Example: `arn:aws:kafka:us-east-1:123456789012:cluster/demo-cluster-1/6357e0b2-0e6a-4b86-a0b4-70df934c2e31-5` .
	// - Amazon ElastiCache replication group - The resource type is `replication-group` and the unique identifier is the replication group name. Example: `replication-group/mycluster` .
	// - Neptune cluster - The resource type is `cluster` and the unique identifier is the cluster name. Example: `cluster:mycluster` .
	ResourceId() *string
	SetResourceId(val *string)
	// Specify the Amazon Resource Name (ARN) of an Identity and Access Management (IAM) role that allows Application Auto Scaling to modify the scalable target on your behalf.
	//
	// This can be either an IAM service role that Application Auto Scaling can assume to make calls to other AWS resources on your behalf, or a service-linked role for the specified service. For more information, see [How Application Auto Scaling works with IAM](https://docs.aws.amazon.com/autoscaling/application/userguide/security_iam_service-with-iam.html) in the *Application Auto Scaling User Guide* .
	//
	// To automatically create a service-linked role (recommended), specify the full ARN of the service-linked role in your stack template. To find the exact ARN of the service-linked role for your AWS or custom resource, see the [Service-linked roles](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-service-linked-roles.html) topic in the *Application Auto Scaling User Guide* . Look for the ARN in the table at the bottom of the page.
	RoleArn() *string
	SetRoleArn(val *string)
	// The scalable dimension associated with the scalable target.
	//
	// This string consists of the service namespace, resource type, and scaling property.
	//
	// - `ecs:service:DesiredCount` - The desired task count of an ECS service.
	// - `elasticmapreduce:instancegroup:InstanceCount` - The instance count of an EMR Instance Group.
	// - `ec2:spot-fleet-request:TargetCapacity` - The target capacity of a Spot Fleet.
	// - `appstream:fleet:DesiredCapacity` - The desired capacity of an AppStream 2.0 fleet.
	// - `dynamodb:table:ReadCapacityUnits` - The provisioned read capacity for a DynamoDB table.
	// - `dynamodb:table:WriteCapacityUnits` - The provisioned write capacity for a DynamoDB table.
	// - `dynamodb:index:ReadCapacityUnits` - The provisioned read capacity for a DynamoDB global secondary index.
	// - `dynamodb:index:WriteCapacityUnits` - The provisioned write capacity for a DynamoDB global secondary index.
	// - `rds:cluster:ReadReplicaCount` - The count of Aurora Replicas in an Aurora DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible edition.
	// - `sagemaker:variant:DesiredInstanceCount` - The number of EC2 instances for a SageMaker model endpoint variant.
	// - `custom-resource:ResourceType:Property` - The scalable dimension for a custom resource provided by your own application or service.
	// - `comprehend:document-classifier-endpoint:DesiredInferenceUnits` - The number of inference units for an Amazon Comprehend document classification endpoint.
	// - `comprehend:entity-recognizer-endpoint:DesiredInferenceUnits` - The number of inference units for an Amazon Comprehend entity recognizer endpoint.
	// - `lambda:function:ProvisionedConcurrency` - The provisioned concurrency for a Lambda function.
	// - `cassandra:table:ReadCapacityUnits` - The provisioned read capacity for an Amazon Keyspaces table.
	// - `cassandra:table:WriteCapacityUnits` - The provisioned write capacity for an Amazon Keyspaces table.
	// - `kafka:broker-storage:VolumeSize` - The provisioned volume size (in GiB) for brokers in an Amazon MSK cluster.
	// - `elasticache:replication-group:NodeGroups` - The number of node groups for an Amazon ElastiCache replication group.
	// - `elasticache:replication-group:Replicas` - The number of replicas per node group for an Amazon ElastiCache replication group.
	// - `neptune:cluster:ReadReplicaCount` - The count of read replicas in an Amazon Neptune DB cluster.
	ScalableDimension() *string
	SetScalableDimension(val *string)
	// The scheduled actions for the scalable target. Duplicates aren't allowed.
	//
	// For more information about using scheduled scaling, see [Scheduled scaling](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-scheduled-scaling.html) in the *Application Auto Scaling User Guide* .
	ScheduledActions() interface{}
	SetScheduledActions(val interface{})
	// The namespace of the AWS service that provides the resource, or a `custom-resource` .
	ServiceNamespace() *string
	SetServiceNamespace(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// An embedded object that contains attributes and attribute values that are used to suspend and resume automatic scaling.
	//
	// Setting the value of an attribute to `true` suspends the specified scaling activities. Setting it to `false` (default) resumes the specified scaling activities.
	//
	// *Suspension Outcomes*
	//
	// - For `DynamicScalingInSuspended` , while a suspension is in effect, all scale-in activities that are triggered by a scaling policy are suspended.
	// - For `DynamicScalingOutSuspended` , while a suspension is in effect, all scale-out activities that are triggered by a scaling policy are suspended.
	// - For `ScheduledScalingSuspended` , while a suspension is in effect, all scaling activities that involve scheduled actions are suspended.
	//
	// For more information, see [Suspending and resuming scaling](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-suspend-resume-scaling.html) in the *Application Auto Scaling User Guide* .
	SuspendedState() interface{}
	SetSuspendedState(val interface{})
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnScalableTarget
type jsiiProxy_CfnScalableTarget struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnScalableTarget) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalableTarget) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalableTarget) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalableTarget) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalableTarget) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalableTarget) MaxCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalableTarget) MinCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalableTarget) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalableTarget) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalableTarget) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalableTarget) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalableTarget) ScalableDimension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalableDimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalableTarget) ScheduledActions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduledActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalableTarget) ServiceNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalableTarget) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalableTarget) SuspendedState() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suspendedState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalableTarget) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApplicationAutoScaling::ScalableTarget`.
func NewCfnScalableTarget(scope awscdk.Construct, id *string, props *CfnScalableTargetProps) CfnScalableTarget {
	_init_.Initialize()

	j := jsiiProxy_CfnScalableTarget{}

	_jsii_.Create(
		"monocdk.aws_applicationautoscaling.CfnScalableTarget",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApplicationAutoScaling::ScalableTarget`.
func NewCfnScalableTarget_Override(c CfnScalableTarget, scope awscdk.Construct, id *string, props *CfnScalableTargetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_applicationautoscaling.CfnScalableTarget",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnScalableTarget) SetMaxCapacity(val *float64) {
	_jsii_.Set(
		j,
		"maxCapacity",
		val,
	)
}

func (j *jsiiProxy_CfnScalableTarget) SetMinCapacity(val *float64) {
	_jsii_.Set(
		j,
		"minCapacity",
		val,
	)
}

func (j *jsiiProxy_CfnScalableTarget) SetResourceId(val *string) {
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_CfnScalableTarget) SetRoleArn(val *string) {
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_CfnScalableTarget) SetScalableDimension(val *string) {
	_jsii_.Set(
		j,
		"scalableDimension",
		val,
	)
}

func (j *jsiiProxy_CfnScalableTarget) SetScheduledActions(val interface{}) {
	_jsii_.Set(
		j,
		"scheduledActions",
		val,
	)
}

func (j *jsiiProxy_CfnScalableTarget) SetServiceNamespace(val *string) {
	_jsii_.Set(
		j,
		"serviceNamespace",
		val,
	)
}

func (j *jsiiProxy_CfnScalableTarget) SetSuspendedState(val interface{}) {
	_jsii_.Set(
		j,
		"suspendedState",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnScalableTarget_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_applicationautoscaling.CfnScalableTarget",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnScalableTarget_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_applicationautoscaling.CfnScalableTarget",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnScalableTarget_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_applicationautoscaling.CfnScalableTarget",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnScalableTarget_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_applicationautoscaling.CfnScalableTarget",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnScalableTarget) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnScalableTarget) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnScalableTarget) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnScalableTarget) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnScalableTarget) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnScalableTarget) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnScalableTarget) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnScalableTarget) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScalableTarget) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScalableTarget) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnScalableTarget) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnScalableTarget) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnScalableTarget) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScalableTarget) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnScalableTarget) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnScalableTarget) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScalableTarget) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScalableTarget) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnScalableTarget) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScalableTarget) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScalableTarget) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// `ScalableTargetAction` specifies the minimum and maximum capacity for the `ScalableTargetAction` property of the [AWS::ApplicationAutoScaling::ScalableTarget ScheduledAction](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalabletarget-scheduledaction.html) property type.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import applicationautoscaling "github.com/aws/aws-cdk-go/awscdk/aws_applicationautoscaling"
//   scalableTargetActionProperty := &scalableTargetActionProperty{
//   	maxCapacity: jsii.Number(123),
//   	minCapacity: jsii.Number(123),
//   }
//
type CfnScalableTarget_ScalableTargetActionProperty struct {
	// The maximum capacity.
	MaxCapacity *float64 `json:"maxCapacity" yaml:"maxCapacity"`
	// The minimum capacity.
	MinCapacity *float64 `json:"minCapacity" yaml:"minCapacity"`
}

// `ScheduledAction` is a property of the [AWS::ApplicationAutoScaling::ScalableTarget](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html) resource that specifies a scheduled action for a scalable target.
//
// For more information, see [PutScheduledAction](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PutScheduledAction.html) in the *Application Auto Scaling API Reference* . For more information about scheduled scaling, including the format for cron expressions, see [Scheduled scaling](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-scheduled-scaling.html) in the *Application Auto Scaling User Guide* .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import applicationautoscaling "github.com/aws/aws-cdk-go/awscdk/aws_applicationautoscaling"
//   scheduledActionProperty := &scheduledActionProperty{
//   	schedule: jsii.String("schedule"),
//   	scheduledActionName: jsii.String("scheduledActionName"),
//
//   	// the properties below are optional
//   	endTime: NewDate(),
//   	scalableTargetAction: &scalableTargetActionProperty{
//   		maxCapacity: jsii.Number(123),
//   		minCapacity: jsii.Number(123),
//   	},
//   	startTime: NewDate(),
//   	timezone: jsii.String("timezone"),
//   }
//
type CfnScalableTarget_ScheduledActionProperty struct {
	// The schedule for this action. The following formats are supported:.
	//
	// - At expressions - " `at( *yyyy* - *mm* - *dd* T *hh* : *mm* : *ss* )` "
	// - Rate expressions - " `rate( *value* *unit* )` "
	// - Cron expressions - " `cron( *fields* )` "
	//
	// At expressions are useful for one-time schedules. Cron expressions are useful for scheduled actions that run periodically at a specified date and time, and rate expressions are useful for scheduled actions that run at a regular interval.
	//
	// At and cron expressions use Universal Coordinated Time (UTC) by default.
	//
	// The cron format consists of six fields separated by white spaces: [Minutes] [Hours] [Day_of_Month] [Month] [Day_of_Week] [Year].
	//
	// For rate expressions, *value* is a positive integer and *unit* is `minute` | `minutes` | `hour` | `hours` | `day` | `days` .
	Schedule *string `json:"schedule" yaml:"schedule"`
	// The name of the scheduled action.
	//
	// This name must be unique among all other scheduled actions on the specified scalable target.
	ScheduledActionName *string `json:"scheduledActionName" yaml:"scheduledActionName"`
	// The date and time that the action is scheduled to end, in UTC.
	EndTime interface{} `json:"endTime" yaml:"endTime"`
	// The new minimum and maximum capacity.
	//
	// You can set both values or just one. At the scheduled time, if the current capacity is below the minimum capacity, Application Auto Scaling scales out to the minimum capacity. If the current capacity is above the maximum capacity, Application Auto Scaling scales in to the maximum capacity.
	ScalableTargetAction interface{} `json:"scalableTargetAction" yaml:"scalableTargetAction"`
	// The date and time that the action is scheduled to begin, in UTC.
	StartTime interface{} `json:"startTime" yaml:"startTime"`
	// The time zone used when referring to the date and time of a scheduled action, when the scheduled action uses an at or cron expression.
	Timezone *string `json:"timezone" yaml:"timezone"`
}

// `SuspendedState` is a property of the [AWS::ApplicationAutoScaling::ScalableTarget](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html) resource that specifies whether the scaling activities for a scalable target are in a suspended state.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import applicationautoscaling "github.com/aws/aws-cdk-go/awscdk/aws_applicationautoscaling"
//   suspendedStateProperty := &suspendedStateProperty{
//   	dynamicScalingInSuspended: jsii.Boolean(false),
//   	dynamicScalingOutSuspended: jsii.Boolean(false),
//   	scheduledScalingSuspended: jsii.Boolean(false),
//   }
//
type CfnScalableTarget_SuspendedStateProperty struct {
	// Whether scale in by a target tracking scaling policy or a step scaling policy is suspended.
	//
	// Set the value to `true` if you don't want Application Auto Scaling to remove capacity when a scaling policy is triggered. The default is `false` .
	DynamicScalingInSuspended interface{} `json:"dynamicScalingInSuspended" yaml:"dynamicScalingInSuspended"`
	// Whether scale out by a target tracking scaling policy or a step scaling policy is suspended.
	//
	// Set the value to `true` if you don't want Application Auto Scaling to add capacity when a scaling policy is triggered. The default is `false` .
	DynamicScalingOutSuspended interface{} `json:"dynamicScalingOutSuspended" yaml:"dynamicScalingOutSuspended"`
	// Whether scheduled scaling is suspended.
	//
	// Set the value to `true` if you don't want Application Auto Scaling to add or remove capacity by initiating scheduled actions. The default is `false` .
	ScheduledScalingSuspended interface{} `json:"scheduledScalingSuspended" yaml:"scheduledScalingSuspended"`
}

// Properties for defining a `CfnScalableTarget`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import applicationautoscaling "github.com/aws/aws-cdk-go/awscdk/aws_applicationautoscaling"
//   cfnScalableTargetProps := &cfnScalableTargetProps{
//   	maxCapacity: jsii.Number(123),
//   	minCapacity: jsii.Number(123),
//   	resourceId: jsii.String("resourceId"),
//   	roleArn: jsii.String("roleArn"),
//   	scalableDimension: jsii.String("scalableDimension"),
//   	serviceNamespace: jsii.String("serviceNamespace"),
//
//   	// the properties below are optional
//   	scheduledActions: []interface{}{
//   		&scheduledActionProperty{
//   			schedule: jsii.String("schedule"),
//   			scheduledActionName: jsii.String("scheduledActionName"),
//
//   			// the properties below are optional
//   			endTime: NewDate(),
//   			scalableTargetAction: &scalableTargetActionProperty{
//   				maxCapacity: jsii.Number(123),
//   				minCapacity: jsii.Number(123),
//   			},
//   			startTime: NewDate(),
//   			timezone: jsii.String("timezone"),
//   		},
//   	},
//   	suspendedState: &suspendedStateProperty{
//   		dynamicScalingInSuspended: jsii.Boolean(false),
//   		dynamicScalingOutSuspended: jsii.Boolean(false),
//   		scheduledScalingSuspended: jsii.Boolean(false),
//   	},
//   }
//
type CfnScalableTargetProps struct {
	// The maximum value that you plan to scale out to.
	//
	// When a scaling policy is in effect, Application Auto Scaling can scale out (expand) as needed to the maximum capacity limit in response to changing demand.
	MaxCapacity *float64 `json:"maxCapacity" yaml:"maxCapacity"`
	// The minimum value that you plan to scale in to.
	//
	// When a scaling policy is in effect, Application Auto Scaling can scale in (contract) as needed to the minimum capacity limit in response to changing demand.
	MinCapacity *float64 `json:"minCapacity" yaml:"minCapacity"`
	// The identifier of the resource associated with the scalable target.
	//
	// This string consists of the resource type and unique identifier.
	//
	// - ECS service - The resource type is `service` and the unique identifier is the cluster name and service name. Example: `service/default/sample-webapp` .
	// - Spot Fleet - The resource type is `spot-fleet-request` and the unique identifier is the Spot Fleet request ID. Example: `spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE` .
	// - EMR cluster - The resource type is `instancegroup` and the unique identifier is the cluster ID and instance group ID. Example: `instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0` .
	// - AppStream 2.0 fleet - The resource type is `fleet` and the unique identifier is the fleet name. Example: `fleet/sample-fleet` .
	// - DynamoDB table - The resource type is `table` and the unique identifier is the table name. Example: `table/my-table` .
	// - DynamoDB global secondary index - The resource type is `index` and the unique identifier is the index name. Example: `table/my-table/index/my-table-index` .
	// - Aurora DB cluster - The resource type is `cluster` and the unique identifier is the cluster name. Example: `cluster:my-db-cluster` .
	// - SageMaker endpoint variant - The resource type is `variant` and the unique identifier is the resource ID. Example: `endpoint/my-end-point/variant/KMeansClustering` .
	// - Custom resources are not supported with a resource type. This parameter must specify the `OutputValue` from the CloudFormation template stack used to access the resources. The unique identifier is defined by the service provider. More information is available in our [GitHub repository](https://docs.aws.amazon.com/https://github.com/aws/aws-auto-scaling-custom-resource) .
	// - Amazon Comprehend document classification endpoint - The resource type and unique identifier are specified using the endpoint ARN. Example: `arn:aws:comprehend:us-west-2:123456789012:document-classifier-endpoint/EXAMPLE` .
	// - Amazon Comprehend entity recognizer endpoint - The resource type and unique identifier are specified using the endpoint ARN. Example: `arn:aws:comprehend:us-west-2:123456789012:entity-recognizer-endpoint/EXAMPLE` .
	// - Lambda provisioned concurrency - The resource type is `function` and the unique identifier is the function name with a function version or alias name suffix that is not `$LATEST` . Example: `function:my-function:prod` or `function:my-function:1` .
	// - Amazon Keyspaces table - The resource type is `table` and the unique identifier is the table name. Example: `keyspace/mykeyspace/table/mytable` .
	// - Amazon MSK cluster - The resource type and unique identifier are specified using the cluster ARN. Example: `arn:aws:kafka:us-east-1:123456789012:cluster/demo-cluster-1/6357e0b2-0e6a-4b86-a0b4-70df934c2e31-5` .
	// - Amazon ElastiCache replication group - The resource type is `replication-group` and the unique identifier is the replication group name. Example: `replication-group/mycluster` .
	// - Neptune cluster - The resource type is `cluster` and the unique identifier is the cluster name. Example: `cluster:mycluster` .
	ResourceId *string `json:"resourceId" yaml:"resourceId"`
	// Specify the Amazon Resource Name (ARN) of an Identity and Access Management (IAM) role that allows Application Auto Scaling to modify the scalable target on your behalf.
	//
	// This can be either an IAM service role that Application Auto Scaling can assume to make calls to other AWS resources on your behalf, or a service-linked role for the specified service. For more information, see [How Application Auto Scaling works with IAM](https://docs.aws.amazon.com/autoscaling/application/userguide/security_iam_service-with-iam.html) in the *Application Auto Scaling User Guide* .
	//
	// To automatically create a service-linked role (recommended), specify the full ARN of the service-linked role in your stack template. To find the exact ARN of the service-linked role for your AWS or custom resource, see the [Service-linked roles](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-service-linked-roles.html) topic in the *Application Auto Scaling User Guide* . Look for the ARN in the table at the bottom of the page.
	RoleArn *string `json:"roleArn" yaml:"roleArn"`
	// The scalable dimension associated with the scalable target.
	//
	// This string consists of the service namespace, resource type, and scaling property.
	//
	// - `ecs:service:DesiredCount` - The desired task count of an ECS service.
	// - `elasticmapreduce:instancegroup:InstanceCount` - The instance count of an EMR Instance Group.
	// - `ec2:spot-fleet-request:TargetCapacity` - The target capacity of a Spot Fleet.
	// - `appstream:fleet:DesiredCapacity` - The desired capacity of an AppStream 2.0 fleet.
	// - `dynamodb:table:ReadCapacityUnits` - The provisioned read capacity for a DynamoDB table.
	// - `dynamodb:table:WriteCapacityUnits` - The provisioned write capacity for a DynamoDB table.
	// - `dynamodb:index:ReadCapacityUnits` - The provisioned read capacity for a DynamoDB global secondary index.
	// - `dynamodb:index:WriteCapacityUnits` - The provisioned write capacity for a DynamoDB global secondary index.
	// - `rds:cluster:ReadReplicaCount` - The count of Aurora Replicas in an Aurora DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible edition.
	// - `sagemaker:variant:DesiredInstanceCount` - The number of EC2 instances for a SageMaker model endpoint variant.
	// - `custom-resource:ResourceType:Property` - The scalable dimension for a custom resource provided by your own application or service.
	// - `comprehend:document-classifier-endpoint:DesiredInferenceUnits` - The number of inference units for an Amazon Comprehend document classification endpoint.
	// - `comprehend:entity-recognizer-endpoint:DesiredInferenceUnits` - The number of inference units for an Amazon Comprehend entity recognizer endpoint.
	// - `lambda:function:ProvisionedConcurrency` - The provisioned concurrency for a Lambda function.
	// - `cassandra:table:ReadCapacityUnits` - The provisioned read capacity for an Amazon Keyspaces table.
	// - `cassandra:table:WriteCapacityUnits` - The provisioned write capacity for an Amazon Keyspaces table.
	// - `kafka:broker-storage:VolumeSize` - The provisioned volume size (in GiB) for brokers in an Amazon MSK cluster.
	// - `elasticache:replication-group:NodeGroups` - The number of node groups for an Amazon ElastiCache replication group.
	// - `elasticache:replication-group:Replicas` - The number of replicas per node group for an Amazon ElastiCache replication group.
	// - `neptune:cluster:ReadReplicaCount` - The count of read replicas in an Amazon Neptune DB cluster.
	ScalableDimension *string `json:"scalableDimension" yaml:"scalableDimension"`
	// The namespace of the AWS service that provides the resource, or a `custom-resource` .
	ServiceNamespace *string `json:"serviceNamespace" yaml:"serviceNamespace"`
	// The scheduled actions for the scalable target. Duplicates aren't allowed.
	//
	// For more information about using scheduled scaling, see [Scheduled scaling](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-scheduled-scaling.html) in the *Application Auto Scaling User Guide* .
	ScheduledActions interface{} `json:"scheduledActions" yaml:"scheduledActions"`
	// An embedded object that contains attributes and attribute values that are used to suspend and resume automatic scaling.
	//
	// Setting the value of an attribute to `true` suspends the specified scaling activities. Setting it to `false` (default) resumes the specified scaling activities.
	//
	// *Suspension Outcomes*
	//
	// - For `DynamicScalingInSuspended` , while a suspension is in effect, all scale-in activities that are triggered by a scaling policy are suspended.
	// - For `DynamicScalingOutSuspended` , while a suspension is in effect, all scale-out activities that are triggered by a scaling policy are suspended.
	// - For `ScheduledScalingSuspended` , while a suspension is in effect, all scaling activities that involve scheduled actions are suspended.
	//
	// For more information, see [Suspending and resuming scaling](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-suspend-resume-scaling.html) in the *Application Auto Scaling User Guide* .
	SuspendedState interface{} `json:"suspendedState" yaml:"suspendedState"`
}

// A CloudFormation `AWS::ApplicationAutoScaling::ScalingPolicy`.
//
// The `AWS::ApplicationAutoScaling::ScalingPolicy` resource defines a scaling policy that Application Auto Scaling uses to adjust the capacity of a scalable target.
//
// For more information, see [PutScalingPolicy](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PutScalingPolicy.html) in the *Application Auto Scaling API Reference* . For more information about Application Auto Scaling scaling policies, see [Target tracking scaling policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-target-tracking.html) and [Step scaling policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html) in the *Application Auto Scaling User Guide* .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import applicationautoscaling "github.com/aws/aws-cdk-go/awscdk/aws_applicationautoscaling"
//   cfnScalingPolicy := applicationautoscaling.NewCfnScalingPolicy(this, jsii.String("MyCfnScalingPolicy"), &cfnScalingPolicyProps{
//   	policyName: jsii.String("policyName"),
//   	policyType: jsii.String("policyType"),
//
//   	// the properties below are optional
//   	resourceId: jsii.String("resourceId"),
//   	scalableDimension: jsii.String("scalableDimension"),
//   	scalingTargetId: jsii.String("scalingTargetId"),
//   	serviceNamespace: jsii.String("serviceNamespace"),
//   	stepScalingPolicyConfiguration: &stepScalingPolicyConfigurationProperty{
//   		adjustmentType: jsii.String("adjustmentType"),
//   		cooldown: jsii.Number(123),
//   		metricAggregationType: jsii.String("metricAggregationType"),
//   		minAdjustmentMagnitude: jsii.Number(123),
//   		stepAdjustments: []interface{}{
//   			&stepAdjustmentProperty{
//   				scalingAdjustment: jsii.Number(123),
//
//   				// the properties below are optional
//   				metricIntervalLowerBound: jsii.Number(123),
//   				metricIntervalUpperBound: jsii.Number(123),
//   			},
//   		},
//   	},
//   	targetTrackingScalingPolicyConfiguration: &targetTrackingScalingPolicyConfigurationProperty{
//   		targetValue: jsii.Number(123),
//
//   		// the properties below are optional
//   		customizedMetricSpecification: &customizedMetricSpecificationProperty{
//   			metricName: jsii.String("metricName"),
//   			namespace: jsii.String("namespace"),
//   			statistic: jsii.String("statistic"),
//
//   			// the properties below are optional
//   			dimensions: []interface{}{
//   				&metricDimensionProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			unit: jsii.String("unit"),
//   		},
//   		disableScaleIn: jsii.Boolean(false),
//   		predefinedMetricSpecification: &predefinedMetricSpecificationProperty{
//   			predefinedMetricType: jsii.String("predefinedMetricType"),
//
//   			// the properties below are optional
//   			resourceLabel: jsii.String("resourceLabel"),
//   		},
//   		scaleInCooldown: jsii.Number(123),
//   		scaleOutCooldown: jsii.Number(123),
//   	},
//   })
//
type CfnScalingPolicy interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The name of the scaling policy.
	//
	// Updates to the name of a target tracking scaling policy are not supported, unless you also update the metric used for scaling. To change only a target tracking scaling policy's name, first delete the policy by removing the existing `AWS::ApplicationAutoScaling::ScalingPolicy` resource from the template and updating the stack. Then, recreate the resource with the same settings and a different name.
	PolicyName() *string
	SetPolicyName(val *string)
	// The scaling policy type.
	//
	// The following policy types are supported:
	//
	// `TargetTrackingScaling` —Not supported for Amazon EMR
	//
	// `StepScaling` —Not supported for DynamoDB, Amazon Comprehend, Lambda, Amazon Keyspaces, Amazon MSK, Amazon ElastiCache, or Neptune.
	PolicyType() *string
	SetPolicyType(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The identifier of the resource associated with the scaling policy.
	//
	// This string consists of the resource type and unique identifier.
	//
	// - ECS service - The resource type is `service` and the unique identifier is the cluster name and service name. Example: `service/default/sample-webapp` .
	// - Spot Fleet - The resource type is `spot-fleet-request` and the unique identifier is the Spot Fleet request ID. Example: `spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE` .
	// - EMR cluster - The resource type is `instancegroup` and the unique identifier is the cluster ID and instance group ID. Example: `instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0` .
	// - AppStream 2.0 fleet - The resource type is `fleet` and the unique identifier is the fleet name. Example: `fleet/sample-fleet` .
	// - DynamoDB table - The resource type is `table` and the unique identifier is the table name. Example: `table/my-table` .
	// - DynamoDB global secondary index - The resource type is `index` and the unique identifier is the index name. Example: `table/my-table/index/my-table-index` .
	// - Aurora DB cluster - The resource type is `cluster` and the unique identifier is the cluster name. Example: `cluster:my-db-cluster` .
	// - SageMaker endpoint variant - The resource type is `variant` and the unique identifier is the resource ID. Example: `endpoint/my-end-point/variant/KMeansClustering` .
	// - Custom resources are not supported with a resource type. This parameter must specify the `OutputValue` from the CloudFormation template stack used to access the resources. The unique identifier is defined by the service provider. More information is available in our [GitHub repository](https://docs.aws.amazon.com/https://github.com/aws/aws-auto-scaling-custom-resource) .
	// - Amazon Comprehend document classification endpoint - The resource type and unique identifier are specified using the endpoint ARN. Example: `arn:aws:comprehend:us-west-2:123456789012:document-classifier-endpoint/EXAMPLE` .
	// - Amazon Comprehend entity recognizer endpoint - The resource type and unique identifier are specified using the endpoint ARN. Example: `arn:aws:comprehend:us-west-2:123456789012:entity-recognizer-endpoint/EXAMPLE` .
	// - Lambda provisioned concurrency - The resource type is `function` and the unique identifier is the function name with a function version or alias name suffix that is not `$LATEST` . Example: `function:my-function:prod` or `function:my-function:1` .
	// - Amazon Keyspaces table - The resource type is `table` and the unique identifier is the table name. Example: `keyspace/mykeyspace/table/mytable` .
	// - Amazon MSK cluster - The resource type and unique identifier are specified using the cluster ARN. Example: `arn:aws:kafka:us-east-1:123456789012:cluster/demo-cluster-1/6357e0b2-0e6a-4b86-a0b4-70df934c2e31-5` .
	// - Amazon ElastiCache replication group - The resource type is `replication-group` and the unique identifier is the replication group name. Example: `replication-group/mycluster` .
	// - Neptune cluster - The resource type is `cluster` and the unique identifier is the cluster name. Example: `cluster:mycluster` .
	ResourceId() *string
	SetResourceId(val *string)
	// The scalable dimension. This string consists of the service namespace, resource type, and scaling property.
	//
	// - `ecs:service:DesiredCount` - The desired task count of an ECS service.
	// - `elasticmapreduce:instancegroup:InstanceCount` - The instance count of an EMR Instance Group.
	// - `ec2:spot-fleet-request:TargetCapacity` - The target capacity of a Spot Fleet.
	// - `appstream:fleet:DesiredCapacity` - The desired capacity of an AppStream 2.0 fleet.
	// - `dynamodb:table:ReadCapacityUnits` - The provisioned read capacity for a DynamoDB table.
	// - `dynamodb:table:WriteCapacityUnits` - The provisioned write capacity for a DynamoDB table.
	// - `dynamodb:index:ReadCapacityUnits` - The provisioned read capacity for a DynamoDB global secondary index.
	// - `dynamodb:index:WriteCapacityUnits` - The provisioned write capacity for a DynamoDB global secondary index.
	// - `rds:cluster:ReadReplicaCount` - The count of Aurora Replicas in an Aurora DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible edition.
	// - `sagemaker:variant:DesiredInstanceCount` - The number of EC2 instances for a SageMaker model endpoint variant.
	// - `custom-resource:ResourceType:Property` - The scalable dimension for a custom resource provided by your own application or service.
	// - `comprehend:document-classifier-endpoint:DesiredInferenceUnits` - The number of inference units for an Amazon Comprehend document classification endpoint.
	// - `comprehend:entity-recognizer-endpoint:DesiredInferenceUnits` - The number of inference units for an Amazon Comprehend entity recognizer endpoint.
	// - `lambda:function:ProvisionedConcurrency` - The provisioned concurrency for a Lambda function.
	// - `cassandra:table:ReadCapacityUnits` - The provisioned read capacity for an Amazon Keyspaces table.
	// - `cassandra:table:WriteCapacityUnits` - The provisioned write capacity for an Amazon Keyspaces table.
	// - `kafka:broker-storage:VolumeSize` - The provisioned volume size (in GiB) for brokers in an Amazon MSK cluster.
	// - `elasticache:replication-group:NodeGroups` - The number of node groups for an Amazon ElastiCache replication group.
	// - `elasticache:replication-group:Replicas` - The number of replicas per node group for an Amazon ElastiCache replication group.
	// - `neptune:cluster:ReadReplicaCount` - The count of read replicas in an Amazon Neptune DB cluster.
	ScalableDimension() *string
	SetScalableDimension(val *string)
	// The CloudFormation-generated ID of an Application Auto Scaling scalable target.
	//
	// For more information about the ID, see the Return Value section of the `AWS::ApplicationAutoScaling::ScalableTarget` resource.
	//
	// > You must specify either the `ScalingTargetId` property, or the `ResourceId` , `ScalableDimension` , and `ServiceNamespace` properties, but not both.
	ScalingTargetId() *string
	SetScalingTargetId(val *string)
	// The namespace of the AWS service that provides the resource, or a `custom-resource` .
	ServiceNamespace() *string
	SetServiceNamespace(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// A step scaling policy.
	StepScalingPolicyConfiguration() interface{}
	SetStepScalingPolicyConfiguration(val interface{})
	// A target tracking scaling policy.
	TargetTrackingScalingPolicyConfiguration() interface{}
	SetTargetTrackingScalingPolicyConfiguration(val interface{})
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnScalingPolicy
type jsiiProxy_CfnScalingPolicy struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnScalingPolicy) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) PolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) PolicyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) ScalableDimension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalableDimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) ScalingTargetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingTargetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) ServiceNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) StepScalingPolicyConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stepScalingPolicyConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) TargetTrackingScalingPolicyConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetTrackingScalingPolicyConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApplicationAutoScaling::ScalingPolicy`.
func NewCfnScalingPolicy(scope awscdk.Construct, id *string, props *CfnScalingPolicyProps) CfnScalingPolicy {
	_init_.Initialize()

	j := jsiiProxy_CfnScalingPolicy{}

	_jsii_.Create(
		"monocdk.aws_applicationautoscaling.CfnScalingPolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApplicationAutoScaling::ScalingPolicy`.
func NewCfnScalingPolicy_Override(c CfnScalingPolicy, scope awscdk.Construct, id *string, props *CfnScalingPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_applicationautoscaling.CfnScalingPolicy",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnScalingPolicy) SetPolicyName(val *string) {
	_jsii_.Set(
		j,
		"policyName",
		val,
	)
}

func (j *jsiiProxy_CfnScalingPolicy) SetPolicyType(val *string) {
	_jsii_.Set(
		j,
		"policyType",
		val,
	)
}

func (j *jsiiProxy_CfnScalingPolicy) SetResourceId(val *string) {
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_CfnScalingPolicy) SetScalableDimension(val *string) {
	_jsii_.Set(
		j,
		"scalableDimension",
		val,
	)
}

func (j *jsiiProxy_CfnScalingPolicy) SetScalingTargetId(val *string) {
	_jsii_.Set(
		j,
		"scalingTargetId",
		val,
	)
}

func (j *jsiiProxy_CfnScalingPolicy) SetServiceNamespace(val *string) {
	_jsii_.Set(
		j,
		"serviceNamespace",
		val,
	)
}

func (j *jsiiProxy_CfnScalingPolicy) SetStepScalingPolicyConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"stepScalingPolicyConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnScalingPolicy) SetTargetTrackingScalingPolicyConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"targetTrackingScalingPolicyConfiguration",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnScalingPolicy_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_applicationautoscaling.CfnScalingPolicy",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnScalingPolicy_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_applicationautoscaling.CfnScalingPolicy",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnScalingPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_applicationautoscaling.CfnScalingPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnScalingPolicy_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_applicationautoscaling.CfnScalingPolicy",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnScalingPolicy) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnScalingPolicy) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnScalingPolicy) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnScalingPolicy) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnScalingPolicy) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnScalingPolicy) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnScalingPolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnScalingPolicy) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScalingPolicy) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScalingPolicy) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnScalingPolicy) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnScalingPolicy) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnScalingPolicy) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScalingPolicy) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnScalingPolicy) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnScalingPolicy) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScalingPolicy) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScalingPolicy) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnScalingPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScalingPolicy) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScalingPolicy) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Contains customized metric specification information for a target tracking scaling policy for Application Auto Scaling.
//
// For information about the available metrics for a service, see [AWS services that publish CloudWatch metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/aws-services-cloudwatch-metrics.html) in the *Amazon CloudWatch User Guide* .
//
// To create your customized metric specification:
//
// - Add values for each required parameter from CloudWatch. You can use an existing metric, or a new metric that you create. To use your own metric, you must first publish the metric to CloudWatch. For more information, see [Publish custom metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/publishingMetrics.html) in the *Amazon CloudWatch User Guide* .
// - Choose a metric that changes proportionally with capacity. The value of the metric should increase or decrease in inverse proportion to the number of capacity units. That is, the value of the metric should decrease when capacity increases, and increase when capacity decreases.
//
// For an example of how creating new metrics can be useful, see [Scaling based on Amazon SQS](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-using-sqs-queue.html) in the *Amazon EC2 Auto Scaling User Guide* . This topic mentions Auto Scaling groups, but the same scenario for Amazon SQS can apply to the target tracking scaling policies that you create for a Spot Fleet by using Application Auto Scaling.
//
// For more information about the CloudWatch terminology below, see [Amazon CloudWatch concepts](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html) .
//
// `CustomizedMetricSpecification` is a property of the [AWS::ApplicationAutoScaling::ScalingPolicy TargetTrackingScalingPolicyConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration.html) property type.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import applicationautoscaling "github.com/aws/aws-cdk-go/awscdk/aws_applicationautoscaling"
//   customizedMetricSpecificationProperty := &customizedMetricSpecificationProperty{
//   	metricName: jsii.String("metricName"),
//   	namespace: jsii.String("namespace"),
//   	statistic: jsii.String("statistic"),
//
//   	// the properties below are optional
//   	dimensions: []interface{}{
//   		&metricDimensionProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	unit: jsii.String("unit"),
//   }
//
type CfnScalingPolicy_CustomizedMetricSpecificationProperty struct {
	// The name of the metric.
	//
	// To get the exact metric name, namespace, and dimensions, inspect the [Metric](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_Metric.html) object that is returned by a call to [ListMetrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_ListMetrics.html) .
	MetricName *string `json:"metricName" yaml:"metricName"`
	// The namespace of the metric.
	Namespace *string `json:"namespace" yaml:"namespace"`
	// The statistic of the metric.
	Statistic *string `json:"statistic" yaml:"statistic"`
	// The dimensions of the metric.
	//
	// Conditional: If you published your metric with dimensions, you must specify the same dimensions in your scaling policy.
	Dimensions interface{} `json:"dimensions" yaml:"dimensions"`
	// The unit of the metric.
	//
	// For a complete list of the units that CloudWatch supports, see the [MetricDatum](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_MetricDatum.html) data type in the *Amazon CloudWatch API Reference* .
	Unit *string `json:"unit" yaml:"unit"`
}

// `MetricDimension` specifies a name/value pair that is part of the identity of a CloudWatch metric for the `Dimensions` property of the [AWS::ApplicationAutoScaling::ScalingPolicy CustomizedMetricSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-customizedmetricspecification.html) property type. Duplicate dimensions are not allowed.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import applicationautoscaling "github.com/aws/aws-cdk-go/awscdk/aws_applicationautoscaling"
//   metricDimensionProperty := &metricDimensionProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnScalingPolicy_MetricDimensionProperty struct {
	// The name of the dimension.
	Name *string `json:"name" yaml:"name"`
	// The value of the dimension.
	Value *string `json:"value" yaml:"value"`
}

// Contains predefined metric specification information for a target tracking scaling policy for Application Auto Scaling.
//
// `PredefinedMetricSpecification` is a property of the [AWS::ApplicationAutoScaling::ScalingPolicy TargetTrackingScalingPolicyConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-targettrackingscalingpolicyconfiguration.html) property type.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import applicationautoscaling "github.com/aws/aws-cdk-go/awscdk/aws_applicationautoscaling"
//   predefinedMetricSpecificationProperty := &predefinedMetricSpecificationProperty{
//   	predefinedMetricType: jsii.String("predefinedMetricType"),
//
//   	// the properties below are optional
//   	resourceLabel: jsii.String("resourceLabel"),
//   }
//
type CfnScalingPolicy_PredefinedMetricSpecificationProperty struct {
	// The metric type.
	//
	// The `ALBRequestCountPerTarget` metric type applies only to Spot fleet requests and ECS services.
	PredefinedMetricType *string `json:"predefinedMetricType" yaml:"predefinedMetricType"`
	// Identifies the resource associated with the metric type.
	//
	// You can't specify a resource label unless the metric type is `ALBRequestCountPerTarget` and there is a target group attached to the Spot Fleet or ECS service.
	//
	// You create the resource label by appending the final portion of the load balancer ARN and the final portion of the target group ARN into a single value, separated by a forward slash (/). The format of the resource label is:
	//
	// `app/my-alb/778d41231b141a0f/targetgroup/my-alb-target-group/943f017f100becff` .
	//
	// Where:
	//
	// - app/<load-balancer-name>/<load-balancer-id> is the final portion of the load balancer ARN
	// - targetgroup/<target-group-name>/<target-group-id> is the final portion of the target group ARN.
	//
	// To find the ARN for an Application Load Balancer, use the [DescribeLoadBalancers](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_DescribeLoadBalancers.html) API operation. To find the ARN for the target group, use the [DescribeTargetGroups](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_DescribeTargetGroups.html) API operation.
	ResourceLabel *string `json:"resourceLabel" yaml:"resourceLabel"`
}

// `StepAdjustment` specifies a step adjustment for the `StepAdjustments` property of the [AWS::ApplicationAutoScaling::ScalingPolicy StepScalingPolicyConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-stepscalingpolicyconfiguration.html) property type.
//
// For the following examples, suppose that you have an alarm with a breach threshold of 50:
//
// - To trigger a step adjustment when the metric is greater than or equal to 50 and less than 60, specify a lower bound of 0 and an upper bound of 10.
// - To trigger a step adjustment when the metric is greater than 40 and less than or equal to 50, specify a lower bound of -10 and an upper bound of 0.
//
// For more information, see [Step adjustments](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html#as-scaling-steps) in the *Application Auto Scaling User Guide* .
//
// You can find a sample template snippet in the [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html#aws-resource-applicationautoscaling-scalingpolicy--examples) section of the `AWS::ApplicationAutoScaling::ScalingPolicy` documentation.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import applicationautoscaling "github.com/aws/aws-cdk-go/awscdk/aws_applicationautoscaling"
//   stepAdjustmentProperty := &stepAdjustmentProperty{
//   	scalingAdjustment: jsii.Number(123),
//
//   	// the properties below are optional
//   	metricIntervalLowerBound: jsii.Number(123),
//   	metricIntervalUpperBound: jsii.Number(123),
//   }
//
type CfnScalingPolicy_StepAdjustmentProperty struct {
	// The amount by which to scale.
	//
	// The adjustment is based on the value that you specified in the `AdjustmentType` property (either an absolute number or a percentage). A positive value adds to the current capacity and a negative number subtracts from the current capacity.
	ScalingAdjustment *float64 `json:"scalingAdjustment" yaml:"scalingAdjustment"`
	// The lower bound for the difference between the alarm threshold and the CloudWatch metric.
	//
	// If the metric value is above the breach threshold, the lower bound is inclusive (the metric must be greater than or equal to the threshold plus the lower bound). Otherwise, it is exclusive (the metric must be greater than the threshold plus the lower bound). A null value indicates negative infinity.
	//
	// You must specify at least one upper or lower bound.
	MetricIntervalLowerBound *float64 `json:"metricIntervalLowerBound" yaml:"metricIntervalLowerBound"`
	// The upper bound for the difference between the alarm threshold and the CloudWatch metric.
	//
	// If the metric value is above the breach threshold, the upper bound is exclusive (the metric must be less than the threshold plus the upper bound). Otherwise, it is inclusive (the metric must be less than or equal to the threshold plus the upper bound). A null value indicates positive infinity.
	//
	// You must specify at least one upper or lower bound.
	MetricIntervalUpperBound *float64 `json:"metricIntervalUpperBound" yaml:"metricIntervalUpperBound"`
}

// `StepScalingPolicyConfiguration` is a property of the [AWS::ApplicationAutoScaling::ScalingPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html) resource that specifies a step scaling policy configuration for Application Auto Scaling.
//
// For more information, see [PutScalingPolicy](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PutScalingPolicy.html) in the *Application Auto Scaling API Reference* . For more information about step scaling policies, see [Step scaling policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html) in the *Application Auto Scaling User Guide* .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import applicationautoscaling "github.com/aws/aws-cdk-go/awscdk/aws_applicationautoscaling"
//   stepScalingPolicyConfigurationProperty := &stepScalingPolicyConfigurationProperty{
//   	adjustmentType: jsii.String("adjustmentType"),
//   	cooldown: jsii.Number(123),
//   	metricAggregationType: jsii.String("metricAggregationType"),
//   	minAdjustmentMagnitude: jsii.Number(123),
//   	stepAdjustments: []interface{}{
//   		&stepAdjustmentProperty{
//   			scalingAdjustment: jsii.Number(123),
//
//   			// the properties below are optional
//   			metricIntervalLowerBound: jsii.Number(123),
//   			metricIntervalUpperBound: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnScalingPolicy_StepScalingPolicyConfigurationProperty struct {
	// Specifies whether the `ScalingAdjustment` value in the `StepAdjustment` property is an absolute number or a percentage of the current capacity.
	AdjustmentType *string `json:"adjustmentType" yaml:"adjustmentType"`
	// The amount of time, in seconds, to wait for a previous scaling activity to take effect.
	//
	// With scale-out policies, the intention is to continuously (but not excessively) scale out. After Application Auto Scaling successfully scales out using a step scaling policy, it starts to calculate the cooldown time. The scaling policy won't increase the desired capacity again unless either a larger scale out is triggered or the cooldown period ends. While the cooldown period is in effect, capacity added by the initiating scale-out activity is calculated as part of the desired capacity for the next scale-out activity. For example, when an alarm triggers a step scaling policy to increase the capacity by 2, the scaling activity completes successfully, and a cooldown period starts. If the alarm triggers again during the cooldown period but at a more aggressive step adjustment of 3, the previous increase of 2 is considered part of the current capacity. Therefore, only 1 is added to the capacity.
	//
	// With scale-in policies, the intention is to scale in conservatively to protect your application’s availability, so scale-in activities are blocked until the cooldown period has expired. However, if another alarm triggers a scale-out activity during the cooldown period after a scale-in activity, Application Auto Scaling scales out the target immediately. In this case, the cooldown period for the scale-in activity stops and doesn't complete.
	//
	// Application Auto Scaling provides a default value of 600 for Amazon ElastiCache replication groups and a default value of 300 for the following scalable targets:
	//
	// - AppStream 2.0 fleets
	// - Aurora DB clusters
	// - ECS services
	// - EMR clusters
	// - Neptune clusters
	// - SageMaker endpoint variants
	// - Spot Fleets
	// - Custom resources
	//
	// For all other scalable targets, the default value is 0:
	//
	// - Amazon Comprehend document classification and entity recognizer endpoints
	// - DynamoDB tables and global secondary indexes
	// - Amazon Keyspaces tables
	// - Lambda provisioned concurrency
	// - Amazon MSK broker storage.
	Cooldown *float64 `json:"cooldown" yaml:"cooldown"`
	// The aggregation type for the CloudWatch metrics.
	//
	// Valid values are `Minimum` , `Maximum` , and `Average` . If the aggregation type is null, the value is treated as `Average` .
	MetricAggregationType *string `json:"metricAggregationType" yaml:"metricAggregationType"`
	// The minimum value to scale by when the adjustment type is `PercentChangeInCapacity` .
	//
	// For example, suppose that you create a step scaling policy to scale out an Amazon ECS service by 25 percent and you specify a `MinAdjustmentMagnitude` of 2. If the service has 4 tasks and the scaling policy is performed, 25 percent of 4 is 1. However, because you specified a `MinAdjustmentMagnitude` of 2, Application Auto Scaling scales out the service by 2 tasks.
	MinAdjustmentMagnitude *float64 `json:"minAdjustmentMagnitude" yaml:"minAdjustmentMagnitude"`
	// A set of adjustments that enable you to scale based on the size of the alarm breach.
	//
	// At least one step adjustment is required if you are adding a new step scaling policy configuration.
	StepAdjustments interface{} `json:"stepAdjustments" yaml:"stepAdjustments"`
}

// `TargetTrackingScalingPolicyConfiguration` is a property of the [AWS::ApplicationAutoScaling::ScalingPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalingpolicy.html) resource that specifies a target tracking scaling policy configuration for Application Auto Scaling. Use a target tracking scaling policy to adjust the capacity of the specified scalable target in response to actual workloads, so that resource utilization remains at or near the target utilization value.
//
// For more information, see [PutScalingPolicy](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PutScalingPolicy.html) in the *Application Auto Scaling API Reference* . For more information about target tracking scaling policies, see [Target tracking scaling policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-target-tracking.html) in the *Application Auto Scaling User Guide* .
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import applicationautoscaling "github.com/aws/aws-cdk-go/awscdk/aws_applicationautoscaling"
//   targetTrackingScalingPolicyConfigurationProperty := &targetTrackingScalingPolicyConfigurationProperty{
//   	targetValue: jsii.Number(123),
//
//   	// the properties below are optional
//   	customizedMetricSpecification: &customizedMetricSpecificationProperty{
//   		metricName: jsii.String("metricName"),
//   		namespace: jsii.String("namespace"),
//   		statistic: jsii.String("statistic"),
//
//   		// the properties below are optional
//   		dimensions: []interface{}{
//   			&metricDimensionProperty{
//   				name: jsii.String("name"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		unit: jsii.String("unit"),
//   	},
//   	disableScaleIn: jsii.Boolean(false),
//   	predefinedMetricSpecification: &predefinedMetricSpecificationProperty{
//   		predefinedMetricType: jsii.String("predefinedMetricType"),
//
//   		// the properties below are optional
//   		resourceLabel: jsii.String("resourceLabel"),
//   	},
//   	scaleInCooldown: jsii.Number(123),
//   	scaleOutCooldown: jsii.Number(123),
//   }
//
type CfnScalingPolicy_TargetTrackingScalingPolicyConfigurationProperty struct {
	// The target value for the metric.
	//
	// Although this property accepts numbers of type Double, it won't accept values that are either too small or too large. Values must be in the range of -2^360 to 2^360. The value must be a valid number based on the choice of metric. For example, if the metric is CPU utilization, then the target value is a percent value that represents how much of the CPU can be used before scaling out.
	TargetValue *float64 `json:"targetValue" yaml:"targetValue"`
	// A customized metric.
	//
	// You can specify either a predefined metric or a customized metric.
	CustomizedMetricSpecification interface{} `json:"customizedMetricSpecification" yaml:"customizedMetricSpecification"`
	// Indicates whether scale in by the target tracking scaling policy is disabled.
	//
	// If the value is `true` , scale in is disabled and the target tracking scaling policy won't remove capacity from the scalable target. Otherwise, scale in is enabled and the target tracking scaling policy can remove capacity from the scalable target. The default value is `false` .
	DisableScaleIn interface{} `json:"disableScaleIn" yaml:"disableScaleIn"`
	// A predefined metric.
	//
	// You can specify either a predefined metric or a customized metric.
	PredefinedMetricSpecification interface{} `json:"predefinedMetricSpecification" yaml:"predefinedMetricSpecification"`
	// The amount of time, in seconds, after a scale-in activity completes before another scale-in activity can start.
	//
	// With the *scale-in cooldown period* , the intention is to scale in conservatively to protect your application’s availability, so scale-in activities are blocked until the cooldown period has expired. However, if another alarm triggers a scale-out activity during the scale-in cooldown period, Application Auto Scaling scales out the target immediately. In this case, the scale-in cooldown period stops and doesn't complete.
	//
	// Application Auto Scaling provides a default value of 600 for Amazon ElastiCache replication groups and a default value of 300 for the following scalable targets:
	//
	// - AppStream 2.0 fleets
	// - Aurora DB clusters
	// - ECS services
	// - EMR clusters
	// - Neptune clusters
	// - SageMaker endpoint variants
	// - Spot Fleets
	// - Custom resources
	//
	// For all other scalable targets, the default value is 0:
	//
	// - Amazon Comprehend document classification and entity recognizer endpoints
	// - DynamoDB tables and global secondary indexes
	// - Amazon Keyspaces tables
	// - Lambda provisioned concurrency
	// - Amazon MSK broker storage.
	ScaleInCooldown *float64 `json:"scaleInCooldown" yaml:"scaleInCooldown"`
	// The amount of time, in seconds, to wait for a previous scale-out activity to take effect.
	//
	// With the *scale-out cooldown period* , the intention is to continuously (but not excessively) scale out. After Application Auto Scaling successfully scales out using a target tracking scaling policy, it starts to calculate the cooldown time. The scaling policy won't increase the desired capacity again unless either a larger scale out is triggered or the cooldown period ends. While the cooldown period is in effect, the capacity added by the initiating scale-out activity is calculated as part of the desired capacity for the next scale-out activity.
	//
	// Application Auto Scaling provides a default value of 600 for Amazon ElastiCache replication groups and a default value of 300 for the following scalable targets:
	//
	// - AppStream 2.0 fleets
	// - Aurora DB clusters
	// - ECS services
	// - EMR clusters
	// - Neptune clusters
	// - SageMaker endpoint variants
	// - Spot Fleets
	// - Custom resources
	//
	// For all other scalable targets, the default value is 0:
	//
	// - Amazon Comprehend document classification and entity recognizer endpoints
	// - DynamoDB tables and global secondary indexes
	// - Amazon Keyspaces tables
	// - Lambda provisioned concurrency
	// - Amazon MSK broker storage.
	ScaleOutCooldown *float64 `json:"scaleOutCooldown" yaml:"scaleOutCooldown"`
}

// Properties for defining a `CfnScalingPolicy`.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import applicationautoscaling "github.com/aws/aws-cdk-go/awscdk/aws_applicationautoscaling"
//   cfnScalingPolicyProps := &cfnScalingPolicyProps{
//   	policyName: jsii.String("policyName"),
//   	policyType: jsii.String("policyType"),
//
//   	// the properties below are optional
//   	resourceId: jsii.String("resourceId"),
//   	scalableDimension: jsii.String("scalableDimension"),
//   	scalingTargetId: jsii.String("scalingTargetId"),
//   	serviceNamespace: jsii.String("serviceNamespace"),
//   	stepScalingPolicyConfiguration: &stepScalingPolicyConfigurationProperty{
//   		adjustmentType: jsii.String("adjustmentType"),
//   		cooldown: jsii.Number(123),
//   		metricAggregationType: jsii.String("metricAggregationType"),
//   		minAdjustmentMagnitude: jsii.Number(123),
//   		stepAdjustments: []interface{}{
//   			&stepAdjustmentProperty{
//   				scalingAdjustment: jsii.Number(123),
//
//   				// the properties below are optional
//   				metricIntervalLowerBound: jsii.Number(123),
//   				metricIntervalUpperBound: jsii.Number(123),
//   			},
//   		},
//   	},
//   	targetTrackingScalingPolicyConfiguration: &targetTrackingScalingPolicyConfigurationProperty{
//   		targetValue: jsii.Number(123),
//
//   		// the properties below are optional
//   		customizedMetricSpecification: &customizedMetricSpecificationProperty{
//   			metricName: jsii.String("metricName"),
//   			namespace: jsii.String("namespace"),
//   			statistic: jsii.String("statistic"),
//
//   			// the properties below are optional
//   			dimensions: []interface{}{
//   				&metricDimensionProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			unit: jsii.String("unit"),
//   		},
//   		disableScaleIn: jsii.Boolean(false),
//   		predefinedMetricSpecification: &predefinedMetricSpecificationProperty{
//   			predefinedMetricType: jsii.String("predefinedMetricType"),
//
//   			// the properties below are optional
//   			resourceLabel: jsii.String("resourceLabel"),
//   		},
//   		scaleInCooldown: jsii.Number(123),
//   		scaleOutCooldown: jsii.Number(123),
//   	},
//   }
//
type CfnScalingPolicyProps struct {
	// The name of the scaling policy.
	//
	// Updates to the name of a target tracking scaling policy are not supported, unless you also update the metric used for scaling. To change only a target tracking scaling policy's name, first delete the policy by removing the existing `AWS::ApplicationAutoScaling::ScalingPolicy` resource from the template and updating the stack. Then, recreate the resource with the same settings and a different name.
	PolicyName *string `json:"policyName" yaml:"policyName"`
	// The scaling policy type.
	//
	// The following policy types are supported:
	//
	// `TargetTrackingScaling` —Not supported for Amazon EMR
	//
	// `StepScaling` —Not supported for DynamoDB, Amazon Comprehend, Lambda, Amazon Keyspaces, Amazon MSK, Amazon ElastiCache, or Neptune.
	PolicyType *string `json:"policyType" yaml:"policyType"`
	// The identifier of the resource associated with the scaling policy.
	//
	// This string consists of the resource type and unique identifier.
	//
	// - ECS service - The resource type is `service` and the unique identifier is the cluster name and service name. Example: `service/default/sample-webapp` .
	// - Spot Fleet - The resource type is `spot-fleet-request` and the unique identifier is the Spot Fleet request ID. Example: `spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE` .
	// - EMR cluster - The resource type is `instancegroup` and the unique identifier is the cluster ID and instance group ID. Example: `instancegroup/j-2EEZNYKUA1NTV/ig-1791Y4E1L8YI0` .
	// - AppStream 2.0 fleet - The resource type is `fleet` and the unique identifier is the fleet name. Example: `fleet/sample-fleet` .
	// - DynamoDB table - The resource type is `table` and the unique identifier is the table name. Example: `table/my-table` .
	// - DynamoDB global secondary index - The resource type is `index` and the unique identifier is the index name. Example: `table/my-table/index/my-table-index` .
	// - Aurora DB cluster - The resource type is `cluster` and the unique identifier is the cluster name. Example: `cluster:my-db-cluster` .
	// - SageMaker endpoint variant - The resource type is `variant` and the unique identifier is the resource ID. Example: `endpoint/my-end-point/variant/KMeansClustering` .
	// - Custom resources are not supported with a resource type. This parameter must specify the `OutputValue` from the CloudFormation template stack used to access the resources. The unique identifier is defined by the service provider. More information is available in our [GitHub repository](https://docs.aws.amazon.com/https://github.com/aws/aws-auto-scaling-custom-resource) .
	// - Amazon Comprehend document classification endpoint - The resource type and unique identifier are specified using the endpoint ARN. Example: `arn:aws:comprehend:us-west-2:123456789012:document-classifier-endpoint/EXAMPLE` .
	// - Amazon Comprehend entity recognizer endpoint - The resource type and unique identifier are specified using the endpoint ARN. Example: `arn:aws:comprehend:us-west-2:123456789012:entity-recognizer-endpoint/EXAMPLE` .
	// - Lambda provisioned concurrency - The resource type is `function` and the unique identifier is the function name with a function version or alias name suffix that is not `$LATEST` . Example: `function:my-function:prod` or `function:my-function:1` .
	// - Amazon Keyspaces table - The resource type is `table` and the unique identifier is the table name. Example: `keyspace/mykeyspace/table/mytable` .
	// - Amazon MSK cluster - The resource type and unique identifier are specified using the cluster ARN. Example: `arn:aws:kafka:us-east-1:123456789012:cluster/demo-cluster-1/6357e0b2-0e6a-4b86-a0b4-70df934c2e31-5` .
	// - Amazon ElastiCache replication group - The resource type is `replication-group` and the unique identifier is the replication group name. Example: `replication-group/mycluster` .
	// - Neptune cluster - The resource type is `cluster` and the unique identifier is the cluster name. Example: `cluster:mycluster` .
	ResourceId *string `json:"resourceId" yaml:"resourceId"`
	// The scalable dimension. This string consists of the service namespace, resource type, and scaling property.
	//
	// - `ecs:service:DesiredCount` - The desired task count of an ECS service.
	// - `elasticmapreduce:instancegroup:InstanceCount` - The instance count of an EMR Instance Group.
	// - `ec2:spot-fleet-request:TargetCapacity` - The target capacity of a Spot Fleet.
	// - `appstream:fleet:DesiredCapacity` - The desired capacity of an AppStream 2.0 fleet.
	// - `dynamodb:table:ReadCapacityUnits` - The provisioned read capacity for a DynamoDB table.
	// - `dynamodb:table:WriteCapacityUnits` - The provisioned write capacity for a DynamoDB table.
	// - `dynamodb:index:ReadCapacityUnits` - The provisioned read capacity for a DynamoDB global secondary index.
	// - `dynamodb:index:WriteCapacityUnits` - The provisioned write capacity for a DynamoDB global secondary index.
	// - `rds:cluster:ReadReplicaCount` - The count of Aurora Replicas in an Aurora DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible edition.
	// - `sagemaker:variant:DesiredInstanceCount` - The number of EC2 instances for a SageMaker model endpoint variant.
	// - `custom-resource:ResourceType:Property` - The scalable dimension for a custom resource provided by your own application or service.
	// - `comprehend:document-classifier-endpoint:DesiredInferenceUnits` - The number of inference units for an Amazon Comprehend document classification endpoint.
	// - `comprehend:entity-recognizer-endpoint:DesiredInferenceUnits` - The number of inference units for an Amazon Comprehend entity recognizer endpoint.
	// - `lambda:function:ProvisionedConcurrency` - The provisioned concurrency for a Lambda function.
	// - `cassandra:table:ReadCapacityUnits` - The provisioned read capacity for an Amazon Keyspaces table.
	// - `cassandra:table:WriteCapacityUnits` - The provisioned write capacity for an Amazon Keyspaces table.
	// - `kafka:broker-storage:VolumeSize` - The provisioned volume size (in GiB) for brokers in an Amazon MSK cluster.
	// - `elasticache:replication-group:NodeGroups` - The number of node groups for an Amazon ElastiCache replication group.
	// - `elasticache:replication-group:Replicas` - The number of replicas per node group for an Amazon ElastiCache replication group.
	// - `neptune:cluster:ReadReplicaCount` - The count of read replicas in an Amazon Neptune DB cluster.
	ScalableDimension *string `json:"scalableDimension" yaml:"scalableDimension"`
	// The CloudFormation-generated ID of an Application Auto Scaling scalable target.
	//
	// For more information about the ID, see the Return Value section of the `AWS::ApplicationAutoScaling::ScalableTarget` resource.
	//
	// > You must specify either the `ScalingTargetId` property, or the `ResourceId` , `ScalableDimension` , and `ServiceNamespace` properties, but not both.
	ScalingTargetId *string `json:"scalingTargetId" yaml:"scalingTargetId"`
	// The namespace of the AWS service that provides the resource, or a `custom-resource` .
	ServiceNamespace *string `json:"serviceNamespace" yaml:"serviceNamespace"`
	// A step scaling policy.
	StepScalingPolicyConfiguration interface{} `json:"stepScalingPolicyConfiguration" yaml:"stepScalingPolicyConfiguration"`
	// A target tracking scaling policy.
	TargetTrackingScalingPolicyConfiguration interface{} `json:"targetTrackingScalingPolicyConfiguration" yaml:"targetTrackingScalingPolicyConfiguration"`
}

// Options to configure a cron expression.
//
// All fields are strings so you can use complex expressions. Absence of
// a field implies '*' or '?', whichever one is appropriate.
//
// Example:
//   var cluster cluster
//   loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &applicationLoadBalancedFargateServiceProps{
//   	cluster: cluster,
//   	memoryLimitMiB: jsii.Number(1024),
//   	desiredCount: jsii.Number(1),
//   	cpu: jsii.Number(512),
//   	taskImageOptions: &applicationLoadBalancedTaskImageOptions{
//   		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   })
//
//   scalableTarget := loadBalancedFargateService.service.autoScaleTaskCount(&enableScalingProps{
//   	minCapacity: jsii.Number(5),
//   	maxCapacity: jsii.Number(20),
//   })
//
//   scalableTarget.scaleOnSchedule(jsii.String("DaytimeScaleDown"), &scalingSchedule{
//   	schedule: appscaling.schedule.cron(&cronOptions{
//   		hour: jsii.String("8"),
//   		minute: jsii.String("0"),
//   	}),
//   	minCapacity: jsii.Number(1),
//   })
//
//   scalableTarget.scaleOnSchedule(jsii.String("EveningRushScaleUp"), &scalingSchedule{
//   	schedule: appscaling.*schedule.cron(&cronOptions{
//   		hour: jsii.String("20"),
//   		minute: jsii.String("0"),
//   	}),
//   	minCapacity: jsii.Number(10),
//   })
//
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html#CronExpressions
//
// Experimental.
type CronOptions struct {
	// The day of the month to run this rule at.
	// Experimental.
	Day *string `json:"day" yaml:"day"`
	// The hour to run this rule at.
	// Experimental.
	Hour *string `json:"hour" yaml:"hour"`
	// The minute to run this rule at.
	// Experimental.
	Minute *string `json:"minute" yaml:"minute"`
	// The month to run this rule at.
	// Experimental.
	Month *string `json:"month" yaml:"month"`
	// The day of the week to run this rule at.
	// Experimental.
	WeekDay *string `json:"weekDay" yaml:"weekDay"`
	// The year to run this rule at.
	// Experimental.
	Year *string `json:"year" yaml:"year"`
}

// Properties for enabling Application Auto Scaling.
//
// Example:
//   var cluster cluster
//   loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &applicationLoadBalancedFargateServiceProps{
//   	cluster: cluster,
//   	memoryLimitMiB: jsii.Number(1024),
//   	desiredCount: jsii.Number(1),
//   	cpu: jsii.Number(512),
//   	taskImageOptions: &applicationLoadBalancedTaskImageOptions{
//   		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   })
//
//   scalableTarget := loadBalancedFargateService.service.autoScaleTaskCount(&enableScalingProps{
//   	minCapacity: jsii.Number(1),
//   	maxCapacity: jsii.Number(20),
//   })
//
//   scalableTarget.scaleOnCpuUtilization(jsii.String("CpuScaling"), &cpuUtilizationScalingProps{
//   	targetUtilizationPercent: jsii.Number(50),
//   })
//
//   scalableTarget.scaleOnMemoryUtilization(jsii.String("MemoryScaling"), &memoryUtilizationScalingProps{
//   	targetUtilizationPercent: jsii.Number(50),
//   })
//
// Experimental.
type EnableScalingProps struct {
	// Maximum capacity to scale to.
	// Experimental.
	MaxCapacity *float64 `json:"maxCapacity" yaml:"maxCapacity"`
	// Minimum capacity to scale to.
	// Experimental.
	MinCapacity *float64 `json:"minCapacity" yaml:"minCapacity"`
}

// Experimental.
type IScalableTarget interface {
	awscdk.IResource
	// Experimental.
	ScalableTargetId() *string
}

// The jsii proxy for IScalableTarget
type jsiiProxy_IScalableTarget struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IScalableTarget) ScalableTargetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalableTargetId",
		&returns,
	)
	return returns
}

// How the scaling metric is going to be aggregated.
// Experimental.
type MetricAggregationType string

const (
	// Average.
	// Experimental.
	MetricAggregationType_AVERAGE MetricAggregationType = "AVERAGE"
	// Minimum.
	// Experimental.
	MetricAggregationType_MINIMUM MetricAggregationType = "MINIMUM"
	// Maximum.
	// Experimental.
	MetricAggregationType_MAXIMUM MetricAggregationType = "MAXIMUM"
)

// One of the predefined autoscaling metrics.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//   var code code
//
//
//   handler := lambda.NewFunction(this, jsii.String("MyFunction"), &functionProps{
//   	runtime: lambda.runtime_PYTHON_3_7(),
//   	handler: jsii.String("index.handler"),
//   	code: code,
//
//   	reservedConcurrentExecutions: jsii.Number(2),
//   })
//
//   fnVer := handler.addVersion(jsii.String("CDKLambdaVersion"), undefined, jsii.String("demo alias"), jsii.Number(10))
//
//   target := appscaling.NewScalableTarget(this, jsii.String("ScalableTarget"), &scalableTargetProps{
//   	serviceNamespace: appscaling.serviceNamespace_LAMBDA,
//   	maxCapacity: jsii.Number(100),
//   	minCapacity: jsii.Number(10),
//   	resourceId: fmt.Sprintf("function:%v:%v", handler.functionName, fnVer.version),
//   	scalableDimension: jsii.String("lambda:function:ProvisionedConcurrency"),
//   })
//
//   target.scaleToTrackMetric(jsii.String("PceTracking"), &basicTargetTrackingScalingPolicyProps{
//   	targetValue: jsii.Number(0.9),
//   	predefinedMetric: appscaling.predefinedMetric_LAMBDA_PROVISIONED_CONCURRENCY_UTILIZATION,
//   })
//
// Experimental.
type PredefinedMetric string

const (
	// DYNAMODB_READ_CAPACITY_UTILIZATIO.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PredefinedMetricSpecification.html
	//
	// Experimental.
	PredefinedMetric_DYNAMODB_READ_CAPACITY_UTILIZATION PredefinedMetric = "DYNAMODB_READ_CAPACITY_UTILIZATION"
	// DYNAMODB_WRITE_CAPACITY_UTILIZATION.
	//
	// Suffix `dummy` is necessary due to jsii bug (https://github.com/aws/jsii/issues/2782).
	// Duplicate values will be dropped, so this suffix is added as a workaround.
	// The value will be replaced when this enum is used.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PredefinedMetricSpecification.html
	//
	// Experimental.
	PredefinedMetric_DYNAMODB_WRITE_CAPACITY_UTILIZATION PredefinedMetric = "DYNAMODB_WRITE_CAPACITY_UTILIZATION"
	// DYANMODB_WRITE_CAPACITY_UTILIZATION.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PredefinedMetricSpecification.html
	//
	// Deprecated: use `PredefinedMetric.DYNAMODB_WRITE_CAPACITY_UTILIZATION`
	PredefinedMetric_DYANMODB_WRITE_CAPACITY_UTILIZATION PredefinedMetric = "DYANMODB_WRITE_CAPACITY_UTILIZATION"
	// ALB_REQUEST_COUNT_PER_TARGET.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PredefinedMetricSpecification.html
	//
	// Experimental.
	PredefinedMetric_ALB_REQUEST_COUNT_PER_TARGET PredefinedMetric = "ALB_REQUEST_COUNT_PER_TARGET"
	// RDS_READER_AVERAGE_CPU_UTILIZATION.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PredefinedMetricSpecification.html
	//
	// Experimental.
	PredefinedMetric_RDS_READER_AVERAGE_CPU_UTILIZATION PredefinedMetric = "RDS_READER_AVERAGE_CPU_UTILIZATION"
	// RDS_READER_AVERAGE_DATABASE_CONNECTIONS.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PredefinedMetricSpecification.html
	//
	// Experimental.
	PredefinedMetric_RDS_READER_AVERAGE_DATABASE_CONNECTIONS PredefinedMetric = "RDS_READER_AVERAGE_DATABASE_CONNECTIONS"
	// EC2_SPOT_FLEET_REQUEST_AVERAGE_CPU_UTILIZATION.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PredefinedMetricSpecification.html
	//
	// Experimental.
	PredefinedMetric_EC2_SPOT_FLEET_REQUEST_AVERAGE_CPU_UTILIZATION PredefinedMetric = "EC2_SPOT_FLEET_REQUEST_AVERAGE_CPU_UTILIZATION"
	// EC2_SPOT_FLEET_REQUEST_AVERAGE_NETWORK_IN.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PredefinedMetricSpecification.html
	//
	// Experimental.
	PredefinedMetric_EC2_SPOT_FLEET_REQUEST_AVERAGE_NETWORK_IN PredefinedMetric = "EC2_SPOT_FLEET_REQUEST_AVERAGE_NETWORK_IN"
	// EC2_SPOT_FLEET_REQUEST_AVERAGE_NETWORK_OUT.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PredefinedMetricSpecification.html
	//
	// Experimental.
	PredefinedMetric_EC2_SPOT_FLEET_REQUEST_AVERAGE_NETWORK_OUT PredefinedMetric = "EC2_SPOT_FLEET_REQUEST_AVERAGE_NETWORK_OUT"
	// SAGEMAKER_VARIANT_INVOCATIONS_PER_INSTANCE.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PredefinedMetricSpecification.html
	//
	// Experimental.
	PredefinedMetric_SAGEMAKER_VARIANT_INVOCATIONS_PER_INSTANCE PredefinedMetric = "SAGEMAKER_VARIANT_INVOCATIONS_PER_INSTANCE"
	// ECS_SERVICE_AVERAGE_CPU_UTILIZATION.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PredefinedMetricSpecification.html
	//
	// Experimental.
	PredefinedMetric_ECS_SERVICE_AVERAGE_CPU_UTILIZATION PredefinedMetric = "ECS_SERVICE_AVERAGE_CPU_UTILIZATION"
	// ECS_SERVICE_AVERAGE_MEMORY_UTILIZATION.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PredefinedMetricSpecification.html
	//
	// Experimental.
	PredefinedMetric_ECS_SERVICE_AVERAGE_MEMORY_UTILIZATION PredefinedMetric = "ECS_SERVICE_AVERAGE_MEMORY_UTILIZATION"
	// LAMBDA_PROVISIONED_CONCURRENCY_UTILIZATION.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/monitoring-metrics.html#monitoring-metrics-concurrency
	//
	// Experimental.
	PredefinedMetric_LAMBDA_PROVISIONED_CONCURRENCY_UTILIZATION PredefinedMetric = "LAMBDA_PROVISIONED_CONCURRENCY_UTILIZATION"
	// KAFKA_BROKER_STORAGE_UTILIZATION.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PredefinedMetricSpecification.html
	//
	// Experimental.
	PredefinedMetric_KAFKA_BROKER_STORAGE_UTILIZATION PredefinedMetric = "KAFKA_BROKER_STORAGE_UTILIZATION"
	// ELASTIC_CACHE_PRIMARY_ENGINE_CPU_UTILIZATION.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PredefinedMetricSpecification.html
	//
	// Experimental.
	PredefinedMetric_ELASTICACHE_PRIMARY_ENGINE_CPU_UTILIZATION PredefinedMetric = "ELASTICACHE_PRIMARY_ENGINE_CPU_UTILIZATION"
	// ELASTIC_CACHE_REPLICA_ENGINE_CPU_UTILIZATION.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PredefinedMetricSpecification.html
	//
	// Experimental.
	PredefinedMetric_ELASTICACHE_REPLICA_ENGINE_CPU_UTILIZATION PredefinedMetric = "ELASTICACHE_REPLICA_ENGINE_CPU_UTILIZATION"
	// ELASTIC_CACHE_REPLICA_ENGINE_CPU_UTILIZATION.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_PredefinedMetricSpecification.html
	//
	// Experimental.
	PredefinedMetric_ELASTICACHE_DATABASE_MEMORY_USAGE_COUNTED_FOR_EVICT_PERCENTAGE PredefinedMetric = "ELASTICACHE_DATABASE_MEMORY_USAGE_COUNTED_FOR_EVICT_PERCENTAGE"
)

// Define a scalable target.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//   var code code
//
//
//   handler := lambda.NewFunction(this, jsii.String("MyFunction"), &functionProps{
//   	runtime: lambda.runtime_PYTHON_3_7(),
//   	handler: jsii.String("index.handler"),
//   	code: code,
//
//   	reservedConcurrentExecutions: jsii.Number(2),
//   })
//
//   fnVer := handler.addVersion(jsii.String("CDKLambdaVersion"), undefined, jsii.String("demo alias"), jsii.Number(10))
//
//   target := appscaling.NewScalableTarget(this, jsii.String("ScalableTarget"), &scalableTargetProps{
//   	serviceNamespace: appscaling.serviceNamespace_LAMBDA,
//   	maxCapacity: jsii.Number(100),
//   	minCapacity: jsii.Number(10),
//   	resourceId: fmt.Sprintf("function:%v:%v", handler.functionName, fnVer.version),
//   	scalableDimension: jsii.String("lambda:function:ProvisionedConcurrency"),
//   })
//
//   target.scaleToTrackMetric(jsii.String("PceTracking"), &basicTargetTrackingScalingPolicyProps{
//   	targetValue: jsii.Number(0.9),
//   	predefinedMetric: appscaling.predefinedMetric_LAMBDA_PROVISIONED_CONCURRENCY_UTILIZATION,
//   })
//
// Experimental.
type ScalableTarget interface {
	awscdk.Resource
	IScalableTarget
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The role used to give AutoScaling permissions to your resource.
	// Experimental.
	Role() awsiam.IRole
	// ID of the Scalable Target.
	//
	// Example value: `service/ecsStack-MyECSCluster-AB12CDE3F4GH/ecsStack-MyECSService-AB12CDE3F4GH|ecs:service:DesiredCount|ecs`.
	// Experimental.
	ScalableTargetId() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Add a policy statement to the role's policy.
	// Experimental.
	AddToRolePolicy(statement awsiam.PolicyStatement)
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Scale out or in, in response to a metric.
	// Experimental.
	ScaleOnMetric(id *string, props *BasicStepScalingPolicyProps) StepScalingPolicy
	// Scale out or in based on time.
	// Experimental.
	ScaleOnSchedule(id *string, action *ScalingSchedule)
	// Scale out or in in order to keep a metric around a target value.
	// Experimental.
	ScaleToTrackMetric(id *string, props *BasicTargetTrackingScalingPolicyProps) TargetTrackingScalingPolicy
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for ScalableTarget
type jsiiProxy_ScalableTarget struct {
	internal.Type__awscdkResource
	jsiiProxy_IScalableTarget
}

func (j *jsiiProxy_ScalableTarget) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScalableTarget) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScalableTarget) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScalableTarget) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScalableTarget) ScalableTargetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalableTargetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ScalableTarget) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewScalableTarget(scope constructs.Construct, id *string, props *ScalableTargetProps) ScalableTarget {
	_init_.Initialize()

	j := jsiiProxy_ScalableTarget{}

	_jsii_.Create(
		"monocdk.aws_applicationautoscaling.ScalableTarget",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewScalableTarget_Override(s ScalableTarget, scope constructs.Construct, id *string, props *ScalableTargetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_applicationautoscaling.ScalableTarget",
		[]interface{}{scope, id, props},
		s,
	)
}

// Experimental.
func ScalableTarget_FromScalableTargetId(scope constructs.Construct, id *string, scalableTargetId *string) IScalableTarget {
	_init_.Initialize()

	var returns IScalableTarget

	_jsii_.StaticInvoke(
		"monocdk.aws_applicationautoscaling.ScalableTarget",
		"fromScalableTargetId",
		[]interface{}{scope, id, scalableTargetId},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func ScalableTarget_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_applicationautoscaling.ScalableTarget",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func ScalableTarget_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_applicationautoscaling.ScalableTarget",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScalableTarget) AddToRolePolicy(statement awsiam.PolicyStatement) {
	_jsii_.InvokeVoid(
		s,
		"addToRolePolicy",
		[]interface{}{statement},
	)
}

func (s *jsiiProxy_ScalableTarget) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_ScalableTarget) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScalableTarget) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScalableTarget) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScalableTarget) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ScalableTarget) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_ScalableTarget) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScalableTarget) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ScalableTarget) ScaleOnMetric(id *string, props *BasicStepScalingPolicyProps) StepScalingPolicy {
	var returns StepScalingPolicy

	_jsii_.Invoke(
		s,
		"scaleOnMetric",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScalableTarget) ScaleOnSchedule(id *string, action *ScalingSchedule) {
	_jsii_.InvokeVoid(
		s,
		"scaleOnSchedule",
		[]interface{}{id, action},
	)
}

func (s *jsiiProxy_ScalableTarget) ScaleToTrackMetric(id *string, props *BasicTargetTrackingScalingPolicyProps) TargetTrackingScalingPolicy {
	var returns TargetTrackingScalingPolicy

	_jsii_.Invoke(
		s,
		"scaleToTrackMetric",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScalableTarget) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_ScalableTarget) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ScalableTarget) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a scalable target.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//   var code code
//
//
//   handler := lambda.NewFunction(this, jsii.String("MyFunction"), &functionProps{
//   	runtime: lambda.runtime_PYTHON_3_7(),
//   	handler: jsii.String("index.handler"),
//   	code: code,
//
//   	reservedConcurrentExecutions: jsii.Number(2),
//   })
//
//   fnVer := handler.addVersion(jsii.String("CDKLambdaVersion"), undefined, jsii.String("demo alias"), jsii.Number(10))
//
//   target := appscaling.NewScalableTarget(this, jsii.String("ScalableTarget"), &scalableTargetProps{
//   	serviceNamespace: appscaling.serviceNamespace_LAMBDA,
//   	maxCapacity: jsii.Number(100),
//   	minCapacity: jsii.Number(10),
//   	resourceId: fmt.Sprintf("function:%v:%v", handler.functionName, fnVer.version),
//   	scalableDimension: jsii.String("lambda:function:ProvisionedConcurrency"),
//   })
//
//   target.scaleToTrackMetric(jsii.String("PceTracking"), &basicTargetTrackingScalingPolicyProps{
//   	targetValue: jsii.Number(0.9),
//   	predefinedMetric: appscaling.predefinedMetric_LAMBDA_PROVISIONED_CONCURRENCY_UTILIZATION,
//   })
//
// Experimental.
type ScalableTargetProps struct {
	// The maximum value that Application Auto Scaling can use to scale a target during a scaling activity.
	// Experimental.
	MaxCapacity *float64 `json:"maxCapacity" yaml:"maxCapacity"`
	// The minimum value that Application Auto Scaling can use to scale a target during a scaling activity.
	// Experimental.
	MinCapacity *float64 `json:"minCapacity" yaml:"minCapacity"`
	// The resource identifier to associate with this scalable target.
	//
	// This string consists of the resource type and unique identifier.
	//
	// Example value: `service/ecsStack-MyECSCluster-AB12CDE3F4GH/ecsStack-MyECSService-AB12CDE3F4GH`.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html
	//
	// Experimental.
	ResourceId *string `json:"resourceId" yaml:"resourceId"`
	// The scalable dimension that's associated with the scalable target.
	//
	// Specify the service namespace, resource type, and scaling property.
	//
	// Example value: `ecs:service:DesiredCount`.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_ScalingPolicy.html
	//
	// Experimental.
	ScalableDimension *string `json:"scalableDimension" yaml:"scalableDimension"`
	// The namespace of the AWS service that provides the resource or custom-resource for a resource provided by your own application or service.
	//
	// For valid AWS service namespace values, see the RegisterScalableTarget
	// action in the Application Auto Scaling API Reference.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html
	//
	// Experimental.
	ServiceNamespace ServiceNamespace `json:"serviceNamespace" yaml:"serviceNamespace"`
	// Role that allows Application Auto Scaling to modify your scalable target.
	// Experimental.
	Role awsiam.IRole `json:"role" yaml:"role"`
}

// A range of metric values in which to apply a certain scaling operation.
//
// Example:
//   import awscdk "github.com/aws/aws-cdk-go/awscdk"import applicationautoscaling "github.com/aws/aws-cdk-go/awscdk/aws_applicationautoscaling"
//   scalingInterval := &scalingInterval{
//   	change: jsii.Number(123),
//
//   	// the properties below are optional
//   	lower: jsii.Number(123),
//   	upper: jsii.Number(123),
//   }
//
// Experimental.
type ScalingInterval struct {
	// The capacity adjustment to apply in this interval.
	//
	// The number is interpreted differently based on AdjustmentType:
	//
	// - ChangeInCapacity: add the adjustment to the current capacity.
	//   The number can be positive or negative.
	// - PercentChangeInCapacity: add or remove the given percentage of the current
	//    capacity to itself. The number can be in the range [-100..100].
	// - ExactCapacity: set the capacity to this number. The number must
	//    be positive.
	// Experimental.
	Change *float64 `json:"change" yaml:"change"`
	// The lower bound of the interval.
	//
	// The scaling adjustment will be applied if the metric is higher than this value.
	// Experimental.
	Lower *float64 `json:"lower" yaml:"lower"`
	// The upper bound of the interval.
	//
	// The scaling adjustment will be applied if the metric is lower than this value.
	// Experimental.
	Upper *float64 `json:"upper" yaml:"upper"`
}

// A scheduled scaling action.
//
// Example:
//   var cluster cluster
//   loadBalancedFargateService := ecsPatterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &applicationLoadBalancedFargateServiceProps{
//   	cluster: cluster,
//   	memoryLimitMiB: jsii.Number(1024),
//   	desiredCount: jsii.Number(1),
//   	cpu: jsii.Number(512),
//   	taskImageOptions: &applicationLoadBalancedTaskImageOptions{
//   		image: ecs.containerImage.fromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	},
//   })
//
//   scalableTarget := loadBalancedFargateService.service.autoScaleTaskCount(&enableScalingProps{
//   	minCapacity: jsii.Number(5),
//   	maxCapacity: jsii.Number(20),
//   })
//
//   scalableTarget.scaleOnSchedule(jsii.String("DaytimeScaleDown"), &scalingSchedule{
//   	schedule: appscaling.schedule.cron(&cronOptions{
//   		hour: jsii.String("8"),
//   		minute: jsii.String("0"),
//   	}),
//   	minCapacity: jsii.Number(1),
//   })
//
//   scalableTarget.scaleOnSchedule(jsii.String("EveningRushScaleUp"), &scalingSchedule{
//   	schedule: appscaling.*schedule.cron(&cronOptions{
//   		hour: jsii.String("20"),
//   		minute: jsii.String("0"),
//   	}),
//   	minCapacity: jsii.Number(10),
//   })
//
// Experimental.
type ScalingSchedule struct {
	// When to perform this action.
	// Experimental.
	Schedule Schedule `json:"schedule" yaml:"schedule"`
	// When this scheduled action expires.
	// Experimental.
	EndTime *time.Time `json:"endTime" yaml:"endTime"`
	// The new maximum capacity.
	//
	// During the scheduled time, the current capacity is above the maximum
	// capacity, Application Auto Scaling scales in to the maximum capacity.
	//
	// At least one of maxCapacity and minCapacity must be supplied.
	// Experimental.
	MaxCapacity *float64 `json:"maxCapacity" yaml:"maxCapacity"`
	// The new minimum capacity.
	//
	// During the scheduled time, if the current capacity is below the minimum
	// capacity, Application Auto Scaling scales out to the minimum capacity.
	//
	// At least one of maxCapacity and minCapacity must be supplied.
	// Experimental.
	MinCapacity *float64 `json:"minCapacity" yaml:"minCapacity"`
	// When this scheduled action becomes active.
	// Experimental.
	StartTime *time.Time `json:"startTime" yaml:"startTime"`
}

// Schedule for scheduled scaling actions.
//
// Example:
//   import autoscaling "github.com/aws/aws-cdk-go/awscdk"
//
//   var fn function
//   alias := lambda.NewAlias(this, jsii.String("Alias"), &aliasProps{
//   	aliasName: jsii.String("prod"),
//   	version: fn.latestVersion,
//   })
//
//   // Create AutoScaling target
//   as := alias.addAutoScaling(&autoScalingOptions{
//   	maxCapacity: jsii.Number(50),
//   })
//
//   // Configure Target Tracking
//   as.scaleOnUtilization(&utilizationScalingOptions{
//   	utilizationTarget: jsii.Number(0.5),
//   })
//
//   // Configure Scheduled Scaling
//   as.scaleOnSchedule(jsii.String("ScaleUpInTheMorning"), &scalingSchedule{
//   	schedule: autoscaling.schedule.cron(&cronOptions{
//   		hour: jsii.String("8"),
//   		minute: jsii.String("0"),
//   	}),
//   	minCapacity: jsii.Number(20),
//   })
//
// Experimental.
type Schedule interface {
	// Retrieve the expression for this schedule.
	// Experimental.
	ExpressionString() *string
}

// The jsii proxy struct for Schedule
type jsiiProxy_Schedule struct {
	_ byte // padding
}

func (j *jsiiProxy_Schedule) ExpressionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressionString",
		&returns,
	)
	return returns
}


// Experimental.
func NewSchedule_Override(s Schedule) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_applicationautoscaling.Schedule",
		nil, // no parameters
		s,
	)
}

// Construct a Schedule from a moment in time.
// Experimental.
func Schedule_At(moment *time.Time) Schedule {
	_init_.Initialize()

	var returns Schedule

	_jsii_.StaticInvoke(
		"monocdk.aws_applicationautoscaling.Schedule",
		"at",
		[]interface{}{moment},
		&returns,
	)

	return returns
}

// Create a schedule from a set of cron fields.
// Experimental.
func Schedule_Cron(options *CronOptions) Schedule {
	_init_.Initialize()

	var returns Schedule

	_jsii_.StaticInvoke(
		"monocdk.aws_applicationautoscaling.Schedule",
		"cron",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Construct a schedule from a literal schedule expression.
// Experimental.
func Schedule_Expression(expression *string) Schedule {
	_init_.Initialize()

	var returns Schedule

	_jsii_.StaticInvoke(
		"monocdk.aws_applicationautoscaling.Schedule",
		"expression",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

// Construct a schedule from an interval and a time unit.
// Experimental.
func Schedule_Rate(duration awscdk.Duration) Schedule {
	_init_.Initialize()

	var returns Schedule

	_jsii_.StaticInvoke(
		"monocdk.aws_applicationautoscaling.Schedule",
		"rate",
		[]interface{}{duration},
		&returns,
	)

	return returns
}

// The service that supports Application AutoScaling.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//   var code code
//
//
//   handler := lambda.NewFunction(this, jsii.String("MyFunction"), &functionProps{
//   	runtime: lambda.runtime_PYTHON_3_7(),
//   	handler: jsii.String("index.handler"),
//   	code: code,
//
//   	reservedConcurrentExecutions: jsii.Number(2),
//   })
//
//   fnVer := handler.addVersion(jsii.String("CDKLambdaVersion"), undefined, jsii.String("demo alias"), jsii.Number(10))
//
//   target := appscaling.NewScalableTarget(this, jsii.String("ScalableTarget"), &scalableTargetProps{
//   	serviceNamespace: appscaling.serviceNamespace_LAMBDA,
//   	maxCapacity: jsii.Number(100),
//   	minCapacity: jsii.Number(10),
//   	resourceId: fmt.Sprintf("function:%v:%v", handler.functionName, fnVer.version),
//   	scalableDimension: jsii.String("lambda:function:ProvisionedConcurrency"),
//   })
//
//   target.scaleToTrackMetric(jsii.String("PceTracking"), &basicTargetTrackingScalingPolicyProps{
//   	targetValue: jsii.Number(0.9),
//   	predefinedMetric: appscaling.predefinedMetric_LAMBDA_PROVISIONED_CONCURRENCY_UTILIZATION,
//   })
//
// Experimental.
type ServiceNamespace string

const (
	// Elastic Container Service.
	// Experimental.
	ServiceNamespace_ECS ServiceNamespace = "ECS"
	// Elastic Map Reduce.
	// Experimental.
	ServiceNamespace_ELASTIC_MAP_REDUCE ServiceNamespace = "ELASTIC_MAP_REDUCE"
	// Elastic Compute Cloud.
	// Experimental.
	ServiceNamespace_EC2 ServiceNamespace = "EC2"
	// App Stream.
	// Experimental.
	ServiceNamespace_APPSTREAM ServiceNamespace = "APPSTREAM"
	// Dynamo DB.
	// Experimental.
	ServiceNamespace_DYNAMODB ServiceNamespace = "DYNAMODB"
	// Relational Database Service.
	// Experimental.
	ServiceNamespace_RDS ServiceNamespace = "RDS"
	// SageMaker.
	// Experimental.
	ServiceNamespace_SAGEMAKER ServiceNamespace = "SAGEMAKER"
	// Custom Resource.
	// Experimental.
	ServiceNamespace_CUSTOM_RESOURCE ServiceNamespace = "CUSTOM_RESOURCE"
	// Lambda.
	// Experimental.
	ServiceNamespace_LAMBDA ServiceNamespace = "LAMBDA"
	// Comprehend.
	// Experimental.
	ServiceNamespace_COMPREHEND ServiceNamespace = "COMPREHEND"
	// Kafka.
	// Experimental.
	ServiceNamespace_KAFKA ServiceNamespace = "KAFKA"
	// ElastiCache.
	// Experimental.
	ServiceNamespace_ELASTICACHE ServiceNamespace = "ELASTICACHE"
)

// Define a step scaling action.
//
// This kind of scaling policy adjusts the target capacity in configurable
// steps. The size of the step is configurable based on the metric's distance
// to its alarm threshold.
//
// This Action must be used as the target of a CloudWatch alarm to take effect.
//
// Example:
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"import awscdk "github.com/aws/aws-cdk-go/awscdk"import applicationautoscaling "github.com/aws/aws-cdk-go/awscdk/aws_applicationautoscaling"
//
//   var duration duration
//   var scalableTarget scalableTarget
//   stepScalingAction := applicationautoscaling.NewStepScalingAction(this, jsii.String("MyStepScalingAction"), &stepScalingActionProps{
//   	scalingTarget: scalableTarget,
//
//   	// the properties below are optional
//   	adjustmentType: applicationautoscaling.adjustmentType_CHANGE_IN_CAPACITY,
//   	cooldown: duration,
//   	metricAggregationType: applicationautoscaling.metricAggregationType_AVERAGE,
//   	minAdjustmentMagnitude: jsii.Number(123),
//   	policyName: jsii.String("policyName"),
//   })
//
// Experimental.
type StepScalingAction interface {
	awscdk.Construct
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// ARN of the scaling policy.
	// Experimental.
	ScalingPolicyArn() *string
	// Add an adjusment interval to the ScalingAction.
	// Experimental.
	AddAdjustment(adjustment *AdjustmentTier)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for StepScalingAction
type jsiiProxy_StepScalingAction struct {
	internal.Type__awscdkConstruct
}

func (j *jsiiProxy_StepScalingAction) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepScalingAction) ScalingPolicyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingPolicyArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewStepScalingAction(scope constructs.Construct, id *string, props *StepScalingActionProps) StepScalingAction {
	_init_.Initialize()

	j := jsiiProxy_StepScalingAction{}

	_jsii_.Create(
		"monocdk.aws_applicationautoscaling.StepScalingAction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewStepScalingAction_Override(s StepScalingAction, scope constructs.Construct, id *string, props *StepScalingActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_applicationautoscaling.StepScalingAction",
		[]interface{}{scope, id, props},
		s,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func StepScalingAction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_applicationautoscaling.StepScalingAction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepScalingAction) AddAdjustment(adjustment *AdjustmentTier) {
	_jsii_.InvokeVoid(
		s,
		"addAdjustment",
		[]interface{}{adjustment},
	)
}

func (s *jsiiProxy_StepScalingAction) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StepScalingAction) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_StepScalingAction) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepScalingAction) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StepScalingAction) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_StepScalingAction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepScalingAction) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a scaling policy.
//
// Example:
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"import awscdk "github.com/aws/aws-cdk-go/awscdk"import applicationautoscaling "github.com/aws/aws-cdk-go/awscdk/aws_applicationautoscaling"
//
//   var duration duration
//   var scalableTarget scalableTarget
//   stepScalingActionProps := &stepScalingActionProps{
//   	scalingTarget: scalableTarget,
//
//   	// the properties below are optional
//   	adjustmentType: applicationautoscaling.adjustmentType_CHANGE_IN_CAPACITY,
//   	cooldown: duration,
//   	metricAggregationType: applicationautoscaling.metricAggregationType_AVERAGE,
//   	minAdjustmentMagnitude: jsii.Number(123),
//   	policyName: jsii.String("policyName"),
//   }
//
// Experimental.
type StepScalingActionProps struct {
	// The scalable target.
	// Experimental.
	ScalingTarget IScalableTarget `json:"scalingTarget" yaml:"scalingTarget"`
	// How the adjustment numbers are interpreted.
	// Experimental.
	AdjustmentType AdjustmentType `json:"adjustmentType" yaml:"adjustmentType"`
	// Grace period after scaling activity.
	//
	// For scale out policies, multiple scale outs during the cooldown period are
	// squashed so that only the biggest scale out happens.
	//
	// For scale in policies, subsequent scale ins during the cooldown period are
	// ignored.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_StepScalingPolicyConfiguration.html
	//
	// Experimental.
	Cooldown awscdk.Duration `json:"cooldown" yaml:"cooldown"`
	// The aggregation type for the CloudWatch metrics.
	// Experimental.
	MetricAggregationType MetricAggregationType `json:"metricAggregationType" yaml:"metricAggregationType"`
	// Minimum absolute number to adjust capacity with as result of percentage scaling.
	//
	// Only when using AdjustmentType = PercentChangeInCapacity, this number controls
	// the minimum absolute effect size.
	// Experimental.
	MinAdjustmentMagnitude *float64 `json:"minAdjustmentMagnitude" yaml:"minAdjustmentMagnitude"`
	// A name for the scaling policy.
	// Experimental.
	PolicyName *string `json:"policyName" yaml:"policyName"`
}

// Define a scaling strategy which scales depending on absolute values of some metric.
//
// You can specify the scaling behavior for various values of the metric.
//
// Implemented using one or more CloudWatch alarms and Step Scaling Policies.
//
// Example:
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"import awscdk "github.com/aws/aws-cdk-go/awscdk"import applicationautoscaling "github.com/aws/aws-cdk-go/awscdk/aws_applicationautoscaling"import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloudwatch "github.com/aws/aws-cdk-go/awscdk/aws_cloudwatch"
//
//   var duration duration
//   var metric metric
//   var scalableTarget scalableTarget
//   stepScalingPolicy := applicationautoscaling.NewStepScalingPolicy(this, jsii.String("MyStepScalingPolicy"), &stepScalingPolicyProps{
//   	metric: metric,
//   	scalingSteps: []scalingInterval{
//   		&scalingInterval{
//   			change: jsii.Number(123),
//
//   			// the properties below are optional
//   			lower: jsii.Number(123),
//   			upper: jsii.Number(123),
//   		},
//   	},
//   	scalingTarget: scalableTarget,
//
//   	// the properties below are optional
//   	adjustmentType: applicationautoscaling.adjustmentType_CHANGE_IN_CAPACITY,
//   	cooldown: duration,
//   	datapointsToAlarm: jsii.Number(123),
//   	evaluationPeriods: jsii.Number(123),
//   	metricAggregationType: applicationautoscaling.metricAggregationType_AVERAGE,
//   	minAdjustmentMagnitude: jsii.Number(123),
//   })
//
// Experimental.
type StepScalingPolicy interface {
	awscdk.Construct
	// Experimental.
	LowerAction() StepScalingAction
	// Experimental.
	LowerAlarm() awscloudwatch.Alarm
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Experimental.
	UpperAction() StepScalingAction
	// Experimental.
	UpperAlarm() awscloudwatch.Alarm
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for StepScalingPolicy
type jsiiProxy_StepScalingPolicy struct {
	internal.Type__awscdkConstruct
}

func (j *jsiiProxy_StepScalingPolicy) LowerAction() StepScalingAction {
	var returns StepScalingAction
	_jsii_.Get(
		j,
		"lowerAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepScalingPolicy) LowerAlarm() awscloudwatch.Alarm {
	var returns awscloudwatch.Alarm
	_jsii_.Get(
		j,
		"lowerAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepScalingPolicy) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepScalingPolicy) UpperAction() StepScalingAction {
	var returns StepScalingAction
	_jsii_.Get(
		j,
		"upperAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepScalingPolicy) UpperAlarm() awscloudwatch.Alarm {
	var returns awscloudwatch.Alarm
	_jsii_.Get(
		j,
		"upperAlarm",
		&returns,
	)
	return returns
}


// Experimental.
func NewStepScalingPolicy(scope constructs.Construct, id *string, props *StepScalingPolicyProps) StepScalingPolicy {
	_init_.Initialize()

	j := jsiiProxy_StepScalingPolicy{}

	_jsii_.Create(
		"monocdk.aws_applicationautoscaling.StepScalingPolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewStepScalingPolicy_Override(s StepScalingPolicy, scope constructs.Construct, id *string, props *StepScalingPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_applicationautoscaling.StepScalingPolicy",
		[]interface{}{scope, id, props},
		s,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func StepScalingPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_applicationautoscaling.StepScalingPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepScalingPolicy) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StepScalingPolicy) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_StepScalingPolicy) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepScalingPolicy) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StepScalingPolicy) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_StepScalingPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepScalingPolicy) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Example:
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"import awscdk "github.com/aws/aws-cdk-go/awscdk"import applicationautoscaling "github.com/aws/aws-cdk-go/awscdk/aws_applicationautoscaling"import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloudwatch "github.com/aws/aws-cdk-go/awscdk/aws_cloudwatch"
//
//   var duration duration
//   var metric metric
//   var scalableTarget scalableTarget
//   stepScalingPolicyProps := &stepScalingPolicyProps{
//   	metric: metric,
//   	scalingSteps: []scalingInterval{
//   		&scalingInterval{
//   			change: jsii.Number(123),
//
//   			// the properties below are optional
//   			lower: jsii.Number(123),
//   			upper: jsii.Number(123),
//   		},
//   	},
//   	scalingTarget: scalableTarget,
//
//   	// the properties below are optional
//   	adjustmentType: applicationautoscaling.adjustmentType_CHANGE_IN_CAPACITY,
//   	cooldown: duration,
//   	datapointsToAlarm: jsii.Number(123),
//   	evaluationPeriods: jsii.Number(123),
//   	metricAggregationType: applicationautoscaling.metricAggregationType_AVERAGE,
//   	minAdjustmentMagnitude: jsii.Number(123),
//   }
//
// Experimental.
type StepScalingPolicyProps struct {
	// Metric to scale on.
	// Experimental.
	Metric awscloudwatch.IMetric `json:"metric" yaml:"metric"`
	// The intervals for scaling.
	//
	// Maps a range of metric values to a particular scaling behavior.
	// Experimental.
	ScalingSteps *[]*ScalingInterval `json:"scalingSteps" yaml:"scalingSteps"`
	// How the adjustment numbers inside 'intervals' are interpreted.
	// Experimental.
	AdjustmentType AdjustmentType `json:"adjustmentType" yaml:"adjustmentType"`
	// Grace period after scaling activity.
	//
	// Subsequent scale outs during the cooldown period are squashed so that only
	// the biggest scale out happens.
	//
	// Subsequent scale ins during the cooldown period are ignored.
	// See: https://docs.aws.amazon.com/autoscaling/application/APIReference/API_StepScalingPolicyConfiguration.html
	//
	// Experimental.
	Cooldown awscdk.Duration `json:"cooldown" yaml:"cooldown"`
	// The number of data points out of the evaluation periods that must be breaching to trigger a scaling action.
	//
	// Creates an "M out of N" alarm, where this property is the M and the value set for
	// `evaluationPeriods` is the N value.
	//
	// Only has meaning if `evaluationPeriods != 1`.
	// Experimental.
	DatapointsToAlarm *float64 `json:"datapointsToAlarm" yaml:"datapointsToAlarm"`
	// How many evaluation periods of the metric to wait before triggering a scaling action.
	//
	// Raising this value can be used to smooth out the metric, at the expense
	// of slower response times.
	//
	// If `datapointsToAlarm` is not set, then all data points in the evaluation period
	// must meet the criteria to trigger a scaling action.
	// Experimental.
	EvaluationPeriods *float64 `json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// Aggregation to apply to all data points over the evaluation periods.
	//
	// Only has meaning if `evaluationPeriods != 1`.
	// Experimental.
	MetricAggregationType MetricAggregationType `json:"metricAggregationType" yaml:"metricAggregationType"`
	// Minimum absolute number to adjust capacity with as result of percentage scaling.
	//
	// Only when using AdjustmentType = PercentChangeInCapacity, this number controls
	// the minimum absolute effect size.
	// Experimental.
	MinAdjustmentMagnitude *float64 `json:"minAdjustmentMagnitude" yaml:"minAdjustmentMagnitude"`
	// The scaling target.
	// Experimental.
	ScalingTarget IScalableTarget `json:"scalingTarget" yaml:"scalingTarget"`
}

// Example:
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"import awscdk "github.com/aws/aws-cdk-go/awscdk"import applicationautoscaling "github.com/aws/aws-cdk-go/awscdk/aws_applicationautoscaling"import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloudwatch "github.com/aws/aws-cdk-go/awscdk/aws_cloudwatch"
//
//   var duration duration
//   var metric metric
//   var scalableTarget scalableTarget
//   targetTrackingScalingPolicy := applicationautoscaling.NewTargetTrackingScalingPolicy(this, jsii.String("MyTargetTrackingScalingPolicy"), &targetTrackingScalingPolicyProps{
//   	scalingTarget: scalableTarget,
//   	targetValue: jsii.Number(123),
//
//   	// the properties below are optional
//   	customMetric: metric,
//   	disableScaleIn: jsii.Boolean(false),
//   	policyName: jsii.String("policyName"),
//   	predefinedMetric: applicationautoscaling.predefinedMetric_DYNAMODB_READ_CAPACITY_UTILIZATION,
//   	resourceLabel: jsii.String("resourceLabel"),
//   	scaleInCooldown: duration,
//   	scaleOutCooldown: duration,
//   })
//
// Experimental.
type TargetTrackingScalingPolicy interface {
	awscdk.Construct
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// ARN of the scaling policy.
	// Experimental.
	ScalingPolicyArn() *string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for TargetTrackingScalingPolicy
type jsiiProxy_TargetTrackingScalingPolicy struct {
	internal.Type__awscdkConstruct
}

func (j *jsiiProxy_TargetTrackingScalingPolicy) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TargetTrackingScalingPolicy) ScalingPolicyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingPolicyArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewTargetTrackingScalingPolicy(scope constructs.Construct, id *string, props *TargetTrackingScalingPolicyProps) TargetTrackingScalingPolicy {
	_init_.Initialize()

	j := jsiiProxy_TargetTrackingScalingPolicy{}

	_jsii_.Create(
		"monocdk.aws_applicationautoscaling.TargetTrackingScalingPolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewTargetTrackingScalingPolicy_Override(t TargetTrackingScalingPolicy, scope constructs.Construct, id *string, props *TargetTrackingScalingPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_applicationautoscaling.TargetTrackingScalingPolicy",
		[]interface{}{scope, id, props},
		t,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func TargetTrackingScalingPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_applicationautoscaling.TargetTrackingScalingPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TargetTrackingScalingPolicy) OnPrepare() {
	_jsii_.InvokeVoid(
		t,
		"onPrepare",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TargetTrackingScalingPolicy) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		t,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (t *jsiiProxy_TargetTrackingScalingPolicy) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TargetTrackingScalingPolicy) Prepare() {
	_jsii_.InvokeVoid(
		t,
		"prepare",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TargetTrackingScalingPolicy) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		t,
		"synthesize",
		[]interface{}{session},
	)
}

func (t *jsiiProxy_TargetTrackingScalingPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TargetTrackingScalingPolicy) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties for a concrete TargetTrackingPolicy.
//
// Adds the scalingTarget.
//
// Example:
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"import awscdk "github.com/aws/aws-cdk-go/awscdk"import applicationautoscaling "github.com/aws/aws-cdk-go/awscdk/aws_applicationautoscaling"import awscdk "github.com/aws/aws-cdk-go/awscdk"import cloudwatch "github.com/aws/aws-cdk-go/awscdk/aws_cloudwatch"
//
//   var duration duration
//   var metric metric
//   var scalableTarget scalableTarget
//   targetTrackingScalingPolicyProps := &targetTrackingScalingPolicyProps{
//   	scalingTarget: scalableTarget,
//   	targetValue: jsii.Number(123),
//
//   	// the properties below are optional
//   	customMetric: metric,
//   	disableScaleIn: jsii.Boolean(false),
//   	policyName: jsii.String("policyName"),
//   	predefinedMetric: applicationautoscaling.predefinedMetric_DYNAMODB_READ_CAPACITY_UTILIZATION,
//   	resourceLabel: jsii.String("resourceLabel"),
//   	scaleInCooldown: duration,
//   	scaleOutCooldown: duration,
//   }
//
// Experimental.
type TargetTrackingScalingPolicyProps struct {
	// Indicates whether scale in by the target tracking policy is disabled.
	//
	// If the value is true, scale in is disabled and the target tracking policy
	// won't remove capacity from the scalable resource. Otherwise, scale in is
	// enabled and the target tracking policy can remove capacity from the
	// scalable resource.
	// Experimental.
	DisableScaleIn *bool `json:"disableScaleIn" yaml:"disableScaleIn"`
	// A name for the scaling policy.
	// Experimental.
	PolicyName *string `json:"policyName" yaml:"policyName"`
	// Period after a scale in activity completes before another scale in activity can start.
	// Experimental.
	ScaleInCooldown awscdk.Duration `json:"scaleInCooldown" yaml:"scaleInCooldown"`
	// Period after a scale out activity completes before another scale out activity can start.
	// Experimental.
	ScaleOutCooldown awscdk.Duration `json:"scaleOutCooldown" yaml:"scaleOutCooldown"`
	// The target value for the metric.
	// Experimental.
	TargetValue *float64 `json:"targetValue" yaml:"targetValue"`
	// A custom metric for application autoscaling.
	//
	// The metric must track utilization. Scaling out will happen if the metric is higher than
	// the target value, scaling in will happen in the metric is lower than the target value.
	//
	// Exactly one of customMetric or predefinedMetric must be specified.
	// Experimental.
	CustomMetric awscloudwatch.IMetric `json:"customMetric" yaml:"customMetric"`
	// A predefined metric for application autoscaling.
	//
	// The metric must track utilization. Scaling out will happen if the metric is higher than
	// the target value, scaling in will happen in the metric is lower than the target value.
	//
	// Exactly one of customMetric or predefinedMetric must be specified.
	// Experimental.
	PredefinedMetric PredefinedMetric `json:"predefinedMetric" yaml:"predefinedMetric"`
	// Identify the resource associated with the metric type.
	//
	// Only used for predefined metric ALBRequestCountPerTarget.
	//
	// Example value: `app/<load-balancer-name>/<load-balancer-id>/targetgroup/<target-group-name>/<target-group-id>`.
	// Experimental.
	ResourceLabel *string `json:"resourceLabel" yaml:"resourceLabel"`
	// Experimental.
	ScalingTarget IScalableTarget `json:"scalingTarget" yaml:"scalingTarget"`
}

