package awstimestream

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awstimestream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::Timestream::Database`.
//
// Creates a new Timestream database. If the AWS KMS key is not specified, the database will be encrypted with a Timestream managed AWS KMS key located in your account. Refer to [AWS managed AWS KMS keys](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk) for more info. [Service quotas apply](https://docs.aws.amazon.com/timestream/latest/developerguide/ts-limits.html) . See [code sample](https://docs.aws.amazon.com/timestream/latest/developerguide/code-samples.create-db.html) for details.
//
// TODO: EXAMPLE
//
type CfnDatabase interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DatabaseName() *string
	SetDatabaseName(val *string)
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	LogicalId() *string
	Node() constructs.Node
	Ref() *string
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
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

// The jsii proxy struct for CfnDatabase
type jsiiProxy_CfnDatabase struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDatabase) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDatabase) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Timestream::Database`.
func NewCfnDatabase(scope constructs.Construct, id *string, props *CfnDatabaseProps) CfnDatabase {
	_init_.Initialize()

	j := jsiiProxy_CfnDatabase{}

	_jsii_.Create(
		"aws-cdk-lib.aws_timestream.CfnDatabase",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Timestream::Database`.
func NewCfnDatabase_Override(c CfnDatabase, scope constructs.Construct, id *string, props *CfnDatabaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_timestream.CfnDatabase",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDatabase) SetDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_CfnDatabase) SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDatabase_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_timestream.CfnDatabase",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDatabase_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_timestream.CfnDatabase",
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
func CfnDatabase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_timestream.CfnDatabase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDatabase_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_timestream.CfnDatabase",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnDatabase) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnDatabase) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnDatabase) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnDatabase) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnDatabase) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnDatabase) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnDatabase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnDatabase) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnDatabase) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnDatabase) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnDatabase) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDatabase) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnDatabase) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnDatabase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDatabase) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnDatabase`.
//
// TODO: EXAMPLE
//
type CfnDatabaseProps struct {
	// The name of the Timestream database.
	//
	// *Length Constraints* : Minimum length of 3 bytes. Maximum length of 256 bytes.
	DatabaseName *string `json:"databaseName" yaml:"databaseName"`
	// The identifier of the AWS KMS key used to encrypt the data stored in the database.
	KmsKeyId *string `json:"kmsKeyId" yaml:"kmsKeyId"`
	// The tags to add to the database.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// A CloudFormation `AWS::Timestream::ScheduledQuery`.
//
// Create a scheduled query that will be run on your behalf at the configured schedule. Timestream assumes the execution role provided as part of the `ScheduledQueryExecutionRoleArn` parameter to run the query. You can use the `NotificationConfiguration` parameter to configure notification for your scheduled query operations.
//
// TODO: EXAMPLE
//
type CfnScheduledQuery interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrSqErrorReportConfiguration() *string
	AttrSqKmsKeyId() *string
	AttrSqName() *string
	AttrSqNotificationConfiguration() *string
	AttrSqQueryString() *string
	AttrSqScheduleConfiguration() *string
	AttrSqScheduledQueryExecutionRoleArn() *string
	AttrSqTargetConfiguration() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ClientToken() *string
	SetClientToken(val *string)
	CreationStack() *[]*string
	ErrorReportConfiguration() interface{}
	SetErrorReportConfiguration(val interface{})
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	LogicalId() *string
	Node() constructs.Node
	NotificationConfiguration() interface{}
	SetNotificationConfiguration(val interface{})
	QueryString() *string
	SetQueryString(val *string)
	Ref() *string
	ScheduleConfiguration() interface{}
	SetScheduleConfiguration(val interface{})
	ScheduledQueryExecutionRoleArn() *string
	SetScheduledQueryExecutionRoleArn(val *string)
	ScheduledQueryName() *string
	SetScheduledQueryName(val *string)
	Stack() awscdk.Stack
	Tags() awscdk.TagManager
	TargetConfiguration() interface{}
	SetTargetConfiguration(val interface{})
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

// The jsii proxy struct for CfnScheduledQuery
type jsiiProxy_CfnScheduledQuery struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnScheduledQuery) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) AttrSqErrorReportConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSqErrorReportConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) AttrSqKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSqKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) AttrSqName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSqName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) AttrSqNotificationConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSqNotificationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) AttrSqQueryString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSqQueryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) AttrSqScheduleConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSqScheduleConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) AttrSqScheduledQueryExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSqScheduledQueryExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) AttrSqTargetConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrSqTargetConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) ClientToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) ErrorReportConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorReportConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) NotificationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) QueryString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) ScheduleConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduleConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) ScheduledQueryExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduledQueryExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) ScheduledQueryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduledQueryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) TargetConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledQuery) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Timestream::ScheduledQuery`.
func NewCfnScheduledQuery(scope constructs.Construct, id *string, props *CfnScheduledQueryProps) CfnScheduledQuery {
	_init_.Initialize()

	j := jsiiProxy_CfnScheduledQuery{}

	_jsii_.Create(
		"aws-cdk-lib.aws_timestream.CfnScheduledQuery",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Timestream::ScheduledQuery`.
func NewCfnScheduledQuery_Override(c CfnScheduledQuery, scope constructs.Construct, id *string, props *CfnScheduledQueryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_timestream.CfnScheduledQuery",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnScheduledQuery) SetClientToken(val *string) {
	_jsii_.Set(
		j,
		"clientToken",
		val,
	)
}

func (j *jsiiProxy_CfnScheduledQuery) SetErrorReportConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"errorReportConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnScheduledQuery) SetKmsKeyId(val *string) {
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_CfnScheduledQuery) SetNotificationConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"notificationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnScheduledQuery) SetQueryString(val *string) {
	_jsii_.Set(
		j,
		"queryString",
		val,
	)
}

func (j *jsiiProxy_CfnScheduledQuery) SetScheduleConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"scheduleConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnScheduledQuery) SetScheduledQueryExecutionRoleArn(val *string) {
	_jsii_.Set(
		j,
		"scheduledQueryExecutionRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnScheduledQuery) SetScheduledQueryName(val *string) {
	_jsii_.Set(
		j,
		"scheduledQueryName",
		val,
	)
}

func (j *jsiiProxy_CfnScheduledQuery) SetTargetConfiguration(val interface{}) {
	_jsii_.Set(
		j,
		"targetConfiguration",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnScheduledQuery_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_timestream.CfnScheduledQuery",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnScheduledQuery_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_timestream.CfnScheduledQuery",
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
func CfnScheduledQuery_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_timestream.CfnScheduledQuery",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnScheduledQuery_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_timestream.CfnScheduledQuery",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnScheduledQuery) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnScheduledQuery) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnScheduledQuery) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnScheduledQuery) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnScheduledQuery) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnScheduledQuery) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnScheduledQuery) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnScheduledQuery) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnScheduledQuery) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnScheduledQuery) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnScheduledQuery) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnScheduledQuery) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnScheduledQuery) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnScheduledQuery) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnScheduledQuery) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// This type is used to map column(s) from the query result to a dimension in the destination table.
//
// TODO: EXAMPLE
//
type CfnScheduledQuery_DimensionMappingProperty struct {
	// Type for the dimension.
	DimensionValueType *string `json:"dimensionValueType" yaml:"dimensionValueType"`
	// Column name from query result.
	Name *string `json:"name" yaml:"name"`
}

// Configuration required for error reporting.
//
// TODO: EXAMPLE
//
type CfnScheduledQuery_ErrorReportConfigurationProperty struct {
	// The S3 configuration for the error reports.
	S3Configuration interface{} `json:"s3Configuration" yaml:"s3Configuration"`
}

// MixedMeasureMappings are mappings that can be used to ingest data into a mixture of narrow and multi measures in the derived table.
//
// TODO: EXAMPLE
//
type CfnScheduledQuery_MixedMeasureMappingProperty struct {
	// Type of the value that is to be read from sourceColumn.
	//
	// If the mapping is for MULTI, use MeasureValueType.MULTI.
	MeasureValueType *string `json:"measureValueType" yaml:"measureValueType"`
	// Refers to the value of measure_name in a result row.
	//
	// This field is required if MeasureNameColumn is provided.
	MeasureName *string `json:"measureName" yaml:"measureName"`
	// Required when measureValueType is MULTI.
	//
	// Attribute mappings for MULTI value measures.
	MultiMeasureAttributeMappings interface{} `json:"multiMeasureAttributeMappings" yaml:"multiMeasureAttributeMappings"`
	// This field refers to the source column from which measure-value is to be read for result materialization.
	SourceColumn *string `json:"sourceColumn" yaml:"sourceColumn"`
	// Target measure name to be used.
	//
	// If not provided, the target measure name by default would be measure-name if provided, or sourceColumn otherwise.
	TargetMeasureName *string `json:"targetMeasureName" yaml:"targetMeasureName"`
}

// Attribute mapping for MULTI value measures.
//
// TODO: EXAMPLE
//
type CfnScheduledQuery_MultiMeasureAttributeMappingProperty struct {
	// Type of the attribute to be read from the source column.
	MeasureValueType *string `json:"measureValueType" yaml:"measureValueType"`
	// Source column from where the attribute value is to be read.
	SourceColumn *string `json:"sourceColumn" yaml:"sourceColumn"`
	// Custom name to be used for attribute name in derived table.
	//
	// If not provided, source column name would be used.
	TargetMultiMeasureAttributeName *string `json:"targetMultiMeasureAttributeName" yaml:"targetMultiMeasureAttributeName"`
}

// Only one of MixedMeasureMappings or MultiMeasureMappings is to be provided.
//
// MultiMeasureMappings can be used to ingest data as multi measures in the derived table.
//
// TODO: EXAMPLE
//
type CfnScheduledQuery_MultiMeasureMappingsProperty struct {
	// Required.
	//
	// Attribute mappings to be used for mapping query results to ingest data for multi-measure attributes.
	MultiMeasureAttributeMappings interface{} `json:"multiMeasureAttributeMappings" yaml:"multiMeasureAttributeMappings"`
	// The name of the target multi-measure name in the derived table.
	//
	// This input is required when measureNameColumn is not provided. If MeasureNameColumn is provided, then value from that column will be used as multi-measure name.
	TargetMultiMeasureName *string `json:"targetMultiMeasureName" yaml:"targetMultiMeasureName"`
}

// Notification configuration for a scheduled query.
//
// A notification is sent by Timestream when a scheduled query is created, its state is updated or when it is deleted.
//
// TODO: EXAMPLE
//
type CfnScheduledQuery_NotificationConfigurationProperty struct {
	// Details on SNS configuration.
	SnsConfiguration interface{} `json:"snsConfiguration" yaml:"snsConfiguration"`
}

// Details on S3 location for error reports that result from running a query.
//
// TODO: EXAMPLE
//
type CfnScheduledQuery_S3ConfigurationProperty struct {
	// Name of the S3 bucket under which error reports will be created.
	BucketName *string `json:"bucketName" yaml:"bucketName"`
	// Encryption at rest options for the error reports.
	//
	// If no encryption option is specified, Timestream will choose SSE_S3 as default.
	EncryptionOption *string `json:"encryptionOption" yaml:"encryptionOption"`
	// Prefix for the error report key.
	//
	// Timestream by default adds the following prefix to the error report path.
	ObjectKeyPrefix *string `json:"objectKeyPrefix" yaml:"objectKeyPrefix"`
}

// Configuration of the schedule of the query.
//
// TODO: EXAMPLE
//
type CfnScheduledQuery_ScheduleConfigurationProperty struct {
	// An expression that denotes when to trigger the scheduled query run.
	//
	// This can be a cron expression or a rate expression.
	ScheduleExpression *string `json:"scheduleExpression" yaml:"scheduleExpression"`
}

// Details on SNS that are required to send the notification.
//
// TODO: EXAMPLE
//
type CfnScheduledQuery_SnsConfigurationProperty struct {
	// SNS topic ARN that the scheduled query status notifications will be sent to.
	TopicArn *string `json:"topicArn" yaml:"topicArn"`
}

// Configuration used for writing the output of a query.
//
// TODO: EXAMPLE
//
type CfnScheduledQuery_TargetConfigurationProperty struct {
	// Configuration needed to write data into the Timestream database and table.
	TimestreamConfiguration interface{} `json:"timestreamConfiguration" yaml:"timestreamConfiguration"`
}

// Configuration to write data into Timestream database and table.
//
// This configuration allows the user to map the query result select columns into the destination table columns.
//
// TODO: EXAMPLE
//
type CfnScheduledQuery_TimestreamConfigurationProperty struct {
	// Name of Timestream database to which the query result will be written.
	DatabaseName *string `json:"databaseName" yaml:"databaseName"`
	// This is to allow mapping column(s) from the query result to the dimension in the destination table.
	DimensionMappings interface{} `json:"dimensionMappings" yaml:"dimensionMappings"`
	// Name of Timestream table that the query result will be written to.
	//
	// The table should be within the same database that is provided in Timestream configuration.
	TableName *string `json:"tableName" yaml:"tableName"`
	// Column from query result that should be used as the time column in destination table.
	//
	// Column type for this should be TIMESTAMP.
	TimeColumn *string `json:"timeColumn" yaml:"timeColumn"`
	// Name of the measure column.
	MeasureNameColumn *string `json:"measureNameColumn" yaml:"measureNameColumn"`
	// Specifies how to map measures to multi-measure records.
	MixedMeasureMappings interface{} `json:"mixedMeasureMappings" yaml:"mixedMeasureMappings"`
	// Multi-measure mappings.
	MultiMeasureMappings interface{} `json:"multiMeasureMappings" yaml:"multiMeasureMappings"`
}

// Properties for defining a `CfnScheduledQuery`.
//
// TODO: EXAMPLE
//
type CfnScheduledQueryProps struct {
	// Configuration for error reporting.
	//
	// Error reports will be generated when a problem is encountered when writing the query results.
	ErrorReportConfiguration interface{} `json:"errorReportConfiguration" yaml:"errorReportConfiguration"`
	// Notification configuration for the scheduled query.
	//
	// A notification is sent by Timestream when a query run finishes, when the state is updated or when you delete it.
	NotificationConfiguration interface{} `json:"notificationConfiguration" yaml:"notificationConfiguration"`
	// The query string to run.
	//
	// Parameter names can be specified in the query string `@` character followed by an identifier. The named Parameter `@scheduled_runtime` is reserved and can be used in the query to get the time at which the query is scheduled to run.
	//
	// The timestamp calculated according to the ScheduleConfiguration parameter, will be the value of `@scheduled_runtime` paramater for each query run. For example, consider an instance of a scheduled query executing on 2021-12-01 00:00:00. For this instance, the `@scheduled_runtime` parameter is initialized to the timestamp 2021-12-01 00:00:00 when invoking the query.
	QueryString *string `json:"queryString" yaml:"queryString"`
	// Schedule configuration.
	ScheduleConfiguration interface{} `json:"scheduleConfiguration" yaml:"scheduleConfiguration"`
	// The ARN for the IAM role that Timestream will assume when running the scheduled query.
	ScheduledQueryExecutionRoleArn *string `json:"scheduledQueryExecutionRoleArn" yaml:"scheduledQueryExecutionRoleArn"`
	// Using a ClientToken makes the call to CreateScheduledQuery idempotent, in other words, making the same request repeatedly will produce the same result.
	//
	// Making multiple identical CreateScheduledQuery requests has the same effect as making a single request.
	//
	// - If CreateScheduledQuery is called without a `ClientToken` , the Query SDK generates a `ClientToken` on your behalf.
	// - After 8 hours, any request with the same `ClientToken` is treated as a new request.
	ClientToken *string `json:"clientToken" yaml:"clientToken"`
	// The Amazon KMS key used to encrypt the scheduled query resource, at-rest.
	//
	// If the Amazon KMS key is not specified, the scheduled query resource will be encrypted with a Timestream owned Amazon KMS key. To specify a KMS key, use the key ID, key ARN, alias name, or alias ARN. When using an alias name, prefix the name with *alias/*
	//
	// If ErrorReportConfiguration uses `SSE_KMS` as encryption type, the same KmsKeyId is used to encrypt the error report at rest.
	KmsKeyId *string `json:"kmsKeyId" yaml:"kmsKeyId"`
	// A name for the query.
	//
	// Scheduled query names must be unique within each Region.
	ScheduledQueryName *string `json:"scheduledQueryName" yaml:"scheduledQueryName"`
	// A list of key-value pairs to label the scheduled query.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// Scheduled query target store configuration.
	TargetConfiguration interface{} `json:"targetConfiguration" yaml:"targetConfiguration"`
}

// A CloudFormation `AWS::Timestream::Table`.
//
// The CreateTable operation adds a new table to an existing database in your account. In an AWS account, table names must be at least unique within each Region if they are in the same database. You may have identical table names in the same Region if the tables are in separate databases. While creating the table, you must specify the table name, database name, and the retention properties. [Service quotas apply](https://docs.aws.amazon.com/timestream/latest/developerguide/ts-limits.html) . See [code sample](https://docs.aws.amazon.com/timestream/latest/developerguide/code-samples.create-table.html) for details.
//
// TODO: EXAMPLE
//
type CfnTable interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttrName() *string
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	DatabaseName() *string
	SetDatabaseName(val *string)
	LogicalId() *string
	MagneticStoreWriteProperties() interface{}
	SetMagneticStoreWriteProperties(val interface{})
	Node() constructs.Node
	Ref() *string
	RetentionProperties() interface{}
	SetRetentionProperties(val interface{})
	Stack() awscdk.Stack
	TableName() *string
	SetTableName(val *string)
	Tags() awscdk.TagManager
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

// The jsii proxy struct for CfnTable
type jsiiProxy_CfnTable struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnTable) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTable) AttrName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTable) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTable) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTable) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTable) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTable) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTable) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTable) MagneticStoreWriteProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"magneticStoreWriteProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTable) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTable) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTable) RetentionProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retentionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTable) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTable) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTable) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTable) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}


// Create a new `AWS::Timestream::Table`.
func NewCfnTable(scope constructs.Construct, id *string, props *CfnTableProps) CfnTable {
	_init_.Initialize()

	j := jsiiProxy_CfnTable{}

	_jsii_.Create(
		"aws-cdk-lib.aws_timestream.CfnTable",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::Timestream::Table`.
func NewCfnTable_Override(c CfnTable, scope constructs.Construct, id *string, props *CfnTableProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_timestream.CfnTable",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTable) SetDatabaseName(val *string) {
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_CfnTable) SetMagneticStoreWriteProperties(val interface{}) {
	_jsii_.Set(
		j,
		"magneticStoreWriteProperties",
		val,
	)
}

func (j *jsiiProxy_CfnTable) SetRetentionProperties(val interface{}) {
	_jsii_.Set(
		j,
		"retentionProperties",
		val,
	)
}

func (j *jsiiProxy_CfnTable) SetTableName(val *string) {
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnTable_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_timestream.CfnTable",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnTable_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_timestream.CfnTable",
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
func CfnTable_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_timestream.CfnTable",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTable_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_timestream.CfnTable",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
func (c *jsiiProxy_CfnTable) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnTable) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnTable) AddMetadata(key *string, value interface{}) {
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
func (c *jsiiProxy_CfnTable) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
func (c *jsiiProxy_CfnTable) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

// Adds an override to a resource property.
//
// Syntactic sugar for `addOverride("Properties.<...>", value)`.
func (c *jsiiProxy_CfnTable) AddPropertyOverride(propertyPath *string, value interface{}) {
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
func (c *jsiiProxy_CfnTable) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnTable) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnTable) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnTable) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Overrides the auto-generated logical ID with a specific ID.
func (c *jsiiProxy_CfnTable) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTable) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnTable) ShouldSynthesize() *bool {
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
func (c *jsiiProxy_CfnTable) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnTable) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Properties for defining a `CfnTable`.
//
// TODO: EXAMPLE
//
type CfnTableProps struct {
	// The name of the Timestream database that contains this table.
	//
	// *Length Constraints* : Minimum length of 3 bytes. Maximum length of 256 bytes.
	DatabaseName *string `json:"databaseName" yaml:"databaseName"`
	// Contains properties to set on the table when enabling magnetic store writes.
	//
	// This object has the following attributes:
	//
	// - *EnableMagneticStoreWrites* : A `boolean` flag to enable magnetic store writes.
	// - *MagneticStoreRejectedDataLocation* : The location to write error reports for records rejected, asynchronously, during magnetic store writes. Only `S3Configuration` objects are allowed. The `S3Configuration` object has the following attributes:
	//
	// - *BucketName* : The bucket name of the S3 bucket.
	// - *EncryptionOption* : The encryption option for the S3 location. Valid values are S3 server-side encryption with an S3 managed key ( `SSE_S3` ) or AWS managed key ( `SSE_KMS` ).
	// - *KmsKeyId* : The AWS KMS key ID to use when encrypting with an AWS managed key.
	// - *ObjectKeyPrefix* :The prefix to use option for the objects stored in S3.
	//
	// Both `BucketName` and `EncryptionOption` are *required* when `S3Configuration` is specified. If you specify `SSE_KMS` as your `EncryptionOption` then `KmsKeyId` is *required* .
	//
	// `EnableMagneticStoreWrites` attribute is *required* when `MagneticStoreWriteProperties` is specified. `MagneticStoreRejectedDataLocation` attribute is *required* when `EnableMagneticStoreWrites` is set to `true` .
	//
	// See the following examples:
	//
	// *JSON*
	//
	// ```json
	// { "Type" : AWS::Timestream::Table", "Properties":{ "DatabaseName":"TestDatabase", "TableName":"TestTable", "MagneticStoreWriteProperties":{ "EnableMagneticStoreWrites":true, "MagneticStoreRejectedDataLocation":{ "S3Configuration":{ "BucketName":"testbucket", "EncryptionOption":"SSE_KMS", "KmsKeyId":"1234abcd-12ab-34cd-56ef-1234567890ab", "ObjectKeyPrefix":"prefix" } } } }
	// }
	// ```
	//
	// *YAML*
	//
	// ```
	// Type: AWS::Timestream::Table
	// DependsOn: TestDatabase
	// Properties: TableName: "TestTable" DatabaseName: "TestDatabase" MagneticStoreWriteProperties: EnableMagneticStoreWrites: true MagneticStoreRejectedDataLocation: S3Configuration: BucketName: "testbucket" EncryptionOption: "SSE_KMS" BucketName: "1234abcd-12ab-34cd-56ef-1234567890ab" EncryptionOption: "prefix"
	// ```
	MagneticStoreWriteProperties interface{} `json:"magneticStoreWriteProperties" yaml:"magneticStoreWriteProperties"`
	// The retention duration for the memory store and magnetic store. This object has the following attributes:.
	//
	// - *MemoryStoreRetentionPeriodInHours* : Retention duration for memory store, in hours.
	// - *MagneticStoreRetentionPeriodInDays* : Retention duration for magnetic store, in days.
	//
	// Both attributes are of type `string` . Both attributes are *required* when `RetentionProperties` is specified.
	//
	// See the following examples:
	//
	// *JSON*
	//
	// `{ "Type" : AWS::Timestream::Table", "Properties" : { "DatabaseName" : "TestDatabase", "TableName" : "TestTable", "RetentionProperties" : { "MemoryStoreRetentionPeriodInHours": "24", "MagneticStoreRetentionPeriodInDays": "7" } } }`
	//
	// *YAML*
	//
	// ```
	// Type: AWS::Timestream::Table
	// DependsOn: TestDatabase
	// Properties: TableName: "TestTable" DatabaseName: "TestDatabase" RetentionProperties: MemoryStoreRetentionPeriodInHours: "24" MagneticStoreRetentionPeriodInDays: "7"
	// ```
	RetentionProperties interface{} `json:"retentionProperties" yaml:"retentionProperties"`
	// The name of the Timestream table.
	//
	// *Length Constraints* : Minimum length of 3 bytes. Maximum length of 256 bytes.
	TableName *string `json:"tableName" yaml:"tableName"`
	// The tags to add to the table.
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

