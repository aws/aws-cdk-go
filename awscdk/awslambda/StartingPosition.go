package awslambda


// The position in the DynamoDB, Kinesis or MSK stream where AWS Lambda should start reading.
//
// Example:
//   import dynamodb "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var table Table
//
//   var fn Function
//
//
//   deadLetterQueue := sqs.NewQueue(this, jsii.String("deadLetterQueue"))
//   fn.AddEventSource(awscdk.NewDynamoEventSource(table, &DynamoEventSourceProps{
//   	StartingPosition: lambda.StartingPosition_TRIM_HORIZON,
//   	BatchSize: jsii.Number(5),
//   	BisectBatchOnError: jsii.Boolean(true),
//   	OnFailure: awscdk.NewSqsDlq(deadLetterQueue),
//   	RetryAttempts: jsii.Number(10),
//   }))
//
type StartingPosition string

const (
	// Start reading at the last untrimmed record in the shard in the system, which is the oldest data record in the shard.
	StartingPosition_TRIM_HORIZON StartingPosition = "TRIM_HORIZON"
	// Start reading just after the most recent record in the shard, so that you always read the most recent data in the shard.
	StartingPosition_LATEST StartingPosition = "LATEST"
	// Start reading from a position defined by a time stamp.
	//
	// Only supported for Amazon Kinesis streams, otherwise an error will occur.
	// If supplied, `startingPositionTimestamp` must also be set.
	StartingPosition_AT_TIMESTAMP StartingPosition = "AT_TIMESTAMP"
)

