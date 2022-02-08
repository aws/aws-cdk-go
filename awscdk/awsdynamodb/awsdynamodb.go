package awsdynamodb

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsapplicationautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awsdynamodb/internal"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awskinesis"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/constructs-go/constructs/v3"
)

// Represents an attribute for describing the key schema for the table and indexes.
//
// TODO: EXAMPLE
//
// Experimental.
type Attribute struct {
	// The name of an attribute.
	// Experimental.
	Name *string `json:"name" yaml:"name"`
	// The data type of an attribute.
	// Experimental.
	Type AttributeType `json:"type" yaml:"type"`
}

// Data types for attributes within a table.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.NamingRulesDataTypes.html#HowItWorks.DataTypes
//
// Experimental.
type AttributeType string

const (
	AttributeType_BINARY AttributeType = "BINARY"
	AttributeType_NUMBER AttributeType = "NUMBER"
	AttributeType_STRING AttributeType = "STRING"
)

// DynamoDB's Read/Write capacity modes.
//
// TODO: EXAMPLE
//
// Experimental.
type BillingMode string

const (
	BillingMode_PAY_PER_REQUEST BillingMode = "PAY_PER_REQUEST"
	BillingMode_PROVISIONED BillingMode = "PROVISIONED"
)

// A CloudFormation `AWS::DynamoDB::GlobalTable`.
//
// The `AWS::DynamoDB::GlobalTable` resource enables you to create and manage a Version 2019.11.21 global table. This resource cannot be used to create or manage a Version 2017.11.29 global table.
//
// > You cannot convert a resource of type `AWS::DynamoDB::Table` into a resource of type `AWS::DynamoDB::GlobalTable` by changing its type in your template. *Doing so might result in the deletion of your DynamoDB table.*
// >
// > You can instead use the GlobalTable resource to create a new table in a single Region. This will be billed the same as a single Region table. If you later update the stack to add other Regions then Global Tables pricing will apply.
//
// You should be aware of the following behaviors when working with DynamoDB global tables.
//
// - The IAM Principal executing the stack operation must have the permissions listed below in all regions where you plan to have a global table replica. The IAM Principal's permissions should not have restrictions based on IP source address. Some global tables operations (for example, adding a replica) are asynchronous, and require that the IAM Principal is valid until they complete. You should not delete the Principal (user or IAM role) until CloudFormation has finished updating your stack.
//
// - `dynamodb:CreateTable`
// - `dynamodb:UpdateTable`
// - `dynamodb:DeleteTable`
// - `dynamodb:DescribeContinuousBackups`
// - `dynamodb:DescribeContributorInsights`
// - `dynamodb:DescribeTable`
// - `dynamodb:DescribeTableReplicaAutoScaling`
// - `dynamodb:DescribeTimeToLive`
// - `dynamodb:ListTables`
// - `dynamodb:UpdateTimeToLive`
// - `dynamodb:UpdateContributorInsights`
// - `dynamodb:UpdateContinuousBackups`
// - `dynamodb:ListTagsOfResource`
// - `dynamodb:TagResource`
// - `dynamodb:UntagResource`
// - `dynamodb:BatchWriteItem`
// - `dynamodb:CreateTableReplica`
// - `dynamodb:DeleteItem`
// - `dynamodb:DeleteTableReplica`
// - `dynamodb:DisableKinesisStreamingDestination`
// - `dynamodb:EnableKinesisStreamingDestination`
// - `dynamodb:GetItem`
// - `dynamodb:PutItem`
// - `dynamodb:Query`
// - `dynamodb:Scan`
// - `dynamodb:UpdateItem`
// - `dynamodb:DescribeTableReplicaAutoScaling`
// - `dynamodb:UpdateTableReplicaAutoScaling`
// - `iam:CreateServiceLinkedRole`
// - `kms:CreateGrant`
// - `kms:DescribeKey`
// - `application-autoscaling:DeleteScalingPolicy`
// - `application-autoscaling:DeleteScheduledAction`
// - `application-autoscaling:DeregisterScalableTarget`
// - `application-autoscaling:DescribeScalingPolicies`
// - `application-autoscaling:DescribeScalableTargets`
// - `application-autoscaling:PutScalingPolicy`
// - `application-autoscaling:PutScheduledAction`
// - `application-autoscaling:RegisterScalableTarget`
// - When using provisioned billing mode, CloudFormation will create an auto scaling policy on each of your replicas to control their write capacities. You must configure this policy using the `WriteProvisionedThroughputSettings` property. CloudFormation will ensure that all replicas have the same write capacity auto scaling property. You cannot directly specify a value for write capacity for a global table.
// - If your table uses provisioned capacity, you must configure auto scaling directly in the `AWS::DynamoDB::GlobalTable` resource. You should not configure additional auto scaling policies on any of the table replicas or global secondary indexes, either via API or via `AWS::ApplicationAutoScaling::ScalableTarget` or `AWS::ApplicationAutoScaling::ScalingPolicy` . Doing so might result in unexpected behavior and is unsupported.
// - In AWS CloudFormation , each global table is controlled by a single stack, in a single region, regardless of the number of replicas. When you deploy your template, CloudFormation will create/update all replicas as part of a single stack operation. You should not deploy the same `AWS::DynamoDB::GlobalTable` resource in multiple regions. Doing so will result in errors, and is unsupported. If you deploy your application template in multiple regions, you can use conditions to only create the resource in a single region. Alternatively, you can choose to define your `AWS::DynamoDB::GlobalTable` resources in a stack separate from your application stack, and make sure it is only deployed to a single region.
//
// TODO: EXAMPLE
//
type CfnGlobalTable interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttributeDefinitions() interface{}
	SetAttributeDefinitions(val interface{})
	AttrStreamArn() *string
	AttrTableId() *string
	BillingMode() *string
	SetBillingMode(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	CreationStack() *[]*string
	GlobalSecondaryIndexes() interface{}
	SetGlobalSecondaryIndexes(val interface{})
	KeySchema() interface{}
	SetKeySchema(val interface{})
	LocalSecondaryIndexes() interface{}
	SetLocalSecondaryIndexes(val interface{})
	LogicalId() *string
	Node() awscdk.ConstructNode
	Ref() *string
	Replicas() interface{}
	SetReplicas(val interface{})
	SseSpecification() interface{}
	SetSseSpecification(val interface{})
	Stack() awscdk.Stack
	StreamSpecification() interface{}
	SetStreamSpecification(val interface{})
	TableName() *string
	SetTableName(val *string)
	TimeToLiveSpecification() interface{}
	SetTimeToLiveSpecification(val interface{})
	UpdatedProperites() *map[string]interface{}
	WriteProvisionedThroughputSettings() interface{}
	SetWriteProvisionedThroughputSettings(val interface{})
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnGlobalTable
type jsiiProxy_CfnGlobalTable struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnGlobalTable) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) AttributeDefinitions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributeDefinitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) AttrStreamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStreamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) AttrTableId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrTableId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) BillingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) GlobalSecondaryIndexes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globalSecondaryIndexes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) KeySchema() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keySchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) LocalSecondaryIndexes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localSecondaryIndexes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) Replicas() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) SseSpecification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sseSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) StreamSpecification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) TimeToLiveSpecification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeToLiveSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalTable) WriteProvisionedThroughputSettings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeProvisionedThroughputSettings",
		&returns,
	)
	return returns
}


// Create a new `AWS::DynamoDB::GlobalTable`.
func NewCfnGlobalTable(scope awscdk.Construct, id *string, props *CfnGlobalTableProps) CfnGlobalTable {
	_init_.Initialize()

	j := jsiiProxy_CfnGlobalTable{}

	_jsii_.Create(
		"monocdk.aws_dynamodb.CfnGlobalTable",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DynamoDB::GlobalTable`.
func NewCfnGlobalTable_Override(c CfnGlobalTable, scope awscdk.Construct, id *string, props *CfnGlobalTableProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_dynamodb.CfnGlobalTable",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnGlobalTable) SetAttributeDefinitions(val interface{}) {
	_jsii_.Set(
		j,
		"attributeDefinitions",
		val,
	)
}

func (j *jsiiProxy_CfnGlobalTable) SetBillingMode(val *string) {
	_jsii_.Set(
		j,
		"billingMode",
		val,
	)
}

func (j *jsiiProxy_CfnGlobalTable) SetGlobalSecondaryIndexes(val interface{}) {
	_jsii_.Set(
		j,
		"globalSecondaryIndexes",
		val,
	)
}

func (j *jsiiProxy_CfnGlobalTable) SetKeySchema(val interface{}) {
	_jsii_.Set(
		j,
		"keySchema",
		val,
	)
}

func (j *jsiiProxy_CfnGlobalTable) SetLocalSecondaryIndexes(val interface{}) {
	_jsii_.Set(
		j,
		"localSecondaryIndexes",
		val,
	)
}

func (j *jsiiProxy_CfnGlobalTable) SetReplicas(val interface{}) {
	_jsii_.Set(
		j,
		"replicas",
		val,
	)
}

func (j *jsiiProxy_CfnGlobalTable) SetSseSpecification(val interface{}) {
	_jsii_.Set(
		j,
		"sseSpecification",
		val,
	)
}

func (j *jsiiProxy_CfnGlobalTable) SetStreamSpecification(val interface{}) {
	_jsii_.Set(
		j,
		"streamSpecification",
		val,
	)
}

func (j *jsiiProxy_CfnGlobalTable) SetTableName(val *string) {
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

func (j *jsiiProxy_CfnGlobalTable) SetTimeToLiveSpecification(val interface{}) {
	_jsii_.Set(
		j,
		"timeToLiveSpecification",
		val,
	)
}

func (j *jsiiProxy_CfnGlobalTable) SetWriteProvisionedThroughputSettings(val interface{}) {
	_jsii_.Set(
		j,
		"writeProvisionedThroughputSettings",
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
func CfnGlobalTable_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_dynamodb.CfnGlobalTable",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnGlobalTable_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_dynamodb.CfnGlobalTable",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnGlobalTable_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_dynamodb.CfnGlobalTable",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGlobalTable_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_dynamodb.CfnGlobalTable",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
func (c *jsiiProxy_CfnGlobalTable) AddDeletionOverride(path *string) {
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
func (c *jsiiProxy_CfnGlobalTable) AddDependsOn(target awscdk.CfnResource) {
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
func (c *jsiiProxy_CfnGlobalTable) AddMetadata(key *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnGlobalTable) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
func (c *jsiiProxy_CfnGlobalTable) AddPropertyDeletionOverride(propertyPath *string) {
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
func (c *jsiiProxy_CfnGlobalTable) AddPropertyOverride(propertyPath *string, value interface{}) {
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
// Experimental.
func (c *jsiiProxy_CfnGlobalTable) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
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
func (c *jsiiProxy_CfnGlobalTable) GetAtt(attributeName *string) awscdk.Reference {
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
func (c *jsiiProxy_CfnGlobalTable) GetMetadata(key *string) interface{} {
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
func (c *jsiiProxy_CfnGlobalTable) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnGlobalTable) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnGlobalTable) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnGlobalTable) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnGlobalTable) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnGlobalTable) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnGlobalTable) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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
func (c *jsiiProxy_CfnGlobalTable) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnGlobalTable) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
func (c *jsiiProxy_CfnGlobalTable) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnGlobalTable) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnGlobalTable) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Represents an attribute for describing the key schema for the table and indexes.
//
// TODO: EXAMPLE
//
type CfnGlobalTable_AttributeDefinitionProperty struct {
	// A name for the attribute.
	AttributeName *string `json:"attributeName" yaml:"attributeName"`
	// The data type for the attribute, where:.
	//
	// - `S` - the attribute is of type String
	// - `N` - the attribute is of type Number
	// - `B` - the attribute is of type Binary
	AttributeType *string `json:"attributeType" yaml:"attributeType"`
}

// Configures a scalable target and an autoscaling policy for a table or global secondary index's read or write capacity.
//
// TODO: EXAMPLE
//
type CfnGlobalTable_CapacityAutoScalingSettingsProperty struct {
	// The maximum provisioned capacity units for the global table.
	MaxCapacity *float64 `json:"maxCapacity" yaml:"maxCapacity"`
	// The minimum provisioned capacity units for the global table.
	MinCapacity *float64 `json:"minCapacity" yaml:"minCapacity"`
	// Defines a target tracking scaling policy.
	TargetTrackingScalingPolicyConfiguration interface{} `json:"targetTrackingScalingPolicyConfiguration" yaml:"targetTrackingScalingPolicyConfiguration"`
	// When switching billing mode from `PAY_PER_REQUEST` to `PROVISIONED` , DynamoDB requires you to specify read and write capacity unit values for the table and for each global secondary index.
	//
	// These values will be applied to all replicas. The table will use these provisioned values until CloudFormation creates the autoscaling policies you configured in your template. CloudFormation cannot determine what capacity the table and its global secondary indexes will require in this time period, since they are application-dependent.
	//
	// If you want to switch a table's billing mode from `PAY_PER_REQUEST` to `PROVISIONED` , you must specify a value for this property for each autoscaled resource. If you specify different values for the same resource in different regions, CloudFormation will use the highest value found in either the `SeedCapacity` or `ReadCapacityUnits` properties. For example, if your global secondary index `myGSI` has a `SeedCapacity` of 10 in us-east-1 and a fixed `ReadCapacityUnits` of 20 in eu-west-1, CloudFormation will initially set the read capacity for `myGSI` to 20. Note that if you disable `ScaleIn` for `myGSI` in us-east-1, its read capacity units might not be set back to 10.
	//
	// You must also specify a value for `SeedCapacity` when you plan to switch a table's billing mode from `PROVISIONED` to `PAY_PER_REQUEST` , because CloudFormation might need to roll back the operation (reverting the billing mode to `PROVISIONED` ) and this cannot succeed without specifying a value for `SeedCapacity` .
	SeedCapacity *float64 `json:"seedCapacity" yaml:"seedCapacity"`
}

// Configures contributor insights settings for a replica or one of its indexes.
//
// TODO: EXAMPLE
//
type CfnGlobalTable_ContributorInsightsSpecificationProperty struct {
	// Indicates whether CloudWatch Contributor Insights are to be enabled (true) or disabled (false).
	Enabled interface{} `json:"enabled" yaml:"enabled"`
}

// Allows you to specify a global secondary index for the global table.
//
// The index will be defined on all replicas.
//
// TODO: EXAMPLE
//
type CfnGlobalTable_GlobalSecondaryIndexProperty struct {
	// The name of the global secondary index.
	//
	// The name must be unique among all other indexes on this table.
	IndexName *string `json:"indexName" yaml:"indexName"`
	// The complete key schema for a global secondary index, which consists of one or more pairs of attribute names and key types:  - `HASH` - partition key - `RANGE` - sort key  > The partition key of an item is also known as its *hash attribute* .
	//
	// The term "hash attribute" derives from DynamoDB's usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values.
	// >
	// > The sort key of an item is also known as its *range attribute* . The term "range attribute" derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.
	KeySchema interface{} `json:"keySchema" yaml:"keySchema"`
	// Represents attributes that are copied (projected) from the table into the global secondary index.
	//
	// These are in addition to the primary key attributes and index key attributes, which are automatically projected.
	Projection interface{} `json:"projection" yaml:"projection"`
	// Defines write capacity settings for the global secondary index.
	//
	// You must specify a value for this property if the table's `BillingMode` is `PROVISIONED` . All replicas will have the same write capacity settings for this global secondary index.
	WriteProvisionedThroughputSettings interface{} `json:"writeProvisionedThroughputSettings" yaml:"writeProvisionedThroughputSettings"`
}

// Represents *a single element* of a key schema.
//
// A key schema specifies the attributes that make up the primary key of a table, or the key attributes of an index.
//
// A `KeySchemaElement` represents exactly one attribute of the primary key. For example, a simple primary key would be represented by one `KeySchemaElement` (for the partition key). A composite primary key would require one `KeySchemaElement` for the partition key, and another `KeySchemaElement` for the sort key.
//
// A `KeySchemaElement` must be a scalar, top-level attribute (not a nested attribute). The data type must be one of String, Number, or Binary. The attribute cannot be nested within a List or a Map.
//
// TODO: EXAMPLE
//
type CfnGlobalTable_KeySchemaProperty struct {
	// The name of a key attribute.
	AttributeName *string `json:"attributeName" yaml:"attributeName"`
	// The role that this key attribute will assume:.
	//
	// - `HASH` - partition key
	// - `RANGE` - sort key
	//
	// > The partition key of an item is also known as its *hash attribute* . The term "hash attribute" derives from DynamoDB's usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values.
	// >
	// > The sort key of an item is also known as its *range attribute* . The term "range attribute" derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.
	KeyType *string `json:"keyType" yaml:"keyType"`
}

// Represents the properties of a local secondary index.
//
// A local secondary index can only be created when its parent table is created.
//
// TODO: EXAMPLE
//
type CfnGlobalTable_LocalSecondaryIndexProperty struct {
	// The name of the local secondary index.
	//
	// The name must be unique among all other indexes on this table.
	IndexName *string `json:"indexName" yaml:"indexName"`
	// The complete key schema for the local secondary index, consisting of one or more pairs of attribute names and key types:  - `HASH` - partition key - `RANGE` - sort key  > The partition key of an item is also known as its *hash attribute* .
	//
	// The term "hash attribute" derives from DynamoDB's usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values.
	// >
	// > The sort key of an item is also known as its *range attribute* . The term "range attribute" derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.
	KeySchema interface{} `json:"keySchema" yaml:"keySchema"`
	// Represents attributes that are copied (projected) from the table into the local secondary index.
	//
	// These are in addition to the primary key attributes and index key attributes, which are automatically projected.
	Projection interface{} `json:"projection" yaml:"projection"`
}

// Represents the settings used to enable point in time recovery.
//
// TODO: EXAMPLE
//
type CfnGlobalTable_PointInTimeRecoverySpecificationProperty struct {
	// Indicates whether point in time recovery is enabled (true) or disabled (false) on the table.
	PointInTimeRecoveryEnabled interface{} `json:"pointInTimeRecoveryEnabled" yaml:"pointInTimeRecoveryEnabled"`
}

// Represents attributes that are copied (projected) from the table into an index.
//
// These are in addition to the primary key attributes and index key attributes, which are automatically projected.
//
// TODO: EXAMPLE
//
type CfnGlobalTable_ProjectionProperty struct {
	// Represents the non-key attribute names which will be projected into the index.
	//
	// For local secondary indexes, the total count of `NonKeyAttributes` summed across all of the local secondary indexes, must not exceed 20. If you project the same attribute into two different indexes, this counts as two distinct attributes when determining the total.
	NonKeyAttributes *[]*string `json:"nonKeyAttributes" yaml:"nonKeyAttributes"`
	// The set of attributes that are projected into the index:.
	//
	// - `KEYS_ONLY` - Only the index and primary keys are projected into the index.
	// - `INCLUDE` - In addition to the attributes described in `KEYS_ONLY` , the secondary index will include other non-key attributes that you specify.
	// - `ALL` - All of the table attributes are projected into the index.
	ProjectionType *string `json:"projectionType" yaml:"projectionType"`
}

// Allows you to specify the read capacity settings for a replica table or a replica global secondary index when the `BillingMode` is set to `PROVISIONED` .
//
// You must specify a value for either `ReadCapacityUnits` or `ReadCapacityAutoScalingSettings` , but not both. You can switch between fixed capacity and auto scaling.
//
// TODO: EXAMPLE
//
type CfnGlobalTable_ReadProvisionedThroughputSettingsProperty struct {
	// Specifies auto scaling settings for the replica table or global secondary index.
	ReadCapacityAutoScalingSettings interface{} `json:"readCapacityAutoScalingSettings" yaml:"readCapacityAutoScalingSettings"`
	// Specifies a fixed read capacity for the replica table or global secondary index.
	ReadCapacityUnits *float64 `json:"readCapacityUnits" yaml:"readCapacityUnits"`
}

// Represents the properties of a global secondary index that can be set on a per-replica basis.
//
// TODO: EXAMPLE
//
type CfnGlobalTable_ReplicaGlobalSecondaryIndexSpecificationProperty struct {
	// The name of the global secondary index.
	//
	// The name must be unique among all other indexes on this table.
	IndexName *string `json:"indexName" yaml:"indexName"`
	// Updates the status for contributor insights for a specific table or index.
	//
	// CloudWatch Contributor Insights for DynamoDB graphs display the partition key and (if applicable) sort key of frequently accessed items and frequently throttled items in plaintext. If you require the use of AWS Key Management Service (KMS) to encrypt this table’s partition key and sort key data with an AWS managed key or customer managed key, you should not enable CloudWatch Contributor Insights for DynamoDB for this table.
	ContributorInsightsSpecification interface{} `json:"contributorInsightsSpecification" yaml:"contributorInsightsSpecification"`
	// Allows you to specify the read capacity settings for a replica global secondary index when the `BillingMode` is set to `PROVISIONED` .
	ReadProvisionedThroughputSettings interface{} `json:"readProvisionedThroughputSettings" yaml:"readProvisionedThroughputSettings"`
}

// Allows you to specify a KMS key identifier to be used for server-side encryption.
//
// The key can be specified via ARN, key ID, or alias. The key must be created in the same region as the replica.
//
// TODO: EXAMPLE
//
type CfnGlobalTable_ReplicaSSESpecificationProperty struct {
	// The AWS KMS key that should be used for the AWS KMS encryption.
	//
	// To specify a key, use its key ID, Amazon Resource Name (ARN), alias name, or alias ARN. Note that you should only provide this parameter if the key is different from the default DynamoDB key `alias/aws/dynamodb` .
	KmsMasterKeyId *string `json:"kmsMasterKeyId" yaml:"kmsMasterKeyId"`
}

// Defines settings specific to a single replica of a global table.
//
// TODO: EXAMPLE
//
type CfnGlobalTable_ReplicaSpecificationProperty struct {
	// The region in which this replica exists.
	Region *string `json:"region" yaml:"region"`
	// The settings used to enable or disable CloudWatch Contributor Insights for the specified replica.
	//
	// When not specified, defaults to contributor insights disabled for the replica.
	ContributorInsightsSpecification interface{} `json:"contributorInsightsSpecification" yaml:"contributorInsightsSpecification"`
	// Defines additional settings for the global secondary indexes of this replica.
	GlobalSecondaryIndexes interface{} `json:"globalSecondaryIndexes" yaml:"globalSecondaryIndexes"`
	// The settings used to enable point in time recovery.
	//
	// When not specified, defaults to point in time recovery disabled for the replica.
	PointInTimeRecoverySpecification interface{} `json:"pointInTimeRecoverySpecification" yaml:"pointInTimeRecoverySpecification"`
	// Defines read capacity settings for the replica table.
	ReadProvisionedThroughputSettings interface{} `json:"readProvisionedThroughputSettings" yaml:"readProvisionedThroughputSettings"`
	// Allows you to specify a customer-managed key for the replica.
	//
	// When using customer-managed keys for server-side encryption, this property must have a value in all replicas.
	SseSpecification interface{} `json:"sseSpecification" yaml:"sseSpecification"`
	// An array of key-value pairs to apply to this replica.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
}

// Represents the settings used to enable server-side encryption.
//
// TODO: EXAMPLE
//
type CfnGlobalTable_SSESpecificationProperty struct {
	// Indicates whether server-side encryption is performed using an AWS managed key or an AWS owned key.
	//
	// If disabled (false) or not specified, server-side encryption uses an AWS owned key. If enabled (true), the server-side encryption type is set to KMS and an AWS managed key is used ( AWS KMS charges apply). If you choose to use KMS encryption, you can also use customer managed KMS keys by specifying them in the `ReplicaSpecification.SSESpecification` object. You cannot mix AWS managed and customer managed KMS keys.
	SseEnabled interface{} `json:"sseEnabled" yaml:"sseEnabled"`
	// Server-side encryption type. The only supported value is:.
	//
	// - `KMS` - Server-side encryption that uses AWS Key Management Service . The key is stored in your account and is managed by AWS KMS ( AWS KMS charges apply).
	SseType *string `json:"sseType" yaml:"sseType"`
}

// Represents the DynamoDB Streams configuration for a table in DynamoDB.
//
// You can only modify this value if your `AWS::DynamoDB::GlobalTable` contains only one entry in `Replicas` . You must specify a value for this property if your `AWS::DynamoDB::GlobalTable` contains more than one replica.
//
// TODO: EXAMPLE
//
type CfnGlobalTable_StreamSpecificationProperty struct {
	// When an item in the table is modified, `StreamViewType` determines what information is written to the stream for this table.
	//
	// Valid values for `StreamViewType` are:
	//
	// - `KEYS_ONLY` - Only the key attributes of the modified item are written to the stream.
	// - `NEW_IMAGE` - The entire item, as it appears after it was modified, is written to the stream.
	// - `OLD_IMAGE` - The entire item, as it appeared before it was modified, is written to the stream.
	// - `NEW_AND_OLD_IMAGES` - Both the new and the old item images of the item are written to the stream.
	StreamViewType *string `json:"streamViewType" yaml:"streamViewType"`
}

// Defines a target tracking scaling policy.
//
// TODO: EXAMPLE
//
type CfnGlobalTable_TargetTrackingScalingPolicyConfigurationProperty struct {
	// Defines a target value for the scaling policy.
	TargetValue *float64 `json:"targetValue" yaml:"targetValue"`
	// Indicates whether scale in by the target tracking scaling policy is disabled.
	//
	// The default value is `false` .
	DisableScaleIn interface{} `json:"disableScaleIn" yaml:"disableScaleIn"`
	// The amount of time, in seconds, after a scale-in activity completes before another scale-in activity can start.
	ScaleInCooldown *float64 `json:"scaleInCooldown" yaml:"scaleInCooldown"`
	// The amount of time, in seconds, after a scale-out activity completes before another scale-out activity can start.
	ScaleOutCooldown *float64 `json:"scaleOutCooldown" yaml:"scaleOutCooldown"`
}

// Represents the settings used to enable or disable Time to Live (TTL) for the specified table.
//
// All replicas will have the same time to live configuration.
//
// TODO: EXAMPLE
//
type CfnGlobalTable_TimeToLiveSpecificationProperty struct {
	// Indicates whether TTL is to be enabled (true) or disabled (false) on the table.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
	// The name of the attribute used to store the expiration time for items in the table.
	//
	// Currently, you cannot directly change the attribute name used to evaluate time to live. In order to do so, you must first disable time to live, and then re-enable it with the new attribute name. It can take up to one hour for changes to time to live to take effect. If you attempt to modify time to live within that time window, your stack operation might be delayed.
	AttributeName *string `json:"attributeName" yaml:"attributeName"`
}

// Specifies an auto scaling policy for write capacity.
//
// This policy will be applied to all replicas. This setting must be specified if `BillingMode` is set to `PROVISIONED` .
//
// TODO: EXAMPLE
//
type CfnGlobalTable_WriteProvisionedThroughputSettingsProperty struct {
	// Specifies auto scaling settings for the replica table or global secondary index.
	WriteCapacityAutoScalingSettings interface{} `json:"writeCapacityAutoScalingSettings" yaml:"writeCapacityAutoScalingSettings"`
}

// Properties for defining a `CfnGlobalTable`.
//
// TODO: EXAMPLE
//
type CfnGlobalTableProps struct {
	// A list of attributes that describe the key schema for the global table and indexes.
	AttributeDefinitions interface{} `json:"attributeDefinitions" yaml:"attributeDefinitions"`
	// Specifies the attributes that make up the primary key for the table.
	//
	// The attributes in the `KeySchema` property must also be defined in the `AttributeDefinitions` property.
	KeySchema interface{} `json:"keySchema" yaml:"keySchema"`
	// Specifies the list of replicas for your global table.
	//
	// The list must contain at least one element, the region where the stack defining the global table is deployed. For example, if you define your table in a stack deployed to us-east-1, you must have an entry in `Replicas` with the region us-east-1. You cannot remove the replica in the stack region.
	//
	// > Adding a replica might take a few minutes for an empty table, or up to several hours for large tables. If you want to add or remove a replica, we recommend submitting an `UpdateStack` operation containing only that change.
	// >
	// > If you add or delete a replica during an update, we recommend that you don't update any other resources. If your stack fails to update and is rolled back while adding a new replica, you might need to manually delete the replica.
	//
	// You can create a new global table with up to two replicas. You can add or remove replicas after table creation, but you can only add or remove a single replica in each update.
	Replicas interface{} `json:"replicas" yaml:"replicas"`
	// Specifies how you are charged for read and write throughput and how you manage capacity. Valid values are:.
	//
	// - `PAY_PER_REQUEST`
	// - `PROVISIONED`
	//
	// All replicas in your global table will have the same billing mode. If you use `PROVISIONED` billing mode, you must provide an auto scaling configuration via the `WriteProvisionedThroughputSettings` property. The default value of this property is `PROVISIONED` .
	BillingMode *string `json:"billingMode" yaml:"billingMode"`
	// Global secondary indexes to be created on the global table.
	//
	// You can create up to 20 global secondary indexes. Each replica in your global table will have the same global secondary index settings. You can only create or delete one global secondary index in a single stack operation.
	//
	// Since the backfilling of an index could take a long time, CloudFormation does not wait for the index to become active. If a stack operation rolls back, CloudFormation might not delete an index that has been added. In that case, you will need to delete the index manually.
	GlobalSecondaryIndexes interface{} `json:"globalSecondaryIndexes" yaml:"globalSecondaryIndexes"`
	// Local secondary indexes to be created on the table.
	//
	// You can create up to five local secondary indexes. Each index is scoped to a given hash key value. The size of each hash key can be up to 10 gigabytes. Each replica in your global table will have the same local secondary index settings.
	LocalSecondaryIndexes interface{} `json:"localSecondaryIndexes" yaml:"localSecondaryIndexes"`
	// Specifies the settings to enable server-side encryption.
	//
	// These settings will be applied to all replicas. If you plan to use customer-managed KMS keys, you must provide a key for each replica using the `ReplicaSpecification.ReplicaSSESpecification` property.
	SseSpecification interface{} `json:"sseSpecification" yaml:"sseSpecification"`
	// Specifies the streams settings on your global table.
	//
	// You must provide a value for this property if your global table contains more than one replica. You can only change the streams settings if your global table has only one replica.
	StreamSpecification interface{} `json:"streamSpecification" yaml:"streamSpecification"`
	// A name for the global table.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID as the table name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	TableName *string `json:"tableName" yaml:"tableName"`
	// Specifies the time to live (TTL) settings for the table.
	//
	// This setting will be applied to all replicas.
	TimeToLiveSpecification interface{} `json:"timeToLiveSpecification" yaml:"timeToLiveSpecification"`
	// Specifies an auto scaling policy for write capacity.
	//
	// This policy will be applied to all replicas. This setting must be specified if `BillingMode` is set to `PROVISIONED` .
	WriteProvisionedThroughputSettings interface{} `json:"writeProvisionedThroughputSettings" yaml:"writeProvisionedThroughputSettings"`
}

// A CloudFormation `AWS::DynamoDB::Table`.
//
// The `AWS::DynamoDB::Table` resource creates a DynamoDB table. For more information, see [CreateTable](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_CreateTable.html) in the *Amazon DynamoDB API Reference* .
//
// You should be aware of the following behaviors when working with DynamoDB tables:
//
// - AWS CloudFormation typically creates DynamoDB tables in parallel. However, if your template includes multiple DynamoDB tables with indexes, you must declare dependencies so that the tables are created sequentially. Amazon DynamoDB limits the number of tables with secondary indexes that are in the creating state. If you create multiple tables with indexes at the same time, DynamoDB returns an error and the stack operation fails. For an example, see [DynamoDB Table with a DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-examples-dependson) .
//
// TODO: EXAMPLE
//
type CfnTable interface {
	awscdk.CfnResource
	awscdk.IInspectable
	AttrArn() *string
	AttributeDefinitions() interface{}
	SetAttributeDefinitions(val interface{})
	AttrStreamArn() *string
	BillingMode() *string
	SetBillingMode(val *string)
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnResourceType() *string
	ContributorInsightsSpecification() interface{}
	SetContributorInsightsSpecification(val interface{})
	CreationStack() *[]*string
	GlobalSecondaryIndexes() interface{}
	SetGlobalSecondaryIndexes(val interface{})
	KeySchema() interface{}
	SetKeySchema(val interface{})
	KinesisStreamSpecification() interface{}
	SetKinesisStreamSpecification(val interface{})
	LocalSecondaryIndexes() interface{}
	SetLocalSecondaryIndexes(val interface{})
	LogicalId() *string
	Node() awscdk.ConstructNode
	PointInTimeRecoverySpecification() interface{}
	SetPointInTimeRecoverySpecification(val interface{})
	ProvisionedThroughput() interface{}
	SetProvisionedThroughput(val interface{})
	Ref() *string
	SseSpecification() interface{}
	SetSseSpecification(val interface{})
	Stack() awscdk.Stack
	StreamSpecification() interface{}
	SetStreamSpecification(val interface{})
	TableClass() *string
	SetTableClass(val *string)
	TableName() *string
	SetTableName(val *string)
	Tags() awscdk.TagManager
	TimeToLiveSpecification() interface{}
	SetTimeToLiveSpecification(val interface{})
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
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	OverrideLogicalId(newLogicalId *string)
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	ShouldSynthesize() *bool
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_CfnTable) AttributeDefinitions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributeDefinitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTable) AttrStreamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStreamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTable) BillingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingMode",
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

func (j *jsiiProxy_CfnTable) ContributorInsightsSpecification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contributorInsightsSpecification",
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

func (j *jsiiProxy_CfnTable) GlobalSecondaryIndexes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globalSecondaryIndexes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTable) KeySchema() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keySchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTable) KinesisStreamSpecification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kinesisStreamSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTable) LocalSecondaryIndexes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localSecondaryIndexes",
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

func (j *jsiiProxy_CfnTable) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTable) PointInTimeRecoverySpecification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pointInTimeRecoverySpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTable) ProvisionedThroughput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionedThroughput",
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

func (j *jsiiProxy_CfnTable) SseSpecification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sseSpecification",
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

func (j *jsiiProxy_CfnTable) StreamSpecification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTable) TableClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableClass",
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

func (j *jsiiProxy_CfnTable) TimeToLiveSpecification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeToLiveSpecification",
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


// Create a new `AWS::DynamoDB::Table`.
func NewCfnTable(scope awscdk.Construct, id *string, props *CfnTableProps) CfnTable {
	_init_.Initialize()

	j := jsiiProxy_CfnTable{}

	_jsii_.Create(
		"monocdk.aws_dynamodb.CfnTable",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DynamoDB::Table`.
func NewCfnTable_Override(c CfnTable, scope awscdk.Construct, id *string, props *CfnTableProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_dynamodb.CfnTable",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnTable) SetAttributeDefinitions(val interface{}) {
	_jsii_.Set(
		j,
		"attributeDefinitions",
		val,
	)
}

func (j *jsiiProxy_CfnTable) SetBillingMode(val *string) {
	_jsii_.Set(
		j,
		"billingMode",
		val,
	)
}

func (j *jsiiProxy_CfnTable) SetContributorInsightsSpecification(val interface{}) {
	_jsii_.Set(
		j,
		"contributorInsightsSpecification",
		val,
	)
}

func (j *jsiiProxy_CfnTable) SetGlobalSecondaryIndexes(val interface{}) {
	_jsii_.Set(
		j,
		"globalSecondaryIndexes",
		val,
	)
}

func (j *jsiiProxy_CfnTable) SetKeySchema(val interface{}) {
	_jsii_.Set(
		j,
		"keySchema",
		val,
	)
}

func (j *jsiiProxy_CfnTable) SetKinesisStreamSpecification(val interface{}) {
	_jsii_.Set(
		j,
		"kinesisStreamSpecification",
		val,
	)
}

func (j *jsiiProxy_CfnTable) SetLocalSecondaryIndexes(val interface{}) {
	_jsii_.Set(
		j,
		"localSecondaryIndexes",
		val,
	)
}

func (j *jsiiProxy_CfnTable) SetPointInTimeRecoverySpecification(val interface{}) {
	_jsii_.Set(
		j,
		"pointInTimeRecoverySpecification",
		val,
	)
}

func (j *jsiiProxy_CfnTable) SetProvisionedThroughput(val interface{}) {
	_jsii_.Set(
		j,
		"provisionedThroughput",
		val,
	)
}

func (j *jsiiProxy_CfnTable) SetSseSpecification(val interface{}) {
	_jsii_.Set(
		j,
		"sseSpecification",
		val,
	)
}

func (j *jsiiProxy_CfnTable) SetStreamSpecification(val interface{}) {
	_jsii_.Set(
		j,
		"streamSpecification",
		val,
	)
}

func (j *jsiiProxy_CfnTable) SetTableClass(val *string) {
	_jsii_.Set(
		j,
		"tableClass",
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

func (j *jsiiProxy_CfnTable) SetTimeToLiveSpecification(val interface{}) {
	_jsii_.Set(
		j,
		"timeToLiveSpecification",
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
func CfnTable_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_dynamodb.CfnTable",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnTable_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_dynamodb.CfnTable",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnTable_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_dynamodb.CfnTable",
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
		"monocdk.aws_dynamodb.CfnTable",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

// Syntactic sugar for `addOverride(path, undefined)`.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
func (c *jsiiProxy_CfnTable) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

// Adds an override that deletes the value of a property from the resource definition.
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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
// Experimental.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnTable) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnTable) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnTable) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Overrides the auto-generated logical ID with a specific ID.
// Experimental.
func (c *jsiiProxy_CfnTable) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CfnTable) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
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
// Experimental.
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CfnTable) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
//
// Returns: a string representation of this resource
// Experimental.
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CfnTable) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Experimental.
func (c *jsiiProxy_CfnTable) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

// Represents an attribute for describing the key schema for the table and indexes.
//
// TODO: EXAMPLE
//
type CfnTable_AttributeDefinitionProperty struct {
	// A name for the attribute.
	AttributeName *string `json:"attributeName" yaml:"attributeName"`
	// The data type for the attribute, where:.
	//
	// - `S` - the attribute is of type String
	// - `N` - the attribute is of type Number
	// - `B` - the attribute is of type Binary
	AttributeType *string `json:"attributeType" yaml:"attributeType"`
}

// The settings used to enable or disable CloudWatch Contributor Insights.
//
// TODO: EXAMPLE
//
type CfnTable_ContributorInsightsSpecificationProperty struct {
	// Indicates whether CloudWatch Contributor Insights are to be enabled (true) or disabled (false).
	Enabled interface{} `json:"enabled" yaml:"enabled"`
}

// Represents the properties of a global secondary index.
//
// TODO: EXAMPLE
//
type CfnTable_GlobalSecondaryIndexProperty struct {
	// The name of the global secondary index.
	//
	// The name must be unique among all other indexes on this table.
	IndexName *string `json:"indexName" yaml:"indexName"`
	// The complete key schema for a global secondary index, which consists of one or more pairs of attribute names and key types:  - `HASH` - partition key - `RANGE` - sort key  > The partition key of an item is also known as its *hash attribute* .
	//
	// The term "hash attribute" derives from DynamoDB's usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values.
	// >
	// > The sort key of an item is also known as its *range attribute* . The term "range attribute" derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.
	KeySchema interface{} `json:"keySchema" yaml:"keySchema"`
	// Represents attributes that are copied (projected) from the table into the global secondary index.
	//
	// These are in addition to the primary key attributes and index key attributes, which are automatically projected.
	Projection interface{} `json:"projection" yaml:"projection"`
	// The settings used to enable or disable CloudWatch Contributor Insights for the specified global secondary index.
	ContributorInsightsSpecification interface{} `json:"contributorInsightsSpecification" yaml:"contributorInsightsSpecification"`
	// Represents the provisioned throughput settings for the specified global secondary index.
	//
	// For current minimum and maximum provisioned throughput values, see [Service, Account, and Table Quotas](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html) in the *Amazon DynamoDB Developer Guide* .
	ProvisionedThroughput interface{} `json:"provisionedThroughput" yaml:"provisionedThroughput"`
}

// Represents *a single element* of a key schema.
//
// A key schema specifies the attributes that make up the primary key of a table, or the key attributes of an index.
//
// A `KeySchemaElement` represents exactly one attribute of the primary key. For example, a simple primary key would be represented by one `KeySchemaElement` (for the partition key). A composite primary key would require one `KeySchemaElement` for the partition key, and another `KeySchemaElement` for the sort key.
//
// A `KeySchemaElement` must be a scalar, top-level attribute (not a nested attribute). The data type must be one of String, Number, or Binary. The attribute cannot be nested within a List or a Map.
//
// TODO: EXAMPLE
//
type CfnTable_KeySchemaProperty struct {
	// The name of a key attribute.
	AttributeName *string `json:"attributeName" yaml:"attributeName"`
	// The role that this key attribute will assume:.
	//
	// - `HASH` - partition key
	// - `RANGE` - sort key
	//
	// > The partition key of an item is also known as its *hash attribute* . The term "hash attribute" derives from DynamoDB's usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values.
	// >
	// > The sort key of an item is also known as its *range attribute* . The term "range attribute" derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.
	KeyType *string `json:"keyType" yaml:"keyType"`
}

// The Kinesis Data Streams configuration for the specified table.
//
// TODO: EXAMPLE
//
type CfnTable_KinesisStreamSpecificationProperty struct {
	// The ARN for a specific Kinesis data stream.
	//
	// Length Constraints: Minimum length of 37. Maximum length of 1024.
	StreamArn *string `json:"streamArn" yaml:"streamArn"`
}

// Represents the properties of a local secondary index.
//
// A local secondary index can only be created when its parent table is created.
//
// TODO: EXAMPLE
//
type CfnTable_LocalSecondaryIndexProperty struct {
	// The name of the local secondary index.
	//
	// The name must be unique among all other indexes on this table.
	IndexName *string `json:"indexName" yaml:"indexName"`
	// The complete key schema for the local secondary index, consisting of one or more pairs of attribute names and key types:  - `HASH` - partition key - `RANGE` - sort key  > The partition key of an item is also known as its *hash attribute* .
	//
	// The term "hash attribute" derives from DynamoDB's usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values.
	// >
	// > The sort key of an item is also known as its *range attribute* . The term "range attribute" derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.
	KeySchema interface{} `json:"keySchema" yaml:"keySchema"`
	// Represents attributes that are copied (projected) from the table into the local secondary index.
	//
	// These are in addition to the primary key attributes and index key attributes, which are automatically projected.
	Projection interface{} `json:"projection" yaml:"projection"`
}

// The settings used to enable point in time recovery.
//
// TODO: EXAMPLE
//
type CfnTable_PointInTimeRecoverySpecificationProperty struct {
	// Indicates whether point in time recovery is enabled (true) or disabled (false) on the table.
	PointInTimeRecoveryEnabled interface{} `json:"pointInTimeRecoveryEnabled" yaml:"pointInTimeRecoveryEnabled"`
}

// Represents attributes that are copied (projected) from the table into an index.
//
// These are in addition to the primary key attributes and index key attributes, which are automatically projected.
//
// TODO: EXAMPLE
//
type CfnTable_ProjectionProperty struct {
	// Represents the non-key attribute names which will be projected into the index.
	//
	// For local secondary indexes, the total count of `NonKeyAttributes` summed across all of the local secondary indexes, must not exceed 20. If you project the same attribute into two different indexes, this counts as two distinct attributes when determining the total.
	NonKeyAttributes *[]*string `json:"nonKeyAttributes" yaml:"nonKeyAttributes"`
	// The set of attributes that are projected into the index:.
	//
	// - `KEYS_ONLY` - Only the index and primary keys are projected into the index.
	// - `INCLUDE` - In addition to the attributes described in `KEYS_ONLY` , the secondary index will include other non-key attributes that you specify.
	// - `ALL` - All of the table attributes are projected into the index.
	ProjectionType *string `json:"projectionType" yaml:"projectionType"`
}

// Throughput for the specified table, which consists of values for `ReadCapacityUnits` and `WriteCapacityUnits` .
//
// For more information about the contents of a provisioned throughput structure, see [Amazon DynamoDB Table ProvisionedThroughput](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-provisionedthroughput.html) .
//
// TODO: EXAMPLE
//
type CfnTable_ProvisionedThroughputProperty struct {
	// The maximum number of strongly consistent reads consumed per second before DynamoDB returns a `ThrottlingException` .
	//
	// For more information, see [Specifying Read and Write Requirements](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.html#ProvisionedThroughput) in the *Amazon DynamoDB Developer Guide* .
	//
	// If read/write capacity mode is `PAY_PER_REQUEST` the value is set to 0.
	ReadCapacityUnits *float64 `json:"readCapacityUnits" yaml:"readCapacityUnits"`
	// The maximum number of writes consumed per second before DynamoDB returns a `ThrottlingException` .
	//
	// For more information, see [Specifying Read and Write Requirements](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.html#ProvisionedThroughput) in the *Amazon DynamoDB Developer Guide* .
	//
	// If read/write capacity mode is `PAY_PER_REQUEST` the value is set to 0.
	WriteCapacityUnits *float64 `json:"writeCapacityUnits" yaml:"writeCapacityUnits"`
}

// Represents the settings used to enable server-side encryption.
//
// TODO: EXAMPLE
//
type CfnTable_SSESpecificationProperty struct {
	// Indicates whether server-side encryption is done using an AWS managed key or an AWS owned key.
	//
	// If enabled (true), server-side encryption type is set to `KMS` and an AWS managed key is used ( AWS KMS charges apply). If disabled (false) or not specified, server-side encryption is set to AWS owned key.
	SseEnabled interface{} `json:"sseEnabled" yaml:"sseEnabled"`
	// The AWS KMS key that should be used for the AWS KMS encryption.
	//
	// To specify a key, use its key ID, Amazon Resource Name (ARN), alias name, or alias ARN. Note that you should only provide this parameter if the key is different from the default DynamoDB key `alias/aws/dynamodb` .
	KmsMasterKeyId *string `json:"kmsMasterKeyId" yaml:"kmsMasterKeyId"`
	// Server-side encryption type. The only supported value is:.
	//
	// - `KMS` - Server-side encryption that uses AWS Key Management Service . The key is stored in your account and is managed by AWS KMS ( AWS KMS charges apply).
	SseType *string `json:"sseType" yaml:"sseType"`
}

// Represents the DynamoDB Streams configuration for a table in DynamoDB.
//
// TODO: EXAMPLE
//
type CfnTable_StreamSpecificationProperty struct {
	// When an item in the table is modified, `StreamViewType` determines what information is written to the stream for this table.
	//
	// Valid values for `StreamViewType` are:
	//
	// - `KEYS_ONLY` - Only the key attributes of the modified item are written to the stream.
	// - `NEW_IMAGE` - The entire item, as it appears after it was modified, is written to the stream.
	// - `OLD_IMAGE` - The entire item, as it appeared before it was modified, is written to the stream.
	// - `NEW_AND_OLD_IMAGES` - Both the new and the old item images of the item are written to the stream.
	StreamViewType *string `json:"streamViewType" yaml:"streamViewType"`
}

// Represents the settings used to enable or disable Time to Live (TTL) for the specified table.
//
// TODO: EXAMPLE
//
type CfnTable_TimeToLiveSpecificationProperty struct {
	// The name of the TTL attribute used to store the expiration time for items in the table.
	//
	// > To update this property, you must first disable TTL then enable TTL with the new attribute name.
	AttributeName *string `json:"attributeName" yaml:"attributeName"`
	// Indicates whether TTL is to be enabled (true) or disabled (false) on the table.
	Enabled interface{} `json:"enabled" yaml:"enabled"`
}

// Properties for defining a `CfnTable`.
//
// TODO: EXAMPLE
//
type CfnTableProps struct {
	// Specifies the attributes that make up the primary key for the table.
	//
	// The attributes in the `KeySchema` property must also be defined in the `AttributeDefinitions` property.
	KeySchema interface{} `json:"keySchema" yaml:"keySchema"`
	// A list of attributes that describe the key schema for the table and indexes.
	//
	// This property is required to create a DynamoDB table.
	//
	// Update requires: [Some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt) . Replacement if you edit an existing AttributeDefinition.
	AttributeDefinitions interface{} `json:"attributeDefinitions" yaml:"attributeDefinitions"`
	// Specify how you are charged for read and write throughput and how you manage capacity.
	//
	// Valid values include:
	//
	// - `PROVISIONED` - We recommend using `PROVISIONED` for predictable workloads. `PROVISIONED` sets the billing mode to [Provisioned Mode](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.ProvisionedThroughput.Manual) .
	// - `PAY_PER_REQUEST` - We recommend using `PAY_PER_REQUEST` for unpredictable workloads. `PAY_PER_REQUEST` sets the billing mode to [On-Demand Mode](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.OnDemand) .
	//
	// If not specified, the default is `PROVISIONED` .
	BillingMode *string `json:"billingMode" yaml:"billingMode"`
	// The settings used to enable or disable CloudWatch Contributor Insights for the specified table.
	ContributorInsightsSpecification interface{} `json:"contributorInsightsSpecification" yaml:"contributorInsightsSpecification"`
	// Global secondary indexes to be created on the table. You can create up to 20 global secondary indexes.
	//
	// > If you update a table to include a new global secondary index, AWS CloudFormation initiates the index creation and then proceeds with the stack update. AWS CloudFormation doesn't wait for the index to complete creation because the backfilling phase can take a long time, depending on the size of the table. You can't use the index or update the table until the index's status is `ACTIVE` . You can track its status by using the DynamoDB [DescribeTable](https://docs.aws.amazon.com/cli/latest/reference/dynamodb/describe-table.html) command.
	// >
	// > If you add or delete an index during an update, we recommend that you don't update any other resources. If your stack fails to update and is rolled back while adding a new index, you must manually delete the index.
	// >
	// > Updates are not supported. The following are exceptions:
	// >
	// > - If you update either the contributor insights specification or the provisioned throughput values of global secondary indexes, you can update the table without interruption.
	// > - You can delete or add one global secondary index without interruption. If you do both in the same update (for example, by changing the index's logical ID), the update fails.
	GlobalSecondaryIndexes interface{} `json:"globalSecondaryIndexes" yaml:"globalSecondaryIndexes"`
	// The Kinesis Data Streams configuration for the specified table.
	KinesisStreamSpecification interface{} `json:"kinesisStreamSpecification" yaml:"kinesisStreamSpecification"`
	// Local secondary indexes to be created on the table.
	//
	// You can create up to 5 local secondary indexes. Each index is scoped to a given hash key value. The size of each hash key can be up to 10 gigabytes.
	LocalSecondaryIndexes interface{} `json:"localSecondaryIndexes" yaml:"localSecondaryIndexes"`
	// The settings used to enable point in time recovery.
	PointInTimeRecoverySpecification interface{} `json:"pointInTimeRecoverySpecification" yaml:"pointInTimeRecoverySpecification"`
	// Throughput for the specified table, which consists of values for `ReadCapacityUnits` and `WriteCapacityUnits` .
	//
	// For more information about the contents of a provisioned throughput structure, see [Amazon DynamoDB Table ProvisionedThroughput](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-provisionedthroughput.html) .
	//
	// If you set `BillingMode` as `PROVISIONED` , you must specify this property. If you set `BillingMode` as `PAY_PER_REQUEST` , you cannot specify this property.
	ProvisionedThroughput interface{} `json:"provisionedThroughput" yaml:"provisionedThroughput"`
	// Specifies the settings to enable server-side encryption.
	SseSpecification interface{} `json:"sseSpecification" yaml:"sseSpecification"`
	// The settings for the DynamoDB table stream, which capture changes to items stored in the table.
	StreamSpecification interface{} `json:"streamSpecification" yaml:"streamSpecification"`
	// The table class of the new table.
	//
	// Valid values are `STANDARD` and `STANDARD_INFREQUENT_ACCESS` .
	TableClass *string `json:"tableClass" yaml:"tableClass"`
	// A name for the table.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the table name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	TableName *string `json:"tableName" yaml:"tableName"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `json:"tags" yaml:"tags"`
	// Specifies the Time to Live (TTL) settings for the table.
	//
	// > For detailed information about the limits in DynamoDB, see [Limits in Amazon DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html) in the Amazon DynamoDB Developer Guide.
	TimeToLiveSpecification interface{} `json:"timeToLiveSpecification" yaml:"timeToLiveSpecification"`
}

// Properties for enabling DynamoDB capacity scaling.
//
// TODO: EXAMPLE
//
// Experimental.
type EnableScalingProps struct {
	// Maximum capacity to scale to.
	// Experimental.
	MaxCapacity *float64 `json:"maxCapacity" yaml:"maxCapacity"`
	// Minimum capacity to scale to.
	// Experimental.
	MinCapacity *float64 `json:"minCapacity" yaml:"minCapacity"`
}

// Properties for a global secondary index.
//
// TODO: EXAMPLE
//
// Experimental.
type GlobalSecondaryIndexProps struct {
	// The name of the secondary index.
	// Experimental.
	IndexName *string `json:"indexName" yaml:"indexName"`
	// The non-key attributes that are projected into the secondary index.
	// Experimental.
	NonKeyAttributes *[]*string `json:"nonKeyAttributes" yaml:"nonKeyAttributes"`
	// The set of attributes that are projected into the secondary index.
	// Experimental.
	ProjectionType ProjectionType `json:"projectionType" yaml:"projectionType"`
	// Partition key attribute definition.
	// Experimental.
	PartitionKey *Attribute `json:"partitionKey" yaml:"partitionKey"`
	// Sort key attribute definition.
	// Experimental.
	SortKey *Attribute `json:"sortKey" yaml:"sortKey"`
	// The read capacity for the global secondary index.
	//
	// Can only be provided if table billingMode is Provisioned or undefined.
	// Experimental.
	ReadCapacity *float64 `json:"readCapacity" yaml:"readCapacity"`
	// The write capacity for the global secondary index.
	//
	// Can only be provided if table billingMode is Provisioned or undefined.
	// Experimental.
	WriteCapacity *float64 `json:"writeCapacity" yaml:"writeCapacity"`
}

// Interface for scalable attributes.
// Experimental.
type IScalableTableAttribute interface {
	// Add scheduled scaling for this scaling attribute.
	// Experimental.
	ScaleOnSchedule(id *string, actions *awsapplicationautoscaling.ScalingSchedule)
	// Scale out or in to keep utilization at a given level.
	// Experimental.
	ScaleOnUtilization(props *UtilizationScalingProps)
}

// The jsii proxy for IScalableTableAttribute
type jsiiProxy_IScalableTableAttribute struct {
	_ byte // padding
}

func (i *jsiiProxy_IScalableTableAttribute) ScaleOnSchedule(id *string, actions *awsapplicationautoscaling.ScalingSchedule) {
	_jsii_.InvokeVoid(
		i,
		"scaleOnSchedule",
		[]interface{}{id, actions},
	)
}

func (i *jsiiProxy_IScalableTableAttribute) ScaleOnUtilization(props *UtilizationScalingProps) {
	_jsii_.InvokeVoid(
		i,
		"scaleOnUtilization",
		[]interface{}{props},
	)
}

// An interface that represents a DynamoDB Table - either created with the CDK, or an existing one.
// Experimental.
type ITable interface {
	awscdk.IResource
	// Adds an IAM policy statement associated with this table to an IAM principal's policy.
	//
	// If `encryptionKey` is present, appropriate grants to the key needs to be added
	// separately using the `table.encryptionKey.grant*` methods.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Permits all DynamoDB operations ("dynamodb:*") to an IAM principal.
	//
	// Appropriate grants will also be added to the customer-managed KMS key
	// if one was configured.
	// Experimental.
	GrantFullAccess(grantee awsiam.IGrantable) awsiam.Grant
	// Permits an IAM principal all data read operations from this table: BatchGetItem, GetRecords, GetShardIterator, Query, GetItem, Scan.
	//
	// Appropriate grants will also be added to the customer-managed KMS key
	// if one was configured.
	// Experimental.
	GrantReadData(grantee awsiam.IGrantable) awsiam.Grant
	// Permits an IAM principal to all data read/write operations to this table.
	//
	// BatchGetItem, GetRecords, GetShardIterator, Query, GetItem, Scan,
	// BatchWriteItem, PutItem, UpdateItem, DeleteItem
	//
	// Appropriate grants will also be added to the customer-managed KMS key
	// if one was configured.
	// Experimental.
	GrantReadWriteData(grantee awsiam.IGrantable) awsiam.Grant
	// Adds an IAM policy statement associated with this table's stream to an IAM principal's policy.
	//
	// If `encryptionKey` is present, appropriate grants to the key needs to be added
	// separately using the `table.encryptionKey.grant*` methods.
	// Experimental.
	GrantStream(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Permits an IAM principal all stream data read operations for this table's stream: DescribeStream, GetRecords, GetShardIterator, ListStreams.
	//
	// Appropriate grants will also be added to the customer-managed KMS key
	// if one was configured.
	// Experimental.
	GrantStreamRead(grantee awsiam.IGrantable) awsiam.Grant
	// Permits an IAM Principal to list streams attached to current dynamodb table.
	// Experimental.
	GrantTableListStreams(grantee awsiam.IGrantable) awsiam.Grant
	// Permits an IAM principal all data write operations to this table: BatchWriteItem, PutItem, UpdateItem, DeleteItem.
	//
	// Appropriate grants will also be added to the customer-managed KMS key
	// if one was configured.
	// Experimental.
	GrantWriteData(grantee awsiam.IGrantable) awsiam.Grant
	// Metric for the number of Errors executing all Lambdas.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the conditional check failed requests.
	// Experimental.
	MetricConditionalCheckFailedRequests(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the consumed read capacity units.
	// Experimental.
	MetricConsumedReadCapacityUnits(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the consumed write capacity units.
	// Experimental.
	MetricConsumedWriteCapacityUnits(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the successful request latency.
	// Experimental.
	MetricSuccessfulRequestLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the system errors.
	// Deprecated: use `metricSystemErrorsForOperations`
	MetricSystemErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the system errors this table.
	// Experimental.
	MetricSystemErrorsForOperations(props *SystemErrorsForOperationsMetricOptions) awscloudwatch.IMetric
	// Metric for throttled requests.
	// Experimental.
	MetricThrottledRequests(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the user errors.
	// Experimental.
	MetricUserErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Optional KMS encryption key associated with this table.
	// Experimental.
	EncryptionKey() awskms.IKey
	// Arn of the dynamodb table.
	// Experimental.
	TableArn() *string
	// Table name of the dynamodb table.
	// Experimental.
	TableName() *string
	// ARN of the table's stream, if there is one.
	// Experimental.
	TableStreamArn() *string
}

// The jsii proxy for ITable
type jsiiProxy_ITable struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_ITable) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITable) GrantFullAccess(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantFullAccess",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITable) GrantReadData(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantReadData",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITable) GrantReadWriteData(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantReadWriteData",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITable) GrantStream(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantStream",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITable) GrantStreamRead(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantStreamRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITable) GrantTableListStreams(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantTableListStreams",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITable) GrantWriteData(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantWriteData",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITable) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITable) MetricConditionalCheckFailedRequests(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricConditionalCheckFailedRequests",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITable) MetricConsumedReadCapacityUnits(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricConsumedReadCapacityUnits",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITable) MetricConsumedWriteCapacityUnits(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricConsumedWriteCapacityUnits",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITable) MetricSuccessfulRequestLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSuccessfulRequestLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITable) MetricSystemErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSystemErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITable) MetricSystemErrorsForOperations(props *SystemErrorsForOperationsMetricOptions) awscloudwatch.IMetric {
	var returns awscloudwatch.IMetric

	_jsii_.Invoke(
		i,
		"metricSystemErrorsForOperations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITable) MetricThrottledRequests(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricThrottledRequests",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITable) MetricUserErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricUserErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ITable) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITable) TableArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITable) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITable) TableStreamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableStreamArn",
		&returns,
	)
	return returns
}

// Properties for a local secondary index.
//
// TODO: EXAMPLE
//
// Experimental.
type LocalSecondaryIndexProps struct {
	// The name of the secondary index.
	// Experimental.
	IndexName *string `json:"indexName" yaml:"indexName"`
	// The non-key attributes that are projected into the secondary index.
	// Experimental.
	NonKeyAttributes *[]*string `json:"nonKeyAttributes" yaml:"nonKeyAttributes"`
	// The set of attributes that are projected into the secondary index.
	// Experimental.
	ProjectionType ProjectionType `json:"projectionType" yaml:"projectionType"`
	// The attribute of a sort key for the local secondary index.
	// Experimental.
	SortKey *Attribute `json:"sortKey" yaml:"sortKey"`
}

// Supported DynamoDB table operations.
// Experimental.
type Operation string

const (
	Operation_GET_ITEM Operation = "GET_ITEM"
	Operation_BATCH_GET_ITEM Operation = "BATCH_GET_ITEM"
	Operation_SCAN Operation = "SCAN"
	Operation_QUERY Operation = "QUERY"
	Operation_GET_RECORDS Operation = "GET_RECORDS"
	Operation_PUT_ITEM Operation = "PUT_ITEM"
	Operation_DELETE_ITEM Operation = "DELETE_ITEM"
	Operation_UPDATE_ITEM Operation = "UPDATE_ITEM"
	Operation_BATCH_WRITE_ITEM Operation = "BATCH_WRITE_ITEM"
	Operation_TRANSACT_WRITE_ITEMS Operation = "TRANSACT_WRITE_ITEMS"
	Operation_TRANSACT_GET_ITEMS Operation = "TRANSACT_GET_ITEMS"
	Operation_EXECUTE_TRANSACTION Operation = "EXECUTE_TRANSACTION"
	Operation_BATCH_EXECUTE_STATEMENT Operation = "BATCH_EXECUTE_STATEMENT"
	Operation_EXECUTE_STATEMENT Operation = "EXECUTE_STATEMENT"
)

// The set of attributes that are projected into the index.
// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_Projection.html
//
// Experimental.
type ProjectionType string

const (
	ProjectionType_KEYS_ONLY ProjectionType = "KEYS_ONLY"
	ProjectionType_INCLUDE ProjectionType = "INCLUDE"
	ProjectionType_ALL ProjectionType = "ALL"
)

// Represents the table schema attributes.
//
// TODO: EXAMPLE
//
// Experimental.
type SchemaOptions struct {
	// Partition key attribute definition.
	// Experimental.
	PartitionKey *Attribute `json:"partitionKey" yaml:"partitionKey"`
	// Sort key attribute definition.
	// Experimental.
	SortKey *Attribute `json:"sortKey" yaml:"sortKey"`
}

// Properties for a secondary index.
//
// TODO: EXAMPLE
//
// Experimental.
type SecondaryIndexProps struct {
	// The name of the secondary index.
	// Experimental.
	IndexName *string `json:"indexName" yaml:"indexName"`
	// The non-key attributes that are projected into the secondary index.
	// Experimental.
	NonKeyAttributes *[]*string `json:"nonKeyAttributes" yaml:"nonKeyAttributes"`
	// The set of attributes that are projected into the secondary index.
	// Experimental.
	ProjectionType ProjectionType `json:"projectionType" yaml:"projectionType"`
}

// When an item in the table is modified, StreamViewType determines what information is written to the stream for this table.
// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_StreamSpecification.html
//
// Experimental.
type StreamViewType string

const (
	StreamViewType_NEW_IMAGE StreamViewType = "NEW_IMAGE"
	StreamViewType_OLD_IMAGE StreamViewType = "OLD_IMAGE"
	StreamViewType_NEW_AND_OLD_IMAGES StreamViewType = "NEW_AND_OLD_IMAGES"
	StreamViewType_KEYS_ONLY StreamViewType = "KEYS_ONLY"
)

// Options for configuring a system errors metric that considers multiple operations.
//
// TODO: EXAMPLE
//
// Experimental.
type SystemErrorsForOperationsMetricOptions struct {
	// Account which this metric comes from.
	// Experimental.
	Account *string `json:"account" yaml:"account"`
	// The hex color code, prefixed with '#' (e.g. '#00ff00'), to use when this metric is rendered on a graph. The `Color` class has a set of standard colors that can be used here.
	// Experimental.
	Color *string `json:"color" yaml:"color"`
	// Dimensions of the metric.
	// Deprecated: Use 'dimensionsMap' instead.
	Dimensions *map[string]interface{} `json:"dimensions" yaml:"dimensions"`
	// Dimensions of the metric.
	// Experimental.
	DimensionsMap *map[string]*string `json:"dimensionsMap" yaml:"dimensionsMap"`
	// Label for this metric when added to a Graph in a Dashboard.
	// Experimental.
	Label *string `json:"label" yaml:"label"`
	// The period over which the specified statistic is applied.
	// Experimental.
	Period awscdk.Duration `json:"period" yaml:"period"`
	// Region which this metric comes from.
	// Experimental.
	Region *string `json:"region" yaml:"region"`
	// What function to use for aggregating.
	//
	// Can be one of the following:
	//
	// - "Minimum" | "min"
	// - "Maximum" | "max"
	// - "Average" | "avg"
	// - "Sum" | "sum"
	// - "SampleCount | "n"
	// - "pNN.NN"
	// Experimental.
	Statistic *string `json:"statistic" yaml:"statistic"`
	// Unit used to filter the metric stream.
	//
	// Only refer to datums emitted to the metric stream with the given unit and
	// ignore all others. Only useful when datums are being emitted to the same
	// metric stream under different units.
	//
	// The default is to use all matric datums in the stream, regardless of unit,
	// which is recommended in nearly all cases.
	//
	// CloudWatch does not honor this property for graphs.
	// Experimental.
	Unit awscloudwatch.Unit `json:"unit" yaml:"unit"`
	// The operations to apply the metric to.
	// Experimental.
	Operations *[]Operation `json:"operations" yaml:"operations"`
}

// Provides a DynamoDB table.
//
// TODO: EXAMPLE
//
// Experimental.
type Table interface {
	awscdk.Resource
	ITable
	EncryptionKey() awskms.IKey
	Env() *awscdk.ResourceEnvironment
	HasIndex() *bool
	Node() awscdk.ConstructNode
	PhysicalName() *string
	RegionalArns() *[]*string
	Stack() awscdk.Stack
	TableArn() *string
	TableName() *string
	TableStreamArn() *string
	AddGlobalSecondaryIndex(props *GlobalSecondaryIndexProps)
	AddLocalSecondaryIndex(props *LocalSecondaryIndexProps)
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	AutoScaleGlobalSecondaryIndexReadCapacity(indexName *string, props *EnableScalingProps) IScalableTableAttribute
	AutoScaleGlobalSecondaryIndexWriteCapacity(indexName *string, props *EnableScalingProps) IScalableTableAttribute
	AutoScaleReadCapacity(props *EnableScalingProps) IScalableTableAttribute
	AutoScaleWriteCapacity(props *EnableScalingProps) IScalableTableAttribute
	GeneratePhysicalName() *string
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	GetResourceNameAttribute(nameAttr *string) *string
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	GrantFullAccess(grantee awsiam.IGrantable) awsiam.Grant
	GrantReadData(grantee awsiam.IGrantable) awsiam.Grant
	GrantReadWriteData(grantee awsiam.IGrantable) awsiam.Grant
	GrantStream(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	GrantStreamRead(grantee awsiam.IGrantable) awsiam.Grant
	GrantTableListStreams(grantee awsiam.IGrantable) awsiam.Grant
	GrantWriteData(grantee awsiam.IGrantable) awsiam.Grant
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricConditionalCheckFailedRequests(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricConsumedReadCapacityUnits(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricConsumedWriteCapacityUnits(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSuccessfulRequestLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSystemErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSystemErrorsForOperations(props *SystemErrorsForOperationsMetricOptions) awscloudwatch.IMetric
	MetricThrottledRequests(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricThrottledRequestsForOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricUserErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	Schema(indexName *string) *SchemaOptions
	Synthesize(session awscdk.ISynthesisSession)
	ToString() *string
	Validate() *[]*string
}

// The jsii proxy struct for Table
type jsiiProxy_Table struct {
	internal.Type__awscdkResource
	jsiiProxy_ITable
}

func (j *jsiiProxy_Table) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Table) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Table) HasIndex() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"hasIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Table) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Table) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Table) RegionalArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionalArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Table) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Table) TableArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Table) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Table) TableStreamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableStreamArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewTable(scope constructs.Construct, id *string, props *TableProps) Table {
	_init_.Initialize()

	j := jsiiProxy_Table{}

	_jsii_.Create(
		"monocdk.aws_dynamodb.Table",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewTable_Override(t Table, scope constructs.Construct, id *string, props *TableProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_dynamodb.Table",
		[]interface{}{scope, id, props},
		t,
	)
}

// Creates a Table construct that represents an external table via table arn.
// Experimental.
func Table_FromTableArn(scope constructs.Construct, id *string, tableArn *string) ITable {
	_init_.Initialize()

	var returns ITable

	_jsii_.StaticInvoke(
		"monocdk.aws_dynamodb.Table",
		"fromTableArn",
		[]interface{}{scope, id, tableArn},
		&returns,
	)

	return returns
}

// Creates a Table construct that represents an external table.
// Experimental.
func Table_FromTableAttributes(scope constructs.Construct, id *string, attrs *TableAttributes) ITable {
	_init_.Initialize()

	var returns ITable

	_jsii_.StaticInvoke(
		"monocdk.aws_dynamodb.Table",
		"fromTableAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Creates a Table construct that represents an external table via table name.
// Experimental.
func Table_FromTableName(scope constructs.Construct, id *string, tableName *string) ITable {
	_init_.Initialize()

	var returns ITable

	_jsii_.StaticInvoke(
		"monocdk.aws_dynamodb.Table",
		"fromTableName",
		[]interface{}{scope, id, tableName},
		&returns,
	)

	return returns
}

// Permits an IAM Principal to list all DynamoDB Streams.
// Deprecated: Use {@link #grantTableListStreams} for more granular permission
func Table_GrantListStreams(grantee awsiam.IGrantable) awsiam.Grant {
	_init_.Initialize()

	var returns awsiam.Grant

	_jsii_.StaticInvoke(
		"monocdk.aws_dynamodb.Table",
		"grantListStreams",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func Table_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_dynamodb.Table",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Table_IsResource(construct awscdk.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_dynamodb.Table",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Add a global secondary index of table.
// Experimental.
func (t *jsiiProxy_Table) AddGlobalSecondaryIndex(props *GlobalSecondaryIndexProps) {
	_jsii_.InvokeVoid(
		t,
		"addGlobalSecondaryIndex",
		[]interface{}{props},
	)
}

// Add a local secondary index of table.
// Experimental.
func (t *jsiiProxy_Table) AddLocalSecondaryIndex(props *LocalSecondaryIndexProps) {
	_jsii_.InvokeVoid(
		t,
		"addLocalSecondaryIndex",
		[]interface{}{props},
	)
}

// Apply the given removal policy to this resource.
//
// The Removal Policy controls what happens to this resource when it stops
// being managed by CloudFormation, either because you've removed it from the
// CDK application or because you've made a change that requires the resource
// to be replaced.
//
// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
// Experimental.
func (t *jsiiProxy_Table) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	_jsii_.InvokeVoid(
		t,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

// Enable read capacity scaling for the given GSI.
//
// Returns: An object to configure additional AutoScaling settings for this attribute
// Experimental.
func (t *jsiiProxy_Table) AutoScaleGlobalSecondaryIndexReadCapacity(indexName *string, props *EnableScalingProps) IScalableTableAttribute {
	var returns IScalableTableAttribute

	_jsii_.Invoke(
		t,
		"autoScaleGlobalSecondaryIndexReadCapacity",
		[]interface{}{indexName, props},
		&returns,
	)

	return returns
}

// Enable write capacity scaling for the given GSI.
//
// Returns: An object to configure additional AutoScaling settings for this attribute
// Experimental.
func (t *jsiiProxy_Table) AutoScaleGlobalSecondaryIndexWriteCapacity(indexName *string, props *EnableScalingProps) IScalableTableAttribute {
	var returns IScalableTableAttribute

	_jsii_.Invoke(
		t,
		"autoScaleGlobalSecondaryIndexWriteCapacity",
		[]interface{}{indexName, props},
		&returns,
	)

	return returns
}

// Enable read capacity scaling for this table.
//
// Returns: An object to configure additional AutoScaling settings
// Experimental.
func (t *jsiiProxy_Table) AutoScaleReadCapacity(props *EnableScalingProps) IScalableTableAttribute {
	var returns IScalableTableAttribute

	_jsii_.Invoke(
		t,
		"autoScaleReadCapacity",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Enable write capacity scaling for this table.
//
// Returns: An object to configure additional AutoScaling settings for this attribute
// Experimental.
func (t *jsiiProxy_Table) AutoScaleWriteCapacity(props *EnableScalingProps) IScalableTableAttribute {
	var returns IScalableTableAttribute

	_jsii_.Invoke(
		t,
		"autoScaleWriteCapacity",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Experimental.
func (t *jsiiProxy_Table) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
//
// Normally, this token will resolve to `arnAttr`, but if the resource is
// referenced across environments, `arnComponents` will be used to synthesize
// a concrete ARN with the resource's physical name. Make sure to reference
// `this.physicalName` in `arnComponents`.
// Experimental.
func (t *jsiiProxy_Table) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
//
// Normally, this token will resolve to `nameAttr`, but if the resource is
// referenced across environments, it will be resolved to `this.physicalName`,
// which will be a concrete name.
// Experimental.
func (t *jsiiProxy_Table) GetResourceNameAttribute(nameAttr *string) *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

// Adds an IAM policy statement associated with this table to an IAM principal's policy.
//
// If `encryptionKey` is present, appropriate grants to the key needs to be added
// separately using the `table.encryptionKey.grant*` methods.
// Experimental.
func (t *jsiiProxy_Table) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		t,
		"grant",
		args,
		&returns,
	)

	return returns
}

// Permits all DynamoDB operations ("dynamodb:*") to an IAM principal.
//
// Appropriate grants will also be added to the customer-managed KMS key
// if one was configured.
// Experimental.
func (t *jsiiProxy_Table) GrantFullAccess(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		t,
		"grantFullAccess",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Permits an IAM principal all data read operations from this table: BatchGetItem, GetRecords, GetShardIterator, Query, GetItem, Scan.
//
// Appropriate grants will also be added to the customer-managed KMS key
// if one was configured.
// Experimental.
func (t *jsiiProxy_Table) GrantReadData(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		t,
		"grantReadData",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Permits an IAM principal to all data read/write operations to this table.
//
// BatchGetItem, GetRecords, GetShardIterator, Query, GetItem, Scan,
// BatchWriteItem, PutItem, UpdateItem, DeleteItem
//
// Appropriate grants will also be added to the customer-managed KMS key
// if one was configured.
// Experimental.
func (t *jsiiProxy_Table) GrantReadWriteData(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		t,
		"grantReadWriteData",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Adds an IAM policy statement associated with this table's stream to an IAM principal's policy.
//
// If `encryptionKey` is present, appropriate grants to the key needs to be added
// separately using the `table.encryptionKey.grant*` methods.
// Experimental.
func (t *jsiiProxy_Table) GrantStream(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		t,
		"grantStream",
		args,
		&returns,
	)

	return returns
}

// Permits an IAM principal all stream data read operations for this table's stream: DescribeStream, GetRecords, GetShardIterator, ListStreams.
//
// Appropriate grants will also be added to the customer-managed KMS key
// if one was configured.
// Experimental.
func (t *jsiiProxy_Table) GrantStreamRead(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		t,
		"grantStreamRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Permits an IAM Principal to list streams attached to current dynamodb table.
// Experimental.
func (t *jsiiProxy_Table) GrantTableListStreams(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		t,
		"grantTableListStreams",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Permits an IAM principal all data write operations to this table: BatchWriteItem, PutItem, UpdateItem, DeleteItem.
//
// Appropriate grants will also be added to the customer-managed KMS key
// if one was configured.
// Experimental.
func (t *jsiiProxy_Table) GrantWriteData(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		t,
		"grantWriteData",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

// Return the given named metric for this Table.
//
// By default, the metric will be calculated as a sum over a period of 5 minutes.
// You can customize this by using the `statistic` and `period` properties.
// Experimental.
func (t *jsiiProxy_Table) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the conditional check failed requests this table.
//
// By default, the metric will be calculated as a sum over a period of 5 minutes.
// You can customize this by using the `statistic` and `period` properties.
// Experimental.
func (t *jsiiProxy_Table) MetricConditionalCheckFailedRequests(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricConditionalCheckFailedRequests",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the consumed read capacity units this table.
//
// By default, the metric will be calculated as a sum over a period of 5 minutes.
// You can customize this by using the `statistic` and `period` properties.
// Experimental.
func (t *jsiiProxy_Table) MetricConsumedReadCapacityUnits(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricConsumedReadCapacityUnits",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the consumed write capacity units this table.
//
// By default, the metric will be calculated as a sum over a period of 5 minutes.
// You can customize this by using the `statistic` and `period` properties.
// Experimental.
func (t *jsiiProxy_Table) MetricConsumedWriteCapacityUnits(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricConsumedWriteCapacityUnits",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the successful request latency this table.
//
// By default, the metric will be calculated as an average over a period of 5 minutes.
// You can customize this by using the `statistic` and `period` properties.
// Experimental.
func (t *jsiiProxy_Table) MetricSuccessfulRequestLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricSuccessfulRequestLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the system errors this table.
// Deprecated: use `metricSystemErrorsForOperations`.
func (t *jsiiProxy_Table) MetricSystemErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricSystemErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the system errors this table.
//
// This will sum errors across all possible operations.
// Note that by default, each individual metric will be calculated as a sum over a period of 5 minutes.
// You can customize this by using the `statistic` and `period` properties.
// Experimental.
func (t *jsiiProxy_Table) MetricSystemErrorsForOperations(props *SystemErrorsForOperationsMetricOptions) awscloudwatch.IMetric {
	var returns awscloudwatch.IMetric

	_jsii_.Invoke(
		t,
		"metricSystemErrorsForOperations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// How many requests are throttled on this table.
//
// Default: sum over 5 minutes
// Deprecated: Do not use this function. It returns an invalid metric. Use `metricThrottledRequestsForOperation` instead.
func (t *jsiiProxy_Table) MetricThrottledRequests(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricThrottledRequests",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// How many requests are throttled on this table, for the given operation.
//
// Default: sum over 5 minutes
// Experimental.
func (t *jsiiProxy_Table) MetricThrottledRequestsForOperation(operation *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricThrottledRequestsForOperation",
		[]interface{}{operation, props},
		&returns,
	)

	return returns
}

// Metric for the user errors.
//
// Note that this metric reports user errors across all
// the tables in the account and region the table resides in.
//
// By default, the metric will be calculated as a sum over a period of 5 minutes.
// You can customize this by using the `statistic` and `period` properties.
// Experimental.
func (t *jsiiProxy_Table) MetricUserErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricUserErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (t *jsiiProxy_Table) OnPrepare() {
	_jsii_.InvokeVoid(
		t,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (t *jsiiProxy_Table) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		t,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (t *jsiiProxy_Table) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (t *jsiiProxy_Table) Prepare() {
	_jsii_.InvokeVoid(
		t,
		"prepare",
		nil, // no parameters
	)
}

// Get schema attributes of table or index.
//
// Returns: Schema of table or index.
// Experimental.
func (t *jsiiProxy_Table) Schema(indexName *string) *SchemaOptions {
	var returns *SchemaOptions

	_jsii_.Invoke(
		t,
		"schema",
		[]interface{}{indexName},
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (t *jsiiProxy_Table) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		t,
		"synthesize",
		[]interface{}{session},
	)
}

// Returns a string representation of this construct.
// Experimental.
func (t *jsiiProxy_Table) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the table construct.
//
// Returns: an array of validation error message
// Experimental.
func (t *jsiiProxy_Table) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Reference to a dynamodb table.
//
// TODO: EXAMPLE
//
// Experimental.
type TableAttributes struct {
	// KMS encryption key, if this table uses a customer-managed encryption key.
	// Experimental.
	EncryptionKey awskms.IKey `json:"encryptionKey" yaml:"encryptionKey"`
	// The name of the global indexes set for this Table.
	//
	// Note that you need to set either this property,
	// or {@link localIndexes},
	// if you want methods like grantReadData()
	// to grant permissions for indexes as well as the table itself.
	// Experimental.
	GlobalIndexes *[]*string `json:"globalIndexes" yaml:"globalIndexes"`
	// The name of the local indexes set for this Table.
	//
	// Note that you need to set either this property,
	// or {@link globalIndexes},
	// if you want methods like grantReadData()
	// to grant permissions for indexes as well as the table itself.
	// Experimental.
	LocalIndexes *[]*string `json:"localIndexes" yaml:"localIndexes"`
	// The ARN of the dynamodb table.
	//
	// One of this, or {@link tableName}, is required.
	// Experimental.
	TableArn *string `json:"tableArn" yaml:"tableArn"`
	// The table name of the dynamodb table.
	//
	// One of this, or {@link tableArn}, is required.
	// Experimental.
	TableName *string `json:"tableName" yaml:"tableName"`
	// The ARN of the table's stream.
	// Experimental.
	TableStreamArn *string `json:"tableStreamArn" yaml:"tableStreamArn"`
}

// DynamoDB's table class.
//
// TODO: EXAMPLE
//
// See: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.TableClasses.html
//
// Experimental.
type TableClass string

const (
	TableClass_STANDARD TableClass = "STANDARD"
	TableClass_STANDARD_INFREQUENT_ACCESS TableClass = "STANDARD_INFREQUENT_ACCESS"
)

// What kind of server-side encryption to apply to this table.
//
// TODO: EXAMPLE
//
// Experimental.
type TableEncryption string

const (
	TableEncryption_DEFAULT TableEncryption = "DEFAULT"
	TableEncryption_CUSTOMER_MANAGED TableEncryption = "CUSTOMER_MANAGED"
	TableEncryption_AWS_MANAGED TableEncryption = "AWS_MANAGED"
)

// Properties of a DynamoDB Table.
//
// Use {@link TableProps} for all table properties
//
// TODO: EXAMPLE
//
// Experimental.
type TableOptions struct {
	// Partition key attribute definition.
	// Experimental.
	PartitionKey *Attribute `json:"partitionKey" yaml:"partitionKey"`
	// Sort key attribute definition.
	// Experimental.
	SortKey *Attribute `json:"sortKey" yaml:"sortKey"`
	// Specify how you are charged for read and write throughput and how you manage capacity.
	// Experimental.
	BillingMode BillingMode `json:"billingMode" yaml:"billingMode"`
	// Whether CloudWatch contributor insights is enabled.
	// Experimental.
	ContributorInsightsEnabled *bool `json:"contributorInsightsEnabled" yaml:"contributorInsightsEnabled"`
	// Whether server-side encryption with an AWS managed customer master key is enabled.
	//
	// This property cannot be set if `serverSideEncryption` is set.
	//
	// > **NOTE**: if you set this to `CUSTOMER_MANAGED` and `encryptionKey` is not
	// > specified, the key that the Tablet generates for you will be created with
	// > default permissions. If you are using CDKv2, these permissions will be
	// > sufficient to enable the key for use with DynamoDB tables.  If you are
	// > using CDKv1, make sure the feature flag
	// > `@aws-cdk/aws-kms:defaultKeyPolicies` is set to `true` in your `cdk.json`.
	// Experimental.
	Encryption TableEncryption `json:"encryption" yaml:"encryption"`
	// External KMS key to use for table encryption.
	//
	// This property can only be set if `encryption` is set to `TableEncryption.CUSTOMER_MANAGED`.
	// Experimental.
	EncryptionKey awskms.IKey `json:"encryptionKey" yaml:"encryptionKey"`
	// Whether point-in-time recovery is enabled.
	// Experimental.
	PointInTimeRecovery *bool `json:"pointInTimeRecovery" yaml:"pointInTimeRecovery"`
	// The read capacity for the table.
	//
	// Careful if you add Global Secondary Indexes, as
	// those will share the table's provisioned throughput.
	//
	// Can only be provided if billingMode is Provisioned.
	// Experimental.
	ReadCapacity *float64 `json:"readCapacity" yaml:"readCapacity"`
	// The removal policy to apply to the DynamoDB Table.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy" yaml:"removalPolicy"`
	// Regions where replica tables will be created.
	// Experimental.
	ReplicationRegions *[]*string `json:"replicationRegions" yaml:"replicationRegions"`
	// The timeout for a table replication operation in a single region.
	// Experimental.
	ReplicationTimeout awscdk.Duration `json:"replicationTimeout" yaml:"replicationTimeout"`
	// Whether server-side encryption with an AWS managed customer master key is enabled.
	//
	// This property cannot be set if `encryption` and/or `encryptionKey` is set.
	// Deprecated: This property is deprecated. In order to obtain the same behavior as
	// enabling this, set the `encryption` property to `TableEncryption.AWS_MANAGED` instead.
	ServerSideEncryption *bool `json:"serverSideEncryption" yaml:"serverSideEncryption"`
	// When an item in the table is modified, StreamViewType determines what information is written to the stream for this table.
	// Experimental.
	Stream StreamViewType `json:"stream" yaml:"stream"`
	// Specify the table class.
	// Experimental.
	TableClass TableClass `json:"tableClass" yaml:"tableClass"`
	// The name of TTL attribute.
	// Experimental.
	TimeToLiveAttribute *string `json:"timeToLiveAttribute" yaml:"timeToLiveAttribute"`
	// Indicates whether CloudFormation stack waits for replication to finish.
	//
	// If set to false, the CloudFormation resource will mark the resource as
	// created and replication will be completed asynchronously. This property is
	// ignored if replicationRegions property is not set.
	//
	// DO NOT UNSET this property if adding/removing multiple replicationRegions
	// in one deployment, as CloudFormation only supports one region replication
	// at a time. CDK overcomes this limitation by waiting for replication to
	// finish before starting new replicationRegion.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-replicas
	//
	// Experimental.
	WaitForReplicationToFinish *bool `json:"waitForReplicationToFinish" yaml:"waitForReplicationToFinish"`
	// The write capacity for the table.
	//
	// Careful if you add Global Secondary Indexes, as
	// those will share the table's provisioned throughput.
	//
	// Can only be provided if billingMode is Provisioned.
	// Experimental.
	WriteCapacity *float64 `json:"writeCapacity" yaml:"writeCapacity"`
}

// Properties for a DynamoDB Table.
//
// TODO: EXAMPLE
//
// Experimental.
type TableProps struct {
	// Partition key attribute definition.
	// Experimental.
	PartitionKey *Attribute `json:"partitionKey" yaml:"partitionKey"`
	// Sort key attribute definition.
	// Experimental.
	SortKey *Attribute `json:"sortKey" yaml:"sortKey"`
	// Specify how you are charged for read and write throughput and how you manage capacity.
	// Experimental.
	BillingMode BillingMode `json:"billingMode" yaml:"billingMode"`
	// Whether CloudWatch contributor insights is enabled.
	// Experimental.
	ContributorInsightsEnabled *bool `json:"contributorInsightsEnabled" yaml:"contributorInsightsEnabled"`
	// Whether server-side encryption with an AWS managed customer master key is enabled.
	//
	// This property cannot be set if `serverSideEncryption` is set.
	//
	// > **NOTE**: if you set this to `CUSTOMER_MANAGED` and `encryptionKey` is not
	// > specified, the key that the Tablet generates for you will be created with
	// > default permissions. If you are using CDKv2, these permissions will be
	// > sufficient to enable the key for use with DynamoDB tables.  If you are
	// > using CDKv1, make sure the feature flag
	// > `@aws-cdk/aws-kms:defaultKeyPolicies` is set to `true` in your `cdk.json`.
	// Experimental.
	Encryption TableEncryption `json:"encryption" yaml:"encryption"`
	// External KMS key to use for table encryption.
	//
	// This property can only be set if `encryption` is set to `TableEncryption.CUSTOMER_MANAGED`.
	// Experimental.
	EncryptionKey awskms.IKey `json:"encryptionKey" yaml:"encryptionKey"`
	// Whether point-in-time recovery is enabled.
	// Experimental.
	PointInTimeRecovery *bool `json:"pointInTimeRecovery" yaml:"pointInTimeRecovery"`
	// The read capacity for the table.
	//
	// Careful if you add Global Secondary Indexes, as
	// those will share the table's provisioned throughput.
	//
	// Can only be provided if billingMode is Provisioned.
	// Experimental.
	ReadCapacity *float64 `json:"readCapacity" yaml:"readCapacity"`
	// The removal policy to apply to the DynamoDB Table.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy" yaml:"removalPolicy"`
	// Regions where replica tables will be created.
	// Experimental.
	ReplicationRegions *[]*string `json:"replicationRegions" yaml:"replicationRegions"`
	// The timeout for a table replication operation in a single region.
	// Experimental.
	ReplicationTimeout awscdk.Duration `json:"replicationTimeout" yaml:"replicationTimeout"`
	// Whether server-side encryption with an AWS managed customer master key is enabled.
	//
	// This property cannot be set if `encryption` and/or `encryptionKey` is set.
	// Deprecated: This property is deprecated. In order to obtain the same behavior as
	// enabling this, set the `encryption` property to `TableEncryption.AWS_MANAGED` instead.
	ServerSideEncryption *bool `json:"serverSideEncryption" yaml:"serverSideEncryption"`
	// When an item in the table is modified, StreamViewType determines what information is written to the stream for this table.
	// Experimental.
	Stream StreamViewType `json:"stream" yaml:"stream"`
	// Specify the table class.
	// Experimental.
	TableClass TableClass `json:"tableClass" yaml:"tableClass"`
	// The name of TTL attribute.
	// Experimental.
	TimeToLiveAttribute *string `json:"timeToLiveAttribute" yaml:"timeToLiveAttribute"`
	// Indicates whether CloudFormation stack waits for replication to finish.
	//
	// If set to false, the CloudFormation resource will mark the resource as
	// created and replication will be completed asynchronously. This property is
	// ignored if replicationRegions property is not set.
	//
	// DO NOT UNSET this property if adding/removing multiple replicationRegions
	// in one deployment, as CloudFormation only supports one region replication
	// at a time. CDK overcomes this limitation by waiting for replication to
	// finish before starting new replicationRegion.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-replicas
	//
	// Experimental.
	WaitForReplicationToFinish *bool `json:"waitForReplicationToFinish" yaml:"waitForReplicationToFinish"`
	// The write capacity for the table.
	//
	// Careful if you add Global Secondary Indexes, as
	// those will share the table's provisioned throughput.
	//
	// Can only be provided if billingMode is Provisioned.
	// Experimental.
	WriteCapacity *float64 `json:"writeCapacity" yaml:"writeCapacity"`
	// Kinesis Data Stream to capture item-level changes for the table.
	// Experimental.
	KinesisStream awskinesis.IStream `json:"kinesisStream" yaml:"kinesisStream"`
	// Enforces a particular physical table name.
	// Experimental.
	TableName *string `json:"tableName" yaml:"tableName"`
}

// Properties for enabling DynamoDB utilization tracking.
//
// TODO: EXAMPLE
//
// Experimental.
type UtilizationScalingProps struct {
	// Indicates whether scale in by the target tracking policy is disabled.
	//
	// If the value is true, scale in is disabled and the target tracking policy
	// won't remove capacity from the scalable resource. Otherwise, scale in is
	// enabled and the target tracking policy can remove capacity from the
	// scalable resource.
	// Experimental.
	DisableScaleIn *bool `json:"disableScaleIn" yaml:"disableScaleIn"`
	// A name for the scaling policy.
	// Experimental.
	PolicyName *string `json:"policyName" yaml:"policyName"`
	// Period after a scale in activity completes before another scale in activity can start.
	// Experimental.
	ScaleInCooldown awscdk.Duration `json:"scaleInCooldown" yaml:"scaleInCooldown"`
	// Period after a scale out activity completes before another scale out activity can start.
	// Experimental.
	ScaleOutCooldown awscdk.Duration `json:"scaleOutCooldown" yaml:"scaleOutCooldown"`
	// Target utilization percentage for the attribute.
	// Experimental.
	TargetUtilizationPercent *float64 `json:"targetUtilizationPercent" yaml:"targetUtilizationPercent"`
}

