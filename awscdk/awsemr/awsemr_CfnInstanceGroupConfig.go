package awsemr

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsemr/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::EMR::InstanceGroupConfig`.
//
// Use `InstanceGroupConfig` to define instance groups for an EMR cluster. A cluster can not use both instance groups and instance fleets. For more information, see [Create a Cluster with Instance Fleets or Uniform Instance Groups](https://docs.aws.amazon.com//emr/latest/ManagementGuide/emr-instance-group-configuration.html) in the *Amazon EMR Management Guide* .
//
// > You can currently only add task instance groups to a cluster with this resource. If you use this resource, CloudFormation waits for the cluster launch to complete before adding the task instance group to the cluster. In order to add task instance groups to the cluster as part of the cluster launch and minimize delays in provisioning task nodes, use the `TaskInstanceGroups` subproperty for the [AWS::EMR::Cluster JobFlowInstancesConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-jobflowinstancesconfig.html) property instead. To use this subproperty, see [AWS::EMR::Cluster](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-cluster.html) for examples.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configurationProperty_ configurationProperty
//
//   cfnInstanceGroupConfig := awscdk.Aws_emr.NewCfnInstanceGroupConfig(this, jsii.String("MyCfnInstanceGroupConfig"), &CfnInstanceGroupConfigProps{
//   	InstanceCount: jsii.Number(123),
//   	InstanceRole: jsii.String("instanceRole"),
//   	InstanceType: jsii.String("instanceType"),
//   	JobFlowId: jsii.String("jobFlowId"),
//
//   	// the properties below are optional
//   	AutoScalingPolicy: &AutoScalingPolicyProperty{
//   		Constraints: &ScalingConstraintsProperty{
//   			MaxCapacity: jsii.Number(123),
//   			MinCapacity: jsii.Number(123),
//   		},
//   		Rules: []interface{}{
//   			&ScalingRuleProperty{
//   				Action: &ScalingActionProperty{
//   					SimpleScalingPolicyConfiguration: &SimpleScalingPolicyConfigurationProperty{
//   						ScalingAdjustment: jsii.Number(123),
//
//   						// the properties below are optional
//   						AdjustmentType: jsii.String("adjustmentType"),
//   						CoolDown: jsii.Number(123),
//   					},
//
//   					// the properties below are optional
//   					Market: jsii.String("market"),
//   				},
//   				Name: jsii.String("name"),
//   				Trigger: &ScalingTriggerProperty{
//   					CloudWatchAlarmDefinition: &CloudWatchAlarmDefinitionProperty{
//   						ComparisonOperator: jsii.String("comparisonOperator"),
//   						MetricName: jsii.String("metricName"),
//   						Period: jsii.Number(123),
//   						Threshold: jsii.Number(123),
//
//   						// the properties below are optional
//   						Dimensions: []interface{}{
//   							&MetricDimensionProperty{
//   								Key: jsii.String("key"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						EvaluationPeriods: jsii.Number(123),
//   						Namespace: jsii.String("namespace"),
//   						Statistic: jsii.String("statistic"),
//   						Unit: jsii.String("unit"),
//   					},
//   				},
//
//   				// the properties below are optional
//   				Description: jsii.String("description"),
//   			},
//   		},
//   	},
//   	BidPrice: jsii.String("bidPrice"),
//   	Configurations: []interface{}{
//   		&configurationProperty{
//   			Classification: jsii.String("classification"),
//   			ConfigurationProperties: map[string]*string{
//   				"configurationPropertiesKey": jsii.String("configurationProperties"),
//   			},
//   			Configurations: []interface{}{
//   				configurationProperty_,
//   			},
//   		},
//   	},
//   	CustomAmiId: jsii.String("customAmiId"),
//   	EbsConfiguration: &EbsConfigurationProperty{
//   		EbsBlockDeviceConfigs: []interface{}{
//   			&EbsBlockDeviceConfigProperty{
//   				VolumeSpecification: &VolumeSpecificationProperty{
//   					SizeInGb: jsii.Number(123),
//   					VolumeType: jsii.String("volumeType"),
//
//   					// the properties below are optional
//   					Iops: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				VolumesPerInstance: jsii.Number(123),
//   			},
//   		},
//   		EbsOptimized: jsii.Boolean(false),
//   	},
//   	Market: jsii.String("market"),
//   	Name: jsii.String("name"),
//   })
//
type CfnInstanceGroupConfig interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// `AutoScalingPolicy` is a subproperty of `InstanceGroupConfig` .
	//
	// `AutoScalingPolicy` defines how an instance group dynamically adds and terminates EC2 instances in response to the value of a CloudWatch metric. For more information, see [Using Automatic Scaling in Amazon EMR](https://docs.aws.amazon.com//emr/latest/ManagementGuide/emr-automatic-scaling.html) in the *Amazon EMR Management Guide* .
	AutoScalingPolicy() interface{}
	SetAutoScalingPolicy(val interface{})
	// If specified, indicates that the instance group uses Spot Instances.
	//
	// This is the maximum price you are willing to pay for Spot Instances. Specify `OnDemandPrice` to set the amount equal to the On-Demand price, or specify an amount in USD.
	BidPrice() *string
	SetBidPrice(val *string)
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// > Amazon EMR releases 4.x or later.
	//
	// The list of configurations supplied for an EMR cluster instance group. You can specify a separate configuration for each instance group (master, core, and task).
	Configurations() interface{}
	SetConfigurations(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The custom AMI ID to use for the provisioned instance group.
	CustomAmiId() *string
	SetCustomAmiId(val *string)
	// `EbsConfiguration` determines the EBS volumes to attach to EMR cluster instances.
	EbsConfiguration() interface{}
	SetEbsConfiguration(val interface{})
	// Target number of instances for the instance group.
	InstanceCount() *float64
	SetInstanceCount(val *float64)
	// The role of the instance group in the cluster.
	//
	// *Allowed Values* : TASK.
	InstanceRole() *string
	SetInstanceRole(val *string)
	// The EC2 instance type for all instances in the instance group.
	InstanceType() *string
	SetInstanceType(val *string)
	// The ID of an Amazon EMR cluster that you want to associate this instance group with.
	JobFlowId() *string
	SetJobFlowId(val *string)
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
	// Market type of the EC2 instances used to create a cluster node.
	Market() *string
	SetMarket(val *string)
	// Friendly name given to the instance group.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
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

// The jsii proxy struct for CfnInstanceGroupConfig
type jsiiProxy_CfnInstanceGroupConfig struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnInstanceGroupConfig) AutoScalingPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoScalingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) BidPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bidPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) Configurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) CustomAmiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customAmiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) EbsConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) InstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) InstanceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) JobFlowId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobFlowId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) Market() *string {
	var returns *string
	_jsii_.Get(
		j,
		"market",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceGroupConfig) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::EMR::InstanceGroupConfig`.
func NewCfnInstanceGroupConfig(scope awscdk.Construct, id *string, props *CfnInstanceGroupConfigProps) CfnInstanceGroupConfig {
	_init_.Initialize()

	if err := validateNewCfnInstanceGroupConfigParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnInstanceGroupConfig{}

	_jsii_.Create(
		"monocdk.aws_emr.CfnInstanceGroupConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::EMR::InstanceGroupConfig`.
func NewCfnInstanceGroupConfig_Override(c CfnInstanceGroupConfig, scope awscdk.Construct, id *string, props *CfnInstanceGroupConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_emr.CfnInstanceGroupConfig",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnInstanceGroupConfig)SetAutoScalingPolicy(val interface{}) {
	if err := j.validateSetAutoScalingPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoScalingPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceGroupConfig)SetBidPrice(val *string) {
	_jsii_.Set(
		j,
		"bidPrice",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceGroupConfig)SetConfigurations(val interface{}) {
	if err := j.validateSetConfigurationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurations",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceGroupConfig)SetCustomAmiId(val *string) {
	_jsii_.Set(
		j,
		"customAmiId",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceGroupConfig)SetEbsConfiguration(val interface{}) {
	if err := j.validateSetEbsConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceGroupConfig)SetInstanceCount(val *float64) {
	if err := j.validateSetInstanceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceCount",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceGroupConfig)SetInstanceRole(val *string) {
	if err := j.validateSetInstanceRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceRole",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceGroupConfig)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceGroupConfig)SetJobFlowId(val *string) {
	if err := j.validateSetJobFlowIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobFlowId",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceGroupConfig)SetMarket(val *string) {
	_jsii_.Set(
		j,
		"market",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceGroupConfig)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
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
func CfnInstanceGroupConfig_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInstanceGroupConfig_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_emr.CfnInstanceGroupConfig",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnInstanceGroupConfig_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnInstanceGroupConfig_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_emr.CfnInstanceGroupConfig",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnInstanceGroupConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInstanceGroupConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_emr.CfnInstanceGroupConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInstanceGroupConfig_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_emr.CfnInstanceGroupConfig",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnInstanceGroupConfig) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnInstanceGroupConfig) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnInstanceGroupConfig) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnInstanceGroupConfig) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnInstanceGroupConfig) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnInstanceGroupConfig) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnInstanceGroupConfig) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnInstanceGroupConfig) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInstanceGroupConfig) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnInstanceGroupConfig) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnInstanceGroupConfig) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnInstanceGroupConfig) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnInstanceGroupConfig) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInstanceGroupConfig) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnInstanceGroupConfig) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnInstanceGroupConfig) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnInstanceGroupConfig) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInstanceGroupConfig) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnInstanceGroupConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInstanceGroupConfig) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInstanceGroupConfig) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

