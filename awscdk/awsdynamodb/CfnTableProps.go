package awsdynamodb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTable`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policyDocument interface{}
//
//   cfnTableProps := &CfnTableProps{
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
//   			OnDemandThroughput: &OnDemandThroughputProperty{
//   				MaxReadRequestUnits: jsii.Number(123),
//   				MaxWriteRequestUnits: jsii.Number(123),
//   			},
//   			ProvisionedThroughput: &ProvisionedThroughputProperty{
//   				ReadCapacityUnits: jsii.Number(123),
//   				WriteCapacityUnits: jsii.Number(123),
//   			},
//   			WarmThroughput: &WarmThroughputProperty{
//   				ReadUnitsPerSecond: jsii.Number(123),
//   				WriteUnitsPerSecond: jsii.Number(123),
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
//
//   		// the properties below are optional
//   		ApproximateCreationDateTimePrecision: jsii.String("approximateCreationDateTimePrecision"),
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
//   	OnDemandThroughput: &OnDemandThroughputProperty{
//   		MaxReadRequestUnits: jsii.Number(123),
//   		MaxWriteRequestUnits: jsii.Number(123),
//   	},
//   	PointInTimeRecoverySpecification: &PointInTimeRecoverySpecificationProperty{
//   		PointInTimeRecoveryEnabled: jsii.Boolean(false),
//   		RecoveryPeriodInDays: jsii.Number(123),
//   	},
//   	ProvisionedThroughput: &ProvisionedThroughputProperty{
//   		ReadCapacityUnits: jsii.Number(123),
//   		WriteCapacityUnits: jsii.Number(123),
//   	},
//   	ResourcePolicy: &ResourcePolicyProperty{
//   		PolicyDocument: policyDocument,
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
//
//   		// the properties below are optional
//   		ResourcePolicy: &ResourcePolicyProperty{
//   			PolicyDocument: policyDocument,
//   		},
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
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		AttributeName: jsii.String("attributeName"),
//   	},
//   	WarmThroughput: &WarmThroughputProperty{
//   		ReadUnitsPerSecond: jsii.Number(123),
//   		WriteUnitsPerSecond: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html
//
type CfnTableProps struct {
	// Specifies the attributes that make up the primary key for the table.
	//
	// The attributes in the `KeySchema` property must also be defined in the `AttributeDefinitions` property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-keyschema
	//
	KeySchema interface{} `field:"required" json:"keySchema" yaml:"keySchema"`
	// A list of attributes that describe the key schema for the table and indexes.
	//
	// This property is required to create a DynamoDB table.
	//
	// Update requires: [Some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt) . Replacement if you edit an existing AttributeDefinition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-attributedefinitions
	//
	AttributeDefinitions interface{} `field:"optional" json:"attributeDefinitions" yaml:"attributeDefinitions"`
	// Specify how you are charged for read and write throughput and how you manage capacity.
	//
	// Valid values include:
	//
	// - `PAY_PER_REQUEST` - We recommend using `PAY_PER_REQUEST` for most DynamoDB workloads. `PAY_PER_REQUEST` sets the billing mode to [On-demand capacity mode](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/on-demand-capacity-mode.html) .
	// - `PROVISIONED` - We recommend using `PROVISIONED` for steady workloads with predictable growth where capacity requirements can be reliably forecasted. `PROVISIONED` sets the billing mode to [Provisioned capacity mode](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/provisioned-capacity-mode.html) .
	//
	// If not specified, the default is `PROVISIONED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-billingmode
	//
	BillingMode *string `field:"optional" json:"billingMode" yaml:"billingMode"`
	// The settings used to enable or disable CloudWatch Contributor Insights for the specified table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-contributorinsightsspecification
	//
	ContributorInsightsSpecification interface{} `field:"optional" json:"contributorInsightsSpecification" yaml:"contributorInsightsSpecification"`
	// Determines if a table is protected from deletion.
	//
	// When enabled, the table cannot be deleted by any user or process. This setting is disabled by default. For more information, see [Using deletion protection](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.Basics.html#WorkingWithTables.Basics.DeletionProtection) in the *Amazon DynamoDB Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-deletionprotectionenabled
	//
	DeletionProtectionEnabled interface{} `field:"optional" json:"deletionProtectionEnabled" yaml:"deletionProtectionEnabled"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-globalsecondaryindexes
	//
	GlobalSecondaryIndexes interface{} `field:"optional" json:"globalSecondaryIndexes" yaml:"globalSecondaryIndexes"`
	// Specifies the properties of data being imported from the S3 bucket source to the" table.
	//
	// > If you specify the `ImportSourceSpecification` property, and also specify either the `StreamSpecification` , the `TableClass` property, the `DeletionProtectionEnabled` property, or the `WarmThroughput` property, the IAM entity creating/updating stack must have `UpdateTable` permission.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-importsourcespecification
	//
	ImportSourceSpecification interface{} `field:"optional" json:"importSourceSpecification" yaml:"importSourceSpecification"`
	// The Kinesis Data Streams configuration for the specified table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-kinesisstreamspecification
	//
	KinesisStreamSpecification interface{} `field:"optional" json:"kinesisStreamSpecification" yaml:"kinesisStreamSpecification"`
	// Local secondary indexes to be created on the table.
	//
	// You can create up to 5 local secondary indexes. Each index is scoped to a given hash key value. The size of each hash key can be up to 10 gigabytes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-localsecondaryindexes
	//
	LocalSecondaryIndexes interface{} `field:"optional" json:"localSecondaryIndexes" yaml:"localSecondaryIndexes"`
	// Sets the maximum number of read and write units for the specified on-demand table.
	//
	// If you use this property, you must specify `MaxReadRequestUnits` , `MaxWriteRequestUnits` , or both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-ondemandthroughput
	//
	OnDemandThroughput interface{} `field:"optional" json:"onDemandThroughput" yaml:"onDemandThroughput"`
	// The settings used to enable point in time recovery.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-pointintimerecoveryspecification
	//
	PointInTimeRecoverySpecification interface{} `field:"optional" json:"pointInTimeRecoverySpecification" yaml:"pointInTimeRecoverySpecification"`
	// Throughput for the specified table, which consists of values for `ReadCapacityUnits` and `WriteCapacityUnits` .
	//
	// For more information about the contents of a provisioned throughput structure, see [Amazon DynamoDB Table ProvisionedThroughput](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_ProvisionedThroughput.html) .
	//
	// If you set `BillingMode` as `PROVISIONED` , you must specify this property. If you set `BillingMode` as `PAY_PER_REQUEST` , you cannot specify this property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-provisionedthroughput
	//
	ProvisionedThroughput interface{} `field:"optional" json:"provisionedThroughput" yaml:"provisionedThroughput"`
	// A resource-based policy document that contains permissions to add to the specified table.
	//
	// In a CloudFormation template, you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to DynamoDB . For more information about resource-based policies, see [Using resource-based policies for DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/access-control-resource-based.html) and [Resource-based policy examples](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/rbac-examples.html) .
	//
	// When you attach a resource-based policy while creating a table, the policy creation is *strongly consistent* . For information about the considerations that you should keep in mind while attaching a resource-based policy, see [Resource-based policy considerations](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/rbac-considerations.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-resourcepolicy
	//
	ResourcePolicy interface{} `field:"optional" json:"resourcePolicy" yaml:"resourcePolicy"`
	// Specifies the settings to enable server-side encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-ssespecification
	//
	SseSpecification interface{} `field:"optional" json:"sseSpecification" yaml:"sseSpecification"`
	// The settings for the DynamoDB table stream, which capture changes to items stored in the table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-streamspecification
	//
	StreamSpecification interface{} `field:"optional" json:"streamSpecification" yaml:"streamSpecification"`
	// The table class of the new table.
	//
	// Valid values are `STANDARD` and `STANDARD_INFREQUENT_ACCESS` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-tableclass
	//
	TableClass *string `field:"optional" json:"tableClass" yaml:"tableClass"`
	// A name for the table.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the table name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-tablename
	//
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Specifies the Time to Live (TTL) settings for the table.
	//
	// > For detailed information about the limits in DynamoDB, see [Limits in Amazon DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html) in the Amazon DynamoDB Developer Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-timetolivespecification
	//
	TimeToLiveSpecification interface{} `field:"optional" json:"timeToLiveSpecification" yaml:"timeToLiveSpecification"`
	// Represents the warm throughput (in read units per second and write units per second) for creating a table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-warmthroughput
	//
	WarmThroughput interface{} `field:"optional" json:"warmThroughput" yaml:"warmThroughput"`
}

