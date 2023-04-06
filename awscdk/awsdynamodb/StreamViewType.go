package awsdynamodb


// When an item in the table is modified, StreamViewType determines what information is written to the stream for this table.
//
// Example:
//   import eventsources "github.com/aws/aws-cdk-go/awscdk"
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
//   fn.AddEventSource(eventsources.NewDynamoEventSource(table, &DynamoEventSourceProps{
//   	StartingPosition: lambda.StartingPosition_LATEST,
//   	Filters: []map[string]interface{}{
//   		lambda.FilterCriteria_Filter(map[string]interface{}{
//   			"eventName": lambda.FilterRule_isEqual(jsii.String("INSERT")),
//   		}),
//   	},
//   }))
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

