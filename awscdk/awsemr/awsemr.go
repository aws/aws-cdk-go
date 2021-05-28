package awsemr

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsemr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::EMR::Cluster`.
type CfnCluster interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AdditionalInfo() interface{}
	SetAdditionalInfo(val interface{})
	Applications() interface{}
	SetApplications(val interface{})
	AttrMasterPublicDns() *string
	AutoScalingRole() *string
	SetAutoScalingRole(val *string)
	BootstrapActions() interface{}
	SetBootstrapActions(val interface{})
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	Configurations() interface{}
	SetConfigurations(val interface{})
	CreationStack() *[]*string
	CustomAmiId() *string
	SetCustomAmiId(val *string)
	EbsRootVolumeSize() *float64
	SetEbsRootVolumeSize(val *float64)
	Instances() interface{}
	SetInstances(val interface{})
	JobFlowRole() *string
	SetJobFlowRole(val *string)
	KerberosAttributes() interface{}
	SetKerberosAttributes(val interface{})
	LogEncryptionKmsKeyId() *string
	SetLogEncryptionKmsKeyId(val *string)
	LogicalId() *string
	LogUri() *string
	SetLogUri(val *string)
	ManagedScalingPolicy() interface{}
	SetManagedScalingPolicy(val interface{})
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	ReleaseLabel() *string
	SetReleaseLabel(val *string)
	ScaleDownBehavior() *string
	SetScaleDownBehavior(val *string)
	SecurityConfiguration() *string
	SetSecurityConfiguration(val *string)
	ServiceRole() *string
	SetServiceRole(val *string)
	Stack() awscdk.Stack
	StepConcurrencyLevel() *float64
	SetStepConcurrencyLevel(val *float64)
	Steps() interface{}
	SetSteps(val interface{})
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	VisibleToAllUsers() interface{}
	SetVisibleToAllUsers(val interface{})
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

// The jsii proxy struct for CfnCluster
type jsiiProxy_CfnCluster struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnCluster) AdditionalInfo() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Applications() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) AttrMasterPublicDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrMasterPublicDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) AutoScalingRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) BootstrapActions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bootstrapActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Configurations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) CustomAmiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customAmiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) EbsRootVolumeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ebsRootVolumeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Instances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) JobFlowRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobFlowRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) KerberosAttributes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberosAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) LogEncryptionKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logEncryptionKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) LogUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) ManagedScalingPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedScalingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) ReleaseLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) ScaleDownBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) SecurityConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) ServiceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) StepConcurrencyLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stepConcurrencyLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Steps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"steps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCluster) VisibleToAllUsers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"visibleToAllUsers",
		&returns,
	)
	return returns
}


// Create a new `AWS::EMR::Cluster`.
func NewCfnCluster(scope constructs.Construct, id *string, props *CfnClusterProps) CfnCluster {
	_init_.Initialize()

	j := jsiiProxy_CfnCluster{}

	_jsii_.Create(
		"aws-cdk-lib.aws_emr.CfnCluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::EMR::Cluster`.
func NewCfnCluster_Override(c CfnCluster, scope constructs.Construct, id *string, props *CfnClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_emr.CfnCluster",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnCluster) SetAdditionalInfo(val interface{}) {
	_jsii_.Set(
		j,
		"additionalInfo",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetApplications(val interface{}) {
	_jsii_.Set(
		j,
		"applications",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetAutoScalingRole(val *string) {
	_jsii_.Set(
		j,
		"autoScalingRole",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetBootstrapActions(val interface{}) {
	_jsii_.Set(
		j,
		"bootstrapActions",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetConfigurations(val interface{}) {
	_jsii_.Set(
		j,
		"configurations",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetCustomAmiId(val *string) {
	_jsii_.Set(
		j,
		"customAmiId",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetEbsRootVolumeSize(val *float64) {
	_jsii_.Set(
		j,
		"ebsRootVolumeSize",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetInstances(val interface{}) {
	_jsii_.Set(
		j,
		"instances",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetJobFlowRole(val *string) {
	_jsii_.Set(
		j,
		"jobFlowRole",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetKerberosAttributes(val interface{}) {
	_jsii_.Set(
		j,
		"kerberosAttributes",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetLogEncryptionKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"logEncryptionKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetLogUri(val *string) {
	_jsii_.Set(
		j,
		"logUri",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetManagedScalingPolicy(val interface{}) {
	_jsii_.Set(
		j,
		"managedScalingPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetReleaseLabel(val *string) {
	_jsii_.Set(
		j,
		"releaseLabel",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetScaleDownBehavior(val *string) {
	_jsii_.Set(
		j,
		"scaleDownBehavior",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetSecurityConfiguration(val *string) {
	_jsii_.Set(
		j,
		"securityConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetServiceRole(val *string) {
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetStepConcurrencyLevel(val *float64) {
	_jsii_.Set(
		j,
		"stepConcurrencyLevel",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetSteps(val interface{}) {
	_jsii_.Set(
		j,
		"steps",
		val,
	)
}

func (j *jsiiProxy_CfnCluster) SetVisibleToAllUsers(val interface{}) {
	_jsii_.Set(
		j,
		"visibleToAllUsers",
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
func CfnCluster_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnCluster",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnCluster_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnCluster",
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
func CfnCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCluster_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_emr.CfnCluster",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnCluster) AddDeletionOverride(path *string) {
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
// Experimental.
func (c *jsiiProxy_CfnCluster) AddDependsOn(target awscdk.CfnResource) {
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
// Experimental.
func (c *jsiiProxy_CfnCluster) AddMetadata(key *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnCluster) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnCluster) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnCluster) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnCluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
// Experimental.
func (c *jsiiProxy_CfnCluster) GetAtt(attributeName *string) awscdk.Reference {
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
// Experimental.
func (c *jsiiProxy_CfnCluster) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnCluster) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnCluster) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnCluster) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
// Experimental.
func (c *jsiiProxy_CfnCluster) ShouldSynthesize() *bool {
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
// Experimental.
func (c *jsiiProxy_CfnCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnCluster) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnCluster_ApplicationProperty struct {
	// `CfnCluster.ApplicationProperty.AdditionalInfo`.
	AdditionalInfo interface{} `json:"additionalInfo"`
	// `CfnCluster.ApplicationProperty.Args`.
	Args *[]*string `json:"args"`
	// `CfnCluster.ApplicationProperty.Name`.
	Name *string `json:"name"`
	// `CfnCluster.ApplicationProperty.Version`.
	Version *string `json:"version"`
}

type CfnCluster_AutoScalingPolicyProperty struct {
	// `CfnCluster.AutoScalingPolicyProperty.Constraints`.
	Constraints interface{} `json:"constraints"`
	// `CfnCluster.AutoScalingPolicyProperty.Rules`.
	Rules interface{} `json:"rules"`
}

type CfnCluster_BootstrapActionConfigProperty struct {
	// `CfnCluster.BootstrapActionConfigProperty.Name`.
	Name *string `json:"name"`
	// `CfnCluster.BootstrapActionConfigProperty.ScriptBootstrapAction`.
	ScriptBootstrapAction interface{} `json:"scriptBootstrapAction"`
}

type CfnCluster_CloudWatchAlarmDefinitionProperty struct {
	// `CfnCluster.CloudWatchAlarmDefinitionProperty.ComparisonOperator`.
	ComparisonOperator *string `json:"comparisonOperator"`
	// `CfnCluster.CloudWatchAlarmDefinitionProperty.MetricName`.
	MetricName *string `json:"metricName"`
	// `CfnCluster.CloudWatchAlarmDefinitionProperty.Period`.
	Period *float64 `json:"period"`
	// `CfnCluster.CloudWatchAlarmDefinitionProperty.Threshold`.
	Threshold *float64 `json:"threshold"`
	// `CfnCluster.CloudWatchAlarmDefinitionProperty.Dimensions`.
	Dimensions interface{} `json:"dimensions"`
	// `CfnCluster.CloudWatchAlarmDefinitionProperty.EvaluationPeriods`.
	EvaluationPeriods *float64 `json:"evaluationPeriods"`
	// `CfnCluster.CloudWatchAlarmDefinitionProperty.Namespace`.
	Namespace *string `json:"namespace"`
	// `CfnCluster.CloudWatchAlarmDefinitionProperty.Statistic`.
	Statistic *string `json:"statistic"`
	// `CfnCluster.CloudWatchAlarmDefinitionProperty.Unit`.
	Unit *string `json:"unit"`
}

type CfnCluster_ComputeLimitsProperty struct {
	// `CfnCluster.ComputeLimitsProperty.MaximumCapacityUnits`.
	MaximumCapacityUnits *float64 `json:"maximumCapacityUnits"`
	// `CfnCluster.ComputeLimitsProperty.MinimumCapacityUnits`.
	MinimumCapacityUnits *float64 `json:"minimumCapacityUnits"`
	// `CfnCluster.ComputeLimitsProperty.UnitType`.
	UnitType *string `json:"unitType"`
	// `CfnCluster.ComputeLimitsProperty.MaximumCoreCapacityUnits`.
	MaximumCoreCapacityUnits *float64 `json:"maximumCoreCapacityUnits"`
	// `CfnCluster.ComputeLimitsProperty.MaximumOnDemandCapacityUnits`.
	MaximumOnDemandCapacityUnits *float64 `json:"maximumOnDemandCapacityUnits"`
}

type CfnCluster_ConfigurationProperty struct {
	// `CfnCluster.ConfigurationProperty.Classification`.
	Classification *string `json:"classification"`
	// `CfnCluster.ConfigurationProperty.ConfigurationProperties`.
	ConfigurationProperties interface{} `json:"configurationProperties"`
	// `CfnCluster.ConfigurationProperty.Configurations`.
	Configurations interface{} `json:"configurations"`
}

type CfnCluster_EbsBlockDeviceConfigProperty struct {
	// `CfnCluster.EbsBlockDeviceConfigProperty.VolumeSpecification`.
	VolumeSpecification interface{} `json:"volumeSpecification"`
	// `CfnCluster.EbsBlockDeviceConfigProperty.VolumesPerInstance`.
	VolumesPerInstance *float64 `json:"volumesPerInstance"`
}

type CfnCluster_EbsConfigurationProperty struct {
	// `CfnCluster.EbsConfigurationProperty.EbsBlockDeviceConfigs`.
	EbsBlockDeviceConfigs interface{} `json:"ebsBlockDeviceConfigs"`
	// `CfnCluster.EbsConfigurationProperty.EbsOptimized`.
	EbsOptimized interface{} `json:"ebsOptimized"`
}

type CfnCluster_HadoopJarStepConfigProperty struct {
	// `CfnCluster.HadoopJarStepConfigProperty.Jar`.
	Jar *string `json:"jar"`
	// `CfnCluster.HadoopJarStepConfigProperty.Args`.
	Args *[]*string `json:"args"`
	// `CfnCluster.HadoopJarStepConfigProperty.MainClass`.
	MainClass *string `json:"mainClass"`
	// `CfnCluster.HadoopJarStepConfigProperty.StepProperties`.
	StepProperties interface{} `json:"stepProperties"`
}

type CfnCluster_InstanceFleetConfigProperty struct {
	// `CfnCluster.InstanceFleetConfigProperty.InstanceTypeConfigs`.
	InstanceTypeConfigs interface{} `json:"instanceTypeConfigs"`
	// `CfnCluster.InstanceFleetConfigProperty.LaunchSpecifications`.
	LaunchSpecifications interface{} `json:"launchSpecifications"`
	// `CfnCluster.InstanceFleetConfigProperty.Name`.
	Name *string `json:"name"`
	// `CfnCluster.InstanceFleetConfigProperty.TargetOnDemandCapacity`.
	TargetOnDemandCapacity *float64 `json:"targetOnDemandCapacity"`
	// `CfnCluster.InstanceFleetConfigProperty.TargetSpotCapacity`.
	TargetSpotCapacity *float64 `json:"targetSpotCapacity"`
}

type CfnCluster_InstanceFleetProvisioningSpecificationsProperty struct {
	// `CfnCluster.InstanceFleetProvisioningSpecificationsProperty.OnDemandSpecification`.
	OnDemandSpecification interface{} `json:"onDemandSpecification"`
	// `CfnCluster.InstanceFleetProvisioningSpecificationsProperty.SpotSpecification`.
	SpotSpecification interface{} `json:"spotSpecification"`
}

type CfnCluster_InstanceGroupConfigProperty struct {
	// `CfnCluster.InstanceGroupConfigProperty.InstanceCount`.
	InstanceCount *float64 `json:"instanceCount"`
	// `CfnCluster.InstanceGroupConfigProperty.InstanceType`.
	InstanceType *string `json:"instanceType"`
	// `CfnCluster.InstanceGroupConfigProperty.AutoScalingPolicy`.
	AutoScalingPolicy interface{} `json:"autoScalingPolicy"`
	// `CfnCluster.InstanceGroupConfigProperty.BidPrice`.
	BidPrice *string `json:"bidPrice"`
	// `CfnCluster.InstanceGroupConfigProperty.Configurations`.
	Configurations interface{} `json:"configurations"`
	// `CfnCluster.InstanceGroupConfigProperty.EbsConfiguration`.
	EbsConfiguration interface{} `json:"ebsConfiguration"`
	// `CfnCluster.InstanceGroupConfigProperty.Market`.
	Market *string `json:"market"`
	// `CfnCluster.InstanceGroupConfigProperty.Name`.
	Name *string `json:"name"`
}

type CfnCluster_InstanceTypeConfigProperty struct {
	// `CfnCluster.InstanceTypeConfigProperty.InstanceType`.
	InstanceType *string `json:"instanceType"`
	// `CfnCluster.InstanceTypeConfigProperty.BidPrice`.
	BidPrice *string `json:"bidPrice"`
	// `CfnCluster.InstanceTypeConfigProperty.BidPriceAsPercentageOfOnDemandPrice`.
	BidPriceAsPercentageOfOnDemandPrice *float64 `json:"bidPriceAsPercentageOfOnDemandPrice"`
	// `CfnCluster.InstanceTypeConfigProperty.Configurations`.
	Configurations interface{} `json:"configurations"`
	// `CfnCluster.InstanceTypeConfigProperty.EbsConfiguration`.
	EbsConfiguration interface{} `json:"ebsConfiguration"`
	// `CfnCluster.InstanceTypeConfigProperty.WeightedCapacity`.
	WeightedCapacity *float64 `json:"weightedCapacity"`
}

type CfnCluster_JobFlowInstancesConfigProperty struct {
	// `CfnCluster.JobFlowInstancesConfigProperty.AdditionalMasterSecurityGroups`.
	AdditionalMasterSecurityGroups *[]*string `json:"additionalMasterSecurityGroups"`
	// `CfnCluster.JobFlowInstancesConfigProperty.AdditionalSlaveSecurityGroups`.
	AdditionalSlaveSecurityGroups *[]*string `json:"additionalSlaveSecurityGroups"`
	// `CfnCluster.JobFlowInstancesConfigProperty.CoreInstanceFleet`.
	CoreInstanceFleet interface{} `json:"coreInstanceFleet"`
	// `CfnCluster.JobFlowInstancesConfigProperty.CoreInstanceGroup`.
	CoreInstanceGroup interface{} `json:"coreInstanceGroup"`
	// `CfnCluster.JobFlowInstancesConfigProperty.Ec2KeyName`.
	Ec2KeyName *string `json:"ec2KeyName"`
	// `CfnCluster.JobFlowInstancesConfigProperty.Ec2SubnetId`.
	Ec2SubnetId *string `json:"ec2SubnetId"`
	// `CfnCluster.JobFlowInstancesConfigProperty.Ec2SubnetIds`.
	Ec2SubnetIds *[]*string `json:"ec2SubnetIds"`
	// `CfnCluster.JobFlowInstancesConfigProperty.EmrManagedMasterSecurityGroup`.
	EmrManagedMasterSecurityGroup *string `json:"emrManagedMasterSecurityGroup"`
	// `CfnCluster.JobFlowInstancesConfigProperty.EmrManagedSlaveSecurityGroup`.
	EmrManagedSlaveSecurityGroup *string `json:"emrManagedSlaveSecurityGroup"`
	// `CfnCluster.JobFlowInstancesConfigProperty.HadoopVersion`.
	HadoopVersion *string `json:"hadoopVersion"`
	// `CfnCluster.JobFlowInstancesConfigProperty.KeepJobFlowAliveWhenNoSteps`.
	KeepJobFlowAliveWhenNoSteps interface{} `json:"keepJobFlowAliveWhenNoSteps"`
	// `CfnCluster.JobFlowInstancesConfigProperty.MasterInstanceFleet`.
	MasterInstanceFleet interface{} `json:"masterInstanceFleet"`
	// `CfnCluster.JobFlowInstancesConfigProperty.MasterInstanceGroup`.
	MasterInstanceGroup interface{} `json:"masterInstanceGroup"`
	// `CfnCluster.JobFlowInstancesConfigProperty.Placement`.
	Placement interface{} `json:"placement"`
	// `CfnCluster.JobFlowInstancesConfigProperty.ServiceAccessSecurityGroup`.
	ServiceAccessSecurityGroup *string `json:"serviceAccessSecurityGroup"`
	// `CfnCluster.JobFlowInstancesConfigProperty.TerminationProtected`.
	TerminationProtected interface{} `json:"terminationProtected"`
}

type CfnCluster_KerberosAttributesProperty struct {
	// `CfnCluster.KerberosAttributesProperty.KdcAdminPassword`.
	KdcAdminPassword *string `json:"kdcAdminPassword"`
	// `CfnCluster.KerberosAttributesProperty.Realm`.
	Realm *string `json:"realm"`
	// `CfnCluster.KerberosAttributesProperty.ADDomainJoinPassword`.
	AdDomainJoinPassword *string `json:"adDomainJoinPassword"`
	// `CfnCluster.KerberosAttributesProperty.ADDomainJoinUser`.
	AdDomainJoinUser *string `json:"adDomainJoinUser"`
	// `CfnCluster.KerberosAttributesProperty.CrossRealmTrustPrincipalPassword`.
	CrossRealmTrustPrincipalPassword *string `json:"crossRealmTrustPrincipalPassword"`
}

type CfnCluster_KeyValueProperty struct {
	// `CfnCluster.KeyValueProperty.Key`.
	Key *string `json:"key"`
	// `CfnCluster.KeyValueProperty.Value`.
	Value *string `json:"value"`
}

type CfnCluster_ManagedScalingPolicyProperty struct {
	// `CfnCluster.ManagedScalingPolicyProperty.ComputeLimits`.
	ComputeLimits interface{} `json:"computeLimits"`
}

type CfnCluster_MetricDimensionProperty struct {
	// `CfnCluster.MetricDimensionProperty.Key`.
	Key *string `json:"key"`
	// `CfnCluster.MetricDimensionProperty.Value`.
	Value *string `json:"value"`
}

type CfnCluster_OnDemandProvisioningSpecificationProperty struct {
	// `CfnCluster.OnDemandProvisioningSpecificationProperty.AllocationStrategy`.
	AllocationStrategy *string `json:"allocationStrategy"`
}

type CfnCluster_PlacementTypeProperty struct {
	// `CfnCluster.PlacementTypeProperty.AvailabilityZone`.
	AvailabilityZone *string `json:"availabilityZone"`
}

type CfnCluster_ScalingActionProperty struct {
	// `CfnCluster.ScalingActionProperty.SimpleScalingPolicyConfiguration`.
	SimpleScalingPolicyConfiguration interface{} `json:"simpleScalingPolicyConfiguration"`
	// `CfnCluster.ScalingActionProperty.Market`.
	Market *string `json:"market"`
}

type CfnCluster_ScalingConstraintsProperty struct {
	// `CfnCluster.ScalingConstraintsProperty.MaxCapacity`.
	MaxCapacity *float64 `json:"maxCapacity"`
	// `CfnCluster.ScalingConstraintsProperty.MinCapacity`.
	MinCapacity *float64 `json:"minCapacity"`
}

type CfnCluster_ScalingRuleProperty struct {
	// `CfnCluster.ScalingRuleProperty.Action`.
	Action interface{} `json:"action"`
	// `CfnCluster.ScalingRuleProperty.Name`.
	Name *string `json:"name"`
	// `CfnCluster.ScalingRuleProperty.Trigger`.
	Trigger interface{} `json:"trigger"`
	// `CfnCluster.ScalingRuleProperty.Description`.
	Description *string `json:"description"`
}

type CfnCluster_ScalingTriggerProperty struct {
	// `CfnCluster.ScalingTriggerProperty.CloudWatchAlarmDefinition`.
	CloudWatchAlarmDefinition interface{} `json:"cloudWatchAlarmDefinition"`
}

type CfnCluster_ScriptBootstrapActionConfigProperty struct {
	// `CfnCluster.ScriptBootstrapActionConfigProperty.Path`.
	Path *string `json:"path"`
	// `CfnCluster.ScriptBootstrapActionConfigProperty.Args`.
	Args *[]*string `json:"args"`
}

type CfnCluster_SimpleScalingPolicyConfigurationProperty struct {
	// `CfnCluster.SimpleScalingPolicyConfigurationProperty.ScalingAdjustment`.
	ScalingAdjustment *float64 `json:"scalingAdjustment"`
	// `CfnCluster.SimpleScalingPolicyConfigurationProperty.AdjustmentType`.
	AdjustmentType *string `json:"adjustmentType"`
	// `CfnCluster.SimpleScalingPolicyConfigurationProperty.CoolDown`.
	CoolDown *float64 `json:"coolDown"`
}

type CfnCluster_SpotProvisioningSpecificationProperty struct {
	// `CfnCluster.SpotProvisioningSpecificationProperty.TimeoutAction`.
	TimeoutAction *string `json:"timeoutAction"`
	// `CfnCluster.SpotProvisioningSpecificationProperty.TimeoutDurationMinutes`.
	TimeoutDurationMinutes *float64 `json:"timeoutDurationMinutes"`
	// `CfnCluster.SpotProvisioningSpecificationProperty.AllocationStrategy`.
	AllocationStrategy *string `json:"allocationStrategy"`
	// `CfnCluster.SpotProvisioningSpecificationProperty.BlockDurationMinutes`.
	BlockDurationMinutes *float64 `json:"blockDurationMinutes"`
}

type CfnCluster_StepConfigProperty struct {
	// `CfnCluster.StepConfigProperty.HadoopJarStep`.
	HadoopJarStep interface{} `json:"hadoopJarStep"`
	// `CfnCluster.StepConfigProperty.Name`.
	Name *string `json:"name"`
	// `CfnCluster.StepConfigProperty.ActionOnFailure`.
	ActionOnFailure *string `json:"actionOnFailure"`
}

type CfnCluster_VolumeSpecificationProperty struct {
	// `CfnCluster.VolumeSpecificationProperty.SizeInGB`.
	SizeInGb *float64 `json:"sizeInGb"`
	// `CfnCluster.VolumeSpecificationProperty.VolumeType`.
	VolumeType *string `json:"volumeType"`
	// `CfnCluster.VolumeSpecificationProperty.Iops`.
	Iops *float64 `json:"iops"`
}

// Properties for defining a `AWS::EMR::Cluster`.
type CfnClusterProps struct {
	// `AWS::EMR::Cluster.Instances`.
	Instances interface{} `json:"instances"`
	// `AWS::EMR::Cluster.JobFlowRole`.
	JobFlowRole *string `json:"jobFlowRole"`
	// `AWS::EMR::Cluster.Name`.
	Name *string `json:"name"`
	// `AWS::EMR::Cluster.ServiceRole`.
	ServiceRole *string `json:"serviceRole"`
	// `AWS::EMR::Cluster.AdditionalInfo`.
	AdditionalInfo interface{} `json:"additionalInfo"`
	// `AWS::EMR::Cluster.Applications`.
	Applications interface{} `json:"applications"`
	// `AWS::EMR::Cluster.AutoScalingRole`.
	AutoScalingRole *string `json:"autoScalingRole"`
	// `AWS::EMR::Cluster.BootstrapActions`.
	BootstrapActions interface{} `json:"bootstrapActions"`
	// `AWS::EMR::Cluster.Configurations`.
	Configurations interface{} `json:"configurations"`
	// `AWS::EMR::Cluster.CustomAmiId`.
	CustomAmiId *string `json:"customAmiId"`
	// `AWS::EMR::Cluster.EbsRootVolumeSize`.
	EbsRootVolumeSize *float64 `json:"ebsRootVolumeSize"`
	// `AWS::EMR::Cluster.KerberosAttributes`.
	KerberosAttributes interface{} `json:"kerberosAttributes"`
	// `AWS::EMR::Cluster.LogEncryptionKmsKeyId`.
	LogEncryptionKmsKeyId *string `json:"logEncryptionKmsKeyId"`
	// `AWS::EMR::Cluster.LogUri`.
	LogUri *string `json:"logUri"`
	// `AWS::EMR::Cluster.ManagedScalingPolicy`.
	ManagedScalingPolicy interface{} `json:"managedScalingPolicy"`
	// `AWS::EMR::Cluster.ReleaseLabel`.
	ReleaseLabel *string `json:"releaseLabel"`
	// `AWS::EMR::Cluster.ScaleDownBehavior`.
	ScaleDownBehavior *string `json:"scaleDownBehavior"`
	// `AWS::EMR::Cluster.SecurityConfiguration`.
	SecurityConfiguration *string `json:"securityConfiguration"`
	// `AWS::EMR::Cluster.StepConcurrencyLevel`.
	StepConcurrencyLevel *float64 `json:"stepConcurrencyLevel"`
	// `AWS::EMR::Cluster.Steps`.
	Steps interface{} `json:"steps"`
	// `AWS::EMR::Cluster.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// `AWS::EMR::Cluster.VisibleToAllUsers`.
	VisibleToAllUsers interface{} `json:"visibleToAllUsers"`
}

// A CloudFormation `AWS::EMR::InstanceFleetConfig`.
type CfnInstanceFleetConfig interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ClusterId() *string
	SetClusterId(val *string)
	CreationStack() *[]*string
	InstanceFleetType() *string
	SetInstanceFleetType(val *string)
	InstanceTypeConfigs() interface{}
	SetInstanceTypeConfigs(val interface{})
	LaunchSpecifications() interface{}
	SetLaunchSpecifications(val interface{})
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	TargetOnDemandCapacity() *float64
	SetTargetOnDemandCapacity(val *float64)
	TargetSpotCapacity() *float64
	SetTargetSpotCapacity(val *float64)
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

// The jsii proxy struct for CfnInstanceFleetConfig
type jsiiProxy_CfnInstanceFleetConfig struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnInstanceFleetConfig) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceFleetConfig) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceFleetConfig) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceFleetConfig) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceFleetConfig) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceFleetConfig) InstanceFleetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceFleetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceFleetConfig) InstanceTypeConfigs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceTypeConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceFleetConfig) LaunchSpecifications() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"launchSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceFleetConfig) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceFleetConfig) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceFleetConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceFleetConfig) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceFleetConfig) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceFleetConfig) TargetOnDemandCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetOnDemandCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceFleetConfig) TargetSpotCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetSpotCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInstanceFleetConfig) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::EMR::InstanceFleetConfig`.
func NewCfnInstanceFleetConfig(scope constructs.Construct, id *string, props *CfnInstanceFleetConfigProps) CfnInstanceFleetConfig {
	_init_.Initialize()

	j := jsiiProxy_CfnInstanceFleetConfig{}

	_jsii_.Create(
		"aws-cdk-lib.aws_emr.CfnInstanceFleetConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::EMR::InstanceFleetConfig`.
func NewCfnInstanceFleetConfig_Override(c CfnInstanceFleetConfig, scope constructs.Construct, id *string, props *CfnInstanceFleetConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_emr.CfnInstanceFleetConfig",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnInstanceFleetConfig) SetClusterId(val *string) {
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceFleetConfig) SetInstanceFleetType(val *string) {
	_jsii_.Set(
		j,
		"instanceFleetType",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceFleetConfig) SetInstanceTypeConfigs(val interface{}) {
	_jsii_.Set(
		j,
		"instanceTypeConfigs",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceFleetConfig) SetLaunchSpecifications(val interface{}) {
	_jsii_.Set(
		j,
		"launchSpecifications",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceFleetConfig) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceFleetConfig) SetTargetOnDemandCapacity(val *float64) {
	_jsii_.Set(
		j,
		"targetOnDemandCapacity",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceFleetConfig) SetTargetSpotCapacity(val *float64) {
	_jsii_.Set(
		j,
		"targetSpotCapacity",
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
func CfnInstanceFleetConfig_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnInstanceFleetConfig",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnInstanceFleetConfig_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnInstanceFleetConfig",
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
func CfnInstanceFleetConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnInstanceFleetConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInstanceFleetConfig_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_emr.CfnInstanceFleetConfig",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnInstanceFleetConfig) AddDeletionOverride(path *string) {
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
// Experimental.
func (c *jsiiProxy_CfnInstanceFleetConfig) AddDependsOn(target awscdk.CfnResource) {
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
// Experimental.
func (c *jsiiProxy_CfnInstanceFleetConfig) AddMetadata(key *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnInstanceFleetConfig) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnInstanceFleetConfig) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnInstanceFleetConfig) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnInstanceFleetConfig) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
// Experimental.
func (c *jsiiProxy_CfnInstanceFleetConfig) GetAtt(attributeName *string) awscdk.Reference {
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
// Experimental.
func (c *jsiiProxy_CfnInstanceFleetConfig) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnInstanceFleetConfig) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnInstanceFleetConfig) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnInstanceFleetConfig) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
// Experimental.
func (c *jsiiProxy_CfnInstanceFleetConfig) ShouldSynthesize() *bool {
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
// Experimental.
func (c *jsiiProxy_CfnInstanceFleetConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnInstanceFleetConfig) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnInstanceFleetConfig_ConfigurationProperty struct {
	// `CfnInstanceFleetConfig.ConfigurationProperty.Classification`.
	Classification *string `json:"classification"`
	// `CfnInstanceFleetConfig.ConfigurationProperty.ConfigurationProperties`.
	ConfigurationProperties interface{} `json:"configurationProperties"`
	// `CfnInstanceFleetConfig.ConfigurationProperty.Configurations`.
	Configurations interface{} `json:"configurations"`
}

type CfnInstanceFleetConfig_EbsBlockDeviceConfigProperty struct {
	// `CfnInstanceFleetConfig.EbsBlockDeviceConfigProperty.VolumeSpecification`.
	VolumeSpecification interface{} `json:"volumeSpecification"`
	// `CfnInstanceFleetConfig.EbsBlockDeviceConfigProperty.VolumesPerInstance`.
	VolumesPerInstance *float64 `json:"volumesPerInstance"`
}

type CfnInstanceFleetConfig_EbsConfigurationProperty struct {
	// `CfnInstanceFleetConfig.EbsConfigurationProperty.EbsBlockDeviceConfigs`.
	EbsBlockDeviceConfigs interface{} `json:"ebsBlockDeviceConfigs"`
	// `CfnInstanceFleetConfig.EbsConfigurationProperty.EbsOptimized`.
	EbsOptimized interface{} `json:"ebsOptimized"`
}

type CfnInstanceFleetConfig_InstanceFleetProvisioningSpecificationsProperty struct {
	// `CfnInstanceFleetConfig.InstanceFleetProvisioningSpecificationsProperty.OnDemandSpecification`.
	OnDemandSpecification interface{} `json:"onDemandSpecification"`
	// `CfnInstanceFleetConfig.InstanceFleetProvisioningSpecificationsProperty.SpotSpecification`.
	SpotSpecification interface{} `json:"spotSpecification"`
}

type CfnInstanceFleetConfig_InstanceTypeConfigProperty struct {
	// `CfnInstanceFleetConfig.InstanceTypeConfigProperty.InstanceType`.
	InstanceType *string `json:"instanceType"`
	// `CfnInstanceFleetConfig.InstanceTypeConfigProperty.BidPrice`.
	BidPrice *string `json:"bidPrice"`
	// `CfnInstanceFleetConfig.InstanceTypeConfigProperty.BidPriceAsPercentageOfOnDemandPrice`.
	BidPriceAsPercentageOfOnDemandPrice *float64 `json:"bidPriceAsPercentageOfOnDemandPrice"`
	// `CfnInstanceFleetConfig.InstanceTypeConfigProperty.Configurations`.
	Configurations interface{} `json:"configurations"`
	// `CfnInstanceFleetConfig.InstanceTypeConfigProperty.EbsConfiguration`.
	EbsConfiguration interface{} `json:"ebsConfiguration"`
	// `CfnInstanceFleetConfig.InstanceTypeConfigProperty.WeightedCapacity`.
	WeightedCapacity *float64 `json:"weightedCapacity"`
}

type CfnInstanceFleetConfig_OnDemandProvisioningSpecificationProperty struct {
	// `CfnInstanceFleetConfig.OnDemandProvisioningSpecificationProperty.AllocationStrategy`.
	AllocationStrategy *string `json:"allocationStrategy"`
}

type CfnInstanceFleetConfig_SpotProvisioningSpecificationProperty struct {
	// `CfnInstanceFleetConfig.SpotProvisioningSpecificationProperty.TimeoutAction`.
	TimeoutAction *string `json:"timeoutAction"`
	// `CfnInstanceFleetConfig.SpotProvisioningSpecificationProperty.TimeoutDurationMinutes`.
	TimeoutDurationMinutes *float64 `json:"timeoutDurationMinutes"`
	// `CfnInstanceFleetConfig.SpotProvisioningSpecificationProperty.AllocationStrategy`.
	AllocationStrategy *string `json:"allocationStrategy"`
	// `CfnInstanceFleetConfig.SpotProvisioningSpecificationProperty.BlockDurationMinutes`.
	BlockDurationMinutes *float64 `json:"blockDurationMinutes"`
}

type CfnInstanceFleetConfig_VolumeSpecificationProperty struct {
	// `CfnInstanceFleetConfig.VolumeSpecificationProperty.SizeInGB`.
	SizeInGb *float64 `json:"sizeInGb"`
	// `CfnInstanceFleetConfig.VolumeSpecificationProperty.VolumeType`.
	VolumeType *string `json:"volumeType"`
	// `CfnInstanceFleetConfig.VolumeSpecificationProperty.Iops`.
	Iops *float64 `json:"iops"`
}

// Properties for defining a `AWS::EMR::InstanceFleetConfig`.
type CfnInstanceFleetConfigProps struct {
	// `AWS::EMR::InstanceFleetConfig.ClusterId`.
	ClusterId *string `json:"clusterId"`
	// `AWS::EMR::InstanceFleetConfig.InstanceFleetType`.
	InstanceFleetType *string `json:"instanceFleetType"`
	// `AWS::EMR::InstanceFleetConfig.InstanceTypeConfigs`.
	InstanceTypeConfigs interface{} `json:"instanceTypeConfigs"`
	// `AWS::EMR::InstanceFleetConfig.LaunchSpecifications`.
	LaunchSpecifications interface{} `json:"launchSpecifications"`
	// `AWS::EMR::InstanceFleetConfig.Name`.
	Name *string `json:"name"`
	// `AWS::EMR::InstanceFleetConfig.TargetOnDemandCapacity`.
	TargetOnDemandCapacity *float64 `json:"targetOnDemandCapacity"`
	// `AWS::EMR::InstanceFleetConfig.TargetSpotCapacity`.
	TargetSpotCapacity *float64 `json:"targetSpotCapacity"`
}

// A CloudFormation `AWS::EMR::InstanceGroupConfig`.
type CfnInstanceGroupConfig interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AutoScalingPolicy() interface{}
	SetAutoScalingPolicy(val interface{})
	BidPrice() *string
	SetBidPrice(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	Configurations() interface{}
	SetConfigurations(val interface{})
	CreationStack() *[]*string
	EbsConfiguration() interface{}
	SetEbsConfiguration(val interface{})
	InstanceCount() *float64
	SetInstanceCount(val *float64)
	InstanceRole() *string
	SetInstanceRole(val *string)
	InstanceType() *string
	SetInstanceType(val *string)
	JobFlowId() *string
	SetJobFlowId(val *string)
	LogicalId() *string
	Market() *string
	SetMarket(val *string)
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
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

func (j *jsiiProxy_CfnInstanceGroupConfig) Node() constructs.Node {
	var returns constructs.Node
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
func NewCfnInstanceGroupConfig(scope constructs.Construct, id *string, props *CfnInstanceGroupConfigProps) CfnInstanceGroupConfig {
	_init_.Initialize()

	j := jsiiProxy_CfnInstanceGroupConfig{}

	_jsii_.Create(
		"aws-cdk-lib.aws_emr.CfnInstanceGroupConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::EMR::InstanceGroupConfig`.
func NewCfnInstanceGroupConfig_Override(c CfnInstanceGroupConfig, scope constructs.Construct, id *string, props *CfnInstanceGroupConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_emr.CfnInstanceGroupConfig",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnInstanceGroupConfig) SetAutoScalingPolicy(val interface{}) {
	_jsii_.Set(
		j,
		"autoScalingPolicy",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceGroupConfig) SetBidPrice(val *string) {
	_jsii_.Set(
		j,
		"bidPrice",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceGroupConfig) SetConfigurations(val interface{}) {
	_jsii_.Set(
		j,
		"configurations",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceGroupConfig) SetEbsConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"ebsConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceGroupConfig) SetInstanceCount(val *float64) {
	_jsii_.Set(
		j,
		"instanceCount",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceGroupConfig) SetInstanceRole(val *string) {
	_jsii_.Set(
		j,
		"instanceRole",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceGroupConfig) SetInstanceType(val *string) {
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceGroupConfig) SetJobFlowId(val *string) {
	_jsii_.Set(
		j,
		"jobFlowId",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceGroupConfig) SetMarket(val *string) {
	_jsii_.Set(
		j,
		"market",
		val,
	)
}

func (j *jsiiProxy_CfnInstanceGroupConfig) SetName(val *string) {
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

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnInstanceGroupConfig",
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

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnInstanceGroupConfig",
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
func CfnInstanceGroupConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnInstanceGroupConfig",
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
		"aws-cdk-lib.aws_emr.CfnInstanceGroupConfig",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnInstanceGroupConfig) AddDeletionOverride(path *string) {
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
// Experimental.
func (c *jsiiProxy_CfnInstanceGroupConfig) AddDependsOn(target awscdk.CfnResource) {
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
// Experimental.
func (c *jsiiProxy_CfnInstanceGroupConfig) AddMetadata(key *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnInstanceGroupConfig) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnInstanceGroupConfig) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnInstanceGroupConfig) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnInstanceGroupConfig) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
// Experimental.
func (c *jsiiProxy_CfnInstanceGroupConfig) GetAtt(attributeName *string) awscdk.Reference {
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
// Experimental.
func (c *jsiiProxy_CfnInstanceGroupConfig) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnInstanceGroupConfig) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnInstanceGroupConfig) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnInstanceGroupConfig) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
// Experimental.
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

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
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

// Experimental.
func (c *jsiiProxy_CfnInstanceGroupConfig) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnInstanceGroupConfig_AutoScalingPolicyProperty struct {
	// `CfnInstanceGroupConfig.AutoScalingPolicyProperty.Constraints`.
	Constraints interface{} `json:"constraints"`
	// `CfnInstanceGroupConfig.AutoScalingPolicyProperty.Rules`.
	Rules interface{} `json:"rules"`
}

type CfnInstanceGroupConfig_CloudWatchAlarmDefinitionProperty struct {
	// `CfnInstanceGroupConfig.CloudWatchAlarmDefinitionProperty.ComparisonOperator`.
	ComparisonOperator *string `json:"comparisonOperator"`
	// `CfnInstanceGroupConfig.CloudWatchAlarmDefinitionProperty.MetricName`.
	MetricName *string `json:"metricName"`
	// `CfnInstanceGroupConfig.CloudWatchAlarmDefinitionProperty.Period`.
	Period *float64 `json:"period"`
	// `CfnInstanceGroupConfig.CloudWatchAlarmDefinitionProperty.Threshold`.
	Threshold *float64 `json:"threshold"`
	// `CfnInstanceGroupConfig.CloudWatchAlarmDefinitionProperty.Dimensions`.
	Dimensions interface{} `json:"dimensions"`
	// `CfnInstanceGroupConfig.CloudWatchAlarmDefinitionProperty.EvaluationPeriods`.
	EvaluationPeriods *float64 `json:"evaluationPeriods"`
	// `CfnInstanceGroupConfig.CloudWatchAlarmDefinitionProperty.Namespace`.
	Namespace *string `json:"namespace"`
	// `CfnInstanceGroupConfig.CloudWatchAlarmDefinitionProperty.Statistic`.
	Statistic *string `json:"statistic"`
	// `CfnInstanceGroupConfig.CloudWatchAlarmDefinitionProperty.Unit`.
	Unit *string `json:"unit"`
}

type CfnInstanceGroupConfig_ConfigurationProperty struct {
	// `CfnInstanceGroupConfig.ConfigurationProperty.Classification`.
	Classification *string `json:"classification"`
	// `CfnInstanceGroupConfig.ConfigurationProperty.ConfigurationProperties`.
	ConfigurationProperties interface{} `json:"configurationProperties"`
	// `CfnInstanceGroupConfig.ConfigurationProperty.Configurations`.
	Configurations interface{} `json:"configurations"`
}

type CfnInstanceGroupConfig_EbsBlockDeviceConfigProperty struct {
	// `CfnInstanceGroupConfig.EbsBlockDeviceConfigProperty.VolumeSpecification`.
	VolumeSpecification interface{} `json:"volumeSpecification"`
	// `CfnInstanceGroupConfig.EbsBlockDeviceConfigProperty.VolumesPerInstance`.
	VolumesPerInstance *float64 `json:"volumesPerInstance"`
}

type CfnInstanceGroupConfig_EbsConfigurationProperty struct {
	// `CfnInstanceGroupConfig.EbsConfigurationProperty.EbsBlockDeviceConfigs`.
	EbsBlockDeviceConfigs interface{} `json:"ebsBlockDeviceConfigs"`
	// `CfnInstanceGroupConfig.EbsConfigurationProperty.EbsOptimized`.
	EbsOptimized interface{} `json:"ebsOptimized"`
}

type CfnInstanceGroupConfig_MetricDimensionProperty struct {
	// `CfnInstanceGroupConfig.MetricDimensionProperty.Key`.
	Key *string `json:"key"`
	// `CfnInstanceGroupConfig.MetricDimensionProperty.Value`.
	Value *string `json:"value"`
}

type CfnInstanceGroupConfig_ScalingActionProperty struct {
	// `CfnInstanceGroupConfig.ScalingActionProperty.SimpleScalingPolicyConfiguration`.
	SimpleScalingPolicyConfiguration interface{} `json:"simpleScalingPolicyConfiguration"`
	// `CfnInstanceGroupConfig.ScalingActionProperty.Market`.
	Market *string `json:"market"`
}

type CfnInstanceGroupConfig_ScalingConstraintsProperty struct {
	// `CfnInstanceGroupConfig.ScalingConstraintsProperty.MaxCapacity`.
	MaxCapacity *float64 `json:"maxCapacity"`
	// `CfnInstanceGroupConfig.ScalingConstraintsProperty.MinCapacity`.
	MinCapacity *float64 `json:"minCapacity"`
}

type CfnInstanceGroupConfig_ScalingRuleProperty struct {
	// `CfnInstanceGroupConfig.ScalingRuleProperty.Action`.
	Action interface{} `json:"action"`
	// `CfnInstanceGroupConfig.ScalingRuleProperty.Name`.
	Name *string `json:"name"`
	// `CfnInstanceGroupConfig.ScalingRuleProperty.Trigger`.
	Trigger interface{} `json:"trigger"`
	// `CfnInstanceGroupConfig.ScalingRuleProperty.Description`.
	Description *string `json:"description"`
}

type CfnInstanceGroupConfig_ScalingTriggerProperty struct {
	// `CfnInstanceGroupConfig.ScalingTriggerProperty.CloudWatchAlarmDefinition`.
	CloudWatchAlarmDefinition interface{} `json:"cloudWatchAlarmDefinition"`
}

type CfnInstanceGroupConfig_SimpleScalingPolicyConfigurationProperty struct {
	// `CfnInstanceGroupConfig.SimpleScalingPolicyConfigurationProperty.ScalingAdjustment`.
	ScalingAdjustment *float64 `json:"scalingAdjustment"`
	// `CfnInstanceGroupConfig.SimpleScalingPolicyConfigurationProperty.AdjustmentType`.
	AdjustmentType *string `json:"adjustmentType"`
	// `CfnInstanceGroupConfig.SimpleScalingPolicyConfigurationProperty.CoolDown`.
	CoolDown *float64 `json:"coolDown"`
}

type CfnInstanceGroupConfig_VolumeSpecificationProperty struct {
	// `CfnInstanceGroupConfig.VolumeSpecificationProperty.SizeInGB`.
	SizeInGb *float64 `json:"sizeInGb"`
	// `CfnInstanceGroupConfig.VolumeSpecificationProperty.VolumeType`.
	VolumeType *string `json:"volumeType"`
	// `CfnInstanceGroupConfig.VolumeSpecificationProperty.Iops`.
	Iops *float64 `json:"iops"`
}

// Properties for defining a `AWS::EMR::InstanceGroupConfig`.
type CfnInstanceGroupConfigProps struct {
	// `AWS::EMR::InstanceGroupConfig.InstanceCount`.
	InstanceCount *float64 `json:"instanceCount"`
	// `AWS::EMR::InstanceGroupConfig.InstanceRole`.
	InstanceRole *string `json:"instanceRole"`
	// `AWS::EMR::InstanceGroupConfig.InstanceType`.
	InstanceType *string `json:"instanceType"`
	// `AWS::EMR::InstanceGroupConfig.JobFlowId`.
	JobFlowId *string `json:"jobFlowId"`
	// `AWS::EMR::InstanceGroupConfig.AutoScalingPolicy`.
	AutoScalingPolicy interface{} `json:"autoScalingPolicy"`
	// `AWS::EMR::InstanceGroupConfig.BidPrice`.
	BidPrice *string `json:"bidPrice"`
	// `AWS::EMR::InstanceGroupConfig.Configurations`.
	Configurations interface{} `json:"configurations"`
	// `AWS::EMR::InstanceGroupConfig.EbsConfiguration`.
	EbsConfiguration interface{} `json:"ebsConfiguration"`
	// `AWS::EMR::InstanceGroupConfig.Market`.
	Market *string `json:"market"`
	// `AWS::EMR::InstanceGroupConfig.Name`.
	Name *string `json:"name"`
}

// A CloudFormation `AWS::EMR::SecurityConfiguration`.
type CfnSecurityConfiguration interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	SecurityConfiguration() interface{}
	SetSecurityConfiguration(val interface{})
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

// The jsii proxy struct for CfnSecurityConfiguration
type jsiiProxy_CfnSecurityConfiguration struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnSecurityConfiguration) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityConfiguration) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityConfiguration) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityConfiguration) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityConfiguration) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityConfiguration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityConfiguration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityConfiguration) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityConfiguration) SecurityConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"securityConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityConfiguration) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSecurityConfiguration) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::EMR::SecurityConfiguration`.
func NewCfnSecurityConfiguration(scope constructs.Construct, id *string, props *CfnSecurityConfigurationProps) CfnSecurityConfiguration {
	_init_.Initialize()

	j := jsiiProxy_CfnSecurityConfiguration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_emr.CfnSecurityConfiguration",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::EMR::SecurityConfiguration`.
func NewCfnSecurityConfiguration_Override(c CfnSecurityConfiguration, scope constructs.Construct, id *string, props *CfnSecurityConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_emr.CfnSecurityConfiguration",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnSecurityConfiguration) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnSecurityConfiguration) SetSecurityConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"securityConfiguration",
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
func CfnSecurityConfiguration_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnSecurityConfiguration",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnSecurityConfiguration_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnSecurityConfiguration",
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
func CfnSecurityConfiguration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnSecurityConfiguration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSecurityConfiguration_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_emr.CfnSecurityConfiguration",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnSecurityConfiguration) AddDeletionOverride(path *string) {
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
// Experimental.
func (c *jsiiProxy_CfnSecurityConfiguration) AddDependsOn(target awscdk.CfnResource) {
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
// Experimental.
func (c *jsiiProxy_CfnSecurityConfiguration) AddMetadata(key *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnSecurityConfiguration) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnSecurityConfiguration) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnSecurityConfiguration) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnSecurityConfiguration) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
// Experimental.
func (c *jsiiProxy_CfnSecurityConfiguration) GetAtt(attributeName *string) awscdk.Reference {
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
// Experimental.
func (c *jsiiProxy_CfnSecurityConfiguration) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnSecurityConfiguration) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnSecurityConfiguration) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnSecurityConfiguration) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
// Experimental.
func (c *jsiiProxy_CfnSecurityConfiguration) ShouldSynthesize() *bool {
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
// Experimental.
func (c *jsiiProxy_CfnSecurityConfiguration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnSecurityConfiguration) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::EMR::SecurityConfiguration`.
type CfnSecurityConfigurationProps struct {
	// `AWS::EMR::SecurityConfiguration.SecurityConfiguration`.
	SecurityConfiguration interface{} `json:"securityConfiguration"`
	// `AWS::EMR::SecurityConfiguration.Name`.
	Name *string `json:"name"`
}

// A CloudFormation `AWS::EMR::Step`.
type CfnStep interface {
	awscdk.CfnResource
	awscdk.IInspectable
	ActionOnFailure() *string
	SetActionOnFailure(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	HadoopJarStep() interface{}
	SetHadoopJarStep(val interface{})
	JobFlowId() *string
	SetJobFlowId(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
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

// The jsii proxy struct for CfnStep
type jsiiProxy_CfnStep struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnStep) ActionOnFailure() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionOnFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStep) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStep) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStep) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStep) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStep) HadoopJarStep() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hadoopJarStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStep) JobFlowId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobFlowId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStep) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStep) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStep) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStep) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStep) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStep) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::EMR::Step`.
func NewCfnStep(scope constructs.Construct, id *string, props *CfnStepProps) CfnStep {
	_init_.Initialize()

	j := jsiiProxy_CfnStep{}

	_jsii_.Create(
		"aws-cdk-lib.aws_emr.CfnStep",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::EMR::Step`.
func NewCfnStep_Override(c CfnStep, scope constructs.Construct, id *string, props *CfnStepProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_emr.CfnStep",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnStep) SetActionOnFailure(val *string) {
	_jsii_.Set(
		j,
		"actionOnFailure",
		val,
	)
}

func (j *jsiiProxy_CfnStep) SetHadoopJarStep(val interface{}) {
	_jsii_.Set(
		j,
		"hadoopJarStep",
		val,
	)
}

func (j *jsiiProxy_CfnStep) SetJobFlowId(val *string) {
	_jsii_.Set(
		j,
		"jobFlowId",
		val,
	)
}

func (j *jsiiProxy_CfnStep) SetName(val *string) {
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
func CfnStep_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnStep",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnStep_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnStep",
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
func CfnStep_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnStep",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStep_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_emr.CfnStep",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnStep) AddDeletionOverride(path *string) {
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
// Experimental.
func (c *jsiiProxy_CfnStep) AddDependsOn(target awscdk.CfnResource) {
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
// Experimental.
func (c *jsiiProxy_CfnStep) AddMetadata(key *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnStep) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnStep) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnStep) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnStep) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
// Experimental.
func (c *jsiiProxy_CfnStep) GetAtt(attributeName *string) awscdk.Reference {
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
// Experimental.
func (c *jsiiProxy_CfnStep) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnStep) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnStep) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnStep) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
// Experimental.
func (c *jsiiProxy_CfnStep) ShouldSynthesize() *bool {
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
// Experimental.
func (c *jsiiProxy_CfnStep) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnStep) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnStep_HadoopJarStepConfigProperty struct {
	// `CfnStep.HadoopJarStepConfigProperty.Jar`.
	Jar *string `json:"jar"`
	// `CfnStep.HadoopJarStepConfigProperty.Args`.
	Args *[]*string `json:"args"`
	// `CfnStep.HadoopJarStepConfigProperty.MainClass`.
	MainClass *string `json:"mainClass"`
	// `CfnStep.HadoopJarStepConfigProperty.StepProperties`.
	StepProperties interface{} `json:"stepProperties"`
}

type CfnStep_KeyValueProperty struct {
	// `CfnStep.KeyValueProperty.Key`.
	Key *string `json:"key"`
	// `CfnStep.KeyValueProperty.Value`.
	Value *string `json:"value"`
}

// Properties for defining a `AWS::EMR::Step`.
type CfnStepProps struct {
	// `AWS::EMR::Step.ActionOnFailure`.
	ActionOnFailure *string `json:"actionOnFailure"`
	// `AWS::EMR::Step.HadoopJarStep`.
	HadoopJarStep interface{} `json:"hadoopJarStep"`
	// `AWS::EMR::Step.JobFlowId`.
	JobFlowId *string `json:"jobFlowId"`
	// `AWS::EMR::Step.Name`.
	Name *string `json:"name"`
}

// A CloudFormation `AWS::EMR::Studio`.
type CfnStudio interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrStudioId() *string
	AttrUrl() *string
	AuthMode() *string
	SetAuthMode(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DefaultS3Location() *string
	SetDefaultS3Location(val *string)
	Description() *string
	SetDescription(val *string)
	EngineSecurityGroupId() *string
	SetEngineSecurityGroupId(val *string)
	LogicalId() *string
	Name() *string
	SetName(val *string)
	Node() constructs.Node
	Ref() *string
	ServiceRole() *string
	SetServiceRole(val *string)
	Stack() awscdk.Stack
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	Tags() awscdk.TagManager
	UpdatedProperites() *map[string]interface{}
	UserRole() *string
	SetUserRole(val *string)
	VpcId() *string
	SetVpcId(val *string)
	WorkspaceSecurityGroupId() *string
	SetWorkspaceSecurityGroupId(val *string)
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

// The jsii proxy struct for CfnStudio
type jsiiProxy_CfnStudio struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnStudio) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) AttrStudioId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStudioId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) AttrUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) AuthMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) DefaultS3Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultS3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) EngineSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) ServiceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) UserRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudio) WorkspaceSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceSecurityGroupId",
		&returns,
	)
	return returns
}


// Create a new `AWS::EMR::Studio`.
func NewCfnStudio(scope constructs.Construct, id *string, props *CfnStudioProps) CfnStudio {
	_init_.Initialize()

	j := jsiiProxy_CfnStudio{}

	_jsii_.Create(
		"aws-cdk-lib.aws_emr.CfnStudio",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::EMR::Studio`.
func NewCfnStudio_Override(c CfnStudio, scope constructs.Construct, id *string, props *CfnStudioProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_emr.CfnStudio",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnStudio) SetAuthMode(val *string) {
	_jsii_.Set(
		j,
		"authMode",
		val,
	)
}

func (j *jsiiProxy_CfnStudio) SetDefaultS3Location(val *string) {
	_jsii_.Set(
		j,
		"defaultS3Location",
		val,
	)
}

func (j *jsiiProxy_CfnStudio) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnStudio) SetEngineSecurityGroupId(val *string) {
	_jsii_.Set(
		j,
		"engineSecurityGroupId",
		val,
	)
}

func (j *jsiiProxy_CfnStudio) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnStudio) SetServiceRole(val *string) {
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
}

func (j *jsiiProxy_CfnStudio) SetSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_CfnStudio) SetUserRole(val *string) {
	_jsii_.Set(
		j,
		"userRole",
		val,
	)
}

func (j *jsiiProxy_CfnStudio) SetVpcId(val *string) {
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

func (j *jsiiProxy_CfnStudio) SetWorkspaceSecurityGroupId(val *string) {
	_jsii_.Set(
		j,
		"workspaceSecurityGroupId",
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
func CfnStudio_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnStudio",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnStudio_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnStudio",
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
func CfnStudio_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnStudio",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStudio_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_emr.CfnStudio",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnStudio) AddDeletionOverride(path *string) {
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
// Experimental.
func (c *jsiiProxy_CfnStudio) AddDependsOn(target awscdk.CfnResource) {
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
// Experimental.
func (c *jsiiProxy_CfnStudio) AddMetadata(key *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnStudio) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnStudio) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnStudio) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnStudio) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
// Experimental.
func (c *jsiiProxy_CfnStudio) GetAtt(attributeName *string) awscdk.Reference {
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
// Experimental.
func (c *jsiiProxy_CfnStudio) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnStudio) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnStudio) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnStudio) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
// Experimental.
func (c *jsiiProxy_CfnStudio) ShouldSynthesize() *bool {
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
// Experimental.
func (c *jsiiProxy_CfnStudio) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnStudio) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::EMR::Studio`.
type CfnStudioProps struct {
	// `AWS::EMR::Studio.AuthMode`.
	AuthMode *string `json:"authMode"`
	// `AWS::EMR::Studio.DefaultS3Location`.
	DefaultS3Location *string `json:"defaultS3Location"`
	// `AWS::EMR::Studio.EngineSecurityGroupId`.
	EngineSecurityGroupId *string `json:"engineSecurityGroupId"`
	// `AWS::EMR::Studio.Name`.
	Name *string `json:"name"`
	// `AWS::EMR::Studio.ServiceRole`.
	ServiceRole *string `json:"serviceRole"`
	// `AWS::EMR::Studio.SubnetIds`.
	SubnetIds *[]*string `json:"subnetIds"`
	// `AWS::EMR::Studio.UserRole`.
	UserRole *string `json:"userRole"`
	// `AWS::EMR::Studio.VpcId`.
	VpcId *string `json:"vpcId"`
	// `AWS::EMR::Studio.WorkspaceSecurityGroupId`.
	WorkspaceSecurityGroupId *string `json:"workspaceSecurityGroupId"`
	// `AWS::EMR::Studio.Description`.
	Description *string `json:"description"`
	// `AWS::EMR::Studio.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::EMR::StudioSessionMapping`.
type CfnStudioSessionMapping interface {
	awscdk.CfnResource
	awscdk.IInspectable
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	IdentityName() *string
	SetIdentityName(val *string)
	IdentityType() *string
	SetIdentityType(val *string)
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	SessionPolicyArn() *string
	SetSessionPolicyArn(val *string)
	Stack() awscdk.Stack
	StudioId() *string
	SetStudioId(val *string)
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

// The jsii proxy struct for CfnStudioSessionMapping
type jsiiProxy_CfnStudioSessionMapping struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnStudioSessionMapping) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioSessionMapping) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioSessionMapping) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioSessionMapping) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioSessionMapping) IdentityName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioSessionMapping) IdentityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioSessionMapping) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioSessionMapping) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioSessionMapping) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioSessionMapping) SessionPolicyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionPolicyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioSessionMapping) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioSessionMapping) StudioId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"studioId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStudioSessionMapping) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::EMR::StudioSessionMapping`.
func NewCfnStudioSessionMapping(scope constructs.Construct, id *string, props *CfnStudioSessionMappingProps) CfnStudioSessionMapping {
	_init_.Initialize()

	j := jsiiProxy_CfnStudioSessionMapping{}

	_jsii_.Create(
		"aws-cdk-lib.aws_emr.CfnStudioSessionMapping",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::EMR::StudioSessionMapping`.
func NewCfnStudioSessionMapping_Override(c CfnStudioSessionMapping, scope constructs.Construct, id *string, props *CfnStudioSessionMappingProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_emr.CfnStudioSessionMapping",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnStudioSessionMapping) SetIdentityName(val *string) {
	_jsii_.Set(
		j,
		"identityName",
		val,
	)
}

func (j *jsiiProxy_CfnStudioSessionMapping) SetIdentityType(val *string) {
	_jsii_.Set(
		j,
		"identityType",
		val,
	)
}

func (j *jsiiProxy_CfnStudioSessionMapping) SetSessionPolicyArn(val *string) {
	_jsii_.Set(
		j,
		"sessionPolicyArn",
		val,
	)
}

func (j *jsiiProxy_CfnStudioSessionMapping) SetStudioId(val *string) {
	_jsii_.Set(
		j,
		"studioId",
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
func CfnStudioSessionMapping_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnStudioSessionMapping",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnStudioSessionMapping_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnStudioSessionMapping",
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
func CfnStudioSessionMapping_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_emr.CfnStudioSessionMapping",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStudioSessionMapping_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_emr.CfnStudioSessionMapping",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnStudioSessionMapping) AddDeletionOverride(path *string) {
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
// Experimental.
func (c *jsiiProxy_CfnStudioSessionMapping) AddDependsOn(target awscdk.CfnResource) {
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
// Experimental.
func (c *jsiiProxy_CfnStudioSessionMapping) AddMetadata(key *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnStudioSessionMapping) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnStudioSessionMapping) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
// Experimental.
func (c *jsiiProxy_CfnStudioSessionMapping) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnStudioSessionMapping) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
// Experimental.
func (c *jsiiProxy_CfnStudioSessionMapping) GetAtt(attributeName *string) awscdk.Reference {
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
// Experimental.
func (c *jsiiProxy_CfnStudioSessionMapping) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnStudioSessionMapping) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnStudioSessionMapping) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnStudioSessionMapping) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
// Experimental.
func (c *jsiiProxy_CfnStudioSessionMapping) ShouldSynthesize() *bool {
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
// Experimental.
func (c *jsiiProxy_CfnStudioSessionMapping) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnStudioSessionMapping) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `AWS::EMR::StudioSessionMapping`.
type CfnStudioSessionMappingProps struct {
	// `AWS::EMR::StudioSessionMapping.IdentityName`.
	IdentityName *string `json:"identityName"`
	// `AWS::EMR::StudioSessionMapping.IdentityType`.
	IdentityType *string `json:"identityType"`
	// `AWS::EMR::StudioSessionMapping.SessionPolicyArn`.
	SessionPolicyArn *string `json:"sessionPolicyArn"`
	// `AWS::EMR::StudioSessionMapping.StudioId`.
	StudioId *string `json:"studioId"`
}

