package awslogs


// The method used to distribute log data to the destination.
//
// Example:
//   import destinations "github.com/aws/aws-cdk-go/awscdk"
//   import kinesis "github.com/aws/aws-cdk-go/awscdk"
//
//   var stream stream
//   var logGroup logGroup
//
//
//   logs.NewSubscriptionFilter(this, jsii.String("Subscription"), &SubscriptionFilterProps{
//   	LogGroup: LogGroup,
//   	Destination: destinations.NewKinesisDestination(stream),
//   	FilterPattern: logs.FilterPattern_AllTerms(jsii.String("ERROR"), jsii.String("MainThread")),
//   	FilterName: jsii.String("ErrorInMainThread"),
//   	Distribution: logs.Distribution_RANDOM,
//   })
//
type Distribution string

const (
	// Log events from the same log stream are kept together and sent to the same destination.
	Distribution_BY_LOG_STREAM Distribution = "BY_LOG_STREAM"
	// Log events are distributed across the log destinations randomly.
	Distribution_RANDOM Distribution = "RANDOM"
)

