package awsdynamodb

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnTable`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTableProps := &cfnTableProps{
//   	keySchema: []interface{}{
//   		&keySchemaProperty{
//   			attributeName: jsii.String("attributeName"),
//   			keyType: jsii.String("keyType"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	attributeDefinitions: []interface{}{
//   		&attributeDefinitionProperty{
//   			attributeName: jsii.String("attributeName"),
//   			attributeType: jsii.String("attributeType"),
//   		},
//   	},
//   	billingMode: jsii.String("billingMode"),
//   	contributorInsightsSpecification: &contributorInsightsSpecificationProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   	globalSecondaryIndexes: []interface{}{
//   		&globalSecondaryIndexProperty{
//   			indexName: jsii.String("indexName"),
//   			keySchema: []interface{}{
//   				&keySchemaProperty{
//   					attributeName: jsii.String("attributeName"),
//   					keyType: jsii.String("keyType"),
//   				},
//   			},
//   			projection: &projectionProperty{
//   				nonKeyAttributes: []*string{
//   					jsii.String("nonKeyAttributes"),
//   				},
//   				projectionType: jsii.String("projectionType"),
//   			},
//
//   			// the properties below are optional
//   			contributorInsightsSpecification: &contributorInsightsSpecificationProperty{
//   				enabled: jsii.Boolean(false),
//   			},
//   			provisionedThroughput: &provisionedThroughputProperty{
//   				readCapacityUnits: jsii.Number(123),
//   				writeCapacityUnits: jsii.Number(123),
//   			},
//   		},
//   	},
//   	importSourceSpecification: &importSourceSpecificationProperty{
//   		inputFormat: jsii.String("inputFormat"),
//   		s3BucketSource: &s3BucketSourceProperty{
//   			s3Bucket: jsii.String("s3Bucket"),
//
//   			// the properties below are optional
//   			s3BucketOwner: jsii.String("s3BucketOwner"),
//   			s3KeyPrefix: jsii.String("s3KeyPrefix"),
//   		},
//
//   		// the properties below are optional
//   		inputCompressionType: jsii.String("inputCompressionType"),
//   		inputFormatOptions: &inputFormatOptionsProperty{
//   			csv: &csvProperty{
//   				delimiter: jsii.String("delimiter"),
//   				headerList: []*string{
//   					jsii.String("headerList"),
//   				},
//   			},
//   		},
//   	},
//   	kinesisStreamSpecification: &kinesisStreamSpecificationProperty{
//   		streamArn: jsii.String("streamArn"),
//   	},
//   	localSecondaryIndexes: []interface{}{
//   		&localSecondaryIndexProperty{
//   			indexName: jsii.String("indexName"),
//   			keySchema: []interface{}{
//   				&keySchemaProperty{
//   					attributeName: jsii.String("attributeName"),
//   					keyType: jsii.String("keyType"),
//   				},
//   			},
//   			projection: &projectionProperty{
//   				nonKeyAttributes: []*string{
//   					jsii.String("nonKeyAttributes"),
//   				},
//   				projectionType: jsii.String("projectionType"),
//   			},
//   		},
//   	},
//   	pointInTimeRecoverySpecification: &pointInTimeRecoverySpecificationProperty{
//   		pointInTimeRecoveryEnabled: jsii.Boolean(false),
//   	},
//   	provisionedThroughput: &provisionedThroughputProperty{
//   		readCapacityUnits: jsii.Number(123),
//   		writeCapacityUnits: jsii.Number(123),
//   	},
//   	sseSpecification: &sSESpecificationProperty{
//   		sseEnabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		kmsMasterKeyId: jsii.String("kmsMasterKeyId"),
//   		sseType: jsii.String("sseType"),
//   	},
//   	streamSpecification: &streamSpecificationProperty{
//   		streamViewType: jsii.String("streamViewType"),
//   	},
//   	tableClass: jsii.String("tableClass"),
//   	tableName: jsii.String("tableName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	timeToLiveSpecification: &timeToLiveSpecificationProperty{
//   		attributeName: jsii.String("attributeName"),
//   		enabled: jsii.Boolean(false),
//   	},
//   }
//
type CfnTableProps struct {
	// Specifies the attributes that make up the primary key for the table.
	//
	// The attributes in the `KeySchema` property must also be defined in the `AttributeDefinitions` property.
	KeySchema interface{} `field:"required" json:"keySchema" yaml:"keySchema"`
	// A list of attributes that describe the key schema for the table and indexes.
	//
	// This property is required to create a DynamoDB table.
	//
	// Update requires: [Some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt) . Replacement if you edit an existing AttributeDefinition.
	AttributeDefinitions interface{} `field:"optional" json:"attributeDefinitions" yaml:"attributeDefinitions"`
	// Specify how you are charged for read and write throughput and how you manage capacity.
	//
	// Valid values include:
	//
	// - `PROVISIONED` - We recommend using `PROVISIONED` for predictable workloads. `PROVISIONED` sets the billing mode to [Provisioned Mode](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.ProvisionedThroughput.Manual) .
	// - `PAY_PER_REQUEST` - We recommend using `PAY_PER_REQUEST` for unpredictable workloads. `PAY_PER_REQUEST` sets the billing mode to [On-Demand Mode](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.OnDemand) .
	//
	// If not specified, the default is `PROVISIONED` .
	BillingMode *string `field:"optional" json:"billingMode" yaml:"billingMode"`
	// The settings used to enable or disable CloudWatch Contributor Insights for the specified table.
	ContributorInsightsSpecification interface{} `field:"optional" json:"contributorInsightsSpecification" yaml:"contributorInsightsSpecification"`
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
	GlobalSecondaryIndexes interface{} `field:"optional" json:"globalSecondaryIndexes" yaml:"globalSecondaryIndexes"`
	// `AWS::DynamoDB::Table.ImportSourceSpecification`.
	ImportSourceSpecification interface{} `field:"optional" json:"importSourceSpecification" yaml:"importSourceSpecification"`
	// The Kinesis Data Streams configuration for the specified table.
	KinesisStreamSpecification interface{} `field:"optional" json:"kinesisStreamSpecification" yaml:"kinesisStreamSpecification"`
	// Local secondary indexes to be created on the table.
	//
	// You can create up to 5 local secondary indexes. Each index is scoped to a given hash key value. The size of each hash key can be up to 10 gigabytes.
	LocalSecondaryIndexes interface{} `field:"optional" json:"localSecondaryIndexes" yaml:"localSecondaryIndexes"`
	// The settings used to enable point in time recovery.
	PointInTimeRecoverySpecification interface{} `field:"optional" json:"pointInTimeRecoverySpecification" yaml:"pointInTimeRecoverySpecification"`
	// Throughput for the specified table, which consists of values for `ReadCapacityUnits` and `WriteCapacityUnits` .
	//
	// For more information about the contents of a provisioned throughput structure, see [Amazon DynamoDB Table ProvisionedThroughput](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-provisionedthroughput.html) .
	//
	// If you set `BillingMode` as `PROVISIONED` , you must specify this property. If you set `BillingMode` as `PAY_PER_REQUEST` , you cannot specify this property.
	ProvisionedThroughput interface{} `field:"optional" json:"provisionedThroughput" yaml:"provisionedThroughput"`
	// Specifies the settings to enable server-side encryption.
	SseSpecification interface{} `field:"optional" json:"sseSpecification" yaml:"sseSpecification"`
	// The settings for the DynamoDB table stream, which capture changes to items stored in the table.
	StreamSpecification interface{} `field:"optional" json:"streamSpecification" yaml:"streamSpecification"`
	// The table class of the new table.
	//
	// Valid values are `STANDARD` and `STANDARD_INFREQUENT_ACCESS` .
	TableClass *string `field:"optional" json:"tableClass" yaml:"tableClass"`
	// A name for the table.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the table name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Specifies the Time to Live (TTL) settings for the table.
	//
	// > For detailed information about the limits in DynamoDB, see [Limits in Amazon DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html) in the Amazon DynamoDB Developer Guide.
	TimeToLiveSpecification interface{} `field:"optional" json:"timeToLiveSpecification" yaml:"timeToLiveSpecification"`
}

