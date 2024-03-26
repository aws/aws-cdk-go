package awsautoscaling

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::AutoScaling::AutoScalingGroup` resource defines an Amazon EC2 Auto Scaling group, which is a collection of Amazon EC2 instances that are treated as a logical grouping for the purposes of automatic scaling and management.
//
// For more information about Amazon EC2 Auto Scaling, see the [Amazon EC2 Auto Scaling User Guide](https://docs.aws.amazon.com/autoscaling/ec2/userguide/what-is-amazon-ec2-auto-scaling.html) .
//
// > Amazon EC2 Auto Scaling configures instances launched as part of an Auto Scaling group using either a [launch template](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-launchtemplate.html) or a launch configuration. We strongly recommend that you do not use launch configurations. For more information, see [Launch configurations](https://docs.aws.amazon.com/autoscaling/ec2/userguide/launch-configurations.html) in the *Amazon EC2 Auto Scaling User Guide* .
// >
// > For help migrating from launch configurations to launch templates, see [Migrate AWS CloudFormation stacks from launch configurations to launch templates](https://docs.aws.amazon.com/autoscaling/ec2/userguide/migrate-launch-configurations-with-cloudformation.html) in the *Amazon EC2 Auto Scaling User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAutoScalingGroup := awscdk.Aws_autoscaling.NewCfnAutoScalingGroup(this, jsii.String("MyCfnAutoScalingGroup"), &CfnAutoScalingGroupProps{
//   	MaxSize: jsii.String("maxSize"),
//   	MinSize: jsii.String("minSize"),
//
//   	// the properties below are optional
//   	AutoScalingGroupName: jsii.String("autoScalingGroupName"),
//   	AvailabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	CapacityRebalance: jsii.Boolean(false),
//   	Context: jsii.String("context"),
//   	Cooldown: jsii.String("cooldown"),
//   	DefaultInstanceWarmup: jsii.Number(123),
//   	DesiredCapacity: jsii.String("desiredCapacity"),
//   	DesiredCapacityType: jsii.String("desiredCapacityType"),
//   	HealthCheckGracePeriod: jsii.Number(123),
//   	HealthCheckType: jsii.String("healthCheckType"),
//   	InstanceId: jsii.String("instanceId"),
//   	InstanceMaintenancePolicy: &InstanceMaintenancePolicyProperty{
//   		MaxHealthyPercentage: jsii.Number(123),
//   		MinHealthyPercentage: jsii.Number(123),
//   	},
//   	LaunchConfigurationName: jsii.String("launchConfigurationName"),
//   	LaunchTemplate: &LaunchTemplateSpecificationProperty{
//   		Version: jsii.String("version"),
//
//   		// the properties below are optional
//   		LaunchTemplateId: jsii.String("launchTemplateId"),
//   		LaunchTemplateName: jsii.String("launchTemplateName"),
//   	},
//   	LifecycleHookSpecificationList: []interface{}{
//   		&LifecycleHookSpecificationProperty{
//   			LifecycleHookName: jsii.String("lifecycleHookName"),
//   			LifecycleTransition: jsii.String("lifecycleTransition"),
//
//   			// the properties below are optional
//   			DefaultResult: jsii.String("defaultResult"),
//   			HeartbeatTimeout: jsii.Number(123),
//   			NotificationMetadata: jsii.String("notificationMetadata"),
//   			NotificationTargetArn: jsii.String("notificationTargetArn"),
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   	},
//   	LoadBalancerNames: []*string{
//   		jsii.String("loadBalancerNames"),
//   	},
//   	MaxInstanceLifetime: jsii.Number(123),
//   	MetricsCollection: []interface{}{
//   		&MetricsCollectionProperty{
//   			Granularity: jsii.String("granularity"),
//
//   			// the properties below are optional
//   			Metrics: []*string{
//   				jsii.String("metrics"),
//   			},
//   		},
//   	},
//   	MixedInstancesPolicy: &MixedInstancesPolicyProperty{
//   		LaunchTemplate: &LaunchTemplateProperty{
//   			LaunchTemplateSpecification: &LaunchTemplateSpecificationProperty{
//   				Version: jsii.String("version"),
//
//   				// the properties below are optional
//   				LaunchTemplateId: jsii.String("launchTemplateId"),
//   				LaunchTemplateName: jsii.String("launchTemplateName"),
//   			},
//
//   			// the properties below are optional
//   			Overrides: []interface{}{
//   				&LaunchTemplateOverridesProperty{
//   					InstanceRequirements: &InstanceRequirementsProperty{
//   						MemoryMiB: &MemoryMiBRequestProperty{
//   							Max: jsii.Number(123),
//   							Min: jsii.Number(123),
//   						},
//   						VCpuCount: &VCpuCountRequestProperty{
//   							Max: jsii.Number(123),
//   							Min: jsii.Number(123),
//   						},
//
//   						// the properties below are optional
//   						AcceleratorCount: &AcceleratorCountRequestProperty{
//   							Max: jsii.Number(123),
//   							Min: jsii.Number(123),
//   						},
//   						AcceleratorManufacturers: []*string{
//   							jsii.String("acceleratorManufacturers"),
//   						},
//   						AcceleratorNames: []*string{
//   							jsii.String("acceleratorNames"),
//   						},
//   						AcceleratorTotalMemoryMiB: &AcceleratorTotalMemoryMiBRequestProperty{
//   							Max: jsii.Number(123),
//   							Min: jsii.Number(123),
//   						},
//   						AcceleratorTypes: []*string{
//   							jsii.String("acceleratorTypes"),
//   						},
//   						AllowedInstanceTypes: []*string{
//   							jsii.String("allowedInstanceTypes"),
//   						},
//   						BareMetal: jsii.String("bareMetal"),
//   						BaselineEbsBandwidthMbps: &BaselineEbsBandwidthMbpsRequestProperty{
//   							Max: jsii.Number(123),
//   							Min: jsii.Number(123),
//   						},
//   						BurstablePerformance: jsii.String("burstablePerformance"),
//   						CpuManufacturers: []*string{
//   							jsii.String("cpuManufacturers"),
//   						},
//   						ExcludedInstanceTypes: []*string{
//   							jsii.String("excludedInstanceTypes"),
//   						},
//   						InstanceGenerations: []*string{
//   							jsii.String("instanceGenerations"),
//   						},
//   						LocalStorage: jsii.String("localStorage"),
//   						LocalStorageTypes: []*string{
//   							jsii.String("localStorageTypes"),
//   						},
//   						MaxSpotPriceAsPercentageOfOptimalOnDemandPrice: jsii.Number(123),
//   						MemoryGiBPerVCpu: &MemoryGiBPerVCpuRequestProperty{
//   							Max: jsii.Number(123),
//   							Min: jsii.Number(123),
//   						},
//   						NetworkBandwidthGbps: &NetworkBandwidthGbpsRequestProperty{
//   							Max: jsii.Number(123),
//   							Min: jsii.Number(123),
//   						},
//   						NetworkInterfaceCount: &NetworkInterfaceCountRequestProperty{
//   							Max: jsii.Number(123),
//   							Min: jsii.Number(123),
//   						},
//   						OnDemandMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   						RequireHibernateSupport: jsii.Boolean(false),
//   						SpotMaxPricePercentageOverLowestPrice: jsii.Number(123),
//   						TotalLocalStorageGb: &TotalLocalStorageGBRequestProperty{
//   							Max: jsii.Number(123),
//   							Min: jsii.Number(123),
//   						},
//   					},
//   					InstanceType: jsii.String("instanceType"),
//   					LaunchTemplateSpecification: &LaunchTemplateSpecificationProperty{
//   						Version: jsii.String("version"),
//
//   						// the properties below are optional
//   						LaunchTemplateId: jsii.String("launchTemplateId"),
//   						LaunchTemplateName: jsii.String("launchTemplateName"),
//   					},
//   					WeightedCapacity: jsii.String("weightedCapacity"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		InstancesDistribution: &InstancesDistributionProperty{
//   			OnDemandAllocationStrategy: jsii.String("onDemandAllocationStrategy"),
//   			OnDemandBaseCapacity: jsii.Number(123),
//   			OnDemandPercentageAboveBaseCapacity: jsii.Number(123),
//   			SpotAllocationStrategy: jsii.String("spotAllocationStrategy"),
//   			SpotInstancePools: jsii.Number(123),
//   			SpotMaxPrice: jsii.String("spotMaxPrice"),
//   		},
//   	},
//   	NewInstancesProtectedFromScaleIn: jsii.Boolean(false),
//   	NotificationConfiguration: &NotificationConfigurationProperty{
//   		TopicArn: jsii.String("topicArn"),
//
//   		// the properties below are optional
//   		NotificationTypes: []*string{
//   			jsii.String("notificationTypes"),
//   		},
//   	},
//   	NotificationConfigurations: []interface{}{
//   		&NotificationConfigurationProperty{
//   			TopicArn: jsii.String("topicArn"),
//
//   			// the properties below are optional
//   			NotificationTypes: []*string{
//   				jsii.String("notificationTypes"),
//   			},
//   		},
//   	},
//   	PlacementGroup: jsii.String("placementGroup"),
//   	ServiceLinkedRoleArn: jsii.String("serviceLinkedRoleArn"),
//   	Tags: []tagPropertyProperty{
//   		&tagPropertyProperty{
//   			Key: jsii.String("key"),
//   			PropagateAtLaunch: jsii.Boolean(false),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetGroupArns: []*string{
//   		jsii.String("targetGroupArns"),
//   	},
//   	TerminationPolicies: []*string{
//   		jsii.String("terminationPolicies"),
//   	},
//   	VpcZoneIdentifier: []*string{
//   		jsii.String("vpcZoneIdentifier"),
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html
//
type CfnAutoScalingGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggable
	// The name of the Auto Scaling group.
	//
	// This name must be unique per Region per account.
	AutoScalingGroupName() *string
	SetAutoScalingGroupName(val *string)
	// A list of Availability Zones where instances in the Auto Scaling group can be created.
	AvailabilityZones() *[]*string
	SetAvailabilityZones(val *[]*string)
	// Indicates whether Capacity Rebalancing is enabled.
	CapacityRebalance() interface{}
	SetCapacityRebalance(val interface{})
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Reserved.
	Context() *string
	SetContext(val *string)
	// *Only needed if you use simple scaling policies.*.
	Cooldown() *string
	SetCooldown(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The amount of time, in seconds, until a new instance is considered to have finished initializing and resource consumption to become stable after it enters the `InService` state.
	DefaultInstanceWarmup() *float64
	SetDefaultInstanceWarmup(val *float64)
	// The desired capacity is the initial capacity of the Auto Scaling group at the time of its creation and the capacity it attempts to maintain.
	DesiredCapacity() *string
	SetDesiredCapacity(val *string)
	// The unit of measurement for the value specified for desired capacity.
	DesiredCapacityType() *string
	SetDesiredCapacityType(val *string)
	// The amount of time, in seconds, that Amazon EC2 Auto Scaling waits before checking the health status of an EC2 instance that has come into service and marking it unhealthy due to a failed health check.
	HealthCheckGracePeriod() *float64
	SetHealthCheckGracePeriod(val *float64)
	// A comma-separated value string of one or more health check types.
	HealthCheckType() *string
	SetHealthCheckType(val *string)
	// The ID of the instance used to base the launch configuration on.
	InstanceId() *string
	SetInstanceId(val *string)
	// An instance maintenance policy.
	InstanceMaintenancePolicy() interface{}
	SetInstanceMaintenancePolicy(val interface{})
	// The name of the launch configuration to use to launch instances.
	LaunchConfigurationName() *string
	SetLaunchConfigurationName(val *string)
	// Information used to specify the launch template and version to use to launch instances.
	LaunchTemplate() interface{}
	SetLaunchTemplate(val interface{})
	// One or more lifecycle hooks to add to the Auto Scaling group before instances are launched.
	LifecycleHookSpecificationList() interface{}
	SetLifecycleHookSpecificationList(val interface{})
	// A list of Classic Load Balancers associated with this Auto Scaling group.
	LoadBalancerNames() *[]*string
	SetLoadBalancerNames(val *[]*string)
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
	// The maximum amount of time, in seconds, that an instance can be in service.
	MaxInstanceLifetime() *float64
	SetMaxInstanceLifetime(val *float64)
	// The maximum size of the group.
	MaxSize() *string
	SetMaxSize(val *string)
	// Enables the monitoring of group metrics of an Auto Scaling group.
	MetricsCollection() interface{}
	SetMetricsCollection(val interface{})
	// The minimum size of the group.
	MinSize() *string
	SetMinSize(val *string)
	// An embedded object that specifies a mixed instances policy.
	MixedInstancesPolicy() interface{}
	SetMixedInstancesPolicy(val interface{})
	// Indicates whether newly launched instances are protected from termination by Amazon EC2 Auto Scaling when scaling in.
	NewInstancesProtectedFromScaleIn() interface{}
	SetNewInstancesProtectedFromScaleIn(val interface{})
	// The tree node.
	Node() constructs.Node
	// A structure that specifies an Amazon SNS notification configuration for the ``NotificationConfigurations`` property of the [AWS::AutoScaling::AutoScalingGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html) resource.  For an example template snippet, see [Auto scaling template snippets](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/quickref-autoscaling.html).  For more information, see [Get Amazon SNS notifications when your Auto Scaling group scales](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ASGettingNotifications.html) in the *Amazon EC2 Auto Scaling User Guide*.
	// Deprecated: this property has been deprecated.
	NotificationConfiguration() interface{}
	// Deprecated: this property has been deprecated.
	SetNotificationConfiguration(val interface{})
	// Configures an Auto Scaling group to send notifications when specified events take place.
	NotificationConfigurations() interface{}
	SetNotificationConfigurations(val interface{})
	// The name of the placement group into which to launch your instances.
	PlacementGroup() *string
	SetPlacementGroup(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The Amazon Resource Name (ARN) of the service-linked role that the Auto Scaling group uses to call other AWS service on your behalf.
	ServiceLinkedRoleArn() *string
	SetServiceLinkedRoleArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// One or more tags.
	TagsRaw() *[]*CfnAutoScalingGroup_TagPropertyProperty
	SetTagsRaw(val *[]*CfnAutoScalingGroup_TagPropertyProperty)
	// The Amazon Resource Names (ARN) of the Elastic Load Balancing target groups to associate with the Auto Scaling group.
	TargetGroupArns() *[]*string
	SetTargetGroupArns(val *[]*string)
	// A policy or a list of policies that are used to select the instance to terminate.
	TerminationPolicies() *[]*string
	SetTerminationPolicies(val *[]*string)
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
	// A list of subnet IDs for a virtual private cloud (VPC) where instances in the Auto Scaling group can be created.
	VpcZoneIdentifier() *[]*string
	SetVpcZoneIdentifier(val *[]*string)
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

// The jsii proxy struct for CfnAutoScalingGroup
type jsiiProxy_CfnAutoScalingGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
}

func (j *jsiiProxy_CfnAutoScalingGroup) AutoScalingGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) CapacityRebalance() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacityRebalance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) Cooldown() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) DefaultInstanceWarmup() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultInstanceWarmup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) DesiredCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) DesiredCapacityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredCapacityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) HealthCheckGracePeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckGracePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) HealthCheckType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) InstanceMaintenancePolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceMaintenancePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) LaunchConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) LaunchTemplate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"launchTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) LifecycleHookSpecificationList() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lifecycleHookSpecificationList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) LoadBalancerNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancerNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) MaxInstanceLifetime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstanceLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) MaxSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) MetricsCollection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricsCollection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) MinSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) MixedInstancesPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mixedInstancesPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) NewInstancesProtectedFromScaleIn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"newInstancesProtectedFromScaleIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) NotificationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) NotificationConfigurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) PlacementGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) ServiceLinkedRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceLinkedRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) TagsRaw() *[]*CfnAutoScalingGroup_TagPropertyProperty {
	var returns *[]*CfnAutoScalingGroup_TagPropertyProperty
	_jsii_.Get(
		j,
		"tagsRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) TargetGroupArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetGroupArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) TerminationPolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"terminationPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAutoScalingGroup) VpcZoneIdentifier() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcZoneIdentifier",
		&returns,
	)
	return returns
}


func NewCfnAutoScalingGroup(scope constructs.Construct, id *string, props *CfnAutoScalingGroupProps) CfnAutoScalingGroup {
	_init_.Initialize()

	if err := validateNewCfnAutoScalingGroupParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAutoScalingGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.CfnAutoScalingGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnAutoScalingGroup_Override(c CfnAutoScalingGroup, scope constructs.Construct, id *string, props *CfnAutoScalingGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_autoscaling.CfnAutoScalingGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup)SetAutoScalingGroupName(val *string) {
	_jsii_.Set(
		j,
		"autoScalingGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup)SetAvailabilityZones(val *[]*string) {
	_jsii_.Set(
		j,
		"availabilityZones",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup)SetCapacityRebalance(val interface{}) {
	if err := j.validateSetCapacityRebalanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityRebalance",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup)SetContext(val *string) {
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup)SetCooldown(val *string) {
	_jsii_.Set(
		j,
		"cooldown",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup)SetDefaultInstanceWarmup(val *float64) {
	_jsii_.Set(
		j,
		"defaultInstanceWarmup",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup)SetDesiredCapacity(val *string) {
	_jsii_.Set(
		j,
		"desiredCapacity",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup)SetDesiredCapacityType(val *string) {
	_jsii_.Set(
		j,
		"desiredCapacityType",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup)SetHealthCheckGracePeriod(val *float64) {
	_jsii_.Set(
		j,
		"healthCheckGracePeriod",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup)SetHealthCheckType(val *string) {
	_jsii_.Set(
		j,
		"healthCheckType",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup)SetInstanceId(val *string) {
	_jsii_.Set(
		j,
		"instanceId",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup)SetInstanceMaintenancePolicy(val interface{}) {
	if err := j.validateSetInstanceMaintenancePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceMaintenancePolicy",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup)SetLaunchConfigurationName(val *string) {
	_jsii_.Set(
		j,
		"launchConfigurationName",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup)SetLaunchTemplate(val interface{}) {
	if err := j.validateSetLaunchTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchTemplate",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup)SetLifecycleHookSpecificationList(val interface{}) {
	if err := j.validateSetLifecycleHookSpecificationListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycleHookSpecificationList",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup)SetLoadBalancerNames(val *[]*string) {
	_jsii_.Set(
		j,
		"loadBalancerNames",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup)SetMaxInstanceLifetime(val *float64) {
	_jsii_.Set(
		j,
		"maxInstanceLifetime",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup)SetMaxSize(val *string) {
	if err := j.validateSetMaxSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSize",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup)SetMetricsCollection(val interface{}) {
	if err := j.validateSetMetricsCollectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsCollection",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup)SetMinSize(val *string) {
	if err := j.validateSetMinSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minSize",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup)SetMixedInstancesPolicy(val interface{}) {
	if err := j.validateSetMixedInstancesPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mixedInstancesPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup)SetNewInstancesProtectedFromScaleIn(val interface{}) {
	if err := j.validateSetNewInstancesProtectedFromScaleInParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newInstancesProtectedFromScaleIn",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup)SetNotificationConfiguration(val interface{}) {
	if err := j.validateSetNotificationConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup)SetNotificationConfigurations(val interface{}) {
	if err := j.validateSetNotificationConfigurationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationConfigurations",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup)SetPlacementGroup(val *string) {
	_jsii_.Set(
		j,
		"placementGroup",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup)SetServiceLinkedRoleArn(val *string) {
	_jsii_.Set(
		j,
		"serviceLinkedRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup)SetTagsRaw(val *[]*CfnAutoScalingGroup_TagPropertyProperty) {
	if err := j.validateSetTagsRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup)SetTargetGroupArns(val *[]*string) {
	_jsii_.Set(
		j,
		"targetGroupArns",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup)SetTerminationPolicies(val *[]*string) {
	_jsii_.Set(
		j,
		"terminationPolicies",
		val,
	)
}

func (j *jsiiProxy_CfnAutoScalingGroup)SetVpcZoneIdentifier(val *[]*string) {
	_jsii_.Set(
		j,
		"vpcZoneIdentifier",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnAutoScalingGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAutoScalingGroup_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.CfnAutoScalingGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnAutoScalingGroup_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAutoScalingGroup_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.CfnAutoScalingGroup",
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
func CfnAutoScalingGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAutoScalingGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_autoscaling.CfnAutoScalingGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAutoScalingGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_autoscaling.CfnAutoScalingGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAutoScalingGroup) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnAutoScalingGroup) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAutoScalingGroup) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAutoScalingGroup) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnAutoScalingGroup) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnAutoScalingGroup) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnAutoScalingGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnAutoScalingGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnAutoScalingGroup) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnAutoScalingGroup) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnAutoScalingGroup) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnAutoScalingGroup) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAutoScalingGroup) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAutoScalingGroup) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnAutoScalingGroup) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnAutoScalingGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnAutoScalingGroup) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnAutoScalingGroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAutoScalingGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAutoScalingGroup) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

