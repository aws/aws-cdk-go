package awsdynamodb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for a DynamoDB Table.
//
// Example:
//   import eventsources "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var fn function
//
//   table := dynamodb.NewTable(this, jsii.String("Table"), &TableProps{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("id"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   	Stream: dynamodb.StreamViewType_NEW_IMAGE,
//   })
//   // Your self managed KMS key
//   myKey := awscdk.Key_FromKeyArn(this, jsii.String("SourceBucketEncryptionKey"), jsii.String("arn:aws:kms:us-east-1:123456789012:key/<key-id>"))
//
//   fn.AddEventSource(eventsources.NewDynamoEventSource(table, &DynamoEventSourceProps{
//   	StartingPosition: lambda.StartingPosition_LATEST,
//   	Filters: []map[string]interface{}{
//   		lambda.FilterCriteria_Filter(map[string]interface{}{
//   			"eventName": lambda.FilterRule_isEqual(jsii.String("INSERT")),
//   		}),
//   	},
//   	FilterEncryption: myKey,
//   }))
//
type TableProps struct {
	// Partition key attribute definition.
	PartitionKey *Attribute `field:"required" json:"partitionKey" yaml:"partitionKey"`
	// Sort key attribute definition.
	// Default: no sort key.
	//
	SortKey *Attribute `field:"optional" json:"sortKey" yaml:"sortKey"`
	// Specify how you are charged for read and write throughput and how you manage capacity.
	// Default: PROVISIONED if `replicationRegions` is not specified, PAY_PER_REQUEST otherwise.
	//
	BillingMode BillingMode `field:"optional" json:"billingMode" yaml:"billingMode"`
	// Whether CloudWatch contributor insights is enabled.
	// Default: false.
	//
	ContributorInsightsEnabled *bool `field:"optional" json:"contributorInsightsEnabled" yaml:"contributorInsightsEnabled"`
	// Enables deletion protection for the table.
	// Default: false.
	//
	DeletionProtection *bool `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
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
	// Default: - The table is encrypted with an encryption key managed by DynamoDB, and you are not charged any fee for using it.
	//
	Encryption TableEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// External KMS key to use for table encryption.
	//
	// This property can only be set if `encryption` is set to `TableEncryption.CUSTOMER_MANAGED`.
	// Default: - If `encryption` is set to `TableEncryption.CUSTOMER_MANAGED` and this
	// property is undefined, a new KMS key will be created and associated with this table.
	// If `encryption` and this property are both undefined, then the table is encrypted with
	// an encryption key managed by DynamoDB, and you are not charged any fee for using it.
	//
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The properties of data being imported from the S3 bucket source to the table.
	// Default: - no data import from the S3 bucket.
	//
	ImportSource *ImportSourceSpecification `field:"optional" json:"importSource" yaml:"importSource"`
	// The maximum read request units for the table.
	//
	// Careful if you add Global Secondary Indexes, as
	// those will share the table's maximum on-demand throughput.
	//
	// Can only be provided if billingMode is PAY_PER_REQUEST.
	// Default: - on-demand throughput is disabled.
	//
	MaxReadRequestUnits *float64 `field:"optional" json:"maxReadRequestUnits" yaml:"maxReadRequestUnits"`
	// The write request units for the table.
	//
	// Careful if you add Global Secondary Indexes, as
	// those will share the table's maximum on-demand throughput.
	//
	// Can only be provided if billingMode is PAY_PER_REQUEST.
	// Default: - on-demand throughput is disabled.
	//
	MaxWriteRequestUnits *float64 `field:"optional" json:"maxWriteRequestUnits" yaml:"maxWriteRequestUnits"`
	// Whether point-in-time recovery is enabled.
	// Default: false - point in time recovery is not enabled.
	//
	// Deprecated: use `pointInTimeRecoverySpecification` instead.
	PointInTimeRecovery *bool `field:"optional" json:"pointInTimeRecovery" yaml:"pointInTimeRecovery"`
	// Whether point-in-time recovery is enabled and recoveryPeriodInDays is set.
	// Default: - point in time recovery is not enabled.
	//
	PointInTimeRecoverySpecification *PointInTimeRecoverySpecification `field:"optional" json:"pointInTimeRecoverySpecification" yaml:"pointInTimeRecoverySpecification"`
	// The read capacity for the table.
	//
	// Careful if you add Global Secondary Indexes, as
	// those will share the table's provisioned throughput.
	//
	// Can only be provided if billingMode is Provisioned.
	// Default: 5.
	//
	ReadCapacity *float64 `field:"optional" json:"readCapacity" yaml:"readCapacity"`
	// The removal policy to apply to the DynamoDB Table.
	// Default: RemovalPolicy.RETAIN
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Regions where replica tables will be created.
	// Default: - no replica tables are created.
	//
	ReplicationRegions *[]*string `field:"optional" json:"replicationRegions" yaml:"replicationRegions"`
	// The timeout for a table replication operation in a single region.
	// Default: Duration.minutes(30)
	//
	ReplicationTimeout awscdk.Duration `field:"optional" json:"replicationTimeout" yaml:"replicationTimeout"`
	// Resource policy to assign to table.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-resourcepolicy
	//
	// Default: - No resource policy statement.
	//
	ResourcePolicy awsiam.PolicyDocument `field:"optional" json:"resourcePolicy" yaml:"resourcePolicy"`
	// When an item in the table is modified, StreamViewType determines what information is written to the stream for this table.
	// Default: - streams are disabled unless `replicationRegions` is specified.
	//
	Stream StreamViewType `field:"optional" json:"stream" yaml:"stream"`
	// Specify the table class.
	// Default: STANDARD.
	//
	TableClass TableClass `field:"optional" json:"tableClass" yaml:"tableClass"`
	// The name of TTL attribute.
	// Default: - TTL is disabled.
	//
	TimeToLiveAttribute *string `field:"optional" json:"timeToLiveAttribute" yaml:"timeToLiveAttribute"`
	// [WARNING: Use this flag with caution, misusing this flag may cause deleting existing replicas, refer to the detailed documentation for more information] Indicates whether CloudFormation stack waits for replication to finish.
	//
	// If set to false, the CloudFormation resource will mark the resource as
	// created and replication will be completed asynchronously. This property is
	// ignored if replicationRegions property is not set.
	//
	// WARNING:
	// DO NOT UNSET this property if adding/removing multiple replicationRegions
	// in one deployment, as CloudFormation only supports one region replication
	// at a time. CDK overcomes this limitation by waiting for replication to
	// finish before starting new replicationRegion.
	//
	// If the custom resource which handles replication has a physical resource
	// ID with the format `region` instead of `tablename-region` (this would happen
	// if the custom resource hasn't received an event since v1.91.0), DO NOT SET
	// this property to false without making a change to the table name.
	// This will cause the existing replicas to be deleted.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-globaltable.html#cfn-dynamodb-globaltable-replicas
	//
	// Default: true.
	//
	WaitForReplicationToFinish *bool `field:"optional" json:"waitForReplicationToFinish" yaml:"waitForReplicationToFinish"`
	// Specify values to pre-warm you DynamoDB Table Warm Throughput feature is not available for Global Table replicas using the `Table` construct.
	//
	// To enable Warm Throughput, use the `TableV2` construct instead.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dynamodb-table.html#cfn-dynamodb-table-warmthroughput
	//
	// Default: - warm throughput is not configured.
	//
	WarmThroughput *WarmThroughput `field:"optional" json:"warmThroughput" yaml:"warmThroughput"`
	// The write capacity for the table.
	//
	// Careful if you add Global Secondary Indexes, as
	// those will share the table's provisioned throughput.
	//
	// Can only be provided if billingMode is Provisioned.
	// Default: 5.
	//
	WriteCapacity *float64 `field:"optional" json:"writeCapacity" yaml:"writeCapacity"`
	// Kinesis Data Stream approximate creation timestamp precision.
	// Default: ApproximateCreationDateTimePrecision.MICROSECOND
	//
	KinesisPrecisionTimestamp ApproximateCreationDateTimePrecision `field:"optional" json:"kinesisPrecisionTimestamp" yaml:"kinesisPrecisionTimestamp"`
	// Kinesis Data Stream to capture item-level changes for the table.
	// Default: - no Kinesis Data Stream.
	//
	KinesisStream awskinesis.IStream `field:"optional" json:"kinesisStream" yaml:"kinesisStream"`
	// Enforces a particular physical table name.
	// Default: <generated>.
	//
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

