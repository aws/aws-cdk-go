package awsdatasync

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::DataSync::Task`.
//
// The `AWS::DataSync::Task` resource specifies a task. A task is a set of two locations (source and destination) and a set of `Options` that you use to control the behavior of a task. If you don't specify `Options` when you create a task, AWS DataSync populates them with service defaults.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTask := awscdk.Aws_datasync.NewCfnTask(this, jsii.String("MyCfnTask"), &CfnTaskProps{
//   	DestinationLocationArn: jsii.String("destinationLocationArn"),
//   	SourceLocationArn: jsii.String("sourceLocationArn"),
//
//   	// the properties below are optional
//   	CloudWatchLogGroupArn: jsii.String("cloudWatchLogGroupArn"),
//   	Excludes: []interface{}{
//   		&FilterRuleProperty{
//   			FilterType: jsii.String("filterType"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Includes: []interface{}{
//   		&FilterRuleProperty{
//   			FilterType: jsii.String("filterType"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Options: &OptionsProperty{
//   		Atime: jsii.String("atime"),
//   		BytesPerSecond: jsii.Number(123),
//   		Gid: jsii.String("gid"),
//   		LogLevel: jsii.String("logLevel"),
//   		Mtime: jsii.String("mtime"),
//   		ObjectTags: jsii.String("objectTags"),
//   		OverwriteMode: jsii.String("overwriteMode"),
//   		PosixPermissions: jsii.String("posixPermissions"),
//   		PreserveDeletedFiles: jsii.String("preserveDeletedFiles"),
//   		PreserveDevices: jsii.String("preserveDevices"),
//   		SecurityDescriptorCopyFlags: jsii.String("securityDescriptorCopyFlags"),
//   		TaskQueueing: jsii.String("taskQueueing"),
//   		TransferMode: jsii.String("transferMode"),
//   		Uid: jsii.String("uid"),
//   		VerifyMode: jsii.String("verifyMode"),
//   	},
//   	Schedule: &TaskScheduleProperty{
//   		ScheduleExpression: jsii.String("scheduleExpression"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnTask interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ARNs of the destination elastic network interfaces (ENIs) that were created for your subnet.
	AttrDestinationNetworkInterfaceArns() *[]*string
	// The ARNs of the source ENIs that were created for your subnet.
	AttrSourceNetworkInterfaceArns() *[]*string
	// The status of the task that was described.
	AttrStatus() *string
	// The ARN of the task.
	AttrTaskArn() *string
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// The Amazon Resource Name (ARN) of the Amazon CloudWatch log group that is used to monitor and log events in the task.
	//
	// For more information about how to use CloudWatch Logs with DataSync, see [Monitoring Your Task](https://docs.aws.amazon.com/datasync/latest/userguide/monitor-datasync.html#cloudwatchlogs) in the *AWS DataSync User Guide.*
	//
	// For more information about these groups, see [Working with Log Groups and Log Streams](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/Working-with-log-groups-and-streams.html) in the *Amazon CloudWatch Logs User Guide* .
	CloudWatchLogGroupArn() *string
	SetCloudWatchLogGroupArn(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// The Amazon Resource Name (ARN) of an AWS storage resource's location.
	DestinationLocationArn() *string
	SetDestinationLocationArn(val *string)
	// Specifies a list of filter rules that exclude specific data during your transfer.
	//
	// For more information and examples, see [Filtering data transferred by DataSync](https://docs.aws.amazon.com/datasync/latest/userguide/filtering.html) .
	Excludes() interface{}
	SetExcludes(val interface{})
	// Specifies a list of filter rules that include specific data during your transfer.
	//
	// For more information and examples, see [Filtering data transferred by DataSync](https://docs.aws.amazon.com/datasync/latest/userguide/filtering.html) .
	Includes() interface{}
	SetIncludes(val interface{})
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
	// The name of a task.
	//
	// This value is a text reference that is used to identify the task in the console.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Specifies the configuration options for a task. Some options include preserving file or object metadata and verifying data integrity.
	//
	// You can also override these options before starting an individual run of a task (also known as a *task execution* ). For more information, see [StartTaskExecution](https://docs.aws.amazon.com/datasync/latest/userguide/API_StartTaskExecution.html) .
	Options() interface{}
	SetOptions(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// Specifies a schedule used to periodically transfer files from a source to a destination location.
	//
	// The schedule should be specified in UTC time. For more information, see [Scheduling your task](https://docs.aws.amazon.com/datasync/latest/userguide/task-scheduling.html) .
	Schedule() interface{}
	SetSchedule(val interface{})
	// The Amazon Resource Name (ARN) of the source location for the task.
	SourceLocationArn() *string
	SetSourceLocationArn(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// Specifies the tags that you want to apply to the Amazon Resource Name (ARN) representing the task.
	//
	// *Tags* are key-value pairs that help you manage, filter, and search for your DataSync resources.
	Tags() awscdk.TagManager
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

// The jsii proxy struct for CfnTask
type jsiiProxy_CfnTask struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnTask) AttrDestinationNetworkInterfaceArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrDestinationNetworkInterfaceArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) AttrSourceNetworkInterfaceArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"attrSourceNetworkInterfaceArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) AttrTaskArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrTaskArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) CloudWatchLogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudWatchLogGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) DestinationLocationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationLocationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) Excludes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) Includes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) Options() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) Schedule() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) SourceLocationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceLocationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTask) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::DataSync::Task`.
func NewCfnTask(scope awscdk.Construct, id *string, props *CfnTaskProps) CfnTask {
	_init_.Initialize()

	if err := validateNewCfnTaskParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTask{}

	_jsii_.Create(
		"monocdk.aws_datasync.CfnTask",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DataSync::Task`.
func NewCfnTask_Override(c CfnTask, scope awscdk.Construct, id *string, props *CfnTaskProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_datasync.CfnTask",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTask)SetCloudWatchLogGroupArn(val *string) {
	_jsii_.Set(
		j,
		"cloudWatchLogGroupArn",
		val,
	)
}

func (j *jsiiProxy_CfnTask)SetDestinationLocationArn(val *string) {
	if err := j.validateSetDestinationLocationArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationLocationArn",
		val,
	)
}

func (j *jsiiProxy_CfnTask)SetExcludes(val interface{}) {
	if err := j.validateSetExcludesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludes",
		val,
	)
}

func (j *jsiiProxy_CfnTask)SetIncludes(val interface{}) {
	if err := j.validateSetIncludesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includes",
		val,
	)
}

func (j *jsiiProxy_CfnTask)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnTask)SetOptions(val interface{}) {
	if err := j.validateSetOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"options",
		val,
	)
}

func (j *jsiiProxy_CfnTask)SetSchedule(val interface{}) {
	if err := j.validateSetScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_CfnTask)SetSourceLocationArn(val *string) {
	if err := j.validateSetSourceLocationArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceLocationArn",
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
func CfnTask_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTask_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_datasync.CfnTask",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnTask_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnTask_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_datasync.CfnTask",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnTask_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTask_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_datasync.CfnTask",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTask_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_datasync.CfnTask",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTask) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnTask) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTask) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnTask) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnTask) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnTask) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnTask) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnTask) GetAtt(attributeName *string) awscdk.Reference {
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

func (c *jsiiProxy_CfnTask) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnTask) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnTask) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnTask) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnTask) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTask) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTask) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnTask) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnTask) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTask) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnTask) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTask) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTask) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

