package awsdynamodb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis"
)

// Properties used to configure a replica table.
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
//   globalTable := dynamodb.NewTableV2(stack, jsii.String("GlobalTable"), &TablePropsV2{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("pk"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   	Replicas: []replicaTableProps{
//   		&replicaTableProps{
//   			Region: jsii.String("us-east-1"),
//   		},
//   	},
//   })
//
//   globalTable.AddReplica(&replicaTableProps{
//   	Region: jsii.String("us-east-2"),
//   	DeletionProtection: jsii.Boolean(true),
//   })
//
type ReplicaTableProps struct {
	// Whether CloudWatch contributor insights is enabled.
	// Default: false.
	//
	ContributorInsights *bool `field:"optional" json:"contributorInsights" yaml:"contributorInsights"`
	// Whether deletion protection is enabled.
	// Default: false.
	//
	DeletionProtection *bool `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Kinesis Data Stream to capture item level changes.
	// Default: - no Kinesis Data Stream.
	//
	KinesisStream awskinesis.IStream `field:"optional" json:"kinesisStream" yaml:"kinesisStream"`
	// Whether point-in-time recovery is enabled.
	// Default: false.
	//
	PointInTimeRecovery *bool `field:"optional" json:"pointInTimeRecovery" yaml:"pointInTimeRecovery"`
	// The table class.
	// Default: TableClass.STANDARD
	//
	TableClass TableClass `field:"optional" json:"tableClass" yaml:"tableClass"`
	// The region that the replica table will be created in.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Options used to configure global secondary index properties.
	// Default: - inherited from the primary table.
	//
	GlobalSecondaryIndexOptions *map[string]*ReplicaGlobalSecondaryIndexOptions `field:"optional" json:"globalSecondaryIndexOptions" yaml:"globalSecondaryIndexOptions"`
	// The read capacity.
	//
	// Note: This can only be configured if the primary table billing is provisioned.
	// Default: - inherited from the primary table.
	//
	ReadCapacity Capacity `field:"optional" json:"readCapacity" yaml:"readCapacity"`
}

