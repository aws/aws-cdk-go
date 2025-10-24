package awsdynamodb


// Data types for attributes within a table.
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
//   	},
//   	WitnessRegion: jsii.String("us-east-2"),
//   })
//
// See: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.NamingRulesDataTypes.html#HowItWorks.DataTypes
//
type AttributeType string

const (
	// Up to 400KiB of binary data (which must be encoded as base64 before sending to DynamoDB).
	AttributeType_BINARY AttributeType = "BINARY"
	// Numeric values made of up to 38 digits (positive, negative or zero).
	AttributeType_NUMBER AttributeType = "NUMBER"
	// Up to 400KiB of UTF-8 encoded text.
	AttributeType_STRING AttributeType = "STRING"
)

