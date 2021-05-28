package awsmwaa

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmwaa/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::MWAA::Environment`.
type CfnEnvironment interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AirflowConfigurationOptions() interface{}
	SetAirflowConfigurationOptions(val interface{})
	AirflowVersion() *string
	SetAirflowVersion(val *string)
	AttrArn() *string
	AttrWebserverUrl() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DagS3Path() *string
	SetDagS3Path(val *string)
	EnvironmentClass() *string
	SetEnvironmentClass(val *string)
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
	KmsKey() *string
	SetKmsKey(val *string)
	LoggingConfiguration() interface{}
	SetLoggingConfiguration(val interface{})
	LogicalId() *string
	MaxWorkers() *float64
	SetMaxWorkers(val *float64)
	MinWorkers() *float64
	SetMinWorkers(val *float64)
	Name() *string
	SetName(val *string)
	NetworkConfiguration() interface{}
	SetNetworkConfiguration(val interface{})
	Node() constructs.Node
	PluginsS3ObjectVersion() *string
	SetPluginsS3ObjectVersion(val *string)
	PluginsS3Path() *string
	SetPluginsS3Path(val *string)
	Ref() *string
	RequirementsS3ObjectVersion() *string
	SetRequirementsS3ObjectVersion(val *string)
	RequirementsS3Path() *string
	SetRequirementsS3Path(val *string)
	SourceBucketArn() *string
	SetSourceBucketArn(val *string)
	Stack() awscdk.Stack
	Tags() *CfnEnvironment_TagMapProperty
	SetTags(val *CfnEnvironment_TagMapProperty)
	UpdatedProperites() *map[string]interface{}
	WebserverAccessMode() *string
	SetWebserverAccessMode(val *string)
	WeeklyMaintenanceWindowStart() *string
	SetWeeklyMaintenanceWindowStart(val *string)
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

// The jsii proxy struct for CfnEnvironment
type jsiiProxy_CfnEnvironment struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnEnvironment) AirflowConfigurationOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"airflowConfigurationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) AirflowVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"airflowVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) AttrWebserverUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrWebserverUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) DagS3Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dagS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) EnvironmentClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) KmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) LoggingConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) MaxWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) MinWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) NetworkConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) PluginsS3ObjectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginsS3ObjectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) PluginsS3Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginsS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) RequirementsS3ObjectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requirementsS3ObjectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) RequirementsS3Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requirementsS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) SourceBucketArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceBucketArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) Tags() *CfnEnvironment_TagMapProperty {
	var returns *CfnEnvironment_TagMapProperty
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) WebserverAccessMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webserverAccessMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) WeeklyMaintenanceWindowStart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyMaintenanceWindowStart",
		&returns,
	)
	return returns
}


// Create a new `AWS::MWAA::Environment`.
func NewCfnEnvironment(scope constructs.Construct, id *string, props *CfnEnvironmentProps) CfnEnvironment {
	_init_.Initialize()

	j := jsiiProxy_CfnEnvironment{}

	_jsii_.Create(
		"aws-cdk-lib.aws_mwaa.CfnEnvironment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::MWAA::Environment`.
func NewCfnEnvironment_Override(c CfnEnvironment, scope constructs.Construct, id *string, props *CfnEnvironmentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_mwaa.CfnEnvironment",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnEnvironment) SetAirflowConfigurationOptions(val interface{}) {
	_jsii_.Set(
		j,
		"airflowConfigurationOptions",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment) SetAirflowVersion(val *string) {
	_jsii_.Set(
		j,
		"airflowVersion",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment) SetDagS3Path(val *string) {
	_jsii_.Set(
		j,
		"dagS3Path",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment) SetEnvironmentClass(val *string) {
	_jsii_.Set(
		j,
		"environmentClass",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment) SetExecutionRoleArn(val *string) {
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment) SetKmsKey(val *string) {
	_jsii_.Set(
		j,
		"kmsKey",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment) SetLoggingConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"loggingConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment) SetMaxWorkers(val *float64) {
	_jsii_.Set(
		j,
		"maxWorkers",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment) SetMinWorkers(val *float64) {
	_jsii_.Set(
		j,
		"minWorkers",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment) SetNetworkConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"networkConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment) SetPluginsS3ObjectVersion(val *string) {
	_jsii_.Set(
		j,
		"pluginsS3ObjectVersion",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment) SetPluginsS3Path(val *string) {
	_jsii_.Set(
		j,
		"pluginsS3Path",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment) SetRequirementsS3ObjectVersion(val *string) {
	_jsii_.Set(
		j,
		"requirementsS3ObjectVersion",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment) SetRequirementsS3Path(val *string) {
	_jsii_.Set(
		j,
		"requirementsS3Path",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment) SetSourceBucketArn(val *string) {
	_jsii_.Set(
		j,
		"sourceBucketArn",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment) SetTags(val *CfnEnvironment_TagMapProperty) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment) SetWebserverAccessMode(val *string) {
	_jsii_.Set(
		j,
		"webserverAccessMode",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment) SetWeeklyMaintenanceWindowStart(val *string) {
	_jsii_.Set(
		j,
		"weeklyMaintenanceWindowStart",
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
func CfnEnvironment_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mwaa.CfnEnvironment",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnEnvironment_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mwaa.CfnEnvironment",
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
func CfnEnvironment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mwaa.CfnEnvironment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEnvironment_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_mwaa.CfnEnvironment",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnEnvironment) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnEnvironment) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnEnvironment) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnEnvironment) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnEnvironment) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnEnvironment) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

// Sets the deletion policy of the resource based on the removal policy specified.
// Experimental.
func (c *jsiiProxy_CfnEnvironment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnEnvironment) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnEnvironment) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnEnvironment) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnEnvironment) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnEnvironment) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnEnvironment) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnEnvironment) ToString() *string {
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
func (c *jsiiProxy_CfnEnvironment) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnEnvironment_LoggingConfigurationProperty struct {
	// `CfnEnvironment.LoggingConfigurationProperty.DagProcessingLogs`.
	DagProcessingLogs interface{} `json:"dagProcessingLogs"`
	// `CfnEnvironment.LoggingConfigurationProperty.SchedulerLogs`.
	SchedulerLogs interface{} `json:"schedulerLogs"`
	// `CfnEnvironment.LoggingConfigurationProperty.TaskLogs`.
	TaskLogs interface{} `json:"taskLogs"`
	// `CfnEnvironment.LoggingConfigurationProperty.WebserverLogs`.
	WebserverLogs interface{} `json:"webserverLogs"`
	// `CfnEnvironment.LoggingConfigurationProperty.WorkerLogs`.
	WorkerLogs interface{} `json:"workerLogs"`
}

type CfnEnvironment_ModuleLoggingConfigurationProperty struct {
	// `CfnEnvironment.ModuleLoggingConfigurationProperty.CloudWatchLogGroupArn`.
	CloudWatchLogGroupArn *string `json:"cloudWatchLogGroupArn"`
	// `CfnEnvironment.ModuleLoggingConfigurationProperty.Enabled`.
	Enabled interface{} `json:"enabled"`
	// `CfnEnvironment.ModuleLoggingConfigurationProperty.LogLevel`.
	LogLevel *string `json:"logLevel"`
}

type CfnEnvironment_NetworkConfigurationProperty struct {
	// `CfnEnvironment.NetworkConfigurationProperty.SecurityGroupIds`.
	SecurityGroupIds *[]*string `json:"securityGroupIds"`
	// `CfnEnvironment.NetworkConfigurationProperty.SubnetIds`.
	SubnetIds *[]*string `json:"subnetIds"`
}

type CfnEnvironment_TagMapProperty struct {
}

// Properties for defining a `AWS::MWAA::Environment`.
type CfnEnvironmentProps struct {
	// `AWS::MWAA::Environment.Name`.
	Name *string `json:"name"`
	// `AWS::MWAA::Environment.AirflowConfigurationOptions`.
	AirflowConfigurationOptions interface{} `json:"airflowConfigurationOptions"`
	// `AWS::MWAA::Environment.AirflowVersion`.
	AirflowVersion *string `json:"airflowVersion"`
	// `AWS::MWAA::Environment.DagS3Path`.
	DagS3Path *string `json:"dagS3Path"`
	// `AWS::MWAA::Environment.EnvironmentClass`.
	EnvironmentClass *string `json:"environmentClass"`
	// `AWS::MWAA::Environment.ExecutionRoleArn`.
	ExecutionRoleArn *string `json:"executionRoleArn"`
	// `AWS::MWAA::Environment.KmsKey`.
	KmsKey *string `json:"kmsKey"`
	// `AWS::MWAA::Environment.LoggingConfiguration`.
	LoggingConfiguration interface{} `json:"loggingConfiguration"`
	// `AWS::MWAA::Environment.MaxWorkers`.
	MaxWorkers *float64 `json:"maxWorkers"`
	// `AWS::MWAA::Environment.MinWorkers`.
	MinWorkers *float64 `json:"minWorkers"`
	// `AWS::MWAA::Environment.NetworkConfiguration`.
	NetworkConfiguration interface{} `json:"networkConfiguration"`
	// `AWS::MWAA::Environment.PluginsS3ObjectVersion`.
	PluginsS3ObjectVersion *string `json:"pluginsS3ObjectVersion"`
	// `AWS::MWAA::Environment.PluginsS3Path`.
	PluginsS3Path *string `json:"pluginsS3Path"`
	// `AWS::MWAA::Environment.RequirementsS3ObjectVersion`.
	RequirementsS3ObjectVersion *string `json:"requirementsS3ObjectVersion"`
	// `AWS::MWAA::Environment.RequirementsS3Path`.
	RequirementsS3Path *string `json:"requirementsS3Path"`
	// `AWS::MWAA::Environment.SourceBucketArn`.
	SourceBucketArn *string `json:"sourceBucketArn"`
	// `AWS::MWAA::Environment.Tags`.
	Tags *CfnEnvironment_TagMapProperty `json:"tags"`
	// `AWS::MWAA::Environment.WebserverAccessMode`.
	WebserverAccessMode *string `json:"webserverAccessMode"`
	// `AWS::MWAA::Environment.WeeklyMaintenanceWindowStart`.
	WeeklyMaintenanceWindowStart *string `json:"weeklyMaintenanceWindowStart"`
}

