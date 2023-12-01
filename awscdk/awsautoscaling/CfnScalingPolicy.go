package awsautoscaling

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::AutoScaling::ScalingPolicy` resource specifies an Amazon EC2 Auto Scaling scaling policy so that the Auto Scaling group can scale the number of instances available for your application.
//
// For more information about using scaling policies to scale your Auto Scaling group automatically, see [Dynamic scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-scale-based-on-demand.html) and [Predictive scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-predictive-scaling.html) in the *Amazon EC2 Auto Scaling User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnScalingPolicy := awscdk.Aws_autoscaling.NewCfnScalingPolicy(this, jsii.String("MyCfnScalingPolicy"), &CfnScalingPolicyProps{
//   	AutoScalingGroupName: jsii.String("autoScalingGroupName"),
//
//   	// the properties below are optional
//   	AdjustmentType: jsii.String("adjustmentType"),
//   	Cooldown: jsii.String("cooldown"),
//   	EstimatedInstanceWarmup: jsii.Number(123),
//   	MetricAggregationType: jsii.String("metricAggregationType"),
//   	MinAdjustmentMagnitude: jsii.Number(123),
//   	PolicyType: jsii.String("policyType"),
//   	PredictiveScalingConfiguration: &PredictiveScalingConfigurationProperty{
//   		MetricSpecifications: []interface{}{
//   			&PredictiveScalingMetricSpecificationProperty{
//   				TargetValue: jsii.Number(123),
//
//   				// the properties below are optional
//   				CustomizedCapacityMetricSpecification: &PredictiveScalingCustomizedCapacityMetricProperty{
//   					MetricDataQueries: []interface{}{
//   						&MetricDataQueryProperty{
//   							Id: jsii.String("id"),
//
//   							// the properties below are optional
//   							Expression: jsii.String("expression"),
//   							Label: jsii.String("label"),
//   							MetricStat: &MetricStatProperty{
//   								Metric: &MetricProperty{
//   									MetricName: jsii.String("metricName"),
//   									Namespace: jsii.String("namespace"),
//
//   									// the properties below are optional
//   									Dimensions: []interface{}{
//   										&MetricDimensionProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   								Stat: jsii.String("stat"),
//
//   								// the properties below are optional
//   								Unit: jsii.String("unit"),
//   							},
//   							ReturnData: jsii.Boolean(false),
//   						},
//   					},
//   				},
//   				CustomizedLoadMetricSpecification: &PredictiveScalingCustomizedLoadMetricProperty{
//   					MetricDataQueries: []interface{}{
//   						&MetricDataQueryProperty{
//   							Id: jsii.String("id"),
//
//   							// the properties below are optional
//   							Expression: jsii.String("expression"),
//   							Label: jsii.String("label"),
//   							MetricStat: &MetricStatProperty{
//   								Metric: &MetricProperty{
//   									MetricName: jsii.String("metricName"),
//   									Namespace: jsii.String("namespace"),
//
//   									// the properties below are optional
//   									Dimensions: []interface{}{
//   										&MetricDimensionProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   								Stat: jsii.String("stat"),
//
//   								// the properties below are optional
//   								Unit: jsii.String("unit"),
//   							},
//   							ReturnData: jsii.Boolean(false),
//   						},
//   					},
//   				},
//   				CustomizedScalingMetricSpecification: &PredictiveScalingCustomizedScalingMetricProperty{
//   					MetricDataQueries: []interface{}{
//   						&MetricDataQueryProperty{
//   							Id: jsii.String("id"),
//
//   							// the properties below are optional
//   							Expression: jsii.String("expression"),
//   							Label: jsii.String("label"),
//   							MetricStat: &MetricStatProperty{
//   								Metric: &MetricProperty{
//   									MetricName: jsii.String("metricName"),
//   									Namespace: jsii.String("namespace"),
//
//   									// the properties below are optional
//   									Dimensions: []interface{}{
//   										&MetricDimensionProperty{
//   											Name: jsii.String("name"),
//   											Value: jsii.String("value"),
//   										},
//   									},
//   								},
//   								Stat: jsii.String("stat"),
//
//   								// the properties below are optional
//   								Unit: jsii.String("unit"),
//   							},
//   							ReturnData: jsii.Boolean(false),
//   						},
//   					},
//   				},
//   				PredefinedLoadMetricSpecification: &PredictiveScalingPredefinedLoadMetricProperty{
//   					PredefinedMetricType: jsii.String("predefinedMetricType"),
//
//   					// the properties below are optional
//   					ResourceLabel: jsii.String("resourceLabel"),
//   				},
//   				PredefinedMetricPairSpecification: &PredictiveScalingPredefinedMetricPairProperty{
//   					PredefinedMetricType: jsii.String("predefinedMetricType"),
//
//   					// the properties below are optional
//   					ResourceLabel: jsii.String("resourceLabel"),
//   				},
//   				PredefinedScalingMetricSpecification: &PredictiveScalingPredefinedScalingMetricProperty{
//   					PredefinedMetricType: jsii.String("predefinedMetricType"),
//
//   					// the properties below are optional
//   					ResourceLabel: jsii.String("resourceLabel"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		MaxCapacityBreachBehavior: jsii.String("maxCapacityBreachBehavior"),
//   		MaxCapacityBuffer: jsii.Number(123),
//   		Mode: jsii.String("mode"),
//   		SchedulingBufferTime: jsii.Number(123),
//   	},
//   	ScalingAdjustment: jsii.Number(123),
//   	StepAdjustments: []interface{}{
//   		&StepAdjustmentProperty{
//   			ScalingAdjustment: jsii.Number(123),
//
//   			// the properties below are optional
//   			MetricIntervalLowerBound: jsii.Number(123),
//   			MetricIntervalUpperBound: jsii.Number(123),
//   		},
//   	},
//   	TargetTrackingConfiguration: &TargetTrackingConfigurationProperty{
//   		TargetValue: jsii.Number(123),
//
//   		// the properties below are optional
//   		CustomizedMetricSpecification: &CustomizedMetricSpecificationProperty{
//   			MetricName: jsii.String("metricName"),
//   			Namespace: jsii.String("namespace"),
//   			Statistic: jsii.String("statistic"),
//
//   			// the properties below are optional
//   			Dimensions: []interface{}{
//   				&MetricDimensionProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Unit: jsii.String("unit"),
//   		},
//   		DisableScaleIn: jsii.Boolean(false),
//   		PredefinedMetricSpecification: &PredefinedMetricSpecificationProperty{
//   			PredefinedMetricType: jsii.String("predefinedMetricType"),
//
//   			// the properties below are optional
//   			ResourceLabel: jsii.String("resourceLabel"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-scalingpolicy.html
//
type CfnScalingPolicy interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// Specifies how the scaling adjustment is interpreted (for example, an absolute number or a percentage).
	AdjustmentType() *string
	SetAdjustmentType(val *string)
	// Returns the ARN of a scaling policy.
	AttrArn() *string
	// Returns the name of a scaling policy.
	AttrPolicyName() *string
	// The name of the Auto Scaling group.
	AutoScalingGroupName() *string
	SetAutoScalingGroupName(val *string)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// A cooldown period, in seconds, that applies to a specific simple scaling policy.
	Cooldown() *string
	SetCooldown(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// *Not needed if the default instance warmup is defined for the group.*.
	EstimatedInstanceWarmup() *float64
	SetEstimatedInstanceWarmup(val *float64)
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The aggregation type for the CloudWatch metrics.
	MetricAggregationType() *string
	SetMetricAggregationType(val *string)
	// The minimum value to scale by when the adjustment type is `PercentChangeInCapacity` .
	MinAdjustmentMagnitude() *float64
	SetMinAdjustmentMagnitude(val *float64)
	// The tree node.
	Node() constructs.Node
	// One of the following policy types:.
	PolicyType() *string
	SetPolicyType(val *string)
	// A predictive scaling policy.
	//
	// Provides support for predefined and custom metrics.
	PredictiveScalingConfiguration() interface{}
	SetPredictiveScalingConfiguration(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The amount by which to scale, based on the specified adjustment type.
	ScalingAdjustment() *float64
	SetScalingAdjustment(val *float64)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// A set of adjustments that enable you to scale based on the size of the alarm breach.
	StepAdjustments() interface{}
	SetStepAdjustments(val interface{})
	// A target tracking scaling policy.
	//
	// Provides support for predefined or custom metrics.
	TargetTrackingConfiguration() interface{}
	SetTargetTrackingConfiguration(val interface{})
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
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
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnScalingPolicy
type jsiiProxy_CfnScalingPolicy struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnScalingPolicy) AdjustmentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adjustmentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) AttrPolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrPolicyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) AutoScalingGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingGroupName",
		&returns,
	)
	return returns
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

func (j *jsiiProxy_CfnScalingPolicy) Cooldown() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cooldown",
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

func (j *jsiiProxy_CfnScalingPolicy) EstimatedInstanceWarmup() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"estimatedInstanceWarmup",
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

func (j *jsiiProxy_CfnScalingPolicy) MetricAggregationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricAggregationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) MinAdjustmentMagnitude() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minAdjustmentMagnitude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
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

func (j *jsiiProxy_CfnScalingPolicy) PredictiveScalingConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"predictiveScalingConfiguration",
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

func (j *jsiiProxy_CfnScalingPolicy) ScalingAdjustment() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scalingAdjustment",
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

func (j *jsiiProxy_CfnScalingPolicy) StepAdjustments() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stepAdjustments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPolicy) TargetTrackingConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetTrackingConfiguration",
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

func (j *jsiiProxy_CfnScalingPolicy) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


func NewCfnScalingPolicy(scope constructs.Construct, id *string, props *CfnScalingPolicyProps) CfnScalingPolicy {
	_init_.Initialize()

	if err := validateNewCfnScalingPolicyParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnScalingPolicy{}

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.CfnScalingPolicy",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnScalingPolicy_Override(c CfnScalingPolicy, scope constructs.Construct, id *string, props *CfnScalingPolicyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.CfnScalingPolicy",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnScalingPolicy)SetAdjustmentType(val *string) {
	_jsii_.Set(
		j,
		"adjustmentType",
		val,
	)
}

func (j *jsiiProxy_CfnScalingPolicy)SetAutoScalingGroupName(val *string) {
	if err := j.validateSetAutoScalingGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoScalingGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnScalingPolicy)SetCooldown(val *string) {
	_jsii_.Set(
		j,
		"cooldown",
		val,
	)
}

func (j *jsiiProxy_CfnScalingPolicy)SetEstimatedInstanceWarmup(val *float64) {
	_jsii_.Set(
		j,
		"estimatedInstanceWarmup",
		val,
	)
}

func (j *jsiiProxy_CfnScalingPolicy)SetMetricAggregationType(val *string) {
	_jsii_.Set(
		j,
		"metricAggregationType",
		val,
	)
}

func (j *jsiiProxy_CfnScalingPolicy)SetMinAdjustmentMagnitude(val *float64) {
	_jsii_.Set(
		j,
		"minAdjustmentMagnitude",
		val,
	)
}

func (j *jsiiProxy_CfnScalingPolicy)SetPolicyType(val *string) {
	_jsii_.Set(
		j,
		"policyType",
		val,
	)
}

func (j *jsiiProxy_CfnScalingPolicy)SetPredictiveScalingConfiguration(val interface{}) {
	if err := j.validateSetPredictiveScalingConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"predictiveScalingConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnScalingPolicy)SetScalingAdjustment(val *float64) {
	_jsii_.Set(
		j,
		"scalingAdjustment",
		val,
	)
}

func (j *jsiiProxy_CfnScalingPolicy)SetStepAdjustments(val interface{}) {
	if err := j.validateSetStepAdjustmentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stepAdjustments",
		val,
	)
}

func (j *jsiiProxy_CfnScalingPolicy)SetTargetTrackingConfiguration(val interface{}) {
	if err := j.validateSetTargetTrackingConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetTrackingConfiguration",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnScalingPolicy_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnScalingPolicy_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.CfnScalingPolicy",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnScalingPolicy_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnScalingPolicy_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.CfnScalingPolicy",
		"isCfnResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnScalingPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnScalingPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.CfnScalingPolicy",
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
		"aws-cdk-lib.aws_autoscaling.CfnScalingPolicy",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnScalingPolicy) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnScalingPolicy) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnScalingPolicy) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnScalingPolicy) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnScalingPolicy) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnScalingPolicy) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnScalingPolicy) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnScalingPolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnScalingPolicy) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScalingPolicy) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
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
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnScalingPolicy) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScalingPolicy) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScalingPolicy) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnScalingPolicy) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnScalingPolicy) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScalingPolicy) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
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

func (c *jsiiProxy_CfnScalingPolicy) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

