package awscdkpipessourcesalpha


// The position in a DynamoDB stream from which to start reading.
//
// Example:
//   var targetQueue Queue
//   table := ddb.NewTableV2(this, jsii.String("MyTable"), &TablePropsV2{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("id"),
//   		Type: ddb.AttributeType_STRING,
//   	},
//   	DynamoStream: ddb.StreamViewType_NEW_IMAGE,
//   })
//
//   pipeSource := sources.NewDynamoDBSource(table, &DynamoDBSourceParameters{
//   	StartingPosition: sources.DynamoDBStartingPosition_LATEST,
//   })
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: pipeSource,
//   	Target: awscdkpipestargetsalpha.NewSqsTarget(targetQueue),
//   })
//
// Experimental.
type DynamoDBStartingPosition string

const (
	// Start reading at the last (untrimmed) stream record, which is the oldest record in the shard.
	// Experimental.
	DynamoDBStartingPosition_TRIM_HORIZON DynamoDBStartingPosition = "TRIM_HORIZON"
	// Start reading just after the most recent stream record in the shard, so that you always read the most recent data in the shard.
	// Experimental.
	DynamoDBStartingPosition_LATEST DynamoDBStartingPosition = "LATEST"
)

