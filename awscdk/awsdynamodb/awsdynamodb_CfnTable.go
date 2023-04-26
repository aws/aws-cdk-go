package awsdynamodb

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsdynamodb/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::DynamoDB::Table`.
//
// The `AWS::DynamoDB::Table` resource creates a DynamoDB table. For more information, see [CreateTable](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_CreateTable.html) in the *Amazon DynamoDB API Reference* .
//
// You should be aware of the following behaviors when working with DynamoDB tables:
//
// - AWS CloudFormation typically creates DynamoDB tables in parallel. However, if your template includes multiple DynamoDB tables with indexes, you must declare dependencies so that the tables are created sequentially. Amazon DynamoDB limits the number of tables with secondary indexes that are in the creating state. If you create multiple tables with indexes at the same time, DynamoDB returns an error and the stack operation fails. For an example, see [DynamoDB Table with a DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#aws-resource-dynamodb-table--examples--DynamoDB_Table_with_a_DependsOn_Attribute) .
//
// > Our guidance is to use the latest schema documented here for your AWS CloudFormation templates. This schema supports the provisioning of all table settings below. When using this schema in your AWS CloudFormation templates, please ensure that your Identity and Access Management ( IAM ) policies are updated with appropriate permissions to allow for the authorization of these setting changes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTable := awscdk.Aws_dynamodb.NewCfnTable(this, jsii.String("MyCfnTable"), &CfnTableProps{
//   	KeySchema: []interface{}{
//   		&KeySchemaProperty{
//   			AttributeName: jsii.String("attributeName"),
//   			KeyType: jsii.String("keyType"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	AttributeDefinitions: []interface{}{
//   		&AttributeDefinitionProperty{
//   			AttributeName: jsii.String("attributeName"),
//   			AttributeType: jsii.String("attributeType"),
//   		},
//   	},
//   	BillingMode: jsii.String("billingMode"),
//   	ContributorInsightsSpecification: &ContributorInsightsSpecificationProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	DeletionProtectionEnabled: jsii.Boolean(false),
//   	GlobalSecondaryIndexes: []interface{}{
//   		&GlobalSecondaryIndexProperty{
//   			IndexName: jsii.String("indexName"),
//   			KeySchema: []interface{}{
//   				&KeySchemaProperty{
//   					AttributeName: jsii.String("attributeName"),
//   					KeyType: jsii.String("keyType"),
//   				},
//   			},
//   			Projection: &ProjectionProperty{
//   				NonKeyAttributes: []*string{
//   					jsii.String("nonKeyAttributes"),
//   				},
//   				ProjectionType: jsii.String("projectionType"),
//   			},
//
//   			// the properties below are optional
//   			ContributorInsightsSpecification: &ContributorInsightsSpecificationProperty{
//   				Enabled: jsii.Boolean(false),
//   			},
//   			ProvisionedThroughput: &ProvisionedThroughputProperty{
//   				ReadCapacityUnits: jsii.Number(123),
//   				WriteCapacityUnits: jsii.Number(123),
//   			},
//   		},
//   	},
//   	ImportSourceSpecification: &ImportSourceSpecificationProperty{
//   		InputFormat: jsii.String("inputFormat"),
//   		S3BucketSource: &S3BucketSourceProperty{
//   			S3Bucket: jsii.String("s3Bucket"),
//
//   			// the properties below are optional
//   			S3BucketOwner: jsii.String("s3BucketOwner"),
//   			S3KeyPrefix: jsii.String("s3KeyPrefix"),
//   		},
//
//   		// the properties below are optional
//   		InputCompressionType: jsii.String("inputCompressionType"),
//   		InputFormatOptions: &InputFormatOptionsProperty{
//   			Csv: &CsvProperty{
//   				Delimiter: jsii.String("delimiter"),
//   				HeaderList: []*string{
//   					jsii.String("headerList"),
//   				},
//   			},
//   		},
//   	},
//   	KinesisStreamSpecification: &KinesisStreamSpecificationProperty{
//   		StreamArn: jsii.String("streamArn"),
//   	},
//   	LocalSecondaryIndexes: []interface{}{
//   		&LocalSecondaryIndexProperty{
//   			IndexName: jsii.String("indexName"),
//   			KeySchema: []interface{}{
//   				&KeySchemaProperty{
//   					AttributeName: jsii.String("attributeName"),
//   					KeyType: jsii.String("keyType"),
//   				},
//   			},
//   			Projection: &ProjectionProperty{
//   				NonKeyAttributes: []*string{
//   					jsii.String("nonKeyAttributes"),
//   				},
//   				ProjectionType: jsii.String("projectionType"),
//   			},
//   		},
//   	},
//   	PointInTimeRecoverySpecification: &PointInTimeRecoverySpecificationProperty{
//   		PointInTimeRecoveryEnabled: jsii.Boolean(false),
//   	},
//   	ProvisionedThroughput: &ProvisionedThroughputProperty{
//   		ReadCapacityUnits: jsii.Number(123),
//   		WriteCapacityUnits: jsii.Number(123),
//   	},
//   	SseSpecification: &SSESpecificationProperty{
//   		SseEnabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		KmsMasterKeyId: jsii.String("kmsMasterKeyId"),
//   		SseType: jsii.String("sseType"),
//   	},
//   	StreamSpecification: &StreamSpecificationProperty{
//   		StreamViewType: jsii.String("streamViewType"),
//   	},
//   	TableClass: jsii.String("tableClass"),
//   	TableName: jsii.String("tableName"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TimeToLiveSpecification: &TimeToLiveSpecificationProperty{
//   		AttributeName: jsii.String("attributeName"),
//   		Enabled: jsii.Boolean(false),
//   	},
//   })
//
type CfnTable interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the DynamoDB table, such as `arn:aws:dynamodb:us-east-2:123456789012:table/myDynamoDBTable` .
	AttrArn() *string
	// A list of attributes that describe the key schema for the table and indexes.
	//
	// This property is required to create a DynamoDB table.
	//
	// Update requires: [Some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt) . Replacement if you edit an existing AttributeDefinition.
	AttributeDefinitions() interface{}
	SetAttributeDefinitions(val interface{})
	// The ARN of the DynamoDB stream, such as `arn:aws:dynamodb:us-east-1:123456789012:table/testddbstack-myDynamoDBTable-012A1SL7SMP5Q/stream/2015-11-30T20:10:00.000` .
	//
	// > You must specify the `StreamSpecification` property to use this attribute.
	AttrStreamArn() *string
	// Specify how you are charged for read and write throughput and how you manage capacity.
	//
	// Valid values include:
	//
	// - `PROVISIONED` - We recommend using `PROVISIONED` for predictable workloads. `PROVISIONED` sets the billing mode to [Provisioned Mode](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.ProvisionedThroughput.Manual) .
	// - `PAY_PER_REQUEST` - We recommend using `PAY_PER_REQUEST` for unpredictable workloads. `PAY_PER_REQUEST` sets the billing mode to [On-Demand Mode](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.OnDemand) .
	//
	// If not specified, the default is `PROVISIONED` .
	BillingMode() *string
	SetBillingMode(val *string)
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// The settings used to enable or disable CloudWatch Contributor Insights for the specified table.
	ContributorInsightsSpecification() interface{}
	SetContributorInsightsSpecification(val interface{})
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// Determines if a table is protected from deletion.
	//
	// When enabled, the table cannot be deleted by any user or process. This setting is disabled by default. For more information, see [Using deletion protection](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.Basics.html#WorkingWithTables.Basics.DeletionProtection) in the *Amazon DynamoDB Developer Guide* .
	DeletionProtectionEnabled() interface{}
	SetDeletionProtectionEnabled(val interface{})
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
	GlobalSecondaryIndexes() interface{}
	SetGlobalSecondaryIndexes(val interface{})
	// Specifies the properties of data being imported from the S3 bucket source to the table.
	//
	// > If you specify the `ImportSourceSpecification` property, and also specify either the `StreamSpecification` , the `TableClass` property, or the `DeletionProtectionEnabled` property, the IAM entity creating/updating stack must have `UpdateTable` permission.
	ImportSourceSpecification() interface{}
	SetImportSourceSpecification(val interface{})
	// Specifies the attributes that make up the primary key for the table.
	//
	// The attributes in the `KeySchema` property must also be defined in the `AttributeDefinitions` property.
	KeySchema() interface{}
	SetKeySchema(val interface{})
	// The Kinesis Data Streams configuration for the specified table.
	KinesisStreamSpecification() interface{}
	SetKinesisStreamSpecification(val interface{})
	// Local secondary indexes to be created on the table.
	//
	// You can create up to 5 local secondary indexes. Each index is scoped to a given hash key value. The size of each hash key can be up to 10 gigabytes.
	LocalSecondaryIndexes() interface{}
	SetLocalSecondaryIndexes(val interface{})
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
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The settings used to enable point in time recovery.
	PointInTimeRecoverySpecification() interface{}
	SetPointInTimeRecoverySpecification(val interface{})
	// Throughput for the specified table, which consists of values for `ReadCapacityUnits` and `WriteCapacityUnits` .
	//
	// For more information about the contents of a provisioned throughput structure, see [Amazon DynamoDB Table ProvisionedThroughput](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ProvisionedThroughput.html) .
	//
	// If you set `BillingMode` as `PROVISIONED` , you must specify this property. If you set `BillingMode` as `PAY_PER_REQUEST` , you cannot specify this property.
	ProvisionedThroughput() interface{}
	SetProvisionedThroughput(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// Specifies the settings to enable server-side encryption.
	SseSpecification() interface{}
	SetSseSpecification(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// The settings for the DynamoDB table stream, which capture changes to items stored in the table.
	StreamSpecification() interface{}
	SetStreamSpecification(val interface{})
	// The table class of the new table.
	//
	// Valid values are `STANDARD` and `STANDARD_INFREQUENT_ACCESS` .
	TableClass() *string
	SetTableClass(val *string)
	// A name for the table.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the table name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	TableName() *string
	SetTableName(val *string)
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags() awscdk.TagManager
	// Specifies the Time to Live (TTL) settings for the table.
	//
	// > For detailed information about the limits in DynamoDB, see [Limits in Amazon DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html) in the Amazon DynamoDB Developer Guide.
	TimeToLiveSpecification() interface{}
	SetTimeToLiveSpecification(val interface{})
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

func (j *jsiiProxy_CfnTable) DeletionProtectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionEnabled",
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

func (j *jsiiProxy_CfnTable) ImportSourceSpecification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"importSourceSpecification",
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

	if err := validateNewCfnTableParameters(scope, id, props); err != nil {
		panic(err)
	}
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

func (j *jsiiProxy_CfnTable)SetAttributeDefinitions(val interface{}) {
	if err := j.validateSetAttributeDefinitionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributeDefinitions",
		val,
	)
}

func (j *jsiiProxy_CfnTable)SetBillingMode(val *string) {
	_jsii_.Set(
		j,
		"billingMode",
		val,
	)
}

func (j *jsiiProxy_CfnTable)SetContributorInsightsSpecification(val interface{}) {
	if err := j.validateSetContributorInsightsSpecificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contributorInsightsSpecification",
		val,
	)
}

func (j *jsiiProxy_CfnTable)SetDeletionProtectionEnabled(val interface{}) {
	if err := j.validateSetDeletionProtectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtectionEnabled",
		val,
	)
}

func (j *jsiiProxy_CfnTable)SetGlobalSecondaryIndexes(val interface{}) {
	if err := j.validateSetGlobalSecondaryIndexesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globalSecondaryIndexes",
		val,
	)
}

func (j *jsiiProxy_CfnTable)SetImportSourceSpecification(val interface{}) {
	if err := j.validateSetImportSourceSpecificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importSourceSpecification",
		val,
	)
}

func (j *jsiiProxy_CfnTable)SetKeySchema(val interface{}) {
	if err := j.validateSetKeySchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keySchema",
		val,
	)
}

func (j *jsiiProxy_CfnTable)SetKinesisStreamSpecification(val interface{}) {
	if err := j.validateSetKinesisStreamSpecificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kinesisStreamSpecification",
		val,
	)
}

func (j *jsiiProxy_CfnTable)SetLocalSecondaryIndexes(val interface{}) {
	if err := j.validateSetLocalSecondaryIndexesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localSecondaryIndexes",
		val,
	)
}

func (j *jsiiProxy_CfnTable)SetPointInTimeRecoverySpecification(val interface{}) {
	if err := j.validateSetPointInTimeRecoverySpecificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pointInTimeRecoverySpecification",
		val,
	)
}

func (j *jsiiProxy_CfnTable)SetProvisionedThroughput(val interface{}) {
	if err := j.validateSetProvisionedThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionedThroughput",
		val,
	)
}

func (j *jsiiProxy_CfnTable)SetSseSpecification(val interface{}) {
	if err := j.validateSetSseSpecificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sseSpecification",
		val,
	)
}

func (j *jsiiProxy_CfnTable)SetStreamSpecification(val interface{}) {
	if err := j.validateSetStreamSpecificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamSpecification",
		val,
	)
}

func (j *jsiiProxy_CfnTable)SetTableClass(val *string) {
	_jsii_.Set(
		j,
		"tableClass",
		val,
	)
}

func (j *jsiiProxy_CfnTable)SetTableName(val *string) {
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

func (j *jsiiProxy_CfnTable)SetTimeToLiveSpecification(val interface{}) {
	if err := j.validateSetTimeToLiveSpecificationParameters(val); err != nil {
		panic(err)
	}
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

	if err := validateCfnTable_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
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

	if err := validateCfnTable_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
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

	if err := validateCfnTable_IsConstructParameters(x); err != nil {
		panic(err)
	}
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

func (c *jsiiProxy_CfnTable) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnTable) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnTable) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnTable) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnTable) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnTable) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnTable) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnTable) GetAtt(attributeName *string) awscdk.Reference {
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

func (c *jsiiProxy_CfnTable) GetMetadata(key *string) interface{} {
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

func (c *jsiiProxy_CfnTable) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnTable) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnTable) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnTable) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnTable) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnTable) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
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

func (c *jsiiProxy_CfnTable) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CfnTable) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

