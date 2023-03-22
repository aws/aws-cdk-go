package awsglue

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Glue::Job`.
//
// The `AWS::Glue::Job` resource specifies an AWS Glue job in the data catalog. For more information, see [Adding Jobs in AWS Glue](https://docs.aws.amazon.com/glue/latest/dg/add-job.html) and [Job Structure](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-jobs-job.html#aws-glue-api-jobs-job-Job) in the *AWS Glue Developer Guide.*
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var defaultArguments interface{}
//   var nonOverridableArguments interface{}
//   var tags interface{}
//
//   cfnJob := awscdk.Aws_glue.NewCfnJob(this, jsii.String("MyCfnJob"), &CfnJobProps{
//   	Command: &JobCommandProperty{
//   		Name: jsii.String("name"),
//   		PythonVersion: jsii.String("pythonVersion"),
//   		ScriptLocation: jsii.String("scriptLocation"),
//   	},
//   	Role: jsii.String("role"),
//
//   	// the properties below are optional
//   	AllocatedCapacity: jsii.Number(123),
//   	Connections: &ConnectionsListProperty{
//   		Connections: []*string{
//   			jsii.String("connections"),
//   		},
//   	},
//   	DefaultArguments: defaultArguments,
//   	Description: jsii.String("description"),
//   	ExecutionClass: jsii.String("executionClass"),
//   	ExecutionProperty: &ExecutionPropertyProperty{
//   		MaxConcurrentRuns: jsii.Number(123),
//   	},
//   	GlueVersion: jsii.String("glueVersion"),
//   	LogUri: jsii.String("logUri"),
//   	MaxCapacity: jsii.Number(123),
//   	MaxRetries: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	NonOverridableArguments: nonOverridableArguments,
//   	NotificationProperty: &NotificationPropertyProperty{
//   		NotifyDelayAfter: jsii.Number(123),
//   	},
//   	NumberOfWorkers: jsii.Number(123),
//   	SecurityConfiguration: jsii.String("securityConfiguration"),
//   	Tags: tags,
//   	Timeout: jsii.Number(123),
//   	WorkerType: jsii.String("workerType"),
//   })
//
type CfnJob interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// This parameter is no longer supported. Use `MaxCapacity` instead.
	//
	// The number of capacity units that are allocated to this job.
	AllocatedCapacity() *float64
	SetAllocatedCapacity(val *float64)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// The code that executes a job.
	Command() interface{}
	SetCommand(val interface{})
	// The connections used for this job.
	Connections() interface{}
	SetConnections(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The default arguments for this job, specified as name-value pairs.
	//
	// You can specify arguments here that your own job-execution script consumes, in addition to arguments that AWS Glue itself consumes.
	//
	// For information about how to specify and consume your own job arguments, see [Calling AWS Glue APIs in Python](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) in the *AWS Glue Developer Guide* .
	//
	// For information about the key-value pairs that AWS Glue consumes to set up your job, see [Special Parameters Used by AWS Glue](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html) in the *AWS Glue Developer Guide* .
	DefaultArguments() interface{}
	SetDefaultArguments(val interface{})
	// A description of the job.
	Description() *string
	SetDescription(val *string)
	// Indicates whether the job is run with a standard or flexible execution class.
	//
	// The standard execution class is ideal for time-sensitive workloads that require fast job startup and dedicated resources.
	//
	// The flexible execution class is appropriate for time-insensitive jobs whose start and completion times may vary.
	//
	// Only jobs with AWS Glue version 3.0 and above and command type `glueetl` will be allowed to set `ExecutionClass` to `FLEX` . The flexible execution class is available for Spark jobs.
	ExecutionClass() *string
	SetExecutionClass(val *string)
	// The maximum number of concurrent runs that are allowed for this job.
	ExecutionProperty() interface{}
	SetExecutionProperty(val interface{})
	// Glue version determines the versions of Apache Spark and Python that AWS Glue supports.
	//
	// The Python version indicates the version supported for jobs of type Spark.
	//
	// For more information about the available AWS Glue versions and corresponding Spark and Python versions, see [Glue version](https://docs.aws.amazon.com/glue/latest/dg/add-job.html) in the developer guide.
	//
	// Jobs that are created without specifying a Glue version default to Glue 0.9.
	GlueVersion() *string
	SetGlueVersion(val *string)
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
	// This field is reserved for future use.
	LogUri() *string
	SetLogUri(val *string)
	// The number of AWS Glue data processing units (DPUs) that can be allocated when this job runs.
	//
	// A DPU is a relative measure of processing power that consists of 4 vCPUs of compute capacity and 16 GB of memory.
	//
	// Do not set `Max Capacity` if using `WorkerType` and `NumberOfWorkers` .
	//
	// The value that can be allocated for `MaxCapacity` depends on whether you are running a Python shell job or an Apache Spark ETL job:
	//
	// - When you specify a Python shell job ( `JobCommand.Name` ="pythonshell"), you can allocate either 0.0625 or 1 DPU. The default is 0.0625 DPU.
	// - When you specify an Apache Spark ETL job ( `JobCommand.Name` ="glueetl"), you can allocate from 2 to 100 DPUs. The default is 10 DPUs. This job type cannot have a fractional DPU allocation.
	MaxCapacity() *float64
	SetMaxCapacity(val *float64)
	// The maximum number of times to retry this job after a JobRun fails.
	MaxRetries() *float64
	SetMaxRetries(val *float64)
	// The name you assign to this job definition.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Non-overridable arguments for this job, specified as name-value pairs.
	NonOverridableArguments() interface{}
	SetNonOverridableArguments(val interface{})
	// Specifies configuration properties of a notification.
	NotificationProperty() interface{}
	SetNotificationProperty(val interface{})
	// The number of workers of a defined `workerType` that are allocated when a job runs.
	//
	// The maximum number of workers you can define are 299 for `G.1X` , and 149 for `G.2X` .
	NumberOfWorkers() *float64
	SetNumberOfWorkers(val *float64)
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The name or Amazon Resource Name (ARN) of the IAM role associated with this job.
	Role() *string
	SetRole(val *string)
	// The name of the `SecurityConfiguration` structure to be used with this job.
	SecurityConfiguration() *string
	SetSecurityConfiguration(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The tags to use with this job.
	Tags() awscdk.TagManager
	// The job timeout in minutes.
	//
	// This is the maximum time that a job run can consume resources before it is terminated and enters TIMEOUT status. The default is 2,880 minutes (48 hours).
	Timeout() *float64
	SetTimeout(val *float64)
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
	// The type of predefined worker that is allocated when a job runs.
	//
	// Accepts a value of Standard, G.1X, or G.2X.
	//
	// - For the `Standard` worker type, each worker provides 4 vCPU, 16 GB of memory and a 50GB disk, and 2 executors per worker.
	// - For the `G.1X` worker type, each worker maps to 1 DPU (4 vCPU, 16 GB of memory, 64 GB disk), and provides 1 executor per worker. We recommend this worker type for memory-intensive jobs.
	// - For the `G.2X` worker type, each worker maps to 2 DPU (8 vCPU, 32 GB of memory, 128 GB disk), and provides 1 executor per worker. We recommend this worker type for memory-intensive jobs.
	WorkerType() *string
	SetWorkerType(val *string)
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

// The jsii proxy struct for CfnJob
type jsiiProxy_CfnJob struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnJob) AllocatedCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Command() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Connections() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) DefaultArguments() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) ExecutionClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) ExecutionProperty() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"executionProperty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) GlueVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"glueVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) LogUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) MaxCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) NonOverridableArguments() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nonOverridableArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) NotificationProperty() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationProperty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) NumberOfWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) SecurityConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJob) WorkerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workerType",
		&returns,
	)
	return returns
}


// Create a new `AWS::Glue::Job`.
func NewCfnJob(scope constructs.Construct, id *string, props *CfnJobProps) CfnJob {
	_init_.Initialize()

	if err := validateNewCfnJobParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnJob{}

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnJob",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Glue::Job`.
func NewCfnJob_Override(c CfnJob, scope constructs.Construct, id *string, props *CfnJobProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_glue.CfnJob",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnJob)SetAllocatedCapacity(val *float64) {
	_jsii_.Set(
		j,
		"allocatedCapacity",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetCommand(val interface{}) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetConnections(val interface{}) {
	if err := j.validateSetConnectionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connections",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetDefaultArguments(val interface{}) {
	if err := j.validateSetDefaultArgumentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultArguments",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetExecutionClass(val *string) {
	_jsii_.Set(
		j,
		"executionClass",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetExecutionProperty(val interface{}) {
	if err := j.validateSetExecutionPropertyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionProperty",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetGlueVersion(val *string) {
	_jsii_.Set(
		j,
		"glueVersion",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetLogUri(val *string) {
	_jsii_.Set(
		j,
		"logUri",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetMaxCapacity(val *float64) {
	_jsii_.Set(
		j,
		"maxCapacity",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetMaxRetries(val *float64) {
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetNonOverridableArguments(val interface{}) {
	if err := j.validateSetNonOverridableArgumentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nonOverridableArguments",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetNotificationProperty(val interface{}) {
	if err := j.validateSetNotificationPropertyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationProperty",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetNumberOfWorkers(val *float64) {
	_jsii_.Set(
		j,
		"numberOfWorkers",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetRole(val *string) {
	if err := j.validateSetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetSecurityConfiguration(val *string) {
	_jsii_.Set(
		j,
		"securityConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetTimeout(val *float64) {
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_CfnJob)SetWorkerType(val *string) {
	_jsii_.Set(
		j,
		"workerType",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnJob_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnJob_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnJob",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnJob_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnJob_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnJob",
		"isCfnResource",
		[]interface{}{construct},
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
func CfnJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_glue.CfnJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnJob_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_glue.CfnJob",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnJob) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnJob) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnJob) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnJob) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnJob) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnJob) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnJob) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnJob) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnJob) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
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

func (c *jsiiProxy_CfnJob) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnJob) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnJob) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnJob) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnJob) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnJob) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnJob) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnJob) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnJob) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnJob) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

