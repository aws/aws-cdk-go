package awsdynamodb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
)

// Properties used to configure a DynamoDB table.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   app := cdk.NewApp()
//   stack := cdk.NewStack(app, jsii.String("Stack"), &StackProps{
//   	Env: &Environment{
//   		Region: jsii.String("us-west-2"),
//   	},
//   })
//
//   mrscTable := dynamodb.NewTableV2(stack, jsii.String("MRSCTable"), &TablePropsV2{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("pk"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   	MultiRegionConsistency: dynamodb.MultiRegionConsistency_STRONG,
//   	Replicas: []ReplicaTableProps{
//   		&ReplicaTableProps{
//   			Region: jsii.String("us-east-1"),
//   		},
//   		&ReplicaTableProps{
//   			Region: jsii.String("us-east-2"),
//   		},
//   	},
//   })
//
type TablePropsV2 struct {
	// Whether CloudWatch contributor insights is enabled.
	// Default: false.
	//
	// Deprecated: use `contributorInsightsSpecification` instead.
	ContributorInsights *bool `field:"optional" json:"contributorInsights" yaml:"contributorInsights"`
	// Whether CloudWatch contributor insights is enabled and what mode is selected.
	// Default: - contributor insights is not enabled.
	//
	ContributorInsightsSpecification *ContributorInsightsSpecification `field:"optional" json:"contributorInsightsSpecification" yaml:"contributorInsightsSpecification"`
	// Whether deletion protection is enabled.
	// Default: false.
	//
	DeletionProtection *bool `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Kinesis Data Stream to capture item level changes.
	// Default: - no Kinesis Data Stream.
	//
	KinesisStream awskinesis.IStream `field:"optional" json:"kinesisStream" yaml:"kinesisStream"`
	// Whether point-in-time recovery is enabled.
	// Default: false - point in time recovery is not enabled.
	//
	// Deprecated: use `pointInTimeRecoverySpecification` instead.
	PointInTimeRecovery *bool `field:"optional" json:"pointInTimeRecovery" yaml:"pointInTimeRecovery"`
	// Whether point-in-time recovery is enabled and recoveryPeriodInDays is set.
	// Default: - point in time recovery is not enabled.
	//
	PointInTimeRecoverySpecification *PointInTimeRecoverySpecification `field:"optional" json:"pointInTimeRecoverySpecification" yaml:"pointInTimeRecoverySpecification"`
	// Resource policy to assign to DynamoDB Table.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-replicaspecification.html#cfn-dynamodb-globaltable-replicaspecification-resourcepolicy
	//
	// Default: - No resource policy statements are added to the created table.
	//
	ResourcePolicy awsiam.PolicyDocument `field:"optional" json:"resourcePolicy" yaml:"resourcePolicy"`
	// The table class.
	// Default: TableClass.STANDARD
	//
	TableClass TableClass `field:"optional" json:"tableClass" yaml:"tableClass"`
	// Tags to be applied to the primary table (default replica table).
	// Default: - no tags.
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Partition key attribute definition.
	PartitionKey *Attribute `field:"required" json:"partitionKey" yaml:"partitionKey"`
	// The billing mode and capacity settings to apply to the table.
	// Default: Billing.onDemand()
	//
	Billing Billing `field:"optional" json:"billing" yaml:"billing"`
	// When an item in the table is modified, StreamViewType determines what information is written to the stream.
	// Default: - streams are disabled if replicas are not configured and this property is
	// not specified. If this property is not specified when replicas are configured, then
	// NEW_AND_OLD_IMAGES will be the StreamViewType for all replicas.
	//
	DynamoStream StreamViewType `field:"optional" json:"dynamoStream" yaml:"dynamoStream"`
	// The server-side encryption.
	// Default: TableEncryptionV2.dynamoOwnedKey()
	//
	Encryption TableEncryptionV2 `field:"optional" json:"encryption" yaml:"encryption"`
	// Global secondary indexes.
	//
	// Note: You can provide a maximum of 20 global secondary indexes.
	// Default: - no global secondary indexes.
	//
	GlobalSecondaryIndexes *[]*GlobalSecondaryIndexPropsV2 `field:"optional" json:"globalSecondaryIndexes" yaml:"globalSecondaryIndexes"`
	// Controls whether table settings are synchronized across replicas.
	//
	// When set to ALL, synchronizable settings (billing mode, throughput, TTL, streams view type, GSIs)
	// are automatically replicated across all replicas. When set to NONE, each replica manages its own
	// settings independently (billing mode must be PAY_PER_REQUEST).
	//
	// Note: Some settings are always synchronized (key schema, LSIs) regardless of this setting,
	// and some are never synchronized (table class, SSE, deletion protection, PITR, tags, resource policy).
	// Default: GlobalTableSettingsReplicationMode.NONE
	//
	GlobalTableSettingsReplicationMode GlobalTableSettingsReplicationMode `field:"optional" json:"globalTableSettingsReplicationMode" yaml:"globalTableSettingsReplicationMode"`
	// Local secondary indexes.
	//
	// Note: You can only provide a maximum of 5 local secondary indexes.
	// Default: - no local secondary indexes.
	//
	LocalSecondaryIndexes *[]*LocalSecondaryIndexProps `field:"optional" json:"localSecondaryIndexes" yaml:"localSecondaryIndexes"`
	// Specifies the consistency mode for a new global table.
	// Default: MultiRegionConsistency.EVENTUAL
	//
	MultiRegionConsistency MultiRegionConsistency `field:"optional" json:"multiRegionConsistency" yaml:"multiRegionConsistency"`
	// The removal policy applied to the table.
	// Default: RemovalPolicy.RETAIN
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Replica tables to deploy with the primary table.
	//
	// Note: Adding replica tables allows you to use your table as a global table. You
	// cannot specify a replica table in the region that the primary table will be deployed
	// to. Replica tables will only be supported if the stack deployment region is defined.
	// Default: - no replica tables.
	//
	Replicas *[]*ReplicaTableProps `field:"optional" json:"replicas" yaml:"replicas"`
	// Sort key attribute definition.
	// Default: - no sort key.
	//
	SortKey *Attribute `field:"optional" json:"sortKey" yaml:"sortKey"`
	// The name of the table.
	// Default: - generated by CloudFormation.
	//
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// The name of the TTL attribute.
	// Default: - TTL is disabled.
	//
	TimeToLiveAttribute *string `field:"optional" json:"timeToLiveAttribute" yaml:"timeToLiveAttribute"`
	// The warm throughput configuration for the table.
	// Default: - no warm throughput is configured.
	//
	WarmThroughput *WarmThroughput `field:"optional" json:"warmThroughput" yaml:"warmThroughput"`
	// The witness Region for the MRSC global table.
	//
	// A MRSC global table can be configured with either three replicas, or with two replicas and one witness.
	//
	// Note: Witness region cannot be specified for a Multi-Region Eventual Consistency (MREC) Global Table.
	// Witness regions are only supported for Multi-Region Strong Consistency (MRSC) Global Tables.
	// Default: - no witness region.
	//
	WitnessRegion *string `field:"optional" json:"witnessRegion" yaml:"witnessRegion"`
}

