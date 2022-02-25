package awsautoscalingplans

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscalingplans/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::AutoScalingPlans::ScalingPlan`.
//
// The `AWS::AutoScalingPlans::ScalingPlan` resource defines an AWS Auto Scaling scaling plan. A scaling plan is used to scale application resources to size them appropriately to ensure that enough resource is available in the application at peak times and to reduce allocated resource during periods of low utilization. The following resources can be added to a scaling plan:
//
// - Amazon EC2 Auto Scaling groups
// - Amazon EC2 Spot Fleet requests
// - Amazon ECS services
// - Amazon DynamoDB tables and global secondary indexes
// - Amazon Aurora Replicas
//
// For more information, see the [AWS Auto Scaling User Guide](https://docs.aws.amazon.com/autoscaling/plans/userguide/what-is-aws-auto-scaling.html) .
//
// TODO: EXAMPLE
//
type CfnScalingPlan interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ApplicationSource() interface{}
	SetApplicationSource(val interface{})
	AttrScalingPlanName() *string
	AttrScalingPlanVersion() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	ScalingInstructions() interface{}
	SetScalingInstructions(val interface{})
	Stack() awscdk.Stack
	UpdatedProperites() *map[string]interface{}
	AddDeletionOverride(path *string)
	AddDependsOn(target awscdk.CfnResource)
	AddMetadata(key *string, value interface{})
	AddOverride(path *string, value interface{})
	AddPropertyDeletionOverride(propertyPath *string)
	AddPropertyOverride(propertyPath *string, value interface{})
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	GetAtt(attributeName *string) awscdk.Reference
	GetMetadata(key *string) interface{}
	Inspect(inspector awscdk.TreeInspector)
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnScalingPlan
type jsiiProxy_CfnScalingPlan struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnScalingPlan) ApplicationSource() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applicationSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPlan) AttrScalingPlanName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrScalingPlanName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPlan) AttrScalingPlanVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrScalingPlanVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPlan) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPlan) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPlan) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPlan) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPlan) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPlan) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPlan) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPlan) ScalingInstructions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingInstructions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPlan) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalingPlan) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::AutoScalingPlans::ScalingPlan`.
func NewCfnScalingPlan(scope constructs.Construct, id *string, props *CfnScalingPlanProps) CfnScalingPlan {
	_init_.Initialize()

	j := jsiiProxy_CfnScalingPlan{}

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscalingplans.CfnScalingPlan",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AutoScalingPlans::ScalingPlan`.
func NewCfnScalingPlan_Override(c CfnScalingPlan, scope constructs.Construct, id *string, props *CfnScalingPlanProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscalingplans.CfnScalingPlan",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnScalingPlan) SetApplicationSource(val interface{}) {
	_jsii_.Set(
		j,
		"applicationSource",
		val,
	)
}

func (j *jsiiProxy_CfnScalingPlan) SetScalingInstructions(val interface{}) {
	_jsii_.Set(
		j,
		"scalingInstructions",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnScalingPlan_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscalingplans.CfnScalingPlan",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnScalingPlan_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscalingplans.CfnScalingPlan",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func CfnScalingPlan_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscalingplans.CfnScalingPlan",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnScalingPlan_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_autoscalingplans.CfnScalingPlan",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnScalingPlan) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
//
// This can be used for resources across stacks (or nested stack) boundaries
// and the dependency will automatically be transferred to the relevant scope.
func (c *jsiiProxy_CfnScalingPlan) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

// Add a value to the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnScalingPlan) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

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
func (c *jsiiProxy_CfnScalingPlan) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnScalingPlan) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnScalingPlan) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
func (c *jsiiProxy_CfnScalingPlan) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

// Returns a token for an runtime attribute of this resource.
//
// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
// in case there is no generated attribute.
func (c *jsiiProxy_CfnScalingPlan) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

// Retrieve a value value from the CloudFormation Resource Metadata.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
//
// Note that this is a different set of metadata from CDK node metadata; this
// metadata ends up in the stack template under the resource, whereas CDK
// node metadata ends up in the Cloud Assembly.
//
func (c *jsiiProxy_CfnScalingPlan) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Examines the CloudFormation resource and discloses attributes.
func (c *jsiiProxy_CfnScalingPlan) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnScalingPlan) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnScalingPlan) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
//
// Returns: `true` if the resource should be included or `false` is the resource
// should be omitted.
func (c *jsiiProxy_CfnScalingPlan) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
func (c *jsiiProxy_CfnScalingPlan) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScalingPlan) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// `ApplicationSource` is a property of [ScalingPlan](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscalingplans-scalingplan.html) that specifies the application source to use with AWS Auto Scaling ( Auto Scaling Plans ). You can create one scaling plan per application source.
//
// TODO: EXAMPLE
//
type CfnScalingPlan_ApplicationSourceProperty struct {
	// The Amazon Resource Name (ARN) of a CloudFormation stack.
	//
	// You must specify either a `CloudFormationStackARN` or `TagFilters` .
	CloudFormationStackArn *string `json:"cloudFormationStackArn" yaml:"cloudFormationStackArn"`
	// A set of tag filters (keys and values).
	//
	// Each tag filter specified must contain a key with values as optional. Each scaling plan can include up to 50 keys, and each key can include up to 20 values.
	//
	// Tags are part of the syntax that you use to specify the resources you want returned when configuring a scaling plan from the AWS Auto Scaling console. You do not need to specify valid tag filter values when you create a scaling plan with CloudFormation. The `Key` and `Values` properties can accept any value as long as the combination of values is unique across scaling plans. However, if you also want to use the AWS Auto Scaling console to edit the scaling plan, then you must specify valid values.
	//
	// You must specify either a `CloudFormationStackARN` or `TagFilters` .
	TagFilters interface{} `json:"tagFilters" yaml:"tagFilters"`
}

// `CustomizedLoadMetricSpecification` is a subproperty of [ScalingInstruction](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html) that specifies a customized load metric for predictive scaling to use with AWS Auto Scaling ( Auto Scaling Plans ).
//
// For predictive scaling to work with a customized load metric specification, AWS Auto Scaling needs access to the `Sum` and `Average` statistics that CloudWatch computes from metric data.
//
// When you choose a load metric, make sure that the required `Sum` and `Average` statistics for your metric are available in CloudWatch and that they provide relevant data for predictive scaling. The `Sum` statistic must represent the total load on the resource, and the `Average` statistic must represent the average load per capacity unit of the resource. For example, there is a metric that counts the number of requests processed by your Auto Scaling group. If the `Sum` statistic represents the total request count processed by the group, then the `Average` statistic for the specified metric must represent the average request count processed by each instance of the group.
//
// If you publish your own metrics, you can aggregate the data points at a given interval and then publish the aggregated data points to CloudWatch. Before AWS Auto Scaling generates the forecast, it sums up all the metric data points that occurred within each hour to match the granularity period that is used in the forecast (60 minutes).
//
// For information about terminology, available metrics, or how to publish new metrics, see [Amazon CloudWatch Concepts](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html) in the *Amazon CloudWatch User Guide* .
//
// After creating your scaling plan, you can use the AWS Auto Scaling console to visualize forecasts for the specified metric. For more information, see [View Scaling Information for a Resource](https://docs.aws.amazon.com/autoscaling/plans/userguide/gs-create-scaling-plan.html#gs-view-resource) in the *AWS Auto Scaling User Guide* .
//
// TODO: EXAMPLE
//
type CfnScalingPlan_CustomizedLoadMetricSpecificationProperty struct {
	// The name of the metric.
	MetricName *string `json:"metricName" yaml:"metricName"`
	// The namespace of the metric.
	Namespace *string `json:"namespace" yaml:"namespace"`
	// The statistic of the metric.
	//
	// *Allowed Values* : `Sum`
	Statistic *string `json:"statistic" yaml:"statistic"`
	// The dimensions of the metric.
	//
	// Conditional: If you published your metric with dimensions, you must specify the same dimensions in your customized load metric specification.
	Dimensions interface{} `json:"dimensions" yaml:"dimensions"`
	// The unit of the metric.
	Unit *string `json:"unit" yaml:"unit"`
}

// `CustomizedScalingMetricSpecification` is a subproperty of [TargetTrackingConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-targettrackingconfiguration.html) that specifies a customized scaling metric for a target tracking configuration to use with AWS Auto Scaling ( Auto Scaling Plans ).
//
// To create your customized scaling metric specification:
//
// - Add values for each required property from CloudWatch. You can use an existing metric, or a new metric that you create. To use your own metric, you must first publish the metric to CloudWatch. For more information, see [Publish Custom Metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/publishingMetrics.html) in the *Amazon CloudWatch User Guide* .
// - Choose a metric that changes proportionally with capacity. The value of the metric should increase or decrease in inverse proportion to the number of capacity units. That is, the value of the metric should decrease when capacity increases.
//
// For information about terminology, available metrics, or how to publish new metrics, see [Amazon CloudWatch Concepts](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch_concepts.html) in the *Amazon CloudWatch User Guide* .
//
// TODO: EXAMPLE
//
type CfnScalingPlan_CustomizedScalingMetricSpecificationProperty struct {
	// The name of the metric.
	MetricName *string `json:"metricName" yaml:"metricName"`
	// The namespace of the metric.
	Namespace *string `json:"namespace" yaml:"namespace"`
	// The statistic of the metric.
	Statistic *string `json:"statistic" yaml:"statistic"`
	// The dimensions of the metric.
	//
	// Conditional: If you published your metric with dimensions, you must specify the same dimensions in your customized scaling metric specification.
	Dimensions interface{} `json:"dimensions" yaml:"dimensions"`
	// The unit of the metric.
	Unit *string `json:"unit" yaml:"unit"`
}

// `MetricDimension` is a subproperty of [CustomizedScalingMetricSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-customizedscalingmetricspecification.html) that specifies a dimension for a customized metric to use with AWS Auto Scaling ( Auto Scaling Plans ). Dimensions are arbitrary name/value pairs that can be associated with a CloudWatch metric. Duplicate dimensions are not allowed.
//
// TODO: EXAMPLE
//
type CfnScalingPlan_MetricDimensionProperty struct {
	// The name of the dimension.
	Name *string `json:"name" yaml:"name"`
	// The value of the dimension.
	Value *string `json:"value" yaml:"value"`
}

// `PredefinedLoadMetricSpecification` is a subproperty of [ScalingInstruction](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html) that specifies a predefined load metric for predictive scaling to use with AWS Auto Scaling ( Auto Scaling Plans ).
//
// After creating your scaling plan, you can use the AWS Auto Scaling console to visualize forecasts for the specified metric. For more information, see [View Scaling Information for a Resource](https://docs.aws.amazon.com/autoscaling/plans/userguide/gs-create-scaling-plan.html#gs-view-resource) in the *AWS Auto Scaling User Guide* .
//
// TODO: EXAMPLE
//
type CfnScalingPlan_PredefinedLoadMetricSpecificationProperty struct {
	// The metric type.
	PredefinedLoadMetricType *string `json:"predefinedLoadMetricType" yaml:"predefinedLoadMetricType"`
	// Identifies the resource associated with the metric type.
	//
	// You can't specify a resource label unless the metric type is `ALBTargetGroupRequestCount` and there is a target group for an Application Load Balancer attached to the Auto Scaling group.
	//
	// You create the resource label by appending the final portion of the load balancer ARN and the final portion of the target group ARN into a single value, separated by a forward slash (/). The format is app/<load-balancer-name>/<load-balancer-id>/targetgroup/<target-group-name>/<target-group-id>, where:
	//
	// - app/<load-balancer-name>/<load-balancer-id> is the final portion of the load balancer ARN
	// - targetgroup/<target-group-name>/<target-group-id> is the final portion of the target group ARN.
	//
	// This is an example: app/EC2Co-EcsEl-1TKLTMITMM0EO/f37c06a68c1748aa/targetgroup/EC2Co-Defau-LDNM7Q3ZH1ZN/6d4ea56ca2d6a18d.
	//
	// To find the ARN for an Application Load Balancer, use the [DescribeLoadBalancers](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_DescribeLoadBalancers.html) API operation. To find the ARN for the target group, use the [DescribeTargetGroups](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_DescribeTargetGroups.html) API operation.
	ResourceLabel *string `json:"resourceLabel" yaml:"resourceLabel"`
}

// `PredefinedScalingMetricSpecification` is a subproperty of [TargetTrackingConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-targettrackingconfiguration.html) that specifies a customized scaling metric for a target tracking configuration to use with AWS Auto Scaling ( Auto Scaling Plans ).
//
// TODO: EXAMPLE
//
type CfnScalingPlan_PredefinedScalingMetricSpecificationProperty struct {
	// The metric type.
	//
	// The `ALBRequestCountPerTarget` metric type applies only to Auto Scaling groups, Spot Fleet requests, and ECS services.
	PredefinedScalingMetricType *string `json:"predefinedScalingMetricType" yaml:"predefinedScalingMetricType"`
	// Identifies the resource associated with the metric type.
	//
	// You can't specify a resource label unless the metric type is `ALBRequestCountPerTarget` and there is a target group for an Application Load Balancer attached to the Auto Scaling group, Spot Fleet request, or ECS service.
	//
	// You create the resource label by appending the final portion of the load balancer ARN and the final portion of the target group ARN into a single value, separated by a forward slash (/). The format is app/<load-balancer-name>/<load-balancer-id>/targetgroup/<target-group-name>/<target-group-id>, where:
	//
	// - app/<load-balancer-name>/<load-balancer-id> is the final portion of the load balancer ARN
	// - targetgroup/<target-group-name>/<target-group-id> is the final portion of the target group ARN.
	//
	// This is an example: app/EC2Co-EcsEl-1TKLTMITMM0EO/f37c06a68c1748aa/targetgroup/EC2Co-Defau-LDNM7Q3ZH1ZN/6d4ea56ca2d6a18d.
	//
	// To find the ARN for an Application Load Balancer, use the [DescribeLoadBalancers](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_DescribeLoadBalancers.html) API operation. To find the ARN for the target group, use the [DescribeTargetGroups](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_DescribeTargetGroups.html) API operation.
	ResourceLabel *string `json:"resourceLabel" yaml:"resourceLabel"`
}

// `ScalingInstruction` is a property of [ScalingPlan](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscalingplans-scalingplan.html) that specifies the scaling instruction for a scalable resource in a scaling plan. Each scaling instruction applies to one resource.
//
// AWS Auto Scaling creates target tracking scaling policies based on the scaling instructions. Target tracking scaling policies adjust the capacity of your scalable resource as required to maintain resource utilization at the target value that you specified.
//
// AWS Auto Scaling also configures predictive scaling for your Amazon EC2 Auto Scaling groups using a subset of properties, including the load metric, the scaling metric, the target value for the scaling metric, the predictive scaling mode (forecast and scale or forecast only), and the desired behavior when the forecast capacity exceeds the maximum capacity of the resource. With predictive scaling, AWS Auto Scaling generates forecasts with traffic predictions for the two days ahead and schedules scaling actions that proactively add and remove resource capacity to match the forecast.
//
// > We recommend waiting a minimum of 24 hours after creating an Auto Scaling group to configure predictive scaling. At minimum, there must be 24 hours of historical data to generate a forecast. For more information, see [Best Practices for AWS Auto Scaling](https://docs.aws.amazon.com/autoscaling/plans/userguide/gs-best-practices.html) in the *AWS Auto Scaling User Guide* .
//
// TODO: EXAMPLE
//
type CfnScalingPlan_ScalingInstructionProperty struct {
	// The maximum capacity of the resource.
	//
	// The exception to this upper limit is if you specify a non-default setting for *PredictiveScalingMaxCapacityBehavior* .
	MaxCapacity *float64 `json:"maxCapacity" yaml:"maxCapacity"`
	// The minimum capacity of the resource.
	MinCapacity *float64 `json:"minCapacity" yaml:"minCapacity"`
	// The ID of the resource. This string consists of the resource type and unique identifier.
	//
	// - Auto Scaling group - The resource type is `autoScalingGroup` and the unique identifier is the name of the Auto Scaling group. Example: `autoScalingGroup/my-asg` .
	// - ECS service - The resource type is `service` and the unique identifier is the cluster name and service name. Example: `service/default/sample-webapp` .
	// - Spot Fleet request - The resource type is `spot-fleet-request` and the unique identifier is the Spot Fleet request ID. Example: `spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE` .
	// - DynamoDB table - The resource type is `table` and the unique identifier is the resource ID. Example: `table/my-table` .
	// - DynamoDB global secondary index - The resource type is `index` and the unique identifier is the resource ID. Example: `table/my-table/index/my-table-index` .
	// - Aurora DB cluster - The resource type is `cluster` and the unique identifier is the cluster name. Example: `cluster:my-db-cluster` .
	ResourceId *string `json:"resourceId" yaml:"resourceId"`
	// The scalable dimension associated with the resource.
	//
	// - `autoscaling:autoScalingGroup:DesiredCapacity` - The desired capacity of an Auto Scaling group.
	// - `ecs:service:DesiredCount` - The desired task count of an ECS service.
	// - `ec2:spot-fleet-request:TargetCapacity` - The target capacity of a Spot Fleet request.
	// - `dynamodb:table:ReadCapacityUnits` - The provisioned read capacity for a DynamoDB table.
	// - `dynamodb:table:WriteCapacityUnits` - The provisioned write capacity for a DynamoDB table.
	// - `dynamodb:index:ReadCapacityUnits` - The provisioned read capacity for a DynamoDB global secondary index.
	// - `dynamodb:index:WriteCapacityUnits` - The provisioned write capacity for a DynamoDB global secondary index.
	// - `rds:cluster:ReadReplicaCount` - The count of Aurora Replicas in an Aurora DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible edition.
	ScalableDimension *string `json:"scalableDimension" yaml:"scalableDimension"`
	// The namespace of the AWS service.
	ServiceNamespace *string `json:"serviceNamespace" yaml:"serviceNamespace"`
	// The target tracking configurations (up to 10).
	//
	// Each of these structures must specify a unique scaling metric and a target value for the metric.
	TargetTrackingConfigurations interface{} `json:"targetTrackingConfigurations" yaml:"targetTrackingConfigurations"`
	// The customized load metric to use for predictive scaling.
	//
	// This property or a *PredefinedLoadMetricSpecification* is required when configuring predictive scaling, and cannot be used otherwise.
	CustomizedLoadMetricSpecification interface{} `json:"customizedLoadMetricSpecification" yaml:"customizedLoadMetricSpecification"`
	// Controls whether dynamic scaling by AWS Auto Scaling is disabled.
	//
	// When dynamic scaling is enabled, AWS Auto Scaling creates target tracking scaling policies based on the specified target tracking configurations.
	//
	// The default is enabled ( `false` ).
	DisableDynamicScaling interface{} `json:"disableDynamicScaling" yaml:"disableDynamicScaling"`
	// The predefined load metric to use for predictive scaling.
	//
	// This property or a *CustomizedLoadMetricSpecification* is required when configuring predictive scaling, and cannot be used otherwise.
	PredefinedLoadMetricSpecification interface{} `json:"predefinedLoadMetricSpecification" yaml:"predefinedLoadMetricSpecification"`
	// Defines the behavior that should be applied if the forecast capacity approaches or exceeds the maximum capacity specified for the resource.
	//
	// The default value is `SetForecastCapacityToMaxCapacity` .
	//
	// The following are possible values:
	//
	// - `SetForecastCapacityToMaxCapacity` - AWS Auto Scaling cannot scale resource capacity higher than the maximum capacity. The maximum capacity is enforced as a hard limit.
	// - `SetMaxCapacityToForecastCapacity` - AWS Auto Scaling can scale resource capacity higher than the maximum capacity to equal but not exceed forecast capacity.
	// - `SetMaxCapacityAboveForecastCapacity` - AWS Auto Scaling can scale resource capacity higher than the maximum capacity by a specified buffer value. The intention is to give the target tracking scaling policy extra capacity if unexpected traffic occurs.
	//
	// Valid only when configuring predictive scaling.
	PredictiveScalingMaxCapacityBehavior *string `json:"predictiveScalingMaxCapacityBehavior" yaml:"predictiveScalingMaxCapacityBehavior"`
	// The size of the capacity buffer to use when the forecast capacity is close to or exceeds the maximum capacity.
	//
	// The value is specified as a percentage relative to the forecast capacity. For example, if the buffer is 10, this means a 10 percent buffer. With a 10 percent buffer, if the forecast capacity is 50, and the maximum capacity is 40, then the effective maximum capacity is 55.
	//
	// Valid only when configuring predictive scaling. Required if *PredictiveScalingMaxCapacityBehavior* is set to `SetMaxCapacityAboveForecastCapacity` , and cannot be used otherwise.
	//
	// The range is 1-100.
	PredictiveScalingMaxCapacityBuffer *float64 `json:"predictiveScalingMaxCapacityBuffer" yaml:"predictiveScalingMaxCapacityBuffer"`
	// The predictive scaling mode.
	//
	// The default value is `ForecastAndScale` . Otherwise, AWS Auto Scaling forecasts capacity but does not apply any scheduled scaling actions based on the capacity forecast.
	PredictiveScalingMode *string `json:"predictiveScalingMode" yaml:"predictiveScalingMode"`
	// Controls whether your scaling policies that are external to AWS Auto Scaling are deleted and new target tracking scaling policies created.
	//
	// The default value is `KeepExternalPolicies` .
	//
	// Valid only when configuring dynamic scaling.
	ScalingPolicyUpdateBehavior *string `json:"scalingPolicyUpdateBehavior" yaml:"scalingPolicyUpdateBehavior"`
	// The amount of time, in seconds, to buffer the run time of scheduled scaling actions when scaling out.
	//
	// For example, if the forecast says to add capacity at 10:00 AM, and the buffer time is 5 minutes, then the run time of the corresponding scheduled scaling action will be 9:55 AM. The intention is to give resources time to be provisioned. For example, it can take a few minutes to launch an EC2 instance. The actual amount of time required depends on several factors, such as the size of the instance and whether there are startup scripts to complete.
	//
	// The value must be less than the forecast interval duration of 3600 seconds (60 minutes). The default is 300 seconds.
	//
	// Valid only when configuring predictive scaling.
	ScheduledActionBufferTime *float64 `json:"scheduledActionBufferTime" yaml:"scheduledActionBufferTime"`
}

// `TagFilter` is a subproperty of [ApplicationSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-applicationsource.html) that specifies a tag for an application source to use with AWS Auto Scaling ( Auto Scaling Plans ).
//
// TODO: EXAMPLE
//
type CfnScalingPlan_TagFilterProperty struct {
	// The tag key.
	Key *string `json:"key" yaml:"key"`
	// The tag values (0 to 20).
	Values *[]*string `json:"values" yaml:"values"`
}

// `TargetTrackingConfiguration` is a subproperty of [ScalingInstruction](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html) that specifies a target tracking configuration to use with AWS Auto Scaling ( Auto Scaling Plans ).
//
// TODO: EXAMPLE
//
type CfnScalingPlan_TargetTrackingConfigurationProperty struct {
	// The target value for the metric.
	//
	// Although this property accepts numbers of type Double, it won't accept values that are either too small or too large. Values must be in the range of -2^360 to 2^360.
	TargetValue *float64 `json:"targetValue" yaml:"targetValue"`
	// A customized metric.
	//
	// You can specify either a predefined metric or a customized metric.
	CustomizedScalingMetricSpecification interface{} `json:"customizedScalingMetricSpecification" yaml:"customizedScalingMetricSpecification"`
	// Indicates whether scale in by the target tracking scaling policy is disabled.
	//
	// If the value is `true` , scale in is disabled and the target tracking scaling policy doesn't remove capacity from the scalable resource. Otherwise, scale in is enabled and the target tracking scaling policy can remove capacity from the scalable resource.
	//
	// The default value is `false` .
	DisableScaleIn interface{} `json:"disableScaleIn" yaml:"disableScaleIn"`
	// The estimated time, in seconds, until a newly launched instance can contribute to the CloudWatch metrics.
	//
	// This value is used only if the resource is an Auto Scaling group.
	EstimatedInstanceWarmup *float64 `json:"estimatedInstanceWarmup" yaml:"estimatedInstanceWarmup"`
	// A predefined metric.
	//
	// You can specify either a predefined metric or a customized metric.
	PredefinedScalingMetricSpecification interface{} `json:"predefinedScalingMetricSpecification" yaml:"predefinedScalingMetricSpecification"`
	// The amount of time, in seconds, after a scale-in activity completes before another scale in activity can start.
	//
	// This value is not used if the scalable resource is an Auto Scaling group.
	ScaleInCooldown *float64 `json:"scaleInCooldown" yaml:"scaleInCooldown"`
	// The amount of time, in seconds, after a scale-out activity completes before another scale-out activity can start.
	//
	// This value is not used if the scalable resource is an Auto Scaling group.
	ScaleOutCooldown *float64 `json:"scaleOutCooldown" yaml:"scaleOutCooldown"`
}

// Properties for defining a `CfnScalingPlan`.
//
// TODO: EXAMPLE
//
type CfnScalingPlanProps struct {
	// A CloudFormation stack or a set of tags.
	//
	// You can create one scaling plan per application source. The `ApplicationSource` property must be present to ensure interoperability with the AWS Auto Scaling console.
	ApplicationSource interface{} `json:"applicationSource" yaml:"applicationSource"`
	// The scaling instructions.
	ScalingInstructions interface{} `json:"scalingInstructions" yaml:"scalingInstructions"`
}

