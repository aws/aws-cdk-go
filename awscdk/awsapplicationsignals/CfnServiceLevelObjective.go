package awsapplicationsignals

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationsignals/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapplicationsignals"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates or updates a service level objective (SLO), which can help you ensure that your critical business operations are meeting customer expectations.
//
// Use SLOs to set and track specific target levels for the reliability and availability of your applications and services. SLOs use service level indicators (SLIs) to calculate whether the application is performing at the level that you want.
//
// Create an SLO to set a target for a service operation, or service dependency's availability or latency. CloudWatch measures this target frequently you can find whether it has been breached.
//
// The target performance quality that is defined for an SLO is the *attainment goal* . An attainment goal is the percentage of time or requests that the SLI is expected to meet the threshold over each time interval. For example, an attainment goal of 99.9% means that within your interval, you are targeting 99.9% of the periods to be in healthy state.
//
// When you create an SLO, you specify whether it is a *period-based SLO* or a *request-based SLO* . Each type of SLO has a different way of evaluating your application's performance against its attainment goal.
//
// - A *period-based SLO* uses defined *periods* of time within a specified total time interval. For each period of time, Application Signals determines whether the application met its goal. The attainment rate is calculated as the `number of good periods/number of total periods` .
//
// For example, for a period-based SLO, meeting an attainment goal of 99.9% means that within your interval, your application must meet its performance goal during at least 99.9% of the time periods.
// - A *request-based SLO* doesn't use pre-defined periods of time. Instead, the SLO measures `number of good requests/number of total requests` during the interval. At any time, you can find the ratio of good requests to total requests for the interval up to the time stamp that you specify, and measure that ratio against the goal set in your SLO.
//
// After you have created an SLO, you can retrieve error budget reports for it. An *error budget* is the amount of time or amount of requests that your application can be non-compliant with the SLO's goal, and still have your application meet the goal.
//
// - For a period-based SLO, the error budget starts at a number defined by the highest number of periods that can fail to meet the threshold, while still meeting the overall goal. The *remaining error budget* decreases with every failed period that is recorded. The error budget within one interval can never increase.
//
// For example, an SLO with a threshold that 99.95% of requests must be completed under 2000ms every month translates to an error budget of 21.9 minutes of downtime per month.
// - For a request-based SLO, the remaining error budget is dynamic and can increase or decrease, depending on the ratio of good requests to total requests.
//
// When you call this operation, Application Signals creates the *AWSServiceRoleForCloudWatchApplicationSignals* service-linked role, if it doesn't already exist in your account. This service- linked role has the following permissions:
//
// - `xray:GetServiceGraph`
// - `logs:StartQuery`
// - `logs:GetQueryResults`
// - `cloudwatch:GetMetricData`
// - `cloudwatch:ListMetrics`
// - `tag:GetResources`
// - `autoscaling:DescribeAutoScalingGroups`
//
// You can easily set SLO targets for your applications, and their dependencies, that are discovered by Application Signals, using critical metrics such as latency and availability. You can also set SLOs against any CloudWatch metric or math expression that produces a time series.
//
// > You can't create an SLO for a service operation that was discovered by Application Signals until after that operation has reported standard metrics to Application Signals.
//
// You cannot change from a period-based SLO to a request-based SLO, or change from a request-based SLO to a period-based SLO.
//
// For more information about SLOs, see [Service level objectives (SLOs)](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-ServiceLevelObjectives.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServiceLevelObjective := awscdk.Aws_applicationsignals.NewCfnServiceLevelObjective(this, jsii.String("MyCfnServiceLevelObjective"), &CfnServiceLevelObjectiveProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	BurnRateConfigurations: []interface{}{
//   		&BurnRateConfigurationProperty{
//   			LookBackWindowMinutes: jsii.Number(123),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	ExclusionWindows: []interface{}{
//   		&ExclusionWindowProperty{
//   			Window: &WindowProperty{
//   				Duration: jsii.Number(123),
//   				DurationUnit: jsii.String("durationUnit"),
//   			},
//
//   			// the properties below are optional
//   			Reason: jsii.String("reason"),
//   			RecurrenceRule: &RecurrenceRuleProperty{
//   				Expression: jsii.String("expression"),
//   			},
//   			StartTime: jsii.String("startTime"),
//   		},
//   	},
//   	Goal: &GoalProperty{
//   		AttainmentGoal: jsii.Number(123),
//   		Interval: &IntervalProperty{
//   			CalendarInterval: &CalendarIntervalProperty{
//   				Duration: jsii.Number(123),
//   				DurationUnit: jsii.String("durationUnit"),
//   				StartTime: jsii.Number(123),
//   			},
//   			RollingInterval: &RollingIntervalProperty{
//   				Duration: jsii.Number(123),
//   				DurationUnit: jsii.String("durationUnit"),
//   			},
//   		},
//   		WarningThreshold: jsii.Number(123),
//   	},
//   	RequestBasedSli: &RequestBasedSliProperty{
//   		RequestBasedSliMetric: &RequestBasedSliMetricProperty{
//   			DependencyConfig: &DependencyConfigProperty{
//   				DependencyKeyAttributes: map[string]*string{
//   					"dependencyKeyAttributesKey": jsii.String("dependencyKeyAttributes"),
//   				},
//   				DependencyOperationName: jsii.String("dependencyOperationName"),
//   			},
//   			KeyAttributes: map[string]*string{
//   				"keyAttributesKey": jsii.String("keyAttributes"),
//   			},
//   			MetricType: jsii.String("metricType"),
//   			MonitoredRequestCountMetric: &MonitoredRequestCountMetricProperty{
//   				BadCountMetric: []interface{}{
//   					&MetricDataQueryProperty{
//   						Id: jsii.String("id"),
//
//   						// the properties below are optional
//   						AccountId: jsii.String("accountId"),
//   						Expression: jsii.String("expression"),
//   						MetricStat: &MetricStatProperty{
//   							Metric: &MetricProperty{
//   								Dimensions: []interface{}{
//   									&DimensionProperty{
//   										Name: jsii.String("name"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   								MetricName: jsii.String("metricName"),
//   								Namespace: jsii.String("namespace"),
//   							},
//   							Period: jsii.Number(123),
//   							Stat: jsii.String("stat"),
//
//   							// the properties below are optional
//   							Unit: jsii.String("unit"),
//   						},
//   						ReturnData: jsii.Boolean(false),
//   					},
//   				},
//   				GoodCountMetric: []interface{}{
//   					&MetricDataQueryProperty{
//   						Id: jsii.String("id"),
//
//   						// the properties below are optional
//   						AccountId: jsii.String("accountId"),
//   						Expression: jsii.String("expression"),
//   						MetricStat: &MetricStatProperty{
//   							Metric: &MetricProperty{
//   								Dimensions: []interface{}{
//   									&DimensionProperty{
//   										Name: jsii.String("name"),
//   										Value: jsii.String("value"),
//   									},
//   								},
//   								MetricName: jsii.String("metricName"),
//   								Namespace: jsii.String("namespace"),
//   							},
//   							Period: jsii.Number(123),
//   							Stat: jsii.String("stat"),
//
//   							// the properties below are optional
//   							Unit: jsii.String("unit"),
//   						},
//   						ReturnData: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			OperationName: jsii.String("operationName"),
//   			TotalRequestCountMetric: []interface{}{
//   				&MetricDataQueryProperty{
//   					Id: jsii.String("id"),
//
//   					// the properties below are optional
//   					AccountId: jsii.String("accountId"),
//   					Expression: jsii.String("expression"),
//   					MetricStat: &MetricStatProperty{
//   						Metric: &MetricProperty{
//   							Dimensions: []interface{}{
//   								&DimensionProperty{
//   									Name: jsii.String("name"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							MetricName: jsii.String("metricName"),
//   							Namespace: jsii.String("namespace"),
//   						},
//   						Period: jsii.Number(123),
//   						Stat: jsii.String("stat"),
//
//   						// the properties below are optional
//   						Unit: jsii.String("unit"),
//   					},
//   					ReturnData: jsii.Boolean(false),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		ComparisonOperator: jsii.String("comparisonOperator"),
//   		MetricThreshold: jsii.Number(123),
//   	},
//   	Sli: &SliProperty{
//   		ComparisonOperator: jsii.String("comparisonOperator"),
//   		MetricThreshold: jsii.Number(123),
//   		SliMetric: &SliMetricProperty{
//   			DependencyConfig: &DependencyConfigProperty{
//   				DependencyKeyAttributes: map[string]*string{
//   					"dependencyKeyAttributesKey": jsii.String("dependencyKeyAttributes"),
//   				},
//   				DependencyOperationName: jsii.String("dependencyOperationName"),
//   			},
//   			KeyAttributes: map[string]*string{
//   				"keyAttributesKey": jsii.String("keyAttributes"),
//   			},
//   			MetricDataQueries: []interface{}{
//   				&MetricDataQueryProperty{
//   					Id: jsii.String("id"),
//
//   					// the properties below are optional
//   					AccountId: jsii.String("accountId"),
//   					Expression: jsii.String("expression"),
//   					MetricStat: &MetricStatProperty{
//   						Metric: &MetricProperty{
//   							Dimensions: []interface{}{
//   								&DimensionProperty{
//   									Name: jsii.String("name"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							MetricName: jsii.String("metricName"),
//   							Namespace: jsii.String("namespace"),
//   						},
//   						Period: jsii.Number(123),
//   						Stat: jsii.String("stat"),
//
//   						// the properties below are optional
//   						Unit: jsii.String("unit"),
//   					},
//   					ReturnData: jsii.Boolean(false),
//   				},
//   			},
//   			MetricType: jsii.String("metricType"),
//   			OperationName: jsii.String("operationName"),
//   			PeriodSeconds: jsii.Number(123),
//   			Statistic: jsii.String("statistic"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationsignals-servicelevelobjective.html
//
type CfnServiceLevelObjective interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsapplicationsignals.IServiceLevelObjectiveRef
	awscdk.ITaggableV2
	// The ARN of this SLO.
	AttrArn() *string
	// The date and time that this SLO was created.
	AttrCreatedTime() *float64
	// Displays whether this is a period-based SLO or a request-based SLO.
	AttrEvaluationType() *string
	// The time that this SLO was most recently updated.
	AttrLastUpdatedTime() *float64
	// Each object in this array defines the length of the look-back window used to calculate one burn rate metric for this SLO.
	BurnRateConfigurations() interface{}
	SetBurnRateConfigurations(val interface{})
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// An optional description for this SLO.
	Description() *string
	SetDescription(val *string)
	Env() *interfaces.ResourceEnvironment
	// The time window to be excluded from the SLO performance metrics.
	ExclusionWindows() interface{}
	SetExclusionWindows(val interface{})
	// This structure contains the attributes that determine the goal of an SLO.
	Goal() interface{}
	SetGoal(val interface{})
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
	// A name for this SLO.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// A structure containing information about the performance metric that this SLO monitors, if this is a request-based SLO.
	RequestBasedSli() interface{}
	SetRequestBasedSli(val interface{})
	// A reference to a ServiceLevelObjective resource.
	ServiceLevelObjectiveRef() *interfacesawsapplicationsignals.ServiceLevelObjectiveReference
	// A structure containing information about the performance metric that this SLO monitors, if this is a period-based SLO.
	Sli() interface{}
	SetSli(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// A list of key-value pairs to associate with the SLO.
	Tags() *[]*awscdk.CfnTag
	SetTags(val *[]*awscdk.CfnTag)
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
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for CfnServiceLevelObjective
type jsiiProxy_CfnServiceLevelObjective struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsapplicationsignalsIServiceLevelObjectiveRef
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnServiceLevelObjective) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceLevelObjective) AttrCreatedTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrCreatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceLevelObjective) AttrEvaluationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEvaluationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceLevelObjective) AttrLastUpdatedTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"attrLastUpdatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceLevelObjective) BurnRateConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"burnRateConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceLevelObjective) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceLevelObjective) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceLevelObjective) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceLevelObjective) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceLevelObjective) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceLevelObjective) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceLevelObjective) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceLevelObjective) ExclusionWindows() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exclusionWindows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceLevelObjective) Goal() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"goal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceLevelObjective) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceLevelObjective) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceLevelObjective) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceLevelObjective) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceLevelObjective) RequestBasedSli() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestBasedSli",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceLevelObjective) ServiceLevelObjectiveRef() *interfacesawsapplicationsignals.ServiceLevelObjectiveReference {
	var returns *interfacesawsapplicationsignals.ServiceLevelObjectiveReference
	_jsii_.Get(
		j,
		"serviceLevelObjectiveRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceLevelObjective) Sli() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sli",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceLevelObjective) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceLevelObjective) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceLevelObjective) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceLevelObjective) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::ApplicationSignals::ServiceLevelObjective`.
func NewCfnServiceLevelObjective(scope constructs.Construct, id *string, props *CfnServiceLevelObjectiveProps) CfnServiceLevelObjective {
	_init_.Initialize()

	if err := validateNewCfnServiceLevelObjectiveParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnServiceLevelObjective{}

	_jsii_.Create(
		"aws-cdk-lib.aws_applicationsignals.CfnServiceLevelObjective",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::ApplicationSignals::ServiceLevelObjective`.
func NewCfnServiceLevelObjective_Override(c CfnServiceLevelObjective, scope constructs.Construct, id *string, props *CfnServiceLevelObjectiveProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_applicationsignals.CfnServiceLevelObjective",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnServiceLevelObjective)SetBurnRateConfigurations(val interface{}) {
	if err := j.validateSetBurnRateConfigurationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"burnRateConfigurations",
		val,
	)
}

func (j *jsiiProxy_CfnServiceLevelObjective)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnServiceLevelObjective)SetExclusionWindows(val interface{}) {
	if err := j.validateSetExclusionWindowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exclusionWindows",
		val,
	)
}

func (j *jsiiProxy_CfnServiceLevelObjective)SetGoal(val interface{}) {
	if err := j.validateSetGoalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"goal",
		val,
	)
}

func (j *jsiiProxy_CfnServiceLevelObjective)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnServiceLevelObjective)SetRequestBasedSli(val interface{}) {
	if err := j.validateSetRequestBasedSliParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestBasedSli",
		val,
	)
}

func (j *jsiiProxy_CfnServiceLevelObjective)SetSli(val interface{}) {
	if err := j.validateSetSliParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sli",
		val,
	)
}

func (j *jsiiProxy_CfnServiceLevelObjective)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func CfnServiceLevelObjective_ArnForServiceLevelObjective(resource interfacesawsapplicationsignals.IServiceLevelObjectiveRef) *string {
	_init_.Initialize()

	if err := validateCfnServiceLevelObjective_ArnForServiceLevelObjectiveParameters(resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_applicationsignals.CfnServiceLevelObjective",
		"arnForServiceLevelObjective",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnServiceLevelObjective_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnServiceLevelObjective_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_applicationsignals.CfnServiceLevelObjective",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnServiceLevelObjective_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnServiceLevelObjective_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_applicationsignals.CfnServiceLevelObjective",
		"isCfnResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnServiceLevelObjective.
func CfnServiceLevelObjective_IsCfnServiceLevelObjective(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnServiceLevelObjective_IsCfnServiceLevelObjectiveParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_applicationsignals.CfnServiceLevelObjective",
		"isCfnServiceLevelObjective",
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
func CfnServiceLevelObjective_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnServiceLevelObjective_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_applicationsignals.CfnServiceLevelObjective",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnServiceLevelObjective_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_applicationsignals.CfnServiceLevelObjective",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnServiceLevelObjective) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnServiceLevelObjective) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnServiceLevelObjective) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnServiceLevelObjective) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnServiceLevelObjective) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnServiceLevelObjective) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnServiceLevelObjective) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnServiceLevelObjective) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnServiceLevelObjective) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnServiceLevelObjective) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnServiceLevelObjective) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnServiceLevelObjective) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnServiceLevelObjective) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnServiceLevelObjective) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnServiceLevelObjective) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnServiceLevelObjective) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnServiceLevelObjective) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnServiceLevelObjective) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnServiceLevelObjective) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnServiceLevelObjective) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

func (c *jsiiProxy_CfnServiceLevelObjective) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"with",
		args,
		&returns,
	)

	return returns
}

