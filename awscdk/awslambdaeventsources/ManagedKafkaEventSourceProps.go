package awslambdaeventsources

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Properties for a MSK event source.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var myFunction function
//
//
//   // Your MSK cluster arn
//   clusterArn := "arn:aws:kafka:us-east-1:0123456789019:cluster/SalesCluster/abcd1234-abcd-cafe-abab-9876543210ab-4"
//
//   // The Kafka topic you want to subscribe to
//   topic := "some-cool-topic"
//
//   // The secret that allows access to your MSK cluster
//   // You still have to make sure that it is associated with your cluster as described in the documentation
//   secret := awscdk.NewSecret(this, jsii.String("Secret"), &SecretProps{
//   	SecretName: jsii.String("AmazonMSK_KafkaSecret"),
//   })
//   myFunction.AddEventSource(awscdk.NewManagedKafkaEventSource(&ManagedKafkaEventSourceProps{
//   	ClusterArn: jsii.String(ClusterArn),
//   	Topic: topic,
//   	Secret: secret,
//   	BatchSize: jsii.Number(100),
//   	 // default
//   	StartingPosition: lambda.StartingPosition_TRIM_HORIZON,
//   }))
//
type ManagedKafkaEventSourceProps struct {
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
	//   * 1000 for `DynamoEventSource`
	// * 10000 for `KinesisEventSource`, `ManagedKafkaEventSource` and `SelfManagedKafkaEventSource`.
	// Default: 100.
	//
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// If the stream event source mapping should be enabled.
	// Default: true.
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The maximum amount of time to gather records before invoking the function.
	//
	// Maximum of Duration.minutes(5).
	// See: https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventsourcemapping.html#invocation-eventsourcemapping-batching
	//
	// Default: - Duration.seconds(0) for Kinesis, DynamoDB, and SQS event sources, Duration.millis(500) for MSK, self-managed Kafka, and Amazon MQ.
	//
	MaxBatchingWindow awscdk.Duration `field:"optional" json:"maxBatchingWindow" yaml:"maxBatchingWindow"`
	// Configuration for provisioned pollers that read from the event source.
	//
	// When specified, allows control over the minimum and maximum number of pollers
	// that can be provisioned to process events from the source.
	// Default: - no provisioned pollers.
	//
	ProvisionedPollerConfig *ProvisionedPollerConfig `field:"optional" json:"provisionedPollerConfig" yaml:"provisionedPollerConfig"`
	// The Kafka topic to subscribe to.
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// The identifier for the Kafka consumer group to join.
	//
	// The consumer group ID must be unique among all your Kafka event sources. After creating a Kafka event source mapping with the consumer group ID specified, you cannot update this value.  The value must have a length between 1 and 200 and full the pattern '[a-zA-Z0-9-\/*:_+=.@-]*'.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/with-msk.html#services-msk-consumer-group-id
	//
	// Default: - none.
	//
	ConsumerGroupId *string `field:"optional" json:"consumerGroupId" yaml:"consumerGroupId"`
	// Add Customer managed KMS key to encrypt Filter Criteria.
	// See: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk
	//
	// Default: - none.
	//
	FilterEncryption awskms.IKey `field:"optional" json:"filterEncryption" yaml:"filterEncryption"`
	// Add filter criteria to Event Source.
	// See: https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html
	//
	// Default: - none.
	//
	Filters *[]*map[string]interface{} `field:"optional" json:"filters" yaml:"filters"`
	// Add an on Failure Destination for this Kafka event.
	//
	// SNS/SQS/S3 are supported.
	// Default: - discarded records are ignored.
	//
	OnFailure awslambda.IEventSourceDlq `field:"optional" json:"onFailure" yaml:"onFailure"`
	// The secret with the Kafka credentials, see https://docs.aws.amazon.com/msk/latest/developerguide/msk-password.html for details This field is required if your Kafka brokers are accessed over the Internet.
	// Default: none.
	//
	Secret awssecretsmanager.ISecret `field:"optional" json:"secret" yaml:"secret"`
	// An MSK cluster construct.
	ClusterArn *string `field:"required" json:"clusterArn" yaml:"clusterArn"`
}

