package awsdynamodb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for a DynamoDB Table.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import cloudwatch "github.com/aws/aws-cdk-go/awscdk"
//
//
//   table := dynamodb.NewTable(this, jsii.String("Table"), &tableProps{
//   	partitionKey: &attribute{
//   		name: jsii.String("id"),
//   		type: dynamodb.attributeType_STRING,
//   	},
//   })
//
//   metric := table.metricThrottledRequestsForOperations(&operationsMetricOptions{
//   	operations: []operation{
//   		dynamodb.*operation_PUT_ITEM,
//   	},
//   	period: awscdk.Duration.minutes(jsii.Number(1)),
//   })
//
//   cloudwatch.NewAlarm(stack, jsii.String("Alarm"), &alarmProps{
//   	metric: metric,
//   	evaluationPeriods: jsii.Number(1),
//   	threshold: jsii.Number(1),
//   })
//
type TableProps struct {
	// Partition key attribute definition.
	PartitionKey *Attribute `field:"required" json:"partitionKey" yaml:"partitionKey"`
	// Sort key attribute definition.
	SortKey *Attribute `field:"optional" json:"sortKey" yaml:"sortKey"`
	// Specify how you are charged for read and write throughput and how you manage capacity.
	BillingMode BillingMode `field:"optional" json:"billingMode" yaml:"billingMode"`
	// Whether CloudWatch contributor insights is enabled.
	ContributorInsightsEnabled *bool `field:"optional" json:"contributorInsightsEnabled" yaml:"contributorInsightsEnabled"`
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
	Encryption TableEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// External KMS key to use for table encryption.
	//
	// This property can only be set if `encryption` is set to `TableEncryption.CUSTOMER_MANAGED`.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Whether point-in-time recovery is enabled.
	PointInTimeRecovery *bool `field:"optional" json:"pointInTimeRecovery" yaml:"pointInTimeRecovery"`
	// The read capacity for the table.
	//
	// Careful if you add Global Secondary Indexes, as
	// those will share the table's provisioned throughput.
	//
	// Can only be provided if billingMode is Provisioned.
	ReadCapacity *float64 `field:"optional" json:"readCapacity" yaml:"readCapacity"`
	// The removal policy to apply to the DynamoDB Table.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Regions where replica tables will be created.
	ReplicationRegions *[]*string `field:"optional" json:"replicationRegions" yaml:"replicationRegions"`
	// The timeout for a table replication operation in a single region.
	ReplicationTimeout awscdk.Duration `field:"optional" json:"replicationTimeout" yaml:"replicationTimeout"`
	// When an item in the table is modified, StreamViewType determines what information is written to the stream for this table.
	Stream StreamViewType `field:"optional" json:"stream" yaml:"stream"`
	// Specify the table class.
	TableClass TableClass `field:"optional" json:"tableClass" yaml:"tableClass"`
	// The name of TTL attribute.
	TimeToLiveAttribute *string `field:"optional" json:"timeToLiveAttribute" yaml:"timeToLiveAttribute"`
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
	WaitForReplicationToFinish *bool `field:"optional" json:"waitForReplicationToFinish" yaml:"waitForReplicationToFinish"`
	// The write capacity for the table.
	//
	// Careful if you add Global Secondary Indexes, as
	// those will share the table's provisioned throughput.
	//
	// Can only be provided if billingMode is Provisioned.
	WriteCapacity *float64 `field:"optional" json:"writeCapacity" yaml:"writeCapacity"`
	// Kinesis Data Stream to capture item-level changes for the table.
	KinesisStream awskinesis.IStream `field:"optional" json:"kinesisStream" yaml:"kinesisStream"`
	// Enforces a particular physical table name.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

