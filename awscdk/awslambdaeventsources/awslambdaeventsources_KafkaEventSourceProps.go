package awslambdaeventsources

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Properties for a Kafka event source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var secret secret
//
//   kafkaEventSourceProps := &kafkaEventSourceProps{
//   	startingPosition: awscdk.Aws_lambda.startingPosition_TRIM_HORIZON,
//   	topic: jsii.String("topic"),
//
//   	// the properties below are optional
//   	batchSize: jsii.Number(123),
//   	enabled: jsii.Boolean(false),
//   	maxBatchingWindow: cdk.duration.minutes(jsii.Number(30)),
//   	secret: secret,
//   }
//
type KafkaEventSourceProps struct {
	// Where to begin consuming the stream.
	StartingPosition awslambda.StartingPosition `field:"required" json:"startingPosition" yaml:"startingPosition"`
	// The largest number of records that AWS Lambda will retrieve from your event source at the time of invoking your function.
	//
	// Your function receives an
	// event with all the retrieved records.
	//
	// Valid Range:
	// * Minimum value of 1
	// * Maximum value of:
	//    * 1000 for {@link DynamoEventSource}
	// * 10000 for {@link KinesisEventSource}, {@link ManagedKafkaEventSource} and {@link SelfManagedKafkaEventSource}.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// If the stream event source mapping should be enabled.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The maximum amount of time to gather records before invoking the function.
	//
	// Maximum of Duration.minutes(5)
	MaxBatchingWindow awscdk.Duration `field:"optional" json:"maxBatchingWindow" yaml:"maxBatchingWindow"`
	// The Kafka topic to subscribe to.
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// The secret with the Kafka credentials, see https://docs.aws.amazon.com/msk/latest/developerguide/msk-password.html for details This field is required if your Kafka brokers are accessed over the Internet.
	Secret awssecretsmanager.ISecret `field:"optional" json:"secret" yaml:"secret"`
}

