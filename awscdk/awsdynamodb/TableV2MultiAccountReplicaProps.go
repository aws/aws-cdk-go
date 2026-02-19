package awsdynamodb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
)

// Properties for creating a multi-account replica table.
//
// Note: partitionKey, sortKey, and localSecondaryIndexes are not options because CloudFormation
// automatically inherits the key schema and LSIs from the source table via globalTableSourceArn.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   app := cdk.NewApp()
//
//   // Source table in Account A
//   sourceStack := cdk.NewStack(app, jsii.String("SourceStack"), &StackProps{
//   	Env: &Environment{
//   		Region: jsii.String("us-east-2"),
//   		Account: jsii.String("111111111111"),
//   	},
//   })
//
//   sourceTable := dynamodb.NewTableV2(sourceStack, jsii.String("SourceTable"), &TablePropsV2{
//   	TableName: jsii.String("MyMultiAccountTable"),
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("pk"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   	GlobalTableSettingsReplicationMode: dynamodb.GlobalTableSettingsReplicationMode_ALL,
//   })
//
//   // Replica stack in Account B
//   replicaStack := cdk.NewStack(app, jsii.String("ReplicaStack"), &StackProps{
//   	Env: &Environment{
//   		Region: jsii.String("us-east-1"),
//   		Account: jsii.String("222222222222"),
//   	},
//   })
//
//   // Create replica - permissions are automatically configured
//   replica := dynamodb.NewTableV2MultiAccountReplica(replicaStack, jsii.String("ReplicaTable"), &TableV2MultiAccountReplicaProps{
//   	TableName: jsii.String("MyMultiAccountTable"),
//   	ReplicaSourceTable: sourceTable,
//   	GlobalTableSettingsReplicationMode: dynamodb.GlobalTableSettingsReplicationMode_ALL,
//   })
//
type TableV2MultiAccountReplicaProps struct {
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
	// The server-side encryption configuration for the replica table.
	//
	// Note: Each replica manages its own encryption independently. This is not synchronized
	// across replicas.
	// Default: TableEncryptionV2.dynamoOwnedKey()
	//
	Encryption TableEncryptionV2 `field:"optional" json:"encryption" yaml:"encryption"`
	// Controls whether table settings are synchronized across replicas.
	//
	// When set to ALL, synchronizable settings (billing mode, throughput, TTL, streams view type, GSIs)
	// are automatically replicated across all replicas. When set to NONE, each replica manages its own
	// settings independently (billing mode must be PAY_PER_REQUEST).
	//
	// Note: Some settings are always synchronized (key schema, LSIs) regardless of this setting,
	// and some are never synchronized (table class, SSE, deletion protection, PITR, tags, resource policy).
	// Default: GlobalTableSettingsReplicationMode.ALL
	//
	GlobalTableSettingsReplicationMode GlobalTableSettingsReplicationMode `field:"optional" json:"globalTableSettingsReplicationMode" yaml:"globalTableSettingsReplicationMode"`
	// Whether or not to grant permissions for all indexes of the table.
	//
	// Note: If false, permissions will only be granted to indexes when `globalIndexes` is specified.
	// Default: false.
	//
	GrantIndexPermissions *bool `field:"optional" json:"grantIndexPermissions" yaml:"grantIndexPermissions"`
	// The removal policy applied to the table.
	// Default: RemovalPolicy.RETAIN
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The source table to replicate from.
	//
	// [disable-awslint:prefer-ref-interface].
	// Default: - must be provided.
	//
	ReplicaSourceTable ITableV2 `field:"optional" json:"replicaSourceTable" yaml:"replicaSourceTable"`
	// Enforces a particular physical table name.
	// Default: - generated by CloudFormation.
	//
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

