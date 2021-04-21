package awsdynamodb

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents an attribute for describing the key schema for the table and indexes.
// Experimental.
type Attribute struct {
	// The name of an attribute.
	// Experimental.
	Name *string `json:"name"`
	// The data type of an attribute.
	// Experimental.
	Type AttributeType `json:"type"`
}

// Data types for attributes within a table.
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
// Experimental.
type BillingMode string

const (
	BillingMode_PAY_PER_REQUEST BillingMode = "PAY_PER_REQUEST"
	BillingMode_PROVISIONED BillingMode = "PROVISIONED"
)

// A CloudFormation `AWS::DynamoDB::Table`.
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
	LocalSecondaryIndexes() interface{}
	SetLocalSecondaryIndexes(val interface{})
	LogicalId() *string
	Node() constructs.Node
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

func (j *jsiiProxy_CfnTable) Node() constructs.Node {
	var returns constructs.Node
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
func NewCfnTable(scope constructs.Construct, id *string, props *CfnTableProps) CfnTable {
	_init_.Initialize()

	j := jsiiProxy_CfnTable{}

	_jsii_.Create(
		"aws-cdk-lib.aws_dynamodb.CfnTable",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::DynamoDB::Table`.
func NewCfnTable_Override(c CfnTable, scope constructs.Construct, id *string, props *CfnTableProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_dynamodb.CfnTable",
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
		"aws-cdk-lib.aws_dynamodb.CfnTable",
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
		"aws-cdk-lib.aws_dynamodb.CfnTable",
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
		"aws-cdk-lib.aws_dynamodb.CfnTable",
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
		"aws-cdk-lib.aws_dynamodb.CfnTable",
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
// Experimental.
func (c *jsiiProxy_CfnTable) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
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

// Experimental.
func (c *jsiiProxy_CfnTable) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

type CfnTable_AttributeDefinitionProperty struct {
	// `CfnTable.AttributeDefinitionProperty.AttributeName`.
	AttributeName *string `json:"attributeName"`
	// `CfnTable.AttributeDefinitionProperty.AttributeType`.
	AttributeType *string `json:"attributeType"`
}

type CfnTable_ContributorInsightsSpecificationProperty struct {
	// `CfnTable.ContributorInsightsSpecificationProperty.Enabled`.
	Enabled interface{} `json:"enabled"`
}

type CfnTable_GlobalSecondaryIndexProperty struct {
	// `CfnTable.GlobalSecondaryIndexProperty.IndexName`.
	IndexName *string `json:"indexName"`
	// `CfnTable.GlobalSecondaryIndexProperty.KeySchema`.
	KeySchema interface{} `json:"keySchema"`
	// `CfnTable.GlobalSecondaryIndexProperty.Projection`.
	Projection interface{} `json:"projection"`
	// `CfnTable.GlobalSecondaryIndexProperty.ContributorInsightsSpecification`.
	ContributorInsightsSpecification interface{} `json:"contributorInsightsSpecification"`
	// `CfnTable.GlobalSecondaryIndexProperty.ProvisionedThroughput`.
	ProvisionedThroughput interface{} `json:"provisionedThroughput"`
}

type CfnTable_KeySchemaProperty struct {
	// `CfnTable.KeySchemaProperty.AttributeName`.
	AttributeName *string `json:"attributeName"`
	// `CfnTable.KeySchemaProperty.KeyType`.
	KeyType *string `json:"keyType"`
}

type CfnTable_LocalSecondaryIndexProperty struct {
	// `CfnTable.LocalSecondaryIndexProperty.IndexName`.
	IndexName *string `json:"indexName"`
	// `CfnTable.LocalSecondaryIndexProperty.KeySchema`.
	KeySchema interface{} `json:"keySchema"`
	// `CfnTable.LocalSecondaryIndexProperty.Projection`.
	Projection interface{} `json:"projection"`
}

type CfnTable_PointInTimeRecoverySpecificationProperty struct {
	// `CfnTable.PointInTimeRecoverySpecificationProperty.PointInTimeRecoveryEnabled`.
	PointInTimeRecoveryEnabled interface{} `json:"pointInTimeRecoveryEnabled"`
}

type CfnTable_ProjectionProperty struct {
	// `CfnTable.ProjectionProperty.NonKeyAttributes`.
	NonKeyAttributes *[]*string `json:"nonKeyAttributes"`
	// `CfnTable.ProjectionProperty.ProjectionType`.
	ProjectionType *string `json:"projectionType"`
}

type CfnTable_ProvisionedThroughputProperty struct {
	// `CfnTable.ProvisionedThroughputProperty.ReadCapacityUnits`.
	ReadCapacityUnits *float64 `json:"readCapacityUnits"`
	// `CfnTable.ProvisionedThroughputProperty.WriteCapacityUnits`.
	WriteCapacityUnits *float64 `json:"writeCapacityUnits"`
}

type CfnTable_SSESpecificationProperty struct {
	// `CfnTable.SSESpecificationProperty.SSEEnabled`.
	SseEnabled interface{} `json:"sseEnabled"`
	// `CfnTable.SSESpecificationProperty.KMSMasterKeyId`.
	KmsMasterKeyId *string `json:"kmsMasterKeyId"`
	// `CfnTable.SSESpecificationProperty.SSEType`.
	SseType *string `json:"sseType"`
}

type CfnTable_StreamSpecificationProperty struct {
	// `CfnTable.StreamSpecificationProperty.StreamViewType`.
	StreamViewType *string `json:"streamViewType"`
}

type CfnTable_TimeToLiveSpecificationProperty struct {
	// `CfnTable.TimeToLiveSpecificationProperty.AttributeName`.
	AttributeName *string `json:"attributeName"`
	// `CfnTable.TimeToLiveSpecificationProperty.Enabled`.
	Enabled interface{} `json:"enabled"`
}

// Properties for defining a `AWS::DynamoDB::Table`.
type CfnTableProps struct {
	// `AWS::DynamoDB::Table.KeySchema`.
	KeySchema interface{} `json:"keySchema"`
	// `AWS::DynamoDB::Table.AttributeDefinitions`.
	AttributeDefinitions interface{} `json:"attributeDefinitions"`
	// `AWS::DynamoDB::Table.BillingMode`.
	BillingMode *string `json:"billingMode"`
	// `AWS::DynamoDB::Table.ContributorInsightsSpecification`.
	ContributorInsightsSpecification interface{} `json:"contributorInsightsSpecification"`
	// `AWS::DynamoDB::Table.GlobalSecondaryIndexes`.
	GlobalSecondaryIndexes interface{} `json:"globalSecondaryIndexes"`
	// `AWS::DynamoDB::Table.LocalSecondaryIndexes`.
	LocalSecondaryIndexes interface{} `json:"localSecondaryIndexes"`
	// `AWS::DynamoDB::Table.PointInTimeRecoverySpecification`.
	PointInTimeRecoverySpecification interface{} `json:"pointInTimeRecoverySpecification"`
	// `AWS::DynamoDB::Table.ProvisionedThroughput`.
	ProvisionedThroughput interface{} `json:"provisionedThroughput"`
	// `AWS::DynamoDB::Table.SSESpecification`.
	SseSpecification interface{} `json:"sseSpecification"`
	// `AWS::DynamoDB::Table.StreamSpecification`.
	StreamSpecification interface{} `json:"streamSpecification"`
	// `AWS::DynamoDB::Table.TableName`.
	TableName *string `json:"tableName"`
	// `AWS::DynamoDB::Table.Tags`.
	Tags *[]*awscdk.CfnTag `json:"tags"`
	// `AWS::DynamoDB::Table.TimeToLiveSpecification`.
	TimeToLiveSpecification interface{} `json:"timeToLiveSpecification"`
}

// Properties for enabling DynamoDB capacity scaling.
// Experimental.
type EnableScalingProps struct {
	// Maximum capacity to scale to.
	// Experimental.
	MaxCapacity *float64 `json:"maxCapacity"`
	// Minimum capacity to scale to.
	// Experimental.
	MinCapacity *float64 `json:"minCapacity"`
}

// Properties for a global secondary index.
// Experimental.
type GlobalSecondaryIndexProps struct {
	// The name of the secondary index.
	// Experimental.
	IndexName *string `json:"indexName"`
	// The non-key attributes that are projected into the secondary index.
	// Experimental.
	NonKeyAttributes *[]*string `json:"nonKeyAttributes"`
	// The set of attributes that are projected into the secondary index.
	// Experimental.
	ProjectionType ProjectionType `json:"projectionType"`
	// The attribute of a partition key for the global secondary index.
	// Experimental.
	PartitionKey *Attribute `json:"partitionKey"`
	// The read capacity for the global secondary index.
	//
	// Can only be provided if table billingMode is Provisioned or undefined.
	// Experimental.
	ReadCapacity *float64 `json:"readCapacity"`
	// The attribute of a sort key for the global secondary index.
	// Experimental.
	SortKey *Attribute `json:"sortKey"`
	// The write capacity for the global secondary index.
	//
	// Can only be provided if table billingMode is Provisioned or undefined.
	// Experimental.
	WriteCapacity *float64 `json:"writeCapacity"`
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
// Experimental.
type LocalSecondaryIndexProps struct {
	// The name of the secondary index.
	// Experimental.
	IndexName *string `json:"indexName"`
	// The non-key attributes that are projected into the secondary index.
	// Experimental.
	NonKeyAttributes *[]*string `json:"nonKeyAttributes"`
	// The set of attributes that are projected into the secondary index.
	// Experimental.
	ProjectionType ProjectionType `json:"projectionType"`
	// The attribute of a sort key for the local secondary index.
	// Experimental.
	SortKey *Attribute `json:"sortKey"`
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

// Properties for a secondary index.
// Experimental.
type SecondaryIndexProps struct {
	// The name of the secondary index.
	// Experimental.
	IndexName *string `json:"indexName"`
	// The non-key attributes that are projected into the secondary index.
	// Experimental.
	NonKeyAttributes *[]*string `json:"nonKeyAttributes"`
	// The set of attributes that are projected into the secondary index.
	// Experimental.
	ProjectionType ProjectionType `json:"projectionType"`
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
// Experimental.
type SystemErrorsForOperationsMetricOptions struct {
	// Account which this metric comes from.
	// Experimental.
	Account *string `json:"account"`
	// The hex color code, prefixed with '#' (e.g. '#00ff00'), to use when this metric is rendered on a graph. The `Color` class has a set of standard colors that can be used here.
	// Experimental.
	Color *string `json:"color"`
	// Dimensions of the metric.
	// Experimental.
	Dimensions *map[string]interface{} `json:"dimensions"`
	// Label for this metric when added to a Graph in a Dashboard.
	// Experimental.
	Label *string `json:"label"`
	// The period over which the specified statistic is applied.
	// Experimental.
	Period awscdk.Duration `json:"period"`
	// Region which this metric comes from.
	// Experimental.
	Region *string `json:"region"`
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
	Statistic *string `json:"statistic"`
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
	Unit awscloudwatch.Unit `json:"unit"`
	// The operations to apply the metric to.
	// Experimental.
	Operations *[]Operation `json:"operations"`
}

// Provides a DynamoDB table.
// Experimental.
type Table interface {
	awscdk.Resource
	ITable
	EncryptionKey() awskms.IKey
	Env() *awscdk.ResourceEnvironment
	HasIndex() *bool
	Node() constructs.Node
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
	MetricSystemErrorsForOperations(props *SystemErrorsForOperationsMetricOptions) awscloudwatch.IMetric
	MetricThrottledRequests(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricUserErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	ToString() *string
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

func (j *jsiiProxy_Table) Node() constructs.Node {
	var returns constructs.Node
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
		"aws-cdk-lib.aws_dynamodb.Table",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewTable_Override(t Table, scope constructs.Construct, id *string, props *TableProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_dynamodb.Table",
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
		"aws-cdk-lib.aws_dynamodb.Table",
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
		"aws-cdk-lib.aws_dynamodb.Table",
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
		"aws-cdk-lib.aws_dynamodb.Table",
		"fromTableName",
		[]interface{}{scope, id, tableName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead
func Table_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dynamodb.Table",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func Table_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_dynamodb.Table",
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
// The resource can be deleted (`RemovalPolicy.DELETE`), or left in your AWS
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
// Experimental.
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

// Reference to a dynamodb table.
// Experimental.
type TableAttributes struct {
	// KMS encryption key, if this table uses a customer-managed encryption key.
	// Experimental.
	EncryptionKey awskms.IKey `json:"encryptionKey"`
	// The name of the global indexes set for this Table.
	//
	// Note that you need to set either this property,
	// or {@link localIndexes},
	// if you want methods like grantReadData()
	// to grant permissions for indexes as well as the table itself.
	// Experimental.
	GlobalIndexes *[]*string `json:"globalIndexes"`
	// The name of the local indexes set for this Table.
	//
	// Note that you need to set either this property,
	// or {@link globalIndexes},
	// if you want methods like grantReadData()
	// to grant permissions for indexes as well as the table itself.
	// Experimental.
	LocalIndexes *[]*string `json:"localIndexes"`
	// The ARN of the dynamodb table.
	//
	// One of this, or {@link tableName}, is required.
	// Experimental.
	TableArn *string `json:"tableArn"`
	// The table name of the dynamodb table.
	//
	// One of this, or {@link tableArn}, is required.
	// Experimental.
	TableName *string `json:"tableName"`
	// The ARN of the table's stream.
	// Experimental.
	TableStreamArn *string `json:"tableStreamArn"`
}

// What kind of server-side encryption to apply to this table.
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
// Experimental.
type TableOptions struct {
	// Partition key attribute definition.
	// Experimental.
	PartitionKey *Attribute `json:"partitionKey"`
	// Specify how you are charged for read and write throughput and how you manage capacity.
	// Experimental.
	BillingMode BillingMode `json:"billingMode"`
	// Whether server-side encryption with an AWS managed customer master key is enabled.
	//
	// This property cannot be set if `serverSideEncryption` is set.
	// Experimental.
	Encryption TableEncryption `json:"encryption"`
	// External KMS key to use for table encryption.
	//
	// This property can only be set if `encryption` is set to `TableEncryption.CUSTOMER_MANAGED`.
	// Experimental.
	EncryptionKey awskms.IKey `json:"encryptionKey"`
	// Whether point-in-time recovery is enabled.
	// Experimental.
	PointInTimeRecovery *bool `json:"pointInTimeRecovery"`
	// The read capacity for the table.
	//
	// Careful if you add Global Secondary Indexes, as
	// those will share the table's provisioned throughput.
	//
	// Can only be provided if billingMode is Provisioned.
	// Experimental.
	ReadCapacity *float64 `json:"readCapacity"`
	// The removal policy to apply to the DynamoDB Table.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy"`
	// Regions where replica tables will be created.
	// Experimental.
	ReplicationRegions *[]*string `json:"replicationRegions"`
	// The timeout for a table replication operation in a single region.
	// Experimental.
	ReplicationTimeout awscdk.Duration `json:"replicationTimeout"`
	// Table sort key attribute definition.
	// Experimental.
	SortKey *Attribute `json:"sortKey"`
	// When an item in the table is modified, StreamViewType determines what information is written to the stream for this table.
	// Experimental.
	Stream StreamViewType `json:"stream"`
	// The name of TTL attribute.
	// Experimental.
	TimeToLiveAttribute *string `json:"timeToLiveAttribute"`
	// The write capacity for the table.
	//
	// Careful if you add Global Secondary Indexes, as
	// those will share the table's provisioned throughput.
	//
	// Can only be provided if billingMode is Provisioned.
	// Experimental.
	WriteCapacity *float64 `json:"writeCapacity"`
}

// Properties for a DynamoDB Table.
// Experimental.
type TableProps struct {
	// Partition key attribute definition.
	// Experimental.
	PartitionKey *Attribute `json:"partitionKey"`
	// Specify how you are charged for read and write throughput and how you manage capacity.
	// Experimental.
	BillingMode BillingMode `json:"billingMode"`
	// Whether server-side encryption with an AWS managed customer master key is enabled.
	//
	// This property cannot be set if `serverSideEncryption` is set.
	// Experimental.
	Encryption TableEncryption `json:"encryption"`
	// External KMS key to use for table encryption.
	//
	// This property can only be set if `encryption` is set to `TableEncryption.CUSTOMER_MANAGED`.
	// Experimental.
	EncryptionKey awskms.IKey `json:"encryptionKey"`
	// Whether point-in-time recovery is enabled.
	// Experimental.
	PointInTimeRecovery *bool `json:"pointInTimeRecovery"`
	// The read capacity for the table.
	//
	// Careful if you add Global Secondary Indexes, as
	// those will share the table's provisioned throughput.
	//
	// Can only be provided if billingMode is Provisioned.
	// Experimental.
	ReadCapacity *float64 `json:"readCapacity"`
	// The removal policy to apply to the DynamoDB Table.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `json:"removalPolicy"`
	// Regions where replica tables will be created.
	// Experimental.
	ReplicationRegions *[]*string `json:"replicationRegions"`
	// The timeout for a table replication operation in a single region.
	// Experimental.
	ReplicationTimeout awscdk.Duration `json:"replicationTimeout"`
	// Table sort key attribute definition.
	// Experimental.
	SortKey *Attribute `json:"sortKey"`
	// When an item in the table is modified, StreamViewType determines what information is written to the stream for this table.
	// Experimental.
	Stream StreamViewType `json:"stream"`
	// The name of TTL attribute.
	// Experimental.
	TimeToLiveAttribute *string `json:"timeToLiveAttribute"`
	// The write capacity for the table.
	//
	// Careful if you add Global Secondary Indexes, as
	// those will share the table's provisioned throughput.
	//
	// Can only be provided if billingMode is Provisioned.
	// Experimental.
	WriteCapacity *float64 `json:"writeCapacity"`
	// Enforces a particular physical table name.
	// Experimental.
	TableName *string `json:"tableName"`
}

// Properties for enabling DynamoDB utilization tracking.
// Experimental.
type UtilizationScalingProps struct {
	// Indicates whether scale in by the target tracking policy is disabled.
	//
	// If the value is true, scale in is disabled and the target tracking policy
	// won't remove capacity from the scalable resource. Otherwise, scale in is
	// enabled and the target tracking policy can remove capacity from the
	// scalable resource.
	// Experimental.
	DisableScaleIn *bool `json:"disableScaleIn"`
	// A name for the scaling policy.
	// Experimental.
	PolicyName *string `json:"policyName"`
	// Period after a scale in activity completes before another scale in activity can start.
	// Experimental.
	ScaleInCooldown awscdk.Duration `json:"scaleInCooldown"`
	// Period after a scale out activity completes before another scale out activity can start.
	// Experimental.
	ScaleOutCooldown awscdk.Duration `json:"scaleOutCooldown"`
	// Target utilization percentage for the attribute.
	// Experimental.
	TargetUtilizationPercent *float64 `json:"targetUtilizationPercent"`
}

