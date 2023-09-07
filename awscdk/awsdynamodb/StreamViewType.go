package awsdynamodb


// When an item in the table is modified, StreamViewType determines what information is written to the stream for this table.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import kinesis "github.com/aws/aws-cdk-go/awscdk"
//
//
//   app := cdk.NewApp()
//   stack := cdk.NewStack(app, jsii.String("Stack"), &StackProps{
//   	Env: &Environment{
//   		Region: jsii.String("us-west-2"),
//   	},
//   })
//
//   globalTable := dynamodb.NewTableV2(this, jsii.String("GlobalTable"), &TablePropsV2{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("id"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   	DynamoStream: dynamodb.StreamViewType_OLD_IMAGE,
//   	// tables in us-west-2, us-east-1, and us-east-2 all have dynamo stream type of OLD_IMAGES
//   	Replicas: []replicaTableProps{
//   		&replicaTableProps{
//   			Region: jsii.String("us-east-1"),
//   		},
//   		&replicaTableProps{
//   			Region: jsii.String("us-east-2"),
//   		},
//   	},
//   })
//
// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_StreamSpecification.html
//
type StreamViewType string

const (
	// The entire item, as it appears after it was modified, is written to the stream.
	StreamViewType_NEW_IMAGE StreamViewType = "NEW_IMAGE"
	// The entire item, as it appeared before it was modified, is written to the stream.
	StreamViewType_OLD_IMAGE StreamViewType = "OLD_IMAGE"
	// Both the new and the old item images of the item are written to the stream.
	StreamViewType_NEW_AND_OLD_IMAGES StreamViewType = "NEW_AND_OLD_IMAGES"
	// Only the key attributes of the modified item are written to the stream.
	StreamViewType_KEYS_ONLY StreamViewType = "KEYS_ONLY"
)

