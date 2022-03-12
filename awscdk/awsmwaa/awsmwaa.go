package awsmwaa

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmwaa/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::MWAA::Environment`.
//
// The `AWS::MWAA::Environment` resource creates an Amazon Managed Workflows for Apache Airflow (MWAA) environment.
//
// TODO: EXAMPLE
//
type CfnEnvironment interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AirflowConfigurationOptions() interface{}
	SetAirflowConfigurationOptions(val interface{})
	AirflowVersion() *string
	SetAirflowVersion(val *string)
	AttrArn() *string
	AttrLoggingConfigurationDagProcessingLogsCloudWatchLogGroupArn() *string
	AttrLoggingConfigurationSchedulerLogsCloudWatchLogGroupArn() *string
	AttrLoggingConfigurationTaskLogsCloudWatchLogGroupArn() *string
	AttrLoggingConfigurationWebserverLogsCloudWatchLogGroupArn() *string
	AttrLoggingConfigurationWorkerLogsCloudWatchLogGroupArn() *string
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
	Schedulers() *float64
	SetSchedulers(val *float64)
	SourceBucketArn() *string
	SetSourceBucketArn(val *string)
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
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

func (j *jsiiProxy_CfnEnvironment) AttrLoggingConfigurationDagProcessingLogsCloudWatchLogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLoggingConfigurationDagProcessingLogsCloudWatchLogGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) AttrLoggingConfigurationSchedulerLogsCloudWatchLogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLoggingConfigurationSchedulerLogsCloudWatchLogGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) AttrLoggingConfigurationTaskLogsCloudWatchLogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLoggingConfigurationTaskLogsCloudWatchLogGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) AttrLoggingConfigurationWebserverLogsCloudWatchLogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLoggingConfigurationWebserverLogsCloudWatchLogGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) AttrLoggingConfigurationWorkerLogsCloudWatchLogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLoggingConfigurationWorkerLogsCloudWatchLogGroupArn",
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

func (j *jsiiProxy_CfnEnvironment) Schedulers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"schedulers",
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

func (j *jsiiProxy_CfnEnvironment) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
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

func (j *jsiiProxy_CfnEnvironment) SetSchedulers(val *float64) {
	_jsii_.Set(
		j,
		"schedulers",
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
//
// The `value` argument to `addOverride` will not be processed or translated
// in any way. Pass raw JSON values in here with the correct capitalization
// for CloudFormation. If you pass CDK classes or structs, they will be
// rendered with lowercased key names, and CloudFormation will reject the
// template.
func (c *jsiiProxy_CfnEnvironment) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
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
func (c *jsiiProxy_CfnEnvironment) AddPropertyOverride(propertyPath *string, value interface{}) {
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

func (c *jsiiProxy_CfnEnvironment) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// The type of Apache Airflow logs to send to CloudWatch Logs.
//
// TODO: EXAMPLE
//
type CfnEnvironment_LoggingConfigurationProperty struct {
	// Defines the processing logs sent to CloudWatch Logs and the logging level to send.
	DagProcessingLogs interface{} `json:"dagProcessingLogs" yaml:"dagProcessingLogs"`
	// Defines the scheduler logs sent to CloudWatch Logs and the logging level to send.
	SchedulerLogs interface{} `json:"schedulerLogs" yaml:"schedulerLogs"`
	// Defines the task logs sent to CloudWatch Logs and the logging level to send.
	TaskLogs interface{} `json:"taskLogs" yaml:"taskLogs"`
	// Defines the web server logs sent to CloudWatch Logs and the logging level to send.
	WebserverLogs interface{} `json:"webserverLogs" yaml:"webserverLogs"`
	// Defines the worker logs sent to CloudWatch Logs and the logging level to send.
	WorkerLogs interface{} `json:"workerLogs" yaml:"workerLogs"`
}

// Defines the type of logs to send for the Apache Airflow log type (e.g. `DagProcessingLogs` ).
//
// TODO: EXAMPLE
//
type CfnEnvironment_ModuleLoggingConfigurationProperty struct {
	// The ARN of the CloudWatch Logs log group for each type of Apache Airflow log type that you have enabled.
	//
	// > `CloudWatchLogGroupArn` is available only as a return value, accessible when specified as an attribute in the [`Fn:GetAtt`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html#aws-resource-mwaa-environment-return-values) intrinsic function. Any value you provide for `CloudWatchLogGroupArn` is discarded by Amazon MWAA.
	CloudWatchLogGroupArn *string `json:"cloudWatchLogGroupArn" yaml:"cloudWatchLogGroupArn"`
	// Indicates whether to enable the Apache Airflow log type (e.g. `DagProcessingLogs` ) in CloudWatch Logs.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// Defines the Apache Airflow logs to send for the log type (e.g. `DagProcessingLogs` ) to CloudWatch Logs. Valid values: `CRITICAL` , `ERROR` , `WARNING` , `INFO` .
	LogLevel *string `json:"logLevel" yaml:"logLevel"`
}

// The VPC networking components used to secure and enable network traffic between the AWS resources for your environment.
//
// To learn more, see [About networking on Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/networking-about.html) .
//
// TODO: EXAMPLE
//
type CfnEnvironment_NetworkConfigurationProperty struct {
	// A list of one or more security group IDs.
	//
	// Accepts up to 5 security group IDs. A security group must be attached to the same VPC as the subnets. To learn more, see [Security in your VPC on Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/vpc-security.html) .
	SecurityGroupIds *[]*string `json:"securityGroupIds" yaml:"securityGroupIds"`
	// A list of subnet IDs.
	//
	// *Required* to create an environment. Must be private subnets in two different availability zones. A subnet must be attached to the same VPC as the security group. To learn more, see [About networking on Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/networking-about.html) .
	SubnetIds *[]*string `json:"subnetIds" yaml:"subnetIds"`
}

// Properties for defining a `CfnEnvironment`.
//
// TODO: EXAMPLE
//
type CfnEnvironmentProps struct {
	// The name of your Amazon MWAA environment.
	Name *string `json:"name" yaml:"name"`
	// A list of key-value pairs containing the Airflow configuration options for your environment.
	//
	// For example, `core.default_timezone: utc` . To learn more, see [Apache Airflow configuration options](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-env-variables.html) .
	AirflowConfigurationOptions interface{} `json:"airflowConfigurationOptions" yaml:"airflowConfigurationOptions"`
	// The version of Apache Airflow to use for the environment.
	//
	// If no value is specified, defaults to the latest version. Valid values: `2.0.2` , `1.10.12` .
	AirflowVersion *string `json:"airflowVersion" yaml:"airflowVersion"`
	// The relative path to the DAGs folder on your Amazon S3 bucket.
	//
	// For example, `dags` . To learn more, see [Adding or updating DAGs](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-folder.html) .
	DagS3Path *string `json:"dagS3Path" yaml:"dagS3Path"`
	// The environment class type.
	//
	// Valid values: `mw1.small` , `mw1.medium` , `mw1.large` . To learn more, see [Amazon MWAA environment class](https://docs.aws.amazon.com/mwaa/latest/userguide/environment-class.html) .
	EnvironmentClass *string `json:"environmentClass" yaml:"environmentClass"`
	// The Amazon Resource Name (ARN) of the execution role in IAM that allows MWAA to access AWS resources in your environment.
	//
	// For example, `arn:aws:iam::123456789:role/my-execution-role` . To learn more, see [Amazon MWAA Execution role](https://docs.aws.amazon.com/mwaa/latest/userguide/mwaa-create-role.html) .
	ExecutionRoleArn *string `json:"executionRoleArn" yaml:"executionRoleArn"`
	// The AWS Key Management Service (KMS) key to encrypt and decrypt the data in your environment.
	//
	// You can use an AWS KMS key managed by MWAA, or a customer-managed KMS key (advanced).
	KmsKey *string `json:"kmsKey" yaml:"kmsKey"`
	// The Apache Airflow logs being sent to CloudWatch Logs: `DagProcessingLogs` , `SchedulerLogs` , `TaskLogs` , `WebserverLogs` , `WorkerLogs` .
	LoggingConfiguration interface{} `json:"loggingConfiguration" yaml:"loggingConfiguration"`
	// The maximum number of workers that you want to run in your environment.
	//
	// MWAA scales the number of Apache Airflow workers up to the number you specify in the `MaxWorkers` field. For example, `20` . When there are no more tasks running, and no more in the queue, MWAA disposes of the extra workers leaving the one worker that is included with your environment, or the number you specify in `MinWorkers` .
	MaxWorkers *float64 `json:"maxWorkers" yaml:"maxWorkers"`
	// The minimum number of workers that you want to run in your environment.
	//
	// MWAA scales the number of Apache Airflow workers up to the number you specify in the `MaxWorkers` field. When there are no more tasks running, and no more in the queue, MWAA disposes of the extra workers leaving the worker count you specify in the `MinWorkers` field. For example, `2` .
	MinWorkers *float64 `json:"minWorkers" yaml:"minWorkers"`
	// The VPC networking components used to secure and enable network traffic between the AWS resources for your environment.
	//
	// To learn more, see [About networking on Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/networking-about.html) .
	NetworkConfiguration interface{} `json:"networkConfiguration" yaml:"networkConfiguration"`
	// The version of the plugins.zip file on your Amazon S3 bucket. To learn more, see [Installing custom plugins](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import-plugins.html) .
	PluginsS3ObjectVersion *string `json:"pluginsS3ObjectVersion" yaml:"pluginsS3ObjectVersion"`
	// The relative path to the `plugins.zip` file on your Amazon S3 bucket. For example, `plugins.zip` . To learn more, see [Installing custom plugins](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import-plugins.html) .
	PluginsS3Path *string `json:"pluginsS3Path" yaml:"pluginsS3Path"`
	// The version of the requirements.txt file on your Amazon S3 bucket. To learn more, see [Installing Python dependencies](https://docs.aws.amazon.com/mwaa/latest/userguide/working-dags-dependencies.html) .
	RequirementsS3ObjectVersion *string `json:"requirementsS3ObjectVersion" yaml:"requirementsS3ObjectVersion"`
	// The relative path to the `requirements.txt` file on your Amazon S3 bucket. For example, `requirements.txt` . To learn more, see [Installing Python dependencies](https://docs.aws.amazon.com/mwaa/latest/userguide/working-dags-dependencies.html) .
	RequirementsS3Path *string `json:"requirementsS3Path" yaml:"requirementsS3Path"`
	// The number of schedulers that you want to run in your environment. Valid values:.
	//
	// - *v2.0.2* - Accepts between 2 to 5. Defaults to 2.
	// - *v1.10.12* - Accepts 1.
	Schedulers *float64 `json:"schedulers" yaml:"schedulers"`
	// The Amazon Resource Name (ARN) of the Amazon S3 bucket where your DAG code and supporting files are stored.
	//
	// For example, `arn:aws:s3:::my-airflow-bucket-unique-name` . To learn more, see [Create an Amazon S3 bucket for Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/mwaa-s3-bucket.html) .
	SourceBucketArn *string `json:"sourceBucketArn" yaml:"sourceBucketArn"`
	// The key-value tag pairs associated to your environment.
	//
	// For example, `"Environment": "Staging"` . To learn more, see [Tagging](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) .
	Tags interface{} `json:"tags" yaml:"tags"`
	// The Apache Airflow *Web server* access mode.
	//
	// To learn more, see [Apache Airflow access modes](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-networking.html) . Valid values: `PRIVATE_ONLY` or `PUBLIC_ONLY` .
	WebserverAccessMode *string `json:"webserverAccessMode" yaml:"webserverAccessMode"`
	// The day and time of the week to start weekly maintenance updates of your environment in the following format: `DAY:HH:MM` .
	//
	// For example: `TUE:03:30` . You can specify a start time in 30 minute increments only. Supported input includes the following:
	//
	// - MON|TUE|WED|THU|FRI|SAT|SUN:([01]\\d|2[0-3]):(00|30)
	WeeklyMaintenanceWindowStart *string `json:"weeklyMaintenanceWindowStart" yaml:"weeklyMaintenanceWindowStart"`
}

