package awsdynamodb


// Global table multi-region consistency mode.
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
// See: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/V2globaltables_HowItWorks.html#V2globaltables_HowItWorks.consistency-modes-mrsc
//
type MultiRegionConsistency string

const (
	// Default consistency mode for Global Tables.
	//
	// Multi-region eventual consistency.
	MultiRegionConsistency_EVENTUAL MultiRegionConsistency = "EVENTUAL"
	// Multi-region strong consistency.
	MultiRegionConsistency_STRONG MultiRegionConsistency = "STRONG"
)

