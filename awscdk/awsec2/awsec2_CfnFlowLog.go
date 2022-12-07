package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::EC2::FlowLog`.
//
// Specifies a VPC flow log that captures IP traffic for a specified network interface, subnet, or VPC. To view the log data, use Amazon CloudWatch Logs (CloudWatch Logs) to help troubleshoot connection issues. For example, you can use a flow log to investigate why certain traffic isn't reaching an instance, which can help you diagnose overly restrictive security group rules. For more information, see [VPC Flow Logs](https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs.html) in the *Amazon VPC User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var destinationOptions interface{}
//
//   cfnFlowLog := awscdk.Aws_ec2.NewCfnFlowLog(this, jsii.String("MyCfnFlowLog"), &cfnFlowLogProps{
//   	resourceId: jsii.String("resourceId"),
//   	resourceType: jsii.String("resourceType"),
//
//   	// the properties below are optional
//   	deliverLogsPermissionArn: jsii.String("deliverLogsPermissionArn"),
//   	destinationOptions: destinationOptions,
//   	logDestination: jsii.String("logDestination"),
//   	logDestinationType: jsii.String("logDestinationType"),
//   	logFormat: jsii.String("logFormat"),
//   	logGroupName: jsii.String("logGroupName"),
//   	maxAggregationInterval: jsii.Number(123),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	trafficType: jsii.String("trafficType"),
//   })
//
type CfnFlowLog interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The ID of the flow log.
	//
	// For example, `fl-123456abc123abc1` .
	AttrId() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// The ARN for the IAM role that permits Amazon EC2 to publish flow logs to a CloudWatch Logs log group in your account.
	//
	// If you specify `LogDestinationType` as `s3` , do not specify `DeliverLogsPermissionArn` or `LogGroupName` .
	DeliverLogsPermissionArn() *string
	SetDeliverLogsPermissionArn(val *string)
	// The destination options. The following options are supported:.
	//
	// - `FileFormat` - The format for the flow log ( `plain-text` | `parquet` ). The default is `plain-text` .
	// - `HiveCompatiblePartitions` - Indicates whether to use Hive-compatible prefixes for flow logs stored in Amazon S3 ( `true` | `false` ). The default is `false` .
	// - `PerHourPartition` - Indicates whether to partition the flow log per hour ( `true` | `false` ). The default is `false` .
	DestinationOptions() interface{}
	SetDestinationOptions(val interface{})
	// The destination to which the flow log data is to be published.
	//
	// Flow log data can be published to a CloudWatch Logs log group or an Amazon S3 bucket. The value specified for this parameter depends on the value specified for `LogDestinationType` .
	//
	// If `LogDestinationType` is not specified or `cloud-watch-logs` , specify the Amazon Resource Name (ARN) of the CloudWatch Logs log group. For example, to publish to a log group called `my-logs` , specify `arn:aws:logs:us-east-1:123456789012:log-group:my-logs` . Alternatively, use `LogGroupName` instead.
	//
	// If LogDestinationType is `s3` , specify the ARN of the Amazon S3 bucket. You can also specify a subfolder in the bucket. To specify a subfolder in the bucket, use the following ARN format: `bucket_ARN/subfolder_name/` . For example, to specify a subfolder named `my-logs` in a bucket named `my-bucket` , use the following ARN: `arn:aws:s3:::my-bucket/my-logs/` . You cannot use `AWSLogs` as a subfolder name. This is a reserved term.
	LogDestination() *string
	SetLogDestination(val *string)
	// The type of destination to which the flow log data is to be published.
	//
	// Flow log data can be published to CloudWatch Logs or Amazon S3. To publish flow log data to CloudWatch Logs, specify `cloud-watch-logs` . To publish flow log data to Amazon S3, specify `s3` .
	//
	// If you specify `LogDestinationType` as `s3` , do not specify `DeliverLogsPermissionArn` or `LogGroupName` .
	//
	// Default: `cloud-watch-logs`.
	LogDestinationType() *string
	SetLogDestinationType(val *string)
	// The fields to include in the flow log record, in the order in which they should appear.
	//
	// For a list of available fields, see [Flow Log Records](https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs.html#flow-log-records) . If you omit this parameter, the flow log is created using the default format. If you specify this parameter, you must specify at least one field.
	//
	// Specify the fields using the `${field-id}` format, separated by spaces.
	LogFormat() *string
	SetLogFormat(val *string)
	// The name of a new or existing CloudWatch Logs log group where Amazon EC2 publishes your flow logs.
	//
	// If you specify `LogDestinationType` as `s3` , do not specify `DeliverLogsPermissionArn` or `LogGroupName` .
	LogGroupName() *string
	SetLogGroupName(val *string)
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
	// The maximum interval of time during which a flow of packets is captured and aggregated into a flow log record.
	//
	// You can specify 60 seconds (1 minute) or 600 seconds (10 minutes).
	//
	// When a network interface is attached to a [Nitro-based instance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#ec2-nitro-instances) , the aggregation interval is always 60 seconds or less, regardless of the value that you specify.
	//
	// Default: 600.
	MaxAggregationInterval() *float64
	SetMaxAggregationInterval(val *float64)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The ID of the subnet, network interface, or VPC for which you want to create a flow log.
	ResourceId() *string
	SetResourceId(val *string)
	// The type of resource for which to create the flow log.
	//
	// For example, if you specified a VPC ID for the `ResourceId` property, specify `VPC` for this property.
	ResourceType() *string
	SetResourceType(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The tags to apply to the flow logs.
	Tags() awscdk.TagManager
	// The type of traffic to log.
	//
	// You can log traffic that the resource accepts or rejects, or all traffic.
	TrafficType() *string
	SetTrafficType(val *string)
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
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnFlowLog
type jsiiProxy_CfnFlowLog struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnFlowLog) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowLog) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowLog) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowLog) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowLog) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowLog) DeliverLogsPermissionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliverLogsPermissionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowLog) DestinationOptions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowLog) LogDestination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowLog) LogDestinationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logDestinationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowLog) LogFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowLog) LogGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowLog) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowLog) MaxAggregationInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAggregationInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowLog) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowLog) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowLog) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowLog) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowLog) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowLog) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowLog) TrafficType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trafficType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowLog) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowLog) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::EC2::FlowLog`.
func NewCfnFlowLog(scope constructs.Construct, id *string, props *CfnFlowLogProps) CfnFlowLog {
	_init_.Initialize()

	if err := validateNewCfnFlowLogParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFlowLog{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.CfnFlowLog",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::EC2::FlowLog`.
func NewCfnFlowLog_Override(c CfnFlowLog, scope constructs.Construct, id *string, props *CfnFlowLogProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.CfnFlowLog",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnFlowLog)SetDeliverLogsPermissionArn(val *string) {
	_jsii_.Set(
		j,
		"deliverLogsPermissionArn",
		val,
	)
}

func (j *jsiiProxy_CfnFlowLog)SetDestinationOptions(val interface{}) {
	if err := j.validateSetDestinationOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationOptions",
		val,
	)
}

func (j *jsiiProxy_CfnFlowLog)SetLogDestination(val *string) {
	_jsii_.Set(
		j,
		"logDestination",
		val,
	)
}

func (j *jsiiProxy_CfnFlowLog)SetLogDestinationType(val *string) {
	_jsii_.Set(
		j,
		"logDestinationType",
		val,
	)
}

func (j *jsiiProxy_CfnFlowLog)SetLogFormat(val *string) {
	_jsii_.Set(
		j,
		"logFormat",
		val,
	)
}

func (j *jsiiProxy_CfnFlowLog)SetLogGroupName(val *string) {
	_jsii_.Set(
		j,
		"logGroupName",
		val,
	)
}

func (j *jsiiProxy_CfnFlowLog)SetMaxAggregationInterval(val *float64) {
	_jsii_.Set(
		j,
		"maxAggregationInterval",
		val,
	)
}

func (j *jsiiProxy_CfnFlowLog)SetResourceId(val *string) {
	if err := j.validateSetResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_CfnFlowLog)SetResourceType(val *string) {
	if err := j.validateSetResourceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceType",
		val,
	)
}

func (j *jsiiProxy_CfnFlowLog)SetTrafficType(val *string) {
	_jsii_.Set(
		j,
		"trafficType",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnFlowLog_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFlowLog_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CfnFlowLog",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnFlowLog_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnFlowLog_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CfnFlowLog",
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
func CfnFlowLog_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFlowLog_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.CfnFlowLog",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFlowLog_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ec2.CfnFlowLog",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFlowLog) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnFlowLog) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnFlowLog) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnFlowLog) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnFlowLog) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnFlowLog) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnFlowLog) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnFlowLog) GetAtt(attributeName *string) awscdk.Reference {
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

func (c *jsiiProxy_CfnFlowLog) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnFlowLog) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnFlowLog) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnFlowLog) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnFlowLog) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFlowLog) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnFlowLog) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

