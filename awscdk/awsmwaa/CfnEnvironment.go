package awsmwaa

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmwaa/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::MWAA::Environment` resource creates an Amazon Managed Workflows for Apache Airflow (MWAA) environment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var airflowConfigurationOptions interface{}
//   var tags interface{}
//
//   cfnEnvironment := awscdk.Aws_mwaa.NewCfnEnvironment(this, jsii.String("MyCfnEnvironment"), &CfnEnvironmentProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	AirflowConfigurationOptions: airflowConfigurationOptions,
//   	AirflowVersion: jsii.String("airflowVersion"),
//   	DagS3Path: jsii.String("dagS3Path"),
//   	EndpointManagement: jsii.String("endpointManagement"),
//   	EnvironmentClass: jsii.String("environmentClass"),
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	KmsKey: jsii.String("kmsKey"),
//   	LoggingConfiguration: &LoggingConfigurationProperty{
//   		DagProcessingLogs: &ModuleLoggingConfigurationProperty{
//   			CloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   			Enabled: jsii.Boolean(false),
//   			LogLevel: jsii.String("logLevel"),
//   		},
//   		SchedulerLogs: &ModuleLoggingConfigurationProperty{
//   			CloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   			Enabled: jsii.Boolean(false),
//   			LogLevel: jsii.String("logLevel"),
//   		},
//   		TaskLogs: &ModuleLoggingConfigurationProperty{
//   			CloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   			Enabled: jsii.Boolean(false),
//   			LogLevel: jsii.String("logLevel"),
//   		},
//   		WebserverLogs: &ModuleLoggingConfigurationProperty{
//   			CloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   			Enabled: jsii.Boolean(false),
//   			LogLevel: jsii.String("logLevel"),
//   		},
//   		WorkerLogs: &ModuleLoggingConfigurationProperty{
//   			CloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   			Enabled: jsii.Boolean(false),
//   			LogLevel: jsii.String("logLevel"),
//   		},
//   	},
//   	MaxWebservers: jsii.Number(123),
//   	MaxWorkers: jsii.Number(123),
//   	MinWebservers: jsii.Number(123),
//   	MinWorkers: jsii.Number(123),
//   	NetworkConfiguration: &NetworkConfigurationProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   	PluginsS3ObjectVersion: jsii.String("pluginsS3ObjectVersion"),
//   	PluginsS3Path: jsii.String("pluginsS3Path"),
//   	RequirementsS3ObjectVersion: jsii.String("requirementsS3ObjectVersion"),
//   	RequirementsS3Path: jsii.String("requirementsS3Path"),
//   	Schedulers: jsii.Number(123),
//   	SourceBucketArn: jsii.String("sourceBucketArn"),
//   	StartupScriptS3ObjectVersion: jsii.String("startupScriptS3ObjectVersion"),
//   	StartupScriptS3Path: jsii.String("startupScriptS3Path"),
//   	Tags: tags,
//   	WebserverAccessMode: jsii.String("webserverAccessMode"),
//   	WeeklyMaintenanceWindowStart: jsii.String("weeklyMaintenanceWindowStart"),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mwaa-environment.html
//
type CfnEnvironment interface {
	awscdk.CfnResource
	awscdk.IInspectable
	awscdk.ITaggable
	// A list of key-value pairs containing the Airflow configuration options for your environment.
	AirflowConfigurationOptions() interface{}
	SetAirflowConfigurationOptions(val interface{})
	// The version of Apache Airflow to use for the environment.
	AirflowVersion() *string
	SetAirflowVersion(val *string)
	// The ARN for the Amazon MWAA environment.
	AttrArn() *string
	// The queue ARN for the environment's [Celery Executor](https://docs.aws.amazon.com/https://airflow.apache.org/docs/apache-airflow/stable/core-concepts/executor/celery.html) . Amazon MWAA uses a Celery Executor to distribute tasks across multiple workers. When you create an environment in a shared VPC, you must provide access to the Celery Executor queue from your VPC.
	AttrCeleryExecutorQueue() *string
	// The VPC endpoint for the environment's Amazon RDS database.
	AttrDatabaseVpcEndpointService() *string
	// The ARN for the CloudWatch Logs group where the Apache Airflow DAG processing logs are published.
	AttrLoggingConfigurationDagProcessingLogsCloudWatchLogGroupArn() *string
	// The ARN for the CloudWatch Logs group where the Apache Airflow Scheduler logs are published.
	AttrLoggingConfigurationSchedulerLogsCloudWatchLogGroupArn() *string
	// The ARN for the CloudWatch Logs group where the Apache Airflow task logs are published.
	AttrLoggingConfigurationTaskLogsCloudWatchLogGroupArn() *string
	// The ARN for the CloudWatch Logs group where the Apache Airflow Web server logs are published.
	AttrLoggingConfigurationWebserverLogsCloudWatchLogGroupArn() *string
	// The ARN for the CloudWatch Logs group where the Apache Airflow Worker logs are published.
	AttrLoggingConfigurationWorkerLogsCloudWatchLogGroupArn() *string
	// The URL of your Apache Airflow UI.
	AttrWebserverUrl() *string
	// The VPC endpoint for the environment's web server.
	AttrWebserverVpcEndpointService() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The relative path to the DAGs folder on your Amazon S3 bucket.
	DagS3Path() *string
	SetDagS3Path(val *string)
	// Defines whether the VPC endpoints configured for the environment are created, and managed, by the customer or by Amazon MWAA.
	EndpointManagement() *string
	SetEndpointManagement(val *string)
	// The environment class type.
	EnvironmentClass() *string
	SetEnvironmentClass(val *string)
	// The Amazon Resource Name (ARN) of the execution role in IAM that allows MWAA to access AWS resources in your environment.
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
	// The AWS Key Management Service (KMS) key to encrypt and decrypt the data in your environment.
	KmsKey() *string
	SetKmsKey(val *string)
	// The Apache Airflow logs being sent to CloudWatch Logs: `DagProcessingLogs` , `SchedulerLogs` , `TaskLogs` , `WebserverLogs` , `WorkerLogs` .
	LoggingConfiguration() interface{}
	SetLoggingConfiguration(val interface{})
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
	// The maximum number of web servers that you want to run in your environment.
	MaxWebservers() *float64
	SetMaxWebservers(val *float64)
	// The maximum number of workers that you want to run in your environment.
	MaxWorkers() *float64
	SetMaxWorkers(val *float64)
	// The minimum number of web servers that you want to run in your environment.
	MinWebservers() *float64
	SetMinWebservers(val *float64)
	// The minimum number of workers that you want to run in your environment.
	MinWorkers() *float64
	SetMinWorkers(val *float64)
	// The name of your Amazon MWAA environment.
	Name() *string
	SetName(val *string)
	// The VPC networking components used to secure and enable network traffic between the AWS resources for your environment.
	NetworkConfiguration() interface{}
	SetNetworkConfiguration(val interface{})
	// The tree node.
	Node() constructs.Node
	// The version of the plugins.zip file on your Amazon S3 bucket. To learn more, see [Installing custom plugins](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import-plugins.html) .
	PluginsS3ObjectVersion() *string
	SetPluginsS3ObjectVersion(val *string)
	// The relative path to the `plugins.zip` file on your Amazon S3 bucket. For example, `plugins.zip` . To learn more, see [Installing custom plugins](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import-plugins.html) .
	PluginsS3Path() *string
	SetPluginsS3Path(val *string)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The version of the requirements.txt file on your Amazon S3 bucket. To learn more, see [Installing Python dependencies](https://docs.aws.amazon.com/mwaa/latest/userguide/working-dags-dependencies.html) .
	RequirementsS3ObjectVersion() *string
	SetRequirementsS3ObjectVersion(val *string)
	// The relative path to the `requirements.txt` file on your Amazon S3 bucket. For example, `requirements.txt` . To learn more, see [Installing Python dependencies](https://docs.aws.amazon.com/mwaa/latest/userguide/working-dags-dependencies.html) .
	RequirementsS3Path() *string
	SetRequirementsS3Path(val *string)
	// The number of schedulers that you want to run in your environment.
	//
	// Valid values:.
	Schedulers() *float64
	SetSchedulers(val *float64)
	// The Amazon Resource Name (ARN) of the Amazon S3 bucket where your DAG code and supporting files are stored.
	SourceBucketArn() *string
	SetSourceBucketArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The version of the startup shell script in your Amazon S3 bucket.
	StartupScriptS3ObjectVersion() *string
	SetStartupScriptS3ObjectVersion(val *string)
	// The relative path to the startup shell script in your Amazon S3 bucket.
	//
	// For example, `s3://mwaa-environment/startup.sh` .
	StartupScriptS3Path() *string
	SetStartupScriptS3Path(val *string)
	// Tag Manager which manages the tags for this resource.
	Tags() awscdk.TagManager
	// The key-value tag pairs associated to your environment.
	//
	// For example, `"Environment": "Staging"` . To learn more, see [Tagging](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) .
	TagsRaw() interface{}
	SetTagsRaw(val interface{})
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
	// The Apache Airflow *Web server* access mode.
	WebserverAccessMode() *string
	SetWebserverAccessMode(val *string)
	// The day and time of the week to start weekly maintenance updates of your environment in the following format: `DAY:HH:MM` .
	WeeklyMaintenanceWindowStart() *string
	SetWeeklyMaintenanceWindowStart(val *string)
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
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/resources-section-structure.html#resources-section-structure-logicalid
	//
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

// The jsii proxy struct for CfnEnvironment
type jsiiProxy_CfnEnvironment struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__awscdkITaggable
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

func (j *jsiiProxy_CfnEnvironment) AttrCeleryExecutorQueue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCeleryExecutorQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) AttrDatabaseVpcEndpointService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrDatabaseVpcEndpointService",
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

func (j *jsiiProxy_CfnEnvironment) AttrWebserverVpcEndpointService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrWebserverVpcEndpointService",
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

func (j *jsiiProxy_CfnEnvironment) EndpointManagement() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointManagement",
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

func (j *jsiiProxy_CfnEnvironment) MaxWebservers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWebservers",
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

func (j *jsiiProxy_CfnEnvironment) MinWebservers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minWebservers",
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

func (j *jsiiProxy_CfnEnvironment) StartupScriptS3ObjectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startupScriptS3ObjectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironment) StartupScriptS3Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startupScriptS3Path",
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

func (j *jsiiProxy_CfnEnvironment) TagsRaw() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsRaw",
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

func (j *jsiiProxy_CfnEnvironment) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
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


func NewCfnEnvironment(scope constructs.Construct, id *string, props *CfnEnvironmentProps) CfnEnvironment {
	_init_.Initialize()

	if err := validateNewCfnEnvironmentParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEnvironment{}

	_jsii_.Create(
		"aws-cdk-lib.aws_mwaa.CfnEnvironment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCfnEnvironment_Override(c CfnEnvironment, scope constructs.Construct, id *string, props *CfnEnvironmentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_mwaa.CfnEnvironment",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetAirflowConfigurationOptions(val interface{}) {
	_jsii_.Set(
		j,
		"airflowConfigurationOptions",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetAirflowVersion(val *string) {
	_jsii_.Set(
		j,
		"airflowVersion",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetDagS3Path(val *string) {
	_jsii_.Set(
		j,
		"dagS3Path",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetEndpointManagement(val *string) {
	_jsii_.Set(
		j,
		"endpointManagement",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetEnvironmentClass(val *string) {
	_jsii_.Set(
		j,
		"environmentClass",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetExecutionRoleArn(val *string) {
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetKmsKey(val *string) {
	_jsii_.Set(
		j,
		"kmsKey",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetLoggingConfiguration(val interface{}) {
	if err := j.validateSetLoggingConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggingConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetMaxWebservers(val *float64) {
	_jsii_.Set(
		j,
		"maxWebservers",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetMaxWorkers(val *float64) {
	_jsii_.Set(
		j,
		"maxWorkers",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetMinWebservers(val *float64) {
	_jsii_.Set(
		j,
		"minWebservers",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetMinWorkers(val *float64) {
	_jsii_.Set(
		j,
		"minWorkers",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetNetworkConfiguration(val interface{}) {
	if err := j.validateSetNetworkConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetPluginsS3ObjectVersion(val *string) {
	_jsii_.Set(
		j,
		"pluginsS3ObjectVersion",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetPluginsS3Path(val *string) {
	_jsii_.Set(
		j,
		"pluginsS3Path",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetRequirementsS3ObjectVersion(val *string) {
	_jsii_.Set(
		j,
		"requirementsS3ObjectVersion",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetRequirementsS3Path(val *string) {
	_jsii_.Set(
		j,
		"requirementsS3Path",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetSchedulers(val *float64) {
	_jsii_.Set(
		j,
		"schedulers",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetSourceBucketArn(val *string) {
	_jsii_.Set(
		j,
		"sourceBucketArn",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetStartupScriptS3ObjectVersion(val *string) {
	_jsii_.Set(
		j,
		"startupScriptS3ObjectVersion",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetStartupScriptS3Path(val *string) {
	_jsii_.Set(
		j,
		"startupScriptS3Path",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetTagsRaw(val interface{}) {
	_jsii_.Set(
		j,
		"tagsRaw",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetWebserverAccessMode(val *string) {
	_jsii_.Set(
		j,
		"webserverAccessMode",
		val,
	)
}

func (j *jsiiProxy_CfnEnvironment)SetWeeklyMaintenanceWindowStart(val *string) {
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

	if err := validateCfnEnvironment_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mwaa.CfnEnvironment",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnEnvironment_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEnvironment_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_mwaa.CfnEnvironment",
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
func CfnEnvironment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEnvironment_IsConstructParameters(x); err != nil {
		panic(err)
	}
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

func (c *jsiiProxy_CfnEnvironment) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnEnvironment) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnEnvironment) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnEnvironment) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnEnvironment) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnEnvironment) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnEnvironment) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnEnvironment) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnEnvironment) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnEnvironment) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnEnvironment) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnEnvironment) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEnvironment) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnEnvironment) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnEnvironment) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnEnvironment) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnEnvironment) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

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
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

