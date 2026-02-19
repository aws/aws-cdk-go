package awsdynamodb


// The replication mode for global table settings across multiple accounts.
//
// Note: In a multi-account global table, you cannot make changes to a synchronized setting using CDK.
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
//   		Region: jsii.String("us-east-1"),
//   		Account: jsii.String("111111111111"),
//   	},
//   })
//
//   // Region us-west-2
//   sourceTable := dynamodb.NewTableV2(sourceStack, jsii.String("SourceTable"), &TablePropsV2{
//   	TableName: jsii.String("MyMultiAccountTable"),
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("pk"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   	GlobalTableSettingsReplicationMode: dynamodb.GlobalTableSettingsReplicationMode_ALL,
//   })
//   // After replica is deployed, update source stack with the ARN
//   sourceTable.Grants.MultiAccountReplicationTo(jsii.String("arn:aws:dynamodb:us-east-1:222222222222:table/MyMultiAccountTable"))
//
// See: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/V2globaltables_MA_HowItWorks.html
//
type GlobalTableSettingsReplicationMode string

const (
	// All synchronizable settings are replicated across all replicas.
	//
	// Synchronizable settings include: billing mode, provisioned throughput, auto-scaling,
	// on-demand throughput, warm throughput, TTL, streams view type, and GSIs.
	//
	// Note: Some settings are always synchronized (key schema, LSIs) and some are never
	// synchronized (table class, SSE, deletion protection, PITR, tags, resource policy, CCI).
	GlobalTableSettingsReplicationMode_ALL GlobalTableSettingsReplicationMode = "ALL"
)

