package awslambda


// The position in the DynamoDB, Kinesis or MSK stream where AWS Lambda should start reading.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   // The secret that allows access to your self hosted Kafka cluster
//   var secret secret
//
//   var myFunction function
//
//
//   // The list of Kafka brokers
//   bootstrapServers := []*string{
//   	"kafka-broker:9092",
//   }
//
//   // The Kafka topic you want to subscribe to
//   topic := "some-cool-topic"
//   myFunction.addEventSource(awscdk.NewSelfManagedKafkaEventSource(&selfManagedKafkaEventSourceProps{
//   	bootstrapServers: bootstrapServers,
//   	topic: topic,
//   	secret: secret,
//   	batchSize: jsii.Number(100),
//   	 // default
//   	startingPosition: lambda.startingPosition_TRIM_HORIZON,
//   }))
//
// Experimental.
type StartingPosition string

const (
	// Start reading at the last untrimmed record in the shard in the system, which is the oldest data record in the shard.
	// Experimental.
	StartingPosition_TRIM_HORIZON StartingPosition = "TRIM_HORIZON"
	// Start reading just after the most recent record in the shard, so that you always read the most recent data in the shard.
	// Experimental.
	StartingPosition_LATEST StartingPosition = "LATEST"
)

