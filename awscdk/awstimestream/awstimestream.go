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

// Properties for defining a `AWS::Timestream::Database`.
//
// TODO: EXAMPLE
//
type CfnDatabaseProps struct {
	// `AWS::Timestream::Database.DatabaseName`.
	DatabaseName *string `json:"databaseName"`
	// `AWS::Timestream::Database.KmsKeyId`.
	KmsKeyId *string `json:"kmsKeyId"`
	// `AWS::Timestream::Database.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

// A CloudFormation `AWS::Timestream::ScheduledQuery`.
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

// TODO: EXAMPLE
//
type CfnScheduledQuery_DimensionMappingProperty struct {
	// `CfnScheduledQuery.DimensionMappingProperty.DimensionValueType`.
	DimensionValueType *string `json:"dimensionValueType"`
	// `CfnScheduledQuery.DimensionMappingProperty.Name`.
	Name *string `json:"name"`
}

// TODO: EXAMPLE
//
type CfnScheduledQuery_ErrorReportConfigurationProperty struct {
	// `CfnScheduledQuery.ErrorReportConfigurationProperty.S3Configuration`.
	S3Configuration interface{} `json:"s3Configuration"`
}

// TODO: EXAMPLE
//
type CfnScheduledQuery_MixedMeasureMappingProperty struct {
	// `CfnScheduledQuery.MixedMeasureMappingProperty.MeasureName`.
	MeasureName *string `json:"measureName"`
	// `CfnScheduledQuery.MixedMeasureMappingProperty.MeasureValueType`.
	MeasureValueType *string `json:"measureValueType"`
	// `CfnScheduledQuery.MixedMeasureMappingProperty.MultiMeasureAttributeMappings`.
	MultiMeasureAttributeMappings interface{} `json:"multiMeasureAttributeMappings"`
	// `CfnScheduledQuery.MixedMeasureMappingProperty.SourceColumn`.
	SourceColumn *string `json:"sourceColumn"`
	// `CfnScheduledQuery.MixedMeasureMappingProperty.TargetMeasureName`.
	TargetMeasureName *string `json:"targetMeasureName"`
}

// TODO: EXAMPLE
//
type CfnScheduledQuery_MultiMeasureAttributeMappingProperty struct {
	// `CfnScheduledQuery.MultiMeasureAttributeMappingProperty.MeasureValueType`.
	MeasureValueType *string `json:"measureValueType"`
	// `CfnScheduledQuery.MultiMeasureAttributeMappingProperty.SourceColumn`.
	SourceColumn *string `json:"sourceColumn"`
	// `CfnScheduledQuery.MultiMeasureAttributeMappingProperty.TargetMultiMeasureAttributeName`.
	TargetMultiMeasureAttributeName *string `json:"targetMultiMeasureAttributeName"`
}

// TODO: EXAMPLE
//
type CfnScheduledQuery_MultiMeasureMappingsProperty struct {
	// `CfnScheduledQuery.MultiMeasureMappingsProperty.MultiMeasureAttributeMappings`.
	MultiMeasureAttributeMappings interface{} `json:"multiMeasureAttributeMappings"`
	// `CfnScheduledQuery.MultiMeasureMappingsProperty.TargetMultiMeasureName`.
	TargetMultiMeasureName *string `json:"targetMultiMeasureName"`
}

// TODO: EXAMPLE
//
type CfnScheduledQuery_NotificationConfigurationProperty struct {
	// `CfnScheduledQuery.NotificationConfigurationProperty.SnsConfiguration`.
	SnsConfiguration interface{} `json:"snsConfiguration"`
}

// TODO: EXAMPLE
//
type CfnScheduledQuery_S3ConfigurationProperty struct {
	// `CfnScheduledQuery.S3ConfigurationProperty.BucketName`.
	BucketName *string `json:"bucketName"`
	// `CfnScheduledQuery.S3ConfigurationProperty.EncryptionOption`.
	EncryptionOption *string `json:"encryptionOption"`
	// `CfnScheduledQuery.S3ConfigurationProperty.ObjectKeyPrefix`.
	ObjectKeyPrefix *string `json:"objectKeyPrefix"`
}

// TODO: EXAMPLE
//
type CfnScheduledQuery_ScheduleConfigurationProperty struct {
	// `CfnScheduledQuery.ScheduleConfigurationProperty.ScheduleExpression`.
	ScheduleExpression *string `json:"scheduleExpression"`
}

// TODO: EXAMPLE
//
type CfnScheduledQuery_SnsConfigurationProperty struct {
	// `CfnScheduledQuery.SnsConfigurationProperty.TopicArn`.
	TopicArn *string `json:"topicArn"`
}

// TODO: EXAMPLE
//
type CfnScheduledQuery_TargetConfigurationProperty struct {
	// `CfnScheduledQuery.TargetConfigurationProperty.TimestreamConfiguration`.
	TimestreamConfiguration interface{} `json:"timestreamConfiguration"`
}

// TODO: EXAMPLE
//
type CfnScheduledQuery_TimestreamConfigurationProperty struct {
	// `CfnScheduledQuery.TimestreamConfigurationProperty.DatabaseName`.
	DatabaseName *string `json:"databaseName"`
	// `CfnScheduledQuery.TimestreamConfigurationProperty.DimensionMappings`.
	DimensionMappings interface{} `json:"dimensionMappings"`
	// `CfnScheduledQuery.TimestreamConfigurationProperty.MeasureNameColumn`.
	MeasureNameColumn *string `json:"measureNameColumn"`
	// `CfnScheduledQuery.TimestreamConfigurationProperty.MixedMeasureMappings`.
	MixedMeasureMappings interface{} `json:"mixedMeasureMappings"`
	// `CfnScheduledQuery.TimestreamConfigurationProperty.MultiMeasureMappings`.
	MultiMeasureMappings interface{} `json:"multiMeasureMappings"`
	// `CfnScheduledQuery.TimestreamConfigurationProperty.TableName`.
	TableName *string `json:"tableName"`
	// `CfnScheduledQuery.TimestreamConfigurationProperty.TimeColumn`.
	TimeColumn *string `json:"timeColumn"`
}

// Properties for defining a `AWS::Timestream::ScheduledQuery`.
//
// TODO: EXAMPLE
//
type CfnScheduledQueryProps struct {
	// `AWS::Timestream::ScheduledQuery.ClientToken`.
	ClientToken *string `json:"clientToken"`
	// `AWS::Timestream::ScheduledQuery.ErrorReportConfiguration`.
	ErrorReportConfiguration interface{} `json:"errorReportConfiguration"`
	// `AWS::Timestream::ScheduledQuery.KmsKeyId`.
	KmsKeyId *string `json:"kmsKeyId"`
	// `AWS::Timestream::ScheduledQuery.NotificationConfiguration`.
	NotificationConfiguration interface{} `json:"notificationConfiguration"`
	// `AWS::Timestream::ScheduledQuery.QueryString`.
	QueryString *string `json:"queryString"`
	// `AWS::Timestream::ScheduledQuery.ScheduleConfiguration`.
	ScheduleConfiguration interface{} `json:"scheduleConfiguration"`
	// `AWS::Timestream::ScheduledQuery.ScheduledQueryExecutionRoleArn`.
	ScheduledQueryExecutionRoleArn *string `json:"scheduledQueryExecutionRoleArn"`
	// `AWS::Timestream::ScheduledQuery.ScheduledQueryName`.
	ScheduledQueryName *string `json:"scheduledQueryName"`
	// `AWS::Timestream::ScheduledQuery.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// `AWS::Timestream::ScheduledQuery.TargetConfiguration`.
	TargetConfiguration interface{} `json:"targetConfiguration"`
}

// A CloudFormation `AWS::Timestream::Table`.
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

// Properties for defining a `AWS::Timestream::Table`.
//
// TODO: EXAMPLE
//
type CfnTableProps struct {
	// `AWS::Timestream::Table.DatabaseName`.
	DatabaseName *string `json:"databaseName"`
	// `AWS::Timestream::Table.RetentionProperties`.
	RetentionProperties interface{} `json:"retentionProperties"`
	// `AWS::Timestream::Table.TableName`.
	TableName *string `json:"tableName"`
	// `AWS::Timestream::Table.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
}

